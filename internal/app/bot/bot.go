package bot

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	config *Config
	bot    *tgbotapi.BotAPI
}

func NewBot(config *Config) *Bot {
	return &Bot{config: config}
}

func (b *Bot) Start() error {
	bot, err := tgbotapi.NewBotAPI(b.config.BotToken)
	if err != nil {
		return err
	}
	b.bot = bot
	log.Println("Authorized on account " + b.bot.Self.UserName)
	b.configureHandles()
	return nil
}

func (b *Bot) configureHandles() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := b.bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message == nil {
			continue
		}
		if !update.Message.IsCommand() {
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.From.ID, "")
		switch update.Message.Command() {
		case "hello":
			msg.Text = "Hi"
		}

		if _, err := b.bot.Send(msg); err != nil {
			log.Println(err)
		}
	}
}
