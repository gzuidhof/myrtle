package myrtle

// ButtonOption configures a ButtonBlock.
type ButtonOption func(*ButtonBlock)

// ButtonTone sets the tone of a button.
func ButtonTone(value Tone) ButtonOption {
	return func(block *ButtonBlock) {
		block.Tone = value
	}
}

// ButtonStyle sets the style of a button.
func ButtonStyle(value ButtonStyleValue) ButtonOption {
	return func(block *ButtonBlock) {
		block.Style = value
	}
}

// ButtonFullWidth toggles full-width rendering for a button.
func ButtonFullWidth(value bool) ButtonOption {
	return func(block *ButtonBlock) {
		block.FullWidth = value
	}
}

// ButtonSize sets the size of a button.
func ButtonSize(value ButtonSizeValue) ButtonOption {
	return func(block *ButtonBlock) {
		block.Size = value
	}
}

// ButtonNoWrap toggles label wrapping for a button.
func ButtonNoWrap(value bool) ButtonOption {
	return func(block *ButtonBlock) {
		block.NoWrap = value
	}
}

// ButtonAlign sets the alignment of a button.
func ButtonAlign(value ButtonAlignmentValue) ButtonOption {
	return func(block *ButtonBlock) {
		block.Alignment = value
	}
}

// ButtonGroupOption configures a ButtonGroupBlock.
type ButtonGroupOption func(*ButtonGroupBlock)

// PanelOption configures a PanelBlock.
type PanelOption func(*PanelBlock)

// GridOption configures a GridBlock.
type GridOption func(*GridBlock)

// CardListOption configures a CardListBlock.
type CardListOption func(*CardListBlock)

// SpacerOption configures a SpacerBlock.
type SpacerOption func(*SpacerBlock)

// DividerOption configures a DividerBlock.
type DividerOption func(*DividerBlock)

// ButtonGroupAlign sets the alignment of buttons in the group.
func ButtonGroupAlign(value ButtonAlignmentValue) ButtonGroupOption {
	return func(block *ButtonGroupBlock) {
		block.Alignment = value
	}
}

// ButtonGroupJoined toggles joined rendering for grouped buttons.
func ButtonGroupJoined(value bool) ButtonGroupOption {
	return func(block *ButtonGroupBlock) {
		block.Joined = value
	}
}

// ButtonGroupStackOnMobile toggles stacking grouped buttons on mobile.
func ButtonGroupStackOnMobile(value bool) ButtonGroupOption {
	return func(block *ButtonGroupBlock) {
		block.StackOnMobile = value
	}
}

// ButtonGroupFullWidthOnMobile toggles full-width grouped buttons on mobile.
func ButtonGroupFullWidthOnMobile(value bool) ButtonGroupOption {
	return func(block *ButtonGroupBlock) {
		block.FullWidthOnMobile = value
	}
}

// ButtonGroupGap sets the horizontal gap between grouped buttons.
func ButtonGroupGap(value int) ButtonGroupOption {
	return func(block *ButtonGroupBlock) {
		block.Gap = value
	}
}

// PanelTitle sets the panel title.
func PanelTitle(value string) PanelOption {
	return func(block *PanelBlock) {
		block.Title = value
	}
}

// PanelSubtitle sets the panel subtitle.
func PanelSubtitle(value string) PanelOption {
	return func(block *PanelBlock) {
		block.Subtitle = value
	}
}

// PanelCategory sets the panel category label.
func PanelCategory(value string) PanelOption {
	return func(block *PanelBlock) {
		block.Category = value
	}
}

// PanelBorder toggles the panel border.
func PanelBorder(value bool) PanelOption {
	return func(block *PanelBlock) {
		block.Border = value
	}
}

// PanelPadding sets panel content padding in pixels.
func PanelPadding(value int) PanelOption {
	return func(block *PanelBlock) {
		block.Padding = value
	}
}

// PanelInsetMode sets the inset mode of the panel.
func PanelInsetMode(value InsetMode) PanelOption {
	return func(block *PanelBlock) {
		block.InsetMode = value
	}
}

// PanelHeaderless toggles rendering the panel without its header section.
func PanelHeaderless(value bool) PanelOption {
	return func(block *PanelBlock) {
		block.Headerless = value
	}
}

// GridColumns sets the number of columns in a grid.
func GridColumns(value int) GridOption {
	return func(block *GridBlock) {
		block.Columns = value
	}
}

// GridGap sets the gap between grid items in pixels.
func GridGap(value int) GridOption {
	return func(block *GridBlock) {
		block.Gap = value
	}
}

// GridBorder toggles the grid border.
func GridBorder(value bool) GridOption {
	return func(block *GridBlock) {
		block.Border = value
	}
}

// GridInsetMode sets the inset mode of the grid.
func GridInsetMode(value InsetMode) GridOption {
	return func(block *GridBlock) {
		block.InsetMode = value
	}
}

// CardListColumns sets the number of columns in a card list.
func CardListColumns(value int) CardListOption {
	return func(block *CardListBlock) {
		block.Columns = value
	}
}

// CardListGap sets the gap between card list items in pixels.
func CardListGap(value int) CardListOption {
	return func(block *CardListBlock) {
		block.Gap = value
	}
}

