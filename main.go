package main

import (
	"flag"
	"log"
)

func main() {
	//t := mustToken()
	// client := client.New()
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
