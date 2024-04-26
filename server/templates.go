package main

import (
	"errors"
	"html/template"
	"io"

	"github.com/labstack/echo"
)

// Define the template registry struct
type TemplateRegistry struct {
	templates map[string]*template.Template
}

// Implement e.Renderer interface
func (t *TemplateRegistry) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	tmpl, ok := t.templates[name]
	if !ok {
		err := errors.New("Template not found -> " + name)
		return err
	}
	return tmpl.ExecuteTemplate(w, "base.html", data)
}

func createTemplate(templ string, base string) *template.Template {

	return template.Must(template.ParseFiles("../templates/"+templ+".html", "../templates/"+base+".html"))
}

func ListTemplates() map[string]*template.Template {
	templates := make(map[string]*template.Template)

	templates["home"] = createTemplate("home", "base")
	templates["profile"] = createTemplate("profile/profile", "base")
	templates["editprofile"] = createTemplate("profile/editprofile", "base")
	templates["get"] = createTemplate("get", "base")
	templates["post"] = createTemplate("post", "base")

	return templates
}
