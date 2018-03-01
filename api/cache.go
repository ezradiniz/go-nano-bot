package api

import (
	"reflect"
	"sync"
	"time"
)

var (
	cache = make(map[string]cacheItem)
	lock  = sync.Mutex{}
)

type cacheItem struct {
	Cachetime float64 // minutes
	Time      time.Time
	Data      interface{}
}

func putCache(key string, value interface{}, cacheTime float64) {
	lock.Lock()
	defer lock.Unlock()
	cache[key] = cacheItem{
		Cachetime: cacheTime,
		Time:      time.Now(),
		Data:      reflect.Indirect(reflect.ValueOf(value)).Interface(),
	}
}

func getCache(key string) (interface{}, bool) {
	lock.Lock()
	defer lock.Unlock()
	item, ok := cache[key]
	if ok {
		return item, true
	}
	return item, false
}
