package mutex

import "errors"

const (
	NAME    = "mutex"
	DEFAULT = "default"
)

var (
	errInvalidMutexConnection = errors.New("Invalid mutex connection.")
)
