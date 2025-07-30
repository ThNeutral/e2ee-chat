package client

import (
	"chat/shared"
	"embed"
	"html/template"
)

//go:embed html/*.html
var htmlFiles embed.FS

func initTemplates() (*template.Template, error) {
	eb := shared.NewErrorBuilder().Msg("failed to initialize html templates")

	templs, err := template.ParseFS(htmlFiles, "html/*.html")
	if err != nil {
		return nil, eb.Cause(err).Err()
	}

	return templs, nil
}
