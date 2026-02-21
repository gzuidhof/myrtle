package myrtle

import (
	"fmt"
	"strings"

	"github.com/gzuidhof/myrtle/theme"
)

type SparklineBlock struct {
	Header        string
	Label         string
	Value         string
	Delta         string
	DeltaSemantic StatDeltaSemantic
	Points        []int
}

func (block SparklineBlock) Kind() theme.BlockKind {
	return theme.BlockKindSparkline
}

func (block SparklineBlock) TemplateData() any {
	normalized := block
	normalized.Points = normalizedIntPoints(block.Points)
	normalized.DeltaSemantic = normalizedStatDeltaSemantic(block.DeltaSemantic)
	return normalized
}

func (block SparklineBlock) RenderMarkdown(_ RenderContext) (string, error) {
	normalized := block.TemplateData().(SparklineBlock)
	parts := make([]string, 0, 3)
	if strings.TrimSpace(normalized.Header) != "" {
		parts = append(parts, "### "+strings.TrimSpace(normalized.Header))
	}

	line := strings.TrimSpace(normalized.Label)
	if strings.TrimSpace(normalized.Value) != "" {
		if line != "" {
			line += ": "
		}
		line += strings.TrimSpace(normalized.Value)
	}
	if strings.TrimSpace(normalized.Delta) != "" {
		line += " (" + strings.TrimSpace(normalized.Delta) + ")"
	}
	if strings.TrimSpace(line) != "" {
		parts = append(parts, line)
	}

	if len(normalized.Points) > 0 {
		parts = append(parts, sparklineGlyphs(normalized.Points))
	}

	return strings.Join(parts, "\n"), nil
}

type StackedBarSegment struct {
	Label   string
	Percent int
	Value   string
}

type StackedBarRow struct {
	Label    string
	Segments []StackedBarSegment
}

type StackedBarBlock struct {
	Header     string
	TotalLabel string
	TotalValue string
	Rows       []StackedBarRow
}

func (block StackedBarBlock) Kind() theme.BlockKind {
	return theme.BlockKindStackedBar
}

func (block StackedBarBlock) TemplateData() any {
	normalized := block
	normalized.Rows = make([]StackedBarRow, 0, len(block.Rows))
	for _, row := range block.Rows {
		segments := make([]StackedBarSegment, 0, len(row.Segments))
		for _, segment := range row.Segments {
			percent := segment.Percent
			if percent < 0 {
				percent = 0
			}
			if percent > 100 {
				percent = 100
			}

			segments = append(segments, StackedBarSegment{
				Label:   strings.TrimSpace(segment.Label),
				Percent: percent,
				Value:   strings.TrimSpace(segment.Value),
			})
		}

		normalized.Rows = append(normalized.Rows, StackedBarRow{
			Label:    strings.TrimSpace(row.Label),
			Segments: segments,
		})
	}

	return normalized
}

func (block StackedBarBlock) RenderMarkdown(_ RenderContext) (string, error) {
	normalized := block.TemplateData().(StackedBarBlock)
	parts := make([]string, 0, len(normalized.Rows)+2)
	if strings.TrimSpace(normalized.Header) != "" {
		parts = append(parts, "### "+strings.TrimSpace(normalized.Header))
	}
	if strings.TrimSpace(normalized.TotalLabel) != "" || strings.TrimSpace(normalized.TotalValue) != "" {
		parts = append(parts, fmt.Sprintf("**%s:** %s", strings.TrimSpace(normalized.TotalLabel), strings.TrimSpace(normalized.TotalValue)))
	}

	for _, row := range normalized.Rows {
		if len(row.Segments) == 0 {
			continue
		}

		segmentParts := make([]string, 0, len(row.Segments))
		for _, segment := range row.Segments {
			if segment.Label == "" {
				continue
			}
			value := segment.Value
			if value == "" {
				value = fmt.Sprintf("%d%%", segment.Percent)
			}
			segmentParts = append(segmentParts, fmt.Sprintf("%s %s", segment.Label, value))
		}

		if len(segmentParts) == 0 {
			continue
		}

		label := row.Label
		if label == "" {
			parts = append(parts, "- "+strings.Join(segmentParts, " · "))
			continue
		}
		parts = append(parts, fmt.Sprintf("- **%s:** %s", label, strings.Join(segmentParts, " · ")))
	}

	return strings.Join(parts, "\n"), nil
}

type ProgressItem struct {
	Label   string
	Percent int
	Value   string
}

type ProgressBlock struct {
	Header string
	Items  []ProgressItem
}

func (block ProgressBlock) Kind() theme.BlockKind {
	return theme.BlockKindProgress
}

