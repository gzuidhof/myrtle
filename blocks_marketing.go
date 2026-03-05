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
