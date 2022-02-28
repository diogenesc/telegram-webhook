package bitbucket

import "fmt"

func buildPipelineText(message BuildPipelineMessage) string {
	var text string
	if message.RepositoryFullName != "" {
		text += fmt.Sprintf("*Repository:* %s\n\n", message.RepositoryFullName)
	}
	if message.Title != "" {
		text += fmt.Sprintf("*%s*\n\n", message.Title)
	}
	if message.State != "" {
		var emote string
		switch message.State {
		case "INPROGRESS":
			emote = "🔄"
		case "SUCCESSFUL":
			emote = "✅"
		case "FAILED":
			emote = "❌"
		case "STOPPED":
			emote = "🔴"
		}
		text += fmt.Sprintf("*State:* %s %s\n\n", message.State, emote)
	}
	if message.Author != "" {
		text += fmt.Sprintf("*Author:* %s\n\n", message.Author)
	}
	if message.URL != "" {
		text += fmt.Sprintf("[More information here](%s)", message.URL)
	}

	return text
}

func buildPullRequestText(message BuildPullRequestMessage) string {
	var text string
	if message.RepositoryFullName != "" {
		text += fmt.Sprintf("*Repository:* %s\n\n", message.RepositoryFullName)
	}
	text += fmt.Sprintf("*Pull Request*\n\n")
	if message.Title != "" {
		text += fmt.Sprintf("*Title:* %s\n\n", message.Title)
	}
	if message.Source != "" && message.Destination != "" {
		text += fmt.Sprintf("*%s*  ➡  *%s*\n\n", message.Source, message.Destination)
	}
	if message.State != "" {
		var emote string
		switch message.State {
		case "OPEN":
			emote = "🟢"
		case "MERGED":
			emote = "✅"
		case "DECLINED":
			emote = "❌"
		}
		text += fmt.Sprintf("*State:* %s %s\n\n", message.State, emote)
	}
	if message.URL != "" {
		text += fmt.Sprintf("[More information here](%s)", message.URL)
	}

	return text
}
