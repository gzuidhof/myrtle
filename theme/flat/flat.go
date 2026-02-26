package flat

import (
	"embed"
	"html/template"
	texttemplate "text/template"

	"github.com/gzuidhof/myrtle/theme"
	defaulttheme "github.com/gzuidhof/myrtle/theme/default"
	"github.com/gzuidhof/myrtle/theme/themerender"
)

//go:embed *.tmpl
var templatesFS embed.FS

type Theme struct {
	htmlTemplates *template.Template
	textTemplates *texttemplate.Template
	handlers      map[theme.BlockKind]themerender.BlockRenderHandler
	fallback      theme.Theme
	styles        theme.Styles
}

type Option func(*Theme)

func WithFallback(fallback theme.Theme) Option {
	return func(themeImpl *Theme) {
		themeImpl.fallback = fallback
	}
}

func New(options ...Option) *Theme {
	sharedTemplateFiles := themerender.SharedBlockTemplateFilesAvailableInFS(templatesFS)

	htmlTemplateFiles := append([]string{"layout.html.tmpl"}, sharedTemplateFiles...)

	htmlTemplates := themerender.ParseHTMLTemplates(
		"flat-html",
		templatesFS,
		htmlTemplateFiles...,
	)

	textTemplates := texttemplate.Must(texttemplate.New("flat-text").ParseFS(
		templatesFS,
		"layout.text.tmpl",
	))

	themeImpl := &Theme{
		htmlTemplates: htmlTemplates,
		textTemplates: textTemplates,
		handlers:      themerender.DefaultBlockRenderHandlersForTemplateFiles(sharedTemplateFiles),
		fallback:      defaulttheme.New(),
		styles: theme.Styles{
			ColorPrimary:           "#265cff",
			ColorSecondary:         "#10b981",
			ColorText:              "#111827",
			ColorTextMuted:         "#6b7280",
			ColorBorder:            "#d1d5db",
			ColorCodeBackground:    "#f8fafc",
			ColorPageBackground:    "#ffffff",
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
			BorderMain:             "none",
			WidthMain:              "100%",
			MaxWidthMain:           "640px",
			OuterPadding:           "32px",
			OutsideContentInset:    "20px",
			RadiusMain:             "0px",
			RadiusElement:          "0px",
			RadiusButton:           "0px",
			RadiusPill:             "0px",
			FontFamilyBase:         "system-ui,-apple-system,BlinkMacSystemFont,'Segoe UI',Roboto,'Helvetica Neue',Arial,sans-serif",
			FontFamilyMono:         "ui-monospace,SFMono-Regular,Menlo,Monaco,Consolas,monospace",
			FontSizeBase:           "14px",
			LineHeightBase:         "1.5",
			FontWeightHeading:      "700",
		},
	}

	for _, option := range options {
		option(themeImpl)
	}

	return themeImpl
}

func (themeImpl *Theme) Name() string {
	return "flat"
}

func (themeImpl *Theme) DefaultStyles() theme.Styles {
	return themeImpl.styles
}

func (themeImpl *Theme) RenderHTML(view theme.EmailView) (string, error) {
	return themerender.ExecuteTemplate(themeImpl.htmlTemplates, "layout.html.tmpl", view)
}

func (themeImpl *Theme) RenderBlockHTML(view theme.BlockView) (string, bool, error) {
	return themerender.RenderBlockHTMLWithHandlers(themeImpl.htmlTemplates, view, themeImpl.handlers, themeImpl.fallback)
}

func (themeImpl *Theme) WrapText(view theme.TextView) (string, error) {
	if wrapper, ok := themeImpl.fallback.(theme.TextWrapper); ok {
		return wrapper.WrapText(view)
	}

	return themerender.ExecuteTextTemplate(themeImpl.textTemplates, "layout.text.tmpl", view)
}
