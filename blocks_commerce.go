package myrtle

import (
	"strings"

	"github.com/gzuidhof/myrtle/theme"
)

type PriceLine struct {
	Label string
	Value string
}

// PriceSummaryBlock renders a line-item pricing summary with totals.
type PriceSummaryBlock struct {
	Header     string
	Items      []PriceLine
	TotalLabel string
	TotalValue string
	InsetMode  InsetMode
}

func (block PriceSummaryBlock) Kind() theme.BlockKind {
	return theme.BlockKindPriceSummary
}

func (block PriceSummaryBlock) TemplateData() any {
	return block
}

func (block PriceSummaryBlock) RenderText(_ RenderContext) (string, error) {
	parts := make([]string, 0, len(block.Items)+2)
	if strings.TrimSpace(block.Header) != "" {
		header := strings.TrimSpace(block.Header)
		parts = append(parts, header, strings.Repeat("-", min(48, max(8, len(header)))))
	}

	for _, item := range block.Items {
		label := strings.TrimSpace(item.Label)
		value := strings.TrimSpace(item.Value)
		if label == "" && value == "" {
			continue
		}

		parts = append(parts, "- "+label+": "+value)
	}

	if strings.TrimSpace(block.TotalLabel) != "" || strings.TrimSpace(block.TotalValue) != "" {
		parts = append(parts, strings.TrimSpace(block.TotalLabel)+": "+strings.TrimSpace(block.TotalValue))
	}

	return strings.Join(parts, "\n"), nil
}

// EmptyStateBlock renders placeholder content when data is unavailable.
type EmptyStateBlock struct {
	Title       string
	Body        string
	ActionLabel string
	ActionURL   string
	Tone        Tone
	InsetMode   InsetMode
}

func (block EmptyStateBlock) Kind() theme.BlockKind {
	return theme.BlockKindEmptyState
}

func (block EmptyStateBlock) TemplateData() any {
	normalized := block
	normalized.Tone = normalizedEmptyStateTone(block.Tone)

	return normalized
}

func (block EmptyStateBlock) RenderText(_ RenderContext) (string, error) {
	parts := make([]string, 0, 3)
	if strings.TrimSpace(block.Title) != "" {
		title := strings.TrimSpace(block.Title)
		parts = append(parts, "[ "+title+" ]", strings.Repeat("-", min(48, max(8, len(title)+4))))
	}
	if strings.TrimSpace(block.Body) != "" {
		parts = append(parts, strings.TrimSpace(block.Body))
	}
	if strings.TrimSpace(block.ActionLabel) != "" && strings.TrimSpace(block.ActionURL) != "" {
		parts = append(parts, strings.TrimSpace(block.ActionLabel)+" ("+strings.TrimSpace(block.ActionURL)+")")
	}

	return strings.Join(parts, "\n\n"), nil
}

func (block PriceSummaryBlock) LayoutSpec() LayoutSpec {
	return normalizedLayoutSpec(LayoutSpec{InsetMode: block.InsetMode})
}

func (block EmptyStateBlock) LayoutSpec() LayoutSpec {
	return normalizedLayoutSpec(LayoutSpec{InsetMode: block.InsetMode})
}

func normalizedEmptyStateTone(value Tone) Tone {
	return normalizedTone(value)
}
