package main

import (
	"encoding/json"
	application "github.com/emipochettino/loleros-bot/application/services"
	"github.com/emipochettino/loleros-bot/infrastructure/adapters/bots"
	infrastructure "github.com/emipochettino/loleros-bot/infrastructure/adapters/providers"
	"github.com/patrickmn/go-cache"
	"os"
	"time"
)

func main() {
	ritoToken := os.Getenv("RITO_TOKEN")
	ritoHost := map[string]string{
		"euw1": "https://euw1.api.riotgames.com",
		"na1":  "https://na1.api.riotgames.com",
		"la2":  "https://la2.api.riotgames.com",
		"la1":  "https://la1.api.riotgames.com",
		"br1":  "https://br1.api.riotgames.com",
		"eun1": "https://eun1.api.riotgames.com",
		"jp1":  "https://jp1.api.riotgames.com",
		"oc1":  "https://oc1.api.riotgames.com",
		"ru":   "https://ru.api.riotgames.com",
		"tr1":  "https://tr1.api.riotgames.com",
	}

	c := cache.New(30*time.Minute, 40*time.Minute)

	provider, err := infrastructure.NewRitoProvider(ritoHost, ritoToken, c)
	if err != nil {
		panic(err)
	}

	matchService := application.NewMatchService(provider)

	//result, err := matchService.FindCurrentMatchByRegionAndSummonerName("euw1", "EłOjoNinja")
	//result, err := matchService.FindCurrentMatchByRegionAndSummonerName("oc1", "potatobrush")
	result, err := matchService.FindCurrentMatchByRegionAndSummonerName("na1", "Nightblue3")
	if err != nil {
		print(err.Error())
	} else {
		jsonByte, _ := json.MarshalIndent(result, "", "    ")
		print(string(jsonByte))
	}

	lolerosBot := bots.NewLolerosBot(matchService)
	lolerosBot.Start()
}
