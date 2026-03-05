package example

import (
	"github.com/gzuidhof/myrtle"
	"github.com/gzuidhof/myrtle/theme"
	defaulttheme "github.com/gzuidhof/myrtle/theme/default"
)

func InsetModesEmail() (*myrtle.Email, error) {
	return InsetModesEmailWithTheme(defaulttheme.New())
}

func InsetModesEmailWithTheme(selectedTheme theme.Theme) (*myrtle.Email, error) {
	if selectedTheme == nil {
		selectedTheme = defaulttheme.New()
	}

	chartLabels := []string{"Mon", "Tue", "Wed", "Thu"}
	chartSeries := []myrtle.VerticalBarChartSeries{
		{Key: "sent", Label: "Sent", Values: []float64{32, 28, 40, 36}},
		{Key: "failed", Label: "Failed", Values: []float64{-4, -3, -5, -4}},
	}

	stackedRows := []myrtle.StackedBarRow{
		{Label: "Activation", Segments: []myrtle.StackedBarSegment{{Label: "Email", Percent: 52, Value: "52%"}, {Label: "Push", Percent: 30, Value: "30%"}, {Label: "SMS", Percent: 18, Value: "18%"}}},
	}

	timelineItems := []myrtle.TimelineItem{
		{Time: "09:07", Title: "Detected", Detail: "Latency spike"},
		{Time: "09:18", Title: "Mitigation", Detail: "Scaled workers"},
		{Time: "09:42", Title: "Resolved", Detail: "Back to baseline"},
	}

	tiles := []myrtle.TileEntry{
		{Content: "🚀", Title: "Launch", Subtitle: "Ready", Variant: myrtle.TileVariantHighlight},
		{Content: "12", Title: "Queued", Subtitle: "Jobs", Variant: myrtle.TileVariantWarning},
		{Content: "✅", Title: "Healthy", Subtitle: "All systems", Variant: myrtle.TileVariantSuccess},
	}

	cards := []myrtle.CardItem{
		{Title: "Card one", Body: "Container card with default inset."},
		{Title: "Card two", Body: "Container card with configurable inset."},
	}

	builder := myrtle.NewBuilder(selectedTheme, myrtle.WithStyles(theme.Styles{MainContentBodyTopSpacing: "0"})).
		WithPreheader("InsetMode demo for box components").
		WithHeader(commonHeaderGroupWithAlt("Inset Modes", "Inset modes", selectedTheme), myrtle.HeaderPlacement(myrtle.HeaderPlacementOutside)).
		AddImage("/assets/dark-gradient-hero.png", "Main container start image", myrtle.ImageFullWidth(), myrtle.ImageInsetMode(myrtle.InsetModeNone), myrtle.ImageTopSpacing(0), myrtle.ImageTopCorners()).
		AddHeading("InsetMode: default vs none", myrtle.HeadingLevel(1)).
		AddText("Each block below is rendered twice: first with default inset, then with InsetModeNone.").
		AddHeading("Image", myrtle.HeadingLevel(2)).
		AddImage("/assets/dark-gradient-hero.png", "Image default", myrtle.ImageFullWidth()).
		AddImage("/assets/dark-gradient-hero.png", "Image none", myrtle.ImageFullWidth(), myrtle.ImageInsetMode(myrtle.InsetModeNone))

	builder.
		AddHeading("Charts and graphs", myrtle.HeadingLevel(2)).
		AddHorizontalBarChart("Horizontal bar (default)", []myrtle.HorizontalBarChartItem{{Label: "US", Percent: 52}, {Label: "EMEA", Percent: 31}, {Label: "APAC", Percent: 17}}).
		AddHorizontalBarChart("Horizontal bar (none)", []myrtle.HorizontalBarChartItem{{Label: "US", Percent: 52}, {Label: "EMEA", Percent: 31}, {Label: "APAC", Percent: 17}}, myrtle.HorizontalBarChartInsetMode(myrtle.InsetModeNone)).
		AddVerticalBarChart(chartLabels, chartSeries, myrtle.VerticalBarChartTitle("Vertical bar (default)"), myrtle.VerticalBarChartLegendPlacement(myrtle.VerticalBarChartLegendBottom)).
		AddVerticalBarChart(chartLabels, chartSeries, myrtle.VerticalBarChartTitle("Vertical bar (none)"), myrtle.VerticalBarChartLegendPlacement(myrtle.VerticalBarChartLegendBottom), myrtle.VerticalBarChartInsetMode(myrtle.InsetModeNone)).
		AddSparkline("Sparkline (default)", "Throughput", "1,284", []int{8, 12, 9, 14, 18, 16, 20}, myrtle.SparklineDelta("+8%")).
		AddSparkline("Sparkline (none)", "Throughput", "1,284", []int{8, 12, 9, 14, 18, 16, 20}, myrtle.SparklineDelta("+8%"), myrtle.SparklineInsetMode(myrtle.InsetModeNone)).
		AddStackedBar("Stacked bar (default)", stackedRows, myrtle.StackedBarTotal("Total", "100%")).
		AddStackedBar("Stacked bar (none)", stackedRows, myrtle.StackedBarTotal("Total", "100%"), myrtle.StackedBarInsetMode(myrtle.InsetModeNone)).
		AddProgress("Progress (default)", []myrtle.ProgressItem{{Label: "Schema", Percent: 100}, {Label: "API", Percent: 76}, {Label: "Client", Percent: 48}}).
		AddProgress("Progress (none)", []myrtle.ProgressItem{{Label: "Schema", Percent: 100}, {Label: "API", Percent: 76}, {Label: "Client", Percent: 48}}, myrtle.ProgressInsetMode(myrtle.InsetModeNone)).
		AddDistribution("Distribution (default)", []myrtle.DistributionBucket{{Label: "0-50", Count: 62}, {Label: "51-100", Count: 44}, {Label: "101-200", Count: 21}, {Label: "200+", Count: 8}}).
		AddDistribution("Distribution (none)", []myrtle.DistributionBucket{{Label: "0-50", Count: 62}, {Label: "51-100", Count: 44}, {Label: "101-200", Count: 21}, {Label: "200+", Count: 8}}, myrtle.DistributionInsetMode(myrtle.InsetModeNone))

	builder.
		AddHeading("Container blocks", myrtle.HeadingLevel(2)).
		AddPanel(
			myrtle.NewGroup().
				AddImage("/assets/dark-gradient-hero.png", "Dark gradient hero", myrtle.ImageFullWidth()).
				AddText("Panel default inset."),
			myrtle.PanelTitle("Panel (default)"),
		).
		AddPanel(
			myrtle.NewGroup().
				AddImage("/assets/dark-gradient-hero.png", "Dark gradient hero", myrtle.ImageFullWidth()).
				AddText("Panel no inset."),
			myrtle.PanelTitle("Panel (none)"),
			myrtle.PanelInsetMode(myrtle.InsetModeNone),
		).
		AddHeading("Rounded container image corners", myrtle.HeadingLevel(3)).
		AddPanel(
			myrtle.NewGroup().
				AddImage("/assets/dark-gradient-hero.png", "Dark gradient hero", myrtle.ImageFullWidth(), myrtle.ImageInsetMode(myrtle.InsetModeNone), myrtle.ImageCorners(myrtle.ImageCornerModeTop)).
				AddText("Image uses top corners to match rounded panel header edge."),
			myrtle.PanelTitle("Panel image top corners"),
		).
		AddPanel(
			myrtle.NewGroup().
				AddText("Image uses bottom corners to match rounded panel footer edge.").
				AddImage("/assets/dark-gradient-hero.png", "Dark gradient hero", myrtle.ImageFullWidth(), myrtle.ImageInsetMode(myrtle.InsetModeNone), myrtle.ImageCorners(myrtle.ImageCornerModeBottom)),
			myrtle.PanelTitle("Panel image bottom corners"),
		).
		AddColumns(
			myrtle.NewGroup().
				AddImage("/assets/dark-gradient-hero.png", "Dark gradient hero", myrtle.ImageFullWidth()).
				AddText("Columns default inset."),
			myrtle.NewGroup().
				AddImage("/assets/dark-gradient-hero.png", "Dark gradient hero", myrtle.ImageFullWidth()).
				AddText("Second column."),
		).
		AddColumns(
			myrtle.NewGroup().
				AddImage("/assets/dark-gradient-hero.png", "Dark gradient hero", myrtle.ImageFullWidth()).
				AddText("Columns no inset."),
			myrtle.NewGroup().
				AddImage("/assets/dark-gradient-hero.png", "Dark gradient hero", myrtle.ImageFullWidth()).
				AddText("Second column."),
			myrtle.ColumnsInsetMode(myrtle.InsetModeNone),
		).
		AddGrid(
			[]myrtle.GridItem{
				myrtle.GridItemGroup(myrtle.NewGroup().AddImage("/assets/dark-gradient-hero.png", "Dark gradient hero", myrtle.ImageFullWidth()).AddText("Grid default 1")),
				myrtle.GridItemGroup(myrtle.NewGroup().AddImage("/assets/dark-gradient-hero.png", "Dark gradient hero", myrtle.ImageFullWidth()).AddText("Grid default 2")),
			},
			myrtle.GridColumns(2),
		).
		AddGrid(
			[]myrtle.GridItem{
				myrtle.GridItemGroup(myrtle.NewGroup().AddImage("/assets/dark-gradient-hero.png", "Dark gradient hero", myrtle.ImageFullWidth()).AddText("Grid none 1")),
				myrtle.GridItemGroup(myrtle.NewGroup().AddImage("/assets/dark-gradient-hero.png", "Dark gradient hero", myrtle.ImageFullWidth()).AddText("Grid none 2")),
			},
			myrtle.GridColumns(2),
			myrtle.GridInsetMode(myrtle.InsetModeNone),
		).
		AddCardList(cards, myrtle.CardListColumns(2)).
		AddCardList(cards, myrtle.CardListColumns(2), myrtle.CardListInsetMode(myrtle.InsetModeNone))

	builder.
		AddHeading("Other box blocks", myrtle.HeadingLevel(2)).
		AddMessage(
			myrtle.MessageBlock{SenderName: "Maya", Subject: "Message (default)", Preview: "This standalone message keeps the default inset.", URL: "https://example.com/messages/43"},
		).
		AddMessage(
			myrtle.MessageBlock{SenderName: "Maya", Subject: "Message (none)", Preview: "This standalone message uses InsetModeNone.", URL: "https://example.com/messages/43"},
			myrtle.MessageInsetMode(myrtle.InsetModeNone),
		).
		AddCallout(myrtle.ToneInfo, "Callout (default)", "Default inset callout").
		AddCallout(myrtle.ToneInfo, "Callout (none)", "No inset callout", myrtle.CalloutInsetMode(myrtle.InsetModeNone)).
		AddAttachment("invoice-Feb-2026.pdf", "PDF · 284 KB", "https://example.com/invoices/feb-2026.pdf", "Download").
		AddAttachment("invoice-Feb-2026.pdf", "PDF · 284 KB", "https://example.com/invoices/feb-2026.pdf", "Download", myrtle.AttachmentInsetMode(myrtle.InsetModeNone)).
		AddEmptyState("Empty state (default)", "No activity right now.", "Refresh", "https://example.com/refresh").
		AddEmptyState("Empty state (none)", "No activity right now.", "Refresh", "https://example.com/refresh", myrtle.EmptyStateInsetMode(myrtle.InsetModeNone)).
		AddTable([]string{"Metric", "Value"}, [][]string{{"Send rate", "99.9%"}, {"Bounce", "0.7%"}}, myrtle.TableTitle("Table (default)")).
		AddTable([]string{"Metric", "Value"}, [][]string{{"Send rate", "99.9%"}, {"Bounce", "0.7%"}}, myrtle.TableTitle("Table (none)"), myrtle.TableInsetMode(myrtle.InsetModeNone)).
		AddTiles(tiles, myrtle.TilesColumns(3)).
		AddTiles(tiles, myrtle.TilesColumns(3), myrtle.TilesInsetMode(myrtle.InsetModeNone)).
		AddTimeline("Timeline (default)", timelineItems, myrtle.TimelineCurrentIndex(1)).
		AddTimeline("Timeline (none)", timelineItems, myrtle.TimelineCurrentIndex(1), myrtle.TimelineInsetMode(myrtle.InsetModeNone)).
		AddVerificationCode("Verification (default)", "693028").
		AddVerificationCode("Verification (none)", "693028", myrtle.VerificationCodeInsetMode(myrtle.InsetModeNone)).
		AddPriceSummary("Price summary (default)", []myrtle.PriceLine{{Label: "Subtotal", Value: "$49.00"}, {Label: "Tax", Value: "$4.90"}}, "Total", "$53.90").
		AddPriceSummary("Price summary (none)", []myrtle.PriceLine{{Label: "Subtotal", Value: "$49.00"}, {Label: "Tax", Value: "$4.90"}}, "Total", "$53.90", myrtle.PriceSummaryInsetMode(myrtle.InsetModeNone)).
		AddMessageDigest(
			[]myrtle.MessageBlock{{SenderName: "Maya", Subject: "Launch update", Preview: "Can you review the draft?", URL: "https://example.com/messages/43"}},
			myrtle.MessageDigestTitle("Message digest (default)"),
		).
		AddMessageDigest(
			[]myrtle.MessageBlock{{SenderName: "Maya", Subject: "Launch update", Preview: "Can you review the draft?", URL: "https://example.com/messages/43"}},
			myrtle.MessageDigestTitle("Message digest (none)"),
			myrtle.MessageDigestInsetMode(myrtle.InsetModeNone),
		).
		AddImage("/assets/dark-gradient-hero.png", "Main container end image", myrtle.ImageFullWidth(), myrtle.ImageInsetMode(myrtle.InsetModeNone), myrtle.ImageBottomSpacing(0), myrtle.ImageBottomCorners()).
		WithFooter(commonFooterGroup(), myrtle.FooterPlacement(myrtle.FooterPlacementOutside))

	return builder.Build(), nil
}
