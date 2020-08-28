package infrastructure

// SummonerDTO dto to map answer from rito api
type SummonerDTO struct {
	Id            string `json:"id"`
	AccountId     string `json:"accountId"`
	Name          string `json:"name"`
	SummonerLevel int    `json:"summonerLevel"`
}
