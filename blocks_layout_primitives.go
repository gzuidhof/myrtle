package myrtle

import "github.com/gzuidhof/myrtle/theme"

type DividerBlock struct{}

func (block DividerBlock) Kind() theme.BlockKind {
	return theme.BlockKindDivider
}

func (block DividerBlock) TemplateData() any {
	return block
}

func (block DividerBlock) RenderMarkdown(_ RenderContext) (string, error) {
	return "---", nil
}
