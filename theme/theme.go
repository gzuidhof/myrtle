package theme

// BlockKind identifies the renderer template kind for a block.
type BlockKind string

const (
	// BlockKindText renders a paragraph/text block.
	BlockKindText BlockKind = "text"
	// BlockKindHeading renders a heading block.
	BlockKindHeading BlockKind = "heading"
	// BlockKindSpacer renders a spacing/separator height block.
	BlockKindSpacer BlockKind = "spacer"
	// BlockKindList renders an ordered or unordered list block.
	BlockKindList BlockKind = "list"
	// BlockKindKeyValue renders key-value table-like rows.
	BlockKindKeyValue BlockKind = "key_value"
	// BlockKindHorizontalBarChart renders a horizontal bar chart block.
	BlockKindHorizontalBarChart BlockKind = "bar_chart"
	// BlockKindVerticalBarChart renders a vertical bar chart block.
	BlockKindVerticalBarChart BlockKind = "vertical_bar_chart"
	// BlockKindSparkline renders a compact sparkline chart block.
	BlockKindSparkline BlockKind = "sparkline"
	// BlockKindStackedBar renders a stacked bar visualization block.
	BlockKindStackedBar BlockKind = "stacked_bar"
	// BlockKindProgress renders a progress indicator block.
	BlockKindProgress BlockKind = "progress"
	// BlockKindDistribution renders a distribution/segment summary block.
	BlockKindDistribution BlockKind = "distribution"
	// BlockKindTimeline renders a timeline/milestone block.
	BlockKindTimeline BlockKind = "timeline"
	// BlockKindStatsRow renders a compact row of stats.
	BlockKindStatsRow BlockKind = "stats_row"
	// BlockKindBadge renders a small semantic badge/chip block.
	BlockKindBadge BlockKind = "badge"
	// BlockKindSummaryCard renders a title/body/footer summary card.
	BlockKindSummaryCard BlockKind = "summary_card"
	// BlockKindAttachment renders an attachment metadata block.
	BlockKindAttachment BlockKind = "attachment"
	// BlockKindHero renders a prominent hero/banner block.
	BlockKindHero BlockKind = "hero"
	// BlockKindFooterLinks renders a footer links/navigation block.
	BlockKindFooterLinks BlockKind = "footer_links"
	// BlockKindPriceSummary renders an itemized pricing summary block.
	BlockKindPriceSummary BlockKind = "price_summary"
	// BlockKindEmptyState renders an empty-state placeholder block.
	BlockKindEmptyState BlockKind = "empty_state"
	// BlockKindQuote renders a quote/testimonial block.
	BlockKindQuote BlockKind = "quote"
	// BlockKindCallout renders an informational/warning callout block.
	BlockKindCallout BlockKind = "callout"
	// BlockKindMessage renders a single message item block.
	BlockKindMessage BlockKind = "message"
	// BlockKindMessageDigest renders a grouped message digest block.
	BlockKindMessageDigest BlockKind = "message_digest"
	// BlockKindLegal renders legal/compliance footer content.
	BlockKindLegal BlockKind = "legal"
	// BlockKindColumns renders a two-column layout block.
	BlockKindColumns BlockKind = "columns"
	// BlockKindPanel renders a bordered panel/container block.
	BlockKindPanel BlockKind = "panel"
	// BlockKindGrid renders a generic multi-column grid block.
	BlockKindGrid BlockKind = "grid"
	// BlockKindCardList renders a list of card items in columns.
	BlockKindCardList BlockKind = "card_list"
	// BlockKindButton renders a single CTA button block.
	BlockKindButton BlockKind = "button"
	// BlockKindButtonGroup renders multiple grouped CTA buttons.
	BlockKindButtonGroup BlockKind = "button_group"
	// BlockKindDivider renders a horizontal divider block.
	BlockKindDivider BlockKind = "divider"
	// BlockKindImage renders an image/media block.
	BlockKindImage BlockKind = "image"
	// BlockKindTable renders a data table block.
	BlockKindTable BlockKind = "table"
	// BlockKindVerificationCode renders a one-time verification code block.
	BlockKindVerificationCode BlockKind = "verification_code"
	// BlockKindTiles renders a tiled metrics/content block.
	BlockKindTiles BlockKind = "tiles"
	// BlockKindFreeMarkdown renders arbitrary markdown content.
	BlockKindFreeMarkdown BlockKind = "free_markdown"
)

// AllBlockKinds lists every registered block kind in stable renderer order.
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
	BlockKindPanel,
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

// Direction controls text/layout direction for a rendered email.
type Direction string

const (
	// DirectionLTR is the default left-to-right text and layout direction.
	DirectionLTR Direction = "ltr"
	// DirectionRTL is the right-to-left text and layout direction for languages such as Arabic and Hebrew.
	DirectionRTL Direction = "rtl"
)

// Values are shared render values available to all blocks and layouts.
type Values struct {
	Direction Direction
	Styles    Styles
}

// MSOCompatibilityMode controls whether Outlook-specific compatibility fallbacks are rendered.
type MSOCompatibilityMode string

