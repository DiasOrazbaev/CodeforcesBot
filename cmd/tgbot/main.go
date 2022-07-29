package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/DiasOrazbaev/CodeforcesBot/internal/app/bot"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "", "path to the botconfig.toml file")
}

func main() {
	flag.Parse()

	config := bot.NewConfig()

	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatalln(err)
	}
	bot := bot.NewBot(config)
	if err := bot.Start(); err != nil {
		log.Fatalln(err)
	}
}
