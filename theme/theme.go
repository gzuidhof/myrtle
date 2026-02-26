package theme

type BlockKind string

const (
	BlockKindText               BlockKind = "text"
	BlockKindHeading            BlockKind = "heading"
	BlockKindSpacer             BlockKind = "spacer"
	BlockKindList               BlockKind = "list"
	BlockKindKeyValue           BlockKind = "key_value"
	BlockKindHorizontalBarChart BlockKind = "bar_chart"
	BlockKindVerticalBarChart   BlockKind = "vertical_bar_chart"
	BlockKindSparkline          BlockKind = "sparkline"
	BlockKindStackedBar         BlockKind = "stacked_bar"
	BlockKindProgress           BlockKind = "progress"
	BlockKindDistribution       BlockKind = "distribution"
	BlockKindTimeline           BlockKind = "timeline"
	BlockKindStatsRow           BlockKind = "stats_row"
	BlockKindBadge              BlockKind = "badge"
	BlockKindSummaryCard        BlockKind = "summary_card"
	BlockKindAttachment         BlockKind = "attachment"
	BlockKindHero               BlockKind = "hero"
	BlockKindFooterLinks        BlockKind = "footer_links"
	BlockKindPriceSummary       BlockKind = "price_summary"
	BlockKindEmptyState         BlockKind = "empty_state"
	BlockKindQuote              BlockKind = "quote"
	BlockKindCallout            BlockKind = "callout"
	BlockKindMessage            BlockKind = "message"
	BlockKindMessageDigest      BlockKind = "message_digest"
	BlockKindLegal              BlockKind = "legal"
	BlockKindColumns            BlockKind = "columns"
	BlockKindSection            BlockKind = "section"
	BlockKindGrid               BlockKind = "grid"
	BlockKindCardList           BlockKind = "card_list"
	BlockKindButton             BlockKind = "button"
	BlockKindButtonGroup        BlockKind = "button_group"
	BlockKindDivider            BlockKind = "divider"
	BlockKindImage              BlockKind = "image"
	BlockKindTable              BlockKind = "table"
	BlockKindVerificationCode   BlockKind = "verification_code"
	BlockKindTiles              BlockKind = "tiles"
	BlockKindFreeMarkdown       BlockKind = "free_markdown"
)

var AllBlockKinds = []BlockKind{
	BlockKindText,
	BlockKindHeading,
	BlockKindSpacer,
	BlockKindList,
	BlockKindKeyValue,
	BlockKindHorizontalBarChart,
	BlockKindVerticalBarChart,
	BlockKindSparkline,
	BlockKindStackedBar,
	BlockKindProgress,
	BlockKindDistribution,
	BlockKindTimeline,
	BlockKindStatsRow,
	BlockKindBadge,
	BlockKindSummaryCard,
	BlockKindAttachment,
	BlockKindHero,
	BlockKindFooterLinks,
	BlockKindPriceSummary,
	BlockKindEmptyState,
	BlockKindQuote,
	BlockKindCallout,
	BlockKindMessage,
	BlockKindMessageDigest,
	BlockKindLegal,
	BlockKindColumns,
	BlockKindSection,
	BlockKindGrid,
	BlockKindCardList,
	BlockKindButton,
	BlockKindButtonGroup,
	BlockKindDivider,
	BlockKindImage,
	BlockKindTable,
	BlockKindVerificationCode,
	BlockKindTiles,
	BlockKindFreeMarkdown,
}

// BlockView is the renderer-facing payload for a single block and resolved shared values.
type BlockView struct {
	Kind   BlockKind
	Data   any
	Values Values
}

type Direction string

const (
	DirectionLTR Direction = "ltr"
	DirectionRTL Direction = "rtl"
)

// Values are shared render values available to all blocks and layouts.
type Values struct {
	Direction Direction
	Styles    Styles
}

