package bot

type Config struct {
	BotToken string `toml:"bot_token"`
}

func New() *Config {
	return &Config{}
}
