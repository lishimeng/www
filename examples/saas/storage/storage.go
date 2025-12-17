package storage

import (
	"github.com/lishimeng/x/container"
	"time"
)

type Storage interface {
	Save(key string, value any, ttl time.Duration) error
	Get(key string) (any, error)
	Del(key string) error
}

func _get() (st Storage) {
	err := container.Get(&st)
	if err != nil {
		panic(err)
	}
	return
}

func Get(key string) (any, error) {
	return _get().Get(key)
}

func Save(key string, value any, ttl time.Duration) error {
	return _get().Save(key, value, ttl)
}

func Del(key string) error {
	return _get().Del(key)
}
