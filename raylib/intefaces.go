package raylib

import "context"

type Echo interface {
	Echo(ctx context.Context, payload string) (string, error)
}
