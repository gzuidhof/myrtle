package myrtle

type FooterSection struct {
	Block        Block
	RenderInText bool
	Placement    FooterPlacementValue
}

type FooterPlacementValue string

const (
	FooterPlacementInside  FooterPlacementValue = "inside"
	FooterPlacementOutside FooterPlacementValue = "outside"
)

type FooterOption func(*FooterSection)

func FooterRenderInText(value bool) FooterOption {
	return func(footer *FooterSection) {
		footer.RenderInText = value
	}
}

func FooterPlacement(value FooterPlacementValue) FooterOption {
	return func(footer *FooterSection) {
		footer.Placement = normalizedFooterPlacement(value)
	}
}

func BuildFooter(block Block, options ...FooterOption) FooterSection {
	result := FooterSection{Block: block, Placement: FooterPlacementInside}
	for _, option := range options {
		if option == nil {
			continue
		}

		option(&result)
	}

	return result
}

func normalizedFooterPlacement(value FooterPlacementValue) FooterPlacementValue {
	if value == FooterPlacementOutside {
		return FooterPlacementOutside
	}

	return FooterPlacementInside
}
