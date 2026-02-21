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
		builder.AddText("Content above spacer").AddSpacer(24).AddText("Content below spacer")
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
		builder.AddColumns(
			func(column *myrtle.ColumnBuilder) {
				column.AddHeading("Left column", myrtle.HeadingLevel(3)).
					AddText("Summary and quick context.")
			},
			func(column *myrtle.ColumnBuilder) {
				column.AddHeading("Right column", myrtle.HeadingLevel(3)).
					AddList([]string{"Point one", "Point two"}, false)
			},
			myrtle.ColumnsWidths(60, 40),
		)
	case "button":
		builder.AddHeading("Primary", myrtle.HeadingLevel(3))
		builder.AddButton("Open docs", "https://github.com/gzuidhof/myrtle", myrtle.ButtonStyle(myrtle.ButtonVariantPrimary))
		builder.AddHeading("Primary · centered", myrtle.HeadingLevel(3))
		builder.AddButton("Open docs", "https://github.com/gzuidhof/myrtle", myrtle.ButtonStyle(myrtle.ButtonVariantPrimary), myrtle.ButtonAlign(myrtle.ButtonAlignmentCenter))
		builder.AddHeading("Primary · right", myrtle.HeadingLevel(3))
		builder.AddButton("Open docs", "https://github.com/gzuidhof/myrtle", myrtle.ButtonStyle(myrtle.ButtonVariantPrimary), myrtle.ButtonAlign(myrtle.ButtonAlignmentRight))
		builder.AddHeading("Secondary", myrtle.HeadingLevel(3))
		builder.AddButton("Open docs", "https://github.com/gzuidhof/myrtle", myrtle.ButtonStyle(myrtle.ButtonVariantSecondary))
		builder.AddHeading("Ghost", myrtle.HeadingLevel(3))
		builder.AddButton("Open docs", "https://github.com/gzuidhof/myrtle", myrtle.ButtonStyle(myrtle.ButtonVariantGhost))
		builder.AddHeading("Primary · full width", myrtle.HeadingLevel(3))
		builder.AddButton("Open docs", "https://github.com/gzuidhof/myrtle", myrtle.ButtonStyle(myrtle.ButtonVariantPrimary), myrtle.ButtonFullWidth(true))
		builder.AddHeading("Secondary · full width", myrtle.HeadingLevel(3))
		builder.AddButton("Open docs", "https://github.com/gzuidhof/myrtle", myrtle.ButtonStyle(myrtle.ButtonVariantSecondary), myrtle.ButtonFullWidth(true))
	case "button-group":
		builder.AddHeading("Centered", myrtle.HeadingLevel(3))
		builder.AddButtonGroup([]myrtle.ButtonGroupButton{{Label: "Approve", URL: "https://example.com/approve", Variant: myrtle.ButtonVariantPrimary}, {Label: "Review", URL: "https://example.com/review", Variant: myrtle.ButtonVariantSecondary}, {Label: "Later", URL: "https://example.com/later", Variant: myrtle.ButtonVariantGhost}}, myrtle.ButtonGroupAlign(myrtle.ButtonAlignmentCenter))
		builder.AddHeading("Centered · joined", myrtle.HeadingLevel(3))
		builder.AddButtonGroup([]myrtle.ButtonGroupButton{{Label: "Approve", URL: "https://example.com/approve", Variant: myrtle.ButtonVariantPrimary}, {Label: "Review", URL: "https://example.com/review", Variant: myrtle.ButtonVariantSecondary}, {Label: "Later", URL: "https://example.com/later", Variant: myrtle.ButtonVariantGhost}}, myrtle.ButtonGroupAlign(myrtle.ButtonAlignmentCenter), myrtle.ButtonGroupJoined(true))
		builder.AddHeading("Right", myrtle.HeadingLevel(3))
		builder.AddButtonGroup([]myrtle.ButtonGroupButton{{Label: "Retry", URL: "https://example.com/retry", Variant: myrtle.ButtonVariantPrimary}, {Label: "Details", URL: "https://example.com/details", Variant: myrtle.ButtonVariantSecondary}}, myrtle.ButtonGroupAlign(myrtle.ButtonAlignmentRight))
	case "divider":
		builder.AddDivider()
	case "image":
		builder.AddImage("https://via.placeholder.com/560x180.png?text=Myrtle", "Myrtle placeholder")
	case "table":
		builder.AddHeading("Default", myrtle.HeadingLevel(3))
		builder.AddTable("Quarterly numbers", []string{"Metric", "Q1", "Q2"}, [][]string{{"Active Users", "8922", "10452"}, {"Conversion", "4.1%", "4.6%"}, {"Churn", "2.8%", "2.3%"}})
		builder.AddHeading("Compact only", myrtle.HeadingLevel(3))
		builder.AddTable("Quarterly numbers", []string{"Metric", "Q1", "Q2"}, [][]string{{"Active Users", "8922", "10452"}, {"Conversion", "4.1%", "4.6%"}, {"Churn", "2.8%", "2.3%"}}, myrtle.TableCompact(true))
		builder.AddHeading("Compact · zebra · right-aligned numeric", myrtle.HeadingLevel(3))
		builder.AddTable("Quarterly numbers", []string{"Metric", "Q1", "Q2"}, [][]string{{"Active Users", "8922", "10452"}, {"Conversion", "4.1%", "4.6%"}, {"Churn", "2.8%", "2.3%"}}, myrtle.TableZebraRows(true), myrtle.TableCompact(true), myrtle.TableRightAlignNumericColumns(true))
	case "action":
		builder.AddAction("Complete your setup to get started.", "Finish setup", "https://example.com/setup")
	case "code":
		builder.Add(myrtle.CodeBlock{Label: "Verification code", Code: "493817"})
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
	return toSpacedTitle(key) + " Email"
}

func goBlockName(key string) string {
	return toPascalCase(key)
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

func toPascalCase(value string) string {
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

	return strings.Join(parts, "")
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
