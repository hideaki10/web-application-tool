package main

import "github.com/hideaki10/web-application-tool/pkg/models"

type templateData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}
