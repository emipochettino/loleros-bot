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

var regions = map[string]string{
	"euw_match":  "euw1",
	"na_match":   "na1",
	"las_match":  "la2",
	"lan_match":  "la1",
	"br_match":   "br1",
	"eune_match": "eun1",
	"jp_match":   "jp1",
	"oc_match":   "oc1",
	"ru_match":   "ru",
	"tr_match":   "tr1",
}

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

			command := update.Message.Command()
			if region, existCommand := regions[command]; existCommand {
				match, err = l.matchService.FindCurrentMatchByRegionAndSummonerName(region, update.Message.CommandArguments())
			} else {
				err = fmt.Errorf("command [%s] not found", command)
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
