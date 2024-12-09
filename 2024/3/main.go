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
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	state := &State{
		ParseNext: parseM,
		Sum:       0,
	}

	for {
		char, readErr := reader.ReadByte()
		if readErr == io.EOF {
			break
		}
		parseErr := state.ParseNext(char, state)
		if parseErr != nil {
			state.ParseNext = parseM
			state.FactorA = 0
			state.DigitStr = ""
			// fmt.Println(parseErr)
		}
	}

	fmt.Println("Parsed", state.ExprCount, "expressions. Sum is ", state.Sum)
}

func parseM(char byte, state *State) error {
	if char == 'm' {
		state.ParseNext = parseU
		return nil
	}
	return fmt.Errorf("expected m, got" + string(char))
}

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
		state.Sum += (state.FactorA * factor)
		state.ExprCount++
		state.ParseNext = parseM
		state.FactorA = 0
		state.DigitStr = ""
		return nil
	}

	fmt.Println("factor parse failed", string(char))
	fmt.Printf("%+v\n", *state)
	fmt.Println("-------------")

	return fmt.Errorf("factor parse failed")
}
