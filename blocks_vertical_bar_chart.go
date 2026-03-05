package myrtle

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/gzuidhof/myrtle/theme"
)

// VerticalBarChartLegendPlacementValue controls where the chart legend is rendered.
type VerticalBarChartLegendPlacementValue string

const (
	VerticalBarChartLegendNone   VerticalBarChartLegendPlacementValue = "none"
	VerticalBarChartLegendBottom VerticalBarChartLegendPlacementValue = "bottom"
)

// VerticalBarChartAxisLabelFormatValue controls axis tick label formatting.
type VerticalBarChartAxisLabelFormatValue string

const (
	VerticalBarChartAxisLabelFormatNumber  VerticalBarChartAxisLabelFormatValue = "number"
	VerticalBarChartAxisLabelFormatPercent VerticalBarChartAxisLabelFormatValue = "percent"
)

// VerticalBarChartMagnitudeSuffixValue controls compact magnitude suffix formatting.
type VerticalBarChartMagnitudeSuffixValue string

const (
	VerticalBarChartMagnitudeSuffixNone  VerticalBarChartMagnitudeSuffixValue = "none"
	VerticalBarChartMagnitudeSuffixShort VerticalBarChartMagnitudeSuffixValue = "short"
)

// VerticalBarChartNegativeFormatValue controls negative number formatting style.
type VerticalBarChartNegativeFormatValue string

const (
	VerticalBarChartNegativeFormatMinus       VerticalBarChartNegativeFormatValue = "minus"
	VerticalBarChartNegativeFormatParentheses VerticalBarChartNegativeFormatValue = "parentheses"
)

// VerticalBarChartValueFormatter controls numeric formatting for axis labels and value labels.
// Prefix and Suffix are applied around formatted values, MagnitudeSuffix enables compact scaling
// (for example 1200 -> 1.2K), and NegativeFormat controls how negative numbers are rendered.
type VerticalBarChartValueFormatter struct {
	Prefix          string
	Suffix          string
	MagnitudeSuffix VerticalBarChartMagnitudeSuffixValue
	NegativeFormat  VerticalBarChartNegativeFormatValue
}

// VerticalBarChartSeries defines one stacked series across all chart columns.
// Values are aligned by index with AxisLabels, and optional Color/ValueLabelColor override
// theme-derived defaults for segment fills and in-segment value labels.
type VerticalBarChartSeries struct {
	Key             string
	Label           string
	Color           string
	ValueLabelColor string
	Values          []float64
}

// VerticalBarChartLegendItem represents a single legend entry for a series color/label pair.
type VerticalBarChartLegendItem struct {
	Label string
	Color string
}

// VerticalBarChartAxis configures baseline, tick visibility, label formatting, and range hints.
// Min can force the lower bound, while HasMax/Max can raise the upper bound when needed.
// Category labels and Y-axis line rendering can be independently toggled.
type VerticalBarChartAxis struct {
	ShowBaseline          bool
	ShowYTicks            bool
	HasDrawYAxisLine      bool
	DrawYAxisLine         bool
	HasShowCategoryLabels bool
	ShowCategoryLabels    bool
	LabelFormat           VerticalBarChartAxisLabelFormatValue
	HasMin                bool
	Min                   float64
	// HasMax enables a configured upper bound hint for the chart range.
	// When true, Max is only used to raise the computed max range if needed.
	// It does not clamp bars to a lower maximum than the data-derived max.
	HasMax bool
	// Max is the optional upper bound hint used when HasMax is true.
	// Effective max is max(dataDerivedMax, Max).
	Max float64
}

// VerticalBarChartValueLabels controls in-bar value labels for each segment.
type VerticalBarChartValueLabels struct {
	Show             bool
	MinSegmentHeight int
	Color            string
}

// VerticalBarChartLegendConfig controls legend placement and explicit legend items.
type VerticalBarChartLegendConfig struct {
	Placement VerticalBarChartLegendPlacementValue
	Items     []VerticalBarChartLegendItem
}

// VerticalBarChartBlock renders multi-series vertical columns with axis and legend controls.
type VerticalBarChartBlock struct {
	Title                 string
	Subtitle              string
	AxisLabels            []string
	Series                []VerticalBarChartSeries
	Height                int
	Normalize             bool
	HasColumnGap          bool
	ColumnGap             int
	HasOuterGap           bool
	OuterGap              int
	TransparentBackground bool
	Tone                  Tone
	InsetMode             InsetMode
	LegendPlacement       VerticalBarChartLegendPlacementValue
	Legend                []VerticalBarChartLegendItem
	Axis                  VerticalBarChartAxis
	ValueLabels           VerticalBarChartValueLabels
	ValueFormatter        VerticalBarChartValueFormatter
}

