package main

import (
	"fmt"
	"log"

	cf "github.com/THANHPP/Work_DeltaTeam/Chatbot/Telegram/config"
	name "github.com/THANHPP/Work_DeltaTeam/Chatbot/Telegram/handler/namedotcom"
	rb "github.com/THANHPP/Work_DeltaTeam/Chatbot/Telegram/handler/rebrandly"
	tb "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	// Set up a telegram bot
	teleAPIKey, err := cf.GetEnvKey("TELEGRAMBOTAPIKEY")
	log.Println(teleAPIKey)
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
					//Forward phase
					msg1 := tb.NewMessage(update.Message.Chat.ID, "Start forwarding... Please wait")
					bot.Send(msg1)

					forwardResult, successForwardCount, errorForwardCount := name.ForwardLinks(arg)

					msg2 := tb.NewMessage(update.Message.Chat.ID, "")
					for _, link := range forwardResult {
						msg2.Text = msg2.Text + link + "\n"
					}
					msg2.Text = msg2.Text + fmt.Sprintf("\nSuccess count : %d\nError count : %d\n", successForwardCount, errorForwardCount)
					bot.Send(msg2)

					//ShortLink phase
					msg3 := tb.NewMessage(update.Message.Chat.ID, "Start shorting... Please wait")
					bot.Send(msg3)

					shortLinkResult, successShortLinkCount, errorShortLinkCount, usedCount := rb.CreateShortLinkRebrandly(arg, forwardResult)

					msg4 := tb.NewMessage(update.Message.Chat.ID, "")
					for _, link := range shortLinkResult {
						msg4.Text = msg4.Text + link + "\n"
					}
					msg4.Text = msg4.Text + fmt.Sprintf("\n Success count : %d\nError count : %d\n", successShortLinkCount, errorShortLinkCount)
					msg4.Text = msg4.Text + fmt.Sprintf("\n Create %+v links with Rebrandly\n%+v links left", usedCount, 500-usedCount)
					bot.Send(msg4)

					msg.Text = "Short link success"
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
