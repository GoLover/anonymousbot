package commands

import (
	tb "gopkg.in/tucnak/telebot.v2"
)
var handlerBot *tb.Bot
func InitHandler(bot *tb.Bot) {
	handlerBot = bot
	handlerBot.Handle("/start", Start)
	handlerBot.Handle("/stop", Stop)
	handlerBot.Handle(tb.OnText,TextMessageHandle)
}
