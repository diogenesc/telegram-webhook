package bitbucket

import (
	"net/http"
	"strconv"
	"strings"

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

	c.BindQuery(&query)
	c.BindJSON(&body)

	chatId64, _ := strconv.ParseInt(query.ChatId, 10, 64)
	event := c.Request.Header.Get("X-Event-Key")
	if strings.HasPrefix(event, "repo:commit_status_") {
		notifyBuildStatus(body, chatId64, query.BotToken)
		c.Status(http.StatusOK)
		return
	}
	if strings.HasPrefix(event, "pullrequest:") {
		notifyPullRequestStatus(body, chatId64, query.BotToken)
		c.Status(http.StatusOK)
		return
	}

	c.Status(http.StatusConflict)
}

func notifyBuildStatus(webhook Body, chatId int64, botToken string) {
	text := buildPipelineText(BuildPipelineMessage{
		webhook.Repository.FullName,
		webhook.CommitStatus.Name,
		webhook.CommitStatus.State,
		webhook.CommitStatus.Commit.Author.User.DisplayName,
		webhook.CommitStatus.URL,
	})

	telegram.SendMessage(chatId, text, botToken)
}

func notifyPullRequestStatus(webhook Body, chatId int64, botToken string) {
	text := buildPullRequestText(BuildPullRequestMessage{
		webhook.Repository.FullName,
		webhook.PullRequest.Title,
		webhook.PullRequest.Source.Branch.Name,
		webhook.PullRequest.Destination.Branch.Name,
		webhook.PullRequest.State,
		webhook.PullRequest.Links.Html.Href,
	})

	telegram.SendMessage(chatId, text, botToken)
}
