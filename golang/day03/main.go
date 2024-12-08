package main

import (
        "fmt"
        "regexp"
        "strconv"
        
        "ilyasabdell.me/advent-code/PuzzleReader"
)

func main() {
    url := "https://adventofcode.com/2024/day/3/input"
    
    puzzleInput := PuzzleReader.ReadPuzzle(url)
    sumMults := 0
    
    matches := extracteMul(puzzleInput)
    
    for _,match := range matches {
        mult := 1
        nbrs := extractFactors(match)
        for _,nbr := range nbrs {
            nb,_ := strconv.Atoi(nbr)
            mult *= int(nb)
        }
        sumMults += mult
    }
    fmt.Println(sumMults)
}

func extracteMul(input string) []string {
    re := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
    matches := re.FindAllString(input, -1)
    return matches
}

func extractFactors(match string) []string {
    re := regexp.MustCompile(`\d{1,3}`)
    nbrs := re.FindAllString(match, -1)
    return nbrs
}