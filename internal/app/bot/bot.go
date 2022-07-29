package bot

import (
	"log"

	"github.com/DiasOrazbaev/CodeforcesBot/internal/app/store"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	config *Config
	bot    *tgbotapi.BotAPI
	store  *store.Store
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
	log.Println("Start configure store")
	if err := b.configureStore(); err != nil {
		return err
	}
	log.Println("Store configured successfully")
	log.Println("Start migrations")
	if err := b.store.MigrationUp(); err != nil {
		log.Fatalln("Failed on migrations: ", err)
	}
	log.Println("Migration do successfully")
	b.configureHandles()

	return nil
}

func (b *Bot) configureStore() error {
	st := store.NewStore(b.config.storeConfig)
	if err := st.Open(); err != nil {
		return err
	}
	b.store = st
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
