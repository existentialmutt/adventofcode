/*
	https://adventofcode.com/2024/day/2#part2
  This was my first attempt which turned out to be incorrect.
  It allows the report to keep going after 1 failed level.

  Unfortunately, sometimes we can make the report good by throwing out the previous level,
  And this solution does not account for that

  e.g.

  2 5 8* 6 7

  removing 8 makes the report good, but we don't know that until we're processing 6
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
		report := Report{
			Content:         line,
			CanSkipBadLevel: true,
		}
		if isReportSafe(&report) {
			safeReports++
		}
	}

	fmt.Println("Parsed " + strconv.Itoa(lineCount) + " lines. Total Safe Reports: " + strconv.Itoa(safeReports))
}

type Report struct {
	Content         string
	CanSkipBadLevel bool
}

func isReportSafe(report *Report) bool {
	fmt.Println("++++++++++++++")
	fmt.Println(report.Content)

	isSafe := true

	// iterate through each level
	// and determine if the report is safe
	reportDirection := 0 // -1 for decreasing, 1 for increasing
	reportDirectionSet := false
	lastLevel := 0
	for i, levelStr := range strings.Split(report.Content, " ") {
		level, _ := strconv.Atoi(levelStr)
		if i == 0 {
			lastLevel = level
			continue
		}

		// establish report direction and ensure it is
		// increasing or decreasing
		if !reportDirectionSet {
			reportDirection = cmp.Compare[int](level, lastLevel)
			fmt.Println("reportDirection: " + strconv.Itoa(reportDirection))
			if reportDirection == 0 {
				if report.CanSkipBadLevel {
					report.CanSkipBadLevel = false
					continue
				}
				isSafe = false
				fmt.Println("Report Unsafe: reportDirection")
				break
			}
			reportDirectionSet = true
		}

		// ensure we're still moving in the same direction
		stepDirection := cmp.Compare[int](level, lastLevel)
		fmt.Println("*" + strconv.FormatBool(report.CanSkipBadLevel))
		fmt.Println("level: " + strconv.Itoa(level))
		fmt.Println("lastLevel: " + strconv.Itoa(lastLevel))
		fmt.Println("stepDirection: " + strconv.Itoa(stepDirection))
		if stepDirection != reportDirection {
			if report.CanSkipBadLevel {
				report.CanSkipBadLevel = false
				continue
			}
			isSafe = false
			fmt.Println("Report Unsafe: stepDirection")
			break
		}

		// ensure we moved by between 1 and 3
		delta := (level - lastLevel) * reportDirection
		fmt.Println("delta: " + strconv.Itoa(delta))
		if delta < 1 || delta > 3 {
			if report.CanSkipBadLevel {
				report.CanSkipBadLevel = false
				continue
			}
			isSafe = false
			fmt.Println("Report Unsafe: delta")
			break
		}
		lastLevel = level
	}
	fmt.Println("Report Safe!")
	return isSafe
}
