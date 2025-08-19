package middlewares

import (
	"bufio"
	"chat/shared/ctxinjector"
	"fmt"
	"log/slog"
	"net"
	"net/http"
)

type responseWriterWrapper struct {
	impl   http.ResponseWriter
	logger *slog.Logger
}

func (w *responseWriterWrapper) Header() http.Header {
	return w.impl.Header()
}

func (w *responseWriterWrapper) Write(bytes []byte) (int, error) {
	w.logger.Info("writing response", "payload", string(bytes))
	return w.impl.Write(bytes)
}

func (w *responseWriterWrapper) WriteHeader(status int) {
	w.logger.Info("writing response header", "code", status)
	w.impl.WriteHeader(status)
}

func (w *responseWriterWrapper) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	if h, ok := w.impl.(http.Hijacker); ok {
		return h.Hijack()
	}
	return nil, nil, fmt.Errorf("underlying ResponseWriter does not support hijacking")
}

func (w *responseWriterWrapper) Flush() {
	if f, ok := w.impl.(http.Flusher); ok {
		f.Flush()
	}
}

func ResponseWriterWrapper(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		wrapper := &responseWriterWrapper{
			impl:   w,
			logger: ctxinjector.GetLogger(r.Context()),
		}

		next.ServeHTTP(wrapper, r)
	})
}
