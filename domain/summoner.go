package domain

type summoner struct {
	name   string
	level  int
	league League
	team   *string
}

const (
	BlueTeam = "BLUE"
	RedTeam  = "RED"
)

func (s summoner) GetName() string {
	return s.name
}

func (s summoner) GetLevel() int {
	return s.level
}

func (s summoner) GetLeague() League {
	return s.league
}

func (s summoner) GetTeam() *string {
	return s.team
}

type Summoner interface {
	GetName() string
	GetLevel() int
	GetLeague() League
	GetTeam() *string
}

func NewSummoner(name string, level int, league *League, teamId *int64) Summoner {
	var team *string
	if teamId != nil {
		temp := "BLUE"
		if *teamId == 200 {
			temp = "RED"
		}
		team = &temp
	}

	return summoner{
		name:   name,
		level:  level,
		league: *league,
		team:   team,
	}
}
