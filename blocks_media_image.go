package myrtle

import (
	"fmt"
	"strings"

	"github.com/gzuidhof/myrtle/theme"
)

// ImageBlock renders an image with optional link and spacing controls.
type ImageBlock struct {
	Src              string
	Alt              string
	Href             string
	Width            int            // px, 0 means auto
	Align            ImageAlignment // "center", "start", "end", "full", "" (default)
	HasTopSpacing    bool
	TopSpacing       int
	HasBottomSpacing bool
	BottomSpacing    int
	CornerMode       ImageCornerMode
	InsetMode        InsetMode
}

// ImageCornerMode controls how image corner radii are applied.
type ImageCornerMode string

const (
	ImageCornerModeAuto   ImageCornerMode = "auto"
	ImageCornerModeNone   ImageCornerMode = "none"
	ImageCornerModeAll    ImageCornerMode = "all"
	ImageCornerModeTop    ImageCornerMode = "top"
	ImageCornerModeBottom ImageCornerMode = "bottom"
)

// ImageAlignment controls image horizontal alignment and width behavior.
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
	if normalized.HasTopSpacing && normalized.TopSpacing < 0 {
		normalized.TopSpacing = 0
	}
	if normalized.HasBottomSpacing && normalized.BottomSpacing < 0 {
		normalized.BottomSpacing = 0
	}
	normalized.CornerMode = normalizedImageCornerMode(normalized.CornerMode)
	normalized.InsetMode = normalizedLayoutSpec(LayoutSpec{InsetMode: normalized.InsetMode}).InsetMode
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

func normalizedImageCornerMode(value ImageCornerMode) ImageCornerMode {
	switch value {
	case ImageCornerModeNone, ImageCornerModeAll, ImageCornerModeTop, ImageCornerModeBottom:
		return value
	default:
		return ImageCornerModeAuto
	}
}

func (block ImageBlock) LayoutSpec() LayoutSpec {
	return normalizedLayoutSpec(LayoutSpec{InsetMode: block.InsetMode})
}
