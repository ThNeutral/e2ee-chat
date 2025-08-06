package shared

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Runner struct {
	handler http.Handler
}

func NewRunner() *Runner {
	r := chi.NewMux()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	return &Runner{
		handler: r,
	}
}

func (r *Runner) Post(path string, handler http.HandlerFunc) {
	r.chi().Post(path, handler)
}

func (r *Runner) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.handler.ServeHTTP(w, req)
}

func (r *Runner) chi() *chi.Mux {
	return r.handler.(*chi.Mux)
}
