package myrtle

type HeadingOption func(*HeadingBlock)

type ColumnsOption func(*ColumnsBlock)

type ColumnBuilder struct {
	blocks []Block
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

func (builder *ColumnBuilder) Add(block Block) *ColumnBuilder {
	builder.blocks = append(builder.blocks, block)
	return builder
}

func (builder *ColumnBuilder) AddText(first string, more ...string) *ColumnBuilder {
	builder.Add(TextBlock{Text: first})
	for _, value := range more {
		builder.Add(TextBlock{Text: value})
	}

	return builder
}

func (builder *ColumnBuilder) AddHeading(text string, options ...HeadingOption) *ColumnBuilder {
	block := HeadingBlock{Text: text, Level: 2}
	for _, option := range options {
		option(&block)
	}
	return builder.Add(block)
}

func (builder *ColumnBuilder) AddSpacer(size int) *ColumnBuilder {
	return builder.Add(SpacerBlock{Size: size})
}

func (builder *ColumnBuilder) AddList(items []string, ordered bool) *ColumnBuilder {
	return builder.Add(ListBlock{Items: items, Ordered: ordered})
}

func (builder *ColumnBuilder) AddKeyValue(header string, pairs []KeyValuePair) *ColumnBuilder {
	return builder.Add(KeyValueBlock{Header: header, Pairs: pairs})
}

func (builder *ColumnBuilder) AddBarChart(header string, items []BarChartItem, options ...BarChartOption) *ColumnBuilder {
	block := BarChartBlock{Header: header, Items: append([]BarChartItem(nil), items...)}
	for _, option := range options {
		option(&block)
	}

	return builder.Add(block)
}

func (builder *ColumnBuilder) AddTimeline(header string, items []TimelineItem) *ColumnBuilder {
	return builder.Add(TimelineBlock{Header: header, Items: append([]TimelineItem(nil), items...)})
}

func (builder *ColumnBuilder) AddStatsRow(header string, stats []StatItem) *ColumnBuilder {
	return builder.Add(StatsRowBlock{Header: header, Stats: append([]StatItem(nil), stats...)})
}

func (builder *ColumnBuilder) AddBadge(tone BadgeTone, text string) *ColumnBuilder {
	return builder.Add(BadgeBlock{Tone: tone, Text: text})
}

func (builder *ColumnBuilder) AddSummaryCard(title, body, footer string) *ColumnBuilder {
	return builder.Add(SummaryCardBlock{Title: title, Body: body, Footer: footer})
}

func (builder *ColumnBuilder) AddAttachment(filename, meta, url, cta string) *ColumnBuilder {
	return builder.Add(AttachmentBlock{Filename: filename, Meta: meta, URL: url, CTA: cta})
}

func (builder *ColumnBuilder) AddQuote(text, author string) *ColumnBuilder {
	return builder.Add(QuoteBlock{Text: text, Author: author})
}

func (builder *ColumnBuilder) AddCallout(calloutType CalloutType, title, body string, options ...CalloutOption) *ColumnBuilder {
	block := CalloutBlock{Type: calloutType, Title: title, Body: body}
	for _, option := range options {
		option(&block)
	}

	return builder.Add(block)
}

func (builder *ColumnBuilder) AddLegal(companyName, address, manageURL, unsubscribeURL string) *ColumnBuilder {
	return builder.Add(LegalBlock{
		CompanyName:    companyName,
		Address:        address,
		ManageURL:      manageURL,
		UnsubscribeURL: unsubscribeURL,
	})
}

func (builder *ColumnBuilder) AddButton(label, url string, options ...ButtonOption) *ColumnBuilder {
	block := ButtonBlock{Label: label, URL: url}
	for _, option := range options {
		option(&block)
	}

	return builder.Add(block)
}

func (builder *ColumnBuilder) AddDivider() *ColumnBuilder {
	return builder.Add(DividerBlock{})
}

func (builder *ColumnBuilder) AddImage(src, alt string) *ColumnBuilder {
	return builder.Add(ImageBlock{Src: src, Alt: alt})
}

func (builder *ColumnBuilder) AddTable(header string, columns []string, rows [][]string) *ColumnBuilder {
	return builder.Add(TableBlock{Header: header, Columns: columns, Rows: rows})
}

func (builder *ColumnBuilder) AddAction(instructions, buttonLabel, buttonURL string) *ColumnBuilder {
	return builder.Add(ActionBlock{Instructions: instructions, ButtonLabel: buttonLabel, ButtonURL: buttonURL})
}

func (builder *ColumnBuilder) AddCode(code string) *ColumnBuilder {
	return builder.Add(CodeBlock{Code: code})
}

func (builder *ColumnBuilder) AddFreeMarkdown(markdown string) *ColumnBuilder {
	return builder.Add(FreeMarkdownBlock{Markdown: markdown})
}

func (builder *Builder) AddColumns(left func(*ColumnBuilder), right func(*ColumnBuilder), options ...ColumnsOption) *Builder {
	leftColumn := &ColumnBuilder{}
	rightColumn := &ColumnBuilder{}

	if left != nil {
		left(leftColumn)
	}
	if right != nil {
		right(rightColumn)
	}

	block := ColumnsBlock{
		Left:       append([]Block(nil), leftColumn.blocks...),
		Right:      append([]Block(nil), rightColumn.blocks...),
		LeftWidth:  50,
		RightWidth: 50,
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
