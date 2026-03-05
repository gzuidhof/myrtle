package myrtle

import (
	"strings"

	"github.com/gzuidhof/myrtle/theme"
)

// HeroBlock renders a high-impact marketing hero section.
type HeroBlock struct {
	Eyebrow   string
	Title     string
	Body      string
	CTALabel  string
	CTAURL    string
	ImageURL  string
	ImageAlt  string
	Tone      Tone
	InsetMode InsetMode
}

func (block HeroBlock) Kind() theme.BlockKind {
	return theme.BlockKindHero
}

func (block HeroBlock) TemplateData() any {
	normalized := block
	normalized.Tone = normalizedTone(block.Tone)

	return normalized
}

func (block HeroBlock) RenderText(_ RenderContext) (string, error) {
	eyebrow := strings.TrimSpace(block.Eyebrow)
	title := strings.TrimSpace(block.Title)
	body := strings.TrimSpace(block.Body)
	ctaLabel := strings.TrimSpace(block.CTALabel)
	ctaURL := strings.TrimSpace(block.CTAURL)

	parts := make([]string, 0, 4)
	if eyebrow != "" {
		parts = append(parts, eyebrow)
	}
	if title != "" {
		parts = append(parts, title, strings.Repeat("-", min(48, max(8, len(title)))))
	}
	if body != "" {
		parts = append(parts, body)
	}
	if ctaLabel != "" && ctaURL != "" {
		parts = append(parts, ctaLabel+" ("+ctaURL+")")
	}

	return strings.Join(parts, "\n\n"), nil
}

// FooterLink is one label/URL pair rendered by FooterLinksBlock.
type FooterLink struct {
	Label string
	URL   string
}

// FooterLinksBlock renders navigational links and an optional footer note.
type FooterLinksBlock struct {
	Links []FooterLink
	Note  string
}

func (block FooterLinksBlock) Kind() theme.BlockKind {
	return theme.BlockKindFooterLinks
}

func (block FooterLinksBlock) TemplateData() any {
	return block
}

func (block FooterLinksBlock) RenderText(_ RenderContext) (string, error) {
	links := make([]string, 0, len(block.Links))
	note := strings.TrimSpace(block.Note)
	for _, link := range block.Links {
		label := strings.TrimSpace(link.Label)
		url := strings.TrimSpace(link.URL)
		if label == "" || url == "" {
			continue
		}

		links = append(links, label+" ("+url+")")
	}

	parts := make([]string, 0, 2)
	if len(links) > 0 {
		parts = append(parts, strings.Join(links, " · "))
	}
	if note != "" {
		parts = append(parts, note)
	}

	return strings.Join(parts, "\n\n"), nil
}

func (block HeroBlock) LayoutSpec() LayoutSpec {
	return normalizedLayoutSpec(LayoutSpec{InsetMode: block.InsetMode})
}

func (block FooterLinksBlock) LayoutSpec() LayoutSpec { return defaultLayoutSpec() }

// PriceLine is one line item rendered in a PriceSummaryBlock.
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
