package application

import (
	providers "github.com/emipochettino/loleros-bot/infrastructure/adapters/providers/dtos"
	"log"
)

type MatchService interface {
	FindCurrentMatchBySummonerName(summonerName string)
}

type RitoProvider interface {
	FindSummonerByName(name string) (*providers.SummonerDTO, error)
}

type matchService struct {
	ritoProvider RitoProvider
}

func (m matchService) FindCurrentMatchBySummonerName(summonerName string) {
	summoner, err := m.ritoProvider.FindSummonerByName(summonerName)
	if err != nil {
		panic(err)
	}

	log.Printf("summoner: %+v", summoner)
}

func NewMatchService(provider RitoProvider) MatchService {
	return matchService{ritoProvider: provider}
}
