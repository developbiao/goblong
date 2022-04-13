package bootstrap

import (
	"embed"
	"goblong/pkg/view"
)

// SetupTemplate initializes the templates
func SetupTemplate(tmplFS embed.FS) {
	view.TplFS = tmplFS
}
