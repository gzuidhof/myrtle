package myrtle

import (
	"fmt"
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

// HeadingBlock renders a section heading.
type HeadingBlock struct {
	Text  string
	Level int
}

// HeadingOption configures a HeadingBlock when building content.
type HeadingOption func(*HeadingBlock)

// HeadingLevel sets the heading level on a HeadingBlock.
func HeadingLevel(value int) HeadingOption {
	return func(block *HeadingBlock) {
		if value > 0 {
			block.Level = value
		}
	}
}

func (block HeadingBlock) Kind() theme.BlockKind {
	return theme.BlockKindHeading
}

func (block HeadingBlock) TemplateData() any {
	return block
}

func (block HeadingBlock) RenderText(_ RenderContext) (string, error) {
	text := strings.TrimSpace(block.Text)
	if text == "" {
		return "", nil
	}

	dividerLength := len(text)
	if dividerLength < 5 {
		dividerLength = 5
	}
	if dividerLength > 48 {
		dividerLength = 48
	}

	return text + "\n" + strings.Repeat("-", dividerLength), nil
}

// SpacerBlock inserts vertical spacing between blocks.
type SpacerBlock struct {
	Size int
}

func (block SpacerBlock) Kind() theme.BlockKind {
	return theme.BlockKindSpacer
}

func (block SpacerBlock) TemplateData() any {
	normalized := block
	if normalized.Size <= 0 {
		normalized.Size = 16
	}

	return normalized
}

func (block SpacerBlock) RenderText(_ RenderContext) (string, error) {
	return "", nil
}

// ListBlock renders an ordered or unordered list.
type ListBlock struct {
	Items   []string
	Ordered bool
}

func (block ListBlock) Kind() theme.BlockKind {
	return theme.BlockKindList
}

func (block ListBlock) TemplateData() any {
	return block
}

func (block ListBlock) RenderText(_ RenderContext) (string, error) {
	parts := make([]string, 0, len(block.Items))
	for _, item := range block.Items {
		value := strings.TrimSpace(item)
		if value == "" {
			continue
		}
		if block.Ordered {
			parts = append(parts, fmt.Sprintf("%d. %s", len(parts)+1, value))
			continue
		}
		parts = append(parts, "- "+value)
	}

	return strings.Join(parts, "\n"), nil
}

func (block HeadingBlock) LayoutSpec() LayoutSpec { return defaultLayoutSpec() }

func (block SpacerBlock) LayoutSpec() LayoutSpec { return defaultLayoutSpec() }

func (block ListBlock) LayoutSpec() LayoutSpec { return defaultLayoutSpec() }

// ImageBlock renders an image with optional link and spacing controls.
type ImageBlock struct {
	Src              string
	Alt              string
	Href             string
	Width            int            // px, 0 means auto
	Align            ImageAlignment // "center", "start", "end", "full", "" (default)
	HasTopSpacing    bool
	TopSpacing       int
	HasBottomSpacing bool
	BottomSpacing    int
	CornerMode       ImageCornerMode
	InsetMode        InsetMode
}

// ImageCornerMode controls how image corner radii are applied.
type ImageCornerMode string

const (
	ImageCornerModeAuto   ImageCornerMode = "auto"
	ImageCornerModeNone   ImageCornerMode = "none"
	ImageCornerModeAll    ImageCornerMode = "all"
	ImageCornerModeTop    ImageCornerMode = "top"
	ImageCornerModeBottom ImageCornerMode = "bottom"
)

// ImageAlignment controls image horizontal alignment and width behavior.
type ImageAlignment string

const (
	ImageAlignmentDefault ImageAlignment = ""
	ImageAlignmentCenter  ImageAlignment = "center"
	ImageAlignmentStart   ImageAlignment = "start"
	ImageAlignmentEnd     ImageAlignment = "end"
	ImageAlignmentFull    ImageAlignment = "full"
)

func (block ImageBlock) Kind() theme.BlockKind {
	return theme.BlockKindImage
}

func (block ImageBlock) TemplateData() any {
	normalized := block
	normalized.Align = normalizedImageAlignment(block.Align)
	if normalized.HasTopSpacing && normalized.TopSpacing < 0 {
		normalized.TopSpacing = 0
	}
	if normalized.HasBottomSpacing && normalized.BottomSpacing < 0 {
		normalized.BottomSpacing = 0
	}
	normalized.CornerMode = normalizedImageCornerMode(normalized.CornerMode)
	normalized.InsetMode = normalizedLayoutSpec(LayoutSpec{InsetMode: normalized.InsetMode}).InsetMode
	return normalized
}

func (block ImageBlock) RenderText(_ RenderContext) (string, error) {
	alt := strings.TrimSpace(block.Alt)
	src := strings.TrimSpace(block.Src)
	if alt == "" && src == "" {
		return "", nil
	}
	if alt == "" {
		return src, nil
	}
	if src == "" {
		return alt, nil
	}

	return fmt.Sprintf("%s (%s)", alt, src), nil
}

func normalizedImageAlignment(value ImageAlignment) ImageAlignment {
	switch value {
	case ImageAlignmentCenter, ImageAlignmentStart, ImageAlignmentEnd, ImageAlignmentFull:
		return value
	default:
		return ImageAlignmentDefault
	}
}

func normalizedImageCornerMode(value ImageCornerMode) ImageCornerMode {
	switch value {
	case ImageCornerModeNone, ImageCornerModeAll, ImageCornerModeTop, ImageCornerModeBottom:
		return value
	default:
		return ImageCornerModeAuto
	}
}

func (block ImageBlock) LayoutSpec() LayoutSpec {
	return normalizedLayoutSpec(LayoutSpec{InsetMode: block.InsetMode})
}

// LegalBlock renders company and subscription-management compliance text.
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

func (block LegalBlock) LayoutSpec() LayoutSpec { return defaultLayoutSpec() }
