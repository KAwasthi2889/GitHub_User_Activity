package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
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

	for _, event := range events {
		switch event.Type {
		case "WatchEvent":
			fmt.Println("User starred", event.Repo.Name)
		case "ForkEvent":
			fmt.Println("User forked", event.Repo.Name)
		case "PublicEvent":
			fmt.Println("User open sourced", event.Repo.Name)
		case "GollumEvent":
			fmt.Println("User created/updated the wiki in", event.Repo.Name)
		case "SponsershipEvent":
			fmt.Println("User sponsored", event.Repo.Name)
		case "PushEvent":
			fmt.Println("User pushed", event.Payload.Size, "commits to", event.Repo.Name)
		case "CommitCommentEvent":
			fmt.Println("User", event.Payload.Action, "a commit comment in", event.Repo.Name)
		case "PullRequestEvent":
			fmt.Println("User", event.Payload.Action, "pull request #"+strconv.Itoa(event.Payload.Number)+" on", event.Repo.Name, "for", event.Payload.Reason)
		case "PullRequestReviewEvent":
			fmt.Println("User", event.Payload.Action, "a pull request review in", event.Repo.Name)
		case "PullRequestReviewCommentEvent":
			fmt.Println("User", event.Payload.Action, "a comment on a pull request review in", event.Repo.Name)
		case "PullRequestReviewThreadEvent":
			fmt.Println("User", event.Payload.Action, "a thread on a pull request review in", event.Repo.Name)
		case "IssuesEvent":
			fmt.Println("User", event.Payload.Action, "an issue in", event.Repo.Name)
		case "IssueCommentEvent":
			fmt.Println("User", event.Payload.Action, "a issue/pull request comment in", event.Repo.Name)
		case "ReleaseEvent":
			fmt.Println("User", event.Payload.Action, "a release in", event.Repo.Name)
		case "MemberEvent":
			fmt.Println("User", event.Payload.Action, "a collaborator to", event.Repo.Name)
		case "CreateEvent":
			if event.Payload.Created_type == "repository" {
				fmt.Println("User created", event.Repo.Name)
			} else {
				fmt.Println("User created", event.Payload.Created_type, "branch in", event.Repo.Name)
			}
		case "DeleteEvent":
			if event.Payload.Created_type == "repository" {
				fmt.Println("User deleted", event.Repo.Name)
			} else {
				fmt.Println("User deleted", event.Payload.Created_type, "branch in", event.Repo.Name)
			}
		default:
			fmt.Println("Undefined event occured:", event.Type)
		}
	}
}
