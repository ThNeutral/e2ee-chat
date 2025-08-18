package repository

import (
	"chat/shared/errs"
	"fmt"
	"net"
	"net/http"
)

type RepositoryConfig struct {
	ServerAddr net.Addr
	HTTPClient *http.Client
}
type Repository struct {
	serverAddr net.Addr
	httpClient *http.Client
}

func New(cfg RepositoryConfig) (*Repository, error) {
	eb := errs.B().Msg("failed to initialize repository")

	if cfg.ServerAddr == nil {
		return nil, eb.Causef("server addr not passed").Err()
	}

	if cfg.HTTPClient == nil {
		return nil, eb.Causef("http cause not passed").Err()
	}

	return &Repository{
		serverAddr: cfg.ServerAddr,
		httpClient: cfg.HTTPClient,
	}, nil
}

func (r *Repository) getServerURL(endpoint string) string {
	return fmt.Sprintf("http://%v%v", r.serverAddr.String(), endpoint)
}
