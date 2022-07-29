.PHONY: build
build:
	go build -v ./cmd/tgbot

.PHONY: start
start:
	.\tgbot.exe -config-path .\configs\botconfig.toml

.PHONY: work
work:
	make build start

.DEFAULT := work