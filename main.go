package main

import (
	configParser "github.com/golover/anonymousbot/internal/pkg/configparser"
	"github.com/golover/anonymousbot/internal/pkg/datastore"
	"github.com/golover/anonymousbot/internal/pkg/tgbot"
)

func init() {
	configParser.ParseConfig()
}

func main() {
	datastore.InitDataStore()
	botInstance := tgbot.InitBot(configParser.GetInstance().WebhookMode)
	botInstance.Start()
}