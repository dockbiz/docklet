package src

import (
	"golang.org/x/net/context"
)

func IsTimeoutError(err error) bool {
	if err == context.DeadlineExceeded {
		return true
	}
	if t, ok := err.(interface {
		IsTimeout() bool
	}); ok {
		return t.IsTimeout()
	}
	return false
}
