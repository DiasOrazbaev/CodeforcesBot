package bot

import "github.com/DiasOrazbaev/CodeforcesBot/internal/app/store"

type Config struct {
	BotToken    string `toml:"bot_token"`
	storeConfig *store.Config
}

func NewConfig() *Config {
	return &Config{
		storeConfig: store.NewConfig(),
	}
}
