package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"time"
)

var (
	cache     interface{}
	timestamp time.Time
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
	if cache != nil && time.Now().Sub(timestamp).Minutes() < 5 {
		reflect.ValueOf(response).Elem().Set(reflect.ValueOf(cache))
		return nil
	}
	res, err := http.Get("https://api.coinmarketcap.com/v1/ticker/nano/")
	if err != nil {
		return err
	}
	defer check(res.Body.Close)
	err = json.NewDecoder(res.Body).Decode(response)

	cache = reflect.Indirect(reflect.ValueOf(response)).Interface()
	timestamp = time.Now()

	return err
}

func check(f func() error) {
	if err := f(); err != nil {
		fmt.Println("Error when closing:", err)
	}
}
