package http

import (
	"chat/shared/errs"
	log "chat/shared/slog"
	"chat/shared/utils"
	"context"
	"net/http"
)

type Handler[T Validatable] func(ctx context.Context, req T) (any, error)

func WrapHandler[T Validatable](handler Handler[T]) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := utils.DefaultContext()
		defer cancel()

		ctx = log.AddRequestLogger(ctx, r.URL.Path)
		logger := log.GetLogger(ctx)

		logger.Info("started handling request")
		defer func() {
			logger.Info("finished handling request")
		}()

		req, err := ParseHTTPRequest[T](r)
		if err != nil {
			logger.Error("error", "err", err.Error(), "code", http.StatusBadRequest)
			WriteError(w, http.StatusBadRequest, err.Error())
			return
		}

		logger.Info("request", "req", req)

		resp, err := handler(ctx, req)
		if err != nil {
			e, ok := err.(*errs.Error)
			if !ok {
				logger.Error("error", "err", err.Error(), "code", http.StatusBadRequest)
				WriteError(w, http.StatusInternalServerError, err.Error())
				return
			}

			logger.Error("error", "err", e.Error(), "code", e.Code)
			WriteError(w, e.Code, e.Error())
			return
		}

		logger.Info("request", "resp", resp)

		WriteSuccess(w, resp)
	}
}
