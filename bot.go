package main

import (
	"fmt"
	"os"
	"strings"

	"./process"

	"bitbucket.org/mrd0ll4r/tbotapi"
	"bitbucket.org/mrd0ll4r/tbotapi/examples/boilerplate"
)

// show verbose debug msg
const DEBUG = true

var idFile = "./data/id.csv"
var planetFile = "./data/planet.csv"

var playerData []process.PlayerName
var gameData []process.GameData

func main() {
	ok := false
	playerData, ok = process.LoadUserName(idFile)
	if !ok {
		fmt.Printf("Player data loading fail\n")
		return
	}
	gameData, ok = process.LoadGameData(planetFile)
	if !ok {
		fmt.Printf("Game data loading fail\n")
		return
	}
	if DEBUG {
		for i, data := range playerData {
			fmt.Printf("%d cn:%s tg:`%s` walkr:`%s`\n", i,
				data.CodeName, data.TgName, data.WalkrName)
		}

		for _, data := range gameData {
			fmt.Printf("#%d: %s/%s %s/%s\n",
				data.Number,
				data.Planet, data.PlanetFile,
				data.Satelite, data.SateliteFile)
		}
	}
	startBot()
}

func startBot() {
	updateFunc := func(update tbotapi.Update, api *tbotapi.TelegramBotAPI) {
		switch update.Type() {
		case tbotapi.MessageUpdate:
			msg := update.Message
			typ := msg.Type()
			if typ != tbotapi.TextMessage {
				//ignore non-text messages for now
				fmt.Println("Ignoring non-text message")
				return
			}
			// Note: Bots cannot receive from channels, at least no text messages. So we don't have to distinguish anything here

			// display the incoming message
			// msg.Chat implements fmt.Stringer, so it'll display nicely
			// we know it's a text message, so we can safely use the Message.Text pointer
			fmt.Printf("<-%d, From:\t%s, Text: %s \n", msg.ID, msg.Chat, *msg.Text)

			ok := Command(api, &msg.Chat, *msg.Text)

			if !ok {
				fmt.Printf("Cannot process input command\n")
				return
			}
		case tbotapi.InlineQueryUpdate:
			fmt.Println("Ignoring received inline query: ", update.InlineQuery.Query)
		case tbotapi.ChosenInlineResultUpdate:
			fmt.Println("Ignoring chosen inline query result (ID): ", update.ChosenInlineResult.ID)
		default:
			fmt.Printf("Ignoring unknown Update type.")
		}
	}

	// run the bot, this will block
	boilerplate.RunBot(apiToken, updateFunc, "Echo", "Echoes messages back")
}

// Command parse tg command line
func Command(api *tbotapi.TelegramBotAPI, chat *tbotapi.Chat, msg string) (ok bool) {
	ok = false
	result := strings.Fields(msg)
	if len(result) == 0 {
		return
	}
	command := result[0]

	switch command {
	case "/wp":
		msg = strings.TrimPrefix(msg, "/wp")
		msg = strings.TrimSpace(msg)
		fileName, _, found := process.FindPlanet(gameData, msg)
		if !found {
			fmt.Printf("Planet %s not found\n", msg)
			return
		}
		// send a photo
		file, err := os.Open(fileName)
		if err != nil {
			fmt.Printf("Error opening file: %s\n", err)
			ok = false
			return
		}
		defer file.Close()
		outMsg, err := api.NewOutgoingPhoto(tbotapi.NewRecipientFromChat(*chat), "planet.png", file).Send()
		if err != nil {
			fmt.Printf("Error sending: %s\n", err)
			return
		}
		fmt.Printf("->%d, To:\t%s, (Photo)\n", outMsg.Message.ID, outMsg.Message.Chat)
		ok = true
		return
	case "/ws":
		msg = strings.TrimPrefix(msg, "/ws")
		msg = strings.TrimSpace(msg)
	case "/wn":
		msg = strings.TrimPrefix(msg, "/wn")
		msg = strings.TrimSpace(msg)
		ok = true
	default:
		fmt.Printf("Cannot process %s", msg)
	}
	return ok
}
