package myrtle

import "github.com/gzuidhof/myrtle/theme"

type RenderContext struct {
	Preheader string
	Values    theme.Values
}

// InsetMode controls how a block is horizontally inset within the main content container.
// For example, an image with InsetModeDefault keeps the theme's normal side padding,
// InsetModeNone spans edge-to-edge inside the container, and InsetModeCustom uses
// LayoutSpec.CustomInset as the per-block inset value.
type InsetMode string

const (
	// InsetModeDefault uses the theme default content inset for the block.
	InsetModeDefault InsetMode = "default"
	// InsetModeNone removes side insets so the block renders full-bleed within the container.
	InsetModeNone InsetMode = "none"
	// InsetModeCustom uses LayoutSpec.CustomInset to set a custom side inset for the block.
	InsetModeCustom InsetMode = "custom"
)

// Tone defines semantic visual emphasis used by tone-aware blocks.
type Tone string

const (
	ToneDefault   Tone = "default"
	TonePrimary   Tone = "primary"
	ToneSecondary Tone = "secondary"
	ToneMuted     Tone = "muted"
	ToneInfo      Tone = "info"
	ToneSuccess   Tone = "success"
	ToneWarning   Tone = "warning"
	ToneDanger    Tone = "danger"
	ToneDark      Tone = "dark"
)

func normalizedTone(value Tone) Tone {
	switch value {
	case TonePrimary, ToneSecondary, ToneMuted, ToneInfo, ToneSuccess, ToneWarning, ToneDanger, ToneDark:
		return value
	default:
		return ToneDefault
	}
}

// LayoutSpec describes per-block layout behavior that themes can use when placing the block.
type LayoutSpec struct {
	InsetMode   InsetMode
	CustomInset string
}

// Block is the core content unit in an email, with HTML template data and text fallback rendering behavior.
type Block interface {
	// Kind returns the theme block kind used to select the HTML template for this block.
	Kind() theme.BlockKind
	// TemplateData returns the normalized data payload passed into the block's HTML template.
	TemplateData() any
	// RenderText returns the plain-text representation of the block for text-only email output.
	RenderText(context RenderContext) (string, error)
	// LayoutSpec returns layout metadata, such as inset behavior, used by theme layout templates.
	LayoutSpec() LayoutSpec
}
