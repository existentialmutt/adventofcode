package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// read a line from standard in
	reader := bufio.NewReader(os.Stdin)
	left := make([]int, 0)
	right := make([]int, 0)
	for i := 0; true; i++ {
		// read the next line
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}

		// parse the line and assign to left and right slices
		line = strings.TrimSuffix(line, "\n")
		parsed := strings.Split(line, "   ")

		a, _ := strconv.Atoi(parsed[0])
		left = append(left, a)

		b, _ := strconv.Atoi(parsed[1])
		right = append(right, b)
	}

	// sort the slices
	sort.Ints(left)
	sort.Ints(right)

	// sum up the differences
	var sum int
	for i, a := range left {
		b := right[i]
		diff := a - b
		if diff < 0 {
			diff *= -1
		}
		sum += diff
	}

	fmt.Println("Parsed " + strconv.Itoa(len(left)) + " lines. Total: " + strconv.Itoa(sum))
}
