package view

import (
	"fmt"
	"goblong/pkg/logger"
	"goblong/pkg/route"
	"html/template"
	"io"
	"path/filepath"
	"strings"
)

// Render template
func Render(w io.Writer, data interface{}, tplFiles ...string) {

	// Set related template path
	viewDir := "resources/views/"

	for i, f := range tplFiles {
		// sugar syntax
		tplFiles[i] = viewDir + strings.Replace(f, ".", "/", -1) + ".gohtml"
	}

	// All layouts
	layoutFiles, err := filepath.Glob(viewDir + "/layouts/*.gohtml")
	logger.LogError(err)

	// Append show article page
	allFiles := append(layoutFiles, tplFiles...)

	tmpl, err := template.New("").
		Funcs(template.FuncMap{
			"RouteName2URL": route.Name2URL,
		}).ParseFiles(allFiles...)

	logger.LogError(err)

	// Render
	tmplErr := tmpl.ExecuteTemplate(w, "app", data)
	if tmplErr != nil {
		fmt.Println("Render Error! >>> ", tmplErr)
	}

}
