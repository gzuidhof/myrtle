package myrtle

import (
	"fmt"
	"strings"

	"github.com/gzuidhof/myrtle/theme"
)

type KeyValuePair struct {
	Key   string
	Value string
}

type KeyValueBlock struct {
	Header string
	Pairs  []KeyValuePair
}

func (block KeyValueBlock) Kind() theme.BlockKind {
	return theme.BlockKindKeyValue
}

func (block KeyValueBlock) TemplateData() any {
	return block
}

func (block KeyValueBlock) RenderText(_ RenderContext) (string, error) {
	parts := make([]string, 0, len(block.Pairs)+1)
	if strings.TrimSpace(block.Header) != "" {
		header := strings.TrimSpace(block.Header)
		parts = append(parts, header, strings.Repeat("-", min(48, max(8, len(header)))))
	}
	for _, pair := range block.Pairs {
		key := strings.TrimSpace(pair.Key)
		value := strings.TrimSpace(pair.Value)
		if key == "" && value == "" {
			continue
		}
		if key == "" {
			parts = append(parts, value)
			continue
		}
		parts = append(parts, fmt.Sprintf("- %s: %s", key, value))
	}
	return strings.Join(parts, "\n"), nil
}

type HorizontalBarChartItem struct {
	Label   string
	Value   string
	Percent int
	Color   string
}

type HorizontalBarChartBlock struct {
	Header                string
	Items                 []HorizontalBarChartItem
	Thickness             int
	TransparentBackground bool
	Tone                  ChartToneValue
}

func (block HorizontalBarChartBlock) Kind() theme.BlockKind {
	return theme.BlockKindHorizontalBarChart
}

func (block HorizontalBarChartBlock) TemplateData() any {
	normalized := block
	normalized.Items = block.normalizedItems()
	normalized.Thickness = block.normalizedThickness()
	normalized.Tone = normalizedChartTone(block.Tone)
	return normalized
}

func (block HorizontalBarChartBlock) RenderText(_ RenderContext) (string, error) {
	parts := make([]string, 0, len(block.Items)+1)
	if strings.TrimSpace(block.Header) != "" {
		header := strings.TrimSpace(block.Header)
		parts = append(parts, header, strings.Repeat("-", min(48, max(8, len(header)))))
	}

	for _, item := range block.normalizedItems() {
		filled := (item.Percent + 9) / 10
		if filled < 0 {
			filled = 0
		}
		if filled > 10 {
			filled = 10
		}
		empty := 10 - filled

		parts = append(parts, fmt.Sprintf("- %s: %s %s%s", item.Label, item.Value, strings.Repeat("#", filled), strings.Repeat(".", empty)))
	}

	return strings.Join(parts, "\n"), nil
}

func (block HorizontalBarChartBlock) normalizedItems() []HorizontalBarChartItem {
	items := make([]HorizontalBarChartItem, 0, len(block.Items))
	for _, item := range block.Items {
		label := strings.TrimSpace(item.Label)
		if label == "" {
			continue
		}

		percent := item.Percent
		if percent < 0 {
			percent = 0
		}
		if percent > 100 {
			percent = 100
		}

		value := strings.TrimSpace(item.Value)
		if value == "" {
			value = fmt.Sprintf("%d%%", percent)
		}

		items = append(items, HorizontalBarChartItem{
			Label:   label,
			Value:   value,
			Percent: percent,
			Color:   strings.TrimSpace(item.Color),
		})
	}

	return items
}

func (block HorizontalBarChartBlock) normalizedThickness() int {
	if block.Thickness <= 0 {
		return 8
	}
	if block.Thickness > 24 {
		return 24
	}
	return block.Thickness
}
