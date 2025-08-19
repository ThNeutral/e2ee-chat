package server

import (
	"net/http"
)

type Hub interface {
	Accept(http.ResponseWriter, *http.Request) error
}
