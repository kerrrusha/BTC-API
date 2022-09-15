package config

func GetCurrencyProviderNameArray() []string {
	return []string{
		cfg.GetEnvironmentVarBinanceProviderName(),
		cfg.GetEnvironmentVarCoinapiProviderName(),
	}
}
func GetDefaultCurrencyProviderName() string {
	return cfg.GetEnvironmentVarBinanceProviderName()
}
func CurrencyProviderNameExists(name string) bool {
	names := GetCurrencyProviderNameArray()
	for _, s := range names {
		if name == s {
			return true
		}
	}
	return false
}
