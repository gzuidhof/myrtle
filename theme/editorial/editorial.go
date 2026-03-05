package editorial

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

// Theme is an editorial visual style theme with optional fallback block rendering.
type Theme struct {
	htmlTemplates *template.Template
	textTemplates *texttemplate.Template
	handlers      map[theme.BlockKind]themerender.BlockRenderHandler
	fallback      theme.Theme
	styles        theme.Styles
}

// Option configures an editorial theme during construction.
type Option func(*Theme)

// WithFallback sets a fallback theme for blocks this theme does not render directly.
func WithFallback(fallback theme.Theme) Option {
	return func(themeImpl *Theme) {
		themeImpl.fallback = fallback
	}
}

// New constructs an editorial theme instance and applies optional configuration.
func New(options ...Option) *Theme {
	sharedTemplateFiles := themerender.SharedBlockTemplateFilesAvailableInFS(templatesFS)

	htmlTemplateFiles := append([]string{"layout.html.tmpl"}, sharedTemplateFiles...)

	htmlTemplates := themerender.ParseHTMLTemplatesWithShared(
		"editorial-html",
		templatesFS,
		htmlTemplateFiles...,
	)

	textTemplates := texttemplate.Must(texttemplate.New("editorial-text").ParseFS(
		templatesFS,
		"layout.text.tmpl",
	))

	themeImpl := &Theme{
		htmlTemplates: htmlTemplates,
		textTemplates: textTemplates,
		handlers:      themerender.DefaultBlockRenderHandlersForTemplateFiles(sharedTemplateFiles),
		fallback:      defaulttheme.New(),
		styles: theme.Styles{
			ColorPrimary:              "#ad4f2d",
			ColorSecondary:            "#8f6a3b",
			ColorText:                 "#2f241d",
			ColorTextMuted:            "#6f5f53",
			ColorBorder:               "#ddcbb8",
			ColorCodeBackground:       "#f8efe6",
			ColorPageBackground:       "#f7f0e8",
			ColorMainBackground:       "#fffaf5",
			ColorSurface:              "#fffdf9",
			ColorSurfaceMuted:         "#f5ebe0",
			ColorTextOnSolid:          "#fffaf3",
			ColorInfo:                 "#2f6fb6",
			ColorInfoBorder:           "#9dc0e5",
			ColorInfoBackground:       "#edf4fb",
			ColorInfoText:             "#245488",
			ColorSuccess:              "#3f7d4e",
			ColorSuccessBorder:        "#a6d1ad",
			ColorSuccessBackground:    "#edf8ef",
			ColorSuccessText:          "#2f5f3b",
			ColorWarning:              "#b7791f",
			ColorWarningBorder:        "#efc48b",
			ColorWarningBackground:    "#fff7ea",
			ColorWarningText:          "#8f5f17",
			ColorDanger:               "#b2453f",
			ColorDangerBorder:         "#e4aaa7",
			ColorDangerBackground:     "#fdf0ef",
			ColorDangerText:           "#8e322d",
			BorderMain:                "1px solid #ddcbb8",
			WidthMain:                 "100%",
			MaxWidthMain:              "660px",
			OuterPadding:              "24px",
			OutsideContentInset:       "28px",
			MainContentBodyTopSpacing: "26px",
			RadiusMain:                "14px",
			RadiusElement:             "12px",
			RadiusButton:              "999px",
			RadiusPill:                "999px",
			TableLegendSwatchSize:     "11px",
			TableLegendSwatchRadius:   "4px",
			TableLegendSwatchBorder:   "1px solid #ddcbb8",
			FontFamilyBase:            "\"Iowan Old Style\",\"Palatino Linotype\",\"Book Antiqua\",Georgia,\"Times New Roman\",Times,serif",
			FontFamilyMono:            "ui-monospace,SFMono-Regular,Menlo,Monaco,Consolas,monospace",
			FontSizeBase:              "16px",
			LineHeightBase:            "1.65",
			FontWeightHeading:         "700",
		},
	}

	for _, option := range options {
		option(themeImpl)
	}

	return themeImpl
}

func (themeImpl *Theme) Name() string {
	return "editorial"
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
