# 🔐 sshMachineBot

A powerful **Telegram bot** written in Go that allows you to remotely execute SSH commands on your servers directly from Telegram. Manage your infrastructure, run diagnostics, and automate routine tasks—all from your chat window.

---

## ✨ Features

- 🤖 **Telegram Bot Interface** – send SSH commands via chat  
- 🔒 **Secure SSH Connections** using key‑based authentication  
- 🌐 **Multi‑host Support** – execute on one or multiple configured hosts  
- ⚙️ **Command Templates** – predefine common commands for quick execution  
- 🐳 **Docker‑ready** for hassle‑free deployment  

---

## 🛠️ Technology Stack

- **Go 1.18+** – core bot logic and SSH client  
- **Telegram Bot API** – user interaction  
- **golang.org/x/crypto/ssh** – secure SSH connections  
- **Docker** – containerized deployment  

---

## 🚀 Getting Started

### 1. Prerequisites

- A **Telegram bot token** from [@BotFather](https://t.me/BotFather)  
- SSH private key with access to your target servers  
- Go installed (or Docker)

### 2. Clone the Repository

```bash
git clone https://github.com/belykh-ik/sshMachineBot.git
cd sshMachineBot
```

### 3. Build and run the container

```bash
docker build -t sshmachinebot .
docker run -d --rm \
  -e TELEGRAM_TOKEN=$TELEGRAM_TOKEN \
  sshmachinebot
```

