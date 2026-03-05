package myrtle

import (
	"strings"

	"github.com/gzuidhof/myrtle/theme"
)

// DividerBlock renders a horizontal divider line with optional label.
type DividerBlock struct {
	Variant   DividerVariant
	Thickness int
	Inset     int
	Label     string
	InsetMode InsetMode
}

// DividerVariant defines the line style used by a divider.
type DividerVariant string

const (
	DividerVariantSolid  DividerVariant = "solid"
	DividerVariantDashed DividerVariant = "dashed"
	DividerVariantDotted DividerVariant = "dotted"
)

func (block DividerBlock) Kind() theme.BlockKind {
	return theme.BlockKindDivider
}

func (block DividerBlock) TemplateData() any {
	normalized := block
	if normalized.Thickness <= 0 {
		normalized.Thickness = 1
	}
	if normalized.Inset < 0 {
		normalized.Inset = 0
	}
	if normalized.Variant != DividerVariantDashed && normalized.Variant != DividerVariantDotted {
		normalized.Variant = DividerVariantSolid
	}
	normalized.InsetMode = normalizedLayoutSpec(LayoutSpec{InsetMode: normalized.InsetMode}).InsetMode
	normalized.Label = strings.TrimSpace(normalized.Label)

	return normalized
}

func (block DividerBlock) RenderText(_ RenderContext) (string, error) {
	label := strings.TrimSpace(block.Label)
	if label != "" {
		return "----- " + label + " -----", nil
	}

	return "--------------------", nil
}

func (block DividerBlock) LayoutSpec() LayoutSpec {
	return normalizedLayoutSpec(LayoutSpec{InsetMode: block.InsetMode})
}
