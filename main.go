package main

import (
	"log"
	"os"

	"github.com/eyoelmeles/hello-plus-telegram-bot/utils"
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

		var msg tgbotapi.MessageConfig
		if update.ChatMember != nil {
			if update.ChatMember.NewChatMember.HasLeft() || update.ChatMember.NewChatMember.Status == "kicked" {
				msg = tgbotapi.NewMessage(update.ChatMember.Chat.ID, "GoodBye "+update.ChatMember.NewChatMember.User.FirstName)
			} else {
				msg = tgbotapi.NewMessage(update.ChatMember.Chat.ID, "Welcome "+update.ChatMember.NewChatMember.User.FirstName)
			}
		} else if update.ChannelPost != nil {
			msg = tgbotapi.NewMessage(update.ChannelPost.Chat.ID, utils.Profanity(update.ChannelPost.Text))
			bot.Request(tgbotapi.NewDeleteMessage(update.ChannelPost.Chat.ID, update.ChannelPost.MessageID))
		} else if update.Message != nil {
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, utils.Profanity(update.Message.Text))
			bot.Request(tgbotapi.NewDeleteMessage(update.Message.Chat.ID, update.Message.MessageID))
			} else {
			continue
		}

		if _, err := bot.Send(msg); err != nil {
			panic(err)
		}
	}
}