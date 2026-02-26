package myrtle

type ButtonOption func(*ButtonBlock)

func ButtonTone(value ButtonToneValue) ButtonOption {
	return func(block *ButtonBlock) {
		block.Tone = value
	}
}

func ButtonStyle(value ButtonStyleValue) ButtonOption {
	return func(block *ButtonBlock) {
		block.Style = value
	}
}

func ButtonFullWidth(value bool) ButtonOption {
	return func(block *ButtonBlock) {
		block.FullWidth = value
	}
}

func ButtonSize(value ButtonSizeValue) ButtonOption {
	return func(block *ButtonBlock) {
		block.Size = value
	}
}

func ButtonNoWrap(value bool) ButtonOption {
	return func(block *ButtonBlock) {
		block.NoWrap = value
	}
}

func ButtonAlign(value ButtonAlignmentValue) ButtonOption {
	return func(block *ButtonBlock) {
		block.Alignment = value
	}
}

type ButtonGroupOption func(*ButtonGroupBlock)

type SectionOption func(*SectionBlock)

type GridOption func(*GridBlock)

type CardListOption func(*CardListBlock)

type SpacerOption func(*SpacerBlock)

type DividerOption func(*DividerBlock)

func ButtonGroupAlign(value ButtonAlignmentValue) ButtonGroupOption {
	return func(block *ButtonGroupBlock) {
		block.Alignment = value
	}
}

func ButtonGroupJoined(value bool) ButtonGroupOption {
	return func(block *ButtonGroupBlock) {
		block.Joined = value
	}
}

func ButtonGroupStackOnMobile(value bool) ButtonGroupOption {
	return func(block *ButtonGroupBlock) {
		block.StackOnMobile = value
	}
}

func ButtonGroupFullWidthOnMobile(value bool) ButtonGroupOption {
	return func(block *ButtonGroupBlock) {
		block.FullWidthOnMobile = value
	}
}

func ButtonGroupGap(value int) ButtonGroupOption {
	return func(block *ButtonGroupBlock) {
		block.Gap = value
	}
}

func SectionTitle(value string) SectionOption {
	return func(block *SectionBlock) {
		block.Title = value
	}
}

func SectionSubtitle(value string) SectionOption {
	return func(block *SectionBlock) {
		block.Subtitle = value
	}
}

func SectionCategory(value string) SectionOption {
	return func(block *SectionBlock) {
		block.Category = value
	}
}

func SectionBorder(value bool) SectionOption {
	return func(block *SectionBlock) {
		block.Border = value
	}
}

func SectionPadding(value int) SectionOption {
	return func(block *SectionBlock) {
		block.Padding = value
	}
}

func GridColumns(value int) GridOption {
	return func(block *GridBlock) {
		block.Columns = value
	}
}

func GridGap(value int) GridOption {
	return func(block *GridBlock) {
		block.Gap = value
	}
}

func GridBorder(value bool) GridOption {
	return func(block *GridBlock) {
		block.Border = value
	}
}

func CardListColumns(value int) CardListOption {
	return func(block *CardListBlock) {
		block.Columns = value
	}
}

func CardListGap(value int) CardListOption {
	return func(block *CardListBlock) {
		block.Gap = value
	}
}

func CardListBorder(value bool) CardListOption {
	return func(block *CardListBlock) {
		block.Border = value
	}
}

func SpacerSize(value int) SpacerOption {
	return func(block *SpacerBlock) {
		block.Size = value
	}
}

func DividerStyle(value DividerVariant) DividerOption {
	return func(block *DividerBlock) {
		block.Variant = value
	}
}

func DividerThickness(value int) DividerOption {
	return func(block *DividerBlock) {
		block.Thickness = value
	}
}

func DividerInset(value int) DividerOption {
	return func(block *DividerBlock) {
		block.Inset = value
	}
}

func DividerLabel(value string) DividerOption {
	return func(block *DividerBlock) {
		block.Label = value
	}
}

type CalloutOption func(*CalloutBlock)

type MessageDigestOption func(*MessageDigestBlock)

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

func MessageDigestTitle(value string) MessageDigestOption {
	return func(block *MessageDigestBlock) {
		block.Title = value
	}
}

func MessageDigestSubtitle(value string) MessageDigestOption {
	return func(block *MessageDigestBlock) {
		block.Subtitle = value
	}
}

func MessageDigestFooter(value string) MessageDigestOption {
	return func(block *MessageDigestBlock) {
		block.Footer = value
	}
}

