package myrtle

func (builder *Builder) Preheader(value string) *Builder {
	builder.preheader = value
	return builder
}

func (builder *Builder) Header(value HeaderSection) *Builder {
	builder.headerMode = HeaderModeEnabled
	builder.header = &HeaderSection{
		Title:            value.Title,
		ProductName:      value.ProductName,
		ProductLink:      value.ProductLink,
		LogoURL:          value.LogoURL,
		LogoAlt:          value.LogoAlt,
		RenderInMarkdown: value.RenderInMarkdown,
		ShowTextWithLogo: value.ShowTextWithLogo,
		LogoCentered:     value.LogoCentered,
		Alignment:        value.Alignment,
	}
	builder.syncValuesFromHeader()
	return builder
}

func (builder *Builder) WithHeader(options ...HeaderOption) *Builder {
	header := builder.ensureHeaderExplicit()
	for _, option := range options {
		option(header)
	}
	builder.syncValuesFromHeader()
	return builder
}

func (builder *Builder) NoHeader() *Builder {
	builder.header = nil
	builder.headerMode = HeaderModeDisabled
	return builder
}

func (builder *Builder) WithoutHeader() *Builder {
	return builder.NoHeader()
}

func (builder *Builder) Product(name, link string) *Builder {
	builder.values.ProductName = name
	builder.values.ProductLink = link
	if header := builder.ensureHeaderImplicit(); header != nil {
		header.ProductName = name
		header.ProductLink = link
	}
	return builder
}

func (builder *Builder) ProductName(value string) *Builder {
	builder.values.ProductName = value
	if header := builder.ensureHeaderImplicit(); header != nil {
		header.ProductName = value
	}
	return builder
}

func (builder *Builder) ProductLink(value string) *Builder {
	builder.values.ProductLink = value
	if header := builder.ensureHeaderImplicit(); header != nil {
		header.ProductLink = value
	}
	return builder
}

func (builder *Builder) Logo(url, alt string) *Builder {
	builder.values.LogoURL = url
	builder.values.LogoAlt = alt
	if header := builder.ensureHeaderImplicit(); header != nil {
		header.LogoURL = url
		header.LogoAlt = alt
	}
	return builder
}

func (builder *Builder) ensureHeaderExplicit() *HeaderSection {
	builder.headerMode = HeaderModeEnabled

	if builder.header == nil {
		builder.header = &HeaderSection{}
	}

	return builder.header
}

func (builder *Builder) ensureHeaderImplicit() *HeaderSection {
	if builder.header != nil {
		return builder.header
	}

	if builder.headerMode == HeaderModeDisabled {
		return nil
	}

	builder.header = &HeaderSection{}
	return builder.header
}

func (builder *Builder) syncValuesFromHeader() {
	if builder.header == nil {
		return
	}

	builder.values.ProductName = builder.header.ProductName
	builder.values.ProductLink = builder.header.ProductLink
	builder.values.LogoURL = builder.header.LogoURL
	builder.values.LogoAlt = builder.header.LogoAlt
}

func cloneHeader(value *HeaderSection) *HeaderSection {
	if value == nil {
		return nil
	}

	copy := *value
	return &copy
}
