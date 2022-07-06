package main

import (
	"flag"
	"github.com/denis0108/BotHpd/clients/client"
	"log"
)

const host = "api.telegram.org"

func main() {
	//t := mustToken()
	client := client.New(host, mustToken())
	client.SendMessage()
}

func mustToken() string {
	token := flag.String(
		"token-bot",
		"",
		"token for telegram bot",
	)
	flag.Parse()

	if *token == "" {
		log.Fatal("Token must be specified")
	}
	return *token

}
