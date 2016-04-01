package process

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// FindPlanetByResource get planet from resouce name
func FindPlanetByResource(gameData []GameData, msg string) (planet GameData, ok bool) {
	ok = false
	for _, data := range gameData {
		if strings.Contains(data.Resource, msg) {
			planet = data
			ok = true
			return
		}
	}

	return
}

// FindPlanet get planet picture from number or name
func FindPlanet(gameData []GameData, msg string) (planet GameData, ok bool) {
	ok = false
	num, _ := strconv.Atoi(msg)
	if num > 0 {
		fmt.Printf("finding for planet #%d\n", num)
		for _, data := range gameData {
			if data.Number == num {
				planet = data
				ok = true
				return
			}
		}
		return
	}
	if len(msg) == 0 {
		return
	}
	fmt.Printf("finding for planet `%s`\n", msg)
	for _, data := range gameData {
		if strings.Contains(data.Planet, msg) {
			planet = data
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
