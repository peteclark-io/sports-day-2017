package main

const skillfulStage = "skill"
const smartStage = "smart"
const teamStage = "team"

type Player struct {
	Name   string `json:"name"`
	Skill  int    `json:"skill"`
	Smarts int    `json:"smarts"`
}

type team struct {
	Name    string   `json:"name"`
	Players []Player `json:"players"`
}

type event struct {
	Event  string `json:"event"`
	Style  string `json:"style"`
	Rating int    `json:"rating"`
}

type eventSelection struct {
	Event    event    `json:"event"`
	Entrants []Player `json:"entrants"`
	TopSeeds bool     `json:"topSeeds"`
}

type stage struct {
	Type   string  `json:"type"`
	Events []event `json:"events"`
}

type scores struct {
	GreenJersey   map[Player]int `json:"greenJersey"`
	YellowJersey  map[Player]int `json:"yellowJersey"`
	PolkaDotersey map[Player]int `json:"polkaDotersey"`
	Teams         map[string]int `json:"teams"`

	teamsData map[string]team
}

type score struct {
	Team  string `json:"team"`
	Score int    `json:"score"`
}

type playerScore struct {
	Player Player `json:"player"`
	Score  int    `json:"score"`
}

type byPlayerScore []playerScore

func (a byPlayerScore) Len() int           { return len(a) }
func (a byPlayerScore) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byPlayerScore) Less(i, j int) bool { return a[i].Score < a[j].Score }

type byScore []score

func (a byScore) Len() int           { return len(a) }
func (a byScore) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byScore) Less(i, j int) bool { return a[i].Score < a[j].Score }

type bySkill []Player

func (a bySkill) Len() int           { return len(a) }
func (a bySkill) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a bySkill) Less(i, j int) bool { return a[i].Skill < a[j].Skill }

type bySmarts []Player

func (a bySmarts) Len() int           { return len(a) }
func (a bySmarts) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a bySmarts) Less(i, j int) bool { return a[i].Smarts < a[j].Smarts }
