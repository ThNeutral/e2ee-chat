package log

import (
	"context"
	"log/slog"

	"github.com/google/uuid"
)

type loggerKey struct{}

func AddRequestLogger(ctx context.Context, endpoint string, additionalKeyValues ...any) context.Context {
	if ctx == nil {
		panic("ctx is nil")
	}

	if GetLogger(ctx) != nil {
		panic("ctx already has logger")
	}

	data := []any{
		slog.String("endpoint", endpoint),
		slog.String("request_id", uuid.NewString()),
	}

	if len(additionalKeyValues) != 0 {
		data = append(data, additionalKeyValues...)
	}

	logger := slog.Default().With(data...)
	return addLogger(ctx, logger)
}

func addLogger(ctx context.Context, logger *slog.Logger) context.Context {
	return context.WithValue(ctx, loggerKey{}, logger)
}

func GetLogger(ctx context.Context) *slog.Logger {
	logger, ok := ctx.Value(loggerKey{}).(*slog.Logger)
	if !ok {
		return nil
	}
	return logger
}
