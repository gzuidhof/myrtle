package myrtle

import (
	"fmt"
	"strings"

	"github.com/gzuidhof/myrtle/theme"
)

type SectionBlock struct {
	Title    string
	Subtitle string
	Padding  int
	Border   bool
	Blocks   []Block
}

type GridItem struct {
	Blocks []Block
}

type GridBlock struct {
	Columns int
	Gap     int
	Border  bool
	Items   []GridItem
}

type CardItem struct {
	Title    string
	Body     string
	Subtitle string
	URL      string
	CTALabel string
}

type CardListBlock struct {
	Columns int
	Gap     int
	Border  bool
	Cards   []CardItem
}

func (block SectionBlock) Kind() theme.BlockKind {
	return theme.BlockKindSection
}

func (block SectionBlock) TemplateData() any {
	normalized := block
	if normalized.Padding <= 0 {
		normalized.Padding = 16
	}
	if !normalized.Border {
		normalized.Border = true
	}
	return normalized
}

func (block SectionBlock) RenderMarkdown(context RenderContext) (string, error) {
	parts := make([]string, 0, 3)

	if title := strings.TrimSpace(block.Title); title != "" {
		parts = append(parts, "## "+title)
	}
	if subtitle := strings.TrimSpace(block.Subtitle); subtitle != "" {
		parts = append(parts, subtitle)
	}

	body, err := renderColumnMarkdown(block.Blocks, context)
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
		Columns: normalizedGridColumns(block.Columns),
		Gap:     normalizedGridGap(block.Gap),
		Border:  block.Border,
		Items:   make([]GridItem, 0, len(block.Items)),
	}

	for _, item := range block.Items {
		if len(item.Blocks) == 0 {
			continue
		}
		normalized.Items = append(normalized.Items, GridItem{Blocks: append([]Block(nil), item.Blocks...)})
	}

	return normalized
}

func (block GridBlock) RenderMarkdown(context RenderContext) (string, error) {
	normalized := block.TemplateData().(GridBlock)
	parts := make([]string, 0, len(normalized.Items))

	for index, item := range normalized.Items {
		body, err := renderColumnMarkdown(item.Blocks, context)
		if err != nil {
			return "", err
		}
		if strings.TrimSpace(body) == "" {
			continue
		}

		parts = append(parts, fmt.Sprintf("### Grid item %d\n\n%s", index+1, body))
	}

	return strings.Join(parts, "\n\n"), nil
}

func (block CardListBlock) Kind() theme.BlockKind {
	return theme.BlockKindCardList
}

func (block CardListBlock) TemplateData() any {
	normalized := CardListBlock{
		Columns: normalizedGridColumns(block.Columns),
		Gap:     normalizedGridGap(block.Gap),
		Border:  block.Border,
		Cards:   make([]CardItem, 0, len(block.Cards)),
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

func (block CardListBlock) RenderMarkdown(_ RenderContext) (string, error) {
	normalized := block.TemplateData().(CardListBlock)
	parts := make([]string, 0, len(normalized.Cards))

	for _, card := range normalized.Cards {
		line := "- "
		if card.Title != "" {
			line += "**" + card.Title + "**"
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
			line += fmt.Sprintf(" ([%s](%s))", label, card.URL)
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
	if value <= 0 {
		return 12
	}

	return value
}
