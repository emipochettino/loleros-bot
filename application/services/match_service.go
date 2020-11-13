package application

import (
	"fmt"
	application "github.com/emipochettino/loleros-bot/application/dto"
	providers "github.com/emipochettino/loleros-bot/infrastructure/adapters/providers/dtos"
	"log"
	"sync"
	"time"
)

type MatchService interface {
	FindCurrentMatchByRegionAndSummonerName(region string, summonerName string) ([]application.SummonerDTO, error)
}

type RitoProvider interface {
	FindSummonerByRegionAndName(region string, name string) (*providers.SummonerDTO, error)
	FindMatchBySummonerId(region string, summonerId string) (*providers.MatchDTO, error)
	FindSummonerByRegionAndId(region string, id string) (*providers.SummonerDTO, error)
	FindLeaguesByRegionAndSummonerId(region string, summonerId string) ([]providers.LeagueInfoDTO, error)
}

type matchService struct {
	ritoProvider RitoProvider
	mu           *sync.Mutex
}

func (m matchService) FindCurrentMatchByRegionAndSummonerName(region string, summonerName string) ([]application.SummonerDTO, error) {
	start := time.Now()
	summonerDTO, err := m.ritoProvider.FindSummonerByRegionAndName(region, summonerName)
	if err != nil {
		return nil, err
	}

	matchDTO, err := m.ritoProvider.FindMatchBySummonerId(region, summonerDTO.Id)
	if err != nil {
		return nil, err
	}

	//TODO make this for async
	var wg sync.WaitGroup
	wg.Add(10)
	var summoners []application.SummonerDTO
	for _, participant := range matchDTO.Participants {
		go func(participant providers.ParticipantDTO) {
			defer wg.Done()
			summonerDTO, err := m.ritoProvider.FindSummonerByRegionAndId(region, participant.SummonerId)
			if err != nil {
				log.Println(err)
				return
			}
			leaguesDTO, err := m.ritoProvider.FindLeaguesByRegionAndSummonerId(region, summonerDTO.Id)
			if err != nil {
				return
			}
			summoner := providers.MapToSummonerModel(*summonerDTO, participant, leaguesDTO)
			m.mu.Lock()
			summoners = append(summoners, summoner)
			m.mu.Unlock()
		}(participant)
	}
	wg.Wait()
	print(fmt.Sprintf("\ntime: %.2f seconds\n", time.Now().Sub(start).Seconds()))
	return summoners, nil
}

func NewMatchService(provider RitoProvider) MatchService {
	return matchService{ritoProvider: provider, mu: &sync.Mutex{}}
}
