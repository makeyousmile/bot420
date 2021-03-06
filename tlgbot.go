package main

import (
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"log"
)

func startBot() {

	bot, err := tgbotapi.NewBotAPI("468353575:AAFBaLyb20M3iMmjqBnm3js-qNrRjL81bFk")
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)
	ucfg := tgbotapi.NewUpdate(0)
	ucfg.Timeout = 60
	updates, err := bot.GetUpdatesChan(ucfg)
	if err != nil {
		log.Print(err)
	}

	for update := range updates {
		if update.Message == nil {
			continue
		}
		typing := tgbotapi.NewChatAction(update.Message.Chat.ID, "typing")
		bot.Send(typing)
		if update.Message.Command() == "go" {
			for i, row := range hs {

				msgtext := "<b>" + row.Title + "</b> Цена:<b>" + row.Price + ".</b> " + row.Market
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, msgtext)
				msg.ParseMode = "HTML"
				//msg.ReplyToMessageID = update.Message.MessageID
				bot.Send(msg)

				if i == len(hs)-1 {
					msgtext := "Последние обновление: " + hs[i].UpdateTime.Format("2006-01-02 15:04:05")
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, msgtext)
					msg.ParseMode = "HTML"
					bot.Send(msg)

				}
			}

		}
		if update.Message.Command() == "start" {
			msgtxt := "Привет. Отправь мне команду /go и получи список доступных предложений. "
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, msgtxt)
			bot.Send(msg)
		}
		if update.Message.Command() == "help" {
			msgtxt := "помоги себе сам"
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, msgtxt)
			msg.ReplyToMessageID = update.Message.MessageID
			bot.Send(msg)
		}
		//bot.Send(msg)
	}
}
