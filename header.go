package myrtle

type HeaderSection struct {
	Block        Block
	RenderInText bool
	Placement    HeaderPlacementValue
}

// HeaderPlacementValue controls whether the header renders inside or outside the main container.
type HeaderPlacementValue string

const (
	HeaderPlacementInside  HeaderPlacementValue = "inside"
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

func BuildHeader(block Block, options ...HeaderOption) HeaderSection {
	result := HeaderSection{Block: block, Placement: HeaderPlacementInside}
	for _, option := range options {
		if option == nil {
			continue
		}

		option(&result)
	}

	return result
}

func normalizedHeaderPlacement(value HeaderPlacementValue) HeaderPlacementValue {
	if value == HeaderPlacementOutside {
		return HeaderPlacementOutside
	}

	return HeaderPlacementInside
}
