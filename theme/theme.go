package theme

type BlockKind string

const (
	BlockKindText         BlockKind = "text"
	BlockKindHeading      BlockKind = "heading"
	BlockKindSpacer       BlockKind = "spacer"
	BlockKindList         BlockKind = "list"
	BlockKindKeyValue     BlockKind = "key_value"
	BlockKindBarChart     BlockKind = "bar_chart"
	BlockKindSparkline    BlockKind = "sparkline"
	BlockKindStackedBar   BlockKind = "stacked_bar"
	BlockKindProgress     BlockKind = "progress"
	BlockKindDistribution BlockKind = "distribution"
	BlockKindTimeline     BlockKind = "timeline"
	BlockKindStatsRow     BlockKind = "stats_row"
	BlockKindBadge        BlockKind = "badge"
	BlockKindSummaryCard  BlockKind = "summary_card"
	BlockKindAttachment   BlockKind = "attachment"
	BlockKindHero         BlockKind = "hero"
	BlockKindFooterLinks  BlockKind = "footer_links"
	BlockKindPriceSummary BlockKind = "price_summary"
	BlockKindEmptyState   BlockKind = "empty_state"
	BlockKindQuote        BlockKind = "quote"
	BlockKindCallout      BlockKind = "callout"
	BlockKindLegal        BlockKind = "legal"
	BlockKindColumns      BlockKind = "columns"
	BlockKindButton       BlockKind = "button"
	BlockKindButtonGroup  BlockKind = "button_group"
	BlockKindDivider      BlockKind = "divider"
	BlockKindImage        BlockKind = "image"
	BlockKindTable        BlockKind = "table"
	BlockKindAction       BlockKind = "action"
	BlockKindCode         BlockKind = "code"
	BlockKindFreeMarkdown BlockKind = "free_markdown"
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
	BlockKindLegal,
	BlockKindColumns,
	BlockKindButton,
	BlockKindButtonGroup,
	BlockKindDivider,
	BlockKindImage,
	BlockKindTable,
	BlockKindAction,
	BlockKindCode,
	BlockKindFreeMarkdown,
}

type BlockView struct {
	Kind   BlockKind
	Data   any
	Values Values
}

type Values struct {
	ProductName string
	ProductLink string
	LogoURL     string
	LogoAlt     string
	Styles      Styles
}

type Styles struct {
	PrimaryColor        string
	TextColor           string
	MutedTextColor      string
	BorderColor         string
	CodeBackgroundColor string
}

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

type EmailView struct {
	Header    *HeaderView
	Preheader string
	Values    Values
	Blocks    []string
}

type TextView struct {
	Header    *HeaderView
	Preheader string
	Values    Values
	Body      string
}

type Theme interface {
	Name() string
	RenderHTML(view EmailView) (string, error)
	RenderBlockHTML(block BlockView) (string, bool, error)
}

type MarkdownWrapper interface {
	WrapMarkdown(view TextView) (string, error)
}
