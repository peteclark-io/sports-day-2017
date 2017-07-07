package main

import (
	"math/rand"
	"os"
	"strconv"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/olekukonko/tablewriter"
)

func (s event) PlayTeamEvent(teams map[string]team) []team {
	teamRatings := make(map[interface{}]int)
	for _, t := range teams {
		for _, p := range t.Players {
			teamRatings[t.Name] = teamRatings[t.Name] + p.Smarts + p.Skill
		}
	}

	places := make([]team, len(teams))
	for i := 0; i < len(teams); i++ {
		placed := s.nextPlace(teamRatings).(string)
		places[i] = teams[placed]
		delete(teamRatings, placed)
	}

	return places
}

func (s event) Play(players ...Player) []Player {
	playerRatings := make(map[interface{}]int)
	for _, p := range players {
		switch s.Style {
		case skillfulStage:
			playerRatings[p] = p.Skill
		case smartStage:
			playerRatings[p] = p.Smarts
		case teamStage:
			playerRatings[p] = p.Smarts + p.Skill
		}
	}

	places := make([]Player, len(players))
	for i := 0; i < len(players); i++ {
		places[i] = s.nextPlace(playerRatings).(Player)
		delete(playerRatings, places[i])
	}

	return places
}

func (s event) nextPlace(players map[interface{}]int) interface{} {
	scores := make(map[interface{}]int)

	sum := 0
	for k, v := range players {
		chance := v * s.Rating
		sum += chance
		scores[k] = sum
	}

	result := random(0, sum)
	for k, v := range scores {
		if result < v {
			return k
		}
	}

	log.Warn("bug happened")
	return Player{}
}

func printStageResults(standings []Player) {
	table := tablewriter.NewWriter(os.Stdout)

	table.SetHeader([]string{"Place", "Name"})
	for i, v := range standings {
		table.Append([]string{strconv.Itoa(i + 1), v.Name})
	}
	table.Render()
}

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}
