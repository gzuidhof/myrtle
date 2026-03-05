package example

import (
	"fmt"

	"github.com/gzuidhof/myrtle"
	"github.com/gzuidhof/myrtle/theme"
	defaulttheme "github.com/gzuidhof/myrtle/theme/default"
)

func MonsterEmail() (*myrtle.Email, error) {
	return MonsterEmailWithTheme(defaulttheme.New())
}

func MonsterEmailWithTheme(selectedTheme theme.Theme) (*myrtle.Email, error) {
	return monsterEmailWithThemeAndStyles(selectedTheme, monsterLightStyles(), theme.DirectionLTR)
}

func MonsterDarkModeEmail() (*myrtle.Email, error) {
	return MonsterDarkModeEmailWithTheme(defaulttheme.New())
}

func MonsterDarkModeEmailWithTheme(selectedTheme theme.Theme) (*myrtle.Email, error) {
	return monsterEmailWithThemeAndStyles(selectedTheme, monsterDarkStyles(), theme.DirectionLTR)
}

func MonsterRTLEmail() (*myrtle.Email, error) {
	return MonsterRTLEmailWithTheme(defaulttheme.New())
}

func MonsterRTLEmailWithTheme(selectedTheme theme.Theme) (*myrtle.Email, error) {
	return monsterEmailWithThemeAndStyles(selectedTheme, monsterLightStyles(), theme.DirectionRTL)
}

