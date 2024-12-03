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

func solve(lefts []int, rights []int) int {
    sum := 0
    var left, right int
    sort.Sort(sort.IntSlice(lefts))
    sort.Sort(sort.IntSlice(rights))
    for len(lefts) > 0 && len(rights) > 0 {
        left, lefts = lefts[0], lefts[1:]
        right, rights = rights[0], rights[1:]
        sum += abs(left - right) 
    }
    return sum
}


func main() {
    data := read("input.txt")
    lefts, rights := parse(data)
    answer := solve(lefts, rights)
    
    // Part 1
    fmt.Println(answer)  
}
