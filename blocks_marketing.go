package myrtle

import (
	"strings"

	"github.com/gzuidhof/myrtle/theme"
)

type HeroBlock struct {
	Eyebrow  string
	Title    string
	Body     string
	CTALabel string
	CTAURL   string
	ImageURL string
	ImageAlt string
}

func (block HeroBlock) Kind() theme.BlockKind {
	return theme.BlockKindHero
}

func (block HeroBlock) TemplateData() any {
	return block
}

func (block HeroBlock) RenderMarkdown(_ RenderContext) (string, error) {
	parts := make([]string, 0, 4)
	if strings.TrimSpace(block.Eyebrow) != "" {
		parts = append(parts, "_"+strings.TrimSpace(block.Eyebrow)+"_")
	}
	if strings.TrimSpace(block.Title) != "" {
		parts = append(parts, "## "+strings.TrimSpace(block.Title))
	}
	if strings.TrimSpace(block.Body) != "" {
		parts = append(parts, strings.TrimSpace(block.Body))
	}
	if strings.TrimSpace(block.CTALabel) != "" && strings.TrimSpace(block.CTAURL) != "" {
		parts = append(parts, "["+strings.TrimSpace(block.CTALabel)+"]("+strings.TrimSpace(block.CTAURL)+")")
	}

	return strings.Join(parts, "\n\n"), nil
}

type FooterLink struct {
	Label string
	URL   string
}

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

func (block FooterLinksBlock) RenderMarkdown(_ RenderContext) (string, error) {
	links := make([]string, 0, len(block.Links))
	for _, link := range block.Links {
		label := strings.TrimSpace(link.Label)
		url := strings.TrimSpace(link.URL)
		if label == "" || url == "" {
			continue
		}

		links = append(links, "["+label+"]("+url+")")
	}

	parts := make([]string, 0, 2)
	if len(links) > 0 {
		parts = append(parts, strings.Join(links, " · "))
	}
	if strings.TrimSpace(block.Note) != "" {
		parts = append(parts, strings.TrimSpace(block.Note))
	}

	return strings.Join(parts, "\n\n"), nil
}
