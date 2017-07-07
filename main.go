package main

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Action = func(ctx *cli.Context) error {
		data, err := ioutil.ReadFile("./players.json")
		if err != nil {
			return err
		}

		teams := make(map[string]team)
		err = json.Unmarshal(data, &teams)
		if err != nil {
			return err
		}

		data, err = ioutil.ReadFile("./stages.json")
		if err != nil {
			return err
		}

		stages := make([]stage, 0)
		err = json.Unmarshal(data, &stages)
		if err != nil {
			return err
		}

		overall := newScores()

		for _, s := range stages {
			allSelections := make([]eventSelection, 0)
			for _, t := range teams {
				allSelections = append(allSelections, t.SelectPlayers(s.Events...)...)
			}

			for _, e := range s.Events {
				selections := getSelections(e, true, allSelections)
				standings := e.Play(selections.Entrants...)
				overall.StageResults(e, 10, standings)
				logrus.Infof("Standings in %s for top seeds", e.Event)
				printStageResults(standings)
			}

			for _, e := range s.Events {
				selections := getSelections(e, false, allSelections)
				standings := e.Play(selections.Entrants...)
				overall.StageResults(e, 0, standings)
				logrus.Infof("Standings in %s for second seeds", e.Event)
				printStageResults(standings)
			}
		}

		overall.print()
		return nil
	}

	app.Run(os.Args)
}

func getSelections(e event, seeds bool, selections []eventSelection) eventSelection {
	result := eventSelection{Event: e, TopSeeds: seeds}
	for _, v := range selections {
		if v.Event == e && v.TopSeeds == seeds {
			result.Entrants = append(result.Entrants, v.Entrants...)
		}
	}
	return result
}