// VerticalBarChartSegmentView is the normalized render-time representation of one segment.
type VerticalBarChartSegmentView struct {
	Series          string
	Label           string
	Value           float64
	Color           string
	ValueLabelColor string
	Display         string
	SignedDisplay   string
	Height          int
}

// VerticalBarChartColumnView is the normalized render-time representation of one chart column.
// Positive and negative segments are split to support charts that cross a zero baseline,
// with separate padding/above-label metadata for positive stacks.
type VerticalBarChartColumnView struct {
	Label                   string
	PositiveSegments        []VerticalBarChartSegmentView
	NegativeSegments        []VerticalBarChartSegmentView
	PositiveTopPadding      int
	PositiveAboveLabel      string
	PositiveAboveLabelColor string
}

// VerticalBarChartTickView is a normalized Y-axis tick label used during template rendering.
type VerticalBarChartTickView struct {
	Label string
}

// VerticalBarChartTemplateData is the fully normalized rendering payload passed to templates.
// It includes computed geometry (heights, gaps, axis widths), preformatted labels, and
// pre-split positive/negative column segment data so templates remain mostly presentational.
type VerticalBarChartTemplateData struct {
	Title                 string
	Subtitle              string
	InsetMode             InsetMode
	Columns               []VerticalBarChartColumnView
	Legend                []VerticalBarChartLegendItem
	LegendPlacement       VerticalBarChartLegendPlacementValue
	Height                int
	ColumnGap             int
	OuterGap              int
	LegendStartOffset     int
	PositiveHeight        int
	NegativeHeight        int
	ShowBaseline          bool
	ShowCategoryLabels    bool
	ShowYTicks            bool
	DrawYAxisLine         bool
	YAxisWidth            int
	YAxisMaxLabel         string
	YAxisZeroLabel        string
	ShowValueLabels       bool
	ValueLabelMinHeight   int
	ValueLabelColor       string
	Ticks                 []VerticalBarChartTickView
	TransparentBackground bool
	Tone                  Tone
}

func (block VerticalBarChartBlock) Kind() theme.BlockKind {
	return theme.BlockKindVerticalBarChart
}

func (block VerticalBarChartBlock) TemplateData() any {
	normalized := block.normalized()
	return normalized
}

func (block VerticalBarChartBlock) RenderText(_ RenderContext) (string, error) {
	normalized := block.normalized()
	parts := make([]string, 0, len(normalized.Columns)+4)

	if strings.TrimSpace(normalized.Title) != "" {
		title := strings.TrimSpace(normalized.Title)
		parts = append(parts, title, strings.Repeat("-", min(48, max(8, len(title)))))
		if strings.TrimSpace(normalized.Subtitle) != "" {
			parts = append(parts, strings.TrimSpace(normalized.Subtitle))
		}
	}

	for _, column := range normalized.Columns {
		if column.Label == "" {
			continue
		}

		segmentParts := make([]string, 0, len(column.PositiveSegments)+len(column.NegativeSegments))
		for _, segment := range column.PositiveSegments {
			segmentParts = append(segmentParts, fmt.Sprintf("%s %s", segment.Label, segment.Display))
		}
		for _, segment := range column.NegativeSegments {
			segmentParts = append(segmentParts, fmt.Sprintf("%s %s", segment.Label, segment.SignedDisplay))
		}

		if len(segmentParts) == 0 {
			parts = append(parts, "- "+column.Label)
			continue
		}

		parts = append(parts, fmt.Sprintf("- %s: %s", column.Label, strings.Join(segmentParts, " · ")))
	}

	if normalized.LegendPlacement == VerticalBarChartLegendBottom && len(normalized.Legend) > 0 {
		labels := make([]string, 0, len(normalized.Legend))
		for _, item := range normalized.Legend {
			if item.Label == "" {
				continue
			}
			labels = append(labels, item.Label)
		}
		if len(labels) > 0 {
			parts = append(parts, "Legend: "+strings.Join(labels, " · "))
		}
	}

	return strings.Join(parts, "\n"), nil
}

