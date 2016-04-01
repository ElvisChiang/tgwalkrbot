package main

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"

	"./process"
)

const idFile = "id.csv"
const resourceMapping = "name.csv"

func loadGameData() (gameDataArray []process.GameData, ok bool) {
	gameDataArray = make([]process.GameData, 0)

	f, err := os.Open(resourceMapping)
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

		data := process.GameData{number, planet, planetFile, satelite, sateliteFile}
		gameDataArray = append(gameDataArray, data)
		ok = true
	}
	return
}

func loadName() (playerNameArray []process.PlayerName, ok bool) {
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
		data := process.PlayerName{record[0], record[1], record[2]}
		playerNameArray = append(playerNameArray, data)
		ok = true
	}
	return
}
