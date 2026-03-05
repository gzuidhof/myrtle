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

// HeadingLevel sets the heading level on a HeadingBlock.
func HeadingLevel(value int) HeadingOption {
	return func(block *HeadingBlock) {
		if value > 0 {
			block.Level = value
		}
	}
}

// ColumnsWidths sets relative left/right column widths using percentage normalization.
func ColumnsWidths(leftWidth, rightWidth int) ColumnsOption {
	return func(block *ColumnsBlock) {
		if leftWidth <= 0 || rightWidth <= 0 {
			return
		}
		total := leftWidth + rightWidth
		if total <= 0 {
			return
		}

		block.LeftWidth = (leftWidth * 100) / total
		block.RightWidth = 100 - block.LeftWidth
	}
}

// ColumnsGap sets the horizontal gap between columns.
func ColumnsGap(value int) ColumnsOption {
	return func(block *ColumnsBlock) {
		if value < 0 {
			return
		}
		block.Gap = value
	}
}

// ColumnsAlign sets vertical alignment for both columns.
func ColumnsAlign(value ColumnsVerticalAlign) ColumnsOption {
	return func(block *ColumnsBlock) {
		block.VerticalAlign = value
	}
}

// ColumnsInsetMode sets the layout inset mode for a columns block.
func ColumnsInsetMode(value InsetMode) ColumnsOption {
	return func(block *ColumnsBlock) {
		block.InsetMode = value
	}
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
func (group *Group) AddList(items []string, ordered bool) *Group {
	return group.Add(ListBlock{Items: items, Ordered: ordered})
}

// AddKeyValue appends a key-value block to the group.
// Key-value blocks present compact labeled facts and values.
func (group *Group) AddKeyValue(header string, pairs []KeyValuePair) *Group {
	return group.Add(KeyValueBlock{Header: header, Pairs: pairs})
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
func (group *Group) AddSummaryCard(title, body, footer string) *Group {
	return group.Add(SummaryCardBlock{Title: title, Body: body, Footer: footer})
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
func (group *Group) AddQuote(text, author string) *Group {
	return group.Add(QuoteBlock{Text: text, Author: author})
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
