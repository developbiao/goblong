package view

import (
	"goblong/pkg/logger"
	"goblong/pkg/route"
	"html/template"
	"io"
	"path/filepath"
	"strings"
)

// Render template
func Render(w io.Writer, name string, data interface{}) {

	// Set related template path
	viewDir := "resources/views/"

	// sugar syntax
	name = strings.Replace(name, ".", "/", -1)

	// All layouts
	files, err := filepath.Glob(viewDir + "/layouts/*.gohtml")
	logger.LogError(err)

	// Append show article page
	newFiles := append(files, viewDir+"/articles/show.gohtml")

	tmpl, err := template.New(name).
		Funcs(template.FuncMap{
			"RouteName2URL": route.Name2URL,
		}).ParseFiles(newFiles...)

	logger.LogError(err)

	// Render
	tmpl.ExecuteTemplate(w, "app", data)

}
