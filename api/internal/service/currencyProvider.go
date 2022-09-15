package service

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/kerrrusha/btc-api/api/internal/config"
	"github.com/kerrrusha/btc-api/api/internal/customErrors"
	"github.com/kerrrusha/btc-api/api/internal/model"
	"github.com/kerrrusha/btc-api/api/internal/utils"
)

type currencyProvider struct {
	baseUrl string
	rateKey string
}

func CreateCurrencyProvider(providerUrl string, currencyRateKey string) *currencyProvider {
	return &currencyProvider{
		baseUrl: providerUrl,
		rateKey: currencyRateKey,
	}
}

func (provider *currencyProvider) RequestJson(baseCurrency string, quoteCurrency string) []byte {
	requestUrl := provider.configureUrl(baseCurrency, quoteCurrency)
	return utils.RequestJson(requestUrl)
}
func (provider *currencyProvider) GetCurrencyRate(baseCurrency string, quoteCurrency string) (int, *customErrors.RequestFailureError) {
	jsonResponse := provider.RequestJson(baseCurrency, quoteCurrency)
	rate, err := provider.castResponse(jsonResponse)

	return int(rate), err
}
func (provider *currencyProvider) GetDomain() string {
	return utils.SubstringBetween(provider.baseUrl, "https://", "/")
}

func (provider *currencyProvider) configureUrl(baseCurrency string, quoteCurrency string) string {
	result := provider.baseUrl
	cfg := config.GetConfig()

	result = strings.ReplaceAll(result, cfg.GetBaseCurrencyMark(), baseCurrency)
	result = strings.ReplaceAll(result, cfg.GetQuoteCurrencyMark(), quoteCurrency)

	return result
}

func (provider *currencyProvider) castResponse(jsonBytes []byte) (float64, *customErrors.RequestFailureError) {
	const INVALID_RETURN_VALUE = -1
	const CAST_ERROR_MESSAGE = "Unsuccessful to unmarshal json Response"
	const STRING_TO_FLOAT_ERROR_MESSAGE = "Unsuccessful parsing string to float64"
	const THIRD_PARTY_ERROR_MESSAGE = "Third-party side API caused error"

	rateStr, unmarshalErr := utils.GetJsonStringValueByKey(jsonBytes, provider.rateKey)
	if unmarshalErr != nil {
		return INVALID_RETURN_VALUE, customErrors.CreateRequestFailureError(CAST_ERROR_MESSAGE)
	}

	rate, ParseFloatErr := strconv.ParseFloat(rateStr, 64)
	if ParseFloatErr != nil {
		return INVALID_RETURN_VALUE, customErrors.CreateRequestFailureError(STRING_TO_FLOAT_ERROR_MESSAGE)
	}

	var errorResponse model.ErrorResponse
	err := json.Unmarshal(jsonBytes, &errorResponse)
	if err != nil {
		return INVALID_RETURN_VALUE, customErrors.CreateRequestFailureError(THIRD_PARTY_ERROR_MESSAGE)
	}

	if len(errorResponse.Error) > 0 {
		return INVALID_RETURN_VALUE, customErrors.CreateRequestFailureError(errorResponse.Error)
	}

	return rate, nil
}
