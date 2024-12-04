package main

import (
	"fmt"
	"strconv"
	"strings"
    "sort"
    "aoc-24/src/utils"
)

func parse(data string) ([]int, []int) {
    nums := strings.Fields(string(data))
    var lefts, rights []int
    for i := 0; i < len(nums); i += 2 {
        left, err := strconv.Atoi(nums[i])
        utils.Check(err)
        lefts = append(lefts, left)
        right, err := strconv.Atoi(nums[i + 1])
        utils.Check(err)
        rights = append(rights, right)
    }
    return lefts, rights
}

func solvePartOne(lefts, rights []int) int {
    sum := 0
    sort.Sort(sort.IntSlice(lefts))
    sort.Sort(sort.IntSlice(rights))
    for i := 0; i < len(lefts); i++ {
        sum += utils.Abs(lefts[i] - rights[i])
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
    data := utils.ReadFile("input.txt")
    lefts, rights := parse(data)
    partOneAnswer := solvePartOne(lefts, rights)
    
    // Part 1
    fmt.Println(partOneAnswer)

    //Part 2
    partTwoAnswer := solvePartTwo(lefts, rights)
    fmt.Println(partTwoAnswer)
}
