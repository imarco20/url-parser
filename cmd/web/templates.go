package main

import (
	"marcode.io/url-parser/pkg/forms"
	"marcode.io/url-parser/pkg/models"
	"path/filepath"
	"text/template"
)

// templateData holds any data we want to pass to the template file
type templateData struct {
	Form *forms.Form
	Link models.LinkDetails
}

// cacheAllTemplates scans the UI directory for template files,
// parses them into template sets and saves them to the template cache map
func cacheAllTemplates(dir string) (map[string]*template.Template, error) {

	templatesCache := map[string]*template.Template{}

	pages, err := filepath.Glob(filepath.Join(dir, "*.page.tmpl"))
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		pageName := filepath.Base(page)

		templateSet, err := template.New(pageName).ParseFiles(page)
		if err != nil {
			return nil, err
		}

		templateSet, err = templateSet.ParseGlob(filepath.Join(dir, "*.layout.tmpl"))
		if err != nil {
			return nil, err
		}

		templateSet, err = templateSet.ParseGlob(filepath.Join(dir, "*.partial.tmpl"))
		if err != nil {
			return nil, err
		}

		templatesCache[pageName] = templateSet
	}

	return templatesCache, nil
}