func MessageDigestEmptyText(value string) MessageDigestOption {
	return func(block *MessageDigestBlock) {
		block.EmptyText = value
	}
}

func MessageDigestMaxItems(value int) MessageDigestOption {
	return func(block *MessageDigestBlock) {
		block.MaxItems = value
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

func StackedBarTone(value ChartToneValue) StackedBarOption {
	return func(block *StackedBarBlock) {
		block.Tone = value
	}
}

type TableOption func(*TableBlock)

type HorizontalBarChartOption func(*HorizontalBarChartBlock)

type VerticalBarChartOption func(*VerticalBarChartBlock)

type SparklineOption func(*SparklineBlock)

type TilesOption func(*TilesBlock)

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

func SparklineTone(value ChartToneValue) SparklineOption {
	return func(block *SparklineBlock) {
		block.Tone = value
	}
}

func HorizontalBarChartThickness(value int) HorizontalBarChartOption {
	return func(block *HorizontalBarChartBlock) {
		block.Thickness = value
	}
}

func HorizontalBarChartTransparentBackground(value bool) HorizontalBarChartOption {
	return func(block *HorizontalBarChartBlock) {
		block.TransparentBackground = value
	}
}

func HorizontalBarChartTone(value ChartToneValue) HorizontalBarChartOption {
	return func(block *HorizontalBarChartBlock) {
		block.Tone = value
	}
}

func VerticalBarChartHeight(value int) VerticalBarChartOption {
	return func(block *VerticalBarChartBlock) {
		block.Height = value
	}
}

// VerticalBarChartNormalize enables per-column normalization where segment heights
// fill the available positive/negative region in each column.
//
// For mixed-sign datasets (any negative value present), Myrtle automatically
// falls back to magnitude scaling to preserve cross-column comparability.
func VerticalBarChartNormalize(value bool) VerticalBarChartOption {
	return func(block *VerticalBarChartBlock) {
		block.Normalize = value
	}
}

func VerticalBarChartColumnGap(value int) VerticalBarChartOption {
	return func(block *VerticalBarChartBlock) {
		block.HasColumnGap = true
		block.ColumnGap = value
	}
}

// VerticalBarChartCategoryGap is kept as a compatibility alias.
// Prefer VerticalBarChartColumnGap for clarity.
func VerticalBarChartCategoryGap(value int) VerticalBarChartOption {
	return VerticalBarChartColumnGap(value)
}

func VerticalBarChartTransparentBackground(value bool) VerticalBarChartOption {
	return func(block *VerticalBarChartBlock) {
		block.TransparentBackground = value
	}
}

func VerticalBarChartTone(value ChartToneValue) VerticalBarChartOption {
	return func(block *VerticalBarChartBlock) {
		block.Tone = value
	}
}

func VerticalBarChartLegendPlacement(value VerticalBarChartLegendPlacementValue) VerticalBarChartOption {
	return func(block *VerticalBarChartBlock) {
		block.LegendPlacement = normalizedLegendPlacement(value)
	}
}

func VerticalBarChartLegend(items []VerticalBarChartLegendItem) VerticalBarChartOption {
	return func(block *VerticalBarChartBlock) {
		if len(items) == 0 {
			block.Legend = nil
			return
		}

		block.Legend = append([]VerticalBarChartLegendItem(nil), items...)
	}
}

func VerticalBarChartLegendConfigOption(value VerticalBarChartLegendConfig) VerticalBarChartOption {
	return func(block *VerticalBarChartBlock) {
		block.LegendPlacement = normalizedLegendPlacement(value.Placement)
		if len(value.Items) == 0 {
			block.Legend = nil
			return
		}

		block.Legend = append([]VerticalBarChartLegendItem(nil), value.Items...)
	}
}

func VerticalBarChartAxisShowBaseline(value bool) VerticalBarChartOption {
	return func(block *VerticalBarChartBlock) {
		block.Axis.ShowBaseline = value
	}
}

func VerticalBarChartAxisShowYTicks(value bool) VerticalBarChartOption {
	return func(block *VerticalBarChartBlock) {
		block.Axis.ShowYTicks = value
	}
}

func VerticalBarChartAxisTickCount(value int) VerticalBarChartOption {
	return func(block *VerticalBarChartBlock) {
		block.Axis.YTickCount = value
	}
}

func VerticalBarChartAxisShowCategoryLabels(value bool) VerticalBarChartOption {
	return func(block *VerticalBarChartBlock) {
		block.Axis.HasShowCategoryLabels = true
		block.Axis.ShowCategoryLabels = value
	}
}

func VerticalBarChartAxisLabelFormat(value VerticalBarChartAxisLabelFormatValue) VerticalBarChartOption {
	return func(block *VerticalBarChartBlock) {
		block.Axis.LabelFormat = value
	}
}

func VerticalBarChartAxisMin(value float64) VerticalBarChartOption {
	return func(block *VerticalBarChartBlock) {
		block.Axis.HasMin = true
		block.Axis.Min = value
	}
}

func VerticalBarChartAxisMax(value float64) VerticalBarChartOption {
	return func(block *VerticalBarChartBlock) {
		block.Axis.HasMax = true
		block.Axis.Max = value
	}
}

func VerticalBarChartAxisConfig(value VerticalBarChartAxis) VerticalBarChartOption {
	return func(block *VerticalBarChartBlock) {
		block.Axis = value
	}
}

func VerticalBarChartValueLabelsOption(value VerticalBarChartValueLabels) VerticalBarChartOption {
	return func(block *VerticalBarChartBlock) {
		block.ValueLabels = value
	}
}

func VerticalBarChartValueFormatterOption(value VerticalBarChartValueFormatter) VerticalBarChartOption {
	return func(block *VerticalBarChartBlock) {
		block.ValueFormatter = value
	}
}

func TilesColumns(value int) TilesOption {
	return func(block *TilesBlock) {
		block.Columns = value
	}
}

func TilesBorder(value bool) TilesOption {
	return func(block *TilesBlock) {
		block.Border = value
	}
}

func TilesTransparentBackground(value bool) TilesOption {
	return func(block *TilesBlock) {
		block.TransparentBackground = value
	}
}

func TilesAlign(value TileAlignment) TilesOption {
	return func(block *TilesBlock) {
		block.Alignment = value
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

func TableDensity(value TableDensityValue) TableOption {
	return func(block *TableBlock) {
		block.Density = value
	}
}

func TableHeaderTone(value TableHeaderToneValue) TableOption {
	return func(block *TableBlock) {
		block.HeaderTone = value
	}
}

func TableBorderStyle(value TableBorderStyleValue) TableOption {
	return func(block *TableBlock) {
		block.BorderStyle = value
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

func TableColumnAlignments(value map[int]TableColumnAlignmentValue) TableOption {
	return func(block *TableBlock) {
		if len(value) == 0 {
			block.ColumnAlignments = nil
			return
		}

		alignments := make(map[int]TableColumnAlignmentValue, len(value))
		for index, alignment := range value {
			alignments[index] = alignment
		}

		block.ColumnAlignments = alignments
	}
}

func (builder *Builder) Add(block Block) *Builder {
	builder.mu.Lock()
	defer builder.mu.Unlock()

	builder.blocks = append(builder.blocks, block)
	return builder
}

func (builder *Builder) AddText(text string, options ...TextOption) *Builder {
	block := TextBlock{Text: text}
	for _, option := range options {
		option(&block)
	}
	builder.Add(block)

	return builder
}

func (builder *Builder) AddHeading(text string, options ...HeadingOption) *Builder {
	block := HeadingBlock{Text: text, Level: 2}
	for _, option := range options {
		option(&block)
	}
	return builder.Add(block)
}

func (builder *Builder) AddSpacer(options ...SpacerOption) *Builder {
	block := SpacerBlock{Size: 16}
	for _, option := range options {
		option(&block)
	}

	return builder.Add(block)
}

func (builder *Builder) AddList(items []string, ordered bool) *Builder {
	return builder.Add(ListBlock{Items: items, Ordered: ordered})
}

func (builder *Builder) AddKeyValue(header string, pairs []KeyValuePair) *Builder {
	return builder.Add(KeyValueBlock{Header: header, Pairs: pairs})
}

func (builder *Builder) AddHorizontalBarChart(header string, items []HorizontalBarChartItem, options ...HorizontalBarChartOption) *Builder {
	block := HorizontalBarChartBlock{Header: header, Items: append([]HorizontalBarChartItem(nil), items...)}
	for _, option := range options {
		option(&block)
	}

	return builder.Add(block)
}

func (builder *Builder) AddVerticalBarChart(header string, axisLabels []string, series []VerticalBarChartSeries, options ...VerticalBarChartOption) *Builder {
	block := VerticalBarChartBlock{
		Header:     header,
		AxisLabels: append([]string(nil), axisLabels...),
		Series:     append([]VerticalBarChartSeries(nil), series...),
	}
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

func (builder *Builder) AddMessage(message MessageBlock) *Builder {
	return builder.Add(message)
}

func (builder *Builder) AddMessageDigest(messages []MessageBlock, options ...MessageDigestOption) *Builder {
	block := MessageDigestBlock{Messages: append([]MessageBlock(nil), messages...), EmptyText: "No messages"}
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
	block := ButtonGroupBlock{Buttons: append([]ButtonGroupButton(nil), buttons...), Gap: 8}
	for _, option := range options {
		option(&block)
	}

	return builder.Add(block)
}

func (builder *Builder) AddDivider() *Builder {
	return builder.Add(DividerBlock{})
}

func (builder *Builder) AddDividerStyled(options ...DividerOption) *Builder {
	block := DividerBlock{}
	for _, option := range options {
		option(&block)
	}

	return builder.Add(block)
}

// ImageOption configures an ImageBlock.
type ImageOption func(*ImageBlock)

// ImageWidth sets the width (in px) of the image.
func ImageWidth(px int) ImageOption {
	return func(ib *ImageBlock) {
		ib.Width = px
	}
}

// ImageAlign sets the alignment of the image.
func ImageAlign(align ImageAlignment) ImageOption {
	return func(ib *ImageBlock) {
		ib.Align = normalizedImageAlignment(align)
	}
}

// ImageHref sets the link URL for the image.
func ImageHref(href string) ImageOption {
	return func(ib *ImageBlock) {
		ib.Href = href
	}
}

// ImageFullWidth sets the image to full width.
func ImageFullWidth() ImageOption {
	return func(ib *ImageBlock) {
		ib.Align = ImageAlignmentFull
	}
}

// AddImage adds an image block to the email with options.
func (builder *Builder) AddImage(src, alt string, opts ...ImageOption) *Builder {
	ib := ImageBlock{Src: src, Alt: alt}
	for _, opt := range opts {
		opt(&ib)
	}
	return builder.Add(ib)
}

func (builder *Builder) AddTable(header string, columns []string, rows [][]string, options ...TableOption) *Builder {
	block := TableBlock{Header: header, Columns: columns, Rows: rows}
	for _, option := range options {
		option(&block)
	}

	return builder.Add(block)
}

func (builder *Builder) AddVerificationCode(label, code string) *Builder {
	return builder.Add(VerificationCodeBlock{Label: label, Value: code})
}

func (builder *Builder) AddTiles(entries []TileEntry, options ...TilesOption) *Builder {
	block := TilesBlock{Entries: append([]TileEntry(nil), entries...)}
	for _, option := range options {
		option(&block)
	}

	return builder.Add(block)
}

func (builder *Builder) AddSection(blocks []Block, options ...SectionOption) *Builder {
	block := SectionBlock{Blocks: append([]Block(nil), blocks...), Border: true, Padding: 16}
	for _, option := range options {
		option(&block)
	}

	return builder.Add(block)
}

func (builder *Builder) AddSectionGroup(group *Group, options ...SectionOption) *Builder {
	if group == nil {
		return builder.AddSection(nil, options...)
	}

	return builder.AddSection(group.Blocks(), options...)
}

func (builder *Builder) AddGrid(items []GridItem, options ...GridOption) *Builder {
	block := GridBlock{Items: append([]GridItem(nil), items...), Columns: 2, Gap: 12}
	for _, option := range options {
		option(&block)
	}

	return builder.Add(block)
}

func GridItemGroup(group *Group) GridItem {
	if group == nil {
		return GridItem{}
	}

	return GridItem{Blocks: group.Blocks()}
}

func (builder *Builder) AddGridGroups(groups []*Group, options ...GridOption) *Builder {
	items := make([]GridItem, 0, len(groups))
	for _, group := range groups {
		if group == nil {
			continue
		}
		items = append(items, GridItemGroup(group))
	}

	return builder.AddGrid(items, options...)
}

func (builder *Builder) AddCardList(cards []CardItem, options ...CardListOption) *Builder {
	block := CardListBlock{Cards: append([]CardItem(nil), cards...), Columns: 2, Gap: 12, Border: true}
	for _, option := range options {
		option(&block)
	}

	return builder.Add(block)
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

func (builder *Builder) Build() *Email {
	builder.mu.Lock()
	defer builder.mu.Unlock()

	result := &Email{
		header:    cloneHeader(builder.header),
		footer:    cloneFooter(builder.footer),
		preheader: builder.preheader,
		values:    normalizeValues(builder.values, builder.theme.DefaultStyles()),
		blocks:    append([]Block(nil), builder.blocks...),
		theme:     builder.theme,
	}

	return result
}
