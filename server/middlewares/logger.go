package middlewares

import (
	"chat/shared/logging"
	"log/slog"
	"net/http"

	"github.com/google/uuid"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := slog.Default()
		logger = logger.With(slog.String("request_id", uuid.NewString()))

		logger.Info(
			"started handling request",
			slog.String("method", r.Method),
			slog.String("path", r.URL.Path),
		)

		ctx := logging.InjectLogger(r.Context(), logger)

		next.ServeHTTP(w, r.WithContext(ctx))

		logger.Info("finished handling request")
	})
}
