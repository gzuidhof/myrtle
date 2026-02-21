package server

import (
	"github.com/gzuidhof/myrtle"
	"github.com/gzuidhof/myrtle/example"
	"github.com/gzuidhof/myrtle/theme"
)

type namedEmailBuilder struct {
	Name  string
	Build func(selectedTheme theme.Theme) (*myrtle.Email, error)
}

type pageItem struct {
	Key      string
	Name     string
	HTMLURL  string
	Markdown string
}

type groupedPageItems struct {
	Name  string
	Items []pageItem
}

type blockGroupDefinition struct {
	Name  string
	Items []string
}

type themeOption struct {
	Name     string
	Selected bool
}

type indexViewData struct {
	Title        string
	ThemeOptions []themeOption
	Theme        string
	EmailItems   []pageItem
	BlockGroups  []groupedPageItems
}

type previewViewData struct {
	Title    string
	Theme    string
	Name     string
	Subject  string
	Preview  string
	Markdown string
}

var exampleEmails = []namedEmailBuilder{
	{Name: "welcome", Build: example.WelcomeEmailWithTheme},
	{Name: "security", Build: example.SecurityCodeEmailWithTheme},
	{Name: "password-reset", Build: example.PasswordResetEmailWithTheme},
	{Name: "report", Build: example.WeeklyReportEmailWithTheme},
	{Name: "common-blocks", Build: example.CommonBlocksEmailWithTheme},
	{Name: "onboarding-checklist", Build: example.OnboardingChecklistEmailWithTheme},
	{Name: "billing-receipt", Build: example.BillingReceiptEmailWithTheme},
	{Name: "incident-notice", Build: example.IncidentNoticeEmailWithTheme},
	{Name: "feature-digest", Build: example.FeatureDigestEmailWithTheme},
	{Name: "high-impact", Build: example.HighImpactEmailWithTheme},
	{Name: "columns-complex", Build: example.ColumnsComplexEmailWithTheme},
	{Name: "bar-chart", Build: example.BarChartEmailWithTheme},
	{Name: "product-launch", Build: example.ProductLaunchEmailWithTheme},
	{Name: "invoice-summary", Build: example.InvoiceSummaryEmailWithTheme},
	{Name: "activity-empty-state", Build: example.ActivityEmptyStateEmailWithTheme},
	{Name: "monster", Build: example.MonsterEmailWithTheme},
}

var blockGroups = []blockGroupDefinition{
	{
		Name: "Content",
		Items: []string{
			"code",
			"free-markdown",
			"heading",
			"list",
			"quote",
			"text",
		},
	},
	{
		Name: "Actions",
		Items: []string{
			"action",
			"badge",
			"button",
			"button-group",
			"callout",
		},
	},
	{
		Name: "Layout & Structure",
		Items: []string{
			"columns",
			"divider",
			"hero",
			"image",
			"spacer",
			"table",
		},
	},
	{
		Name: "Data & Metrics",
		Items: []string{
			"bar-chart",
			"distribution",
			"key-value",
			"price-summary",
			"progress",
			"sparkline",
			"stacked-bar",
			"stats-row",
			"timeline",
		},
	},
	{
		Name: "Account & System",
		Items: []string{
			"attachment",
			"empty-state",
			"footer-links",
			"legal",
			"summary-card",
		},
	},
}
