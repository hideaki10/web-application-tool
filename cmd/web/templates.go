package main

import (
	"path/filepath"
	"text/template"
	"time"

	"github.com/hideaki10/web-application-tool/pkg/forms"
	"github.com/hideaki10/web-application-tool/pkg/models"
)

type templateData struct {
	Snippet         *models.Snippet
	Flash           string
	Snippets        []*models.Snippet
	CurrentYear     int
	Form            *forms.Form
	IsAuthenticated bool
	CSRFToken       string
}

func newTemplateCache(dir string) (map[string]*template.Template, error) {

	cache := map[string]*template.Template{}

	pages, err := filepath.Glob(filepath.Join(dir, "*.page.tmpl"))
	if err != nil {
		return nil, err
	}

	for _, page := range pages {

		name := filepath.Base(page)

		// ts, err := template.ParseFiles(page)
		// if err != nil {
		// 	return nil, err
		// }
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return nil, err
		}
		ts, err = ts.ParseGlob(filepath.Join(dir, "*.layout.tmpl"))
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob(filepath.Join(dir, "*.partial.tmpl"))
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	// Return the map.
	return cache, nil
}
func humanDate(t time.Time) string {

	if t.IsZero() {
		return ""
	}
	return t.UTC().Format("02 Jan 2006 at 15:04")
}

var functions = template.FuncMap{
	"humanDate": humanDate,
}
