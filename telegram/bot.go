package telegram

import (
	"log"
	"strconv"

	"github.com/diogenesc/telegram-webhook/env"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var bot *tgbotapi.BotAPI

func initBot(botToken string) *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug, _ = strconv.ParseBool(env.GetEnv("TELEGRAM_BOT_DEBUG", "false"))

	log.Printf("Authorized on account %s", bot.Self.UserName)

	return bot
}

func SendMessage(chatId int64, text string, botToken string) {
	msg := tgbotapi.NewMessage(chatId, text)

	msg.ParseMode = "markdown"
	msg.DisableWebPagePreview = true

	bot = initBot(botToken)

	bot.Send(msg)
}
