package main

import (
        "fmt"
        "net/http"
        "io"
        "os"
        "strings"
        "strconv"
        "slices"
          
        "github.com/joho/godotenv"
)

var client = http.Client{}
func main(){
    url := "https://adventofcode.com/2024/day/1/input"
    err := godotenv.Load(".env")
    if err != nil {
      fmt.Println("Ops, can't load env file: ", err)
      return
    }
    
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        fmt.Println("Ops, can't make request: ", err)
        return
    }
    
    cookies := []http.Cookie{
      {Name: "session", Value: os.Getenv("SESSION")},
    }
    
    for _, cookie := range cookies {
        req.AddCookie(&cookie)
    }
    
    resp, err := client.Do(req)
    if err != nil {
        fmt.Println("Ops, can't fetch  puzzle input: ", err)
        return
    }
    // Closing the body after reading
    defer resp.Body.Close()
    
    bodyResp, err := io.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("Ops, can't read the puzzle body: ", err)
        return
    } 
    
    puzzleInput := string(bodyResp)
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
    
    var Sum int
    
    for len(col1) != 0 {
        Sum += Abs(slices.Min(col1) - slices.Min(col2))
        col1 = removeElement(col1, slices.Min(col1))
        col2 = removeElement(col2, slices.Min(col2))
    }
    
    fmt.Println(Sum)
}

func Abs(num int) int {
    if num < 0 {
        return -num
    }
    return num
}

func removeElement(slice []int, val int) []int{
    for i, v := range slice {
        if val == v {
            return append(slice[:i], slice[i+1:]...)
        }
    }
    return slice
}