package server

import (
	"io"
	"text/template"

	"github.com/labstack/echo/v4"
)

const viewsPath = "views/*.html"

type Templates struct {
	templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

//TODO: files cannot be served like this as we need to put them in the binary

func NewTemplates() *Templates {
	return &Templates{
		templates: template.Must(template.ParseGlob(viewsPath)),
	}
}
