package defaulttheme

import (
	"embed"
	"html/template"
	texttemplate "text/template"

	"github.com/gzuidhof/myrtle/theme"
	"github.com/gzuidhof/myrtle/theme/themerender"
)

//go:embed *.tmpl
var templatesFS embed.FS

type Theme struct {
	htmlTemplates *template.Template
	textTemplates *texttemplate.Template
}

func New() *Theme {
	htmlTemplateFiles := append(
		[]string{"layout.html.tmpl"},
		themerender.SharedBlockTemplateFiles()...,
	)

	htmlTemplates := themerender.ParseHTMLTemplates(
		"default-html",
		templatesFS,
		htmlTemplateFiles...,
	)

	textTemplates := texttemplate.Must(texttemplate.New("default-text").ParseFS(
		templatesFS,
		"layout.text.tmpl",
	))

	return &Theme{
		htmlTemplates: htmlTemplates,
		textTemplates: textTemplates,
	}
}

func (themeImpl *Theme) Name() string {
	return "default"
}

func (themeImpl *Theme) RenderHTML(view theme.EmailView) (string, error) {
	return themerender.ExecuteTemplate(themeImpl.htmlTemplates, "layout.html.tmpl", view)
}

func (themeImpl *Theme) RenderBlockHTML(view theme.BlockView) (string, bool, error) {
	return themerender.RenderBlockHTML(themeImpl.htmlTemplates, view, nil)
}

func (themeImpl *Theme) WrapMarkdown(view theme.TextView) (string, error) {
	return themerender.ExecuteTextTemplate(themeImpl.textTemplates, "layout.text.tmpl", view)
}
