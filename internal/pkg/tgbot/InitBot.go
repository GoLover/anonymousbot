package tgbot

import (
	configParser "github.com/golover/anonymousbot/internal/pkg/configparser"
	"github.com/golover/anonymousbot/internal/pkg/tgbot/commands"
	tb "gopkg.in/tucnak/telebot.v2"
	"time"
)
var botInstance *tb.Bot
func InitBot(webhookMode bool) *tb.Bot{
	if botInstance != nil {
		return botInstance
	}
	var poller tb.Poller
	if webhookMode {
		poller = &tb.Webhook{
			Listen:         configParser.GetInstance().WebhookPort,
			MaxConnections: 16,
			AllowedUpdates: nil,
			HasCustomCert:  false,
			PendingUpdates: 0,
			ErrorUnixtime:  0,
			ErrorMessage:   "",
			TLS:            nil,
			Endpoint:       nil,
		}
	} else {
		poller = &tb.LongPoller{Timeout: 10 * time.Second}
	}
	botInstance, err := tb.NewBot(tb.Settings{
		Token:       configParser.GetInstance().BotToken,
		Poller: poller,
		})

	if err != nil {
		panic(err)
		return nil
	}
	commands.InitHandler(botInstance)
	return botInstance
}

func GetInstance() *tb.Bot {
	if botInstance != nil {
		return botInstance
	}
	// switch to longpolling mode, this mode will work fine.
	return InitBot(false)
}