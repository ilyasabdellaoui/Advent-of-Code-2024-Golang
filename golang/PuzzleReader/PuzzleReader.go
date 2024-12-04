package PuzzleReader

import (
        "fmt"
        "net/http"
        "io"
        "os"
          
        "github.com/joho/godotenv"
)

var client = http.Client{}
func ReadPuzzle(url string) string {
    err := godotenv.Load(".env")
    if err != nil {
      fmt.Println("Ops, can't load env file: ", err)
      return ""
    }
    
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        fmt.Println("Ops, can't make request: ", err)
        return ""
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
        return ""
    }

    defer resp.Body.Close()
    
    bodyResp, err := io.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("Ops, can't read the puzzle body: ", err)
        return ""
    } 
    
    puzzleInput := string(bodyResp)
    return puzzleInput
}