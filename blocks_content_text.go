package myrtle

import (
	"strings"

	"github.com/gzuidhof/myrtle/theme"
)

// TextBlock renders styled paragraph text.
type TextBlock struct {
	Text      string
	Tone      Tone
	Size      TextSizeValue
	Align     TextAlignmentValue
	Weight    TextWeightValue
	NoMargin  bool
	Spacing   TextSpacingValue
	Transform TextTransformValue
}

// TextOption configures a TextBlock.
type TextOption func(*TextBlock)

// TextAlignmentValue defines logical text alignment.
type TextAlignmentValue string

const (
	TextAlignStart  TextAlignmentValue = "start"
	TextAlignCenter TextAlignmentValue = "center"
	TextAlignEnd    TextAlignmentValue = "end"
)

// TextWeightValue defines text font-weight presets.
type TextWeightValue string

const (
	TextWeightNormal   TextWeightValue = "normal"
	TextWeightMedium   TextWeightValue = "medium"
	TextWeightSemibold TextWeightValue = "semibold"
	TextWeightBold     TextWeightValue = "bold"
)

// TextSpacingValue defines line-height spacing presets.
type TextSpacingValue string

const (
	TextSpacingCompact TextSpacingValue = "compact"
	TextSpacingNormal  TextSpacingValue = "normal"
	TextSpacingRelaxed TextSpacingValue = "relaxed"
)

// TextTransformValue defines text transform behavior.
type TextTransformValue string

const (
	TextTransformNone       TextTransformValue = "none"
	TextTransformUppercase  TextTransformValue = "uppercase"
	TextTransformLowercase  TextTransformValue = "lowercase"
	TextTransformCapitalize TextTransformValue = "capitalize"
)

// TextSizeValue defines text size presets.
type TextSizeValue string

const (
	TextSizeSmall TextSizeValue = "small"
	TextSizeBase  TextSizeValue = "base"
	TextSizeLarge TextSizeValue = "large"
)

// TextSize sets the size preset for a text block.
func TextSize(value TextSizeValue) TextOption {
	return func(block *TextBlock) {
		switch value {
		case TextSizeSmall, TextSizeBase, TextSizeLarge:
			block.Size = value
		}
	}
}

// TextTone sets the semantic tone for a text block.
func TextTone(value Tone) TextOption {
	return func(block *TextBlock) {
		switch value {
		case ToneDefault, ToneMuted, ToneInfo, ToneSuccess, ToneWarning, ToneDanger, ToneDark:
			block.Tone = value
		}
	}
}

// TextAlign sets logical alignment for a text block.
func TextAlign(value TextAlignmentValue) TextOption {
	return func(block *TextBlock) {
		switch value {
		case TextAlignStart, TextAlignCenter, TextAlignEnd:
			block.Align = value
		}
	}
}

// TextWeight sets the font weight preset for a text block.
func TextWeight(value TextWeightValue) TextOption {
	return func(block *TextBlock) {
		switch value {
		case TextWeightNormal, TextWeightMedium, TextWeightSemibold, TextWeightBold:
			block.Weight = value
		}
	}
}

// TextNoMargin toggles paragraph bottom margin for a text block.
func TextNoMargin(value bool) TextOption {
	return func(block *TextBlock) {
		block.NoMargin = value
	}
}

// TextSpacing sets line-height spacing preset for a text block.
func TextSpacing(value TextSpacingValue) TextOption {
	return func(block *TextBlock) {
		switch value {
		case TextSpacingCompact, TextSpacingNormal, TextSpacingRelaxed:
			block.Spacing = value
		}
	}
}

// TextTransform sets text transform behavior for a text block.
func TextTransform(value TextTransformValue) TextOption {
	return func(block *TextBlock) {
		switch value {
		case TextTransformNone, TextTransformUppercase, TextTransformLowercase, TextTransformCapitalize:
			block.Transform = value
		}
	}
}

func (block TextBlock) Kind() theme.BlockKind {
	return theme.BlockKindText
}

func (block TextBlock) TemplateData() any {
	return block
}

func (block TextBlock) RenderText(_ RenderContext) (string, error) {
	return strings.TrimSpace(block.Text), nil
}

func (block TextBlock) LayoutSpec() LayoutSpec { return defaultLayoutSpec() }
