package main

import (
    "fmt"
    "aoc-24/src/utils"
    "strings"
    "bufio"
    // "strconv"
)

type Order int

type Reports struct {
    levels []int
    order Order
}

const (
    Ascending Order = iota
    Decreasing
)

func parse(data string)  {
    lists := strings.Split(data, "\n")
    fmt.Println(lists)
    // for _, list := range lists {
    //     if utils.Abs()
    // }

}

func main() {
    data := utils.ReadFile("example.txt")
    parse(data)
}
