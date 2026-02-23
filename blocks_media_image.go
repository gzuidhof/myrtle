package myrtle

import (
	"fmt"

	"github.com/gzuidhof/myrtle/theme"
)

type ImageBlock struct {
	Src   string
	Alt   string
	Width int            // px, 0 means auto
	Align ImageAlignment // "center", "left", "right", "full", "" (default)
}

type ImageAlignment string

const (
	ImageAlignmentDefault ImageAlignment = ""
	ImageAlignmentCenter  ImageAlignment = "center"
	ImageAlignmentLeft    ImageAlignment = "left"
	ImageAlignmentRight   ImageAlignment = "right"
	ImageAlignmentFull    ImageAlignment = "full"
)

func (block ImageBlock) Kind() theme.BlockKind {
	return theme.BlockKindImage
}

func (block ImageBlock) TemplateData() any {
	normalized := block
	normalized.Align = normalizedImageAlignment(block.Align)
	return normalized
}

func (block ImageBlock) RenderMarkdown(_ RenderContext) (string, error) {
	return fmt.Sprintf("![%s](%s)", block.Alt, block.Src), nil
}

func normalizedImageAlignment(value ImageAlignment) ImageAlignment {
	switch value {
	case ImageAlignmentCenter, ImageAlignmentLeft, ImageAlignmentRight, ImageAlignmentFull:
		return value
	default:
		return ImageAlignmentDefault
	}
}
