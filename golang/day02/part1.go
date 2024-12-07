package main 

import (
        "fmt"
        "strings"
        "strconv"
        
        "ilyasabdell.me/advent-code/PuzzleReader"
)

func main() {
    url := "https://adventofcode.com/2024/day/2/input"
    puzzleInput := PuzzleReader.ReadPuzzle(url)
    
    lines := strings.Split(puzzleInput, "\n")
    
    safeCount := 0
    for _,line := range lines {
        if isSafeReport(line) {
            safeCount++
        }
    }
    fmt.Println(safeCount)
}

func isSafeReport(line string) bool {
    fields := strings.Fields(line)
    if len(fields) == 0 {
        return false
    }
    if isDecreasing(fields) || isIncreasing(fields) {
         return true   
    }
    return false
}

func isDecreasing(field []string) bool {
    for idx := 0; idx < len(field)-1; idx++ {
        num1, err1 := strconv.Atoi(field[idx])
        num2, err2 := strconv.Atoi(field[idx+1])
        if err1 != nil || err2 != nil {
            return false
        }
        // checking both conditions
        if num1 < num2 || num1 - num2 > 3 || num1 - num2 < 1 {
            return false
        }
    }
    return true
}

func isIncreasing(field []string) bool {
    for idx := 0; idx < len(field)-1; idx++ {
        num1, err1 := strconv.Atoi(field[idx])
        num2, err2 := strconv.Atoi(field[idx+1])
        if err1 != nil || err2 != nil {
            return false
        }
        if num1 > num2 || num2 - num1 > 3 || num2 - num1 < 1 {
            return false
        }
        
    }
    return true
}