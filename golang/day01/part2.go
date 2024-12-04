package main

import (
        "fmt"
        "strings"
        "strconv"
        
        "ilyasabdell.me/advent-code/PuzzleReader"
)

func main(){
    url := "https://adventofcode.com/2024/day/1/input"
    puzzleInput := PuzzleReader.ReadPuzzle(url)
    
    var col1 []int
    var col2 []int
    
    lines := strings.Split(puzzleInput,"\n")
    for _, line := range lines {
        col := strings.Fields(line)
        
        if len(col) != 2 {
            continue
        }
        
        colVal1, err1 := strconv.Atoi(col[0])
        colVal2, err2 := strconv.Atoi(col[1])
        
        if err1 != nil || err2 != nil{
            continue
        }
        
        col1 = append(col1, colVal1)
        col2 = append(col2, colVal2)
    }
    
    var simScore int
    for _, val := range(col1) {
        simScore += val*countOcc(col2, val)
    }
    
    fmt.Println(simScore)
}

func countOcc(slice []int, val int) int {
    var count int
    for _, v := range slice {
        if v == val {
            count += 1
        }
    }
    return count
}
