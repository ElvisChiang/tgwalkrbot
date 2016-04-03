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
			fmt.Printf("#%d: %s/%s %s/%s %s\n",
				data.Number,
				data.Planet, data.PlanetFile,
				data.Satellite, data.SatelliteFile,
				data.Resource)
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
			text := "(nil)"
			if typ == tbotapi.StickerMessage {
				sticker := update.Message.Sticker
				fmt.Printf("\tSticker id: %s size: %d\n",
					sticker.FileBase.ID, sticker.FileBase.Size)
			}
			if typ == tbotapi.TextMessage {
				text = *msg.Text
			}
			fmt.Printf("<-%d, From:\t%s, Text: %s \n", msg.ID, msg.Chat, text)
			if typ != tbotapi.TextMessage {
				//ignore non-text messages for now
				fmt.Println("Ignoring non-text message")
				return
			}
			// Note: Bots cannot receive from channels, at least no text messages. So we don't have to distinguish anything here

			// display the incoming message
			// msg.Chat implements fmt.Stringer, so it'll display nicely
			// we know it's a text message, so we can safely use the Message.Text pointer
			// fmt.Printf("<-%d, From:\t%s, Text: %s \n", msg.ID, msg.Chat, *msg.Text)

			Command(api, &msg.Chat, *msg.Text)
		case tbotapi.InlineQueryUpdate:
			fmt.Println("Ignoring received inline query: ", update.InlineQuery.Query)
		case tbotapi.ChosenInlineResultUpdate:
			fmt.Println("Ignoring chosen inline query result (ID): ", update.ChosenInlineResult.ID)
		default:
			fmt.Printf("Ignoring unknown Update type.")
		}
	}

	// run the bot, this will block
	boilerplate.RunBot(apiToken, updateFunc, "WalkrBot", "Reply Walkr information")
}

func sendPlanetPic(api *tbotapi.TelegramBotAPI, chat *tbotapi.Chat, planet process.GameData) (ok bool) {
	ok = false
	// send a photo
	file, err := os.Open(planet.PlanetFile)
	if err != nil {
		fmt.Printf("Error opening file: %s\n", err)
		ok = false
		return
	}
	defer file.Close()
	photo := api.NewOutgoingPhoto(tbotapi.NewRecipientFromChat(*chat), "planet.png", file)
	captain := fmt.Sprintf("# %d: %s\n命定衛星: %s\n生產資源: %s",
		planet.Number, planet.Planet, planet.Satellite, planet.Resource)
	fmt.Println(captain)
	photo.SetCaption(captain)
	outMsg, err := photo.Send()
	if err != nil {
		fmt.Printf("Error sending photo: %s\n", err)
		return
	}
	fmt.Printf("->%d, To:\t%s, (Photo)\n", outMsg.Message.ID, outMsg.Message.Chat)
	ok = true
	return
}

func sendSatelitePic(api *tbotapi.TelegramBotAPI, chat *tbotapi.Chat, planet process.GameData) (ok bool) {
	ok = false
	// send a photo
	file, err := os.Open(planet.SatelliteFile)
	if err != nil {
		fmt.Printf("Error opening file: %s\n", err)
		ok = false
		return
	}
	defer file.Close()
	photo := api.NewOutgoingPhoto(tbotapi.NewRecipientFromChat(*chat), "satelite.png", file)
	captain := fmt.Sprintf("命定衛星: %s\n對應星球\n # %d: %s\n生產資源: %s",
		planet.Satellite, planet.Number, planet.Planet, planet.Resource)
	if planet.Planet == "-" {
		captain = fmt.Sprintf("命定衛星: %s\n對應星球: 任意",
			planet.Satellite)
	}
	fmt.Println(captain)
	photo.SetCaption(captain)
	outMsg, err := photo.Send()
	if err != nil {
		fmt.Printf("Error sending photo: %s\n", err)
		return
	}
	fmt.Printf("->%d, To:\t%s, (Photo)\n", outMsg.Message.ID, outMsg.Message.Chat)
	ok = true
	return
}

func sendText(api *tbotapi.TelegramBotAPI, chat *tbotapi.Chat, text string) (ok bool) {
	outMsg, err := api.NewOutgoingMessage(tbotapi.NewRecipientFromChat(*chat), text).Send()
	if err != nil {
		fmt.Printf("Error sending text: %s, err = %s\n", text, err)
		return false
	}
	fmt.Printf("->%d, To:\t%s, %s\n", outMsg.Message.ID, outMsg.Message.Chat, text)
	return true
}

func sendSticker(api *tbotapi.TelegramBotAPI, chat *tbotapi.Chat, id string) (ok bool) {
	outMsg, err := api.NewOutgoingStickerResend(tbotapi.NewRecipientFromChat(*chat), id).Send()
	if err != nil {
		fmt.Printf("Error sending sticker: %s, err = %s\n", id, err)
		return false
	}
	fmt.Printf("->%d, To:\t%s, sticker %s\n", outMsg.Message.ID, outMsg.Message.Chat, id)
	return true
}

// Command parse tg command line
func Command(api *tbotapi.TelegramBotAPI, chat *tbotapi.Chat, msg string) (ok bool) {
	ok = false
	result := strings.Fields(msg)
	if len(result) == 0 {
		return
	}
	command := result[0]
	lowerCmd := strings.ToLower(command)
	if lowerCmd == "/help" || lowerCmd == "/start" {
		text := "/wp 星球編號\n/wp 星球名字(模糊搜尋第一個)\n/wr 資源\n" +
			"\n/wp 1\n/wp 地球\n/wr 笑料"
		sendText(api, chat, text)
		// sendSticker(api, chat, "BQADBQADPwAD_HMCBpRNwQvxuQoDAg")
		ok = true
		return
	}

	msg = strings.TrimPrefix(msg, command)
	msg = strings.TrimSpace(msg)
	if len(msg) == 0 {
		return
	}

	switch lowerCmd {
	case "/wp":
		planet, found := process.FindPlanet(gameData, msg)
		if !found {
			text := "醒醒吧，你沒有" + msg
			sendText(api, chat, text)
			fmt.Println(text)
			return
		}
		ok = sendPlanetPic(api, chat, planet)
	case "/wn":
		ok = true
	case "/wr":
		planet, found := process.FindPlanetByResource(gameData, msg)
		if !found {
			text := "沒有生產" + msg + "的星球"
			sendText(api, chat, text)
			fmt.Println(text)
			return
		}
		ok = sendPlanetPic(api, chat, planet)
	case "/wsp": // Find Satelite picture
		satelite, found := process.FindSatellite(gameData, msg)
		if !found {
			text := "醒醒吧，你沒有" + msg
			sendText(api, chat, text)
			fmt.Println(text)
			return
		}
		ok = sendSatelitePic(api, chat, satelite)
	case "/ws":
		planet, found := process.FindPlanetBySatellite(gameData, msg)
		if !found {
			text := "沒有星球喜歡" + msg
			sendText(api, chat, text)
			fmt.Println(text)
			return
		}
		ok = sendPlanetPic(api, chat, planet)
	default:
		fmt.Printf("我看不懂!! %s\n", msg)
	}
	return ok
}
