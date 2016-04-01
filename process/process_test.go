package process

import (
	"fmt"
	"testing"
)

var idFile = "../data/id.csv"
var planetFile = "../data/planet.csv"

var playerData []PlayerName
var gameData []GameData

func TestCommand(t *testing.T) {
	pcases := []struct {
		in     string
		result bool
	}{
		{"1", true},
		{"隱形基地", true},
		{"瓦肯", false},
		{"水", true},
	}
	ncases := []struct {
		in     string
		result bool
	}{
		//		{"/wn @elvisfb"},
		{"/wn \\T'Elvis Chiang' tg name \\W'Nicole Lai' walkr name @elvisfb codename", true},
		{"/wn \\W'Elvis' walkr name", true},
	}

	playerData, ok := LoadUserName(idFile)
	if !ok {
		t.Errorf("Player data loading fail")
		return
	}
	gameData, ok := LoadGameData(planetFile)
	if !ok {
		t.Errorf("Game data loading fail")
		return
	}
	// for Debug dumping
	if false {
		for _, data := range gameData {
			fmt.Printf("%d,%s,%s,%s,%s\n",
				data.Number, data.Planet, data.PlanetFile,
				data.Satelite, data.Resource)
		}
	}
	fmt.Println("----------------")
	for _, c := range pcases {
		planet, ok := FindPlanet(gameData, c.in)
		if ok != c.result {
			t.Errorf("cannot process " + c.in)
			continue
		}
		fmt.Printf("%s -> `%s`\n", c.in, planet.Planet)
		fmt.Println("----------------")
	}
	for _, c := range ncases {
		processName(playerData, c.in)
	}
}