// CardListBorder toggles borders around cards.
func CardListBorder(value bool) CardListOption {
	return func(block *CardListBlock) {
		block.Border = value
	}
}

// CardListInsetMode sets the inset mode of the card list.
func CardListInsetMode(value InsetMode) CardListOption {
	return func(block *CardListBlock) {
		block.InsetMode = value
	}
}

// SpacerSize sets spacer height in pixels.
func SpacerSize(value int) SpacerOption {
	return func(block *SpacerBlock) {
		block.Size = value
	}
}

// DividerStyle sets the divider variant.
func DividerStyle(value DividerVariant) DividerOption {
	return func(block *DividerBlock) {
		block.Variant = value
	}
}

// DividerThickness sets divider thickness in pixels.
func DividerThickness(value int) DividerOption {
	return func(block *DividerBlock) {
		block.Thickness = value
	}
}

// DividerInset sets divider inset in pixels.
func DividerInset(value int) DividerOption {
	return func(block *DividerBlock) {
		block.Inset = value
	}
}

// DividerLabel sets divider label text.
func DividerLabel(value string) DividerOption {
	return func(block *DividerBlock) {
		block.Label = value
	}
}

// DividerInsetMode sets the inset mode of the divider.
func DividerInsetMode(value InsetMode) DividerOption {
	return func(block *DividerBlock) {
		block.InsetMode = value
	}
}

// CalloutOption configures a CalloutBlock.
type CalloutOption func(*CalloutBlock)

// MessageOption configures a MessageBlock.
type MessageOption func(*MessageBlock)

// MessageDigestOption configures a MessageDigestBlock.
type MessageDigestOption func(*MessageDigestBlock)

// ProgressOption configures a ProgressBlock.
type ProgressOption func(*ProgressBlock)

// DistributionOption configures a DistributionBlock.
type DistributionOption func(*DistributionBlock)

// AttachmentOption configures an AttachmentBlock.
type AttachmentOption func(*AttachmentBlock)

// VerificationCodeOption configures a VerificationCodeBlock.
type VerificationCodeOption func(*VerificationCodeBlock)

// EmptyStateOption configures an EmptyStateBlock.
type EmptyStateOption func(*EmptyStateBlock)

// SummaryCardOption configures a SummaryCardBlock.
type SummaryCardOption func(*SummaryCardBlock)

// PriceSummaryOption configures a PriceSummaryBlock.
type PriceSummaryOption func(*PriceSummaryBlock)

// HeroOption configures a HeroBlock.
type HeroOption func(*HeroBlock)

// CalloutStyle sets the callout variant.
func CalloutStyle(variant CalloutVariant) CalloutOption {
	return func(block *CalloutBlock) {
		block.Variant = variant
	}
}

// CalloutLink sets the callout link label and URL.
func CalloutLink(label, url string) CalloutOption {
	return func(block *CalloutBlock) {
		block.LinkLabel = label
		block.LinkURL = url
	}
}

// CalloutInsetMode sets the inset mode of the callout.
func CalloutInsetMode(value InsetMode) CalloutOption {
	return func(block *CalloutBlock) {
		block.InsetMode = value
	}
}

// MessageInsetMode sets the inset mode of the message block.
func MessageInsetMode(value InsetMode) MessageOption {
	return func(block *MessageBlock) {
		block.InsetMode = value
	}
}

// MessageDigestTitle sets the title of a message digest block.
func MessageDigestTitle(value string) MessageDigestOption {
	return func(block *MessageDigestBlock) {
		block.Title = value
	}
}

// MessageDigestSubtitle sets the subtitle of a message digest block.
func MessageDigestSubtitle(value string) MessageDigestOption {
	return func(block *MessageDigestBlock) {
		block.Subtitle = value
	}
}

// MessageDigestFooter sets the footer text of a message digest block.
func MessageDigestFooter(value string) MessageDigestOption {
	return func(block *MessageDigestBlock) {
		block.Footer = value
	}
}

// MessageDigestEmptyText sets empty-state text for a message digest.
func MessageDigestEmptyText(value string) MessageDigestOption {
	return func(block *MessageDigestBlock) {
		block.EmptyText = value
	}
}

// MessageDigestMaxItems sets the maximum number of digest items to render.
func MessageDigestMaxItems(value int) MessageDigestOption {
	return func(block *MessageDigestBlock) {
		block.MaxItems = value
	}
}

// MessageDigestInsetMode sets the inset mode of the message digest block.
func MessageDigestInsetMode(value InsetMode) MessageDigestOption {
	return func(block *MessageDigestBlock) {
		block.InsetMode = value
	}
}

// TimelineOption configures a TimelineBlock.
type TimelineOption func(*TimelineBlock)

// TimelineCurrentIndex sets the currently active timeline index.
func TimelineCurrentIndex(value int) TimelineOption {
	return func(block *TimelineBlock) {
		block.HasCurrentIndex = true
		block.CurrentIndex = value
	}
}

// TimelineAggregateHeader sets the aggregate header text for timeline groups.
func TimelineAggregateHeader(value string) TimelineOption {
	return func(block *TimelineBlock) {
		block.AggregateHeader = value
	}
}

