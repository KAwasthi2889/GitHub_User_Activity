package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		println("Usage: ./github-activity <username>")
		return
	}
	username := os.Args[1]

	url := "https://api.github.com/users/" + username + "/events"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == 404 {
		fmt.Println("Invalid Username")
		return
	}
	if resp.StatusCode != 200 {
		fmt.Println("Server Busy, Please try again later")
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var events []Event
	err = json.Unmarshal(body, &events)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if len(events) == 0 {
		fmt.Println("No public activity by", username, "within 90 days.")
		return
	}

	fmt.Println("Public activities by", username, "within 90 days are:-")

	for i, event := range events {
		fmt.Printf("%02d. %s\n", i+1, Activity(event))
	}
}