func monsterEmailWithThemeAndStyles(selectedTheme theme.Theme, styles theme.Styles, direction theme.Direction) (*myrtle.Email, error) {
	if selectedTheme == nil {
		selectedTheme = defaulttheme.New()
	}

	metricsRows := [][]string{
		{"Delivered", "126,842", "+2.1%"},
		{"Open rate", "41.8%", "+0.7pp"},
		{"Click rate", "4.2%", "+0.3pp"},
		{"Unsubscribes", "0.18%", "-0.02pp"},
	}

	monthNames := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	axisLabels24Months := make([]string, 0, 24)
	series24Months := []myrtle.VerticalBarChartSeries{
		{Key: "new", Label: "New", Color: "#1d4ed8", ValueLabelColor: "#ffffff", Values: make([]float64, 0, 24)},
		{Key: "expansion", Label: "Expansion", Color: "#3b82f6", ValueLabelColor: "#ffffff", Values: make([]float64, 0, 24)},
		{Key: "churn", Label: "Churn", Color: "#ef4444", ValueLabelColor: "#ffffff", Values: make([]float64, 0, 24)},
	}
	for i := 0; i < 24; i++ {
		year := 2024
		if i >= 12 {
			year = 2025
		}
		axisLabels24Months = append(axisLabels24Months, fmt.Sprintf("%s '%02d", monthNames[i%12], year%100))
		series24Months[0].Values = append(series24Months[0].Values, 26+float64((i*5)%18))
		series24Months[1].Values = append(series24Months[1].Values, 10+float64((i*3)%11))
		series24Months[2].Values = append(series24Months[2].Values, -float64(6+(i*4)%10))
	}

	return myrtle.NewBuilder(
		selectedTheme,
		myrtle.WithStyles(styles),
		myrtle.WithDirection(direction),
	).
		WithoutHeader().
		WithPreheader("Weekly operations snapshot covering delivery, incidents, billing, and team activity").
		WithHeader(myrtle.HeadingBlock{Text: "Myrtle Weekly Operations Snapshot", Level: 1}, myrtle.HeaderPlacement(myrtle.HeaderPlacementOutside)).
		Add(myrtle.HeroBlock{
			Eyebrow:  "Ops Digest",
			Title:    "Delivery healthy, onboarding conversion up, and no active incidents",
			Body:     "Use this synthetic but realistic snapshot to validate layouts while previewing day-to-day operational content.",
			CTALabel: "Open dashboard",
			CTAURL:   "https://app.example.com/dashboard",
			ImageURL: "/assets/hero.png",
			ImageAlt: "Operations dashboard summary",
		}).
		AddBadge(myrtle.ToneInfo, "Info badge variant").
		AddBadge(myrtle.ToneSuccess, "Success badge variant").
		AddBadge(myrtle.ToneWarning, "Warning badge variant").
		AddBadge(myrtle.ToneDanger, "Error badge variant").
		AddHeading("Everything in one place", myrtle.HeadingLevel(1)).
		AddText("This email intentionally demonstrates all built-in blocks and several composition patterns.").
		AddCallout(myrtle.ToneInfo, "Overview", "Scan this message to see all components rendered together.", myrtle.CalloutStyle(myrtle.CalloutVariantSoft)).
		AddCallout(myrtle.ToneWarning, "Variant", "Outline style for warning callouts.", myrtle.CalloutStyle(myrtle.CalloutVariantOutline)).
		AddCallout(myrtle.ToneDanger, "Variant", "Solid style for critical alerts.", myrtle.CalloutStyle(myrtle.CalloutVariantSolid)).
		AddSpacer(myrtle.SpacerSize(10)).
		AddHeading("Buttons", myrtle.HeadingLevel(2)).
		AddButton("Review campaign", "https://app.example.com/campaigns/weekly", myrtle.ButtonTone(myrtle.TonePrimary)).
		AddButton("Open incident queue", "https://app.example.com/incidents", myrtle.ButtonTone(myrtle.TonePrimary), myrtle.ButtonAlign(myrtle.ButtonAlignmentCenter)).
		AddButton("View billing", "https://app.example.com/billing", myrtle.ButtonTone(myrtle.ToneSecondary), myrtle.ButtonAlign(myrtle.ButtonAlignmentEnd)).
		AddButton("Pause send", "https://app.example.com/sends/pause", myrtle.ButtonTone(myrtle.ToneDanger)).
		AddButton("Export CSV", "https://app.example.com/reports/export", myrtle.ButtonStyle(myrtle.ButtonStyleOutline)).
		AddButton("Snooze alerts", "https://app.example.com/alerts/snooze", myrtle.ButtonStyle(myrtle.ButtonStyleGhost)).
		AddButton("Disable integration", "https://app.example.com/integrations/disable", myrtle.ButtonTone(myrtle.ToneDanger), myrtle.ButtonStyle(myrtle.ButtonStyleGhost)).
		AddButton("Open full report", "https://app.example.com/reports/weekly", myrtle.ButtonTone(myrtle.TonePrimary), myrtle.ButtonFullWidth(true)).
		AddButtonGroup([]myrtle.ButtonGroupButton{{Label: "Approve send", URL: "https://app.example.com/sends/approve", Tone: myrtle.TonePrimary}, {Label: "Review changes", URL: "https://app.example.com/sends/review", Tone: myrtle.ToneSecondary}, {Label: "Cancel", URL: "https://app.example.com/sends/cancel", Tone: myrtle.ToneDanger}}, myrtle.ButtonGroupAlign(myrtle.ButtonAlignmentCenter), myrtle.ButtonGroupJoined(true)).
		AddButtonGroup([]myrtle.ButtonGroupButton{{Label: "Primary", URL: "https://example.com/primary-group", Tone: myrtle.TonePrimary}, {Label: "Outline", URL: "https://example.com/outline-group", Style: myrtle.ButtonStyleOutline}, {Label: "Ghost", URL: "https://example.com/ghost-group", Tone: myrtle.ToneDanger, Style: myrtle.ButtonStyleGhost}}, myrtle.ButtonGroupGap(12), myrtle.ButtonGroupStackOnMobile(true), myrtle.ButtonGroupFullWidthOnMobile(true)).
		AddColumns(
			myrtle.NewGroup().
				AddHeading("Quick links", myrtle.HeadingLevel(3)).
				AddList([]string{"Open dashboard", "Review incidents", "Manage preferences"}, false).
				AddButton("Open dashboard", "https://example.com/dashboard"),
			myrtle.NewGroup().
				AddHeading("Verification", myrtle.HeadingLevel(3)).
				Add(myrtle.VerificationCodeBlock{Label: "One-time code", Value: "483920"}).
				AddText("Use this code to authorize billing changes.").
				AddButton("Open security", "https://example.com/security"),
			myrtle.ColumnsWidths(55, 45),
		).
		AddDivider().
		AddDivider(myrtle.DividerStyle(myrtle.DividerVariantDashed), myrtle.DividerThickness(2)).
		AddDivider(myrtle.DividerStyle(myrtle.DividerVariantDotted), myrtle.DividerThickness(2), myrtle.DividerInset(24)).
		AddHeading("Messaging blocks", myrtle.HeadingLevel(2)).
		AddMessage(myrtle.MessageBlock{SenderName: "Alex Johnson", SenderHandle: "@alex", AvatarURL: "/assets/avatar1.png", LogoAlt: "Alex Johnson avatar", LogoHref: "https://app.example.com/messages/42", Subject: "Alert policy update for low-priority webhooks", Preview: "Can we batch non-critical webhook failures into a 15-minute digest?", SentAt: "2m ago", Platform: "Myrtle Chat", URL: "https://app.example.com/messages/42", ActionLabel: "Open thread", ActionURL: "https://app.example.com/messages/42"}).
		AddMessageDigest(
			[]myrtle.MessageBlock{
				{SenderName: "Maya", SenderHandle: "@maya", AvatarURL: "https://i.pravatar.cc/80?img=5", LogoAlt: "Maya avatar", LogoHref: "https://example.com/messages/43", Subject: "**Launch update**", Preview: "Can you check [the draft](https://example.com/draft)?", SentAt: "5m ago", Platform: "Myrtle Chat", URL: "https://example.com/messages/43"},
				{SenderName: "Nina", SenderHandle: "@nina", Subject: "Incident summary", Preview: "Looks good overall, adding one more metric.", SentAt: "11m ago", Platform: "Myrtle Chat", URL: "https://example.com/messages/46"},
				{SenderName: "Ben", SenderHandle: "@ben", Preview: "No subject message to test compact rendering.", SentAt: "20m ago", Platform: "Myrtle Chat", URL: "https://example.com/messages/44"},
			},
			myrtle.MessageDigestTitle("Team inbox"),
			myrtle.MessageDigestSubtitle("Recent engineering and operations discussions"),
			myrtle.MessageDigestFooter("[Open inbox](https://example.com/messages)"),
			myrtle.MessageDigestMaxItems(3),
		).
		AddHeading("Advanced layout blocks", myrtle.HeadingLevel(2)).
		AddPanel(
			myrtle.NewGroup().
				AddText("Panel body with nested button for grouped rendering validation.").
				AddButton("Open panel", "https://example.com/panel", myrtle.ButtonStyle(myrtle.ButtonStyleOutline)),
			myrtle.PanelTitle("Panel block"),
			myrtle.PanelSubtitle("Border + padding variant"),
			myrtle.PanelPadding(20),
			myrtle.PanelBorder(true),
		).
		AddGrid(
			[]myrtle.GridItem{
				myrtle.GridItemGroup(myrtle.NewGroup().AddHeading("Grid item 1", myrtle.HeadingLevel(4)).AddText("Nested content one.")),
				myrtle.GridItemGroup(myrtle.NewGroup().AddHeading("Grid item 2", myrtle.HeadingLevel(4)).AddText("Nested content two.")),
				myrtle.GridItemGroup(myrtle.NewGroup().AddHeading("Grid item 3", myrtle.HeadingLevel(4)).AddText("Wrap behavior check.")),
				myrtle.GridItemGroup(myrtle.NewGroup().AddHeading("Grid item 4", myrtle.HeadingLevel(4)).AddText("Border + spacing stress.")),
			},
			myrtle.GridColumns(3),
			myrtle.GridGap(12),
			myrtle.GridBorder(true),
		).
		AddCardList(
			[]myrtle.CardItem{
				{Title: "Deploy complete", Subtitle: "Production", Body: "No customer impact detected.", URL: "https://example.com/deploy/1", CTALabel: "View"},
				{Title: "Billing updated", Subtitle: "Finance", Body: "Invoice #8241 has been paid.", URL: "https://example.com/billing/8241", CTALabel: "Open"},
				{Title: "Weekly report", Body: "Read the latest metrics and highlights.", URL: "https://example.com/reports/weekly"},
			},
			myrtle.CardListColumns(2),
			myrtle.CardListGap(12),
			myrtle.CardListBorder(true),
		).
		AddTiles(
			[]myrtle.TileEntry{
				{Content: "🚀", Title: "Launch", Subtitle: "Ready", URL: "https://example.com/launch", Variant: myrtle.TileVariantHighlight},
				{Content: "12", Title: "Queued", Subtitle: "Jobs", Variant: myrtle.TileVariantWarning},
				{Content: "✅", Title: "Healthy", Subtitle: "All systems", Variant: myrtle.TileVariantSuccess},
				{Content: "!", Title: "Critical", Subtitle: "Manual action", Variant: myrtle.TileVariantCritical},
			},
			myrtle.TilesColumns(4),
			myrtle.TilesBorder(true),
		).
		AddHeading("Metrics and highlights", myrtle.HeadingLevel(2)).
		AddStatsRow("Core metrics", []myrtle.StatItem{{Label: "Delivery", Value: "99.82%", Delta: "+0.12pp", DeltaSemantic: myrtle.StatDeltaSemanticPositive}, {Label: "Open rate", Value: "41.8%", Delta: "+0.7pp", DeltaSemantic: myrtle.StatDeltaSemanticPositive}, {Label: "Complaint rate", Value: "0.03%", Delta: "-0.01pp", DeltaSemantic: myrtle.StatDeltaSemanticNegative}}).
		AddTable(
			[]string{"Metric", "Value", "Delta"},
			metricsRows,
			myrtle.TableCompact(true),
			myrtle.TableZebraRows(true),
			myrtle.TableRightAlignNumericColumns(true),
			myrtle.TableEmphasizeTotalRow(true),
			myrtle.TableColumnAlignments(map[int]myrtle.TableColumnAlignmentValue{
				0: myrtle.TableColumnAlignmentStart,
				1: myrtle.TableColumnAlignmentEnd,
				2: myrtle.TableColumnAlignmentEnd,
			}),
		).
		AddHorizontalBarChart("Message volume by channel", []myrtle.HorizontalBarChartItem{
			{Label: "Email", Value: "68%", Percent: 68, Color: "#2563eb"},
			{Label: "SMS", Value: "21%", Percent: 21, Color: "#7c3aed"},
			{Label: "Push", Value: "11%", Percent: 11, Color: "#0ea5e9"},
		}).
		AddSparkline("Trend", "Signups", "1,204", []int{8, 12, 9, 14, 18, 16, 20}, myrtle.SparklineDelta("+8%"), myrtle.SparklineDeltaSemantic(myrtle.StatDeltaSemanticPositive)).
		AddStackedBar("Stage composition", []myrtle.StackedBarRow{
			{Label: "Acquisition", Segments: []myrtle.StackedBarSegment{{Label: "Email", Percent: 58, Value: "58%", Color: "#2563eb"}, {Label: "SMS", Percent: 24, Value: "24%", Color: "#7c3aed"}, {Label: "Push", Percent: 18, Value: "18%", Color: "#0ea5e9"}}},
			{Label: "Activation", Segments: []myrtle.StackedBarSegment{{Label: "Email", Percent: 42, Value: "42%", Color: "#16a34a"}, {Label: "SMS", Percent: 33, Value: "33%", Color: "#f59e0b"}, {Label: "Push", Percent: 25, Value: "25%", Color: "#dc2626"}}},
		}, myrtle.StackedBarTotal("Total", "120k")).
		AddProgress("Release checklist", []myrtle.ProgressItem{{Label: "Schema migration", Percent: 100, Color: "#16a34a"}, {Label: "API deploy", Percent: 76, Color: "#f59e0b"}, {Label: "Client rollout", Percent: 48, Color: "#dc2626"}}).
		AddDistribution("Latency buckets (ms)", []myrtle.DistributionBucket{{Label: "0-50", Count: 62, Color: "#16a34a"}, {Label: "51-100", Count: 44, Color: "#0ea5e9"}, {Label: "101-200", Count: 21, Color: "#f59e0b"}, {Label: "200+", Count: 8, Color: "#dc2626"}}).
		AddPriceSummary("Commercial summary", []myrtle.PriceLine{{Label: "Email volume (1.2M)", Value: "$480.00"}, {Label: "Pro support", Value: "$149.00"}, {Label: "Credit", Value: "-$75.00"}}, "Estimated total", "$554.00").
		AddKeyValue("Account summary", []myrtle.KeyValuePair{
			{Key: "Plan", Value: "Scale Annual"},
			{Key: "Workspace", Value: "acme-production"},
			{Key: "Renewal", Value: "2026-04-15"},
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
			myrtle.NewGroup().
				AddCallout(myrtle.ToneSuccess, "All systems nominal", "No active incidents and queue lag is within normal range.").
				AddCallout(myrtle.ToneWarning, "Action recommended", "Rotate API keys older than 90 days."),
			myrtle.NewGroup().
				AddCallout(myrtle.ToneDanger, "Failed webhooks", "2 endpoints failed retries in the last 24h."),
			myrtle.ColumnsWidths(50, 50),
		).
		AddAttachment("weekly-report.pdf", "PDF · 312 KB", "https://example.com/reports/weekly", "Download").
		AddEmptyState("No pending approvals", "Everything is up to date for now.", "Create approval rule", "https://example.com/rules/new").
		AddHeading("InsetMode none examples", myrtle.HeadingLevel(2)).
		AddText("Representative components rendered with InsetModeNone for visual parity checks.").
		AddCallout(myrtle.ToneInfo, "Callout (none)", "Callout rendered edge-to-edge in the main container.", myrtle.CalloutInsetMode(myrtle.InsetModeNone)).
		AddPanel(
			myrtle.NewGroup().
				AddText("Panel content with no inset.").
				AddButton("Open panel", "https://example.com/panel-none", myrtle.ButtonStyle(myrtle.ButtonStyleOutline)),
			myrtle.PanelTitle("Panel (none)"),
			myrtle.PanelInsetMode(myrtle.InsetModeNone),
		).
		AddGrid(
			[]myrtle.GridItem{
				myrtle.GridItemGroup(myrtle.NewGroup().AddHeading("Grid none 1", myrtle.HeadingLevel(4)).AddText("No-inset grid item.")),
				myrtle.GridItemGroup(myrtle.NewGroup().AddHeading("Grid none 2", myrtle.HeadingLevel(4)).AddText("No-inset grid item.")),
			},
			myrtle.GridColumns(2),
			myrtle.GridGap(0),
			myrtle.GridBorder(true),
			myrtle.GridInsetMode(myrtle.InsetModeNone),
		).
		AddCardList(
			[]myrtle.CardItem{{Title: "Card none", Body: "Card list no-inset variant.", URL: "https://example.com/card-none", CTALabel: "View"}, {Title: "Card none 2", Body: "Second card variant."}},
			myrtle.CardListColumns(2),
			myrtle.CardListGap(0),
			myrtle.CardListBorder(true),
			myrtle.CardListInsetMode(myrtle.InsetModeNone),
		).
		AddTiles(
			[]myrtle.TileEntry{{Content: "⚙️", Title: "Tiles none", Subtitle: "Edge-to-edge", Variant: myrtle.TileVariantHighlight}, {Content: "24", Title: "Checks", Subtitle: "No inset", Variant: myrtle.TileVariantWarning}},
			myrtle.TilesColumns(2),
			myrtle.TilesBorder(true),
			myrtle.TilesInsetMode(myrtle.InsetModeNone),
		).
		AddHorizontalBarChart("Horizontal bar (none)", []myrtle.HorizontalBarChartItem{{Label: "Email", Value: "68%", Percent: 68}, {Label: "SMS", Value: "21%", Percent: 21}}, myrtle.HorizontalBarChartInsetMode(myrtle.InsetModeNone)).
		AddVerticalBarChart(
			axisLabels24Months,
			series24Months,
			myrtle.VerticalBarChartTitle("Vertical bar 24 months (none)"),
			myrtle.VerticalBarChartHeight(170),
			myrtle.VerticalBarChartColumnGap(1),
			myrtle.VerticalBarChartAxisShowCategoryLabels(false),
			myrtle.VerticalBarChartAxisShowYTicks(true),
			myrtle.VerticalBarChartLegendPlacement(myrtle.VerticalBarChartLegendBottom),
			myrtle.VerticalBarChartInsetMode(myrtle.InsetModeNone),
		).
		AddVerticalBarChart(
			axisLabels24Months,
			series24Months,
			myrtle.VerticalBarChartTitle("Vertical bar 24 months with in-bar values (none)"),
			myrtle.VerticalBarChartHeight(170),
			myrtle.VerticalBarChartColumnGap(1),
			myrtle.VerticalBarChartAxisShowCategoryLabels(false),
			myrtle.VerticalBarChartAxisShowYTicks(true),
			myrtle.VerticalBarChartLegendPlacement(myrtle.VerticalBarChartLegendBottom),
			myrtle.VerticalBarChartValueLabelsOption(myrtle.VerticalBarChartValueLabels{Show: true, MinSegmentHeight: 14, Color: "#ffffff"}),
			myrtle.VerticalBarChartInsetMode(myrtle.InsetModeNone),
		).
		AddSparkline("Sparkline (none)", "Throughput", "1,284", []int{8, 12, 9, 14, 18, 16, 20}, myrtle.SparklineDelta("+8%"), myrtle.SparklineInsetMode(myrtle.InsetModeNone)).
		AddStackedBar("Stacked bar (none)", []myrtle.StackedBarRow{{Label: "Acquisition", Segments: []myrtle.StackedBarSegment{{Label: "Email", Percent: 58, Value: "58%"}, {Label: "SMS", Percent: 24, Value: "24%"}, {Label: "Push", Percent: 18, Value: "18%"}}}}, myrtle.StackedBarTotal("Total", "100%"), myrtle.StackedBarInsetMode(myrtle.InsetModeNone)).
		AddTable(
			[]string{"Metric", "Value"},
			[][]string{{"Messages", "128,490"}, {"Delivery", "99.82%"}},
			myrtle.TableInsetMode(myrtle.InsetModeNone),
		).
		AddMessage(
			myrtle.MessageBlock{SenderName: "Inset Bot", SenderHandle: "@inset", Subject: "Message (none)", Preview: "Message block rendered with no inset.", SentAt: "Now", Platform: "Myrtle Chat", URL: "https://example.com/messages/inset-none"},
			myrtle.MessageInsetMode(myrtle.InsetModeNone),
		).
		AddMessageDigest(
			[]myrtle.MessageBlock{{SenderName: "Inset Bot", SenderHandle: "@inset", Subject: "Digest (none)", Preview: "Digest row no-inset variant.", SentAt: "1m", Platform: "Myrtle Chat", URL: "https://example.com/messages/inset-digest"}},
			myrtle.MessageDigestTitle("Digest (none)"),
			myrtle.MessageDigestInsetMode(myrtle.InsetModeNone),
		).
		AddProgress("Progress (none)", []myrtle.ProgressItem{{Label: "Schema migration", Percent: 100}, {Label: "Client rollout", Percent: 48}}, myrtle.ProgressInsetMode(myrtle.InsetModeNone)).
		AddDistribution("Distribution (none)", []myrtle.DistributionBucket{{Label: "0-50", Count: 62}, {Label: "51-100", Count: 44}}, myrtle.DistributionInsetMode(myrtle.InsetModeNone)).
		AddAttachment("inset-none.txt", "TXT · 4 KB", "https://example.com/inset-none", "Download", myrtle.AttachmentInsetMode(myrtle.InsetModeNone)).
		AddEmptyState("Empty state (none)", "No queued tasks.", "Create task", "https://example.com/tasks/new", myrtle.EmptyStateInsetMode(myrtle.InsetModeNone)).
		AddPriceSummary("Pricing (none)", []myrtle.PriceLine{{Label: "Subtotal", Value: "$84.00"}, {Label: "Tax", Value: "$6.72"}}, "Total", "$90.72", myrtle.PriceSummaryInsetMode(myrtle.InsetModeNone)).
		AddTimeline("Timeline (none)", []myrtle.TimelineItem{{Time: "09:07 UTC", Title: "Detected", Detail: "Webhook processing latency exceeded 1.2s."}, {Time: "09:18 UTC", Title: "Mitigated", Detail: "Worker pool scaled from 12 to 20 instances."}}, myrtle.TimelineInsetMode(myrtle.InsetModeNone)).
		AddHeading("Markdown escape hatch", myrtle.HeadingLevel(2)).
		AddFreeMarkdown("Use **free markdown** for one-off sections where rich text is convenient.\n\n- Supports lists\n- Supports emphasis\n- Supports links like [Myrtle](https://github.com/gzuidhof/myrtle)").
		AddFooterLinks([]myrtle.FooterLink{{Label: "Docs", URL: "https://github.com/gzuidhof/myrtle"}, {Label: "Support", URL: "https://github.com/gzuidhof/myrtle/discussions"}, {Label: "Status", URL: "https://example.com/status"}}, "You can manage what you receive in preferences.").
		AddLegal("Myrtle Inc.", "Dam Square 1, 1012 JS Amsterdam, Netherlands", "https://example.com/preferences", "https://example.com/unsubscribe").
		Build(), nil
}

func monsterLightStyles() theme.Styles {
	return theme.Styles{
		ColorPrimary:   "#6d28d9",
		ColorSecondary: "#a855f7",
	}
}

func monsterDarkStyles() theme.Styles {
	styles := theme.DefaultDarkModeStyles()
	styles.ColorPrimary = "#8b5cf6"
	styles.ColorSecondary = "#c084fc"

	return styles
}
