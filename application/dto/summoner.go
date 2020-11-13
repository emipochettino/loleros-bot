package application

type SummonerDTO struct {
	Name   string
	Level  int
	League *LeagueDTO
	Team   int
}

type LeagueDTO struct {
	Tier    string
	Rank    string
	Wins    int
	Losses  int
	WinRate float32
}
