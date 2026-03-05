package myrtle

import (
	"strings"

	"github.com/gzuidhof/myrtle/theme"
)

// RenderContext contains resolved email values available to block text renderers.
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
	// ToneDefault renders neutral/default styling.
	ToneDefault Tone = "default"
	// TonePrimary renders primary accent styling.
	TonePrimary Tone = "primary"
	// ToneSecondary renders secondary accent styling.
	ToneSecondary Tone = "secondary"
	// ToneMuted renders de-emphasized styling.
	ToneMuted Tone = "muted"
	// ToneInfo renders informational semantic styling.
	ToneInfo Tone = "info"
	// ToneSuccess renders success semantic styling.
	ToneSuccess Tone = "success"
	// ToneWarning renders warning semantic styling.
	ToneWarning Tone = "warning"
	// ToneDanger renders danger/error semantic styling.
	ToneDanger Tone = "danger"
	// ToneDark renders high-contrast dark styling.
	ToneDark Tone = "dark"
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

func normalizedLayoutSpec(spec LayoutSpec) LayoutSpec {
	if spec.InsetMode != InsetModeNone && spec.InsetMode != InsetModeCustom {
		spec.InsetMode = InsetModeDefault
	}

	spec.CustomInset = strings.TrimSpace(spec.CustomInset)
	if spec.InsetMode == InsetModeCustom && spec.CustomInset == "" {
		spec.InsetMode = InsetModeDefault
	}

	return spec
}

func defaultLayoutSpec() LayoutSpec {
	return LayoutSpec{InsetMode: InsetModeDefault}
}

func (group *Group) Kind() theme.BlockKind {
	return theme.BlockKind("group")
}

func (group *Group) TemplateData() any {
	return group
}

func (group *Group) RenderText(context RenderContext) (string, error) {
	if group == nil {
		return "", nil
	}

	return renderColumnText(group.Blocks(), context)
}

func (group *Group) LayoutSpec() LayoutSpec { return defaultLayoutSpec() }
