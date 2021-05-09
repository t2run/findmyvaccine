package main

import (
	"fmt"
	"net/http"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters"

	handlersc "t2findmyvaccinebot/pkg/handlers"
)

func main() {
	//replace Token with your bot Token
	b, err := gotgbot.NewBot("your-token", &gotgbot.BotOpts{
		Client:      http.Client{},
		GetTimeout:  gotgbot.DefaultGetTimeout,
		PostTimeout: gotgbot.DefaultPostTimeout,
	})
	if err != nil {
		panic("failed to create new bot: " + err.Error())
	}

	// Create updater and dispatcher.
	updater := ext.NewUpdater(nil)
	dispatcher := updater.Dispatcher

	// Add echo handler to reply to all messages.
	dispatcher.AddHandler(handlers.NewCommand("start", handlersc.Start))
	//return states and code
	dispatcher.AddHandler(handlers.NewCommand("states", handlersc.States))
	dispatcher.AddHandler(handlers.NewCallback(filters.Equal("start_callback"), handlersc.StartCB))
	//main handle when Bot is Active
	dispatcher.AddHandler(handlers.NewMessage(filters.All, handlersc.HandleAll))

	// Start receiving updates.
	err = updater.StartPolling(b, &ext.PollingOpts{DropPendingUpdates: true})
	if err != nil {
		panic("failed to start polling: " + err.Error())
	}
	fmt.Printf("%s has been started...\n", b.User.Username)

	// Idle, to keep updates coming in, and avoid bot stopping.
	updater.Idle()
}
