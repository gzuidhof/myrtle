package myrtle

type FooterSection struct {
	Block        Block
	RenderInText bool
	Placement    FooterPlacementValue
}

// FooterPlacementValue controls whether the footer renders inside or outside the main container.
type FooterPlacementValue string

const (
	FooterPlacementInside  FooterPlacementValue = "inside"
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
