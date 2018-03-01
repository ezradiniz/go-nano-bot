package api

import (
	"fmt"
)

// NanoResponse response
type NanoResponse []Coin

// Coin response coin
type Coin struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	Symbol           string `json:"symbol"`
	Rank             string `json:"rank"`
	PriceUSD         string `json:"price_usd"`
	PriceBTC         string `json:"price_btc"`
	Volume24h        string `json:"24h_volume_usd"`
	MarketCap        string `json:"market_cap_usd"`
	AvailableSupply  string `json:"available_supply"`
	TotalSupply      string `json:"total_supply"`
	MaxSupply        string `json:"max_supply"`
	PercentChange1h  string `json:"percent_change_1h"`
	PercentChange24h string `json:"percent_change_24h"`
	PercentChange7d  string `json:"percent_change_7d"`
	LastUpdated      string `json:"last_updated"`
}

// FetchNano fetch coin market cap
func FetchNano(response interface{}) error {
	return fetchJSON("https://api.coinmarketcap.com/v1/ticker/nano/", response, 1)
}

func check(f func() error) {
	if err := f(); err != nil {
		fmt.Println("Error when closing:", err)
	}
}
