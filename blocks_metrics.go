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

type TimelineBlock struct {
	Header          string
	AggregateHeader string
	HasCurrentIndex bool
	CurrentIndex    int
	Items           []TimelineItem
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

func (block TimelineBlock) RenderMarkdown(_ RenderContext) (string, error) {
	parts := make([]string, 0, len(block.Items)+2)
	if strings.TrimSpace(block.Header) != "" {
		parts = append(parts, "### "+strings.TrimSpace(block.Header))
	}
	if strings.TrimSpace(block.AggregateHeader) != "" {
		parts = append(parts, "_"+strings.TrimSpace(block.AggregateHeader)+"_")
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
			line += "**" + strings.TrimSpace(item.Time) + "** — "
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

func (block StatsRowBlock) RenderMarkdown(_ RenderContext) (string, error) {
	parts := make([]string, 0, len(block.Stats)+1)
	if strings.TrimSpace(block.Header) != "" {
		parts = append(parts, "### "+strings.TrimSpace(block.Header))
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
			line += "**" + value + "**"
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

type BadgeTone string

const (
	BadgeToneInfo    BadgeTone = "info"
	BadgeToneSuccess BadgeTone = "success"
	BadgeToneWarning BadgeTone = "warning"
	BadgeToneError   BadgeTone = "error"
)

type BadgeBlock struct {
	Tone BadgeTone
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

func (block BadgeBlock) RenderMarkdown(_ RenderContext) (string, error) {
	text := strings.TrimSpace(block.Text)
	if text == "" {
		return "", nil
	}
	return fmt.Sprintf("**[%s]** %s", strings.ToUpper(string(normalizedBadgeTone(block.Tone))), text), nil
}

type SummaryCardBlock struct {
	Title  string
	Body   string
	Footer string
}

func (block SummaryCardBlock) Kind() theme.BlockKind {
	return theme.BlockKindSummaryCard
}

func (block SummaryCardBlock) TemplateData() any {
	return block
}

func (block SummaryCardBlock) RenderMarkdown(_ RenderContext) (string, error) {
	parts := make([]string, 0, 3)
	if strings.TrimSpace(block.Title) != "" {
		parts = append(parts, "### "+strings.TrimSpace(block.Title))
	}
	if strings.TrimSpace(block.Body) != "" {
		parts = append(parts, strings.TrimSpace(block.Body))
	}
	if strings.TrimSpace(block.Footer) != "" {
		parts = append(parts, "_"+strings.TrimSpace(block.Footer)+"_")
	}
	return strings.Join(parts, "\n\n"), nil
}

type AttachmentBlock struct {
	Filename string
	Meta     string
	URL      string
	CTA      string
}

func (block AttachmentBlock) Kind() theme.BlockKind {
	return theme.BlockKindAttachment
}

func (block AttachmentBlock) TemplateData() any {
	return block
}

func (block AttachmentBlock) RenderMarkdown(_ RenderContext) (string, error) {
	filename := strings.TrimSpace(block.Filename)
	url := strings.TrimSpace(block.URL)
	if filename == "" || url == "" {
		return "", nil
	}

	line := fmt.Sprintf("[%s](%s)", filename, url)
	if strings.TrimSpace(block.Meta) != "" {
		line += " — " + strings.TrimSpace(block.Meta)
	}
	if strings.TrimSpace(block.CTA) != "" {
		line += fmt.Sprintf(" ([%s](%s))", strings.TrimSpace(block.CTA), url)
	}

	return line, nil
}

func normalizedBadgeTone(value BadgeTone) BadgeTone {
	switch value {
	case BadgeToneSuccess, BadgeToneWarning, BadgeToneError:
		return value
	default:
		return BadgeToneInfo
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