const (
	// MSOCompatibilityModeOn enables Outlook compatibility fallbacks.
	MSOCompatibilityModeOn MSOCompatibilityMode = "on"
	// MSOCompatibilityModeOff disables Outlook compatibility fallbacks.
	MSOCompatibilityModeOff MSOCompatibilityMode = "off"
)

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
	// MainContentBodyTopSpacing is the top spacing inserted before the first body block when no inside header is present.
	// Set to "0" to disable this spacing.
	MainContentBodyTopSpacing string
	// MSOCompatibility controls whether Outlook-specific compatibility fallbacks are rendered.
	// Use MSOCompatibilityModeOn (default) to enable or MSOCompatibilityModeOff to disable.
	MSOCompatibility MSOCompatibilityMode
	// RadiusMain is the CSS border-radius value for the main content container.
	RadiusMain string
	// RadiusElement is the default border-radius used by card-like/content elements.
	RadiusElement string
	// RadiusButton is the border-radius used by button and button-group controls.
	RadiusButton string
	// RadiusPill is the fully-rounded border-radius used by chips, badges, and tracks.
	RadiusPill string
	// TableLegendSwatchSize is the size of table legend swatches rendered by TableLegendSwatches.
	TableLegendSwatchSize string
	// TableLegendSwatchRadius is the corner radius of table legend swatches rendered by TableLegendSwatches.
	TableLegendSwatchRadius string
	// TableLegendSwatchBorder is an optional CSS border value for table legend swatches rendered by TableLegendSwatches.
	TableLegendSwatchBorder string
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
	HTML        string
	Text        string
	Placement   string
	InsetMode   string
	CustomInset string
}

// FooterView contains the resolved footer state used by HTML and text theme layouts.
type FooterView struct {
	HTML        string
	Text        string
	Placement   string
	InsetMode   string
	CustomInset string
}

// EmailBlockView is a rendered HTML block with optional layout metadata.
type EmailBlockView struct {
	Kind        BlockKind
	HTML        string
	InsetMode   string
	CustomInset string
}

// EmailView is the complete render input for an HTML email body.
type EmailView struct {
	Header                 *HeaderView
	Footer                 *FooterView
	Preheader              string
	PreheaderPaddingRepeat int
	Values                 Values
	Blocks                 []string
	BlockViews             []EmailBlockView
}

// TextView is the complete render input for wrapped plain-text output.
type TextView struct {
	Header    *HeaderView
	Footer    *FooterView
	Preheader string
	Values    Values
	Body      string
}

// Theme defines the rendering contract for HTML emails and per-block HTML fallbacks.
type Theme interface {
	// Name returns the stable theme identifier used for previews and diagnostics.
	Name() string
	// DefaultStyles returns the base style token set used by this theme.
	DefaultStyles() Styles
	// RenderHTML renders a complete HTML email document from a resolved EmailView.
	RenderHTML(view EmailView) (string, error)
	// RenderBlockHTML renders one block kind in this theme and reports whether it handled the block.
	RenderBlockHTML(block BlockView) (string, bool, error)
}

// TextWrapper defines plain-text output wrapping behavior for rendered TextView payloads.
type TextWrapper interface {
	// WrapText renders and wraps a text email body from resolved text view data.
	WrapText(view TextView) (string, error)
}

// DefaultDarkModeStyles returns a ready-to-use dark mode style preset.
//
// Intended usage:
//
//	myrtle.NewBuilder(selectedTheme, myrtle.WithStyles(theme.DefaultDarkModeStyles()))
func DefaultDarkModeStyles() Styles {
	return Styles{
		ColorPrimary:              "#8b5cf6",
		ColorSecondary:            "#c084fc",
		ColorText:                 "#e5e7eb",
		ColorTextMuted:            "#94a3b8",
		ColorBorder:               "#334155",
		ColorCodeBackground:       "#0f172a",
		ColorPageBackground:       "#020617",
		ColorMainBackground:       "#0b1220",
		ColorSurface:              "#0f172a",
		ColorSurfaceMuted:         "#111827",
		ColorTextOnSolid:          "#f8fafc",
		ColorInfo:                 "#3b82f6",
		ColorInfoBorder:           "#1d4ed8",
		ColorInfoBackground:       "#0b2a4a",
		ColorInfoText:             "#bfdbfe",
		ColorSuccess:              "#22c55e",
		ColorSuccessBorder:        "#15803d",
		ColorSuccessBackground:    "#052e16",
		ColorSuccessText:          "#86efac",
		ColorWarning:              "#f59e0b",
		ColorWarningBorder:        "#b45309",
		ColorWarningBackground:    "#451a03",
		ColorWarningText:          "#fcd34d",
		ColorDanger:               "#ef4444",
		ColorDangerBorder:         "#b91c1c",
		ColorDangerBackground:     "#450a0a",
		ColorDangerText:           "#fca5a5",
		BorderMain:                "1px solid #334155",
		RadiusMain:                "16px",
		MainContentBodyTopSpacing: "24px",
		MSOCompatibility:          MSOCompatibilityModeOn,
		RadiusElement:             "12px",
		RadiusButton:              "8px",
		RadiusPill:                "999px",
		TableLegendSwatchSize:     "11px",
		TableLegendSwatchRadius:   "3px",
		TableLegendSwatchBorder:   "",
	}
}
