package main

import (
	"log"

	ggs "github.com/THANHPP/Work_DeltaTeam/Chatbot/Telegram/handler/googlesheett"
	tb "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	// Set up a telegram bot
	bot, err := tb.NewBotAPI("1032135930:AAG7_bSPwq8ih4rc3cBHsz5UDyYd-q9gr8g")
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
				msg.Text = "/createShortLink [range]"

			case "createShortLink":
				arg := update.Message.CommandArguments()
				if len(arg) < 1 {
					msg.Text = "nothing received"
				} else {
					dataSet := ggs.GetDataFromRage(arg)
					for _, data := range dataSet {
						msg.Text = msg.Text + data + "\n"
					}
				}

			default:
				msg.Text = "WTF :-?"
			}

			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}
		}

	}
}
