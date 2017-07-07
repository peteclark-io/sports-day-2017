package main

import (
	"sort"
)

func (t team) SelectPlayers(events ...event) []eventSelection {
	number := len(t.Players) / len(events)

	selections := make([]eventSelection, 0)
	remainder := t.Players
	for _, e := range events {
		sorted := remainder

		switch e.Style {
		case skillfulStage:
			sort.Sort(sort.Reverse(bySkill(sorted)))
		case smartStage:
			sort.Sort(sort.Reverse(bySmarts(sorted)))
		case teamStage:
			selections = append(selections, eventSelection{e, t.Players, true})
			continue
		}

		result := sorted[:number]
		remainder = sorted[number:]
		for i, player := range result {
			seed := false
			if i == 0 {
				seed = true
			}
			selections = append(selections, eventSelection{Event: e, Entrants: []Player{player}, TopSeeds: seed})
		}
	}

	return selections
}
