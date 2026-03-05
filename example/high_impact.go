package example

import (
	"github.com/gzuidhof/myrtle"
	"github.com/gzuidhof/myrtle/theme"
	defaulttheme "github.com/gzuidhof/myrtle/theme/default"
)

func WeeklyOperationsBriefEmail() (*myrtle.Email, error) {
	return WeeklyOperationsBriefEmailWithTheme(defaulttheme.New())
}

func WeeklyOperationsBriefEmailWithTheme(selectedTheme theme.Theme) (*myrtle.Email, error) {
	if selectedTheme == nil {
		selectedTheme = defaulttheme.New()
	}

	return myrtle.NewBuilder(
		selectedTheme,
		myrtle.WithStyles(theme.Styles{ColorPrimary: "#953ffd"}),
		myrtle.WithHeader(myrtle.HeadingBlock{Text: "Weekly operations brief", Level: 1}),
	).
		WithPreheader("High-impact summary with status, metrics, and next actions").
		AddBadge(myrtle.ToneSuccess, "Operational").
		AddSummaryCard("Status summary", "Queue latency is back within SLA after this morning's mitigation.", "Updated 6 minutes ago").
		AddStatsRow("Core metrics", []myrtle.StatItem{
			{Label: "Delivery", Value: "99.8%", Delta: "+0.3%", DeltaSemantic: myrtle.StatDeltaSemanticPositive},
			{Label: "Open rate", Value: "42.1%", Delta: "-1.2%", DeltaSemantic: myrtle.StatDeltaSemanticNegative},
			{Label: "Bounce", Value: "0.9%", Delta: "-0.1%", DeltaSemantic: myrtle.StatDeltaSemanticPositive},
		}).
		AddVerticalBarChart(
			[]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21", "22", "23", "24", "25", "26", "27", "28", "29", "30"},
			[]myrtle.VerticalBarChartSeries{
				{Key: "processed", Label: "Processed", Color: "#953ffd", Values: []float64{240, 255, 250, 262, 270, 220, 205, 245, 260, 255, 268, 276, 226, 210, 250, 265, 260, 272, 280, 230, 214, 254, 269, 264, 276, 284, 234, 218, 258, 272}},
				{Key: "resolved", Label: "Resolved", Color: "#a662fa", Values: []float64{222, 236, 232, 243, 249, 204, 190, 227, 241, 237, 249, 257, 210, 196, 232, 246, 241, 253, 260, 214, 199, 236, 250, 246, 257, 264, 218, 203, 240, 253}},
				{Key: "sla", Label: "Within SLA", Color: "#c395fb", Values: []float64{206, 219, 215, 226, 231, 189, 176, 210, 223, 220, 231, 238, 194, 181, 215, 228, 224, 235, 242, 199, 185, 219, 232, 228, 238, 245, 202, 188, 222, 234}},
			},
			myrtle.VerticalBarChartTitle("Daily throughput trend (last 30 days)"),
			myrtle.VerticalBarChartHeight(210),
			myrtle.VerticalBarChartColumnGap(1),
			myrtle.VerticalBarChartLegendPlacement(myrtle.VerticalBarChartLegendBottom),
			myrtle.VerticalBarChartAxisShowCategoryLabels(true),
			myrtle.VerticalBarChartAxisLabelFormat(myrtle.VerticalBarChartAxisLabelFormatNumber),
		).
		AddTimeline("Incident timeline", []myrtle.TimelineItem{
			{Time: "09:07", Title: "Detected", Detail: "Webhook retries rose above baseline."},
			{Time: "09:18", Title: "Mitigated", Detail: "Workers scaled and queue drained."},
			{Time: "09:42", Title: "Resolved", Detail: "Latency returned to normal."},
		}, myrtle.TimelineCurrentIndex(2)).
		AddAttachment("weekly-ops-report.pdf", "PDF · 312 KB", "https://example.com/reports/weekly-ops.pdf", "Download report").
		AddText("Review the complete report and approve next sprint priorities:").
		AddButton("Open dashboard", "https://example.com/ops/dashboard").
		AddLegal("Myrtle Inc.", "Dam Square 1, 1012 JS Amsterdam, Netherlands", "https://example.com/preferences", "https://example.com/unsubscribe").
		Build(), nil
}

func HighImpactEmail() (*myrtle.Email, error) {
	return WeeklyOperationsBriefEmail()
}

func HighImpactEmailWithTheme(selectedTheme theme.Theme) (*myrtle.Email, error) {
	return WeeklyOperationsBriefEmailWithTheme(selectedTheme)
}
