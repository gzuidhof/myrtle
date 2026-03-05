package myrtle

import (
	"strings"

	"github.com/gzuidhof/myrtle/theme"
)

// ColumnsBlock renders two side-by-side block columns with configurable widths.
type ColumnsBlock struct {
	Left          []Block
	Right         []Block
	LeftWidth     int
	RightWidth    int
	Gap           int
	VerticalAlign ColumnsVerticalAlign
	InsetMode     InsetMode
}

// ColumnsVerticalAlign controls vertical alignment of content inside each columns cell.
type ColumnsVerticalAlign string

const (
	ColumnsVerticalAlignTop    ColumnsVerticalAlign = "top"
	ColumnsVerticalAlignMiddle ColumnsVerticalAlign = "middle"
	ColumnsVerticalAlignBottom ColumnsVerticalAlign = "bottom"
)

func (block ColumnsBlock) Kind() theme.BlockKind {
	return theme.BlockKindColumns
}

func (block ColumnsBlock) TemplateData() any {
	normalized := block
	if normalized.Gap < 0 {
		normalized.Gap = 0
	}
	if normalized.VerticalAlign != ColumnsVerticalAlignMiddle && normalized.VerticalAlign != ColumnsVerticalAlignBottom {
		normalized.VerticalAlign = ColumnsVerticalAlignTop
	}

	return normalized
}

func (block ColumnsBlock) RenderText(context RenderContext) (string, error) {
	left, err := renderColumnText(block.Left, context)
	if err != nil {
		return "", err
	}

	right, err := renderColumnText(block.Right, context)
	if err != nil {
		return "", err
	}

	parts := make([]string, 0, 2)
	if strings.TrimSpace(left) != "" {
		parts = append(parts, "[ Column 1 ]\n--------------------\n"+left)
	}
	if strings.TrimSpace(right) != "" {
		parts = append(parts, "[ Column 2 ]\n--------------------\n"+right)
	}

	return strings.Join(parts, "\n\n"), nil
}

func renderColumnText(blocks []Block, context RenderContext) (string, error) {
	parts := make([]string, 0, len(blocks))
	for _, block := range blocks {
		if block == nil {
			continue
		}

		text, err := block.RenderText(context)
		if err != nil {
			return "", err
		}
		if strings.TrimSpace(text) == "" {
			continue
		}

		parts = append(parts, text)
	}

	return strings.Join(parts, "\n\n"), nil
}

func (block ColumnsBlock) LayoutSpec() LayoutSpec {
	return normalizedLayoutSpec(LayoutSpec{InsetMode: block.InsetMode})
}
