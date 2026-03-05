package myrtle

import (
	"fmt"
	"strings"

	"github.com/gzuidhof/myrtle/theme"
)

// ButtonBlock renders a single call-to-action button.
type ButtonBlock struct {
	Label     string
	URL       string
	Tone      Tone
	Style     ButtonStyleValue
	Alignment ButtonAlignmentValue
	Size      ButtonSizeValue
	NoWrap    bool
	FullWidth bool
}

type ButtonGroupButton struct {
	Label string
	URL   string
	Tone  Tone
	Style ButtonStyleValue
}

// ButtonGroupBlock renders multiple related buttons as one responsive group.
type ButtonGroupBlock struct {
	Buttons           []ButtonGroupButton
	Alignment         ButtonAlignmentValue
	Joined            bool
	Gap               int
	StackOnMobile     bool
	FullWidthOnMobile bool
}

type (
	ButtonStyleValue     string
	ButtonAlignmentValue string
	ButtonSizeValue      string
)

const (
	ButtonStyleFilled  ButtonStyleValue = "filled"
	ButtonStyleOutline ButtonStyleValue = "outline"
	ButtonStyleGhost   ButtonStyleValue = "ghost"

	ButtonAlignmentStart  ButtonAlignmentValue = "start"
	ButtonAlignmentCenter ButtonAlignmentValue = "center"
	ButtonAlignmentEnd    ButtonAlignmentValue = "end"

	ButtonSizeSmall  ButtonSizeValue = "small"
	ButtonSizeMedium ButtonSizeValue = "medium"
	ButtonSizeLarge  ButtonSizeValue = "large"
)

func (block ButtonBlock) Kind() theme.BlockKind {
	return theme.BlockKindButton
}

func (block ButtonBlock) TemplateData() any {
	normalized := block
	normalized.Tone = normalizedButtonTone(block.Tone)
	normalized.Style = normalizedButtonStyle(block.Style)
	normalized.Alignment = normalizedButtonAlignment(block.Alignment)
	normalized.Size = normalizedButtonSize(block.Size)
	return normalized
}

func (block ButtonBlock) RenderText(_ RenderContext) (string, error) {
	label := strings.TrimSpace(block.Label)
	url := strings.TrimSpace(block.URL)
	if label == "" && url == "" {
		return "", nil
	}
	if label == "" {
		return url, nil
	}
	if url == "" {
		return label, nil
	}

	return fmt.Sprintf("%s (%s)", label, url), nil
}

func normalizedButtonTone(value Tone) Tone {
	switch value {
	case ToneSecondary, ToneDanger, ToneDark:
		return value
	default:
		return TonePrimary
	}
}

func normalizedButtonStyle(value ButtonStyleValue) ButtonStyleValue {
	switch value {
	case ButtonStyleOutline, ButtonStyleGhost:
		return value
	default:
		return ButtonStyleFilled
	}
}

func normalizedButtonAlignment(value ButtonAlignmentValue) ButtonAlignmentValue {
	switch value {
	case ButtonAlignmentCenter, ButtonAlignmentEnd:
		return value
	default:
		return ButtonAlignmentStart
	}
}

func normalizedButtonSize(value ButtonSizeValue) ButtonSizeValue {
	switch value {
	case ButtonSizeSmall, ButtonSizeLarge:
		return value
	default:
		return ButtonSizeMedium
	}
}

func (block ButtonGroupBlock) Kind() theme.BlockKind {
	return theme.BlockKindButtonGroup
}

func (block ButtonGroupBlock) TemplateData() any {
	normalized := ButtonGroupBlock{
		Alignment:         normalizedButtonAlignment(block.Alignment),
		Joined:            block.Joined,
		Gap:               normalizedButtonGroupGap(block.Gap),
		StackOnMobile:     block.StackOnMobile,
		FullWidthOnMobile: block.FullWidthOnMobile,
		Buttons:           make([]ButtonGroupButton, 0, len(block.Buttons)),
	}

	for _, button := range block.Buttons {
		label := strings.TrimSpace(button.Label)
		url := strings.TrimSpace(button.URL)
		if label == "" || url == "" {
			continue
		}

		normalized.Buttons = append(normalized.Buttons, ButtonGroupButton{
			Label: label,
			URL:   url,
			Tone:  normalizedButtonTone(button.Tone),
			Style: normalizedButtonStyle(button.Style),
		})
	}

	return normalized
}

func normalizedButtonGroupGap(value int) int {
	if value < 0 {
		return 8
	}
	if value == 0 {
		return 8
	}

	return value
}

func (block ButtonGroupBlock) RenderText(_ RenderContext) (string, error) {
	normalized := block.TemplateData().(ButtonGroupBlock)
	parts := make([]string, 0, len(normalized.Buttons))
	for _, button := range normalized.Buttons {
		parts = append(parts, fmt.Sprintf("%s (%s)", button.Label, button.URL))
	}

	return strings.Join(parts, " · "), nil
}

func (block ButtonBlock) LayoutSpec() LayoutSpec { return defaultLayoutSpec() }

func (block ButtonGroupBlock) LayoutSpec() LayoutSpec { return defaultLayoutSpec() }
