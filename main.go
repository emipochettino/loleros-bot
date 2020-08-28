package main

import (
	application "github.com/emipochettino/loleros-bot/application/services"
	infrastructure "github.com/emipochettino/loleros-bot/infrastructure/adapters/providers"
	"os"
)

func main() {
	ritoToken := os.Getenv("RITO_TOKEN")

	provider, err := infrastructure.NewRitoProvider(ritoToken)
	if err != nil {
		panic(err)
	}

	matchService := application.NewMatchService(provider)

	//TODO parametrize SERVER (EUW, LAS, LAN, NA)
	matchService.FindCurrentMatchBySummonerName("test")
}
