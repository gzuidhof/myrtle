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
		myrtle.WithStyles(theme.Styles{ColorPrimary: "#6d28d9", ColorSecondary: "#a855f7"}),
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
			ImageURL: "/assets/hero.png",
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
		AddSpacer(myrtle.SpacerSize(10)).
		AddHeading("Buttons", myrtle.HeadingLevel(2)).
		AddButton("Primary CTA", "https://example.com/primary", myrtle.ButtonTone(myrtle.ButtonTonePrimary)).
		AddButton("Centered CTA", "https://example.com/center", myrtle.ButtonTone(myrtle.ButtonTonePrimary), myrtle.ButtonAlign(myrtle.ButtonAlignmentCenter)).
		AddButton("Secondary CTA", "https://example.com/secondary", myrtle.ButtonTone(myrtle.ButtonToneSecondary)).
		AddButton("Danger CTA", "https://example.com/danger", myrtle.ButtonTone(myrtle.ButtonToneDanger)).
		AddButton("Outline CTA", "https://example.com/outline", myrtle.ButtonStyle(myrtle.ButtonStyleOutline)).
		AddButton("Ghost CTA", "https://example.com/ghost", myrtle.ButtonStyle(myrtle.ButtonStyleGhost)).
		AddButton("Danger ghost", "https://example.com/danger-ghost", myrtle.ButtonTone(myrtle.ButtonToneDanger), myrtle.ButtonStyle(myrtle.ButtonStyleGhost)).
		AddButton("Full-width primary", "https://example.com/full", myrtle.ButtonTone(myrtle.ButtonTonePrimary), myrtle.ButtonFullWidth(true)).
		AddButtonGroup([]myrtle.ButtonGroupButton{{Label: "Approve", URL: "https://example.com/approve", Tone: myrtle.ButtonTonePrimary}, {Label: "Review", URL: "https://example.com/review", Tone: myrtle.ButtonToneSecondary}, {Label: "Delete", URL: "https://example.com/delete", Tone: myrtle.ButtonToneDanger}}, myrtle.ButtonGroupAlign(myrtle.ButtonAlignmentCenter), myrtle.ButtonGroupJoined(true)).
		AddButtonGroup([]myrtle.ButtonGroupButton{{Label: "Primary", URL: "https://example.com/primary-group", Tone: myrtle.ButtonTonePrimary}, {Label: "Outline", URL: "https://example.com/outline-group", Style: myrtle.ButtonStyleOutline}, {Label: "Ghost", URL: "https://example.com/ghost-group", Tone: myrtle.ButtonToneDanger, Style: myrtle.ButtonStyleGhost}}, myrtle.ButtonGroupGap(12), myrtle.ButtonGroupStackOnMobile(true), myrtle.ButtonGroupFullWidthOnMobile(true)).
		AddColumns(
			func(column *myrtle.ColumnBuilder) {
				column.AddHeading("Quick links", myrtle.HeadingLevel(3)).
					AddList([]string{"Open dashboard", "Review incidents", "Manage preferences"}, false).
					AddButton("Open dashboard", "https://example.com/dashboard")
			},
			func(column *myrtle.ColumnBuilder) {
				column.AddHeading("Verification", myrtle.HeadingLevel(3)).
					Add(myrtle.VerificationCodeBlock{Label: "One-time code", Value: "483920"}).
					AddText("Use this code to authorize billing changes.").
					AddButton("Open security", "https://example.com/security")
			},
			myrtle.ColumnsWidths(55, 45),
		).
		AddDivider().
		AddDividerStyled(myrtle.DividerStyle(myrtle.DividerVariantDashed), myrtle.DividerThickness(2)).
		AddDividerStyled(myrtle.DividerStyle(myrtle.DividerVariantDotted), myrtle.DividerThickness(2), myrtle.DividerInset(24)).
		AddHeading("Messaging blocks", myrtle.HeadingLevel(2)).
		AddMessage(myrtle.MessageBlock{SenderName: "Alex Johnson", SenderHandle: "@alex", AvatarURL: "/assets/avatar1.png", LogoAlt: "Alex Johnson avatar", LogoHref: "https://example.com/messages/42", Subject: "Build pipeline notifications", Preview: "Can we switch alerts to batch mode for low-priority events?", SentAt: "2m ago", Platform: "Myrtle Chat", URL: "https://example.com/messages/42", ActionLabel: "Open thread", ActionURL: "https://example.com/messages/42"}).
		AddMessageDigest(
			[]myrtle.MessageBlock{
				{SenderName: "Maya", SenderHandle: "@maya", AvatarURL: "https://i.pravatar.cc/80?img=5", LogoAlt: "Maya avatar", LogoHref: "https://example.com/messages/43", Subject: "**Launch update**", Preview: "Can you check [the draft](https://example.com/draft)?", SentAt: "5m ago", Platform: "Myrtle Chat", URL: "https://example.com/messages/43"},
				{SenderName: "Nina", SenderHandle: "@nina", Subject: "Incident summary", Preview: "Looks good overall, adding one more metric.", SentAt: "11m ago", Platform: "Myrtle Chat", URL: "https://example.com/messages/46"},
				{SenderName: "Ben", SenderHandle: "@ben", Preview: "No subject message to test compact rendering.", SentAt: "20m ago", Platform: "Myrtle Chat", URL: "https://example.com/messages/44"},
			},
			myrtle.MessageDigestTitle("Team inbox"),
			myrtle.MessageDigestSubtitle("Recent direct messages and mentions"),
			myrtle.MessageDigestFooter("[Open inbox](https://example.com/messages)"),
			myrtle.MessageDigestMaxItems(3),
		).
		AddHeading("Advanced layout blocks", myrtle.HeadingLevel(2)).
		AddSection(
			[]myrtle.Block{
				myrtle.TextBlock{Text: "Section body with nested button for grouped rendering validation."},
				myrtle.ButtonBlock{Label: "Open section", URL: "https://example.com/section", Style: myrtle.ButtonStyleOutline},
			},
			myrtle.SectionTitle("Section block"),
			myrtle.SectionSubtitle("Border + padding variant"),
			myrtle.SectionPadding(20),
			myrtle.SectionBorder(true),
		).
		AddGrid(
			[]myrtle.GridItem{
				{Blocks: []myrtle.Block{myrtle.HeadingBlock{Text: "Grid item 1", Level: 4}, myrtle.TextBlock{Text: "Nested content one."}}},
				{Blocks: []myrtle.Block{myrtle.HeadingBlock{Text: "Grid item 2", Level: 4}, myrtle.TextBlock{Text: "Nested content two."}}},
				{Blocks: []myrtle.Block{myrtle.HeadingBlock{Text: "Grid item 3", Level: 4}, myrtle.TextBlock{Text: "Wrap behavior check."}}},
				{Blocks: []myrtle.Block{myrtle.HeadingBlock{Text: "Grid item 4", Level: 4}, myrtle.TextBlock{Text: "Border + spacing stress."}}},
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
		AddStatsRow("Core metrics", []myrtle.StatItem{{Label: "Delivery", Value: "99.8%", Delta: "+0.3%", DeltaSemantic: myrtle.StatDeltaSemanticPositive}, {Label: "Open rate", Value: "42.1%", Delta: "+1.2%", DeltaSemantic: myrtle.StatDeltaSemanticPositive}, {Label: "Bounce", Value: "0.9%", Delta: "-0.1%", DeltaSemantic: myrtle.StatDeltaSemanticNegative}}).
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
					AddImage("/assets/chart-preview.png", "Myrtle chart preview")
			},
			myrtle.ColumnsWidths(50, 50),
		).
		AddAttachment("weekly-report.pdf", "PDF · 312 KB", "https://example.com/reports/weekly", "Download").
		AddEmptyState("No pending approvals", "Everything is up to date for now.", "Create approval rule", "https://example.com/rules/new").
		AddHeading("Markdown escape hatch", myrtle.HeadingLevel(2)).
		AddFreeMarkdown("Use **free markdown** for one-off sections where rich text is convenient.\n\n- Supports lists\n- Supports emphasis\n- Supports links like [Myrtle](https://github.com/gzuidhof/myrtle)").
		AddFooterLinks([]myrtle.FooterLink{{Label: "Docs", URL: "https://github.com/gzuidhof/myrtle"}, {Label: "Support", URL: "https://github.com/gzuidhof/myrtle/discussions"}, {Label: "Status", URL: "https://example.com/status"}}, "You can manage what you receive in preferences.").
		AddLegal("Myrtle Inc.", "123 Market St, San Francisco, CA", "https://example.com/preferences", "https://example.com/unsubscribe").
		Build(), nil
}
