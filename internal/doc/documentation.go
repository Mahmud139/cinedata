package doc

import (
	"embed"
	"fmt"
	"html/template"
)

//go:embed "templates"
var templateFS embed.FS

func Documentation() (*template.Template, error) {
	tmpl, err := template.New("doc").ParseFS(templateFS, "templates/doc.tmpl")
	if err != nil {
		fmt.Println("here")
		return nil, err
	}

	return tmpl, err
}
