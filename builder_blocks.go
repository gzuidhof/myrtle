package myrtle

import "github.com/gzuidhof/myrtle/theme"

type ButtonOption func(*ButtonBlock)

func ButtonStyle(variant ButtonVariant) ButtonOption {
	return func(block *ButtonBlock) {
		block.Variant = variant
	}
}

func ButtonFullWidth(value bool) ButtonOption {
	return func(block *ButtonBlock) {
		block.FullWidth = value
	}
}

func ButtonAlign(value ButtonAlignment) ButtonOption {
	return func(block *ButtonBlock) {
		block.Alignment = value
	}
}

type ButtonGroupOption func(*ButtonGroupBlock)

func ButtonGroupAlign(value ButtonAlignment) ButtonGroupOption {
	return func(block *ButtonGroupBlock) {
		block.Alignment = value
	}
}

func ButtonGroupJoined(value bool) ButtonGroupOption {
	return func(block *ButtonGroupBlock) {
		block.Joined = value
	}
}

type CalloutOption func(*CalloutBlock)

func CalloutStyle(variant CalloutVariant) CalloutOption {
	return func(block *CalloutBlock) {
		block.Variant = variant
	}
}

func CalloutLink(label, url string) CalloutOption {
	return func(block *CalloutBlock) {
		block.LinkLabel = label
		block.LinkURL = url
	}
}

type TimelineOption func(*TimelineBlock)

func TimelineCurrentIndex(value int) TimelineOption {
	return func(block *TimelineBlock) {
		block.HasCurrentIndex = true
		block.CurrentIndex = value
	}
}

func TimelineAggregateHeader(value string) TimelineOption {
	return func(block *TimelineBlock) {
		block.AggregateHeader = value
	}
}

type StackedBarOption func(*StackedBarBlock)

func StackedBarTotal(label, value string) StackedBarOption {
	return func(block *StackedBarBlock) {
		block.TotalLabel = label
		block.TotalValue = value
	}
}

type TableOption func(*TableBlock)

type BarChartOption func(*BarChartBlock)

type SparklineOption func(*SparklineBlock)

func SparklineDelta(value string) SparklineOption {
	return func(block *SparklineBlock) {
		block.Delta = value
	}
}

func SparklineDeltaSemantic(value StatDeltaSemantic) SparklineOption {
	return func(block *SparklineBlock) {
		block.DeltaSemantic = value
	}
}

func BarChartThickness(value int) BarChartOption {
	return func(block *BarChartBlock) {
		block.Thickness = value
	}
}

func BarChartTransparentBackground(value bool) BarChartOption {
	return func(block *BarChartBlock) {
		block.TransparentBackground = value
	}
}

func TableZebraRows(value bool) TableOption {
	return func(block *TableBlock) {
		block.ZebraRows = value
	}
}

func TableCompact(value bool) TableOption {
	return func(block *TableBlock) {
		block.Compact = value
	}
}

func TableRightAlignNumericColumns(value bool) TableOption {
	return func(block *TableBlock) {
		block.RightAlignNumericColumns = value
	}
}

func TableEmphasizeTotalRow(value bool) TableOption {
	return func(block *TableBlock) {
		block.EmphasizeTotalRow = value
	}
}

func TableColumnAlignments(value map[int]TableColumnAlignment) TableOption {
	return func(block *TableBlock) {
		if len(value) == 0 {
			block.ColumnAlignments = nil
			return
		}

		alignments := make(map[int]TableColumnAlignment, len(value))
		for index, alignment := range value {
			alignments[index] = alignment
		}

		block.ColumnAlignments = alignments
	}
}

func (builder *Builder) Add(block Block) *Builder {
	builder.blocks = append(builder.blocks, block)
	return builder
}

func (builder *Builder) AddText(first string, more ...string) *Builder {
	builder.Add(TextBlock{Text: first})
	for _, value := range more {
		builder.Add(TextBlock{Text: value})
	}

	return builder
}

func (builder *Builder) AddHeading(text string, options ...HeadingOption) *Builder {
	block := HeadingBlock{Text: text, Level: 2}
	for _, option := range options {
		option(&block)
	}
	return builder.Add(block)
}

func (builder *Builder) AddSpacer(size int) *Builder {
	return builder.Add(SpacerBlock{Size: size})
}

func (builder *Builder) AddList(items []string, ordered bool) *Builder {
	return builder.Add(ListBlock{Items: items, Ordered: ordered})
}

func (builder *Builder) AddKeyValue(header string, pairs []KeyValuePair) *Builder {
	return builder.Add(KeyValueBlock{Header: header, Pairs: pairs})
}

func (builder *Builder) AddBarChart(header string, items []BarChartItem, options ...BarChartOption) *Builder {
	block := BarChartBlock{Header: header, Items: append([]BarChartItem(nil), items...)}
	for _, option := range options {
		option(&block)
	}

	return builder.Add(block)
}

func (builder *Builder) AddSparkline(header, label, value string, points []int, options ...SparklineOption) *Builder {
	block := SparklineBlock{Header: header, Label: label, Value: value, Points: append([]int(nil), points...)}
	for _, option := range options {
		option(&block)
	}

	return builder.Add(block)
}

func (builder *Builder) AddStackedBar(header string, rows []StackedBarRow, options ...StackedBarOption) *Builder {
	block := StackedBarBlock{Header: header, Rows: append([]StackedBarRow(nil), rows...)}
	for _, option := range options {
		option(&block)
	}

	return builder.Add(block)
}

