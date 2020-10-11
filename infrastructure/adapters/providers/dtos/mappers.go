package infrastructure

import "github.com/emipochettino/loleros-bot/domain"

func MapToSummonerModel(summoner SummonerDTO, participant ParticipantDTO, leagues []LeagueInfo) domain.Summoner {
	var soloQueueLeague *LeagueInfo
	for _, league := range leagues {
		if "RANKED_SOLO_5x5" == league.QueueType {
			soloQueueLeague = &league
			break
		}
	}
	var league domain.League
	if soloQueueLeague != nil {
		league = domain.NewLeague(soloQueueLeague.Tier, soloQueueLeague.Rank, soloQueueLeague.Wins, soloQueueLeague.Losses)
	}

	return domain.NewSummoner(summoner.Name, summoner.Level, &league, &participant.TeamId)
}
