package myrtle

// Group is a reusable collection of blocks that can be composed and rendered as one block.
type Group struct {
	blocks []Block
}

// NewGroup creates an empty reusable block group.
// Groups can be composed independently and embedded in columns or grids.
func NewGroup() *Group {
	return &Group{}
}

// Blocks returns a shallow copy of the group's blocks.
// Returning a copy prevents callers from mutating internal group state.
func (group *Group) Blocks() []Block {
	if group == nil || len(group.blocks) == 0 {
		return nil
	}

	return append([]Block(nil), group.blocks...)
}

// Add appends a block to the group.
// Use this for custom or preconstructed block instances.
func (group *Group) Add(block Block) *Group {
	if group == nil {
		return nil
	}
	if blockGroup, ok := block.(*Group); ok && blockGroup == group {
		return group
	}

	group.blocks = append(group.blocks, block)
	return group
}

// AddText appends a text block to the group.
// Text blocks render paragraph-style body copy.
func (group *Group) AddText(text string, options ...TextOption) *Group {
	block := TextBlock{Text: text}
	for _, option := range options {
		option(&block)
	}
	group.Add(block)

	return group
}

// AddHeading appends a heading block to the group.
// Heading blocks introduce and structure content sections.
func (group *Group) AddHeading(text string, options ...HeadingOption) *Group {
	block := HeadingBlock{Text: text, Level: 2}
	for _, option := range options {
		option(&block)
	}
	return group.Add(block)
}

// AddSpacer appends a spacer block to the group.
// Spacer blocks create vertical rhythm between nearby sections.
func (group *Group) AddSpacer(options ...SpacerOption) *Group {
	block := SpacerBlock{Size: 16}
	for _, option := range options {
		option(&block)
	}

	return group.Add(block)
}

// AddList appends a list block to the group.
// List blocks render ordered or unordered bullet content.
func (group *Group) AddList(items []string, ordered bool, options ...ListOption) *Group {
	block := ListBlock{Items: items, Ordered: ordered}
	for _, option := range options {
		option(&block)
	}

	return group.Add(block)
}

// AddKeyValue appends a key-value block to the group.
// Key-value blocks present compact labeled facts and values.
func (group *Group) AddKeyValue(header string, pairs []KeyValuePair, options ...KeyValueOption) *Group {
	block := KeyValueBlock{Header: header, Pairs: pairs}
	for _, option := range options {
		option(&block)
	}

	return group.Add(block)
}

// AddHorizontalBarChart appends a horizontal bar chart block to the group.
// This block compares categories with left-to-right bars.
func (group *Group) AddHorizontalBarChart(header string, items []HorizontalBarChartItem, options ...HorizontalBarChartOption) *Group {
	block := HorizontalBarChartBlock{Header: header, Items: append([]HorizontalBarChartItem(nil), items...)}
	for _, option := range options {
		option(&block)
	}

	return group.Add(block)
}

// AddVerticalBarChart appends a vertical bar chart block to the group.
// This block compares categories with bottom-to-top columns.
func (group *Group) AddVerticalBarChart(axisLabels []string, series []VerticalBarChartSeries, options ...VerticalBarChartOption) *Group {
	block := VerticalBarChartBlock{
		AxisLabels: append([]string(nil), axisLabels...),
		Series:     append([]VerticalBarChartSeries(nil), series...),
	}
	for _, option := range options {
		option(&block)
	}

	return group.Add(block)
}

// AddTimeline appends a timeline block to the group.
// Timeline blocks show ordered milestones or process steps.
func (group *Group) AddTimeline(header string, items []TimelineItem, options ...TimelineOption) *Group {
	block := TimelineBlock{Header: header, CurrentIndex: -1, Items: append([]TimelineItem(nil), items...)}
	for _, option := range options {
		option(&block)
	}

	return group.Add(block)
}

// AddStatsRow appends a stats row block to the group.
// Stats rows present multiple compact KPI values in one line.
func (group *Group) AddStatsRow(header string, stats []StatItem) *Group {
	return group.Add(StatsRowBlock{Header: header, Stats: append([]StatItem(nil), stats...)})
}

