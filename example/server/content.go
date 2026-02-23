package server

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gzuidhof/myrtle"
	"github.com/gzuidhof/myrtle/theme"
	defaulttheme "github.com/gzuidhof/myrtle/theme/default"
	"github.com/gzuidhof/myrtle/theme/flat"
)

func buildExampleEmail(name string, selectedTheme theme.Theme) (*myrtle.Email, error) {
	for _, emailBuilder := range exampleEmails {
		if emailBuilder.Name == name {
			return emailBuilder.Build(selectedTheme)
		}
	}

	return nil, fmt.Errorf("unknown example email: %s", name)
}

func buildBlockEmail(name string, selectedTheme theme.Theme) (*myrtle.Email, error) {
	builder := myrtle.NewBuilder(selectedTheme).
		WithoutHeader()

	switch name {
	case "text":
		builder.AddText("This is a text block.")
	case "heading":
		builder.AddHeading("This is a heading", myrtle.HeadingLevel(2))
	case "spacer":
		builder.AddHeading("Fixed size", myrtle.HeadingLevel(3))
		builder.AddText("Content above spacer").AddSpacer(myrtle.SpacerSize(24)).AddText("Content below spacer")
		builder.AddHeading("Custom sizes", myrtle.HeadingLevel(3))
		builder.AddText("8px").AddSpacer(myrtle.SpacerSize(8)).AddText("12px").AddSpacer(myrtle.SpacerSize(12)).AddText("16px default").AddSpacer().AddText("24px").AddSpacer(myrtle.SpacerSize(24))
	case "list":
		builder.AddHeading("Unordered", myrtle.HeadingLevel(3))
		builder.AddList([]string{"First item", "Second item", "Third item"}, false)
		builder.AddHeading("Ordered", myrtle.HeadingLevel(3))
		builder.AddList([]string{"First item", "Second item", "Third item"}, true)
	case "key-value":
		builder.AddKeyValue("Order details", []myrtle.KeyValuePair{{Key: "Order", Value: "#4821"}, {Key: "Total", Value: "$129.00"}, {Key: "Status", Value: "Shipped"}})
	case "bar-chart":
		builder.AddHeading("Default", myrtle.HeadingLevel(3))
		builder.AddBarChart("Delivery by region", []myrtle.BarChartItem{
			{Label: "US", Value: "52%", Percent: 52},
			{Label: "EMEA", Value: "31%", Percent: 31},
			{Label: "APAC", Value: "17%", Percent: 17},
		})
		builder.AddHeading("Thicker bars", myrtle.HeadingLevel(3))
		builder.AddBarChart("Delivery by region", []myrtle.BarChartItem{
			{Label: "US", Value: "52%", Percent: 52},
			{Label: "EMEA", Value: "31%", Percent: 31},
			{Label: "APAC", Value: "17%", Percent: 17},
		}, myrtle.BarChartThickness(14))
		builder.AddHeading("Transparent background", myrtle.HeadingLevel(3))
		builder.AddBarChart("Delivery by region", []myrtle.BarChartItem{
			{Label: "US", Value: "52%", Percent: 52},
			{Label: "EMEA", Value: "31%", Percent: 31},
			{Label: "APAC", Value: "17%", Percent: 17},
		}, myrtle.BarChartTransparentBackground(true))
	case "sparkline":
		builder.AddHeading("Signups", myrtle.HeadingLevel(3))
		builder.AddSparkline("Weekly trend", "Signups", "1,204", []int{8, 12, 9, 14, 18, 16, 20}, myrtle.SparklineDelta("+8%"), myrtle.SparklineDeltaSemantic(myrtle.StatDeltaSemanticPositive))
		builder.AddHeading("Revenue", myrtle.HeadingLevel(3))
		builder.AddSparkline("Weekly trend", "Revenue", "$18.2k", []int{15, 14, 17, 16, 19, 21, 20})
		builder.AddHeading("Incidents", myrtle.HeadingLevel(3))
		builder.AddSparkline("Weekly trend", "Incidents", "12", []int{20, 18, 16, 15, 13, 12, 10}, myrtle.SparklineDelta("-2"), myrtle.SparklineDeltaSemantic(myrtle.StatDeltaSemanticNegative))
	case "stacked-bar":
		builder.AddStackedBar("Channel mix", []myrtle.StackedBarRow{
			{Label: "Acquisition", Segments: []myrtle.StackedBarSegment{{Label: "Email", Percent: 58, Value: "58%"}, {Label: "SMS", Percent: 24, Value: "24%"}, {Label: "Push", Percent: 18, Value: "18%"}}},
			{Label: "Activation", Segments: []myrtle.StackedBarSegment{{Label: "Email", Percent: 42, Value: "42%"}, {Label: "SMS", Percent: 33, Value: "33%"}, {Label: "Push", Percent: 25, Value: "25%"}}},
		}, myrtle.StackedBarTotal("Total", "120k"))
	case "progress":
		builder.AddProgress("Rollout progress", []myrtle.ProgressItem{{Label: "Schema migration", Percent: 100}, {Label: "API deploy", Percent: 76}, {Label: "Client rollout", Percent: 48}})
	case "distribution":
		builder.AddDistribution("Latency distribution (ms)", []myrtle.DistributionBucket{{Label: "0-50", Count: 62}, {Label: "51-100", Count: 44}, {Label: "101-200", Count: 21}, {Label: "200+", Count: 8}})
	case "timeline":
		builder.AddTimeline("Incident timeline", []myrtle.TimelineItem{
			{Time: "09:07", Title: "Detected", Detail: "Elevated webhook latency"},
			{Time: "09:18", Title: "Mitigation", Detail: "Queue workers scaled to 2x"},
			{Time: "09:42", Title: "Resolved", Detail: "Latency returned to baseline"},
		}, myrtle.TimelineAggregateHeader("3 events · currently mitigating"), myrtle.TimelineCurrentIndex(1))
	case "stats-row":
		builder.AddStatsRow("Weekly KPIs", []myrtle.StatItem{
			{Label: "Delivery", Value: "99.8%", Delta: "+0.3%", DeltaSemantic: myrtle.StatDeltaSemanticPositive},
			{Label: "CTR", Value: "4.1%", Delta: "+0.4%", DeltaSemantic: myrtle.StatDeltaSemanticPositive},
			{Label: "Bounces", Value: "0.9%", Delta: "-0.1%", DeltaSemantic: myrtle.StatDeltaSemanticNegative},
		})
	case "badge":
		builder.AddHeading("Info", myrtle.HeadingLevel(3))
		builder.AddBadge(myrtle.BadgeToneInfo, "Informational")
		builder.AddHeading("Success", myrtle.HeadingLevel(3))
		builder.AddBadge(myrtle.BadgeToneSuccess, "Operational")
		builder.AddHeading("Warning", myrtle.HeadingLevel(3))
		builder.AddBadge(myrtle.BadgeToneWarning, "Needs review")
		builder.AddHeading("Error", myrtle.HeadingLevel(3))
		builder.AddBadge(myrtle.BadgeToneError, "Action required")
	case "summary-card":
		builder.AddSummaryCard("Deployment complete", "The rollout to production finished successfully with no customer-facing impact.", "Updated 5 minutes ago")
	case "attachment":
		builder.AddAttachment("invoice-Feb-2026.pdf", "PDF · 284 KB", "https://example.com/invoices/feb-2026.pdf", "Download")
	case "hero":
		builder.AddHeading("Default", myrtle.HeadingLevel(3))
		builder.Add(myrtle.HeroBlock{
			Eyebrow:  "New",
			Title:    "Faster sends with Myrtle",
			Body:     "Compose transactional emails with reusable blocks and themeable output.",
			CTALabel: "Read docs",
			CTAURL:   "https://github.com/gzuidhof/myrtle",
		})
		builder.AddHeading("With image", myrtle.HeadingLevel(3))
		builder.Add(myrtle.HeroBlock{
			Eyebrow:  "New",
			Title:    "Faster sends with Myrtle",
			Body:     "Compose transactional emails with reusable blocks and themeable output.",
			CTALabel: "Read docs",
			CTAURL:   "https://github.com/gzuidhof/myrtle",
			ImageURL: "/assets/hero.svg",
			ImageAlt: "Hero image",
		})
		builder.AddHeading("Minimal", myrtle.HeadingLevel(3))
		builder.Add(myrtle.HeroBlock{
			Title: "A compact hero",
			Body:  "Useful for announcement emails that do not need an image or CTA.",
		})
	case "footer-links":
		builder.AddHeading("Default", myrtle.HeadingLevel(3))
		builder.AddFooterLinks([]myrtle.FooterLink{
			{Label: "Help", URL: "https://example.com/help"},
			{Label: "Privacy", URL: "https://example.com/privacy"},
			{Label: "Terms", URL: "https://example.com/terms"},
		}, "You’re receiving this because you have an active account.")
		builder.AddHeading("Note only", myrtle.HeadingLevel(3))
		builder.AddFooterLinks(nil, "You’re receiving this because you have an active account.")
		builder.AddHeading("Links only", myrtle.HeadingLevel(3))
		builder.AddFooterLinks([]myrtle.FooterLink{{Label: "Status", URL: "https://example.com/status"}, {Label: "Help", URL: "https://example.com/help"}}, "")
	case "price-summary":
		builder.AddHeading("Default", myrtle.HeadingLevel(3))
		builder.AddPriceSummary("Order summary", []myrtle.PriceLine{
			{Label: "Subtotal", Value: "$89.00"},
			{Label: "Tax", Value: "$7.12"},
			{Label: "Discount", Value: "-$5.00"},
		}, "Total", "$91.12")
		builder.AddHeading("Minimal", myrtle.HeadingLevel(3))
		builder.AddPriceSummary("Order summary", []myrtle.PriceLine{{Label: "Subtotal", Value: "$89.00"}}, "", "")
		builder.AddHeading("With discount", myrtle.HeadingLevel(3))
		builder.AddPriceSummary("Order summary", []myrtle.PriceLine{{Label: "Subtotal", Value: "$100.00"}, {Label: "Discount", Value: "-$12.00"}, {Label: "Tax", Value: "$7.04"}}, "Total", "$95.04")
	case "empty-state":
		builder.AddHeading("Default", myrtle.HeadingLevel(3))
		builder.AddEmptyState("No incidents", "Everything looks healthy right now.", "", "")
		builder.AddHeading("With action", myrtle.HeadingLevel(3))
		builder.AddEmptyState("No incidents", "Everything looks healthy right now.", "View dashboard", "https://example.com/dashboard")
		builder.AddHeading("Long copy", myrtle.HeadingLevel(3))
		builder.AddEmptyState("No open tasks", "You’re all caught up for now. We’ll notify you when something needs your attention.", "Create workflow", "https://example.com/workflows/new")
	case "quote":
		builder.AddQuote("Myrtle made our transactional emails easy to maintain.", "Engineering Team")
	case "callout":
		builder.AddHeading("Info · soft", myrtle.HeadingLevel(3))
		builder.AddCallout(myrtle.CalloutTypeInfo, "FYI", "Routine update with no action required.", myrtle.CalloutStyle(myrtle.CalloutVariantSoft))
		builder.AddHeading("Info · with link", myrtle.HeadingLevel(3))
		builder.AddCallout(myrtle.CalloutTypeInfo, "FYI", "Routine update with no action required.", myrtle.CalloutStyle(myrtle.CalloutVariantSoft), myrtle.CalloutLink("View details", "https://example.com/details"))
		builder.AddHeading("Success · soft", myrtle.HeadingLevel(3))
		builder.AddCallout(myrtle.CalloutTypeSuccess, "Done", "Deployment completed successfully.", myrtle.CalloutStyle(myrtle.CalloutVariantSoft))
		builder.AddHeading("Warning · outline", myrtle.HeadingLevel(3))
		builder.AddCallout(myrtle.CalloutTypeWarning, "Action needed", "Please verify your billing details before March 1.", myrtle.CalloutStyle(myrtle.CalloutVariantOutline))
		builder.AddHeading("Critical · solid", myrtle.HeadingLevel(3))
		builder.AddCallout(myrtle.CalloutTypeCritical, "Action needed", "Please verify your billing details before March 1.", myrtle.CalloutStyle(myrtle.CalloutVariantSolid))
	case "legal":
		builder.AddLegal("Myrtle Inc.", "123 Market St, SF, CA", "https://example.com/preferences", "https://example.com/unsubscribe")
	case "columns":
		builder.AddHeading("Custom widths + gap + middle align", myrtle.HeadingLevel(3))
		builder.AddColumns(
			func(column *myrtle.ColumnBuilder) {
				column.AddHeading("Left column", myrtle.HeadingLevel(3)).
					AddText("Summary and quick context.").
					AddText("Additional details to make this column taller.")
			},
			func(column *myrtle.ColumnBuilder) {
				column.AddHeading("Right column", myrtle.HeadingLevel(3)).
					AddList([]string{"Point one", "Point two"}, false)
			},
			myrtle.ColumnsWidths(60, 40),
			myrtle.ColumnsGap(24),
			myrtle.ColumnsAlign(myrtle.ColumnsVerticalAlignMiddle),
		)
	case "button":
		builder.AddHeading("Primary", myrtle.HeadingLevel(3))
		builder.AddButton("Open docs", "https://github.com/gzuidhof/myrtle", myrtle.ButtonTone(myrtle.ButtonTonePrimary))
		builder.AddHeading("Primary · centered", myrtle.HeadingLevel(3))
		builder.AddButton("Open docs", "https://github.com/gzuidhof/myrtle", myrtle.ButtonTone(myrtle.ButtonTonePrimary), myrtle.ButtonAlign(myrtle.ButtonAlignmentCenter))
		builder.AddHeading("Primary · right", myrtle.HeadingLevel(3))
		builder.AddButton("Open docs", "https://github.com/gzuidhof/myrtle", myrtle.ButtonTone(myrtle.ButtonTonePrimary), myrtle.ButtonAlign(myrtle.ButtonAlignmentRight))
		builder.AddHeading("Secondary", myrtle.HeadingLevel(3))
		builder.AddButton("Open docs", "https://github.com/gzuidhof/myrtle", myrtle.ButtonTone(myrtle.ButtonToneSecondary))
		builder.AddHeading("Danger", myrtle.HeadingLevel(3))
		builder.AddButton("Delete workspace", "https://github.com/gzuidhof/myrtle", myrtle.ButtonTone(myrtle.ButtonToneDanger))
		builder.AddHeading("Outline", myrtle.HeadingLevel(3))
		builder.AddButton("Open docs", "https://github.com/gzuidhof/myrtle", myrtle.ButtonStyle(myrtle.ButtonStyleOutline))
		builder.AddHeading("Ghost", myrtle.HeadingLevel(3))
		builder.AddButton("Open docs", "https://github.com/gzuidhof/myrtle", myrtle.ButtonStyle(myrtle.ButtonStyleGhost))
		builder.AddHeading("Danger · ghost", myrtle.HeadingLevel(3))
		builder.AddButton("Delete workspace", "https://github.com/gzuidhof/myrtle", myrtle.ButtonTone(myrtle.ButtonToneDanger), myrtle.ButtonStyle(myrtle.ButtonStyleGhost))
		builder.AddHeading("Primary · full width", myrtle.HeadingLevel(3))
		builder.AddButton("Open docs", "https://github.com/gzuidhof/myrtle", myrtle.ButtonTone(myrtle.ButtonTonePrimary), myrtle.ButtonFullWidth(true))
		builder.AddHeading("Secondary · full width", myrtle.HeadingLevel(3))
		builder.AddButton("Open docs", "https://github.com/gzuidhof/myrtle", myrtle.ButtonTone(myrtle.ButtonToneSecondary), myrtle.ButtonFullWidth(true))
		builder.AddHeading("Outline · small · no-wrap", myrtle.HeadingLevel(3))
		builder.AddButton("A longer CTA label", "https://github.com/gzuidhof/myrtle", myrtle.ButtonStyle(myrtle.ButtonStyleOutline), myrtle.ButtonSize(myrtle.ButtonSizeSmall), myrtle.ButtonNoWrap(true))
	case "button-group":
		builder.AddHeading("Centered", myrtle.HeadingLevel(3))
		builder.AddButtonGroup([]myrtle.ButtonGroupButton{{Label: "Approve", URL: "https://example.com/approve", Tone: myrtle.ButtonTonePrimary}, {Label: "Review", URL: "https://example.com/review", Tone: myrtle.ButtonToneSecondary}, {Label: "Delete", URL: "https://example.com/delete", Tone: myrtle.ButtonToneDanger}}, myrtle.ButtonGroupAlign(myrtle.ButtonAlignmentCenter))
		builder.AddHeading("Centered · joined", myrtle.HeadingLevel(3))
		builder.AddButtonGroup([]myrtle.ButtonGroupButton{{Label: "Approve", URL: "https://example.com/approve", Tone: myrtle.ButtonTonePrimary}, {Label: "Review", URL: "https://example.com/review", Tone: myrtle.ButtonToneSecondary}, {Label: "Delete", URL: "https://example.com/delete", Tone: myrtle.ButtonToneDanger}}, myrtle.ButtonGroupAlign(myrtle.ButtonAlignmentCenter), myrtle.ButtonGroupJoined(true))
		builder.AddHeading("Right", myrtle.HeadingLevel(3))
		builder.AddButtonGroup([]myrtle.ButtonGroupButton{{Label: "Retry", URL: "https://example.com/retry", Tone: myrtle.ButtonTonePrimary}, {Label: "Details", URL: "https://example.com/details", Tone: myrtle.ButtonToneSecondary}, {Label: "Delete", URL: "https://example.com/delete", Tone: myrtle.ButtonToneDanger}}, myrtle.ButtonGroupAlign(myrtle.ButtonAlignmentRight))
		builder.AddHeading("Full width on mobile", myrtle.HeadingLevel(3))
		builder.AddButtonGroup([]myrtle.ButtonGroupButton{{Label: "Approve", URL: "https://example.com/approve", Tone: myrtle.ButtonTonePrimary}, {Label: "Review", URL: "https://example.com/review", Tone: myrtle.ButtonToneSecondary}}, myrtle.ButtonGroupFullWidthOnMobile(true))
		builder.AddHeading("Stack on mobile · custom gap", myrtle.HeadingLevel(3))
		builder.AddButtonGroup([]myrtle.ButtonGroupButton{{Label: "Approve", URL: "https://example.com/approve", Tone: myrtle.ButtonTonePrimary}, {Label: "Review", URL: "https://example.com/review", Style: myrtle.ButtonStyleOutline}, {Label: "Delete", URL: "https://example.com/delete", Tone: myrtle.ButtonToneDanger, Style: myrtle.ButtonStyleGhost}}, myrtle.ButtonGroupGap(14), myrtle.ButtonGroupStackOnMobile(true))
	case "divider":
		builder.AddHeading("Default", myrtle.HeadingLevel(3))
		builder.AddDivider()
		builder.AddHeading("Dashed", myrtle.HeadingLevel(3))
		builder.AddDividerStyled(myrtle.DividerStyle(myrtle.DividerVariantDashed), myrtle.DividerThickness(2))
		builder.AddHeading("Dotted + inset", myrtle.HeadingLevel(3))
		builder.AddDividerStyled(myrtle.DividerStyle(myrtle.DividerVariantDotted), myrtle.DividerThickness(2), myrtle.DividerInset(32))
	case "image":
		builder.AddHeading("Default", myrtle.HeadingLevel(3))
		builder.AddImage("/assets/myrtle-placeholder.png", "Myrtle placeholder")
		builder.AddHeading("Full width", myrtle.HeadingLevel(3))
		builder.AddImage("/assets/myrtle-placeholder.png", "Myrtle placeholder", myrtle.ImageFullWidth())
		builder.AddHeading("Centered, 320px wide", myrtle.HeadingLevel(3))
		builder.AddImage("/assets/myrtle-placeholder.png", "Myrtle placeholder", myrtle.ImageWidth(320), myrtle.ImageAlign(myrtle.ImageAlignmentCenter))
	case "table":
		builder.AddHeading("Default", myrtle.HeadingLevel(3))
		builder.AddTable("Quarterly numbers", []string{"Metric", "Q1", "Q2"}, [][]string{{"Active Users", "8922", "10452"}, {"Conversion", "4.1%", "4.6%"}, {"Churn", "2.8%", "2.3%"}})
		builder.AddHeading("Compact only", myrtle.HeadingLevel(3))
		builder.AddTable("Quarterly numbers", []string{"Metric", "Q1", "Q2"}, [][]string{{"Active Users", "8922", "10452"}, {"Conversion", "4.1%", "4.6%"}, {"Churn", "2.8%", "2.3%"}}, myrtle.TableCompact(true))
		builder.AddHeading("Compact · zebra · right-aligned numeric", myrtle.HeadingLevel(3))
		builder.AddTable("Quarterly numbers", []string{"Metric", "Q1", "Q2"}, [][]string{{"Active Users", "8922", "10452"}, {"Conversion", "4.1%", "4.6%"}, {"Churn", "2.8%", "2.3%"}}, myrtle.TableZebraRows(true), myrtle.TableCompact(true), myrtle.TableRightAlignNumericColumns(true))
	case "verification_code":
		builder.Add(myrtle.VerificationCodeBlock{Label: "Verification code", Value: "493817"})
	case "message":
		builder.AddMessage(myrtle.MessageBlock{
			SenderName:   "Alex Johnson",
			SenderHandle: "@alex",
			AvatarURL:    "/assets/avatar1.png",
			LogoAlt:      "Alex Johnson avatar",
			LogoHref:     "https://example.com/messages/42",
			Subject:      "New private message",
			Preview:      "Can you review the release notes before 3 PM?",
			SentAt:       "2m ago",
			Platform:     "Myrtle Chat",
			URL:          "https://example.com/messages/42",
			ActionLabel:  "Open thread",
			ActionURL:    "https://example.com/messages/42",
		})
	case "message-digest":
		builder.AddMessageDigest([]myrtle.MessageBlock{
			{SenderName: "Maya", SenderHandle: "@maya", AvatarURL: "/assets/avatar2.png", LogoAlt: "Maya avatar", LogoHref: "https://example.com/messages/43", Subject: "**Launch update**", Preview: "Can you check [the draft](https://example.com/draft)?", SentAt: "5m ago", Platform: "Myrtle Chat", URL: "https://example.com/messages/43"},
			{SenderName: "Nina", SenderHandle: "@nina", AvatarURL: "/assets/avatar3.png", LogoAlt: "Nina avatar", LogoHref: "https://example.com/messages/46", Preview: "Quick ping: can we move this to tomorrow morning?", SentAt: "11m ago", Platform: "Myrtle Chat", URL: "https://example.com/messages/46"},
			{SenderName: "Ben", SenderHandle: "@ben", AvatarURL: "/assets/avatar1.png", LogoAlt: "Ben avatar", LogoHref: "https://example.com/messages/44", Subject: "Design feedback", Preview: "Looks good overall.", SentAt: "20m ago", Platform: "Myrtle Chat", URL: "https://example.com/messages/44"},
			{SenderName: "Ari", SenderHandle: "@ari", Subject: "Follow-up", Preview: "Can we sync tomorrow?", SentAt: "1h ago", Platform: "Myrtle Chat", URL: "https://example.com/messages/45"},
		},
			myrtle.MessageDigestTitle("Inbox"),
			myrtle.MessageDigestSubtitle("Recent direct messages from Myrtle Chat"),
			myrtle.MessageDigestFooter("[Open inbox](https://example.com/messages)"),
		)
	case "tiles":
		builder.AddHeading("Default (3 columns)", myrtle.HeadingLevel(3))
		builder.AddTiles([]myrtle.TileEntry{{Content: "🚀", Title: "Launch", Subtitle: "Ready", URL: "https://example.com/launch", Variant: myrtle.TileVariantHighlight}, {Content: "12", Title: "Queued", Subtitle: "Jobs", Variant: myrtle.TileVariantWarning}, {Content: "✅", Title: "Healthy", Subtitle: "All systems", Variant: myrtle.TileVariantSuccess}})
		builder.AddHeading("Left aligned", myrtle.HeadingLevel(3))
		builder.AddTiles([]myrtle.TileEntry{{Content: "📦", Title: "Shipped", Subtitle: "2h ago"}, {Content: "9", Title: "Pending", Subtitle: "Orders"}, {Content: "⚠️", Title: "Delayed", Subtitle: "1 route", Variant: myrtle.TileVariantWarning}}, myrtle.TilesAlign(myrtle.TileAlignmentLeft))
		builder.AddHeading("Right aligned", myrtle.HeadingLevel(3))
		builder.AddTiles([]myrtle.TileEntry{{Content: "🧾", Title: "Invoices"}, {Content: "3", Title: "Overdue", Variant: myrtle.TileVariantCritical}, {Content: "✅", Title: "Paid", Subtitle: "This week"}}, myrtle.TilesAlign(myrtle.TileAlignmentRight))
		builder.AddHeading("No content examples", myrtle.HeadingLevel(3))
		builder.AddTiles([]myrtle.TileEntry{{Title: "No icon tile", Subtitle: "Title + subtitle only"}, {Title: "Title only"}, {Subtitle: "Subtitle only"}}, myrtle.TilesBorder(true))
		builder.AddHeading("4 columns with border", myrtle.HeadingLevel(3))
		builder.AddTiles([]myrtle.TileEntry{{Content: "1", Title: "One"}, {Content: "2", Title: "Two"}, {Content: "3", Title: "Three", Variant: myrtle.TileVariantCritical}, {Content: "4", Title: "Four"}}, myrtle.TilesColumns(4), myrtle.TilesBorder(true))
		builder.AddHeading("Wraps when items exceed columns", myrtle.HeadingLevel(3))
		builder.AddTiles([]myrtle.TileEntry{{Content: "📥", Title: "Inbox"}, {Content: "📤", Title: "Outbox"}, {Content: "⚙️", Title: "Settings"}, {Content: "🔔", Title: "Alerts"}, {Content: "🧪", Title: "Experiments"}, {Content: "🗂️", Title: "Archive"}}, myrtle.TilesColumns(4), myrtle.TilesBorder(true))
	case "section":
		builder.AddSection(
			[]myrtle.Block{
				myrtle.TextBlock{Text: "Section body content with grouped context."},
				myrtle.ButtonBlock{Label: "Open section", URL: "https://example.com/section", Style: myrtle.ButtonStyleOutline},
			},
			myrtle.SectionTitle("Section block"),
			myrtle.SectionSubtitle("Optional subtitle for context"),
			myrtle.SectionPadding(18),
			myrtle.SectionBorder(true),
		)
	case "grid":
		builder.AddGrid(
			[]myrtle.GridItem{
				{Blocks: []myrtle.Block{myrtle.HeadingBlock{Text: "Item 1", Level: 4}, myrtle.TextBlock{Text: "Grid content one."}}},
				{Blocks: []myrtle.Block{myrtle.HeadingBlock{Text: "Item 2", Level: 4}, myrtle.TextBlock{Text: "Grid content two."}}},
				{Blocks: []myrtle.Block{myrtle.HeadingBlock{Text: "Item 3", Level: 4}, myrtle.TextBlock{Text: "Grid content three."}}},
				{Blocks: []myrtle.Block{myrtle.HeadingBlock{Text: "Item 4", Level: 4}, myrtle.TextBlock{Text: "Grid content four."}}},
				{Blocks: []myrtle.Block{myrtle.HeadingBlock{Text: "Item 5", Level: 4}, myrtle.TextBlock{Text: "Wraps to new row."}}},
			},
			myrtle.GridColumns(3),
			myrtle.GridGap(12),
			myrtle.GridBorder(true),
		)
	case "card-list":
		builder.AddCardList(
			[]myrtle.CardItem{
				{Title: "Deploy complete", Subtitle: "Production", Body: "No customer impact detected.", URL: "https://example.com/deploy/1", CTALabel: "View"},
				{Title: "Billing updated", Subtitle: "Finance", Body: "Invoice #8241 has been paid.", URL: "https://example.com/billing/8241", CTALabel: "Open"},
				{Title: "Weekly report", Body: "Read the latest metrics and highlights.", URL: "https://example.com/reports/weekly"},
			},
			myrtle.CardListColumns(2),
			myrtle.CardListGap(12),
			myrtle.CardListBorder(true),
		)
	case "free-markdown":
		builder.AddFreeMarkdown("### Custom Markdown\n\nYou can use **bold**, lists, and links.\n\n- First\n- Second")
	default:
		return nil, errors.New("unknown block")
	}

	return builder.Build(), nil
}

