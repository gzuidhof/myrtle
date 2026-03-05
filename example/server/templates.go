package server

import (
	"html/template"
	"net/http"
)

type templateSet struct {
	index *template.Template
	email *template.Template
	block *template.Template
}

func newTemplateSet() (templateSet, error) {
	indexTemplate, err := parsePageTemplate("templates/index.html.tmpl")
	if err != nil {
		return templateSet{}, err
	}

	emailTemplate, err := parsePageTemplate("templates/email.html.tmpl")
	if err != nil {
		return templateSet{}, err
	}

	blockTemplate, err := parsePageTemplate("templates/block.html.tmpl")
	if err != nil {
		return templateSet{}, err
	}

	return templateSet{
		index: indexTemplate,
		email: emailTemplate,
		block: blockTemplate,
	}, nil
}

func parsePageTemplate(pageFile string) (*template.Template, error) {
	return template.ParseFS(
		templatesFS,
		"templates/layout.html.tmpl",
		pageFile,
	)
}

func setHTMLHeader(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "text/html; charset=utf-8")
}
