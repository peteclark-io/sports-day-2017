package main

import (
	"os"
	"sort"
	"strconv"

	log "github.com/Sirupsen/logrus"
	"github.com/olekukonko/tablewriter"
)

func newScores() *scores {
	return &scores{GreenJersey: make(map[Player]int), YellowJersey: make(map[Player]int), PolkaDotersey: make(map[Player]int)}
}

func (s *scores) StageResults(e event, boost int, placings []Player) {
	for i, p := range placings {
		switch i {
		case 0:
			s.tally(e, 60+boost, p)
		case 1:
			s.tally(e, 30+boost, p)
		case 2:
			s.tally(e, 20+boost, p)
		case 3:
			s.tally(e, 10+boost, p)
		}
	}
}

func (s *scores) TeamStageResults(e event, boost int, placings []Player) {
	for i, p := range placings {
		switch i {
		case 0:
			s.tally(e, 60+boost, p)
		case 1:
			s.tally(e, 30+boost, p)
		case 2:
			s.tally(e, 20+boost, p)
		case 3:
			s.tally(e, 10+boost, p)
		}
	}
}

func (s *scores) tally(e event, score int, p Player) {
	switch e.Style {
	case skillfulStage:
		s.GreenJersey[p] = s.GreenJersey[p] + score
	case smartStage:
		s.PolkaDotersey[p] = s.PolkaDotersey[p] + score
	}

	s.YellowJersey[p] = s.YellowJersey[p] + score
}

func (s *scores) print() {
	log.Info("Yellow Jersey scores")
	printTable(s.YellowJersey)

	log.Info("Green Jersey scores")
	printTable(s.GreenJersey)

	log.Info("Polka Dot Jersey scores")
	printTable(s.PolkaDotersey)
}

func printTable(scores map[Player]int) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Score"})

	s := make([]playerScore, 0)
	for k, v := range scores {
		s = append(s, playerScore{k, v})
	}
	sort.Sort(sort.Reverse(byScore(s)))
	for _, v := range s {
		table.Append([]string{v.Player.Name, strconv.Itoa(v.Score)})
	}

	table.Render()
}
