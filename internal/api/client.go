package api

import (
	"net/http"
	"time"
	"github.com/alaw22/Gokedex/internal/pokecache"
)

type Client struct {
	httpClient http.Client
	gokeCache  *pokecache.Cache
}


func NewClient (timeout, interval time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		gokeCache: pokecache.NewCache(interval),
	}
}