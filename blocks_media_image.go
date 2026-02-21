package myrtle

import (
	"fmt"

	"github.com/gzuidhof/myrtle/theme"
)

type ImageBlock struct {
	Src string
	Alt string
}

func (block ImageBlock) Kind() theme.BlockKind {
	return theme.BlockKindImage
}

func (block ImageBlock) TemplateData() any {
	return block
}

func (block ImageBlock) RenderMarkdown(_ RenderContext) (string, error) {
	return fmt.Sprintf("![%s](%s)", block.Alt, block.Src), nil
}
