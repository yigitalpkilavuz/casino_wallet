package caching

import "time"

type ICache interface {
	Set(key string, value string, expiration time.Duration) error
	Get(key string) (string, error)
	Clear(key string) error
	Exists(key string) (bool, error)
}
