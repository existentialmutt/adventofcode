/*
	https://adventofcode.com/2024/day/2
*/

package main

import (
	"bufio"
	"cmp"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	// read a line from standard in
	reader := bufio.NewReader(os.Stdin)
	safeReports := 0
	lineCount := 0
	for ; true; lineCount++ {
		// read the next line
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}

		report := strings.TrimSuffix(line, "\n")
		if isReportSafe(report) {
			safeReports++
		}
	}

	fmt.Println("Parsed " + strconv.Itoa(lineCount) + " lines. Total Safe Reports: " + strconv.Itoa(safeReports))
}

func isReportSafe(report string) bool {
	fmt.Println("------------")
	fmt.Println(report)

	isSafe := true

	// iterate through each level
	// and determine if the report is safe
	reportDirection := 0 // -1 for decreasing, 1 for increasing
	lastLevel := 0
	for i, levelStr := range strings.Split(report, " ") {
		level, _ := strconv.Atoi(levelStr)
		if i == 0 {
			lastLevel = level
			continue
		}

		// establish report direction and ensure it is
		// increasing or decreasing
		if i == 1 {
			reportDirection = cmp.Compare[int](level, lastLevel)
			fmt.Println("reportDirection: " + strconv.Itoa(reportDirection))
			if reportDirection == 0 {
				isSafe = false
				fmt.Println("Report Unsafe: reportDirection")
				break
			}
		}

		// ensure we're still moving in the same direction
		stepDirection := cmp.Compare[int](level, lastLevel)
		fmt.Println("*")
		fmt.Println("level: " + strconv.Itoa(level))
		fmt.Println("lastLevel: " + strconv.Itoa(lastLevel))
		fmt.Println("stepDirection: " + strconv.Itoa(stepDirection))
		if stepDirection != reportDirection {
			isSafe = false
			fmt.Println("Report Unsafe: stepDirection")
			break
		}

		// ensure we moved by between 1 and 3
		delta := (level - lastLevel) * reportDirection
		fmt.Println("delta: " + strconv.Itoa(delta))
		if delta < 1 || delta > 3 {
			isSafe = false
			fmt.Println("Report Unsafe: delta")
			break
		}
		lastLevel = level
		fmt.Println("Report Safe!")
	}
	return isSafe
}
