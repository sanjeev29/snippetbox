package main

import "snippetbox/pkg/models"

// templateData struct for dynamic data for templates
type templateData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}
