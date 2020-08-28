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
}

func (r ritoProvider) FindSummonerByName(name string) (*providers.SummonerDTO, error) {
	request, err := http.NewRequest("GET", fmt.Sprintf("https://euw1.api.riotgames.com/lol/summoner/v4/summoners/by-name/%s", name), nil)
	if err != nil {
		return nil, err
	}
	request.Header.Add("X-Riot-Token", r.token)

	response, err := r.client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode == 401 {
		return nil, fmt.Errorf("rito token can be expired")
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
