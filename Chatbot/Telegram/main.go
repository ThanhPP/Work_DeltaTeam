package main

import (
	"fmt"
	"log"

	cf "github.com/THANHPP/Work_DeltaTeam/Chatbot/Telegram/config"
	name "github.com/THANHPP/Work_DeltaTeam/Chatbot/Telegram/handler/namedotcom"
	tb "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	// Set up a telegram bot
	teleAPIKey, err := cf.GetEnvKey("TELEGRAMBOTAPIKEY")
	bot, err := tb.NewBotAPI(teleAPIKey)
	if err != nil {
		panic(err)
	}
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tb.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		//handling command
		if update.Message.IsCommand() {
			msg := tb.NewMessage(update.Message.Chat.ID, "")
			switch update.Message.Command() {
			case "help":
				msg.Text = "/createshortlink [range]"

			case "createshortlink":
				arg := update.Message.CommandArguments()
				if len(arg) < 1 {
					msg.Text = "nothing received"
				} else {
					msg1 := tb.NewMessage(update.Message.Chat.ID, "Start forwarding... Please wait")
					bot.Send(msg1)
					forwardResult, successForwardCount, errorForwardCount := name.ForwardLinks(arg)
					for _, link := range forwardResult {
						msg.Text = msg.Text + link + "\n"
					}
					msg.Text = msg.Text + fmt.Sprintf("\nSuccess count : %d\nError count : %d\n", successForwardCount, errorForwardCount)
				}

			default:
				msg.Text = "WTF :-?"
				bot.Send(tb.NewMessage(update.Message.Chat.ID, "test"))
			}

			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}
		}

	}
}
