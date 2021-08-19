package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/joho/godotenv"

	"github.com/diogenesc/telegram-webhook/env"
	"github.com/diogenesc/telegram-webhook/providers/bitbucket"
	"github.com/diogenesc/telegram-webhook/telegram"
)

func main() {
	godotenv.Load()

	router := gin.Default()
	router.POST("/bitbucket", bitbucketHandler)

	port := env.GetEnv("HOST_PORT", "8080")

	router.Run(":" + port)
}

func bitbucketHandler(c *gin.Context) {
	var body bitbucket.Body
	c.BindJSON(&body)

	bot_token := c.Query("bot_token")
	chat_id := c.Query("chat_id")

	if bot_token == "" || chat_id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Parameter missing",
		})
		return
	}

	chat_id_64, _ := strconv.ParseInt(chat_id, 10, 64)

	telegram.NotifyBuildStatus(body, chat_id_64, bot_token)

	c.Status(http.StatusOK)
}
