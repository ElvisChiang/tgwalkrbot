package main

import (
	"fmt"

	"bitbucket.org/mrd0ll4r/tbotapi"
	"bitbucket.org/mrd0ll4r/tbotapi/examples/boilerplate"
)

// show verbose debug msg
const DEBUG = true

func main() {
	playerData, ok := loadName()
	if !ok {
		fmt.Printf("Player data loading fail\n")
		return
	}
	gameData, ok := loadGameData()
	if !ok {
		fmt.Printf("Game data loading fail\n")
		return
	}
	if DEBUG {
		for i, data := range playerData {
			fmt.Printf("%d cn:%s tg:`%s` walkr:`%s`\n", i,
				data.codeName, data.tgName, data.walkrName)
		}

		for _, data := range gameData {
			fmt.Printf("#%d: %s/%s %s/%s\n",
				data.number,
				data.planet, data.planetFile,
				data.satelite, data.sateliteFile)
		}
	}
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

			// now simply echo that back
			outMsg, err := api.NewOutgoingMessage(tbotapi.NewRecipientFromChat(msg.Chat), *msg.Text).Send()

			if err != nil {
				fmt.Printf("Error sending: %s\n", err)
				return
			}
			fmt.Printf("->%d, To:\t%s, Text: %s\n", outMsg.Message.ID, outMsg.Message.Chat, *outMsg.Message.Text)
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
