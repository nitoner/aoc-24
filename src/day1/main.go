package main

import (
	"fmt"
	"strconv"
	"strings"
	"os"
    "sort"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func read(file string) string {
    data, err := os.ReadFile(file) 
    check(err)
    return string(data)
}

func parse(data string) ([]int, []int) {
    nums := strings.Fields(string(data))
    var lefts, rights []int
    for i := 0; i < len(nums); i += 2 {
        left, err := strconv.Atoi(nums[i])
        check(err)
        lefts = append(lefts, left)
        right, err := strconv.Atoi(nums[i + 1])
        check(err)
        rights = append(rights, right)
    }
    return lefts, rights
}

func abs(value int) int {
    if value < 0 {
        return -value
    }
    return value
}

func solvePartOne(lefts, rights []int) int {
    sum := 0
    sort.Sort(sort.IntSlice(lefts))
    sort.Sort(sort.IntSlice(rights))
    for i := 0; i < len(lefts); i++ {
        sum += abs(lefts[i] - rights[i])
    }
    return sum
}

func count(n int, elems []int,) int {
    count := 0
    for _, val := range elems {
        if val == n {
            count++
        } 
    }
    return count
}

func solvePartTwo(lefts, rights []int) int {
    sum := 0
    for i := 0; i < len(lefts); i++ {
        sum += lefts[i] * count(lefts[i], rights)
    }
    return sum
}

func main() {
    data := read("input.txt")
    lefts, rights := parse(data)
    partOneAnswer := solvePartOne(lefts, rights)
    
    // Part 1
    fmt.Println(partOneAnswer)

    //Part 2
    partTwoAnswer := solvePartTwo(lefts, rights)
    fmt.Println(partTwoAnswer)
}
