package bot

import (
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

	return nil
}
