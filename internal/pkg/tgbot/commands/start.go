package commands

import (
	configParser "github.com/golover/anonymousbot/internal/pkg/configparser"
	"github.com/golover/anonymousbot/internal/pkg/usermanagement"
	tb "gopkg.in/tucnak/telebot.v2"
	"strconv"
)

func Start(m *tb.Message) {
	if !m.Private() {
		return
	}
	if m.Payload == "" {
		uniqueId := usermanagement.AddUser(strconv.Itoa(m.Sender.ID))
		handlerBot.Send(m.Sender, "Welcome to Anonymous Message bot, You can share this link to others to receive messages."+
			"\nYour links is: https://t.me/"+configParser.GetInstance().BotUsername+"?start="+uniqueId)
		return
	}
	contactUniqueId := usermanagement.LookupUserByUniqueId(m.Payload)
	if contactUniqueId != "" {
		usermanagement.SetCurrentContact(strconv.Itoa(m.Sender.ID),contactUniqueId)
		handlerBot.Send(m.Sender, "You are currently send message to your receptor, feel free.")
		return
	}
}