// AddBadge appends a badge block to the group.
// Badges highlight short status labels with visual tone.
func (group *Group) AddBadge(tone Tone, text string) *Group {
	return group.Add(BadgeBlock{Tone: tone, Text: text})
}

// AddSummaryCard appends a summary card block to the group.
// Summary cards combine a title, message, and optional footer note.
func (group *Group) AddSummaryCard(title, body, footer string, options ...SummaryCardOption) *Group {
	block := SummaryCardBlock{Title: title, Body: body, Footer: footer, Tone: ToneDefault}
	for _, option := range options {
		option(&block)
	}

	return group.Add(block)
}

// AddAttachment appends an attachment block to the group.
// Attachment blocks describe downloadable files with metadata and CTA.
func (group *Group) AddAttachment(filename, meta, url, cta string, options ...AttachmentOption) *Group {
	block := AttachmentBlock{Filename: filename, Meta: meta, URL: url, CTA: cta}
	for _, option := range options {
		option(&block)
	}

	return group.Add(block)
}

// AddQuote appends a quote block to the group.
// Quote blocks emphasize testimonial or attribution-style text.
func (group *Group) AddQuote(text, author string, options ...QuoteOption) *Group {
	block := QuoteBlock{Text: text, Author: author}
	for _, option := range options {
		option(&block)
	}

	return group.Add(block)
}

// AddCallout appends a callout block to the group.
// Callout blocks surface important notices with semantic styling.
func (group *Group) AddCallout(tone Tone, title, body string, options ...CalloutOption) *Group {
	block := CalloutBlock{Tone: tone, Title: title, Body: body}
	for _, option := range options {
		option(&block)
	}

	return group.Add(block)
}

// AddMessage appends a single message block to the group.
// Message blocks render conversational items in a digest/thread style.
func (group *Group) AddMessage(message MessageBlock, options ...MessageOption) *Group {
	block := message
	for _, option := range options {
		option(&block)
	}

	return group.Add(block)
}

// AddMessageDigest appends a digest block containing multiple messages.
// Digest blocks group multiple messages under a shared header/footer.
func (group *Group) AddMessageDigest(messages []MessageBlock, options ...MessageDigestOption) *Group {
	block := MessageDigestBlock{Messages: append([]MessageBlock(nil), messages...), EmptyText: "No messages"}
	for _, option := range options {
		option(&block)
	}

	return group.Add(block)
}

// AddLegal appends a legal/footer-compliance block to the group.
// Legal blocks include company address and subscription management links.
func (group *Group) AddLegal(companyName, address, manageURL, unsubscribeURL string) *Group {
	return group.Add(LegalBlock{
		CompanyName:    companyName,
		Address:        address,
		ManageURL:      manageURL,
		UnsubscribeURL: unsubscribeURL,
	})
}

// AddButton appends a button block to the group.
// Button blocks render a primary call-to-action link.
func (group *Group) AddButton(label, url string, options ...ButtonOption) *Group {
	block := ButtonBlock{Label: label, URL: url}
	for _, option := range options {
		option(&block)
	}

	return group.Add(block)
}

// AddDivider appends a divider block to the group.
// Divider blocks separate sections with a horizontal rule or label.
func (group *Group) AddDivider(options ...DividerOption) *Group {
	block := DividerBlock{}
	for _, option := range options {
		option(&block)
	}

	return group.Add(block)
}

// AddImage appends an image block to the group.
// Image blocks render visual media with alignment and corner controls.
func (group *Group) AddImage(src, alt string, opts ...ImageOption) *Group {
	ib := ImageBlock{Src: src, Alt: alt}
	for _, opt := range opts {
		opt(&ib)
	}

	return group.Add(ib)
}

// AddTable appends a table block to the group.
// Table blocks present structured rows and columns of data.
func (group *Group) AddTable(columns []string, rows [][]string, options ...TableOption) *Group {
	block := TableBlock{Columns: columns, Rows: rows}
	for _, option := range options {
		option(&block)
	}

	return group.Add(block)
}

