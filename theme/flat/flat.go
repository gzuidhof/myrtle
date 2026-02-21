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
	}

	for _, option := range options {
		option(themeImpl)
	}

	return themeImpl
}

func (themeImpl *Theme) Name() string {
	return "flat"
}

func (themeImpl *Theme) RenderHTML(view theme.EmailView) (string, error) {
	return themerender.ExecuteTemplate(themeImpl.htmlTemplates, "layout.html.tmpl", view)
}

func (themeImpl *Theme) RenderBlockHTML(view theme.BlockView) (string, bool, error) {
	return themerender.RenderBlockHTMLWithHandlers(themeImpl.htmlTemplates, view, themeImpl.handlers, themeImpl.fallback)
}

func (themeImpl *Theme) WrapMarkdown(view theme.TextView) (string, error) {
	if wrapper, ok := themeImpl.fallback.(theme.MarkdownWrapper); ok {
		return wrapper.WrapMarkdown(view)
	}

	return themerender.ExecuteTextTemplate(themeImpl.textTemplates, "layout.text.tmpl", view)
}
