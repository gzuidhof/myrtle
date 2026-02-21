package example

import (
	"github.com/gzuidhof/myrtle"
	"github.com/gzuidhof/myrtle/theme"
	defaulttheme "github.com/gzuidhof/myrtle/theme/default"
)

func HighImpactEmail() (*myrtle.Email, error) {
	return HighImpactEmailWithTheme(defaulttheme.New())
}

func HighImpactEmailWithTheme(selectedTheme theme.Theme) (*myrtle.Email, error) {
	if selectedTheme == nil {
		selectedTheme = defaulttheme.New()
	}

	return myrtle.NewBuilder(
		selectedTheme,
		myrtle.WithStyles(theme.Styles{PrimaryColor: "#2563eb"}),
		myrtle.WithHeaderOptions(
			myrtle.HeaderTitle("Weekly operations brief"),
			myrtle.HeaderProduct("Myrtle Ops", "https://example.com/ops"),
			myrtle.HeaderLogo("/assets/logo.png", "Myrtle Ops"),
			myrtle.HeaderLogoCentered(true),
		),
	).
		Preheader("High-impact summary with status, metrics, and next actions").
		AddBadge(myrtle.BadgeToneSuccess, "Operational").
		AddSummaryCard("Status summary", "Queue latency is back within SLA after this morning's mitigation.", "Updated 6 minutes ago").
		AddStatsRow("Core metrics", []myrtle.StatItem{
			{Label: "Delivery", Value: "99.8%", Delta: "+0.3%", DeltaSemantic: myrtle.StatDeltaSemanticPositive},
			{Label: "Open rate", Value: "42.1%", Delta: "+1.2%", DeltaSemantic: myrtle.StatDeltaSemanticPositive},
			{Label: "Bounce", Value: "0.9%", Delta: "-0.1%", DeltaSemantic: myrtle.StatDeltaSemanticNegative},
		}).
		AddTimeline("Incident timeline", []myrtle.TimelineItem{
			{Time: "09:07", Title: "Detected", Detail: "Webhook retries rose above baseline."},
			{Time: "09:18", Title: "Mitigated", Detail: "Workers scaled and queue drained."},
			{Time: "09:42", Title: "Resolved", Detail: "Latency returned to normal."},
		}).
		AddAttachment("weekly-ops-report.pdf", "PDF · 312 KB", "https://example.com/reports/weekly-ops.pdf", "Download report").
		AddAction("Review the complete report and approve next sprint priorities:", "Open dashboard", "https://example.com/ops/dashboard").
		AddLegal("Myrtle Inc.", "123 Market St, San Francisco, CA", "https://example.com/preferences", "https://example.com/unsubscribe").
		Build(), nil
}
