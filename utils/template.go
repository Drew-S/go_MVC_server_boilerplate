package utils

import (
	"html/template"
	"path"
)

// CreateTemplate returns an html/template to execute against with the shared
//   html templates in views/shared
func CreateTemplate(name string) (*template.Template, error) {
	return template.New(path.Base(name)).
		Funcs(template.FuncMap{
			"title": func() string {
				if GetConfig().Name == "" {
					return "Site name"
				}
				return GetConfig().Name
			},
			"dev": func() bool {
				return GetConfig().Dev
			},
		}).
		ParseFiles([]string{
			name,
			"views/shared/layout.html",
		}...)
}
