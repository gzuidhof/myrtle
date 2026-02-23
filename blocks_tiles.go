package myrtle

import (
	"fmt"
	"strings"

	"github.com/gzuidhof/myrtle/theme"
)

type TileVariant string

const (
	TileVariantDefault   TileVariant = "default"
	TileVariantHighlight TileVariant = "highlight"
	TileVariantSuccess   TileVariant = "success"
	TileVariantWarning   TileVariant = "warning"
	TileVariantCritical  TileVariant = "critical"
)

type TileAlignment string

const (
	TileAlignmentCenter TileAlignment = "center"
	TileAlignmentLeft   TileAlignment = "left"
	TileAlignmentRight  TileAlignment = "right"
)

type TileEntry struct {
	Content  string
	Title    string
	Subtitle string
	URL      string
	Variant  TileVariant
}

type TilesBlock struct {
	Columns               int
	Border                bool
	TransparentBackground bool
	Alignment             TileAlignment
	Entries               []TileEntry
}

func (block TilesBlock) Kind() theme.BlockKind {
	return theme.BlockKindTiles
}

func (block TilesBlock) TemplateData() any {
	normalized := block
	normalized.Columns = normalizedTilesColumns(block.Columns)
	normalized.Alignment = normalizedTileAlignment(block.Alignment)
	normalized.Entries = make([]TileEntry, 0, len(block.Entries))

	for _, entry := range block.Entries {
		content := strings.TrimSpace(entry.Content)
		title := strings.TrimSpace(entry.Title)
		subtitle := strings.TrimSpace(entry.Subtitle)
		url := strings.TrimSpace(entry.URL)
		if content == "" && title == "" && subtitle == "" {
			continue
		}

		normalized.Entries = append(normalized.Entries, TileEntry{
			Content:  content,
			Title:    title,
			Subtitle: subtitle,
			URL:      url,
			Variant:  normalizedTileVariant(entry.Variant),
		})
	}

	return normalized
}

func (block TilesBlock) RenderMarkdown(_ RenderContext) (string, error) {
	normalized := block.TemplateData().(TilesBlock)
	parts := make([]string, 0, len(normalized.Entries))

	for _, entry := range normalized.Entries {
		line := "- "
		if entry.Content != "" {
			line += fmt.Sprintf("[%s]", entry.Content)
		}
		if entry.Title != "" {
			if entry.Content != "" {
				line += " "
			}
			if entry.URL != "" {
				line += fmt.Sprintf("[**%s**](%s)", entry.Title, entry.URL)
			} else {
				line += "**" + entry.Title + "**"
			}
		}
		if entry.Subtitle != "" {
			if entry.Content != "" || entry.Title != "" {
				line += " — "
			}
			line += entry.Subtitle
		}

		parts = append(parts, strings.TrimSpace(line))
	}

	return strings.Join(parts, "\n"), nil
}

func normalizedTileVariant(value TileVariant) TileVariant {
	switch value {
	case TileVariantHighlight, TileVariantSuccess, TileVariantWarning, TileVariantCritical:
		return value
	default:
		return TileVariantDefault
	}
}

func normalizedTilesColumns(value int) int {
	if value <= 0 {
		return 3
	}
	if value > 6 {
		return 6
	}

	return value
}

func normalizedTileAlignment(value TileAlignment) TileAlignment {
	switch value {
	case TileAlignmentLeft, TileAlignmentRight:
		return value
	default:
		return TileAlignmentCenter
	}
}
