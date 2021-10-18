package bitbucket

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

	if body.CommitStatus.Type == "build" {
		notifyBuildStatus(body, chatId64)
		c.Status(http.StatusOK)
		return
	}

	c.Status(http.StatusBadRequest)
}

func notifyBuildStatus(webhook Body, chatId int64) {
	text := buildStatusText(BuildStatusMessage{
		webhook.Repository.FullName,
		webhook.CommitStatus.Name,
		webhook.CommitStatus.State,
		webhook.CommitStatus.Commit.Author.User.DisplayName,
		webhook.CommitStatus.URL,
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
