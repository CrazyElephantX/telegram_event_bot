package main

import (
	"fmt"
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"log"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("")
	if err != nil {
		log.Panic(err)
	}
	const LineBreak = "\r\n"
	//	bot.Debug = false
	//	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	Text := "sorry bro"

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		//log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		switch update.Message.Text {
		case "/start":
			Text = fmt.Sprintf(`Привет @%s! Я знаю все ответы на вопросы по мероприятию. Нажми /help для получения списка команд`,
				update.Message.From.UserName)
		case "/help":
			Text = "/Location" + LineBreak +
				"/org" + LineBreak +
				"/teh" + LineBreak +
				"/muah" + LineBreak
		case "/Location":
			Text = "Россия. Москва"
		case "/org":
			Text = "+79265223959 Михаил"
		case "/teh":
			Text = "https://rig-ru.tech/"
		case "/muah":
			Text = "+79636860686 Алена"
		default:
			Text = "Пока =)"
		}

		//button := tgbotapi.KeyboardButton{}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, Text)
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}
}