func (block VerticalBarChartBlock) normalized() VerticalBarChartTemplateData {
	height := block.Height
	if height <= 0 {
		height = 160
	}
	if height < 80 {
		height = 80
	}
	if height > 360 {
		height = 360
	}

	axis := block.Axis
	if axis.LabelFormat != VerticalBarChartAxisLabelFormatPercent {
		axis.LabelFormat = VerticalBarChartAxisLabelFormatNumber
	}
	if !axis.HasShowCategoryLabels {
		axis.ShowCategoryLabels = true
	}
	if !axis.HasDrawYAxisLine {
		axis.DrawYAxisLine = axis.ShowYTicks
	}

	valueLabels := block.ValueLabels
	if valueLabels.MinSegmentHeight <= 0 {
		valueLabels.MinSegmentHeight = 16
	}
	if valueLabels.MinSegmentHeight < 10 {
		valueLabels.MinSegmentHeight = 10
	}
	valueLabels.Color = strings.TrimSpace(valueLabels.Color)

	valueFormatter := block.ValueFormatter
	if valueFormatter.MagnitudeSuffix != VerticalBarChartMagnitudeSuffixShort {
		valueFormatter.MagnitudeSuffix = VerticalBarChartMagnitudeSuffixNone
	}
	if valueFormatter.NegativeFormat != VerticalBarChartNegativeFormatParentheses {
		valueFormatter.NegativeFormat = VerticalBarChartNegativeFormatMinus
	}
	valueFormatter.Prefix = strings.TrimSpace(valueFormatter.Prefix)
	valueFormatter.Suffix = strings.TrimSpace(valueFormatter.Suffix)

	tone := normalizedChartTone(block.Tone)
	normalizedSeries := make([]VerticalBarChartSeries, 0, len(block.Series))
	for seriesIndex, series := range block.Series {
		key := strings.TrimSpace(series.Key)
		if key == "" {
			key = fmt.Sprintf("series_%d", seriesIndex+1)
		}
		label := strings.TrimSpace(series.Label)
		if label == "" {
			label = key
		}
		normalizedSeries = append(normalizedSeries, VerticalBarChartSeries{
			Key:             key,
			Label:           label,
			Color:           strings.TrimSpace(series.Color),
			ValueLabelColor: strings.TrimSpace(series.ValueLabelColor),
			Values:          append([]float64(nil), series.Values...),
		})
	}

	columnCount := len(block.AxisLabels)
	for _, series := range normalizedSeries {
		if len(series.Values) > columnCount {
			columnCount = len(series.Values)
		}
	}

	columns := make([]VerticalBarChartColumnView, 0, columnCount)
	inferredLegend := make([]VerticalBarChartLegendItem, 0, len(block.Series))
	legendSeen := map[string]struct{}{}
	maxPositiveTotal := 0.0
	maxNegativeAbsTotal := 0.0

	for columnIndex := 0; columnIndex < columnCount; columnIndex++ {
		label := ""
		if columnIndex < len(block.AxisLabels) {
			label = strings.TrimSpace(block.AxisLabels[columnIndex])
		}
		if label == "" {
			label = fmt.Sprintf("Column %d", columnIndex+1)
		}

		positiveSegments := make([]VerticalBarChartSegmentView, 0, len(normalizedSeries))
		negativeSegments := make([]VerticalBarChartSegmentView, 0, len(normalizedSeries))
		positiveTotal := 0.0
		negativeAbsTotal := 0.0

		for _, series := range normalizedSeries {
			if columnIndex >= len(series.Values) {
				continue
			}

			value := series.Values[columnIndex]
			if math.IsNaN(value) || math.IsInf(value, 0) || value == 0 {
				continue
			}

			display := formatVerticalBarChartValue(math.Abs(value), axis.LabelFormat, valueFormatter)
			signedDisplay := display
			if value < 0 {
				signedDisplay = formatVerticalBarChartSignedValue(display, valueFormatter)
			}

			segment := VerticalBarChartSegmentView{
				Series:          series.Key,
				Label:           series.Label,
				Value:           value,
				Color:           series.Color,
				ValueLabelColor: series.ValueLabelColor,
				Display:         display,
				SignedDisplay:   signedDisplay,
			}

			if value > 0 {
				positiveTotal += value
				positiveSegments = append(positiveSegments, segment)
			} else {
				negativeAbsTotal += math.Abs(value)
				negativeSegments = append(negativeSegments, segment)
			}

			if _, exists := legendSeen[series.Key]; !exists {
				legendSeen[series.Key] = struct{}{}
				inferredLegend = append(inferredLegend, VerticalBarChartLegendItem{Label: series.Label, Color: series.Color})
			}
		}

		if len(positiveSegments) == 0 && len(negativeSegments) == 0 {
			continue
		}

		if positiveTotal > maxPositiveTotal {
			maxPositiveTotal = positiveTotal
		}
		if negativeAbsTotal > maxNegativeAbsTotal {
			maxNegativeAbsTotal = negativeAbsTotal
		}

		columns = append(columns, VerticalBarChartColumnView{
			Label:            label,
			PositiveSegments: positiveSegments,
			NegativeSegments: negativeSegments,
		})
	}

	if len(columns) == 0 {
		columnGap := 0
		if block.HasColumnGap {
			columnGap = block.ColumnGap
			columnGap = normalizeVerticalBarChartGap(columnGap)
		}

		outerGap := columnGap
		if block.HasOuterGap {
			outerGap = block.OuterGap
			outerGap = normalizeVerticalBarChartGap(outerGap)
		}

		return VerticalBarChartTemplateData{
			Title:                 strings.TrimSpace(block.Title),
			Subtitle:              strings.TrimSpace(block.Subtitle),
			InsetMode:             block.InsetMode,
			Columns:               nil,
			Legend:                nil,
			LegendPlacement:       normalizedLegendPlacement(block.LegendPlacement),
			Height:                height,
			ColumnGap:             columnGap,
			OuterGap:              outerGap,
			LegendStartOffset:     outerGap,
			PositiveHeight:        height,
			NegativeHeight:        0,
			ShowBaseline:          false,
			ShowCategoryLabels:    axis.ShowCategoryLabels,
			ShowYTicks:            false,
			DrawYAxisLine:         axis.DrawYAxisLine,
			YAxisWidth:            0,
			YAxisMaxLabel:         "",
			YAxisZeroLabel:        "",
			ShowValueLabels:       valueLabels.Show,
			ValueLabelMinHeight:   valueLabels.MinSegmentHeight,
			ValueLabelColor:       valueLabels.Color,
			Ticks:                 nil,
			TransparentBackground: block.TransparentBackground,
			Tone:                  tone,
		}
	}

	columnGap := defaultVerticalBarChartColumnGap(len(columns))
	if block.HasColumnGap {
		columnGap = block.ColumnGap
		columnGap = normalizeVerticalBarChartGap(columnGap)
	}

	outerGap := columnGap
	if block.HasOuterGap {
		outerGap = block.OuterGap
		outerGap = normalizeVerticalBarChartGap(outerGap)
	}

	minValue := -maxNegativeAbsTotal
	maxValue := maxPositiveTotal
	if axis.HasMin {
		minValue = axis.Min
	}
	if axis.HasMax {
		// Max is intentionally treated as a floor for the upper range so explicit
		// axis configs can guarantee headroom without clipping higher data values.
		if axis.Max > maxValue {
			maxValue = axis.Max
		}
	}
	if maxValue <= minValue {
		maxValue = minValue + 1
	}

	positiveHeight, negativeHeight := resolveVerticalChartHeights(height, minValue, maxValue)

	for columnIndex := range columns {
		canNormalize := block.Normalize && maxNegativeAbsTotal <= 0
		columns[columnIndex].PositiveSegments = applyVerticalChartSegmentHeights(columns[columnIndex].PositiveSegments, maxValue, positiveHeight, canNormalize)
		columns[columnIndex].NegativeSegments = applyVerticalChartSegmentHeights(columns[columnIndex].NegativeSegments, math.Abs(minValue), negativeHeight, canNormalize)

		positiveUsedHeight := 0
		for _, segment := range columns[columnIndex].PositiveSegments {
			if segment.Height > 0 {
				positiveUsedHeight += segment.Height
			}
		}
		if positiveUsedHeight > positiveHeight {
			positiveUsedHeight = positiveHeight
		}
		columns[columnIndex].PositiveTopPadding = positiveHeight - positiveUsedHeight

		if !valueLabels.Show || columns[columnIndex].PositiveTopPadding < 12 {
			continue
		}

		topPositiveSegmentIndex := -1
		for segmentIndex, segment := range columns[columnIndex].PositiveSegments {
			if segment.Height > 0 {
				topPositiveSegmentIndex = segmentIndex
				break
			}
		}
		if topPositiveSegmentIndex >= 0 {
			topPositiveSegment := columns[columnIndex].PositiveSegments[topPositiveSegmentIndex]
			if topPositiveSegment.Height < valueLabels.MinSegmentHeight {
				columns[columnIndex].PositiveAboveLabel = topPositiveSegment.Display
				columns[columnIndex].PositiveAboveLabelColor = topPositiveSegment.ValueLabelColor
			}
			continue
		}

		topNegativeSegmentIndex := -1
		for segmentIndex, segment := range columns[columnIndex].NegativeSegments {
			if segment.Height > 0 {
				topNegativeSegmentIndex = segmentIndex
				break
			}
		}
		if topNegativeSegmentIndex < 0 {
			continue
		}

		topNegativeSegment := columns[columnIndex].NegativeSegments[topNegativeSegmentIndex]
		if topNegativeSegment.Height >= valueLabels.MinSegmentHeight {
			continue
		}

		columns[columnIndex].PositiveAboveLabel = topNegativeSegment.SignedDisplay
		columns[columnIndex].PositiveAboveLabelColor = topNegativeSegment.ValueLabelColor
	}

	legend := make([]VerticalBarChartLegendItem, 0, len(inferredLegend))
	if len(block.Legend) > 0 {
		for _, item := range block.Legend {
			label := strings.TrimSpace(item.Label)
			if label == "" {
				continue
			}
			legend = append(legend, VerticalBarChartLegendItem{Label: label, Color: strings.TrimSpace(item.Color)})
		}
	} else {
		legend = append(legend, inferredLegend...)
	}

	ticks := []VerticalBarChartTickView{}
	yAxisWidth := 0
	yAxisMaxLabel := ""
	yAxisZeroLabel := ""
	if axis.ShowYTicks {
		yAxisMaxLabel = formatVerticalBarChartValue(maxValue, axis.LabelFormat, valueFormatter)
		yAxisZeroLabel = formatVerticalBarChartValue(0, axis.LabelFormat, valueFormatter)
		yAxisWidth = resolveVerticalBarChartAxisWidth(yAxisMaxLabel, yAxisZeroLabel)
		ticks = buildVerticalBarChartAxisTicks(maxValue, axis.LabelFormat, valueFormatter)
	}

	showBaseline := axis.ShowBaseline || (minValue < 0 && maxValue > 0)
	if maxValue <= 0 || minValue >= 0 {
		showBaseline = axis.ShowBaseline
	}

	legendStartOffset := outerGap + yAxisWidth

	return VerticalBarChartTemplateData{
		Title:                 strings.TrimSpace(block.Title),
		Subtitle:              strings.TrimSpace(block.Subtitle),
		InsetMode:             block.InsetMode,
		Columns:               columns,
		Legend:                legend,
		LegendPlacement:       normalizedLegendPlacement(block.LegendPlacement),
		Height:                height,
		ColumnGap:             columnGap,
		OuterGap:              outerGap,
		LegendStartOffset:     legendStartOffset,
		PositiveHeight:        positiveHeight,
		NegativeHeight:        negativeHeight,
		ShowBaseline:          showBaseline,
		ShowCategoryLabels:    axis.ShowCategoryLabels,
		ShowYTicks:            axis.ShowYTicks,
		DrawYAxisLine:         axis.DrawYAxisLine,
		YAxisWidth:            yAxisWidth,
		YAxisMaxLabel:         yAxisMaxLabel,
		YAxisZeroLabel:        yAxisZeroLabel,
		ShowValueLabels:       valueLabels.Show,
		ValueLabelMinHeight:   valueLabels.MinSegmentHeight,
		ValueLabelColor:       valueLabels.Color,
		Ticks:                 ticks,
		TransparentBackground: block.TransparentBackground,
		Tone:                  tone,
	}
}

