package documentation

import (
	"html/template"
)


//go:embed "templates"
var templateFS embed.FS

func Documentation() (*templates.Template, error) {
	tmpl, err := template.New("email").ParseFS(templateFS, "templates/"+templateFile)
	if err != nil {
		return err
	}
}