// TimelineInsetMode sets the inset mode of the timeline.
func TimelineInsetMode(value InsetMode) TimelineOption {
	return func(block *TimelineBlock) {
		block.InsetMode = value
	}
}

// StackedBarOption configures a StackedBarBlock.
type StackedBarOption func(*StackedBarBlock)

// StackedBarTotal sets the summary label and value shown with the stacked bar.
func StackedBarTotal(label, value string) StackedBarOption {
	return func(block *StackedBarBlock) {
		block.TotalLabel = label
		block.TotalValue = value
	}
}

// StackedBarInsetMode sets the inset mode of the stacked bar block.
func StackedBarInsetMode(value InsetMode) StackedBarOption {
	return func(block *StackedBarBlock) {
		block.InsetMode = value
	}
}

// TableOption configures a TableBlock.
type TableOption func(*TableBlock)

// HorizontalBarChartOption configures a HorizontalBarChartBlock.
type HorizontalBarChartOption func(*HorizontalBarChartBlock)

// VerticalBarChartOption configures a VerticalBarChartBlock.
type VerticalBarChartOption func(*VerticalBarChartBlock)

// SparklineOption configures a SparklineBlock.
type SparklineOption func(*SparklineBlock)

// TilesOption configures a TilesBlock.
type TilesOption func(*TilesBlock)

// SparklineDelta sets the delta label shown with the sparkline.
func SparklineDelta(value string) SparklineOption {
	return func(block *SparklineBlock) {
		block.Delta = value
	}
}

// SparklineDeltaSemantic sets semantic styling for the sparkline delta value.
func SparklineDeltaSemantic(value StatDeltaSemantic) SparklineOption {
	return func(block *SparklineBlock) {
		block.DeltaSemantic = value
	}
}

// SparklineTone sets the tone of the sparkline.
func SparklineTone(value Tone) SparklineOption {
	return func(block *SparklineBlock) {
		block.Tone = value
	}
}

// SparklineInsetMode sets the inset mode of the sparkline block.
func SparklineInsetMode(value InsetMode) SparklineOption {
	return func(block *SparklineBlock) {
		block.InsetMode = value
	}
}

// HorizontalBarChartThickness sets bar thickness in pixels.
func HorizontalBarChartThickness(value int) HorizontalBarChartOption {
	return func(block *HorizontalBarChartBlock) {
		block.Thickness = value
	}
}

// HorizontalBarChartLabelsInsideBars toggles labels inside bars.
func HorizontalBarChartLabelsInsideBars(value bool) HorizontalBarChartOption {
	return func(block *HorizontalBarChartBlock) {
		block.ShowLabelsInsideBars = value
	}
}

// HorizontalBarChartTransparentBackground toggles transparent chart background.
func HorizontalBarChartTransparentBackground(value bool) HorizontalBarChartOption {
	return func(block *HorizontalBarChartBlock) {
		block.TransparentBackground = value
	}
}

// HorizontalBarChartTone sets the tone of the horizontal bar chart.
func HorizontalBarChartTone(value Tone) HorizontalBarChartOption {
	return func(block *HorizontalBarChartBlock) {
		block.Tone = value
	}
}

// HorizontalBarChartInsetMode sets the inset mode of the horizontal bar chart.
func HorizontalBarChartInsetMode(value InsetMode) HorizontalBarChartOption {
	return func(block *HorizontalBarChartBlock) {
		block.InsetMode = value
	}
}

// VerticalBarChartHeight sets chart height in pixels.
func VerticalBarChartHeight(value int) VerticalBarChartOption {
	return func(block *VerticalBarChartBlock) {
		block.Height = value
	}
}

// VerticalBarChartTitle sets the chart title.
func VerticalBarChartTitle(value string) VerticalBarChartOption {
	return func(block *VerticalBarChartBlock) {
		block.Title = value
	}
}