// Styles defines shared color tokens used by themes and block templates.
type Styles struct {
	// ColorPrimary is the primary accent color for buttons, links, and highlights.
	ColorPrimary string
	// ColorSecondary is the secondary accent color for alternate emphasis/tone.
	ColorSecondary string
	// ColorText is the default foreground color for main content text.
	ColorText string
	// ColorTextMuted is the muted foreground color for metadata and secondary text.
	ColorTextMuted string
	// ColorBorder is the standard border color used by blocks and separators.
	ColorBorder string
	// ColorCodeBackground is the background color used by code/verification-style blocks.
	ColorCodeBackground string
	// ColorPageBackground is the page/shell background color outside the main content container.
	ColorPageBackground string
	// ColorMainBackground is the background color of the main content container/card.
	ColorMainBackground string
	// ColorSurface is the default surface/card background color used by blocks.
	ColorSurface string
	// ColorSurfaceMuted is the subtle surface background used for secondary areas and tracks.
	ColorSurfaceMuted string
	// ColorTextOnSolid is the text color used on strong/filled semantic or accent backgrounds.
	ColorTextOnSolid string
	// ColorInfo is the semantic info/accent strong color.
	ColorInfo string
	// ColorInfoBorder is the semantic info border color.
	ColorInfoBorder string
	// ColorInfoBackground is the semantic info soft background color.
	ColorInfoBackground string
	// ColorInfoText is the semantic info text color for soft variants.
	ColorInfoText string
	// ColorSuccess is the semantic success strong color.
	ColorSuccess string
	// ColorSuccessBorder is the semantic success border color.
	ColorSuccessBorder string
	// ColorSuccessBackground is the semantic success soft background color.
	ColorSuccessBackground string
	// ColorSuccessText is the semantic success text color for soft variants and deltas.
	ColorSuccessText string
	// ColorWarning is the semantic warning strong color.
	ColorWarning string
	// ColorWarningBorder is the semantic warning border color.
	ColorWarningBorder string
	// ColorWarningBackground is the semantic warning soft background color.
	ColorWarningBackground string
	// ColorWarningText is the semantic warning text color for soft variants.
	ColorWarningText string
	// ColorDanger is the semantic danger/error strong color.
	ColorDanger string
	// ColorDangerBorder is the semantic danger/error border color.
	ColorDangerBorder string
	// ColorDangerBackground is the semantic danger/error soft background color.
	ColorDangerBackground string
	// ColorDangerText is the semantic danger/error text color for soft variants and deltas.
	ColorDangerText string
	// BorderMain is the CSS border value applied to the main content container.
	BorderMain string
	// WidthMain is the CSS width value applied to the main content container (typically 100%).
	WidthMain string
	// MaxWidthMain is the CSS max-width value applied to the main content container.
	MaxWidthMain string
	// OuterPadding is the CSS padding value applied to the page/body wrapper around the email container.
	OuterPadding string
	// OutsideContentInset is the horizontal inset used by header/footer content rendered outside the main container.
	OutsideContentInset string
	// RadiusMain is the CSS border-radius value for the main content container.
	RadiusMain string
	// RadiusElement is the default border-radius used by card-like/content elements.
	RadiusElement string
	// RadiusButton is the border-radius used by button and button-group controls.
	RadiusButton string
	// RadiusPill is the fully-rounded border-radius used by chips, badges, and tracks.
	RadiusPill string
	// FontFamilyBase is the default font-family used for email body text.
	FontFamilyBase string
	// FontFamilyMono is the monospace font-family used for code/verification text.
	FontFamilyMono string
	// FontSizeBase is the default font-size used for body text.
	FontSizeBase string
	// LineHeightBase is the default line-height used for body text.
	LineHeightBase string
	// FontWeightHeading is the default font-weight used for heading blocks.
	FontWeightHeading string
}

// HeaderView contains the resolved header state used by HTML and text theme layouts.
type HeaderView struct {
	HTML      string
	Text      string
	Placement string
}

// FooterView contains the resolved footer state used by HTML and text theme layouts.
type FooterView struct {
	HTML      string
	Text      string
	Placement string
}

// EmailView is the complete render input for an HTML email body.
type EmailView struct {
	Header    *HeaderView
	Footer    *FooterView
	Preheader string
	Values    Values
	Blocks    []string
}

// TextView is the complete render input for wrapped plain-text output.
type TextView struct {
	Header    *HeaderView
	Footer    *FooterView
	Preheader string
	Values    Values
	Body      string
}

type Theme interface {
	Name() string
	DefaultStyles() Styles
	RenderHTML(view EmailView) (string, error)
	RenderBlockHTML(block BlockView) (string, bool, error)
}

type TextWrapper interface {
	WrapText(view TextView) (string, error)
}

// DefaultDarkModeStyles returns a ready-to-use dark mode style preset.
//
// Intended usage:
//
//	myrtle.NewBuilder(selectedTheme, myrtle.WithStyles(theme.DefaultDarkModeStyles()))
func DefaultDarkModeStyles() Styles {
	return Styles{
		ColorPrimary:           "#8b5cf6",
		ColorSecondary:         "#c084fc",
		ColorText:              "#e5e7eb",
		ColorTextMuted:         "#94a3b8",
		ColorBorder:            "#334155",
		ColorCodeBackground:    "#0f172a",
		ColorPageBackground:    "#020617",
		ColorMainBackground:    "#0b1220",
		ColorSurface:           "#0f172a",
		ColorSurfaceMuted:      "#111827",
		ColorTextOnSolid:       "#f8fafc",
		ColorInfo:              "#3b82f6",
		ColorInfoBorder:        "#1d4ed8",
		ColorInfoBackground:    "#0b2a4a",
		ColorInfoText:          "#bfdbfe",
		ColorSuccess:           "#22c55e",
		ColorSuccessBorder:     "#15803d",
		ColorSuccessBackground: "#052e16",
		ColorSuccessText:       "#86efac",
		ColorWarning:           "#f59e0b",
		ColorWarningBorder:     "#b45309",
		ColorWarningBackground: "#451a03",
		ColorWarningText:       "#fcd34d",
		ColorDanger:            "#ef4444",
		ColorDangerBorder:      "#b91c1c",
		ColorDangerBackground:  "#450a0a",
		ColorDangerText:        "#fca5a5",
		BorderMain:             "1px solid #334155",
		RadiusMain:             "16px",
		RadiusElement:          "12px",
		RadiusButton:           "8px",
		RadiusPill:             "999px",
	}
}
