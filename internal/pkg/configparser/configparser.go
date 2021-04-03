package configParser
import (
	"encoding/json"
	"github.com/joho/godotenv"
	"io/ioutil"
	"os"
)

type Configuration struct {
	BotUsername string `json:"BotUsername"`
	BotToken string `json:"BotToken"`
	WebhookPort string `json:"ListenAddressPort"`
	WebhookMode bool `json:"WebhookMode"`
	DataStore string `json:"DataStore"`
}
var configurationInstance *Configuration = nil
func ParseConfig() Configuration{
	if configurationInstance != nil {
		return *configurationInstance
	}
	err := godotenv.Load()
	if err != nil {
		configInBytes, secondErr := ioutil.ReadFile("config.json")
		if secondErr != nil {
			panic("cannot find configuration file")
		}
		secondErr = json.Unmarshal(configInBytes,configurationInstance)
		if secondErr != nil {
			panic("config isn't valid")
		}
	}
	whmode := true
	if os.Getenv("WebhookMode") == "" || os.Getenv("WebhookMode") == "false" {
		whmode = false
	}
	configurationInstance = &Configuration{
		BotUsername: os.Getenv("BotUsername"),
		BotToken:    os.Getenv("BotToken"),
		WebhookPort: os.Getenv("ListenAddressPort"),
		WebhookMode: whmode,
		DataStore: os.Getenv("DataStore"),
	}
	return *configurationInstance
}

func GetInstance() Configuration {
	if configurationInstance != nil {
		return *configurationInstance
	}
	return ParseConfig()
}