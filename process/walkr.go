package process

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
)

// LoadGameData from nameFile
func LoadGameData(nameFile string) (gameDataArray []GameData, ok bool) {
	gameDataArray = make([]GameData, 0)

	f, err := os.Open(nameFile)
	if err != nil {
		log.Print(err)
		return nil, false
	}

	r := csv.NewReader(bufio.NewReader(f))
	for {
		record, err := r.Read()
		// Stop at EOF.
		if err == io.EOF {
			break
		}
		if len(record) < 5 {
			continue
		}

		number, _ := strconv.Atoi(record[0])
		planet := record[1]
		planetFile := "resources/planets/default-" + record[4] + "-placeholder@2x.png"
		satelite := record[3]
		sateliteFile := ""

		data := GameData{number, planet, planetFile, satelite, sateliteFile}
		gameDataArray = append(gameDataArray, data)
		ok = true
	}
	return
}

// LoadUserName from idFile
func LoadUserName(idFile string) (playerNameArray []PlayerName, ok bool) {
	f, err := os.Open(idFile)
	if err != nil {
		log.Print(err)
		return nil, false
	}
	r := csv.NewReader(bufio.NewReader(f))
	for {
		record, err := r.Read()
		// Stop at EOF.
		if err == io.EOF {
			break
		}
		if len(record) < 3 {
			continue
		}
		data := PlayerName{record[0], record[1], record[2]}
		playerNameArray = append(playerNameArray, data)
		ok = true
	}
	return
}
