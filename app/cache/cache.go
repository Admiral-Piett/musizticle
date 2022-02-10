package app_cache

import (
	"github.com/Admiral-Piett/musizticle/app/utils"
	"github.com/patrickmn/go-cache"
	"net/http"
	"time"
)

type AppCache struct {
	client *cache.Cache
}

type User struct {
	Username string `json:"name"`
	Password string `json:"name"`
}

func InitCache() *AppCache {
	expiration := time.Duration(utils.CACHE_EXPIRATION_MINUTES) * time.Minute
	cleanup := time.Duration(utils.CACHE_CLEANUP_MINUTES) * time.Minute
	return &AppCache{client: cache.New(expiration, cleanup)}
}

func (a *AppCache) ValidateSession(w http.ResponseWriter, r *http.Request) error {

	return nil
}

func (a *AppCache) Get() {

}

// Shadow cache.Cache's Add method since it will not overwrite a valid/non-expired entry in the cache.
func (a *AppCache) Add() {

}