func (block ProgressBlock) TemplateData() any {
	normalized := block
	normalized.Items = make([]ProgressItem, 0, len(block.Items))
	for _, item := range block.Items {
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

		normalized.Items = append(normalized.Items, ProgressItem{
			Label:   strings.TrimSpace(item.Label),
			Percent: percent,
			Value:   value,
		})
	}

	return normalized
}

func (block ProgressBlock) RenderMarkdown(_ RenderContext) (string, error) {
	normalized := block.TemplateData().(ProgressBlock)
	parts := make([]string, 0, len(normalized.Items)+1)
	if strings.TrimSpace(normalized.Header) != "" {
		parts = append(parts, "### "+strings.TrimSpace(normalized.Header))
	}

	for _, item := range normalized.Items {
		if item.Label == "" && item.Value == "" {
			continue
		}
		filled := (item.Percent + 9) / 10
		if filled < 0 {
			filled = 0
		}
		if filled > 10 {
			filled = 10
		}
		empty := 10 - filled
		parts = append(parts, fmt.Sprintf("- **%s:** %s %s%s", item.Label, item.Value, strings.Repeat("█", filled), strings.Repeat("░", empty)))
	}

	return strings.Join(parts, "\n"), nil
}

type DistributionBucket struct {
	Label        string
	Count        int
	WidthPercent int
}

type DistributionBlock struct {
	Header  string
	Buckets []DistributionBucket
}

func (block DistributionBlock) Kind() theme.BlockKind {
	return theme.BlockKindDistribution
}

func (block DistributionBlock) TemplateData() any {
	normalized := block
	normalized.Buckets = make([]DistributionBucket, 0, len(block.Buckets))
	maxCount := 0
	for _, bucket := range block.Buckets {
		count := bucket.Count
		if count < 0 {
			count = 0
		}
		if count > maxCount {
			maxCount = count
		}
		normalized.Buckets = append(normalized.Buckets, DistributionBucket{
			Label: strings.TrimSpace(bucket.Label),
			Count: count,
		})
	}

	if maxCount <= 0 {
		maxCount = 1
	}

	for index := range normalized.Buckets {
		width := int(float64(normalized.Buckets[index].Count) / float64(maxCount) * 100)
		if width < 0 {
			width = 0
		}
		if width > 100 {
			width = 100
		}
		normalized.Buckets[index].WidthPercent = width
	}

	return normalized
}

func (block DistributionBlock) RenderMarkdown(_ RenderContext) (string, error) {
	normalized := block.TemplateData().(DistributionBlock)
	parts := make([]string, 0, len(normalized.Buckets)+1)
	if strings.TrimSpace(normalized.Header) != "" {
		parts = append(parts, "### "+strings.TrimSpace(normalized.Header))
	}

	maxCount := 0
	for _, bucket := range normalized.Buckets {
		if bucket.Count > maxCount {
			maxCount = bucket.Count
		}
	}
	if maxCount == 0 {
		maxCount = 1
	}

	for _, bucket := range normalized.Buckets {
		if bucket.Label == "" {
			continue
		}
		filled := int(float64(bucket.Count) / float64(maxCount) * 10)
		if filled < 0 {
			filled = 0
		}
		if filled > 10 {
			filled = 10
		}
		empty := 10 - filled
		parts = append(parts, fmt.Sprintf("- %s %s%s (%d)", bucket.Label, strings.Repeat("█", filled), strings.Repeat("░", empty), bucket.Count))
	}

	return strings.Join(parts, "\n"), nil
}

func normalizedIntPoints(points []int) []int {
	result := make([]int, 0, len(points))
	for _, point := range points {
		if point < 0 {
			point = 0
		}
		result = append(result, point)
	}
	return result
}

func sparklineGlyphs(points []int) string {
	if len(points) == 0 {
		return ""
	}

	glyphs := []rune{'▁', '▂', '▃', '▄', '▅', '▆', '▇', '█'}
	minValue := points[0]
	maxValue := points[0]
	for _, point := range points {
		if point < minValue {
			minValue = point
		}
		if point > maxValue {
			maxValue = point
		}
	}

	if maxValue == minValue {
		return strings.Repeat(string(glyphs[len(glyphs)/2]), len(points))
	}

	var output strings.Builder
	for _, point := range points {
		normalized := float64(point-minValue) / float64(maxValue-minValue)
		index := int(normalized * float64(len(glyphs)-1))
		if index < 0 {
			index = 0
		}
		if index >= len(glyphs) {
			index = len(glyphs) - 1
		}
		output.WriteRune(glyphs[index])
	}

	return output.String()
}
