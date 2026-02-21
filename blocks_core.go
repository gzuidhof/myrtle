package myrtle

import "github.com/gzuidhof/myrtle/theme"

type RenderContext struct {
	Preheader string
	Values    theme.Values
}

type Block interface {
	Kind() theme.BlockKind
	TemplateData() any
	RenderMarkdown(context RenderContext) (string, error)
}
