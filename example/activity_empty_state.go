package example

import (
	"github.com/gzuidhof/myrtle"
	"github.com/gzuidhof/myrtle/theme"
	defaulttheme "github.com/gzuidhof/myrtle/theme/default"
)

func ActivityEmptyStateEmail() (*myrtle.Email, error) {
	return ActivityEmptyStateEmailWithTheme(defaulttheme.New())
}

func ActivityEmptyStateEmailWithTheme(selectedTheme theme.Theme) (*myrtle.Email, error) {
	if selectedTheme == nil {
		selectedTheme = defaulttheme.New()
	}

	return myrtle.NewBuilder(
		selectedTheme,
		myrtle.WithStyles(theme.Styles{ColorPrimary: "#2563eb"}),
	).
		WithPreheader("Set up alerts to get notified when things change").
		WithHeader(commonHeaderGroup("Myrtle Ops")).
		AddEmptyState("All clear", "No incidents or anomalies were detected in your monitored services.", "Configure alerts", "https://example.com/ops/alerts").
		AddCallout(myrtle.CalloutTypeWarning, "Stay prepared", "Create escalation policies before the next incident.", myrtle.CalloutStyle(myrtle.CalloutVariantOutline)).
		AddButton("Open dashboard", "https://example.com/ops/dashboard", myrtle.ButtonStyle(myrtle.ButtonStyleGhost)).
		AddFooterLinks(
			[]myrtle.FooterLink{{Label: "Status page", URL: "https://example.com/status"}, {Label: "Preferences", URL: "https://example.com/preferences"}},
			"You can disable these reminders in notification settings.",
		).
		Build(), nil
}
