package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

type State struct {
	ParseNext func(byte, *State) error
	Sum       int
	FactorA   int
	DigitStr  string
	ExprCount int
	Active    bool
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	state := &State{
		ParseNext: parseStart,
		Sum:       0,
		Active:    true,
	}

	for {
		char, readErr := reader.ReadByte()
		if readErr == io.EOF {
			break
		}
		parseErr := state.ParseNext(char, state)
		if parseErr != nil {
			state.ParseNext = parseStart
			state.FactorA = 0
			state.DigitStr = ""
			// fmt.Println(parseErr)
		}
	}

	fmt.Println("Parsed", state.ExprCount, "expressions. Sum is ", state.Sum)
}

func parseStart(char byte, state *State) error {
	if char == 'm' {
		state.ParseNext = parseU
		return nil
	}

	if char == 'd' {
		state.ParseNext = parseO
		return nil
	}
	return fmt.Errorf("expected m, got" + string(char))
}

///////// Parse mul instructions

func parseU(char byte, state *State) error {
	if char == 'u' {
		state.ParseNext = parseL
		return nil
	}
	return fmt.Errorf("expected u, got" + string(char))
}

func parseL(char byte, state *State) error {
	if char == 'l' {
		state.ParseNext = parseOpenParen
		return nil
	}
	return fmt.Errorf("expected l, got" + string(char))
}

func parseOpenParen(char byte, state *State) error {
	if char == '(' {
		state.ParseNext = parseFactor
		return nil
	}
	return fmt.Errorf("expected (, got" + string(char))
}

func parseFactor(char byte, state *State) error {
	if char >= '0' && char <= '9' {
		state.DigitStr += string(char)
		fmt.Println(state.DigitStr)
		return nil
	}

	// end of first factor
	if char == ',' && state.FactorA == 0 && state.DigitStr != "" {
		factor, err := strconv.Atoi(state.DigitStr)
		if err != nil {
			panic("error parsing factor" + state.DigitStr)
		}
		state.FactorA = factor
		state.DigitStr = ""
		return nil
	}

	// end of second factor
	if char == ')' && state.FactorA != 0 && state.DigitStr != "" {
		factor, err := strconv.Atoi(state.DigitStr)
		if err != nil {
			panic("error parsing factor" + state.DigitStr)
		}

		// add to sum and reset
		if state.Active {
			state.Sum += (state.FactorA * factor)
			state.ExprCount++
		}
		state.ParseNext = parseStart
		state.FactorA = 0
		state.DigitStr = ""
		return nil
	}

	fmt.Println("factor parse failed", string(char))
	fmt.Printf("%+v\n", *state)
	fmt.Println("-------------")

	return fmt.Errorf("factor parse failed")
}

// /////// Parse do / don't instructions
func parseO(char byte, state *State) error {
	if char == 'o' {
		state.ParseNext = parseDoParenOrDontN
		return nil
	}
	return fmt.Errorf("expected o, got" + string(char))
}

func parseDoParenOrDontN(char byte, state *State) error {
	if char == '(' {
		state.ParseNext = parseDoCloseParen
		return nil
	}

	if char == 'n' {
		state.ParseNext = parseApost
		return nil
	}
	return fmt.Errorf("expected ( or n, got" + string(char))
}

func parseDoCloseParen(char byte, state *State) error {
	if char == ')' {
		state.Active = true
		state.ParseNext = parseStart
		return nil
	}
	return fmt.Errorf("expected ), got" + string(char))
}

func parseApost(char byte, state *State) error {
	if char == '\'' {
		state.ParseNext = parseT
		return nil
	}
	return fmt.Errorf("expected ', got" + string(char))
}

func parseT(char byte, state *State) error {
	if char == 't' {
		state.ParseNext = parseDontOpenParen
		return nil
	}
	return fmt.Errorf("expected ', got" + string(char))
}

func parseDontOpenParen(char byte, state *State) error {
	if char == '(' {
		state.ParseNext = parseDontClosedParen
		return nil
	}
	return fmt.Errorf("expected (, got" + string(char))
}

func parseDontClosedParen(char byte, state *State) error {
	if char == ')' {
		state.Active = false
		state.ParseNext = parseStart
		return nil
	}
	return fmt.Errorf("expected ), got" + string(char))
}