func normalizeVerticalBarChartGap(value int) int {
	if value < 0 {
		return 0
	}
	if value > 28 {
		return 28
	}

	return value
}

func normalizedLegendPlacement(value VerticalBarChartLegendPlacementValue) VerticalBarChartLegendPlacementValue {
	if value == VerticalBarChartLegendBottom {
		return VerticalBarChartLegendBottom
	}

	return VerticalBarChartLegendNone
}

func resolveVerticalChartHeights(totalHeight int, minValue, maxValue float64) (int, int) {
	if minValue < 0 && maxValue > 0 {
		rangeSize := maxValue - minValue
		positiveHeight := int(math.Round((maxValue / rangeSize) * float64(totalHeight)))
		if positiveHeight < 1 {
			positiveHeight = 1
		}
		if positiveHeight > totalHeight-1 {
			positiveHeight = totalHeight - 1
		}
		negativeHeight := totalHeight - positiveHeight
		return positiveHeight, negativeHeight
	}

	if maxValue <= 0 {
		return 0, totalHeight
	}

	return totalHeight, 0
}

func applyVerticalChartSegmentHeights(segments []VerticalBarChartSegmentView, scale float64, regionHeight int, normalize bool) []VerticalBarChartSegmentView {
	if len(segments) == 0 || scale <= 0 || regionHeight <= 0 {
		for index := range segments {
			segments[index].Height = 0
		}
		return segments
	}

	if !normalize {
		for index := range segments {
			value := math.Abs(segments[index].Value)
			height := int(math.Round((value / scale) * float64(regionHeight)))
			if height < 0 {
				height = 0
			}
			if height > regionHeight {
				height = regionHeight
			}
			segments[index].Height = height
		}

		return segments
	}

	remaining := regionHeight
	for index := range segments {
		value := math.Abs(segments[index].Value)
		height := int(math.Round((value / scale) * float64(regionHeight)))
		if index == len(segments)-1 {
			height = remaining
		}
		if height < 0 {
			height = 0
		}
		if height > remaining {
			height = remaining
		}
		segments[index].Height = height
		remaining -= height
	}

	return segments
}

