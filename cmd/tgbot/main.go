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

	config := bot.New()

	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatalln(err)
	}
}
