package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"time"
)

func fetchJSON(url string, response interface{}, cacheTime float64) error {
	if res, ok := getCache(url); ok != false {
		item := res.(cacheItem)
		if time.Now().Sub(item.Time).Minutes() < item.Cachetime {
			reflect.ValueOf(response).Elem().Set(reflect.ValueOf(item.Data))
			return nil
		}
	}

	res, err := http.Get(url)
	if err != nil {
		return err
	}

	defer check(res.Body.Close)
	err = json.NewDecoder(res.Body).Decode(response)

	if err == nil {
		putCache(url, response, cacheTime)
	}

	return err
}

func check(f func() error) {
	if err := f(); err != nil {
		fmt.Println("Error when closing:", err)
	}
}
