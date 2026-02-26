package myrtle

import (
	"strings"

	"github.com/gzuidhof/myrtle/theme"
)

type TextBlock struct {
	Text      string
	Tone      TextToneValue
	Size      TextSizeValue
	Align     TextAlignmentValue
	Weight    TextWeightValue
	NoMargin  bool
	Spacing   TextSpacingValue
	Transform TextTransformValue
}

type TextOption func(*TextBlock)

type TextToneValue string

const (
	TextToneDefault TextToneValue = "default"
	TextToneMuted   TextToneValue = "muted"
	TextToneInfo    TextToneValue = "info"
	TextToneSuccess TextToneValue = "success"
	TextToneWarning TextToneValue = "warning"
	TextToneDanger  TextToneValue = "danger"
)

type TextAlignmentValue string

const (
	TextAlignStart  TextAlignmentValue = "start"
	TextAlignCenter TextAlignmentValue = "center"
	TextAlignEnd    TextAlignmentValue = "end"
)

type TextWeightValue string

const (
	TextWeightNormal   TextWeightValue = "normal"
	TextWeightMedium   TextWeightValue = "medium"
	TextWeightSemibold TextWeightValue = "semibold"
	TextWeightBold     TextWeightValue = "bold"
)

type TextSpacingValue string

const (
	TextSpacingCompact TextSpacingValue = "compact"
	TextSpacingNormal  TextSpacingValue = "normal"
	TextSpacingRelaxed TextSpacingValue = "relaxed"
)

type TextTransformValue string

const (
	TextTransformNone       TextTransformValue = "none"
	TextTransformUppercase  TextTransformValue = "uppercase"
	TextTransformLowercase  TextTransformValue = "lowercase"
	TextTransformCapitalize TextTransformValue = "capitalize"
)

type TextSizeValue string

const (
	TextSizeSmall TextSizeValue = "small"
	TextSizeBase  TextSizeValue = "base"
	TextSizeLarge TextSizeValue = "large"
)

func TextSize(value TextSizeValue) TextOption {
	return func(block *TextBlock) {
		switch value {
		case TextSizeSmall, TextSizeBase, TextSizeLarge:
			block.Size = value
		}
	}
}

func TextTone(value TextToneValue) TextOption {
	return func(block *TextBlock) {
		switch value {
		case TextToneDefault, TextToneMuted, TextToneInfo, TextToneSuccess, TextToneWarning, TextToneDanger:
			block.Tone = value
		}
	}
}

func TextAlign(value TextAlignmentValue) TextOption {
	return func(block *TextBlock) {
		switch value {
		case TextAlignStart, TextAlignCenter, TextAlignEnd:
			block.Align = value
		}
	}
}

func TextWeight(value TextWeightValue) TextOption {
	return func(block *TextBlock) {
		switch value {
		case TextWeightNormal, TextWeightMedium, TextWeightSemibold, TextWeightBold:
			block.Weight = value
		}
	}
}

func TextNoMargin(value bool) TextOption {
	return func(block *TextBlock) {
		block.NoMargin = value
	}
}

func TextSpacing(value TextSpacingValue) TextOption {
	return func(block *TextBlock) {
		switch value {
		case TextSpacingCompact, TextSpacingNormal, TextSpacingRelaxed:
			block.Spacing = value
		}
	}
}

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
