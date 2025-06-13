package main

import (
	"fmt"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {

	botToken := os.Getenv("TOKEN")
	if botToken == "" {
		log.Fatalf("BOT TOKEN not set")
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

	//Состояние переключателя для записи данных от пользователя
	userSwitch := 0

	//Входная точка
	var command string

	//Переменные для подключения к серверу
	var sshHost string
	var sshUser string
	var sshPassword string

	for update := range updates {
		if update.Message != nil && update.Message.Text != "" {
			chatID := update.Message.Chat.ID
			if command == "" {
				command = update.Message.Text
			}

			if command == "/start" {
				switch userSwitch {
				case 0:
					msg := tgbotapi.NewMessage(chatID, "Привет! Чтобы отправлять команды на сервер, давай к нему подключимся:\nДля начала введи адрес и порт хоста в формате:\nip:port\n")
					bot.Send(msg)
					userSwitch = 1
				case 1:
					sshHost = update.Message.Text
					msg := tgbotapi.NewMessage(chatID, "Отлично, теперь введи под каким поьзователем ты хочешь подключиться:\n")
					bot.Send(msg)
					userSwitch = 2
				case 2:
					sshUser = update.Message.Text
					msg := tgbotapi.NewMessage(chatID, "Отлично, теперь введи пароль от этого пользователя:\n")
					bot.Send(msg)
					userSwitch = 3
				case 3:
					sshPassword = update.Message.Text
					msg := tgbotapi.NewMessage(chatID, "Отлично, теперь введи команду, которую необходимо выполнить:\n")
					bot.Send(msg)
					command = ""
					userSwitch = 0
				}
			} else {
				// Выполняем SSH-команду и отправляем результат
				output, err := executeSSHCommand(sshHost, sshUser, sshPassword, command)
				if err != nil {
					msg := tgbotapi.NewMessage(chatID, fmt.Sprintf("Ошибка при выполнении команды: %v", err))
					bot.Send(msg)
				} else {
					msg := tgbotapi.NewMessage(chatID, fmt.Sprintf("Результат команды:\n%s", output))
					bot.Send(msg)
				}
				command = ""
			}
		}
	}
}
