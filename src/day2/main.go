package main

import (
	"aoc-24/src/utils"
	"bufio"
	"fmt"
    "strings"
	// "fmt"
	"os"
	// "strings"
	// "bufio"
	"strconv"
)


type Order int

const (
    Ascending Order = iota
    Decreasing
)

type Report struct {
    levels []int
    unsafeIdx []int
    order Order
}

func order(levels []int) Order {
    if levels[0] > levels[len(levels) - 1] {
        return Decreasing
    }
    return Ascending
}

func isSafe(report Report) bool {
    if report.order == Ascending {
        for i := 1; i < len(report.levels); i++ {
            gap := report.levels[i] - report.levels[i - 1]
            if gap > 3 || gap < 1 {
                report.unsafeIdx = append(report.unsafeIdx, i - 1)
                report.unsafeIdx = append(report.unsafeIdx, i)
                return false
            }
        }
    } else {
        for i := 0; i < len(report.levels) - 1; i++ {
            gap := report.levels[i] - report.levels[i + 1]
            if gap > 3 || gap < 1 {
                report.unsafeIdx = append(report.unsafeIdx, i)
                report.unsafeIdx = append(report.unsafeIdx, i + 1)
                return false
            }
        }
    }
    return true 
}

func solvePartOne(reports []Report) int {
    count := 0
    for _, report := range reports {
        if isSafe(report) {
            count++
        }
    }
    return count
}

func solvePartTwo(reports []Report) int {
    count := 0

    for _, report := range reports {
        tempReport := Report{copy(report.levels), report.unsafeIdx, report.order}
        copy() 
    }
    return count++
}

func parse(file *os.File) []Report {
    var reports []Report
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := strings.TrimSpace(scanner.Text())
        if line == "" {
            return reports
        }
        nums := strings.Fields(line)
        var levels []int
        for _, n := range nums {
            level, err := strconv.Atoi(n)
            utils.Check(err) 
            levels = append(levels, level)
        }
        reports = append(reports, Report{levels, []int{}, order(levels)})
    }
    return reports
}

func main() {
    file, err := os.Open("input.txt")
    utils.Check(err)
    defer file.Close()
    reports := parse(file)
    numOfSafeReportsPartOne := solvePartOne(reports)
    fmt.Println(numOfSafeReports)
    numOfSafeReportsPartTwo := solvePartTwo(reports)
    totalOfSageReports := numOfSafeReportsPartOne + numOfSafeReportsPartTwo
    fmt.Println(totalOfSageReports)
}
