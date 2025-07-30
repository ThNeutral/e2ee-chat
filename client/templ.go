package client

import (
	"chat/shared"
	"embed"
	"fmt"
	"html/template"
)

type templateName string

func (tn templateName) String() string {
	return string(tn)
}

const (
	pageTemplateName       templateName = "page"
	componentsTemplateName templateName = "component"
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

func getTemplateName(part templateName, name string) string {
	return fmt.Sprintf("%v_%v.html", part, name)
}
