package myrtle

// HeaderSection stores header content and placement settings for an email.
type HeaderSection struct {
	Block        Block
	RenderInText bool
	Placement    HeaderPlacementValue
}

// HeaderPlacementValue controls whether the header renders inside or outside the main container.
type HeaderPlacementValue string

const (
	// HeaderPlacementInside renders the header within the main content container.
	HeaderPlacementInside HeaderPlacementValue = "inside"
	// HeaderPlacementOutside renders the header outside the main content container.
	HeaderPlacementOutside HeaderPlacementValue = "outside"
)

// HeaderOption configures a HeaderSection.
type HeaderOption func(*HeaderSection)

// HeaderRenderInText controls whether the header is included in text output.
func HeaderRenderInText(value bool) HeaderOption {
	return func(header *HeaderSection) {
		header.RenderInText = value
	}
}

// HeaderPlacement sets whether the header renders inside or outside the main container.
func HeaderPlacement(value HeaderPlacementValue) HeaderOption {
	return func(header *HeaderSection) {
		header.Placement = normalizedHeaderPlacement(value)
	}
}

func normalizedHeaderPlacement(value HeaderPlacementValue) HeaderPlacementValue {
	if value == HeaderPlacementOutside {
		return HeaderPlacementOutside
	}

	return HeaderPlacementInside
}

// FooterSection stores footer content and placement settings for an email.
type FooterSection struct {
	Block        Block
	RenderInText bool
	Placement    FooterPlacementValue
}

// FooterPlacementValue controls whether the footer renders inside or outside the main container.
type FooterPlacementValue string

const (
	// FooterPlacementInside renders the footer within the main content container.
	FooterPlacementInside FooterPlacementValue = "inside"
	// FooterPlacementOutside renders the footer outside the main content container.
	FooterPlacementOutside FooterPlacementValue = "outside"
)

// FooterOption configures a FooterSection.
type FooterOption func(*FooterSection)

// FooterRenderInText controls whether the footer is included in text output.
func FooterRenderInText(value bool) FooterOption {
	return func(footer *FooterSection) {
		footer.RenderInText = value
	}
}

// FooterPlacement sets whether the footer renders inside or outside the main container.
func FooterPlacement(value FooterPlacementValue) FooterOption {
	return func(footer *FooterSection) {
		footer.Placement = normalizedFooterPlacement(value)
	}
}

func normalizedFooterPlacement(value FooterPlacementValue) FooterPlacementValue {
	if value == FooterPlacementOutside {
		return FooterPlacementOutside
	}

	return FooterPlacementInside
}
