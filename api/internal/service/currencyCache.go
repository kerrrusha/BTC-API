package service

import (
	"sync"
	"time"

	"github.com/bluele/gcache"
	"github.com/kerrrusha/btc-api/api/internal/customErrors"
)

const (
	CACHE_SIZE            = 10
	CACHE_EXPIRATION_TIME = 300 * time.Second
	INVALID_RESULT        = -1
	DEFAULT_KEY           = 0
)

type currencyCache struct {
	currencyRate gcache.Cache
}

var lockCacheCreation = &sync.Mutex{}

var cache *currencyCache

func GetCurrencyCache() *currencyCache {
	if cache != nil {
		return cache
	}

	tryCreateCurrencyCacheSingleton()

	return cache
}

func tryCreateCurrencyCacheSingleton() {
	lockCacheCreation.Lock()
	defer lockCacheCreation.Unlock()
	if cache != nil {
		return
	}

	createCurrencyCache()
}
func createCurrencyCache() {
	cache = &currencyCache{
		currencyRate: gcache.New(CACHE_SIZE).Expiration(CACHE_EXPIRATION_TIME).ARC().Build(),
	}
}

func (cache *currencyCache) Set(rate int) error {
	return cache.currencyRate.SetWithExpire(DEFAULT_KEY, rate, CACHE_EXPIRATION_TIME)
}

func (cache *currencyCache) IsEmpty() bool {
	return cache.currencyRate.Len(false) == 0
}

func (cache *currencyCache) Get() (int, *customErrors.RateNotInCacheError) {
	ERROR_MESSAGE := "Rate not in cache."

	rate, err := cache.currencyRate.Get(DEFAULT_KEY)
	if err != nil {
		return INVALID_RESULT, customErrors.CreateRateNotInCacheError(ERROR_MESSAGE)
	}

	return rate.(int), nil
}

func (cache *currencyCache) Clear() {
	cache.currencyRate.Remove(DEFAULT_KEY)
}
