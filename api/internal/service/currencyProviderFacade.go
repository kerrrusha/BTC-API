package service

import (
	"github.com/kerrrusha/btc-api/api/internal/customErrors"
	"log"
)

type currencyProviderFacade struct {
	provider *currencyProvider
	cache    *currencyCache
}

func CreateCurrencyProviderFacade(provider *currencyProvider, cache *currencyCache) *currencyProviderFacade {
	return &currencyProviderFacade{provider: provider, cache: cache}
}

func (providerFacade *currencyProviderFacade) GetCurrencyRate(baseCurrency string, quoteCurrency string) (int, *customErrors.RequestFailureError) {
	cachedRate, absentErr := providerFacade.cache.Get()
	if absentErr == nil {
		return cachedRate, nil
	}

	jsonResponse := providerFacade.provider.RequestJson(baseCurrency, quoteCurrency)
	providerFacade.logProviderResponse(jsonResponse)

	rate, err := providerFacade.provider.castResponse(jsonResponse)
	if err != nil {
		return int(rate), err
	}

	cacheSetErr := providerFacade.cache.Set(int(rate))
	if cacheSetErr != nil {
		return INVALID_RESULT, nil
	}

	return int(rate), err
}
func (providerFacade *currencyProviderFacade) logProviderResponse(response []byte) {
	responseMsg := "Response from " + providerFacade.provider.GetDomain() + ": " + string(response)
	log.Println(responseMsg)
}
