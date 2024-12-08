package main

import (
        "fmt"
        "regexp"
        "strconv"
        "strings"
        
        "ilyasabdell.me/advent-code/PuzzleReader"
)

func main() {
    url := "https://adventofcode.com/2024/day/3/input"
    
    puzzleInput := PuzzleReader.ReadPuzzle(url)
    sumMults := 0
    
    matches := extractMul(puzzleInput)
    
    for _,match := range matches {
        mult := 1
        nbrs := extractFactors(match)
        for _,nbr := range nbrs {
            nb,_ := strconv.Atoi(nbr)
            mult *= int(nb)
        }
        sumMults += mult
    }
    fmt.Println("Part 1 solution : ", sumMults)
    
    // part2: extracting only do multiplicatons
    doMultsSum := 0
    doMultiplications := extrcatDoMuls(puzzleInput)
    
    doMatches := extractMul(doMultiplications)
    for _,doMatch := range doMatches {
        doMult := 1
        doNbrs := extractFactors(doMatch)
        for _,doNbr := range doNbrs {
            doNb,_ := strconv.Atoi(doNbr)
            doMult *= int(doNb)
        }
        doMultsSum += doMult
    }
    fmt.Println("Part 2 solution : ", doMultsSum)
}

func extractMul(input string) []string {
    re := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
    matches := re.FindAllString(input, -1)
    return matches
}

func extractFactors(match string) []string {
    re := regexp.MustCompile(`\d{1,3}`)
    nbrs := re.FindAllString(match, -1)
    return nbrs
}

func extrcatDoMuls(input string) string {
  	re := regexp.MustCompile(`do\(\)|don't\(\)|mul\(\d{1,3},\d{1,3}\)`)
    matches := re.FindAllString(input, -1)
    
  	mulEnabled := true
  	var result []string
  
  	for _, match := range matches {
  		if match == "do()" {
  			mulEnabled = true
  		} else if match == "don't()" {
  			mulEnabled = false
  		} else if mulEnabled && match[:4] == "mul(" {
  			result = append(result, match)
  		}
  	}
    return strings.Join(result, " ")
}