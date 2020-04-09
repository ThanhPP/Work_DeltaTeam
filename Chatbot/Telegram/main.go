package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"

	ggsheet "github.com/THANHPP/Work_DeltaTeam/Chatbot/Telegram/handler/googlesheet"
	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {
	// bot settings
	bot1, err := tb.NewBot(tb.Settings{
		Token:  "1032135930:AAG7_bSPwq8ih4rc3cBHsz5UDyYd-q9gr8g",
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	// bot handling
	bot1.Handle("/hello", func(m *tb.Message) {
		bot1.Send(m.Sender, "hello world")
	})

	// Start the bot
	//bot1.Start()
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Import data range : ")
	datarange, _ := reader.ReadString('\n')
	datarange = datarange[:len(datarange)-1]
	ggsheet.GetDataFromRage(datarange)
}
