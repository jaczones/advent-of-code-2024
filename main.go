package main

import (
	"advent-of-code/day3"
	"fmt"
	"io"
	"net/http"
	"os"
)

func FetchWebPage(url, sessionCookie string) (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	// Add the session cookie to the request
	req.Header.Set("Cookie", fmt.Sprintf("session=%s", sessionCookie))

	// Make the request
	response, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	// Read the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func main() {
	sessionCookie := os.Getenv("SESSION_COOKIE")
	url := "https://adventofcode.com/2024/day/3/input"

	content, err := FetchWebPage(url, sessionCookie)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	day3.Day3Part2(content)
}
