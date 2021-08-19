package telegram

import (
	"log"
	"strconv"

	"github.com/diogenesc/telegram-webhook/env"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func InitBot(botToken string) *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug, _ = strconv.ParseBool(env.GetEnv("TELEGRAM_BOT_DEBUG", "false"))

	log.Printf("Authorized on account %s", bot.Self.UserName)

	return bot
}
