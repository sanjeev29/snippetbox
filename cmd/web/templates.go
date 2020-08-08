package main

import (
	"html/template"
	"path/filepath"
	"time"

	"snippetbox/pkg/forms"
	"snippetbox/pkg/models"
)

// templateData struct for dynamic data for templates
type templateData struct {
	CurrentYear       int
	Flash             string
	Form              *forms.Form
	Snippet           *models.Snippet
	Snippets          []*models.Snippet
	Users             *models.User
	AuthenticatedUser int
	CSRFToken         string
}

func humanDate(t time.Time) string {
	return t.Format("02 Jan 2006 at 15:04")
}

var functions = template.FuncMap{
	"humanDate": humanDate,
}

// Template caching
func newTemplateCache(dir string) (map[string]*template.Template, error) {
	// init new map to act as cache
	cache := map[string]*template.Template{}

	// get slice of all filepaths with '.page.tmpl'
	pages, err := filepath.Glob(filepath.Join(dir, "*.page.tmpl"))
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		// extract filename
		name := filepath.Base(page)

		// parse template file in to a template set
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return nil, err
		}

		// add  '.layout.tmpl' files to template set
		ts, err = ts.ParseGlob(filepath.Join(dir, "*.layout.tmpl"))
		if err != nil {
			return nil, err
		}

		// add '.partial.tmpl' files to template set
		ts, err = ts.ParseGlob(filepath.Join(dir, "*.partial.tmpl"))
		if err != nil {
			return nil, err
		}

		// add template set to cache using file name of page as the key
		cache[name] = ts

	}

	return cache, nil
}
