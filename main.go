package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("BOT_TOKEN"))
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60
	updateConfig.AllowedUpdates = []string{"message", "channel_post", "chat_member"}
	updates := bot.GetUpdatesChan(updateConfig)

	for update := range updates {

		if update.ChatMember == nil {
			continue
		}
		msg := tgbotapi.NewMessage(update.ChatMember.Chat.ID, "Welcome "+update.ChatMember.NewChatMember.User.FirstName)
		if _, err := bot.Send(msg); err != nil {
			panic(err)
		}
	} 
}