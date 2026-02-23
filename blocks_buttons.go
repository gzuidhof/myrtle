package myrtle

import (
	"fmt"
	"strings"

	"github.com/gzuidhof/myrtle/theme"
)

type ButtonBlock struct {
	Label     string
	URL       string
	Tone      ButtonToneValue
	Style     ButtonStyleValue
	Alignment ButtonAlignment
	Size      ButtonSizeValue
	NoWrap    bool
	FullWidth bool
}

type ButtonGroupButton struct {
	Label string
	URL   string
	Tone  ButtonToneValue
	Style ButtonStyleValue
}

type ButtonGroupBlock struct {
	Buttons           []ButtonGroupButton
	Alignment         ButtonAlignment
	Joined            bool
	Gap               int
	StackOnMobile     bool
	FullWidthOnMobile bool
}

type (
	ButtonToneValue  string
	ButtonStyleValue string
	ButtonAlignment  string
	ButtonSizeValue  string
)

const (
	ButtonTonePrimary   ButtonToneValue = "primary"
	ButtonToneSecondary ButtonToneValue = "secondary"
	ButtonToneDanger    ButtonToneValue = "danger"

	ButtonStyleFilled  ButtonStyleValue = "filled"
	ButtonStyleOutline ButtonStyleValue = "outline"
	ButtonStyleGhost   ButtonStyleValue = "ghost"

	ButtonAlignmentLeft   ButtonAlignment = "left"
	ButtonAlignmentCenter ButtonAlignment = "center"
	ButtonAlignmentRight  ButtonAlignment = "right"

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

func (block ButtonBlock) RenderMarkdown(_ RenderContext) (string, error) {
	return fmt.Sprintf("[%s](%s)", block.Label, block.URL), nil
}

func normalizedButtonTone(value ButtonToneValue) ButtonToneValue {
	switch value {
	case ButtonToneSecondary, ButtonToneDanger:
		return value
	default:
		return ButtonTonePrimary
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

func normalizedButtonAlignment(value ButtonAlignment) ButtonAlignment {
	switch value {
	case ButtonAlignmentCenter, ButtonAlignmentRight:
		return value
	default:
		return ButtonAlignmentLeft
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

func (block ButtonGroupBlock) RenderMarkdown(_ RenderContext) (string, error) {
	normalized := block.TemplateData().(ButtonGroupBlock)
	parts := make([]string, 0, len(normalized.Buttons))
	for _, button := range normalized.Buttons {
		parts = append(parts, fmt.Sprintf("[%s](%s)", button.Label, button.URL))
	}

	return strings.Join(parts, " · "), nil
}
