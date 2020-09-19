package domain

import "math"

type league struct {
	tier    string
	rank    string
	winRate float32
}

func (l league) GetTier() string {
	return l.tier
}

func (l league) GetRank() string {
	return l.rank
}

func (l league) GetWinRate() float32 {
	return l.winRate
}

type League interface {
	GetTier() string
	GetRank() string
	GetWinRate() float32
}

func NewLeague(tier string, rank string, wins int, losses int) League {
	winRate := float32(math.Round((float64(wins)*100.0/float64(wins+losses))*100) / 100)
	return league{
		tier:    tier,
		rank:    rank,
		winRate: winRate,
	}
}
