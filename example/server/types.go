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
	Key     string
	Name    string
	HTMLURL string
	Text    string
}

type groupedPageItems struct {
	Name  string
	Items []pageItem
}

type blockGroupDefinition struct {
	Name  string
	Items []string
}

type emailGroupDefinition struct {
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
	EmailGroups  []groupedPageItems
	BlockGroups  []groupedPageItems
	SMTPEnabled  bool
	SMTPDefault  string
	SendStatus   *sendStatus
}

type sendStatus struct {
	Name    string
	Success bool
	Message string
}

type previewViewData struct {
	Title   string
	Theme   string
	Name    string
	Subject string
	Preview string
	Text    string
}

var exampleEmails = []namedEmailBuilder{
	{Name: "welcome", Build: example.WelcomeEmailWithTheme},
	{Name: "monster", Build: example.MonsterEmailWithTheme},
	{Name: "monster-dark-mode", Build: example.MonsterDarkModeEmailWithTheme},
	{Name: "monster-rtl", Build: example.MonsterRTLEmailWithTheme},
	{Name: "stress", Build: example.StressTestEmailWithTheme},

	{Name: "security", Build: example.SecurityCodeEmailWithTheme},
	{Name: "password-reset", Build: example.PasswordResetEmailWithTheme},
	{Name: "account-deletion-confirmation", Build: example.AccountDeletionConfirmationEmailWithTheme},
	{Name: "incident-notice", Build: example.IncidentNoticeEmailWithTheme},

	{Name: "report", Build: example.WeeklyReportEmailWithTheme},
	{Name: "feature-digest", Build: example.FeatureDigestEmailWithTheme},
	{Name: "weekly-operations-brief", Build: example.WeeklyOperationsBriefEmailWithTheme},
	{Name: "bar-chart", Build: example.HorizontalBarChartEmailWithTheme},
	{Name: "vertical-bar-chart", Build: example.VerticalBarChartEmailWithTheme},
	{Name: "vertical-bar-chart-ticks", Build: example.VerticalBarChartTicksEmailWithTheme},
	{Name: "inset-modes", Build: example.InsetModesEmailWithTheme},

	{Name: "invoice-summary", Build: example.InvoiceSummaryEmailWithTheme},
	{Name: "billing-receipt", Build: example.BillingReceiptEmailWithTheme},
	{Name: "activity-empty-state", Build: example.ActivityEmptyStateEmailWithTheme},
	{Name: "dark-mode-styles", Build: example.DarkModeStylesEmailWithTheme},
	{Name: "custom-feature-flag-rollout", Build: CustomFeatureFlagRolloutEmailWithTheme},

	{Name: "onboarding-checklist", Build: example.OnboardingChecklistEmailWithTheme},
	{Name: "columns-complex", Build: example.ColumnsComplexEmailWithTheme},
}

var exampleEmailGroups = []emailGroupDefinition{
	{
		Name: "Lifecycle & Onboarding",
		Items: []string{
			"welcome",
			"onboarding-checklist",
			"security",
			"password-reset",
			"account-deletion-confirmation",
		},
	},
	{
		Name: "Operations & Incidents",
		Items: []string{
			"incident-notice",
			"activity-empty-state",
			"report",
			"feature-digest",
			"weekly-operations-brief",
			"custom-feature-flag-rollout",
		},
	},
	{
		Name: "Commerce & Billing",
		Items: []string{
			"invoice-summary",
			"billing-receipt",
		},
	},
	{
		Name: "Data & Charts",
		Items: []string{
			"bar-chart",
			"vertical-bar-chart",
			"vertical-bar-chart-ticks",
		},
	},
	{
		Name: "Showcase & Styling",
		Items: []string{
			"columns-complex",
			"dark-mode-styles",
			"monster",
			"monster-dark-mode",
			"monster-rtl",
			"stress",
		},
	},
	{
		Name: "Inset Modes",
		Items: []string{
			"inset-modes",
		},
	},
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
			"panel",
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
			"vertical-bar-chart",
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
