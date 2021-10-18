package main

import (
	"github.com/gin-gonic/gin"

	"github.com/joho/godotenv"

	"github.com/diogenesc/telegram-webhook/env"
	"github.com/diogenesc/telegram-webhook/providers/bitbucket"
	"github.com/diogenesc/telegram-webhook/providers/github"
)

func main() {
	godotenv.Load()

	router := gin.Default()
	router.POST("/bitbucket", bitbucket.Handler)
	router.POST("/github", github.Handler)

	port := env.GetEnv("PORT", "8080")

	router.Run(":" + port)
}
