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
			ColorPrimary:           "#265cff",
			ColorSecondary:         "#10b981",
			ColorText:              "#111827",
			ColorTextMuted:         "#6b7280",
			ColorBorder:            "#e5e7eb",
			ColorCodeBackground:    "#f8fafc",
			ColorPageBackground:    "#f3f4f6",
			ColorMainBackground:    "#ffffff",
			ColorSurface:           "#ffffff",
			ColorSurfaceMuted:      "#f8fafc",
			ColorTextOnSolid:       "#ffffff",
			ColorInfo:              "#2563eb",
			ColorInfoBorder:        "#93c5fd",
			ColorInfoBackground:    "#eff6ff",
			ColorInfoText:          "#1d4ed8",
			ColorSuccess:           "#16a34a",
			ColorSuccessBorder:     "#86efac",
			ColorSuccessBackground: "#f0fdf4",
			ColorSuccessText:       "#15803d",
			ColorWarning:           "#ca8a04",
			ColorWarningBorder:     "#fcd34d",
			ColorWarningBackground: "#fffbeb",
			ColorWarningText:       "#92400e",
			ColorDanger:            "#dc2626",
			ColorDangerBorder:      "#fca5a5",
			ColorDangerBackground:  "#fef2f2",
			ColorDangerText:        "#b91c1c",
			BorderMain:             "1px solid #e5e7eb",
			WidthMain:              "100%",
			MaxWidthMain:           "640px",
			OuterPadding:           "24px",
			OutsideContentInset:    "24px",
			RadiusMain:             "12px",
			RadiusElement:          "10px",
			RadiusButton:           "8px",
			RadiusPill:             "999px",
			FontFamilyBase:         "system-ui,-apple-system,BlinkMacSystemFont,'Segoe UI',Roboto,'Helvetica Neue',Arial,sans-serif",
			FontFamilyMono:         "ui-monospace,SFMono-Regular,Menlo,Monaco,Consolas,monospace",
			FontSizeBase:           "14px",
			LineHeightBase:         "1.6",
			FontWeightHeading:      "700",
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

func (themeImpl *Theme) WrapText(view theme.TextView) (string, error) {
	return themerender.ExecuteTextTemplate(themeImpl.textTemplates, "layout.text.tmpl", view)
}
