package myrtle

import (
	"fmt"
	"strings"

	"github.com/gzuidhof/myrtle/theme"
)

type ImageBlock struct {
	Src   string
	Alt   string
	Href  string
	Width int            // px, 0 means auto
	Align ImageAlignment // "center", "start", "end", "full", "" (default)
}

type ImageAlignment string

const (
	ImageAlignmentDefault ImageAlignment = ""
	ImageAlignmentCenter  ImageAlignment = "center"
	ImageAlignmentStart   ImageAlignment = "start"
	ImageAlignmentEnd     ImageAlignment = "end"
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

func (block ImageBlock) RenderText(_ RenderContext) (string, error) {
	alt := strings.TrimSpace(block.Alt)
	src := strings.TrimSpace(block.Src)
	if alt == "" && src == "" {
		return "", nil
	}
	if alt == "" {
		return src, nil
	}
	if src == "" {
		return alt, nil
	}

	return fmt.Sprintf("%s (%s)", alt, src), nil
}

func normalizedImageAlignment(value ImageAlignment) ImageAlignment {
	switch value {
	case ImageAlignmentCenter, ImageAlignmentStart, ImageAlignmentEnd, ImageAlignmentFull:
		return value
	default:
		return ImageAlignmentDefault
	}
}
