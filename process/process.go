package process

import (
	"fmt"
	"regexp"
	"strings"
)

// Command parse tg command line
func Command(playerData []PlayerName, gameData []GameData, msg string) (output string, ok bool) {
	output = ""
	ok = false
	result := strings.Fields(msg)
	if len(result) == 0 {
		return
	}
	command := result[0]

	switch command {
	case "/wp":
		output, ok = findPlanet(gameData, msg)
	case "/ws":
	case "/wn":
		ok = true
	default:
		fmt.Printf("Cannot process %s", msg)
	}
	return output, ok
}

func findPlanet(gameData []GameData, msg string) (output string, ok bool) {
	output = ""
	ok = false
	for _, data := range gameData {
		fmt.Printf("#%d: %s/%s %s/%s\n",
			data.Number,
			data.Planet, data.PlanetFile,
			data.Satelite, data.SateliteFile)
	}
	return output, ok
}
