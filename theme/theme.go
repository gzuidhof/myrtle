package theme

type BlockKind string

const (
	BlockKindText             BlockKind = "text"
	BlockKindHeading          BlockKind = "heading"
	BlockKindSpacer           BlockKind = "spacer"
	BlockKindList             BlockKind = "list"
	BlockKindKeyValue         BlockKind = "key_value"
	BlockKindBarChart         BlockKind = "bar_chart"
	BlockKindSparkline        BlockKind = "sparkline"
	BlockKindStackedBar       BlockKind = "stacked_bar"
	BlockKindProgress         BlockKind = "progress"
	BlockKindDistribution     BlockKind = "distribution"
	BlockKindTimeline         BlockKind = "timeline"
	BlockKindStatsRow         BlockKind = "stats_row"
	BlockKindBadge            BlockKind = "badge"
	BlockKindSummaryCard      BlockKind = "summary_card"
	BlockKindAttachment       BlockKind = "attachment"
	BlockKindHero             BlockKind = "hero"
	BlockKindFooterLinks      BlockKind = "footer_links"
	BlockKindPriceSummary     BlockKind = "price_summary"
	BlockKindEmptyState       BlockKind = "empty_state"
	BlockKindQuote            BlockKind = "quote"
	BlockKindCallout          BlockKind = "callout"
	BlockKindMessage          BlockKind = "message"
	BlockKindMessageDigest    BlockKind = "message_digest"
	BlockKindLegal            BlockKind = "legal"
	BlockKindColumns          BlockKind = "columns"
	BlockKindSection          BlockKind = "section"
	BlockKindGrid             BlockKind = "grid"
	BlockKindCardList         BlockKind = "card_list"
	BlockKindButton           BlockKind = "button"
	BlockKindButtonGroup      BlockKind = "button_group"
	BlockKindDivider          BlockKind = "divider"
	BlockKindImage            BlockKind = "image"
	BlockKindTable            BlockKind = "table"
	BlockKindVerificationCode BlockKind = "verification_code"
	BlockKindTiles            BlockKind = "tiles"
	BlockKindFreeMarkdown     BlockKind = "free_markdown"
)

var AllBlockKinds = []BlockKind{
	BlockKindText,
	BlockKindHeading,
	BlockKindSpacer,
	BlockKindList,
	BlockKindKeyValue,
	BlockKindBarChart,
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

// Values are shared render values available to all blocks and layouts.
type Values struct {
	ProductName string
	ProductLink string
	LogoURL     string
	LogoAlt     string
	Styles      Styles
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
	// BorderMain is the CSS border value applied to the main content container.
	BorderMain string
	// RadiusMain is the CSS border-radius value for the main content container.
	RadiusMain string
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
	Title           string
	ShowTitle       bool
	ProductName     string
	ShowProductName bool
	ProductLink     string
	LogoURL         string
	LogoAlt         string
	LogoCentered    bool
	Alignment       string
}

// EmailView is the complete render input for an HTML email body.
type EmailView struct {
	Header    *HeaderView
	Preheader string
	Values    Values
	Blocks    []string
}

// TextView is the complete render input for wrapped markdown/text output.
type TextView struct {
	Header    *HeaderView
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

type MarkdownWrapper interface {
	WrapMarkdown(view TextView) (string, error)
}
