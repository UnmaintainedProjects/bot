package handlers

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
)

var commandStartHandler = handlers.NewCommand("start", commandStart)

func commandStart(b *gotgbot.Bot, ctx *ext.Context) error {
	_, err := ctx.Message.Reply(b, "Hey!", nil)
	return err
}
