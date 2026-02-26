package myrtle

import (
	"strings"

	"github.com/gzuidhof/myrtle/theme"
)

type LegalBlock struct {
	CompanyName    string
	Address        string
	ManageURL      string
	UnsubscribeURL string
}

func (block LegalBlock) Kind() theme.BlockKind {
	return theme.BlockKindLegal
}

func (block LegalBlock) TemplateData() any {
	return block
}

func (block LegalBlock) RenderText(_ RenderContext) (string, error) {
	parts := make([]string, 0, 4)
	if strings.TrimSpace(block.CompanyName) != "" {
		parts = append(parts, strings.TrimSpace(block.CompanyName))
	}
	if strings.TrimSpace(block.Address) != "" {
		parts = append(parts, strings.TrimSpace(block.Address))
	}

	links := make([]string, 0, 2)
	if strings.TrimSpace(block.ManageURL) != "" {
		links = append(links, "Manage preferences ("+strings.TrimSpace(block.ManageURL)+")")
	}
	if strings.TrimSpace(block.UnsubscribeURL) != "" {
		links = append(links, "Unsubscribe ("+strings.TrimSpace(block.UnsubscribeURL)+")")
	}
	if len(links) > 0 {
		parts = append(parts, strings.Join(links, " · "))
	}

	return strings.Join(parts, "\n\n"), nil
}
