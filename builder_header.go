package myrtle

import "github.com/gzuidhof/myrtle/theme"

const defaultPreheaderPaddingRepeat = 10

type preheaderConfig struct {
	paddingRepeat int
}

// PreheaderOption configures preheader rendering behavior.
type PreheaderOption func(*preheaderConfig)

// PreheaderPaddingRepeat sets how many hidden "&nbsp;&zwnj;" pairs are appended to the HTML preheader.
// Values below zero are clamped to zero.
func PreheaderPaddingRepeat(value int) PreheaderOption {
	return func(config *preheaderConfig) {
		config.paddingRepeat = value
	}
}

// WithPreheader sets the preheader text shown by email clients in inbox previews.
// This text appears near the subject line in many inbox list views.
func (builder *Builder) WithPreheader(value string, options ...PreheaderOption) *Builder {
	builder.mu.Lock()
	defer builder.mu.Unlock()

	config := preheaderConfig{paddingRepeat: defaultPreheaderPaddingRepeat}
	for _, option := range options {
		if option == nil {
			continue
		}

		option(&config)
	}
	if config.paddingRepeat < 0 {
		config.paddingRepeat = 0
	}

	builder.preheader = value
	builder.preheaderPaddingRepeat = config.paddingRepeat
	return builder
}

// WithDirection sets text direction for rendered email content.
// Use RTL for right-to-left scripts and LTR for default left-to-right text.
func (builder *Builder) WithDirection(value theme.Direction) *Builder {
	builder.mu.Lock()
	defer builder.mu.Unlock()

	if value == theme.DirectionRTL {
		builder.values.Direction = theme.DirectionRTL
		return builder
	}

	builder.values.Direction = theme.DirectionLTR
	return builder
}

// Header replaces the current header configuration.
// It accepts a fully configured HeaderSection value.
func (builder *Builder) Header(value HeaderSection) *Builder {
	builder.mu.Lock()
	defer builder.mu.Unlock()

	if value.Block == nil {
		builder.header = nil
		builder.headerMode = HeaderModeDisabled
		return builder
	}

	builder.headerMode = HeaderModeEnabled
	builder.header = &HeaderSection{
		Block:        value.Block,
		RenderInText: value.RenderInText,
		Placement:    normalizedHeaderPlacement(value.Placement),
	}
	return builder
}

// WithHeader enables a header block and applies optional header settings.
// This is a convenience helper when starting from a plain block.
func (builder *Builder) WithHeader(block Block, options ...HeaderOption) *Builder {
	builder.mu.Lock()
	defer builder.mu.Unlock()

	if block == nil {
		builder.header = nil
		builder.headerMode = HeaderModeDisabled
		return builder
	}

	header := &HeaderSection{Block: block, Placement: HeaderPlacementInside}
	for _, option := range options {
		if option == nil {
			continue
		}

		option(header)
	}
	builder.headerMode = HeaderModeEnabled
	builder.header = header
	return builder
}

// NoHeader disables the header section.
// The rendered email will omit any header content.
func (builder *Builder) NoHeader() *Builder {
	builder.mu.Lock()
	defer builder.mu.Unlock()

	builder.header = nil
	builder.headerMode = HeaderModeDisabled
	return builder
}

// WithoutHeader is an alias for NoHeader.
// It provides a fluent alternative naming style.
func (builder *Builder) WithoutHeader() *Builder {
	return builder.NoHeader()
}

func cloneHeader(value *HeaderSection) *HeaderSection {
	if value == nil {
		return nil
	}

	copy := *value
	return &copy
}
