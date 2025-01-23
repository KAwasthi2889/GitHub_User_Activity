package main

import "fmt"

func Activity(event Event) string {
	switch event.Type {
	case "WatchEvent":
		return fmt.Sprintf("User starred %s", event.Repo.Name)
	case "ForkEvent":
		return fmt.Sprintf("User forked %s", event.Repo.Name)
	case "PublicEvent":
		return fmt.Sprintf("User open sourced %s", event.Repo.Name)
	case "GollumEvent":
		return fmt.Sprintf("User created/updated the wiki in %s", event.Repo.Name)
	case "SponsershipEvent":
		return fmt.Sprintf("User sponsored %s", event.Repo.Name)
	case "PushEvent":
		ch := ""
		if event.Payload.Size > 1 {
			ch = "s"
		}
		return fmt.Sprintf("User pushed %d commit%s to %s", event.Payload.Size, ch, event.Repo.Name)
	case "CommitCommentEvent":
		return fmt.Sprintf("User %s a commit comment in %s", event.Payload.Action, event.Repo.Name)
	case "PullRequestEvent":
		return fmt.Sprintf("User %s pull request #%d on %s for %s", event.Payload.Action, event.Payload.Number, event.Repo.Name, event.Payload.Reason)
	case "PullRequestReviewEvent":
		return fmt.Sprintf("User %s a pull request review in %s", event.Payload.Action, event.Repo.Name)
	case "PullRequestReviewCommentEvent":
		return fmt.Sprintf("User %s a comment on a pull request review in %s", event.Payload.Action, event.Repo.Name)
	case "PullRequestReviewThreadEvent":
		return fmt.Sprintf("User %s a thread on a pull request review in %s", event.Payload.Action, event.Repo.Name)
	case "IssuesEvent":
		return fmt.Sprintf("User %s an issue in %s", event.Payload.Action, event.Repo.Name)
	case "IssueCommentEvent":
		return fmt.Sprintf("User %s an issue/pull request comment in %s", event.Payload.Action, event.Repo.Name)
	case "ReleaseEvent":
		return fmt.Sprintf("User %s a release in %s", event.Payload.Action, event.Repo.Name)
	case "MemberEvent":
		return fmt.Sprintf("User %s a collaborator to %s", event.Payload.Action, event.Repo.Name)
	case "CreateEvent":
		if event.Payload.Created_type == "repository" {
			return fmt.Sprintf("User created %s", event.Repo.Name)
		}
		return fmt.Sprintf("User created %s in %s", event.Payload.Created_type, event.Repo.Name)
	case "DeleteEvent":
		if event.Payload.Created_type == "repository" {
			return fmt.Sprintf("User deleted %s", event.Repo.Name)
		}
		return fmt.Sprintf("User deleted %s branch in %s", event.Payload.Created_type, event.Repo.Name)
	default:
		return fmt.Sprintf("Undefined event occured: %s", event.Type)
	}
}
