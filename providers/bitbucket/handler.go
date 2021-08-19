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

func BitbucketHandler(c *gin.Context) {
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

	chat_id_64, _ := strconv.ParseInt(chatId, 10, 64)

	bot = telegram.InitBot(botToken)

	if body.CommitStatus.Type == "build" {
		notifyBuildStatus(body, chat_id_64)
		c.Status(http.StatusOK)
		return
	}

	c.Status(http.StatusBadRequest)
}

func notifyBuildStatus(webhook Body, chatId int64) {
	text := buildStatusText(BuildStatusMessage{webhook.CommitStatus.Name, webhook.CommitStatus.State, webhook.CommitStatus.URL})

	msg := tgbotapi.NewMessage(chatId, text)

	msg.ParseMode = "markdown"
	msg.DisableWebPagePreview = true

	bot.Send(msg)
}

func buildStatusText(message BuildStatusMessage) string {
	var text string
	if message.Title != "" {
		text += fmt.Sprintf("*%s*\n\n", message.Title)
	}
	if message.State != "" {
		text += fmt.Sprintf("*State:* %s\n\n", message.State)
	}
	if message.URL != "" {
		text += fmt.Sprintf("[More information here](%s)", message.URL)
	}

	return text
}
