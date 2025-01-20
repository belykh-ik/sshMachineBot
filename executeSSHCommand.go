package main

import (
	"bytes"
	"fmt"

	"golang.org/x/crypto/ssh"
)

func executeSSHCommand(sshHost string, sshUser string, sshPassword string, command string) (string, error) {
	config := &ssh.ClientConfig{
		User: sshUser,
		Auth: []ssh.AuthMethod{
			ssh.Password(sshPassword),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// Подключаемся к серверу
	client, err := ssh.Dial("tcp", sshHost, config)
	if err != nil {
		return "", fmt.Errorf("не удалось подключиться по SSH: %w", err)
	}
	defer client.Close()

	// Открываем сессию
	session, err := client.NewSession()
	if err != nil {
		return "", fmt.Errorf("не удалось создать SSH-сессию: %w", err)
	}
	defer session.Close()

	// Выполняем команду
	var stdout, stderr bytes.Buffer
	session.Stdout = &stdout
	session.Stderr = &stderr

	if err := session.Run(command); err != nil {
		return "", fmt.Errorf("ошибка выполнения команды: %w, stderr: %s", err, stderr.String())
	}

	return stdout.String(), nil
}
