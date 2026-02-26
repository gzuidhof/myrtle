package myrtle

type HeaderSection struct {
	Block        Block
	RenderInText bool
	Placement    HeaderPlacementValue
}

type HeaderPlacementValue string

const (
	HeaderPlacementInside  HeaderPlacementValue = "inside"
	HeaderPlacementOutside HeaderPlacementValue = "outside"
)

type HeaderOption func(*HeaderSection)

func HeaderRenderInText(value bool) HeaderOption {
	return func(header *HeaderSection) {
		header.RenderInText = value
	}
}

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
