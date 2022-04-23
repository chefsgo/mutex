package mutex

import (
	"time"
)

type (
	// Driver 数据驱动
	Driver interface {
		Connect(name string, config Config) (Connect, error)
	}

	// Connect 会话连接
	Connect interface {
		//打开、关闭
		Open() error
		Close() error

		Lock(key string, expiry time.Duration) error
		Unlock(key string) error
	}
)

// Driver 注册驱动
func (module *Module) Driver(name string, driver Driver, override bool) {
	module.mutex.Lock()
	defer module.mutex.Unlock()

	if driver == nil {
		panic("Invalid mutex driver: " + name)
	}

	if override {
		module.drivers[name] = driver
	} else {
		if module.drivers[name] == nil {
			module.drivers[name] = driver
		}
	}
}
