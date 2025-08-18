package utils

import (
	"context"
	"time"
)

const (
	DefaultTimeout = 30 * time.Second
)

func DefaultContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), DefaultTimeout)
}
