package infrastructure

import (
	application "github.com/emipochettino/loleros-bot/application/dto"
)

func MapToSummonerModel(summoner SummonerDTO, participant ParticipantDTO, leagues []LeagueInfoDTO) application.SummonerDTO {
	var soloQueueLeague *LeagueInfoDTO
	for _, league := range leagues {
		if "RANKED_SOLO_5x5" == league.QueueType {
			soloQueueLeague = &league
			break
		}
	}
	var league *application.LeagueDTO
	if soloQueueLeague != nil {
		league = &application.LeagueDTO{
			Tier:    soloQueueLeague.Tier,
			Rank:    soloQueueLeague.Rank,
			Wins:    soloQueueLeague.Wins,
			Losses:  soloQueueLeague.Losses,
			WinRate: float32(soloQueueLeague.Wins) / (float32(soloQueueLeague.Wins) + float32(soloQueueLeague.Losses)),
		}
	}

	return application.SummonerDTO{
		Name:   summoner.Name,
		Level:  summoner.Level,
		Team:   int(participant.TeamId),
		League: league,
	}
}
