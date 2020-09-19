package bots

import (
	"crypto/tls"
	"fmt"
	application "github.com/emipochettino/loleros-bot/application/services"
	"github.com/emipochettino/loleros-bot/domain"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"net/http"
	"os"
	"strings"
)

type lolerosBot struct {
	matchService application.MatchService
	bot          *tgbotapi.BotAPI
}

func (l lolerosBot) Start() {
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60

	updatesChan, err := l.bot.GetUpdatesChan(updateConfig)
	if err != nil {
		panic(err)
	}

	for update := range updatesChan {
		go func(innerUpdate tgbotapi.Update) {
			if innerUpdate.Message == nil || !innerUpdate.Message.IsCommand() {
				return
			}

			var match *domain.Match
			var err error

			switch command := update.Message.Command(); command {
			case "euw_match":
				match, err = l.matchService.FindCurrentMatchByRegionAndSummonerName("euw1", update.Message.CommandArguments())
			case "na_match":
				match, err = l.matchService.FindCurrentMatchByRegionAndSummonerName("na1", update.Message.CommandArguments())
			case "las_match":
				match, err = l.matchService.FindCurrentMatchByRegionAndSummonerName("la2", update.Message.CommandArguments())
			case "lan_match":
				match, err = l.matchService.FindCurrentMatchByRegionAndSummonerName("la1", update.Message.CommandArguments())
			case "br_match":
				match, err = l.matchService.FindCurrentMatchByRegionAndSummonerName("br1", update.Message.CommandArguments())
			case "eune_match":
				match, err = l.matchService.FindCurrentMatchByRegionAndSummonerName("eun1", update.Message.CommandArguments())
			case "jp_match":
				match, err = l.matchService.FindCurrentMatchByRegionAndSummonerName("jp1", update.Message.CommandArguments())
			case "oc_match":
				match, err = l.matchService.FindCurrentMatchByRegionAndSummonerName("oc1", update.Message.CommandArguments())
			case "ru_match":
				match, err = l.matchService.FindCurrentMatchByRegionAndSummonerName("ru", update.Message.CommandArguments())
			case "tr_match":
				match, err = l.matchService.FindCurrentMatchByRegionAndSummonerName("tr1", update.Message.CommandArguments())
			default:
				fmt.Println(command)
			}

			answer := getAnswer(match, err)
			msg := tgbotapi.NewMessage(innerUpdate.Message.Chat.ID, answer)
			msg.ReplyToMessageID = innerUpdate.Message.MessageID
			l.bot.Send(msg)
		}(update)
	}
}

type Bot interface {
	Start()
}

func NewLolerosBot(matchService application.MatchService) Bot {
	token := os.Getenv("TELEGRAM_TOKEN")
	if len(token) == 0 {
		panic(fmt.Errorf("TELEGRAM_TOKEN not set"))
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	bot, err := tgbotapi.NewBotAPIWithClient(token, client)
	if err != nil {
		panic(err)
	}

	bot.Debug = strings.EqualFold("dev", os.Getenv("PROFILE"))

	log.Printf("Authorized on account %s", bot.Self.UserName)

	return lolerosBot{matchService: matchService, bot: bot}
}
