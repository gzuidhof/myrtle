package myrtle

import (
	"fmt"
	"strings"

	"github.com/gzuidhof/myrtle/theme"
)

// PanelBlock renders a bordered container around nested blocks.
type PanelBlock struct {
	Title      string
	Subtitle   string
	Category   string
	Headerless bool
	ShowHeader bool
	Padding    int
	Border     bool
	Blocks     []Block
	InsetMode  InsetMode
}

type GridItem struct {
	Content Block
}

// GridBlock renders grid items in configurable columns.
type GridBlock struct {
	Columns   int
	Gap       int
	Border    bool
	Items     []GridItem
	InsetMode InsetMode
}

type CardItem struct {
	Title    string
	Body     string
	Subtitle string
	URL      string
	CTALabel string
}

// CardListBlock renders a multi-column list of simple cards.
type CardListBlock struct {
	Columns   int
	Gap       int
	Border    bool
	Cards     []CardItem
	InsetMode InsetMode
}

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

// ColumnsOption configures a ColumnsBlock when building content.
type ColumnsOption func(*ColumnsBlock)

// ColumnsVerticalAlign controls vertical alignment of content inside each columns cell.
type ColumnsVerticalAlign string

const (
	ColumnsVerticalAlignTop    ColumnsVerticalAlign = "top"
	ColumnsVerticalAlignMiddle ColumnsVerticalAlign = "middle"
	ColumnsVerticalAlignBottom ColumnsVerticalAlign = "bottom"
)

func (block PanelBlock) Kind() theme.BlockKind {
	return theme.BlockKindPanel
}

func (block PanelBlock) TemplateData() any {
	normalized := block
	normalized.Title = strings.TrimSpace(normalized.Title)
	normalized.Subtitle = strings.TrimSpace(normalized.Subtitle)
	normalized.Category = strings.TrimSpace(normalized.Category)
	normalized.ShowHeader = !normalized.Headerless && (normalized.Category != "" || normalized.Title != "" || normalized.Subtitle != "")
	if normalized.Padding <= 0 {
		normalized.Padding = 16
	}
	if !normalized.Border {
		normalized.Border = true
	}
	return normalized
}

func (block PanelBlock) RenderText(context RenderContext) (string, error) {
	parts := make([]string, 0, 3)
	if category := strings.TrimSpace(block.Category); category != "" {
		parts = append(parts, strings.ToUpper(category))
	}

	if title := strings.TrimSpace(block.Title); title != "" {
		parts = append(parts, "[ "+title+" ]", strings.Repeat("-", min(48, max(8, len(title)+4))))
	}
	if subtitle := strings.TrimSpace(block.Subtitle); subtitle != "" {
		parts = append(parts, subtitle)
	}

	body, err := renderColumnText(block.Blocks, context)
	if err != nil {
		return "", err
	}
	if strings.TrimSpace(body) != "" {
		parts = append(parts, body)
	}

	return strings.Join(parts, "\n\n"), nil
}

func (block GridBlock) Kind() theme.BlockKind {
	return theme.BlockKindGrid
}

func (block GridBlock) TemplateData() any {
	normalized := GridBlock{
		Columns:   normalizedGridColumns(block.Columns),
		Gap:       normalizedGridGap(block.Gap),
		Border:    block.Border,
		Items:     make([]GridItem, 0, len(block.Items)),
		InsetMode: normalizedLayoutSpec(LayoutSpec{InsetMode: block.InsetMode}).InsetMode,
	}

	for _, item := range block.Items {
		if item.Content == nil {
			continue
		}
		normalized.Items = append(normalized.Items, GridItem{Content: item.Content})
	}

	return normalized
}

func (block GridBlock) RenderText(context RenderContext) (string, error) {
	normalized := block.TemplateData().(GridBlock)
	parts := make([]string, 0, len(normalized.Items))

	for index, item := range normalized.Items {
		body, err := item.Content.RenderText(context)
		if err != nil {
			return "", err
		}
		if strings.TrimSpace(body) == "" {
			continue
		}

		parts = append(parts, fmt.Sprintf("[ Grid item %d ]\n--------------------\n%s", index+1, body))
	}

	return strings.Join(parts, "\n\n"), nil
}

func (block CardListBlock) Kind() theme.BlockKind {
	return theme.BlockKindCardList
}

func (block CardListBlock) TemplateData() any {
	normalized := CardListBlock{
		Columns:   normalizedGridColumns(block.Columns),
		Gap:       normalizedGridGap(block.Gap),
		Border:    block.Border,
		Cards:     make([]CardItem, 0, len(block.Cards)),
		InsetMode: normalizedLayoutSpec(LayoutSpec{InsetMode: block.InsetMode}).InsetMode,
	}

	for _, card := range block.Cards {
		title := strings.TrimSpace(card.Title)
		body := strings.TrimSpace(card.Body)
		subtitle := strings.TrimSpace(card.Subtitle)
		url := strings.TrimSpace(card.URL)
		label := strings.TrimSpace(card.CTALabel)

		if title == "" && body == "" && subtitle == "" {
			continue
		}

		normalized.Cards = append(normalized.Cards, CardItem{
			Title:    title,
			Body:     body,
			Subtitle: subtitle,
			URL:      url,
			CTALabel: label,
		})
	}

	return normalized
}

func (block CardListBlock) RenderText(_ RenderContext) (string, error) {
	normalized := block.TemplateData().(CardListBlock)
	parts := make([]string, 0, len(normalized.Cards))

	for _, card := range normalized.Cards {
		line := "- "
		if card.Title != "" {
			line += card.Title
		}
		if card.Body != "" {
			if card.Title != "" {
				line += " — "
			}
			line += card.Body
		}
		if card.URL != "" {
			label := card.CTALabel
			if label == "" {
				label = "Open"
			}
			line += fmt.Sprintf(" (%s: %s)", label, card.URL)
		}

		parts = append(parts, strings.TrimSpace(line))
	}

	return strings.Join(parts, "\n"), nil
}

func normalizedGridColumns(value int) int {
	if value <= 0 {
		return 2
	}
	if value > 4 {
		return 4
	}

	return value
}

func normalizedGridGap(value int) int {
	if value < 0 {
		return 12
	}

	return value
}

func (block PanelBlock) LayoutSpec() LayoutSpec {
	return normalizedLayoutSpec(LayoutSpec{InsetMode: block.InsetMode})
}

func (block GridBlock) LayoutSpec() LayoutSpec {
	return normalizedLayoutSpec(LayoutSpec{InsetMode: block.InsetMode})
}

func (block CardListBlock) LayoutSpec() LayoutSpec {
	return normalizedLayoutSpec(LayoutSpec{InsetMode: block.InsetMode})
}

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
