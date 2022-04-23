package mutex

import (
	"time"
)

func Lock(key string, expiries ...time.Duration) error {
	return module.Lock(key, expiries...)
}
func Unlock(key string) error {
	return module.Unlock(key)
}
func LockTo(conn string, key string, expiries ...time.Duration) error {
	return module.Lock(key, expiries...)
}
func UnlockFrom(conn string, key string) error {
	return module.UnlockFrom(conn, key)
}

func Locked(key string, expiries ...time.Duration) bool {
	return module.Lock(key, expiries...) != nil
}