func buildEmailItems(themeName string, selectedTheme theme.Theme) ([]pageItem, error) {
	items := make([]pageItem, 0, len(exampleEmails))
	for _, emailBuilder := range exampleEmails {
		email, err := emailBuilder.Build(selectedTheme)
		if err != nil {
			return nil, err
		}

		markdown, err := email.Text()
		if err != nil {
			return nil, err
		}

		items = append(items, pageItem{
			Key:      emailBuilder.Name,
			Name:     goEmailName(emailBuilder.Name),
			HTMLURL:  "/emails/" + emailBuilder.Name + "/html?theme=" + themeName,
			Markdown: markdown,
		})
	}

	return items, nil
}

func buildBlockItems(themeName string, selectedTheme theme.Theme) ([]groupedPageItems, error) {
	groups := make([]groupedPageItems, 0, len(blockGroups))
	for _, blockGroup := range blockGroups {
		items := make([]pageItem, 0, len(blockGroup.Items))
		for _, blockName := range blockGroup.Items {
			email, err := buildBlockEmail(blockName, selectedTheme)
			if err != nil {
				return nil, err
			}

			markdown, err := email.Text()
			if err != nil {
				return nil, err
			}

			items = append(items, pageItem{
				Key:      blockName,
				Name:     goBlockName(blockName),
				HTMLURL:  "/blocks/" + blockName + "/html?theme=" + themeName,
				Markdown: markdown,
			})
		}

		groups = append(groups, groupedPageItems{
			Name:  blockGroup.Name,
			Items: items,
		})
	}

	return groups, nil
}

func goEmailName(key string) string {
	return toSpacedTitle(key)
}

func goBlockName(key string) string {
	return toSpacedTitle(key)
}

func toSpacedTitle(value string) string {
	if value == "" {
		return value
	}

	parts := strings.FieldsFunc(value, func(r rune) bool {
		return r == '-' || r == '_'
	})
	for index, part := range parts {
		if part == "" {
			continue
		}

		parts[index] = strings.ToUpper(part[:1]) + part[1:]
	}

	return strings.Join(parts, " ")
}

func selectedThemeFromRequest(queryValue string) (string, theme.Theme) {
	switch queryValue {
	case "flat":
		return "flat", flat.New(flat.WithFallback(defaulttheme.New()))
	default:
		return "default", defaulttheme.New()
	}
}

func themeOptions(selected string) []themeOption {
	return []themeOption{
		{Name: "default", Selected: selected == "default"},
		{Name: "flat", Selected: selected == "flat"},
	}
}
