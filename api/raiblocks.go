package api

import (
	"reflect"
	"strings"

	"github.com/gocolly/colly"
)

// Balance response
type Balance struct {
	Amount string `json:"amount"`
}

// Raiblocks raiblocks (nano) URL
const Raiblocks = "https://raiblocks.net/account/index.php"

// FetchBalance fetch user balance from raiblocks.net
func FetchBalance(address string, response interface{}) error {
	c := colly.NewCollector()

	c.OnHTML(".npage > .row:nth-of-type(3) strong", func(e *colly.HTMLElement) {
		text := strings.TrimSpace(e.Text)
		reflect.ValueOf(response).Elem().FieldByName("Amount").SetString(text)
	})

	err := c.Visit(Raiblocks + "?acc=" + address)

	return err
}
