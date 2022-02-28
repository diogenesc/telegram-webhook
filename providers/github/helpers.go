package github

import "fmt"

func buildStatusText(message BuildStatusMessage) string {
	var text string
	if message.RepositoryFullName != "" {
		text += fmt.Sprintf("*Repository:* %s\n\n", message.RepositoryFullName)
	}
	if message.Title != "" {
		text += fmt.Sprintf("*CI name:* %s\n\n", message.Title)
	}
	if message.Status != "" {
		var emote string
		var statusMessage string
		switch message.Status {
		case "queued":
			statusMessage = message.Status
			emote = "🔄"
		case "completed":
			statusMessage = message.Conclusion
			switch message.Conclusion {
			case "success":
				emote = "✅"
			case "failure":
				emote = "❌"
			}

		}
		text += fmt.Sprintf("*State:* %s %s\n\n", statusMessage, emote)
	}
	if message.Sender != "" {
		text += fmt.Sprintf("*Author:* %s\n\n", message.Sender)
	}
	if message.URL != "" {
		text += fmt.Sprintf("[More information here](%s)", message.URL)
	}

	return text
}
