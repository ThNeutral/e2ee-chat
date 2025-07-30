package client

import (
	"chat/shared"
	"net/http"
	"strings"
)

func (c *Client) handleStatic(w http.ResponseWriter, r *http.Request) {
	eb := shared.NewErrorBuilder().Msg("failed to handle static")

	parts := strings.Split(r.URL.Path, "/")
	name := parts[len(parts)-1]

	err := c.templates.ExecuteTemplate(w, getTemplateName(pageTemplateName, name), nil)
	if err != nil {
		shared.WriteHTTPError(w, http.StatusNotFound, eb.Cause(err).Err())
		return
	}
}
