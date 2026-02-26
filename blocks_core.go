package myrtle

import "github.com/gzuidhof/myrtle/theme"

type RenderContext struct {
	Preheader string
	Values    theme.Values
}

// Block is the core content unit in an email, with HTML template data and text fallback rendering behavior.
type Block interface {
	Kind() theme.BlockKind
	TemplateData() any
	RenderText(context RenderContext) (string, error)
}
