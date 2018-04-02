package api

import (
	"reflect"
	"strings"

	"github.com/gocolly/colly"
)

// Block response
type Block struct {
	Value string `json:"block"`
}

// Nanode nanode.co URL
const Nanode = "https://www.nanode.co"

// FetchBlock fetch nano blocks
func FetchBlock(response interface{}) error {
	c := colly.NewCollector()

	c.OnHTML("p[class='title'] > a[href='/blocks']", func(e *colly.HTMLElement) {
		text := strings.TrimSpace(e.Text)
		reflect.ValueOf(response).Elem().FieldByName("Value").SetString(text)
	})

	err := c.Visit(Nanode)
	return err
}
