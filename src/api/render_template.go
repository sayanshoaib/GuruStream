package api

import (
	"github.com/labstack/echo/v4"
	"html/template"
	"io"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func TempRender(e *echo.Echo) {
	e.Static("/", "/Users/shoaib/GolandProjects/src/github.com/sayanshoaib/softwareLab/src/web")
	t := &Template{
		templates: template.Must(template.ParseGlob(
			"/Users/shoaib/GolandProjects/src/github.com/sayanshoaib/softwareLab/src/web/*.tmpl")),
	}
	e.Renderer = t
}