// AddVerificationCode appends a verification code block to the group.
// Verification code blocks highlight short one-time passcodes.
func (group *Group) AddVerificationCode(label, code string, options ...VerificationCodeOption) *Group {
	block := VerificationCodeBlock{Label: label, Value: code}
	for _, option := range options {
		option(&block)
	}

	return group.Add(block)
}

// AddFreeMarkdown appends a free-markdown block to the group.
// Free markdown blocks allow direct authoring of rich text snippets.
func (group *Group) AddFreeMarkdown(markdown string) *Group {
	return group.Add(FreeMarkdownBlock{Markdown: markdown})
}

// AddColumns appends a two-column layout block to the builder.
// Columns blocks render two side-by-side groups with configurable widths.
func (builder *Builder) AddColumns(leftGroup, rightGroup *Group, options ...ColumnsOption) *Builder {
	leftBlocks := []Block(nil)
	rightBlocks := []Block(nil)
	if leftGroup != nil {
		leftBlocks = leftGroup.Blocks()
	}
	if rightGroup != nil {
		rightBlocks = rightGroup.Blocks()
	}

	block := ColumnsBlock{
		Left:          leftBlocks,
		Right:         rightBlocks,
		LeftWidth:     50,
		RightWidth:    50,
		Gap:           16,
		VerticalAlign: ColumnsVerticalAlignTop,
	}

	for _, option := range options {
		option(&block)
	}

	if block.LeftWidth <= 0 || block.RightWidth <= 0 || block.LeftWidth+block.RightWidth != 100 {
		block.LeftWidth = 50
		block.RightWidth = 50
	}

	return builder.Add(block)
}

// AddButtonGroup appends a grouped button block to the group.
// Button groups place multiple CTAs in one aligned row or stack.
func (group *Group) AddButtonGroup(buttons []ButtonGroupButton, options ...ButtonGroupOption) *Group {
	block := ButtonGroupBlock{Buttons: append([]ButtonGroupButton(nil), buttons...), Gap: 8}
	for _, option := range options {
		option(&block)
	}

	return group.Add(block)
}

// AddSparkline appends a sparkline block to the group.
// Sparkline blocks summarize a short trend inline with key values.
func (group *Group) AddSparkline(header, label, value string, points []int, options ...SparklineOption) *Group {
	block := SparklineBlock{Header: header, Label: label, Value: value, Points: append([]int(nil), points...)}
	for _, option := range options {
		option(&block)
	}

	return group.Add(block)
}

// AddStackedBar appends a stacked bar block to the group.
// Stacked bars show part-to-whole composition per row.
func (group *Group) AddStackedBar(header string, rows []StackedBarRow, options ...StackedBarOption) *Group {
	block := StackedBarBlock{Header: header, Rows: append([]StackedBarRow(nil), rows...)}
	for _, option := range options {
		option(&block)
	}

	return group.Add(block)
}

// AddProgress appends a progress block to the group.
// Progress blocks communicate completion toward one or more goals.
func (group *Group) AddProgress(header string, items []ProgressItem, options ...ProgressOption) *Group {
	block := ProgressBlock{Header: header, Items: append([]ProgressItem(nil), items...)}
	for _, option := range options {
		option(&block)
	}

	return group.Add(block)
}

// AddDistribution appends a distribution block to the group.
// Distribution blocks visualize bucketed values across ranges.
func (group *Group) AddDistribution(header string, buckets []DistributionBucket, options ...DistributionOption) *Group {
	block := DistributionBlock{Header: header, Buckets: append([]DistributionBucket(nil), buckets...)}
	for _, option := range options {
		option(&block)
	}

	return group.Add(block)
}

// AddTiles appends a tiles block to the group.
// Tiles blocks show small metric cards in a compact grid.
func (group *Group) AddTiles(entries []TileEntry, options ...TilesOption) *Group {
	block := TilesBlock{Entries: append([]TileEntry(nil), entries...)}
	for _, option := range options {
		option(&block)
	}

	return group.Add(block)
}

// AddPanel appends a panel block wrapping optional content to the group.
// Panels provide a bordered container to group related content.
func (group *Group) AddPanel(content Block, options ...PanelOption) *Group {
	blocks := []Block{}
	if content != nil {
		blocks = append(blocks, content)
	}

	block := PanelBlock{Blocks: blocks, Border: true, Padding: 16}
	for _, option := range options {
		option(&block)
	}

	return group.Add(block)
}