// VerticalBarChartSubtitle sets the chart subtitle.
func VerticalBarChartSubtitle(value string) VerticalBarChartOption {
	return func(block *VerticalBarChartBlock) {
		block.Subtitle = value
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

// VerticalBarChartColumnGap sets spacing between columns in pixels.
func VerticalBarChartColumnGap(value int) VerticalBarChartOption {
	return func(block *VerticalBarChartBlock) {
		block.HasColumnGap = true
		block.ColumnGap = value
	}
}

// VerticalBarChartOuterGap sets outer chart padding in pixels.
func VerticalBarChartOuterGap(value int) VerticalBarChartOption {
	return func(block *VerticalBarChartBlock) {
		block.HasOuterGap = true
		block.OuterGap = value
	}
}

// VerticalBarChartCategoryGap is kept as a compatibility alias.
// Prefer VerticalBarChartColumnGap for clarity.
func VerticalBarChartCategoryGap(value int) VerticalBarChartOption {
	return VerticalBarChartColumnGap(value)
}

// VerticalBarChartTransparentBackground toggles a transparent chart background.
func VerticalBarChartTransparentBackground(value bool) VerticalBarChartOption {
	return func(block *VerticalBarChartBlock) {
		block.TransparentBackground = value
	}
}

// VerticalBarChartTone sets the tone of the vertical bar chart.
func VerticalBarChartTone(value Tone) VerticalBarChartOption {
	return func(block *VerticalBarChartBlock) {
		block.Tone = value
	}
}

// VerticalBarChartInsetMode sets the inset mode of the vertical bar chart.
func VerticalBarChartInsetMode(value InsetMode) VerticalBarChartOption {
	return func(block *VerticalBarChartBlock) {
		block.InsetMode = value
	}
}

// VerticalBarChartLegendPlacement sets where the legend is rendered.
func VerticalBarChartLegendPlacement(value VerticalBarChartLegendPlacementValue) VerticalBarChartOption {
	return func(block *VerticalBarChartBlock) {
		block.LegendPlacement = normalizedLegendPlacement(value)
	}
}

// VerticalBarChartLegend sets legend items for the chart.
func VerticalBarChartLegend(items []VerticalBarChartLegendItem) VerticalBarChartOption {
	return func(block *VerticalBarChartBlock) {
		if len(items) == 0 {
			block.Legend = nil
			return
		}

		block.Legend = append([]VerticalBarChartLegendItem(nil), items...)
	}
}

// VerticalBarChartLegendConfigOption sets legend placement and items together.
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

// VerticalBarChartAxisShowBaseline toggles rendering the baseline.
func VerticalBarChartAxisShowBaseline(value bool) VerticalBarChartOption {
	return func(block *VerticalBarChartBlock) {
		block.Axis.ShowBaseline = value
	}
}

// VerticalBarChartAxisShowYTicks toggles rendering Y-axis ticks.
func VerticalBarChartAxisShowYTicks(value bool) VerticalBarChartOption {
	return func(block *VerticalBarChartBlock) {
		block.Axis.ShowYTicks = value
	}
}

// VerticalBarChartAxisDrawYAxisLine toggles rendering the Y-axis line.
func VerticalBarChartAxisDrawYAxisLine(value bool) VerticalBarChartOption {
	return func(block *VerticalBarChartBlock) {
		block.Axis.HasDrawYAxisLine = true
		block.Axis.DrawYAxisLine = value
	}
}

// VerticalBarChartAxisShowCategoryLabels toggles category labels on the axis.
func VerticalBarChartAxisShowCategoryLabels(value bool) VerticalBarChartOption {
	return func(block *VerticalBarChartBlock) {
		block.Axis.HasShowCategoryLabels = true
		block.Axis.ShowCategoryLabels = value
	}
}

// VerticalBarChartAxisLabelFormat sets axis label formatting style.
func VerticalBarChartAxisLabelFormat(value VerticalBarChartAxisLabelFormatValue) VerticalBarChartOption {
	return func(block *VerticalBarChartBlock) {
		block.Axis.LabelFormat = value
	}
}

// VerticalBarChartAxisMin sets an explicit minimum axis value.
func VerticalBarChartAxisMin(value float64) VerticalBarChartOption {
	return func(block *VerticalBarChartBlock) {
		block.Axis.HasMin = true
		block.Axis.Min = value
	}
}

// VerticalBarChartAxisMax sets an explicit maximum axis value.
func VerticalBarChartAxisMax(value float64) VerticalBarChartOption {
	return func(block *VerticalBarChartBlock) {
		block.Axis.HasMax = true
		block.Axis.Max = value
	}
}

// VerticalBarChartAxisConfig replaces the full axis configuration.
func VerticalBarChartAxisConfig(value VerticalBarChartAxis) VerticalBarChartOption {
	return func(block *VerticalBarChartBlock) {
		block.Axis = value
	}
}

// VerticalBarChartValueLabelsOption sets value label rendering options.
func VerticalBarChartValueLabelsOption(value VerticalBarChartValueLabels) VerticalBarChartOption {
	return func(block *VerticalBarChartBlock) {
		block.ValueLabels = value
	}
}

// VerticalBarChartValueFormatterOption sets value formatting rules.
func VerticalBarChartValueFormatterOption(value VerticalBarChartValueFormatter) VerticalBarChartOption {
	return func(block *VerticalBarChartBlock) {
		block.ValueFormatter = value
	}
}

// TilesColumns sets the number of tile columns.
func TilesColumns(value int) TilesOption {
	return func(block *TilesBlock) {
		block.Columns = value
	}
}

// TilesBorder toggles tile borders.
func TilesBorder(value bool) TilesOption {
	return func(block *TilesBlock) {
		block.Border = value
	}
}

// TilesTransparentBackground toggles transparent tile backgrounds.
func TilesTransparentBackground(value bool) TilesOption {
	return func(block *TilesBlock) {
		block.TransparentBackground = value
	}
}

// TilesAlign sets alignment for tile content.
func TilesAlign(value TileAlignment) TilesOption {
	return func(block *TilesBlock) {
		block.Alignment = value
	}
}

// TilesInsetMode sets the inset mode of the tiles block.
func TilesInsetMode(value InsetMode) TilesOption {
	return func(block *TilesBlock) {
		block.InsetMode = value
	}
}

// TableZebraRows toggles zebra striping for table rows.
func TableZebraRows(value bool) TableOption {
	return func(block *TableBlock) {
		block.ZebraRows = value
	}
}

// TableTitle sets the table header text.
func TableTitle(value string) TableOption {
	return func(block *TableBlock) {
		block.Header = value
	}
}

// TableLegendSwatches sets legend swatch colors for the table.
func TableLegendSwatches(value []string) TableOption {
	return func(block *TableBlock) {
		block.HasLegendSwatches = true
		if len(value) == 0 {
			block.LegendSwatches = nil
			return
		}

		block.LegendSwatches = append([]string(nil), value...)
	}
}

// TableCompact toggles compact table spacing.
func TableCompact(value bool) TableOption {
	return func(block *TableBlock) {
		block.Compact = value
	}
}

// TableDensity sets the table density mode.
func TableDensity(value TableDensityValue) TableOption {
	return func(block *TableBlock) {
		block.Density = value
	}
}

// TableHeaderTone sets the table header tone.
func TableHeaderTone(value TableHeaderToneValue) TableOption {
	return func(block *TableBlock) {
		block.HeaderTone = value
	}
}

// TableBorderStyle sets the table border style.
func TableBorderStyle(value TableBorderStyleValue) TableOption {
	return func(block *TableBlock) {
		block.BorderStyle = value
	}
}

// TableRightAlignNumericColumns toggles right alignment for numeric columns.
func TableRightAlignNumericColumns(value bool) TableOption {
	return func(block *TableBlock) {
		block.RightAlignNumericColumns = value
	}
}

// TableEmphasizeTotalRow toggles emphasized styling for the total row.
func TableEmphasizeTotalRow(value bool) TableOption {
	return func(block *TableBlock) {
		block.EmphasizeTotalRow = value
	}
}

// TableInsetMode sets the inset mode of the table.
func TableInsetMode(value InsetMode) TableOption {
	return func(block *TableBlock) {
		block.InsetMode = value
	}
}

// TableColumnAlignments sets explicit alignment per table column index.
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

// ProgressInsetMode sets the inset mode of the progress block.
func ProgressInsetMode(value InsetMode) ProgressOption {
	return func(block *ProgressBlock) {
		block.InsetMode = value
	}
}

// DistributionInsetMode sets the inset mode of the distribution block.
func DistributionInsetMode(value InsetMode) DistributionOption {
	return func(block *DistributionBlock) {
		block.InsetMode = value
	}
}

// AttachmentInsetMode sets the inset mode of the attachment block.
func AttachmentInsetMode(value InsetMode) AttachmentOption {
	return func(block *AttachmentBlock) {
		block.InsetMode = value
	}
}

// VerificationCodeInsetMode sets the inset mode of the verification code block.
func VerificationCodeInsetMode(value InsetMode) VerificationCodeOption {
	return func(block *VerificationCodeBlock) {
		block.InsetMode = value
	}
}

// VerificationCodeTone sets the semantic tone of the verification code block.
func VerificationCodeTone(value Tone) VerificationCodeOption {
	return func(block *VerificationCodeBlock) {
		block.Tone = value
	}
}

// VerificationCodeMonospace toggles monospace rendering for the code value.
func VerificationCodeMonospace(value bool) VerificationCodeOption {
	return func(block *VerificationCodeBlock) {
		block.UseMonospace = value
		block.useMonospaceSet = true
	}
}

// VerificationCodeSpacing sets the code letter spacing in em units.
func VerificationCodeSpacing(value float64) VerificationCodeOption {
	return func(block *VerificationCodeBlock) {
		block.CharacterSpacingEm = value
		block.characterSpacingEmSet = true
	}
}

// EmptyStateInsetMode sets the inset mode of the empty state block.
func EmptyStateInsetMode(value InsetMode) EmptyStateOption {
	return func(block *EmptyStateBlock) {
		block.InsetMode = value
	}
}

// EmptyStateTone sets the visual tone of the empty state block.
func EmptyStateTone(value Tone) EmptyStateOption {
	return func(block *EmptyStateBlock) {
		block.Tone = value
	}
}

// SummaryCardInsetMode sets the inset mode of the summary card block.
func SummaryCardInsetMode(value InsetMode) SummaryCardOption {
	return func(block *SummaryCardBlock) {
		block.InsetMode = value
	}
}

// SummaryCardTone sets the visual tone of the summary card block.
func SummaryCardTone(value Tone) SummaryCardOption {
	return func(block *SummaryCardBlock) {
		block.Tone = value
	}
}

// PriceSummaryInsetMode sets the inset mode of the price summary block.
func PriceSummaryInsetMode(value InsetMode) PriceSummaryOption {
	return func(block *PriceSummaryBlock) {
		block.InsetMode = value
	}
}

// HeroInsetMode sets the inset mode of the hero block.
func HeroInsetMode(value InsetMode) HeroOption {
	return func(block *HeroBlock) {
		block.InsetMode = value
	}
}

// HeroTone sets the visual tone of the hero block.
func HeroTone(value Tone) HeroOption {
	return func(block *HeroBlock) {
		block.Tone = value
	}
}

// HeroEyebrow sets the eyebrow text of the hero block.
func HeroEyebrow(value string) HeroOption {
	return func(block *HeroBlock) {
		block.Eyebrow = value
	}
}

// HeroImage sets the hero image URL and alt text.
func HeroImage(url, alt string) HeroOption {
	return func(block *HeroBlock) {
		block.ImageURL = url
		block.ImageAlt = alt
	}
}

// Add appends a block to the builder.
// Use this for custom or preconstructed block instances.
func (builder *Builder) Add(block Block) *Builder {
	builder.mu.Lock()
	defer builder.mu.Unlock()

	builder.blocks = append(builder.blocks, block)
	return builder
}

// AddText appends a text block to the builder.
// Text blocks render paragraph-style body copy.
func (builder *Builder) AddText(text string, options ...TextOption) *Builder {
	block := TextBlock{Text: text}
	for _, option := range options {
		option(&block)
	}
	builder.Add(block)

	return builder
}

// AddHeading appends a heading block to the builder.
// Heading blocks introduce and structure content sections.
func (builder *Builder) AddHeading(text string, options ...HeadingOption) *Builder {
	block := HeadingBlock{Text: text, Level: 2}
	for _, option := range options {
		option(&block)
	}
	return builder.Add(block)
}

// AddSpacer appends a spacer block to the builder.
// Spacer blocks create vertical rhythm between nearby sections.
func (builder *Builder) AddSpacer(options ...SpacerOption) *Builder {
	block := SpacerBlock{Size: 16}
	for _, option := range options {
		option(&block)
	}

	return builder.Add(block)
}

// AddList appends a list block to the builder.
// List blocks render ordered or unordered bullet content.
func (builder *Builder) AddList(items []string, ordered bool) *Builder {
	return builder.Add(ListBlock{Items: items, Ordered: ordered})
}

// AddKeyValue appends a key-value block to the builder.
// Key-value blocks present compact labeled facts and values.
func (builder *Builder) AddKeyValue(header string, pairs []KeyValuePair) *Builder {
	return builder.Add(KeyValueBlock{Header: header, Pairs: pairs})
}

// AddHorizontalBarChart appends a horizontal bar chart block to the builder.
// This block compares categories with left-to-right bars.
func (builder *Builder) AddHorizontalBarChart(header string, items []HorizontalBarChartItem, options ...HorizontalBarChartOption) *Builder {
	block := HorizontalBarChartBlock{Header: header, Items: append([]HorizontalBarChartItem(nil), items...)}
	for _, option := range options {
		option(&block)
	}

	return builder.Add(block)
}

// AddVerticalBarChart appends a vertical bar chart block to the builder.
// This block compares categories with bottom-to-top columns.
func (builder *Builder) AddVerticalBarChart(axisLabels []string, series []VerticalBarChartSeries, options ...VerticalBarChartOption) *Builder {
	block := VerticalBarChartBlock{
		AxisLabels: append([]string(nil), axisLabels...),
		Series:     append([]VerticalBarChartSeries(nil), series...),
	}
	for _, option := range options {
		option(&block)
	}

	return builder.Add(block)
}

// AddSparkline appends a sparkline block to the builder.
// Sparkline blocks summarize a short trend inline with key values.
func (builder *Builder) AddSparkline(header, label, value string, points []int, options ...SparklineOption) *Builder {
	block := SparklineBlock{Header: header, Label: label, Value: value, Points: append([]int(nil), points...)}
	for _, option := range options {
		option(&block)
	}

	return builder.Add(block)
}

// AddStackedBar appends a stacked bar block to the builder.
// Stacked bars show part-to-whole composition per row.
func (builder *Builder) AddStackedBar(header string, rows []StackedBarRow, options ...StackedBarOption) *Builder {
	block := StackedBarBlock{Header: header, Rows: append([]StackedBarRow(nil), rows...)}
	for _, option := range options {
		option(&block)
	}

	return builder.Add(block)
}

// AddProgress appends a progress block to the builder.
// Progress blocks communicate completion toward one or more goals.
func (builder *Builder) AddProgress(header string, items []ProgressItem, options ...ProgressOption) *Builder {
	block := ProgressBlock{Header: header, Items: append([]ProgressItem(nil), items...)}
	for _, option := range options {
		option(&block)
	}

	return builder.Add(block)
}

// AddDistribution appends a distribution block to the builder.
// Distribution blocks visualize bucketed values across ranges.
func (builder *Builder) AddDistribution(header string, buckets []DistributionBucket, options ...DistributionOption) *Builder {
	block := DistributionBlock{Header: header, Buckets: append([]DistributionBucket(nil), buckets...)}
	for _, option := range options {
		option(&block)
	}

	return builder.Add(block)
}

// AddTimeline appends a timeline block to the builder.
// Timeline blocks show ordered milestones or process steps.
func (builder *Builder) AddTimeline(header string, items []TimelineItem, options ...TimelineOption) *Builder {
	block := TimelineBlock{Header: header, CurrentIndex: -1, Items: append([]TimelineItem(nil), items...)}
	for _, option := range options {
		option(&block)
	}

	return builder.Add(block)
}

// AddStatsRow appends a stats row block to the builder.
// Stats rows present multiple compact KPI values in one line.
func (builder *Builder) AddStatsRow(header string, stats []StatItem) *Builder {
	return builder.Add(StatsRowBlock{Header: header, Stats: append([]StatItem(nil), stats...)})
}

// AddBadge appends a badge block to the builder.
// Badges highlight short status labels with visual tone.
func (builder *Builder) AddBadge(tone Tone, text string) *Builder {
	return builder.Add(BadgeBlock{Tone: tone, Text: text})
}

// AddSummaryCard appends a summary card block to the builder.
// Summary cards combine a title, message, and optional footer note.
func (builder *Builder) AddSummaryCard(title, body, footer string, options ...SummaryCardOption) *Builder {
	block := SummaryCardBlock{Title: title, Body: body, Footer: footer, Tone: ToneDefault}
	for _, option := range options {
		option(&block)
	}

	return builder.Add(block)
}

// AddAttachment appends an attachment block to the builder.
// Attachment blocks describe downloadable files with metadata and CTA.
func (builder *Builder) AddAttachment(filename, meta, url, cta string, options ...AttachmentOption) *Builder {
	block := AttachmentBlock{Filename: filename, Meta: meta, URL: url, CTA: cta}
	for _, option := range options {
		option(&block)
	}

	return builder.Add(block)
}

// AddQuote appends a quote block to the builder.
// Quote blocks emphasize testimonial or attribution-style text.
func (builder *Builder) AddQuote(text, author string) *Builder {
	return builder.Add(QuoteBlock{Text: text, Author: author})
}

// AddCallout appends a callout block to the builder.
// Callout blocks surface important notices with semantic styling.
func (builder *Builder) AddCallout(tone Tone, title, body string, options ...CalloutOption) *Builder {
	block := CalloutBlock{Tone: tone, Title: title, Body: body}
	for _, option := range options {
		option(&block)
	}

	return builder.Add(block)
}

// AddMessage appends a message block to the builder.
// Message blocks render conversational items in a digest/thread style.
func (builder *Builder) AddMessage(message MessageBlock, options ...MessageOption) *Builder {
	block := message
	for _, option := range options {
		option(&block)
	}

	return builder.Add(block)
}

// AddMessageDigest appends a message digest block to the builder.
// Digest blocks group multiple messages under a shared header/footer.
func (builder *Builder) AddMessageDigest(messages []MessageBlock, options ...MessageDigestOption) *Builder {
	block := MessageDigestBlock{Messages: append([]MessageBlock(nil), messages...), EmptyText: "No messages"}
	for _, option := range options {
		option(&block)
	}

	return builder.Add(block)
}

// AddLegal appends a legal/compliance block to the builder.
// Legal blocks include company address and subscription management links.
func (builder *Builder) AddLegal(companyName, address, manageURL, unsubscribeURL string) *Builder {
	return builder.Add(LegalBlock{
		CompanyName:    companyName,
		Address:        address,
		ManageURL:      manageURL,
		UnsubscribeURL: unsubscribeURL,
	})
}

// AddButton appends a button block to the builder.
// Button blocks render a primary call-to-action link.
func (builder *Builder) AddButton(label, url string, options ...ButtonOption) *Builder {
	block := ButtonBlock{Label: label, URL: url}
	for _, option := range options {
		option(&block)
	}

	return builder.Add(block)
}

// AddButtonGroup appends a grouped button block to the builder.
// Button groups place multiple CTAs in one aligned row or stack.
func (builder *Builder) AddButtonGroup(buttons []ButtonGroupButton, options ...ButtonGroupOption) *Builder {
	block := ButtonGroupBlock{Buttons: append([]ButtonGroupButton(nil), buttons...), Gap: 8}
	for _, option := range options {
		option(&block)
	}

	return builder.Add(block)
}

// AddDivider appends a divider block to the builder.
// Divider blocks separate sections with a horizontal rule or label.
func (builder *Builder) AddDivider(options ...DividerOption) *Builder {
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

// ImageInsetMode sets the inset mode of the image block.
func ImageInsetMode(value InsetMode) ImageOption {
	return func(ib *ImageBlock) {
		ib.InsetMode = value
	}
}

// ImageTopSpacing sets the top spacing (in px) for the image block.
func ImageTopSpacing(px int) ImageOption {
	return func(ib *ImageBlock) {
		ib.HasTopSpacing = true
		ib.TopSpacing = px
	}
}

// ImageBottomSpacing sets the bottom spacing (in px) for the image block.
func ImageBottomSpacing(px int) ImageOption {
	return func(ib *ImageBlock) {
		ib.HasBottomSpacing = true
		ib.BottomSpacing = px
	}
}

// ImageCorners controls which image corners receive radius.
func ImageCorners(value ImageCornerMode) ImageOption {
	return func(ib *ImageBlock) {
		ib.CornerMode = normalizedImageCornerMode(value)
	}
}

// ImageNoCorners disables corner radius on all image corners.
func ImageNoCorners() ImageOption {
	return ImageCorners(ImageCornerModeNone)
}

// ImageAllCorners enables corner radius on all image corners.
func ImageAllCorners() ImageOption {
	return ImageCorners(ImageCornerModeAll)
}

// ImageTopCorners enables corner radius only on top corners.
func ImageTopCorners() ImageOption {
	return ImageCorners(ImageCornerModeTop)
}

// ImageBottomCorners enables corner radius only on bottom corners.
func ImageBottomCorners() ImageOption {
	return ImageCorners(ImageCornerModeBottom)
}

// AddImage adds an image block to the email with options.
// Image blocks render visual media with alignment and corner controls.
func (builder *Builder) AddImage(src, alt string, opts ...ImageOption) *Builder {
	ib := ImageBlock{Src: src, Alt: alt}
	for _, opt := range opts {
		opt(&ib)
	}
	return builder.Add(ib)
}

// AddTable appends a table block to the builder.
// Table blocks present structured rows and columns of data.
func (builder *Builder) AddTable(columns []string, rows [][]string, options ...TableOption) *Builder {
	block := TableBlock{Columns: columns, Rows: rows}
	for _, option := range options {
		option(&block)
	}

	return builder.Add(block)
}

// AddVerificationCode appends a verification code block to the builder.
// Verification code blocks highlight short one-time passcodes.
func (builder *Builder) AddVerificationCode(label, code string, options ...VerificationCodeOption) *Builder {
	block := VerificationCodeBlock{Label: label, Value: code}
	for _, option := range options {
		option(&block)
	}

	return builder.Add(block)
}

// AddTiles appends a tiles block to the builder.
// Tiles blocks show small metric cards in a compact grid.
func (builder *Builder) AddTiles(entries []TileEntry, options ...TilesOption) *Builder {
	block := TilesBlock{Entries: append([]TileEntry(nil), entries...)}
	for _, option := range options {
		option(&block)
	}

	return builder.Add(block)
}

// AddPanel appends a panel block wrapping optional content.
// Panels provide a bordered container to group related content.
func (builder *Builder) AddPanel(content Block, options ...PanelOption) *Builder {
	blocks := []Block{}
	if content != nil {
		blocks = append(blocks, content)
	}

	block := PanelBlock{Blocks: blocks, Border: true, Padding: 16}
	for _, option := range options {
		option(&block)
	}

	return builder.Add(block)
}

// AddGrid appends a grid block to the builder.
// Grid blocks lay out heterogeneous items across multiple columns.
func (builder *Builder) AddGrid(items []GridItem, options ...GridOption) *Builder {
	block := GridBlock{Items: append([]GridItem(nil), items...), Columns: 2, Gap: 12}
	for _, option := range options {
		option(&block)
	}

	return builder.Add(block)
}

// GridItemGroup wraps a Group as a GridItem.
func GridItemGroup(group *Group) GridItem {
	if group == nil {
		return GridItem{}
	}

	return GridItem{Content: group}
}

// AddGridGroups appends groups as a grid block.
// Each group becomes one grid cell with its own nested content.
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

// AddCardList appends a card list block to the builder.
// Card list blocks render repeated card entries in columns.
func (builder *Builder) AddCardList(cards []CardItem, options ...CardListOption) *Builder {
	block := CardListBlock{Cards: append([]CardItem(nil), cards...), Columns: 2, Gap: 12, Border: true}
	for _, option := range options {
		option(&block)
	}

	return builder.Add(block)
}

// AddFreeMarkdown appends a markdown-rendered content block to the builder.
// Free markdown blocks allow direct authoring of rich text snippets.
func (builder *Builder) AddFreeMarkdown(markdown string) *Builder {
	return builder.Add(FreeMarkdownBlock{Markdown: markdown})
}

// AddHero appends a hero block to the builder.
// Hero blocks present high-impact title, body, and optional CTA.
func (builder *Builder) AddHero(title, body, ctaLabel, ctaURL string, options ...HeroOption) *Builder {
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

	return builder.Add(block)
}

// AddFooterLinks appends a footer links block to the builder.
// Footer links blocks provide secondary navigation and policy links.
func (builder *Builder) AddFooterLinks(links []FooterLink, note string) *Builder {
	return builder.Add(FooterLinksBlock{Links: append([]FooterLink(nil), links...), Note: note})
}

// AddPriceSummary appends a price summary block to the builder.
// AddPriceSummary appends a price summary block to the builder.
// Price summary blocks itemize charges and present an order total.
func (builder *Builder) AddPriceSummary(header string, items []PriceLine, totalLabel, totalValue string, options ...PriceSummaryOption) *Builder {
	block := PriceSummaryBlock{Header: header, Items: append([]PriceLine(nil), items...), TotalLabel: totalLabel, TotalValue: totalValue}
	for _, option := range options {
		option(&block)
	}

	return builder.Add(block)
}

// AddEmptyState appends an empty-state block to the builder.
// Empty state blocks explain missing data and suggest next actions.
func (builder *Builder) AddEmptyState(title, body, actionLabel, actionURL string, options ...EmptyStateOption) *Builder {
	block := EmptyStateBlock{Title: title, Body: body, ActionLabel: actionLabel, ActionURL: actionURL}
	block.Tone = ToneDefault
	for _, option := range options {
		option(&block)
	}

	return builder.Add(block)
}

// Build materializes an immutable Email snapshot from the current builder state.
// Subsequent builder mutations do not affect the returned Email.
func (builder *Builder) Build() *Email {
	builder.mu.Lock()
	defer builder.mu.Unlock()

	result := &Email{
		header:                 cloneHeader(builder.header),
		footer:                 cloneFooter(builder.footer),
		preheader:              builder.preheader,
		preheaderPaddingRepeat: builder.preheaderPaddingRepeat,
		values:                 normalizeValues(builder.values, builder.theme.DefaultStyles()),
		blocks:                 append([]Block(nil), builder.blocks...),
		theme:                  builder.theme,
	}

	return result
}
