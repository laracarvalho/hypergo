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

func createTemplate(templs ...string) *template.Template {
	if len(templs) == 1 {
		return template.Must(template.ParseGlob("../templates/" + templs[0] + ".html"))
	}

	return template.Must(template.ParseFiles("../templates/"+templs[0]+".html", "../templates/"+templs[1]+".html"))

	// var str string

	// for _, tmpl := range templs {
	// 	str += "../templates/" + tmpl + ".html, "
	// }

	// return template.Must(template.ParseFiles(...strings.Split(s, ",")))
}

func ListTemplates() map[string]*template.Template {
	templates := make(map[string]*template.Template)

	templates["home"] = createTemplate("home", "base")
	templates["battle"] = createTemplate("battle", "base")
	templates["profile"] = createTemplate("profile/profile", "base")
	templates["editprofile"] = createTemplate("profile/editprofile", "base")

	return templates
}
