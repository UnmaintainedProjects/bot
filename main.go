package main

import (
	"os"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	token := os.Getenv("TOKEN")
	bot, err := gotgbot.NewBot(token, nil)
	if err != nil {
		panic(err)
	}
	updater := ext.NewUpdater(nil)
	err = updater.StartPolling(
		bot, &ext.PollingOpts{
			// DropPendingUpdates: true,
		},
	)
	if err != nil {
		panic(err)
	}
	updater.Idle()
}
