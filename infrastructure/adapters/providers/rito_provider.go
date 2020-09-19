package infrastructure

import (
	"encoding/json"
	"fmt"
	"github.com/emipochettino/loleros-bot/application/services"
	providers "github.com/emipochettino/loleros-bot/infrastructure/adapters/providers/dtos"
	"net/http"
)

type ritoProvider struct {
	client http.Client
	token  string
	assets providers.Assets
}

func (r ritoProvider) FindSummonerByRegionAndId(region string, id string) (*providers.SummonerDTO, error) {
	request, err := http.NewRequest("GET", fmt.Sprintf("https://%s.api.riotgames.com/lol/summoner/v4/summoners/%s", region, id), nil)
	if err != nil {
		return nil, err
	}
	request.Header.Add("X-Riot-Token", r.token)

	response, err := r.client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode == 403 {
		return nil, fmt.Errorf("rito token can be expired")
	}

	if response.StatusCode == 404 {
		return nil, fmt.Errorf("summoner not found")
	}

	var summonerDTO providers.SummonerDTO
	err = json.NewDecoder(response.Body).Decode(&summonerDTO)
	if err != nil {
		return nil, err
	}

	return &summonerDTO, nil
}

func (r ritoProvider) FindLeaguesByRegionAndSummonerId(region string, summonerId string) ([]providers.LeagueInfo, error) {
	request, err := http.NewRequest("GET", fmt.Sprintf("https://%s.api.riotgames.com/lol/league/v4/entries/by-summoner/%s", region, summonerId), nil)
	if err != nil {
		return nil, err
	}
	request.Header.Add("X-Riot-Token", r.token)

	response, err := r.client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode == 403 {
		return nil, fmt.Errorf("rito token can be expired")
	}

	var leagues []providers.LeagueInfo
	err = json.NewDecoder(response.Body).Decode(&leagues)
	if err != nil {
		return nil, err
	}

	return leagues, nil
}

func (r ritoProvider) FindMatchBySummonerId(region string, summonerId string) (*providers.MatchDTO, error) {
	request, err := http.NewRequest("GET", fmt.Sprintf("https://%s.api.riotgames.com/lol/spectator/v4/active-games/by-summoner/%s", region, summonerId), nil)
	if err != nil {
		return nil, err
	}
	request.Header.Add("X-Riot-Token", r.token)

	response, err := r.client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode == 403 {
		return nil, fmt.Errorf("rito token can be expired")
	}

	if response.StatusCode == 404 {
		return nil, fmt.Errorf("match not found")
	}

	var matchDTO providers.MatchDTO
	err = json.NewDecoder(response.Body).Decode(&matchDTO)
	if err != nil {
		return nil, err
	}

	return &matchDTO, nil
}

func (r ritoProvider) FindQueueById(queueId int64) (*providers.QueueInfoDTO, error) {
	if r.assets.Queues == nil || len(r.assets.Queues) == 0 {
		response, err := r.client.Get("http://static.developer.riotgames.com/docs/lol/queues.json")
		if err != nil {
			panic(err)
		}
		defer response.Body.Close()
		err = json.NewDecoder(response.Body).Decode(&r.assets.Queues)
		if err != nil {
			panic(err)
		}
	}

	for _, queue := range r.assets.Queues {
		if queue.QueueId == queueId {
			return &queue, nil
		}
	}

	return nil, fmt.Errorf("could not find the queue with id: %d", queueId)
}

func (r ritoProvider) FindSummonerByRegionAndName(region string, name string) (*providers.SummonerDTO, error) {
	request, err := http.NewRequest("GET", fmt.Sprintf("https://%s.api.riotgames.com/lol/summoner/v4/summoners/by-name/%s", region, name), nil)
	if err != nil {
		return nil, err
	}
	request.Header.Add("X-Riot-Token", r.token)

	response, err := r.client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode == 403 {
		return nil, fmt.Errorf("rito token can be expired")
	}

	if response.StatusCode == 404 {
		return nil, fmt.Errorf("summoner not found")
	}

	var summonerDTO providers.SummonerDTO
	err = json.NewDecoder(response.Body).Decode(&summonerDTO)
	if err != nil {
		return nil, err
	}

	return &summonerDTO, nil
}

func NewRitoProvider(token string) (application.RitoProvider, error) {
	if len(token) == 0 {
		return nil, fmt.Errorf("rito token should exist")
	}
	return ritoProvider{client: http.Client{}, token: token}, nil
}