func (builder *Builder) AddProgress(header string, items []ProgressItem) *Builder {
	return builder.Add(ProgressBlock{Header: header, Items: append([]ProgressItem(nil), items...)})
}

func (builder *Builder) AddDistribution(header string, buckets []DistributionBucket) *Builder {
	return builder.Add(DistributionBlock{Header: header, Buckets: append([]DistributionBucket(nil), buckets...)})
}

func (builder *Builder) AddTimeline(header string, items []TimelineItem, options ...TimelineOption) *Builder {
	block := TimelineBlock{Header: header, CurrentIndex: -1, Items: append([]TimelineItem(nil), items...)}
	for _, option := range options {
		option(&block)
	}

	return builder.Add(block)
}

func (builder *Builder) AddStatsRow(header string, stats []StatItem) *Builder {
	return builder.Add(StatsRowBlock{Header: header, Stats: append([]StatItem(nil), stats...)})
}

func (builder *Builder) AddBadge(tone BadgeTone, text string) *Builder {
	return builder.Add(BadgeBlock{Tone: tone, Text: text})
}

func (builder *Builder) AddSummaryCard(title, body, footer string) *Builder {
	return builder.Add(SummaryCardBlock{Title: title, Body: body, Footer: footer})
}

func (builder *Builder) AddAttachment(filename, meta, url, cta string) *Builder {
	return builder.Add(AttachmentBlock{Filename: filename, Meta: meta, URL: url, CTA: cta})
}

func (builder *Builder) AddQuote(text, author string) *Builder {
	return builder.Add(QuoteBlock{Text: text, Author: author})
}

func (builder *Builder) AddCallout(calloutType CalloutType, title, body string, options ...CalloutOption) *Builder {
	block := CalloutBlock{Type: calloutType, Title: title, Body: body}
	for _, option := range options {
		option(&block)
	}

	return builder.Add(block)
}

func (builder *Builder) AddLegal(companyName, address, manageURL, unsubscribeURL string) *Builder {
	return builder.Add(LegalBlock{
		CompanyName:    companyName,
		Address:        address,
		ManageURL:      manageURL,
		UnsubscribeURL: unsubscribeURL,
	})
}

func (builder *Builder) AddButton(label, url string, options ...ButtonOption) *Builder {
	block := ButtonBlock{Label: label, URL: url}
	for _, option := range options {
		option(&block)
	}

	return builder.Add(block)
}

func (builder *Builder) AddButtonGroup(buttons []ButtonGroupButton, options ...ButtonGroupOption) *Builder {
	block := ButtonGroupBlock{Buttons: append([]ButtonGroupButton(nil), buttons...)}
	for _, option := range options {
		option(&block)
	}

	return builder.Add(block)
}

func (builder *Builder) AddDivider() *Builder {
	return builder.Add(DividerBlock{})
}

func (builder *Builder) AddImage(src, alt string) *Builder {
	return builder.Add(ImageBlock{Src: src, Alt: alt})
}

func (builder *Builder) AddTable(header string, columns []string, rows [][]string, options ...TableOption) *Builder {
	block := TableBlock{Header: header, Columns: columns, Rows: rows}
	for _, option := range options {
		option(&block)
	}

	return builder.Add(block)
}

func (builder *Builder) AddAction(instructions, buttonLabel, buttonURL string) *Builder {
	return builder.Add(ActionBlock{Instructions: instructions, ButtonLabel: buttonLabel, ButtonURL: buttonURL})
}

func (builder *Builder) AddCode(code string) *Builder {
	return builder.Add(CodeBlock{Code: code})
}

func (builder *Builder) AddFreeMarkdown(markdown string) *Builder {
	return builder.Add(FreeMarkdownBlock{Markdown: markdown})
}

func (builder *Builder) AddHero(title, body, ctaLabel, ctaURL string) *Builder {
	return builder.Add(HeroBlock{
		Title:    title,
		Body:     body,
		CTALabel: ctaLabel,
		CTAURL:   ctaURL,
	})
}

func (builder *Builder) AddFooterLinks(links []FooterLink, note string) *Builder {
	return builder.Add(FooterLinksBlock{Links: append([]FooterLink(nil), links...), Note: note})
}

func (builder *Builder) AddPriceSummary(header string, items []PriceLine, totalLabel, totalValue string) *Builder {
	return builder.Add(PriceSummaryBlock{Header: header, Items: append([]PriceLine(nil), items...), TotalLabel: totalLabel, TotalValue: totalValue})
}

func (builder *Builder) AddEmptyState(title, body, actionLabel, actionURL string) *Builder {
	return builder.Add(EmptyStateBlock{Title: title, Body: body, ActionLabel: actionLabel, ActionURL: actionURL})
}

func (builder *Builder) AddCustom(kind string, data any) (*Builder, error) {
	customBlock, err := builder.registry.create(theme.BlockKind(kind), data)
	if err != nil {
		return builder, err
	}

	return builder.Add(customBlock), nil
}

func (builder *Builder) Build() *Email {
	result := &Email{
		header:    cloneHeader(builder.header),
		preheader: builder.preheader,
		values:    normalizeValues(builder.values),
		blocks:    append([]Block(nil), builder.blocks...),
		theme:     builder.theme,
	}

	if result.header != nil && result.header.LogoAlt == "" {
		result.header.LogoAlt = result.values.LogoAlt
	}

	return result
}
