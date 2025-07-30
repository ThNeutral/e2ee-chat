package client

import (
	"chat/shared"
	"fmt"
	"html/template"
	"net"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Client struct {
	handler http.Handler

	serverAddr net.Addr
	clientPort int

	templates *template.Template

	httpClient *http.Client
}

func New(serverAddr net.Addr, clientPort int) (*Client, error) {
	eb := shared.NewErrorBuilder().Msg("failed to initialize client")

	templ, err := initTemplates()
	if err != nil {
		return nil, eb.Cause(err).Err()
	}

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	cl := &Client{
		handler: r,

		serverAddr: serverAddr,
		clientPort: clientPort,

		templates: templ,

		httpClient: &http.Client{},
	}

	r.Post("/echo", cl.handleEcho)
	r.Get("/static/{name}", cl.handleStatic)

	return cl, nil
}

func (c *Client) Run() error {
	return http.ListenAndServe(fmt.Sprintf(":%v", c.clientPort), c.handler)
}

func (c *Client) getServerURL(endpoint string) string {
	return fmt.Sprintf("http://%v/%v", c.serverAddr.String(), endpoint)
}
