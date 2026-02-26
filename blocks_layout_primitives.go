package myrtle

import (
	"strings"

	"github.com/gzuidhof/myrtle/theme"
)

type DividerBlock struct {
	Variant   DividerVariant
	Thickness int
	Inset     int
	Label     string
}

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
