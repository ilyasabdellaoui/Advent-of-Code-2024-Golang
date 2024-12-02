package main

import (
	"fmt"
	"io"
	"net/http"
)

var client = http.Client{}

func main() {
	// Cookies to send
	cookies := []http.Cookie{
		{Name: "_ga", Value: "GA1.2.1941783799.1733146083"},
		{Name: "_gid", Value: "GA1.2.550765448.1733146083"},
		{Name: "session", Value: "53616c7465645f5f7c6acbe72772bb03756f86fa8be848c20c949b547a54b8ae70b494dc8120e4f3af92638fd24023056f710a97bf2d46f984ef4fb585398ecc"},
		{Name: "_ga_MHSNPJKWC7", Value: "GS1.2.1733146084.1.1.1733146186.0.0.0"},
	}

	// Create the request
	req, err := http.NewRequest("GET", "https://adventofcode.com/2024/day/1/input", nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Add cookies to the request
	for _, cookie := range cookies {
		req.AddCookie(&cookie)
	}

	// Send the request with cookies
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error fetching the input:", err)
		return
	}
	defer resp.Body.Close() // Ensure the body is closed after reading

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	// Print the response body
	fmt.Println(string(body))
}

