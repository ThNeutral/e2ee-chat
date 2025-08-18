package runner

import (
	"chat/shared/http"
	stdhttp "net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Runner struct {
	handler stdhttp.Handler
}

func New() *Runner {
	r := chi.NewMux()

	r.Use(middleware.Recoverer)

	return &Runner{
		handler: r,
	}
}

func Post[T http.Validatable](r *Runner, path string, handler http.Handler[T]) {
	r.chi().Post(path, http.WrapHandler(handler))
}

func (r *Runner) ServeHTTP(w stdhttp.ResponseWriter, req *stdhttp.Request) {
	r.handler.ServeHTTP(w, req)
}

func (r *Runner) chi() *chi.Mux {
	return r.handler.(*chi.Mux)
}
