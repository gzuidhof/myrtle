package myrtle

type HeaderSection struct {
	Title            string
	ProductName      string
	ProductLink      string
	LogoURL          string
	LogoAlt          string
	RenderInMarkdown bool
	ShowTextWithLogo bool
	LogoCentered     bool
	Alignment        HeaderAlignment
}

type HeaderAlignment string

const (
	HeaderAlignmentCenter HeaderAlignment = "center"
	HeaderAlignmentLeft   HeaderAlignment = "left"
)

type HeaderOption func(*HeaderSection)

func HeaderTitle(value string) HeaderOption {
	return func(header *HeaderSection) {
		header.Title = value
	}
}

func HeaderProduct(name, link string) HeaderOption {
	return func(header *HeaderSection) {
		header.ProductName = name
		header.ProductLink = link
	}
}

func HeaderLogo(url, alt string) HeaderOption {
	return func(header *HeaderSection) {
		header.LogoURL = url
		header.LogoAlt = alt
	}
}

func HeaderShowTextWithLogo(value bool) HeaderOption {
	return func(header *HeaderSection) {
		header.ShowTextWithLogo = value
	}
}

func HeaderRenderInMarkdown(value bool) HeaderOption {
	return func(header *HeaderSection) {
		header.RenderInMarkdown = value
	}
}

func HeaderLogoCentered(value bool) HeaderOption {
	return func(header *HeaderSection) {
		header.LogoCentered = value
		if value {
			header.Alignment = HeaderAlignmentCenter
			return
		}

		header.Alignment = HeaderAlignmentLeft
	}
}

func HeaderAlign(alignment HeaderAlignment) HeaderOption {
	return func(header *HeaderSection) {
		header.Alignment = normalizedHeaderAlignment(alignment)
		header.LogoCentered = header.Alignment == HeaderAlignmentCenter
	}
}

func BuildHeader(options ...HeaderOption) HeaderSection {
	result := HeaderSection{}
	for _, option := range options {
		option(&result)
	}

	result.Alignment = normalizedHeaderAlignment(result.Alignment)
	result.LogoCentered = result.Alignment == HeaderAlignmentCenter

	return result
}

func normalizedHeaderAlignment(value HeaderAlignment) HeaderAlignment {
	switch value {
	case HeaderAlignmentLeft:
		return HeaderAlignmentLeft
	default:
		return HeaderAlignmentCenter
	}
}
