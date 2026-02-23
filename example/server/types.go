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
	{Name: "product-launch", Build: example.ProductLaunchEmailWithTheme},
	{Name: "common-blocks", Build: example.CommonBlocksEmailWithTheme},
	{Name: "monster", Build: example.MonsterEmailWithTheme},
	{Name: "stress", Build: example.StressTestEmailWithTheme},

	{Name: "security", Build: example.SecurityCodeEmailWithTheme},
	{Name: "password-reset", Build: example.PasswordResetEmailWithTheme},
	{Name: "account-deletion-confirmation", Build: example.AccountDeletionConfirmationEmailWithTheme},
	{Name: "incident-notice", Build: example.IncidentNoticeEmailWithTheme},

	{Name: "report", Build: example.WeeklyReportEmailWithTheme},
	{Name: "feature-digest", Build: example.FeatureDigestEmailWithTheme},
	{Name: "high-impact", Build: example.HighImpactEmailWithTheme},
	{Name: "bar-chart", Build: example.BarChartEmailWithTheme},

	{Name: "invoice-summary", Build: example.InvoiceSummaryEmailWithTheme},
	{Name: "billing-receipt", Build: example.BillingReceiptEmailWithTheme},
	{Name: "activity-empty-state", Build: example.ActivityEmptyStateEmailWithTheme},
	{Name: "container-styles", Build: example.ContainerStylesEmailWithTheme},
	{Name: "dark-mode-styles", Build: example.DarkModeStylesEmailWithTheme},

	{Name: "onboarding-checklist", Build: example.OnboardingChecklistEmailWithTheme},
	{Name: "columns-complex", Build: example.ColumnsComplexEmailWithTheme},
}

var blockGroups = []blockGroupDefinition{
	{
		Name: "Text & Content",
		Items: []string{
			"text",
			"heading",
			"list",
			"quote",
			"free-markdown",
			"spacer",
		},
	},
	{
		Name: "Messaging & Alerts",
		Items: []string{
			"message",
			"message-digest",
			"verification_code",
			"callout",
			"badge",
		},
	},
	{
		Name: "Actions & Navigation",
		Items: []string{
			"button",
			"button-group",
			"hero",
			"image",
			"footer-links",
			"attachment",
		},
	},
	{
		Name: "Layout & Structure",
		Items: []string{
			"section",
			"columns",
			"grid",
			"card-list",
			"divider",
		},
	},
	{
		Name: "Data & Metrics",
		Items: []string{
			"table",
			"key-value",
			"stats-row",
			"timeline",
			"price-summary",
			"summary-card",
			"tiles",
			"bar-chart",
			"progress",
			"sparkline",
			"stacked-bar",
			"distribution",
		},
	},
	{
		Name: "Account & Legal",
		Items: []string{
			"empty-state",
			"legal",
		},
	},
}
