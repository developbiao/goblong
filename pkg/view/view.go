package view

import (
	"embed"
	"fmt"
	"goblong/app/models/category"
	"goblong/pkg/auth"
	"goblong/pkg/flash"
	"goblong/pkg/logger"
	"goblong/pkg/route"
	"html/template"
	"io"
	"io/fs"
	"strings"
)

// D represent map[string]interface
type D map[string]interface{}

var TplFS embed.FS

// Render common view
func Render(w io.Writer, data D, tplFiles ...string) {
	RenderTemplate(w, "app", data, tplFiles...)
}

// Render simple view
func RenderSimple(w io.Writer, data D, tplFiles ...string) {
	RenderTemplate(w, "simple", data, tplFiles...)
}

// Render template
func RenderTemplate(w io.Writer, name string, data D, tplFiles ...string) {
	// Get common template
	data["isLogined"] = auth.Check()
	data["loginUser"] = auth.User
	data["flash"] = flash.All()
	data["Categories"], _ = category.All()

	// Get template files
	allFiles := getTemplateFiles(tplFiles...)

	tmpl, err := template.New("").
		Funcs(template.FuncMap{
			"RouteName2URL": route.Name2URL,
		}).ParseFS(TplFS, allFiles...)

	logger.LogError(err)

	// Render
	tmplErr := tmpl.ExecuteTemplate(w, name, data)
	if tmplErr != nil {
		fmt.Println("Render Error! >>> ", tmplErr)
	}

}

func getTemplateFiles(tplFiles ...string) []string {
	// Set related template path
	viewDir := "resources/views/"

	for i, f := range tplFiles {
		// sugar syntax
		tplFiles[i] = viewDir + strings.Replace(f, ".", "/", -1) + ".gohtml"
	}

	// All layouts slice
	layoutFiles, err := fs.Glob(TplFS, viewDir+"layouts/*.gohtml")
	logger.LogError(err)

	// Append show article page
	return append(layoutFiles, tplFiles...)

}
