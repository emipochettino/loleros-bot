package domain

import "time"

type match struct {
	startTime time.Time
	blueTeam  []Summoner
	redTeam   []Summoner
}

func (m match) GetStartTime() time.Time {
	return m.GetStartTime()
}

func (m match) GetCurrentTimeInMinutes() int {
	return int(time.Now().Sub(m.startTime).Minutes())
}

func (m match) GetBlueTeam() []Summoner {
	return m.blueTeam
}

func (m match) GetRedTeam() []Summoner {
	return m.redTeam
}

type Match interface {
	GetStartTime() time.Time
	GetCurrentTimeInMinutes() int
	GetBlueTeam() []Summoner
	GetRedTeam() []Summoner
}

func NewMatch(startTimeMillis int64, summoners []Summoner) Match {
	startTime := time.Unix(0, startTimeMillis*int64(time.Millisecond))

	var blueTeam []Summoner
	var redTeam []Summoner
	for _, summoner := range summoners {
		if summoner.GetTeam() != nil {
			if *summoner.GetTeam() == BlueTeam {
				blueTeam = append(blueTeam, summoner)
			} else if *summoner.GetTeam() == RedTeam {
				redTeam = append(redTeam, summoner)
			}
		}
	}

	return match{
		startTime: startTime,
		blueTeam:  blueTeam,
		redTeam:   redTeam,
	}
}
