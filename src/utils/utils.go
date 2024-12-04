package utils

import (
    "os"
)

func Check(e error) {
    if e != nil {
        panic(e)
    }
}

func ReadFile(file string) string {
    data, err := os.ReadFile(file) 
    Check(err)
    return string(data)
}


func Abs(value int) int {
    if value < 0 {
        return -value
    }
    return value
}
