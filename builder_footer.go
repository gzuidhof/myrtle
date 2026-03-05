package myrtle

// Footer replaces the current footer configuration.
// It accepts a fully configured FooterSection value.
func (builder *Builder) Footer(value FooterSection) *Builder {
	builder.mu.Lock()
	defer builder.mu.Unlock()

	if value.Block == nil {
		builder.footer = nil
		builder.footerMode = FooterModeDisabled
		return builder
	}

	builder.footerMode = FooterModeEnabled
	builder.footer = &FooterSection{
		Block:        value.Block,
		RenderInText: value.RenderInText,
		Placement:    normalizedFooterPlacement(value.Placement),
	}
	return builder
}

// WithFooter enables a footer block and applies optional footer settings.
// This is a convenience helper when starting from a plain block.
func (builder *Builder) WithFooter(block Block, options ...FooterOption) *Builder {
	builder.mu.Lock()
	defer builder.mu.Unlock()

	if block == nil {
		builder.footer = nil
		builder.footerMode = FooterModeDisabled
		return builder
	}

	footer := &FooterSection{Block: block, Placement: FooterPlacementInside}
	for _, option := range options {
		if option == nil {
			continue
		}

		option(footer)
	}
	builder.footerMode = FooterModeEnabled
	builder.footer = footer
	return builder
}

// NoFooter disables the footer section.
// The rendered email will omit any footer content.
func (builder *Builder) NoFooter() *Builder {
	builder.mu.Lock()
	defer builder.mu.Unlock()

	builder.footer = nil
	builder.footerMode = FooterModeDisabled
	return builder
}

// WithoutFooter is an alias for NoFooter.
// It provides a fluent alternative naming style.
func (builder *Builder) WithoutFooter() *Builder {
	return builder.NoFooter()
}

func cloneFooter(value *FooterSection) *FooterSection {
	if value == nil {
		return nil
	}

	copy := *value
	return &copy
}