// AddGrid appends a grid block to the group.
// Grid blocks lay out heterogeneous items across multiple columns.
func (group *Group) AddGrid(items []GridItem, options ...GridOption) *Group {
	block := GridBlock{Items: append([]GridItem(nil), items...), Columns: 2, Gap: 12}
	for _, option := range options {
		option(&block)
	}

	return group.Add(block)
}

// AddGridGroups appends groups as a grid block to the group.
// Each group becomes one grid cell with its own nested content.
func (group *Group) AddGridGroups(groups []*Group, options ...GridOption) *Group {
	items := make([]GridItem, 0, len(groups))
	for _, nested := range groups {
		if nested == nil {
			continue
		}
		items = append(items, GridItemGroup(nested))
	}

	return group.AddGrid(items, options...)
}

// AddCardList appends a card list block to the group.
// Card list blocks render repeated card entries in columns.
func (group *Group) AddCardList(cards []CardItem, options ...CardListOption) *Group {
	block := CardListBlock{Cards: append([]CardItem(nil), cards...), Columns: 2, Gap: 12, Border: true}
	for _, option := range options {
		option(&block)
	}

	return group.Add(block)
}

// AddHero appends a hero block to the group.
// Hero blocks present high-impact title, body, and optional CTA.
func (group *Group) AddHero(title, body, ctaLabel, ctaURL string, options ...HeroOption) *Group {
	block := HeroBlock{
		Title:    title,
		Body:     body,
		CTALabel: ctaLabel,
		CTAURL:   ctaURL,
		Tone:     ToneDefault,
	}
	for _, option := range options {
		option(&block)
	}

	return group.Add(block)
}

// AddFooterLinks appends a footer links block to the group.
// Footer links blocks provide secondary navigation and policy links.
func (group *Group) AddFooterLinks(links []FooterLink, note string) *Group {
	return group.Add(FooterLinksBlock{Links: append([]FooterLink(nil), links...), Note: note})
}

// AddPriceSummary appends a price summary block to the group.
// Price summary blocks itemize charges and present an order total.
func (group *Group) AddPriceSummary(header string, items []PriceLine, totalLabel, totalValue string, options ...PriceSummaryOption) *Group {
	block := PriceSummaryBlock{Header: header, Items: append([]PriceLine(nil), items...), TotalLabel: totalLabel, TotalValue: totalValue}
	for _, option := range options {
		option(&block)
	}

	return group.Add(block)
}

// AddEmptyState appends an empty-state block to the group.
// Empty state blocks explain missing data and suggest next actions.
func (group *Group) AddEmptyState(title, body, actionLabel, actionURL string, options ...EmptyStateOption) *Group {
	block := EmptyStateBlock{Title: title, Body: body, ActionLabel: actionLabel, ActionURL: actionURL}
	block.Tone = ToneDefault
	for _, option := range options {
		option(&block)
	}

	return group.Add(block)
}

// AddColumns appends a two-column layout block to the group.
// Columns blocks render two side-by-side groups with configurable widths.
func (group *Group) AddColumns(leftGroup, rightGroup *Group, options ...ColumnsOption) *Group {
	leftBlocks := []Block(nil)
	rightBlocks := []Block(nil)
	if leftGroup != nil {
		leftBlocks = leftGroup.Blocks()
	}
	if rightGroup != nil {
		rightBlocks = rightGroup.Blocks()
	}

	block := ColumnsBlock{
		Left:          leftBlocks,
		Right:         rightBlocks,
		LeftWidth:     50,
		RightWidth:    50,
		Gap:           16,
		VerticalAlign: ColumnsVerticalAlignTop,
	}

	for _, option := range options {
		option(&block)
	}

	if block.LeftWidth <= 0 || block.RightWidth <= 0 || block.LeftWidth+block.RightWidth != 100 {
		block.LeftWidth = 50
		block.RightWidth = 50
	}

	return group.Add(block)
}
