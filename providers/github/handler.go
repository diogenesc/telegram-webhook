package github

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/diogenesc/telegram-webhook/telegram"
	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var bot *tgbotapi.BotAPI

func Handler(c *gin.Context) {
	var body Body
	c.BindJSON(&body)

	botToken := c.Query("bot_token")
	chatId := c.Query("chat_id")

	if botToken == "" || chatId == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Parameter missing",
		})
		return
	}

	chatId64, _ := strconv.ParseInt(chatId, 10, 64)

	bot = telegram.InitBot(botToken)

	if (body.WorkflowRun != WorkflowRun{}) {
		notifyBuildStatus(body, chatId64)
		c.Status(http.StatusOK)
		return
	}

	c.Status(http.StatusBadRequest)
}

func notifyBuildStatus(webhook Body, chatId int64) {
	text := buildStatusText(BuildStatusMessage{
		webhook.Repository.FullName,
		webhook.WorkflowRun.Name,
		webhook.WorkflowRun.Status,
		webhook.WorkflowRun.Conclusion,
		webhook.Sender.Login,
		webhook.WorkflowRun.URL,
	})

	msg := tgbotapi.NewMessage(chatId, text)

	msg.ParseMode = "markdown"
	msg.DisableWebPagePreview = true

	bot.Send(msg)
}

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
