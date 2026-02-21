package myrtle

import (
	"fmt"
	"strings"

	"github.com/gzuidhof/myrtle/theme"
)

type ButtonBlock struct {
	Label     string
	URL       string
	Variant   ButtonVariant
	Alignment ButtonAlignment
	FullWidth bool
}

type ButtonGroupButton struct {
	Label   string
	URL     string
	Variant ButtonVariant
}

type ButtonGroupBlock struct {
	Buttons   []ButtonGroupButton
	Alignment ButtonAlignment
	Joined    bool
}

type (
	ButtonVariant   string
	ButtonAlignment string
)

const (
	ButtonVariantPrimary   ButtonVariant = "primary"
	ButtonVariantSecondary ButtonVariant = "secondary"
	ButtonVariantGhost     ButtonVariant = "ghost"

	ButtonAlignmentLeft   ButtonAlignment = "left"
	ButtonAlignmentCenter ButtonAlignment = "center"
	ButtonAlignmentRight  ButtonAlignment = "right"
)

func (block ButtonBlock) Kind() theme.BlockKind {
	return theme.BlockKindButton
}

func (block ButtonBlock) TemplateData() any {
	normalized := block
	normalized.Variant = normalizedButtonVariant(block.Variant)
	normalized.Alignment = normalizedButtonAlignment(block.Alignment)
	return normalized
}

func (block ButtonBlock) RenderMarkdown(_ RenderContext) (string, error) {
	return fmt.Sprintf("[%s](%s)", block.Label, block.URL), nil
}

func normalizedButtonVariant(value ButtonVariant) ButtonVariant {
	switch value {
	case ButtonVariantSecondary, ButtonVariantGhost:
		return value
	default:
		return ButtonVariantPrimary
	}
}

func normalizedButtonAlignment(value ButtonAlignment) ButtonAlignment {
	switch value {
	case ButtonAlignmentCenter, ButtonAlignmentRight:
		return value
	default:
		return ButtonAlignmentLeft
	}
}

func (block ButtonGroupBlock) Kind() theme.BlockKind {
	return theme.BlockKindButtonGroup
}

func (block ButtonGroupBlock) TemplateData() any {
	normalized := ButtonGroupBlock{
		Alignment: normalizedButtonAlignment(block.Alignment),
		Joined:    block.Joined,
		Buttons:   make([]ButtonGroupButton, 0, len(block.Buttons)),
	}

	for _, button := range block.Buttons {
		label := strings.TrimSpace(button.Label)
		url := strings.TrimSpace(button.URL)
		if label == "" || url == "" {
			continue
		}

		normalized.Buttons = append(normalized.Buttons, ButtonGroupButton{
			Label:   label,
			URL:     url,
			Variant: normalizedButtonVariant(button.Variant),
		})
	}

	return normalized
}

func (block ButtonGroupBlock) RenderMarkdown(_ RenderContext) (string, error) {
	normalized := block.TemplateData().(ButtonGroupBlock)
	parts := make([]string, 0, len(normalized.Buttons))
	for _, button := range normalized.Buttons {
		parts = append(parts, fmt.Sprintf("[%s](%s)", button.Label, button.URL))
	}

	return strings.Join(parts, " · "), nil
}
