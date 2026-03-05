package myrtle

import (
	"fmt"
	"strings"

	"github.com/gzuidhof/myrtle/theme"
)

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
