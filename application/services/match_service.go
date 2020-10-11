package application

import (
	"github.com/emipochettino/loleros-bot/domain"
	providers "github.com/emipochettino/loleros-bot/infrastructure/adapters/providers/dtos"
	"log"
)

type MatchService interface {
	FindCurrentMatchByRegionAndSummonerName(region string, summonerName string) (*domain.Match, error)
}

type RitoProvider interface {
	FindSummonerByRegionAndName(region string, name string) (*providers.SummonerDTO, error)
	FindSummonerByRegionAndId(region string, id string) (*providers.SummonerDTO, error)
	FindMatchBySummonerId(region string, summonerId string) (*providers.MatchDTO, error)
	FindQueueById(queueId int64) (*providers.QueueInfoDTO, error)
	FindLeaguesByRegionAndSummonerId(region string, summonerId string) ([]providers.LeagueInfo, error)
}

type matchService struct {
	ritoProvider RitoProvider
}

func (m matchService) FindCurrentMatchByRegionAndSummonerName(region string, summonerName string) (*domain.Match, error) {
	summoner, err := m.ritoProvider.FindSummonerByRegionAndName(region, summonerName)
	if err != nil {
		return nil, err
	}

	matchDTO, err := m.ritoProvider.FindMatchBySummonerId(region, summoner.Id)
	if err != nil {
		return nil, err
	}

	/*
		//TODO check if it is necessary
		queue, err := m.ritoProvider.FindQueueById(matchDTO.GameQueueConfigId)
		if err != nil {
			return nil, err
		}
	*/

	//TODO make this for async
	var summoners []domain.Summoner
	for _, participant := range matchDTO.Participants {
		summonerDTO, err := m.ritoProvider.FindSummonerByRegionAndId(region, participant.SummonerId)
		if err != nil {
			log.Println(err)
			continue
		}
		leagues, err := m.ritoProvider.FindLeaguesByRegionAndSummonerId(region, summonerDTO.Id)
		if err != nil {
			//log.Printf("error getting league for summonerDTO [%s], error [%s]", summonerDTO.Name, err)
			continue
		}
		summoner := providers.MapToSummonerModel(*summonerDTO, participant, leagues)
		//log.Printf("Summoner name: [%s], level [%d], %+v, %+v", summoner.GetName(), summoner.GetLevel(), summoner.GetLeague(), *summoner.GetTeam())

		summoners = append(summoners, summoner)
	}

	match := domain.NewMatch(matchDTO.GameStartTime, summoners)

	return &match, nil
}

func NewMatchService(provider RitoProvider) MatchService {
	return matchService{ritoProvider: provider}
}
