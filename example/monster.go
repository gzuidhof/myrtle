package example

import (
	"github.com/gzuidhof/myrtle"
	"github.com/gzuidhof/myrtle/theme"
	defaulttheme "github.com/gzuidhof/myrtle/theme/default"
)

func MonsterEmail() (*myrtle.Email, error) {
	return MonsterEmailWithTheme(defaulttheme.New())
}

func MonsterEmailWithTheme(selectedTheme theme.Theme) (*myrtle.Email, error) {
	if selectedTheme == nil {
		selectedTheme = defaulttheme.New()
	}

	metricsRows := [][]string{
		{"Messages", "128,490", "+12%"},
		{"Delivery", "99.82%", "+0.4%"},
		{"CTR", "4.12%", "+0.6%"},
		{"Total", "232,402", "+4.2%"},
	}

	return myrtle.NewBuilder(
		selectedTheme,
		myrtle.WithStyles(theme.Styles{PrimaryColor: "#0ea5e9"}),
	).
		WithoutHeader().
		Preheader("A single email that demonstrates every available block").
		WithHeader(
			myrtle.HeaderTitle("The Myrtle Monster Email"),
			myrtle.HeaderProduct("Myrtle Platform", "https://github.com/gzuidhof/myrtle"),
			myrtle.HeaderLogo("/assets/logo.png", "Myrtle Platform"),
			myrtle.HeaderShowTextWithLogo(true),
		).
		Add(myrtle.HeroBlock{
			Eyebrow:  "Monster Demo",
			Title:    "A single email showing nearly every component",
			Body:     "Use this as a quick visual regression pass for block rendering across themes.",
			CTALabel: "Open docs",
			CTAURL:   "https://github.com/gzuidhof/myrtle",
			ImageURL: "/assets/hero.svg",
			ImageAlt: "Abstract hero",
		}).
		AddBadge(myrtle.BadgeToneInfo, "Info badge variant").
		AddBadge(myrtle.BadgeToneSuccess, "Success badge variant").
		AddBadge(myrtle.BadgeToneWarning, "Warning badge variant").
		AddBadge(myrtle.BadgeToneError, "Error badge variant").
		AddHeading("Everything in one place", myrtle.HeadingLevel(1)).
		AddText("This email intentionally demonstrates all built-in blocks and several composition patterns.").
		AddCallout(myrtle.CalloutTypeInfo, "Overview", "Scan this message to see all components rendered together.", myrtle.CalloutStyle(myrtle.CalloutVariantSoft)).
		AddCallout(myrtle.CalloutTypeWarning, "Variant", "Outline style for warning callouts.", myrtle.CalloutStyle(myrtle.CalloutVariantOutline)).
		AddCallout(myrtle.CalloutTypeCritical, "Variant", "Solid style for critical alerts.", myrtle.CalloutStyle(myrtle.CalloutVariantSolid)).
		AddSpacer(10).
		AddHeading("Buttons", myrtle.HeadingLevel(2)).
		AddButton("Primary CTA", "https://example.com/primary", myrtle.ButtonStyle(myrtle.ButtonVariantPrimary)).
		AddButton("Centered CTA", "https://example.com/center", myrtle.ButtonStyle(myrtle.ButtonVariantPrimary), myrtle.ButtonAlign(myrtle.ButtonAlignmentCenter)).
		AddButton("Secondary CTA", "https://example.com/secondary", myrtle.ButtonStyle(myrtle.ButtonVariantSecondary)).
		AddButton("Ghost CTA", "https://example.com/ghost", myrtle.ButtonStyle(myrtle.ButtonVariantGhost)).
		AddButton("Full-width primary", "https://example.com/full", myrtle.ButtonStyle(myrtle.ButtonVariantPrimary), myrtle.ButtonFullWidth(true)).
		AddButtonGroup([]myrtle.ButtonGroupButton{{Label: "Approve", URL: "https://example.com/approve", Variant: myrtle.ButtonVariantPrimary}, {Label: "Review", URL: "https://example.com/review", Variant: myrtle.ButtonVariantSecondary}, {Label: "Later", URL: "https://example.com/later", Variant: myrtle.ButtonVariantGhost}}, myrtle.ButtonGroupAlign(myrtle.ButtonAlignmentCenter), myrtle.ButtonGroupJoined(true)).
		AddColumns(
			func(column *myrtle.ColumnBuilder) {
				column.AddHeading("Quick links", myrtle.HeadingLevel(3)).
					AddList([]string{"Open dashboard", "Review incidents", "Manage preferences"}, false).
					AddButton("Open dashboard", "https://example.com/dashboard")
			},
			func(column *myrtle.ColumnBuilder) {
				column.AddHeading("Verification", myrtle.HeadingLevel(3)).
					Add(myrtle.CodeBlock{Label: "One-time code", Code: "483920"}).
					AddAction("Use this code to authorize billing changes.", "Open security", "https://example.com/security")
			},
			myrtle.ColumnsWidths(55, 45),
		).
		AddDivider().
		AddHeading("Metrics and highlights", myrtle.HeadingLevel(2)).
		AddTable(
			"Weekly performance",
			[]string{"Metric", "Value", "Delta"},
			metricsRows,
			myrtle.TableCompact(true),
			myrtle.TableZebraRows(true),
			myrtle.TableRightAlignNumericColumns(true),
			myrtle.TableEmphasizeTotalRow(true),
			myrtle.TableColumnAlignments(map[int]myrtle.TableColumnAlignment{
				0: myrtle.TableColumnAlignmentLeft,
				1: myrtle.TableColumnAlignmentRight,
				2: myrtle.TableColumnAlignmentRight,
			}),
		).
		AddBarChart("Message volume by channel", []myrtle.BarChartItem{
			{Label: "Email", Value: "68%", Percent: 68},
			{Label: "SMS", Value: "21%", Percent: 21},
			{Label: "Push", Value: "11%", Percent: 11},
		}).
		AddSparkline("Trend", "Signups", "1,204", []int{8, 12, 9, 14, 18, 16, 20}, myrtle.SparklineDelta("+8%"), myrtle.SparklineDeltaSemantic(myrtle.StatDeltaSemanticPositive)).
		AddStackedBar("Stage composition", []myrtle.StackedBarRow{
			{Label: "Acquisition", Segments: []myrtle.StackedBarSegment{{Label: "Email", Percent: 58, Value: "58%"}, {Label: "SMS", Percent: 24, Value: "24%"}, {Label: "Push", Percent: 18, Value: "18%"}}},
			{Label: "Activation", Segments: []myrtle.StackedBarSegment{{Label: "Email", Percent: 42, Value: "42%"}, {Label: "SMS", Percent: 33, Value: "33%"}, {Label: "Push", Percent: 25, Value: "25%"}}},
		}, myrtle.StackedBarTotal("Total", "120k")).
		AddProgress("Release checklist", []myrtle.ProgressItem{{Label: "Schema migration", Percent: 100}, {Label: "API deploy", Percent: 76}, {Label: "Client rollout", Percent: 48}}).
		AddDistribution("Latency buckets (ms)", []myrtle.DistributionBucket{{Label: "0-50", Count: 62}, {Label: "51-100", Count: 44}, {Label: "101-200", Count: 21}, {Label: "200+", Count: 8}}).
		AddPriceSummary("Commercial summary", []myrtle.PriceLine{{Label: "Subtotal", Value: "$120.00"}, {Label: "Discount", Value: "-$12.00"}, {Label: "Tax", Value: "$8.64"}}, "Total", "$116.64").
		AddKeyValue("Account summary", []myrtle.KeyValuePair{
			{Key: "Plan", Value: "Business"},
			{Key: "Workspace", Value: "acme-prod"},
			{Key: "Renewal", Value: "2026-04-01"},
		}).
		AddQuote("Migrating to composable blocks reduced template maintenance time by 60%.", "Email Engineering Team").
		AddSummaryCard("Summary card variant", "A concise summary card with optional footer text.", "Last updated just now").
		AddTimeline(
			"Recent events",
			[]myrtle.TimelineItem{{Time: "09:07", Title: "Detected", Detail: "Queue latency increased."}, {Time: "09:18", Title: "Mitigated", Detail: "Workers scaled."}, {Time: "09:42", Title: "Resolved", Detail: "Latency restored."}},
			myrtle.TimelineAggregateHeader("3 events · 1 currently active"),
			myrtle.TimelineCurrentIndex(1),
		).
		AddColumns(
			func(column *myrtle.ColumnBuilder) {
				column.AddCallout(myrtle.CalloutTypeSuccess, "All systems nominal", "No active incidents and queue lag is within normal range.").
					AddCallout(myrtle.CalloutTypeWarning, "Action recommended", "Rotate API keys older than 90 days.")
			},
			func(column *myrtle.ColumnBuilder) {
				column.AddCallout(myrtle.CalloutTypeError, "Failed webhooks", "2 endpoints failed retries in the last 24h.").
					AddImage("https://via.placeholder.com/520x240.png?text=Myrtle+Preview", "Myrtle chart preview")
			},
			myrtle.ColumnsWidths(50, 50),
		).
		AddAttachment("weekly-report.pdf", "PDF · 312 KB", "https://example.com/reports/weekly", "Download").
		AddEmptyState("No pending approvals", "Everything is up to date for now.", "Create approval rule", "https://example.com/rules/new").
		AddHeading("Markdown escape hatch", myrtle.HeadingLevel(2)).
		AddFreeMarkdown("Use **free markdown** for one-off sections where rich text is convenient.\n\n- Supports lists\n- Supports emphasis\n- Supports links like [Myrtle](https://github.com/gzuidhof/myrtle)").
		AddAction("Continue in the app:", "Open Myrtle", "https://example.com/app").
		AddFooterLinks([]myrtle.FooterLink{{Label: "Docs", URL: "https://github.com/gzuidhof/myrtle"}, {Label: "Support", URL: "https://github.com/gzuidhof/myrtle/discussions"}, {Label: "Status", URL: "https://example.com/status"}}, "You can manage what you receive in preferences.").
		AddLegal("Myrtle Inc.", "123 Market St, San Francisco, CA", "https://example.com/preferences", "https://example.com/unsubscribe").
		Build(), nil
}
