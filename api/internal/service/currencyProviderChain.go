package service

type currencyProviderChain struct {
	provider *currencyProvider
	next     *currencyProviderChain
}

func CreateCurrencyProviderChain(pr *currencyProvider) *currencyProviderChain {
	return &currencyProviderChain{provider: pr}
}

func (c *currencyProviderChain) GetCurrencyProvider() *currencyProvider {
	return c.provider
}

func (c *currencyProviderChain) SetNext(nextChain *currencyProviderChain) {
	c.next = nextChain
}

func (c *currencyProviderChain) IsEmpty() bool {
	return c.provider == nil
}
