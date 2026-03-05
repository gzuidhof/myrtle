package myrtle

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gzuidhof/myrtle/theme"
)

// SparklineBlock renders a compact inline trend chart with summary values.
type SparklineBlock struct {
	Header        string
	Label         string
	Value         string
	Delta         string
	DeltaSemantic StatDeltaSemantic
	Tone          Tone
	Points        []int
	InsetMode     InsetMode
}

func (block SparklineBlock) Kind() theme.BlockKind {
	return theme.BlockKindSparkline
}

func (block SparklineBlock) TemplateData() any {
	normalized := block
	normalized.Points = normalizedIntPoints(block.Points)
	normalized.DeltaSemantic = normalizedStatDeltaSemantic(block.DeltaSemantic)
	normalized.Tone = normalizedChartTone(block.Tone)
	return normalized
}

func (block SparklineBlock) RenderText(_ RenderContext) (string, error) {
	normalized := block.TemplateData().(SparklineBlock)
	parts := make([]string, 0, 3)
	if strings.TrimSpace(normalized.Header) != "" {
		header := strings.TrimSpace(normalized.Header)
		parts = append(parts, header, strings.Repeat("-", min(48, max(8, len(header)))))
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
	Color   string
}

type StackedBarRow struct {
	Label    string
	Segments []StackedBarSegment
}

// StackedBarBlock renders stacked proportional bars for part-to-whole data.
type StackedBarBlock struct {
	Header     string
	TotalLabel string
	TotalValue string
	Rows       []StackedBarRow
	InsetMode  InsetMode
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
				Color:   strings.TrimSpace(segment.Color),
			})
		}

		normalized.Rows = append(normalized.Rows, StackedBarRow{
			Label:    strings.TrimSpace(row.Label),
			Segments: segments,
		})
	}

	return normalized
}

func normalizedChartTone(value Tone) Tone {
	switch value {
	case ToneMuted, ToneInfo, ToneSuccess, ToneWarning, ToneDanger, ToneDark:
		return value
	default:
		return TonePrimary
	}
}

func (block StackedBarBlock) RenderText(_ RenderContext) (string, error) {
	normalized := block.TemplateData().(StackedBarBlock)
	parts := make([]string, 0, len(normalized.Rows)+2)
	if strings.TrimSpace(normalized.Header) != "" {
		header := strings.TrimSpace(normalized.Header)
		parts = append(parts, header, strings.Repeat("-", min(48, max(8, len(header)))))
	}
	if strings.TrimSpace(normalized.TotalLabel) != "" || strings.TrimSpace(normalized.TotalValue) != "" {
		parts = append(parts, fmt.Sprintf("%s: %s", strings.TrimSpace(normalized.TotalLabel), strings.TrimSpace(normalized.TotalValue)))
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
		parts = append(parts, fmt.Sprintf("- %s: %s", label, strings.Join(segmentParts, " · ")))
	}

	return strings.Join(parts, "\n"), nil
}

type ProgressItem struct {
	Label   string
	Percent int
	Value   string
	Color   string
}

// ProgressBlock renders one or more progress indicators.
type ProgressBlock struct {
	Header    string
	Items     []ProgressItem
	InsetMode InsetMode
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
			Color:   strings.TrimSpace(item.Color),
		})
	}

	return normalized
}

func (block ProgressBlock) RenderText(_ RenderContext) (string, error) {
	normalized := block.TemplateData().(ProgressBlock)
	parts := make([]string, 0, len(normalized.Items)+1)
	if strings.TrimSpace(normalized.Header) != "" {
		header := strings.TrimSpace(normalized.Header)
		parts = append(parts, header, strings.Repeat("-", min(48, max(8, len(header)))))
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
		parts = append(parts, fmt.Sprintf("- %s: %s %s%s", item.Label, item.Value, strings.Repeat("#", filled), strings.Repeat(".", empty)))
	}

	return strings.Join(parts, "\n"), nil
}

type DistributionBucket struct {
	Label        string
	Count        int
	WidthPercent int
	Color        string
}

// DistributionBlock renders bucketed value distributions.
type DistributionBlock struct {
	Header             string
	Buckets            []DistributionBucket
	CountColumnWidthCh int
	InsetMode          InsetMode
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
			Color: strings.TrimSpace(bucket.Color),
		})
	}

	if maxCount <= 0 {
		maxCount = 1
	}

	countWidthCh := len(strconv.Itoa(maxCount))
	if countWidthCh < 2 {
		countWidthCh = 2
	}
	normalized.CountColumnWidthCh = countWidthCh

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

