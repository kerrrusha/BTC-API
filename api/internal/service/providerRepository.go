package service

import (
	"os"
	"sync"

	"github.com/kerrrusha/btc-api/api/internal/config"
	"github.com/kerrrusha/btc-api/api/internal/customErrors"
)

type providerRepository struct {
	current *currencyProviderChain
}

var lockRepoCreation = &sync.Mutex{}

var repo *providerRepository

func GetProviderRepository() *providerRepository {
	if repo != nil {
		return repo
	}

	TryInitProviderRepositorySingleton()

	return repo
}
func getMainCurrencyProviderName() string {
	providerName, presented := os.LookupEnv("CRYPTO_CURRENCY_PROVIDER")
	if !presented || !config.CurrencyProviderNameExists(providerName) {
		return config.GetDefaultCurrencyProviderName()
	}
	return providerName
}
func TryInitProviderRepositorySingleton() {
	lockRepoCreation.Lock()
	defer lockRepoCreation.Unlock()
	if repo != nil {
		return
	}

	createProviderRepository()

	initialiseProviderRepository()
}
func createProviderRepository() {
	repo = &providerRepository{}
}
func initialiseProviderRepository() {
	cfg := config.GetConfig()

	coinapiProvider := CreateCurrencyProvider(cfg.GetCoinapiUrl(), cfg.GetCoinapiRateKey())
	coinapiChain := CreateCurrencyProviderChain(coinapiProvider)

	binanceProvider := CreateCurrencyProvider(cfg.GetBinanceUrl(), cfg.GetBinanceRateKey())
	binanceChain := CreateCurrencyProviderChain(binanceProvider)

	mainProviderName := getMainCurrencyProviderName()
	if mainProviderName == cfg.GetEnvironmentVarBinanceProviderName() {
		repo.addCurrencyProviderChain(binanceChain)
		repo.addCurrencyProviderChain(coinapiChain)
	}
	if mainProviderName == cfg.GetEnvironmentVarCoinapiProviderName() {
		repo.addCurrencyProviderChain(coinapiChain)
		repo.addCurrencyProviderChain(binanceChain)
	}
}

func (c *providerRepository) GetCurrencyProvider() (*currencyProvider, *customErrors.CurrencyProviderChainAreOverError) {
	if c.current.IsEmpty() {
		return nil, customErrors.CreateCurrencyProviderChainAreOverError("Current provider chain is empty.")
	}
	return c.current.GetCurrencyProvider(), nil
}
func (c *providerRepository) isEmpty() bool {
	return c.current == nil
}
func (c *providerRepository) getLastCurrencyProviderChain() *currencyProviderChain {
	if c.isEmpty() {
		return nil
	}
	lastChain := c.current
	for lastChain.next != nil {
		lastChain = lastChain.next
	}
	return lastChain
}
func (c *providerRepository) addCurrencyProviderChain(nextChain *currencyProviderChain) {
	if c.isEmpty() {
		c.current = nextChain
		return
	}
	c.getLastCurrencyProviderChain().next = nextChain
}
