package main

import (
	"fmt"
	"log"
	"time"

	cf "github.com/THANHPP/Work_DeltaTeam/Chatbot/Telegram/config"
	name "github.com/THANHPP/Work_DeltaTeam/Chatbot/Telegram/handler/namedotcom"
	rb "github.com/THANHPP/Work_DeltaTeam/Chatbot/Telegram/handler/rebrandly"
	tb "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	// Set up a telegram bot
	cf.Init()
	config := cf.GetConfig()
	teleAPIKey := config.GetString("TELEGRAMBOTAPIKEY")
	bot, err := tb.NewBotAPI(teleAPIKey)
	if err != nil {
		panic(err)
	}
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tb.NewUpdate(1)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		panic(err)
	}

	for update := range updates {
		//handling command
		if update.Message.IsCommand() {
			msg := tb.NewMessage(update.Message.Chat.ID, "")

			switch update.Message.Command() {
			case "help":
				msg.Text = "/createfwshortlink [range]"

			case "createfwshortlink":
				arg := update.Message.CommandArguments()
				if len(arg) < 1 {
					msg.Text = "createfwshortlink : no argument found"
				} else {
					//FORWARD PHASE
					//	msg 1
					msg1 := tb.NewMessage(update.Message.Chat.ID, "Start forwarding... Please wait")
					bot.Send(msg1)

					//	timer
					createshortlinkTimer := time.Now()

					//	msg 2
					forwardResult, successForwardCount, errorForwardCount := name.ForwardLinks(arg)

					msg2 := tb.NewMessage(update.Message.Chat.ID, "")
					for _, link := range forwardResult {
						msg2.Text = msg2.Text + link + "\n"
					}
					msg2.Text = msg2.Text + fmt.Sprintf("\nSuccess count : %d\nError count : %d\n", successForwardCount, errorForwardCount)
					msg2.DisableWebPagePreview = true

					bot.Send(msg2)

					//SHORTLINK PHASE
					//	msg 3
					msg3 := tb.NewMessage(update.Message.Chat.ID, "Start shorting... Please wait")
					bot.Send(msg3)

					//	msg 4
					shortLinkResult, successShortLinkCount, errorShortLinkCount, usedCount := rb.CreateShortLinkRebrandly(arg, forwardResult)

					msg4 := tb.NewMessage(update.Message.Chat.ID, "")
					for _, link := range shortLinkResult {
						msg4.Text = msg4.Text + link + "\n"
					}
					msg4.Text = msg4.Text + fmt.Sprintf("\nSuccess count : %d\nError count : %d\n", successShortLinkCount, errorShortLinkCount)
					msg4.Text = msg4.Text + fmt.Sprintf("\nCreate %+v links with Rebrandly\n%+v links left", usedCount, 500-usedCount)
					msg4.DisableWebPagePreview = true

					bot.Send(msg4)

					// msg complete
					msg.Text = "Create link complete" + fmt.Sprintf(" in %+v", time.Since(createshortlinkTimer))
				}

			case "setrebrandapi":
				arg := update.Message.CommandArguments()
				if len(arg) < 1 {
					msg.Text = "setrebrandapi : no argument found"
				} else {
					err := rb.SetRebrandlyAPIKey(arg)
					if err != nil {
						msg.Text = fmt.Sprintf("setrebrandapi error : %+v ", err)
					} else {
						msg.Text = "setrebrandapi : set new API key successfully"
					}
				}

			default:
				msg.Text = "Unknown command"
				bot.Send(tb.NewMessage(update.Message.Chat.ID, "Still alive"))
			}

			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}
		}

	}
}
