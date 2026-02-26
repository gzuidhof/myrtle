package myrtle

import "github.com/gzuidhof/myrtle/theme"

func (builder *Builder) WithPreheader(value string) *Builder {
	builder.mu.Lock()
	defer builder.mu.Unlock()

	builder.preheader = value
	return builder
}

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

func (builder *Builder) NoHeader() *Builder {
	builder.mu.Lock()
	defer builder.mu.Unlock()

	builder.header = nil
	builder.headerMode = HeaderModeDisabled
	return builder
}

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