func (block DistributionBlock) RenderText(_ RenderContext) (string, error) {
	normalized := block.TemplateData().(DistributionBlock)
	parts := make([]string, 0, len(normalized.Buckets)+1)
	if strings.TrimSpace(normalized.Header) != "" {
		header := strings.TrimSpace(normalized.Header)
		parts = append(parts, header, strings.Repeat("-", min(48, max(8, len(header)))))
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
		parts = append(parts, fmt.Sprintf("- %s %s%s (%d)", bucket.Label, strings.Repeat("#", filled), strings.Repeat(".", empty), bucket.Count))
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

	glyphs := []rune{'.', ':', '-', '=', '+', '*', '#', '@'}
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

func (block SparklineBlock) LayoutSpec() LayoutSpec {
	return normalizedLayoutSpec(LayoutSpec{InsetMode: block.InsetMode})
}

func (block StackedBarBlock) LayoutSpec() LayoutSpec {
	return normalizedLayoutSpec(LayoutSpec{InsetMode: block.InsetMode})
}

func (block ProgressBlock) LayoutSpec() LayoutSpec {
	return normalizedLayoutSpec(LayoutSpec{InsetMode: block.InsetMode})
}

func (block DistributionBlock) LayoutSpec() LayoutSpec {
	return normalizedLayoutSpec(LayoutSpec{InsetMode: block.InsetMode})
}

type KeyValuePair struct {
	Key   string
	Value string
}

// KeyValueBlock renders labeled key-value pairs.
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
	header := strings.TrimSpace(block.Header)

	parts := make([]string, 0, len(block.Pairs)+1)
	if header != "" {
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

// HorizontalBarChartBlock renders a horizontal category comparison chart.
type HorizontalBarChartBlock struct {
	Header                string
	Items                 []HorizontalBarChartItem
	Thickness             int
	ShowLabelsInsideBars  bool
	TransparentBackground bool
	Tone                  Tone
	InsetMode             InsetMode
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
	header := strings.TrimSpace(block.Header)

	parts := make([]string, 0, len(block.Items)+1)
	if header != "" {
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
	minThickness := 8
	if block.ShowLabelsInsideBars {
		minThickness = 18
	}

	if block.Thickness <= 0 {
		return minThickness
	}
	if block.Thickness > 24 {
		return 24
	}
	if block.Thickness < minThickness {
		return minThickness
	}
	return block.Thickness
}

type TimelineItem struct {
	Time   string
	Title  string
	Detail string
}

// TimelineBlock renders chronological milestones or status updates.
type TimelineBlock struct {
	Header          string
	AggregateHeader string
	HasCurrentIndex bool
	CurrentIndex    int
	Items           []TimelineItem
	InsetMode       InsetMode
}

func (block TimelineBlock) Kind() theme.BlockKind {
	return theme.BlockKindTimeline
}

func (block TimelineBlock) TemplateData() any {
	normalized := block
	if !normalized.HasCurrentIndex || normalized.CurrentIndex < 0 || normalized.CurrentIndex >= len(normalized.Items) {
		normalized.CurrentIndex = -1
	}

	return normalized
}

func (block TimelineBlock) RenderText(_ RenderContext) (string, error) {
	parts := make([]string, 0, len(block.Items)+2)
	if strings.TrimSpace(block.Header) != "" {
		header := strings.TrimSpace(block.Header)
		parts = append(parts, header, strings.Repeat("-", min(48, max(8, len(header)))))
	}
	if strings.TrimSpace(block.AggregateHeader) != "" {
		parts = append(parts, strings.TrimSpace(block.AggregateHeader))
	}

	currentIndex := block.CurrentIndex
	if !block.HasCurrentIndex || currentIndex < 0 || currentIndex >= len(block.Items) {
		currentIndex = -1
	}

	for index, item := range block.Items {
		title := strings.TrimSpace(item.Title)
		if title == "" {
			continue
		}

		line := "- "
		if index == currentIndex {
			line = "- 👉 "
		}
		if strings.TrimSpace(item.Time) != "" {
			line += strings.TrimSpace(item.Time) + " - "
		}
		line += title
		if strings.TrimSpace(item.Detail) != "" {
			line += ": " + strings.TrimSpace(item.Detail)
		}
		parts = append(parts, line)
	}

	return strings.Join(parts, "\n"), nil
}

type StatItem struct {
	Label         string
	Value         string
	Delta         string
	DeltaSemantic StatDeltaSemantic
}

type StatDeltaSemantic string

const (
	StatDeltaSemanticNone     StatDeltaSemantic = "none"
	StatDeltaSemanticPositive StatDeltaSemantic = "positive"
	StatDeltaSemanticNegative StatDeltaSemantic = "negative"
)

// StatsRowBlock renders a row of compact KPI/stat entries.
type StatsRowBlock struct {
	Header string
	Stats  []StatItem
}

func (block StatsRowBlock) Kind() theme.BlockKind {
	return theme.BlockKindStatsRow
}

func (block StatsRowBlock) TemplateData() any {
	normalized := block
	normalized.Stats = make([]StatItem, 0, len(block.Stats))
	for _, stat := range block.Stats {
		stat.DeltaSemantic = normalizedStatDeltaSemantic(stat.DeltaSemantic)
		normalized.Stats = append(normalized.Stats, stat)
	}
	return normalized
}

func (block StatsRowBlock) RenderText(_ RenderContext) (string, error) {
	parts := make([]string, 0, len(block.Stats)+1)
	if strings.TrimSpace(block.Header) != "" {
		header := strings.TrimSpace(block.Header)
		parts = append(parts, header, strings.Repeat("-", min(48, max(8, len(header)))))
	}

	for _, stat := range block.Stats {
		stat.DeltaSemantic = normalizedStatDeltaSemantic(stat.DeltaSemantic)
		label := strings.TrimSpace(stat.Label)
		value := strings.TrimSpace(stat.Value)
		if label == "" && value == "" {
			continue
		}

		line := "- "
		if value != "" {
			line += value
		}
		if label != "" {
			line += " " + label
		}
		if strings.TrimSpace(stat.Delta) != "" {
			line += " (" + strings.TrimSpace(stat.Delta) + ")"
		}
		parts = append(parts, strings.TrimSpace(line))
	}

	return strings.Join(parts, "\n"), nil
}

// BadgeBlock renders a short status label with semantic tone.
type BadgeBlock struct {
	Tone Tone
	Text string
}

func (block BadgeBlock) Kind() theme.BlockKind {
	return theme.BlockKindBadge
}

func (block BadgeBlock) TemplateData() any {
	normalized := block
	normalized.Tone = normalizedBadgeTone(block.Tone)
	return normalized
}

func (block BadgeBlock) RenderText(_ RenderContext) (string, error) {
	text := strings.TrimSpace(block.Text)
	if text == "" {
		return "", nil
	}
	return fmt.Sprintf("[%s] %s", strings.ToUpper(string(normalizedBadgeTone(block.Tone))), text), nil
}

// SummaryCardBlock renders a concise title/body/footer summary card.
type SummaryCardBlock struct {
	Title     string
	Body      string
	Footer    string
	Tone      Tone
	InsetMode InsetMode
}

func (block SummaryCardBlock) Kind() theme.BlockKind {
	return theme.BlockKindSummaryCard
}

func (block SummaryCardBlock) TemplateData() any {
	normalized := block
	normalized.Tone = normalizedSummaryCardTone(block.Tone)

	return normalized
}

func (block SummaryCardBlock) RenderText(_ RenderContext) (string, error) {
	parts := make([]string, 0, 3)
	if strings.TrimSpace(block.Title) != "" {
		title := strings.TrimSpace(block.Title)
		parts = append(parts, "[ "+title+" ]", strings.Repeat("-", min(48, max(8, len(title)+4))))
	}
	if strings.TrimSpace(block.Body) != "" {
		parts = append(parts, strings.TrimSpace(block.Body))
	}
	if strings.TrimSpace(block.Footer) != "" {
		parts = append(parts, strings.TrimSpace(block.Footer))
	}
	return strings.Join(parts, "\n\n"), nil
}

// AttachmentBlock renders file attachment metadata with a CTA link.
type AttachmentBlock struct {
	Filename  string
	Meta      string
	URL       string
	CTA       string
	InsetMode InsetMode
}

func (block AttachmentBlock) Kind() theme.BlockKind {
	return theme.BlockKindAttachment
}

func (block AttachmentBlock) TemplateData() any {
	return block
}

func (block AttachmentBlock) RenderText(_ RenderContext) (string, error) {
	filename := strings.TrimSpace(block.Filename)
	url := strings.TrimSpace(block.URL)
	if filename == "" || url == "" {
		return "", nil
	}

	line := fmt.Sprintf("%s (%s)", filename, url)
	if strings.TrimSpace(block.Meta) != "" {
		line += " — " + strings.TrimSpace(block.Meta)
	}
	if strings.TrimSpace(block.CTA) != "" {
		line += fmt.Sprintf(" (%s: %s)", strings.TrimSpace(block.CTA), url)
	}

	return line, nil
}

func normalizedBadgeTone(value Tone) Tone {
	switch value {
	case ToneSuccess, ToneWarning, ToneDanger, ToneDark:
		return value
	default:
		return ToneInfo
	}
}

func normalizedStatDeltaSemantic(value StatDeltaSemantic) StatDeltaSemantic {
	switch value {
	case StatDeltaSemanticPositive, StatDeltaSemanticNegative:
		return value
	default:
		return StatDeltaSemanticNone
	}
}

func normalizedSummaryCardTone(value Tone) Tone {
	return normalizedTone(value)
}

func (block KeyValueBlock) LayoutSpec() LayoutSpec { return defaultLayoutSpec() }

func (block HorizontalBarChartBlock) LayoutSpec() LayoutSpec {
	return normalizedLayoutSpec(LayoutSpec{InsetMode: block.InsetMode})
}

func (block TimelineBlock) LayoutSpec() LayoutSpec {
	return normalizedLayoutSpec(LayoutSpec{InsetMode: block.InsetMode})
}

func (block StatsRowBlock) LayoutSpec() LayoutSpec { return defaultLayoutSpec() }

func (block BadgeBlock) LayoutSpec() LayoutSpec { return defaultLayoutSpec() }

func (block SummaryCardBlock) LayoutSpec() LayoutSpec {
	return normalizedLayoutSpec(LayoutSpec{InsetMode: block.InsetMode})
}

func (block AttachmentBlock) LayoutSpec() LayoutSpec {
	return normalizedLayoutSpec(LayoutSpec{InsetMode: block.InsetMode})
}
