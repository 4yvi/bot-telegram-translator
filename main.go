package main

import (
	"log"
	"github.com/Syfaro/telegram-bot-api"
	trans "github.com/aerokite/go-google-translate/pkg"
	"strings"
)

func main() {
	// Инициализация
	bot, err := tgbotapi.NewBotAPI("")
	if err != nil {
		log.Panic(err)
	}

	// Режим тестирования
	bot.Debug = false

	// Выводим сообщение об авторизации
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		// переменная в котором будет хранится сообщение
		message := ""

		// Определяем какое действие нужно совершить
		switch {
			case strings.HasPrefix(update.Message.Text, "/start"):
				message = "Привет 🖐\n"
				message += "🤖 Этот бот переводит текст с Русского на Английский и наоборот. \n"
				message += "👉🏻 Для того, чтобы перевести с Русского на Английский, нужно перед фразой добавить /ru \n"
				message += "👉🏻 Для того, чтобы перевести с Английского на Русский, нужно перед фразой добавить /en \n"
				message += "ℹ Для того, чтобы получить информацию о разработчике, нужно послать команду /info \n"
				break

			case strings.HasPrefix(update.Message.Text, "/info"):
				message = "🤖 Я pet проект моего хозяина)\n"
				message += "👉🏻 Создатель: Трофимов Алексей\n"
				message += "👉🏻 Сайт: https://4yvi.github.io\n"
				message += "👉🏻 Технологии: Golang + Google Translate\n"
				break

			case strings.HasPrefix(update.Message.Text, "/en"):
				message = translate( "en","ru", strings.Replace(update.Message.Text, "/en", "",1))
				break

			case strings.HasPrefix(update.Message.Text, "/ru"):
				message = translate("ru", "en", strings.Replace(update.Message.Text, "/ru", "",1))
				break

			default:
				message = "🤖 Я тебя не понял) \n"
				message += "👉🏻 Попробуй использовать команду /en и /ru\n"
				message += "ℹ Для того, чтобы получить информацию о разработчике, нужно послать команду /info"
				break
		}

		// Пишем сообщенеи
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)

		// Добовляем запрос
		msg.ReplyToMessageID = update.Message.MessageID

		// Печатаем информацию в консоль
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		// Посылаем сообщение
		bot.Send(msg)
	}
}

// Переводит переводит текст на другой язык
func translate(sourceLang string, targetLang string , text string) string {

	if (len(text) > 0) {
		req := &trans.TranslateRequest{
			SourceLang: sourceLang,
			TargetLang: targetLang,
			Text:       text,
		}

		translated, err := trans.Translate(req)

		if err != nil {
			log.Fatalln(err)
		}
		return  translated
	} else {
		return "🤖 Ты не ввел ничего! \n 👉🏻 Попробуй так: /" + sourceLang + " {текст который нужно перевести} \n"
	}


}