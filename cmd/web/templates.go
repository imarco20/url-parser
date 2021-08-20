package main

import (
	"marcode.io/url-parser/pkg/forms"
	"marcode.io/url-parser/pkg/models"
	"path/filepath"
	"text/template"
)

type templateData struct {
	Form *forms.Form
	Link models.LinkDetails
}

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
