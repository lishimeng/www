package www

import "time"

type TokenStore interface {
	Save(data string, ttl time.Duration)
}
