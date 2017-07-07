package main

import (
	"os"
	"sort"
	"strconv"

	log "github.com/Sirupsen/logrus"
	"github.com/olekukonko/tablewriter"
)

func newScores(teams map[string]team) *scores {
	return &scores{GreenJersey: make(map[Player]int), YellowJersey: make(map[Player]int), PolkaDotersey: make(map[Player]int), Teams: make(map[string]int), teamsData: teams}
}

func (s *scores) StageResults(e event, boost int, placings []Player) {
	for i, p := range placings {
		switch i {
		case 0:
			s.tally(e, 60, p)
		case 1:
			s.tally(e, 30, p)
		case 2:
			s.tally(e, 20, p)
		case 3:
			s.tally(e, 10, p)
		}
	}
}

func (s *scores) TeamStageResults(e event, boost int, placings []team) {
	for i, p := range placings {
		switch i {
		case 0:
			s.Teams[p.Name] = s.Teams[p.Name] + 50
		case 1:
			s.Teams[p.Name] = s.Teams[p.Name] + 30
		case 2:
			s.Teams[p.Name] = s.Teams[p.Name] + 20
		case 3:
			s.Teams[p.Name] = s.Teams[p.Name] + 10
		}
	}
}

func (s *scores) lookupTeam(p Player) string {
	for _, team := range s.teamsData {
		for _, teamPlayer := range team.Players {
			if p.Name == teamPlayer.Name {
				return team.Name
			}
		}
	}
	log.Warn("bug happened")
	return ""
}

func (s *scores) tally(e event, score int, p Player) {
	switch e.Style {
	case skillfulStage:
		s.GreenJersey[p] = s.GreenJersey[p] + score
	case smartStage:
		s.PolkaDotersey[p] = s.PolkaDotersey[p] + score
	}

	s.YellowJersey[p] = s.YellowJersey[p] + score

	teamName := s.lookupTeam(p)
	s.Teams[teamName] = s.Teams[teamName] + score
}

func (s *scores) jerseys() {
	yellows := determineJerseyOrder(s.YellowJersey)
	greens := determineJerseyOrder(s.GreenJersey)
	polka := determineJerseyOrder(s.PolkaDotersey)

	yellowWinningTeam := s.lookupTeam(yellows[0].Player)
	s.Teams[yellowWinningTeam] = s.Teams[yellowWinningTeam] + 30

	greenWinner := nextJerseyWinner(greens, yellows[0].Player)
	greenWinningTeam := s.lookupTeam(greenWinner)
	s.Teams[greenWinningTeam] = s.Teams[greenWinningTeam] + 20

	polkaWinner := nextJerseyWinner(polka, yellows[0].Player, greenWinner)
	polkaWinningTeam := s.lookupTeam(polkaWinner)
	s.Teams[polkaWinningTeam] = s.Teams[polkaWinningTeam] + 20

	log.Infof("Green Winner! %v", greenWinner.Name)
	log.Infof("Polka Dot Winner! %v", polkaWinner.Name)
	log.Infof("Yellow Winner! %v", yellows[0].Player.Name)
}

func nextJerseyWinner(order []playerScore, prev ...Player) Player {
	for _, v := range order {
		if !arrContains(v.Player, prev) {
			return v.Player
		}
	}
	log.Warn("lol bug happened")
	return Player{}
}

func arrContains(p Player, ps []Player) bool {
	for _, v := range ps {
		if p == v {
			return true
		}
	}
	return false
}

func (s *scores) print() {
	s.jerseys()

	log.Info("Yellow Jersey scores")
	s.printTable(s.YellowJersey)

	log.Info("Green Jersey scores")
	s.printTable(s.GreenJersey)

	log.Info("Polka Dot Jersey scores")
	s.printTable(s.PolkaDotersey)

	log.Info("Final scores")
	printTeamsTable(s.Teams)
}

func printTeamsTable(scores map[string]int) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Score"})

	s := make([]score, 0)
	for k, v := range scores {
		s = append(s, score{k, v})
	}
	sort.Sort(sort.Reverse(byScore(s)))
	for _, v := range s {
		table.Append([]string{v.Team, strconv.Itoa(v.Score)})
	}

	table.Render()
}

func (s *scores) printTable(scores map[Player]int) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Team", "Name", "Score"})

	r := make([]playerScore, 0)
	for k, v := range scores {
		r = append(r, playerScore{k, v})
	}
	sort.Sort(sort.Reverse(byPlayerScore(r)))
	for _, v := range r {
		table.Append([]string{s.lookupTeam(v.Player), v.Player.Name, strconv.Itoa(v.Score)})
	}

	table.Render()
}

func determineJerseyOrder(scores map[Player]int) []playerScore {
	s := make([]playerScore, 0)
	for k, v := range scores {
		s = append(s, playerScore{k, v})
	}
	sort.Sort(sort.Reverse(byPlayerScore(s)))
	return s
}
