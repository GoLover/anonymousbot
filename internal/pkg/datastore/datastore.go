package datastore

import (
	configParser "github.com/golover/anonymousbot/internal/pkg/configparser"
	"github.com/golover/anonymousbot/internal/pkg/datastore/inmemory"
	"github.com/golover/anonymousbot/internal/pkg/datastore/redis"
)

type DSInterface interface {
	Get(string) string
	Set(string, string)
}
var selectedDS DSInterface
func InitDataStore() {
	if configParser.GetInstance().DataStore == "inmemory" {
		selectedDS = inmemory.InitDataStore()
	} else {
		selectedDS = redis.InitDataStore()
	}
}

func Get(key string) string {
	return selectedDS.Get(key)
}

func Set(key,value string) {
	selectedDS.Set(key,value)
}