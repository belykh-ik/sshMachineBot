package main

import (
	"fmt"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	sshUser     = "user"
	sshPassword = "password"
	sshHost     = "ip:port"
)

func main() {

	botToken := os.Getenv("TOKEN")
	if botToken == "" {
		log.Fatalf("bot token not set")
	}

	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Fatalf("Ошибка при создании бота: %v", err)
	}

	bot.Debug = true
	fmt.Printf("Бот %s успешно запущен\n", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil && update.Message.Text != "" {
			chatID := update.Message.Chat.ID
			command := update.Message.Text

			if command == "/start" {
				msg := tgbotapi.NewMessage(chatID, "Привет! Отправьте команду для выполнения на сервере через SSH.")
				bot.Send(msg)
			} else {
				// Выполняем SSH-команду и отправляем результат
				output, err := executeSSHCommand(command)
				if err != nil {
					msg := tgbotapi.NewMessage(chatID, fmt.Sprintf("Ошибка при выполнении команды: %v", err))
					bot.Send(msg)
				} else {
					msg := tgbotapi.NewMessage(chatID, fmt.Sprintf("Результат команды:\n%s", output))
					bot.Send(msg)
				}
			}
		}
	}
}