func buildVerticalBarChartAxisTicks(maxValue float64, format VerticalBarChartAxisLabelFormatValue, valueFormatter VerticalBarChartValueFormatter) []VerticalBarChartTickView {
	maxLabel := formatVerticalBarChartValue(maxValue, format, valueFormatter)
	zeroLabel := formatVerticalBarChartValue(0, format, valueFormatter)
	if maxLabel == zeroLabel {
		return []VerticalBarChartTickView{{Label: maxLabel}}
	}

	return []VerticalBarChartTickView{{Label: maxLabel}, {Label: zeroLabel}}
}

func formatVerticalBarChartValue(value float64, format VerticalBarChartAxisLabelFormatValue, valueFormatter VerticalBarChartValueFormatter) string {
	rounded := value
	if math.Abs(value-math.Round(value)) < 0.001 {
		rounded = math.Round(value)
	}
	formatted := strconv.FormatFloat(rounded, 'f', -1, 64)
	if format == VerticalBarChartAxisLabelFormatPercent {
		return formatted + "%"
	}
	if valueFormatter.MagnitudeSuffix == VerticalBarChartMagnitudeSuffixShort {
		formatted = formatVerticalBarChartCompactNumber(rounded)
	}
	if valueFormatter.Prefix != "" {
		formatted = valueFormatter.Prefix + formatted
	}
	if valueFormatter.Suffix != "" {
		formatted += valueFormatter.Suffix
	}

	return formatted
}

