package mutex

import (
	"sync"
	"time"

	. "github.com/chefsgo/base"
	"github.com/chefsgo/chef"
	"github.com/chefsgo/util"
)

func init() {
	chef.Register(NAME, module)
}

var (
	module = &Module{
		configs:   make(map[string]Config, 0),
		drivers:   make(map[string]Driver, 0),
		instances: make(map[string]Instance, 0),
	}
)

type (
	Module struct {
		mutex sync.Mutex

		connected, initialized, launched bool

		configs map[string]Config

		drivers   map[string]Driver
		instances map[string]Instance

		weights  map[string]int
		hashring *util.HashRing
	}

	Config struct {
		Driver  string
		Weight  int
		Prefix  string
		Expiry  time.Duration
		Setting Map
	}
	Instance struct {
		name    string
		config  Config
		connect Connect
	}
)

// Lock 加锁
func (this Module) Lock(key string, expiries ...time.Duration) error {
	locate := module.hashring.Locate(key)

	if inst, ok := module.instances[locate]; ok {

		expiry := inst.config.Expiry
		if len(expiries) > 0 {
			expiry = expiries[0]
		}

		// 加上前缀
		key := inst.config.Prefix + key

		return inst.connect.Lock(key, expiry)
	}

	return errInvalidMutexConnection
}

// LockTo 加锁到指定的连接
func (this Module) LockTo(conn string, key string, expiries ...time.Duration) error {
	if inst, ok := module.instances[conn]; ok {

		//默认过期时间
		expiry := inst.config.Expiry
		if len(expiries) > 0 {
			expiry = expiries[0]
		}

		// 加上前缀
		key := inst.config.Prefix + key

		return inst.connect.Lock(key, expiry)
	}

	return errInvalidMutexConnection
}

// Unlock 解锁
func (this Module) Unlock(key string) error {
	locate := module.hashring.Locate(key)

	if inst, ok := module.instances[locate]; ok {
		key := inst.config.Prefix + key //加上前缀
		return inst.connect.Unlock(key)
	}

	return errInvalidMutexConnection
}

// UnlockFrom 从指定的连接解锁
func (this Module) UnlockFrom(locate string, key string) error {
	if inst, ok := module.instances[locate]; ok {
		key := inst.config.Prefix + key //加上前缀
		return inst.connect.Unlock(key)
	}

	return errInvalidMutexConnection
}
