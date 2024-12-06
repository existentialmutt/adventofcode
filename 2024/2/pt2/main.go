/*
	Test reports, plus all possible variations with one element removed

	https://adventofcode.com/2024/day/2#part2
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

		line = strings.TrimSuffix(line, "\n")
		levelStrs := strings.Split(line, " ")
		report := make([]int, 0)
		for _, str := range levelStrs {
			level, _ := strconv.Atoi(str)
			report = append(report, level)
		}
		if isReportSafeWithProblemDampener(report) {
			safeReports++
		}
	}

	fmt.Println("Parsed " + strconv.Itoa(lineCount) + " lines. Total Safe Reports: " + strconv.Itoa(safeReports))
}

func isReportSafeWithProblemDampener(report []int) bool {
	fmt.Println("-----------------")
	fmt.Println("base report: ", report)
	if isReportSafe(report) {
		return true
	}

	for i := 0; i < len(report); i++ {
		fmt.Println("++")
		variation := make([]int, 0)
		variation = append(variation, report[:i]...)
		variation = append(variation, report[i+1:]...)
		fmt.Println("report", report)
		fmt.Println("variation: ", variation)
		if isReportSafe(variation) {
			return true
		}
	}
	return false
}

func isReportSafe(report []int) bool {
	isSafe := true

	// iterate through each level
	// and determine if the report is safe
	reportDirection := 0 // -1 for decreasing, 1 for increasing
	lastLevel := 0
	for i, level := range report {
		if i == 0 {
			lastLevel = level
			continue
		}
		// fmt.Println("*")
		// establish report direction and ensure it is
		// increasing or decreasing
		if i == 1 {
			reportDirection = cmp.Compare[int](level, lastLevel)
			if reportDirection == 0 {
				isSafe = false
				fmt.Println("Unsafe reportDirection", lastLevel, level)
				break
			}
		}

		// ensure we're still moving in the same direction
		stepDirection := cmp.Compare[int](level, lastLevel)
		// fmt.Println("stepDirection: ", stepDirection)
		if stepDirection != reportDirection {
			fmt.Println("Unsafe stepDirection!", lastLevel, level, stepDirection, reportDirection)
			isSafe = false
			break
		}

		// ensure we moved by between 1 and 3
		delta := (level - lastLevel) * reportDirection
		if delta < 1 || delta > 3 {
			fmt.Println("Unsafe delta!", lastLevel, level, delta)
			isSafe = false
			break
		}
		lastLevel = level
	}
	if isSafe {
		fmt.Println("Safe!")
	}
	return isSafe
}
