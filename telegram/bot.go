package telegram

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/diogenesc/telegram-webhook/providers/bitbucket"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func initBot(botToken string) *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug, _ = strconv.ParseBool(os.Getenv("TELEGRAM_BOT_DEBUG"))

	log.Printf("Authorized on account %s", bot.Self.UserName)

	return bot
}

func NotifyBuildStatus(webhook bitbucket.Body, chatId int64, botToken string) {
	bot := initBot(botToken)
	var text string
	text += fmt.Sprintf("*%s*\n\n", webhook.CommitStatus.Name)
	text += fmt.Sprintf("*State:* %s\n\n", webhook.CommitStatus.State)
	text += fmt.Sprintf("[More information here](%s)", webhook.CommitStatus.URL)
	msg := tgbotapi.NewMessage(chatId, text)

	msg.ParseMode = "markdown"
	msg.DisableWebPagePreview = true

	bot.Send(msg)
}
