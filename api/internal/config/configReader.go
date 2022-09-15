package config

import (
	"encoding/json"
	"strconv"
	"sync"

	"github.com/kerrrusha/btc-api/api/internal/model/dataStorage/fileStorage"
	"github.com/kerrrusha/btc-api/api/internal/utils"
)

const PROJECT_FOLDER_NAME = "btc-api"
const CONFIG_FILEPATH = "config.json"

type config struct {
	data map[string]json.RawMessage
}

var lock = &sync.Mutex{}

var cfg *config

func GetConfig() *config {
	if cfg != nil {
		return cfg
	}

	TryInitConfigSingleton()

	return cfg
}

func TryInitConfigSingleton() {
	lock.Lock()
	defer lock.Unlock()
	if cfg == nil {
		createConfig()
	}
}

func createConfig() {
	path := utils.GetProjPath(PROJECT_FOLDER_NAME) + CONFIG_FILEPATH
	reader := fileStorage.CreateFileReader(path)

	jsonBytes := reader.Read()
	jsonMap := make(map[string]json.RawMessage)

	err := json.Unmarshal(jsonBytes, &jsonMap)
	if err != nil {
		panic(err)
	}

	cfg = &config{data: jsonMap}
}

func (c *config) GetEmailsFilepath() string {
	return toString(c.data["emailsFilepath"])
}
func (c *config) GetBaseCurrency() string {
	return toString(c.data["baseCurrency"])
}
func (c *config) GetBaseCurrencyMark() string {
	return toString(c.data["baseCurrencyMark"])
}
func (c *config) GetQuoteCurrency() string {
	return toString(c.data["quoteCurrency"])
}
func (c *config) GetQuoteCurrencyMark() string {
	return toString(c.data["quoteCurrencyMark"])
}
func (c *config) GetCoinapiUrl() string {
	return toString(c.data["coinapiUrl"])
}
func (c *config) GetCoinapiRateKey() string {
	return toString(c.data["coinapiRateKey"])
}
func (c *config) GetBinanceUrl() string {
	return toString(c.data["binanceUrl"])
}
func (c *config) GetBinanceRateKey() string {
	return toString(c.data["binanceRateKey"])
}
func (c *config) GetEnvironmentVarBinanceProviderName() string {
	return toString(c.data["environmentVarBinanceProviderName"])
}
func (c *config) GetEnvironmentVarCoinapiProviderName() string {
	return toString(c.data["environmentVarCoinapiProviderName"])
}

func toString(bytes []byte) string {
	return utils.RemoveRedundantGaps(string(bytes))
}
func toInt(bytes []byte) int {
	number, err := strconv.Atoi(string(bytes))
	if err != nil {
		panic(err)
	}
	return number
}
