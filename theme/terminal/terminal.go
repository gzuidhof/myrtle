package terminal

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
		"terminal-html",
		templatesFS,
		htmlTemplateFiles...,
	)

	textTemplates := texttemplate.Must(texttemplate.New("terminal-text").ParseFS(
		templatesFS,
		"layout.text.tmpl",
	))

	themeImpl := &Theme{
		htmlTemplates: htmlTemplates,
		textTemplates: textTemplates,
		handlers:      themerender.DefaultBlockRenderHandlersForTemplateFiles(sharedTemplateFiles),
		fallback:      defaulttheme.New(),
		styles: theme.Styles{
			ColorPrimary:           "#22c55e",
			ColorSecondary:         "#06b6d4",
			ColorText:              "#e2e8f0",
			ColorTextMuted:         "#94a3b8",
			ColorBorder:            "#334155",
			ColorCodeBackground:    "#020617",
			ColorPageBackground:    "#020617",
			ColorMainBackground:    "#0b1220",
			ColorSurface:           "#0f172a",
			ColorSurfaceMuted:      "#111827",
			ColorTextOnSolid:       "#f8fafc",
			ColorInfo:              "#3b82f6",
			ColorInfoBorder:        "#1d4ed8",
			ColorInfoBackground:    "#0b2a4a",
			ColorInfoText:          "#bfdbfe",
			ColorSuccess:           "#22c55e",
			ColorSuccessBorder:     "#15803d",
			ColorSuccessBackground: "#052e16",
			ColorSuccessText:       "#86efac",
			ColorWarning:           "#f59e0b",
			ColorWarningBorder:     "#b45309",
			ColorWarningBackground: "#451a03",
			ColorWarningText:       "#fcd34d",
			ColorDanger:            "#ef4444",
			ColorDangerBorder:      "#b91c1c",
			ColorDangerBackground:  "#450a0a",
			ColorDangerText:        "#fca5a5",
			BorderMain:             "1px solid #334155",
			WidthMain:              "100%",
			MaxWidthMain:           "640px",
			OuterPadding:           "32px",
			OutsideContentInset:    "20px",
			RadiusMain:             "0px",
			RadiusElement:          "0px",
			RadiusButton:           "0px",
			RadiusPill:             "0px",
			FontFamilyBase:         "ui-monospace,SFMono-Regular,Menlo,Monaco,Consolas,'Liberation Mono','Courier New',monospace",
			FontFamilyMono:         "ui-monospace,SFMono-Regular,Menlo,Monaco,Consolas,'Liberation Mono','Courier New',monospace",
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
	return "terminal"
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
