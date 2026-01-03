package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type GitHubEvent struct {
	Type string `json:"type"`
	Repo struct {
		Name string `json:"name"`
	} `json:"repo"`
	Payload struct {
		Size int `json:"size"`
	} `json:"payload"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: github-activity <username>")
		os.Exit(1)
	}

	username := os.Args[1]
	var events []GitHubEvent

	apiURL := fmt.Sprintf("https://api.github.com/users/%s/events", username)

	resp, err := http.Get(apiURL)

	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}

	switch resp.StatusCode {
	case 200:

	case 404:
		fmt.Printf("User %s was not found!", username)
		return

	case 403:
		fmt.Println("Error: API rate limit exceeded. Try again later.")
		return

	default:
		fmt.Printf("Unexpected error: %s (Status: %d)\n", resp.Status, resp.StatusCode)
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Error reading body", err)
		os.Exit(1)
	}

	// result := string(body)

	// err = os.WriteFile("user.json", []byte(result), 0644)
	err = json.Unmarshal(body, &events)
	if err != nil {
		log.Fatal(err)
	}

	if len(events) == 0 {
		fmt.Println("No recent activity found for this user.")
		return
	}

	for _, event := range events {
		switch event.Type {
		case "PushEvent":
			fmt.Printf("- Pushed %d commits to %s\n", event.Payload.Size, event.Repo.Name)
		case "IssuesEvent":
			fmt.Printf("- Opened a new issue in %s\n", event.Repo.Name)
		case "WatchEvent":
			fmt.Printf("- Starred %s\n", event.Repo.Name)
		default:
			fmt.Printf("- %s in %s\n", event.Type, event.Repo.Name)
		}
	}

}