func formatVerticalBarChartCompactNumber(value float64) string {
	absolute := math.Abs(value)
	suffix := ""
	scaled := absolute

	switch {
	case absolute >= 1_000_000_000:
		scaled = absolute / 1_000_000_000
		suffix = "B"
	case absolute >= 1_000_000:
		scaled = absolute / 1_000_000
		suffix = "M"
	case absolute >= 1_000:
		scaled = absolute / 1_000
		suffix = "K"
	}

	formatted := strconv.FormatFloat(scaled, 'f', 1, 64)
	formatted = strings.TrimSuffix(formatted, ".0")
	if value < 0 {
		formatted = "-" + formatted
	}

	return formatted + suffix
}

func formatVerticalBarChartSignedValue(display string, valueFormatter VerticalBarChartValueFormatter) string {
	if valueFormatter.NegativeFormat == VerticalBarChartNegativeFormatParentheses {
		return "(" + display + ")"
	}

	return "-" + display
}

func defaultVerticalBarChartColumnGap(columnCount int) int {
	if columnCount <= 0 {
		return 0
	}
	if columnCount <= 4 {
		return 8
	}
	if columnCount <= 8 {
		return 6
	}
	if columnCount <= 12 {
		return 4
	}
	if columnCount <= 18 {
		return 3
	}
	if columnCount <= 30 {
		return 2
	}

	return 1
}

func resolveVerticalBarChartAxisWidth(maxLabel, zeroLabel string) int {
	maxLen := len(strings.TrimSpace(maxLabel))
	zeroLen := len(strings.TrimSpace(zeroLabel))
	if zeroLen > maxLen {
		maxLen = zeroLen
	}

	if maxLen <= 0 {
		return 36
	}

	width := 14 + maxLen*6
	if width < 36 {
		return 36
	}
	if width > 52 {
		return 52
	}

	return width
}

func (block VerticalBarChartBlock) LayoutSpec() LayoutSpec {
	return normalizedLayoutSpec(LayoutSpec{InsetMode: block.InsetMode})
}
