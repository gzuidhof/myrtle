package myrtle

import (
	"strings"

	"github.com/gzuidhof/myrtle/theme"
)

type TextBlock struct {
	Text string
}

func (block TextBlock) Kind() theme.BlockKind {
	return theme.BlockKindText
}

func (block TextBlock) TemplateData() any {
	return block
}

func (block TextBlock) RenderMarkdown(_ RenderContext) (string, error) {
	return strings.TrimSpace(block.Text), nil
}
