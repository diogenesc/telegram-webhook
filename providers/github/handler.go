package github

import (
	"net/http"
	"strconv"

	"github.com/diogenesc/telegram-webhook/telegram"
	"github.com/gin-gonic/gin"
)

type RequestQueryString struct {
	BotToken string `form:"bot_token" binding:"required"`
	ChatId   string `form:"chat_id" binding:"required"`
}

func Handler(c *gin.Context) {
	var query RequestQueryString
	var body Body

	c.Bind(&query)
	c.BindJSON(&body)

	chatId64, _ := strconv.ParseInt(query.ChatId, 10, 64)

	if (body.WorkflowRun != WorkflowRun{}) {
		notifyBuildStatus(body, chatId64, query.BotToken)
		c.Status(http.StatusOK)
		return
	}

	c.Status(http.StatusBadRequest)
}

func notifyBuildStatus(webhook Body, chatId int64, botToken string) {
	text := buildStatusText(BuildStatusMessage{
		webhook.Repository.FullName,
		webhook.WorkflowRun.Name,
		webhook.WorkflowRun.Status,
		webhook.WorkflowRun.Conclusion,
		webhook.Sender.Login,
		webhook.WorkflowRun.URL,
	})

	telegram.SendMessage(chatId, text, botToken)
}
