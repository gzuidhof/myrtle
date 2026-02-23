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
	styles        theme.Styles
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
		styles: theme.Styles{
			ColorPrimary:        "#265cff",
			ColorSecondary:      "#10b981",
			ColorText:           "#111827",
			ColorTextMuted:      "#6b7280",
			ColorBorder:         "#e5e7eb",
			ColorCodeBackground: "#f8fafc",
			ColorPageBackground: "#f3f4f6",
			ColorMainBackground: "#ffffff",
			BorderMain:          "1px solid #e5e7eb",
			RadiusMain:          "12px",
			FontFamilyBase:      "system-ui,-apple-system,BlinkMacSystemFont,'Segoe UI',Roboto,'Helvetica Neue',Arial,sans-serif",
			FontFamilyMono:      "ui-monospace,SFMono-Regular,Menlo,Monaco,Consolas,monospace",
			FontSizeBase:        "14px",
			LineHeightBase:      "1.6",
			FontWeightHeading:   "700",
		},
	}
}

func (themeImpl *Theme) Name() string {
	return "default"
}

func (themeImpl *Theme) DefaultStyles() theme.Styles {
	return themeImpl.styles
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
