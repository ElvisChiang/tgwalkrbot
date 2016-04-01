package process

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// FindPlanet get planet picture from number or name
func FindPlanet(gameData []GameData, msg string) (file string, name string, ok bool) {
	file = ""
	name = ""
	ok = false
	num, _ := strconv.Atoi(msg)
	if num > 0 {
		fmt.Printf("finding %d planet\n", num)
		for _, data := range gameData {
			if data.Number == num {
				name = data.Planet
				file = data.PlanetFile
				ok = true
				return
			}
		}
		return
	}
	for _, data := range gameData {
		if data.Planet == msg {
			name = data.Planet
			file = data.PlanetFile
			ok = true
			return
		}
	}
	return
}

func findUserName(codename, tgName, walkrName string) (output string, ok bool) {
	ok = false

	return
}

// TODO
func processName(playerData []PlayerName, msg string) (output string) {
	msg = strings.TrimPrefix(msg, "/wn")
	msg = strings.TrimSpace(msg)
	re := regexp.MustCompile("\\\\W'.+?'")
	for {
		loc := re.FindStringIndex(msg)
		if loc == nil {
			break
		}
		fmt.Printf("tg name index = %d\n", re.FindStringIndex(msg))
		tag := re.FindString(msg)
		fmt.Printf("tg name string = %s\n", tag)
		replace, ok := findUserName("", "", tag)
		if !ok {
			fmt.Printf("walkr name cannot found %s\n", tag)
			break
		}
		msg = strings.Replace(msg, tag, replace, 1)
	}

	return msg
}
