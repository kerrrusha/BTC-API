package rest

import (
	"encoding/json"
	"github.com/kerrrusha/btc-api/api/internal/customErrors"
	"log"
	"net/http"

	"github.com/kerrrusha/btc-api/api/internal/config"
	"github.com/kerrrusha/btc-api/api/internal/model"
	"github.com/kerrrusha/btc-api/api/internal/service"
	"github.com/kerrrusha/btc-api/api/internal/utils"
)

func Rate(w http.ResponseWriter, r *http.Request) {
	log.Println("rate endpoint")

	rate, err := tryToGetRate()
	if err != nil {
		utils.SendResponse(w, model.ErrorResponse{Error: err.GetMessage()}, http.StatusBadRequest)
		return
	}

	response := model.RateValue{Rate: uint32(rate)}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	utils.CheckForError(json.NewEncoder(w).Encode(response))
}

func tryToGetRate() (int, *customErrors.CustomError) {
	INVALID_RESULT := -1

	provider, emptyRepoErr := service.GetProviderRepository().GetCurrencyProvider()
	if emptyRepoErr != nil {
		return INVALID_RESULT, customErrors.CreateCustomError(emptyRepoErr.GetMessage())
	}
	cache := service.GetCurrencyCache()
	providerFacade := service.CreateCurrencyProviderFacade(provider, cache)

	cfg := config.GetConfig()
	rate, requestFailErr := providerFacade.GetCurrencyRate(cfg.GetBaseCurrency(), cfg.GetQuoteCurrency())

	if requestFailErr != nil {
		return INVALID_RESULT, customErrors.CreateCustomError(requestFailErr.GetMessage())
	}

	return rate, nil
}
