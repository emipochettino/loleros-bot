package bots

import (
	"fmt"
	"github.com/emipochettino/loleros-bot/domain"
	"strings"
)

//getAnswer map the match to a string
func getAnswer(match *domain.Match, err error) string {
	if err != nil {
		return err.Error()
	}

	var sb strings.Builder
	sb.WriteString("🔵🔵🔵 Blue Team 🔵🔵🔵\n")
	for index, summoner := range (*match).GetBlueTeam() {
		sb.WriteString(writeSummoner(index+1, summoner))
	}
	sb.WriteString("\n🔴🔴🔴 Red Team 🔴🔴🔴\n")
	for index, summoner := range (*match).GetRedTeam() {
		sb.WriteString(writeSummoner(index+1, summoner))
	}

	return sb.String()
}

func writeSummoner(number int, summoner domain.Summoner) string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("%d- %s - lvl: %d", number, summoner.GetName(), summoner.GetLevel()))
	if league := summoner.GetLeague(); league != nil {
		sb.WriteString(fmt.Sprintf(" - league: %s  %s - win rate: %0.2f", league.GetTier(), league.GetRank(), league.GetWinRate()))
	}
	sb.WriteString("\n")

	return sb.String()
}
