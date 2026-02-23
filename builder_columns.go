package myrtle

type HeadingOption func(*HeadingBlock)

type ColumnsOption func(*ColumnsBlock)

// Group is a reusable collection of blocks that can be composed and rendered as one block.
type Group struct {
	blocks []Block
}

type ColumnBuilder = Group

func NewGroup() *Group {
	return &Group{}
}

func (group *Group) Blocks() []Block {
	if group == nil || len(group.blocks) == 0 {
		return nil
	}

	return append([]Block(nil), group.blocks...)
}

func HeadingLevel(value int) HeadingOption {
	return func(block *HeadingBlock) {
		if value > 0 {
			block.Level = value
		}
	}
}

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

func ColumnsGap(value int) ColumnsOption {
	return func(block *ColumnsBlock) {
		if value < 0 {
			return
		}
		block.Gap = value
	}
}

func ColumnsAlign(value ColumnsVerticalAlign) ColumnsOption {
	return func(block *ColumnsBlock) {
		block.VerticalAlign = value
	}
}

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

func (group *Group) AddText(first string, more ...string) *Group {
	group.Add(TextBlock{Text: first})
	for _, value := range more {
		group.Add(TextBlock{Text: value})
	}

	return group
}

func (group *Group) AddHeading(text string, options ...HeadingOption) *Group {
	block := HeadingBlock{Text: text, Level: 2}
	for _, option := range options {
		option(&block)
	}
	return group.Add(block)
}

func (group *Group) AddSpacer(options ...SpacerOption) *Group {
	block := SpacerBlock{Size: 16}
	for _, option := range options {
		option(&block)
	}

	return group.Add(block)
}

func (group *Group) AddList(items []string, ordered bool) *Group {
	return group.Add(ListBlock{Items: items, Ordered: ordered})
}

func (group *Group) AddKeyValue(header string, pairs []KeyValuePair) *Group {
	return group.Add(KeyValueBlock{Header: header, Pairs: pairs})
}

func (group *Group) AddBarChart(header string, items []BarChartItem, options ...BarChartOption) *Group {
	block := BarChartBlock{Header: header, Items: append([]BarChartItem(nil), items...)}
	for _, option := range options {
		option(&block)
	}

	return group.Add(block)
}

func (group *Group) AddTimeline(header string, items []TimelineItem) *Group {
	return group.Add(TimelineBlock{Header: header, Items: append([]TimelineItem(nil), items...)})
}

func (group *Group) AddStatsRow(header string, stats []StatItem) *Group {
	return group.Add(StatsRowBlock{Header: header, Stats: append([]StatItem(nil), stats...)})
}

func (group *Group) AddBadge(tone BadgeTone, text string) *Group {
	return group.Add(BadgeBlock{Tone: tone, Text: text})
}

func (group *Group) AddSummaryCard(title, body, footer string) *Group {
	return group.Add(SummaryCardBlock{Title: title, Body: body, Footer: footer})
}

func (group *Group) AddAttachment(filename, meta, url, cta string) *Group {
	return group.Add(AttachmentBlock{Filename: filename, Meta: meta, URL: url, CTA: cta})
}

func (group *Group) AddQuote(text, author string) *Group {
	return group.Add(QuoteBlock{Text: text, Author: author})
}

func (group *Group) AddCallout(calloutType CalloutType, title, body string, options ...CalloutOption) *Group {
	block := CalloutBlock{Type: calloutType, Title: title, Body: body}
	for _, option := range options {
		option(&block)
	}

	return group.Add(block)
}

func (group *Group) AddMessage(message MessageBlock) *Group {
	return group.Add(message)
}

func (group *Group) AddMessageDigest(messages []MessageBlock, options ...MessageDigestOption) *Group {
	block := MessageDigestBlock{Messages: append([]MessageBlock(nil), messages...), EmptyText: "No messages"}
	for _, option := range options {
		option(&block)
	}

	return group.Add(block)
}

func (group *Group) AddLegal(companyName, address, manageURL, unsubscribeURL string) *Group {
	return group.Add(LegalBlock{
		CompanyName:    companyName,
		Address:        address,
		ManageURL:      manageURL,
		UnsubscribeURL: unsubscribeURL,
	})
}

func (group *Group) AddButton(label, url string, options ...ButtonOption) *Group {
	block := ButtonBlock{Label: label, URL: url}
	for _, option := range options {
		option(&block)
	}

	return group.Add(block)
}

func (group *Group) AddDivider() *Group {
	return group.Add(DividerBlock{})
}

func (group *Group) AddDividerStyled(options ...DividerOption) *Group {
	block := DividerBlock{}
	for _, option := range options {
		option(&block)
	}

	return group.Add(block)
}

func (group *Group) AddImage(src, alt string) *Group {
	return group.Add(ImageBlock{Src: src, Alt: alt})
}

func (group *Group) AddTable(header string, columns []string, rows [][]string) *Group {
	return group.Add(TableBlock{Header: header, Columns: columns, Rows: rows})
}

func (group *Group) AddVerificationCode(label, code string) *Group {
	return group.Add(VerificationCodeBlock{Label: label, Value: code})
}

func (group *Group) AddFreeMarkdown(markdown string) *Group {
	return group.Add(FreeMarkdownBlock{Markdown: markdown})
}

func (builder *Builder) AddColumns(left func(*ColumnBuilder), right func(*ColumnBuilder), options ...ColumnsOption) *Builder {
	leftColumn := &Group{}
	rightColumn := &Group{}

	if left != nil {
		left(leftColumn)
	}
	if right != nil {
		right(rightColumn)
	}

	block := ColumnsBlock{
		Left:          leftColumn.Blocks(),
		Right:         rightColumn.Blocks(),
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

func (builder *Builder) AddColumnsGroups(leftGroup, rightGroup *Group, options ...ColumnsOption) *Builder {
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
