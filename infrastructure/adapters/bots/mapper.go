package bots

import (
	"fmt"
	application "github.com/emipochettino/loleros-bot/application/dto"
	"strings"
)

//getAnswer map the match to a string
func getAnswer(summoners []application.SummonerDTO, err error) string {
	if err != nil {
		return err.Error()
	}
	var blueTeam []application.SummonerDTO
	var redTeam []application.SummonerDTO

	for _, summoner := range summoners {
		if summoner.Team == 200 {
			redTeam = append(redTeam, summoner)
		} else {
			blueTeam = append(blueTeam, summoner)
		}
	}

	var sb strings.Builder
	sb.WriteString("🔵🔵🔵 Blue Team 🔵🔵🔵\n")
	for index, summoner := range blueTeam {
		sb.WriteString(writeSummoner(index+1, summoner))
	}
	sb.WriteString("\n🔴🔴🔴 Red Team 🔴🔴🔴\n")
	for index, summoner := range redTeam {
		sb.WriteString(writeSummoner(index+1, summoner))
	}

	return sb.String()
}

func writeSummoner(number int, summoner application.SummonerDTO) string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("%d- %s - lvl: %d", number, summoner.Name, summoner.Level))
	if league := summoner.League; league != nil {
		sb.WriteString(fmt.Sprintf(" - league: %s  %s - win rate: %0.2f", league.Tier, league.Rank, league.WinRate))
	}
	sb.WriteString("\n")

	return sb.String()
}
