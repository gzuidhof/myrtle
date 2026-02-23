package example

import (
	"github.com/gzuidhof/myrtle"
	"github.com/gzuidhof/myrtle/theme"
	defaulttheme "github.com/gzuidhof/myrtle/theme/default"
)

func AccountDeletionConfirmationEmail() (*myrtle.Email, error) {
	return AccountDeletionConfirmationEmailWithTheme(defaulttheme.New())
}

func AccountDeletionConfirmationEmailWithTheme(selectedTheme theme.Theme) (*myrtle.Email, error) {
	if selectedTheme == nil {
		selectedTheme = defaulttheme.New()
	}

	return myrtle.NewBuilder(
		selectedTheme,
		myrtle.WithStyles(theme.Styles{ColorPrimary: "#2563eb"}),
	).
		Preheader("Confirm account deletion within 30 minutes").
		Product("Myrtle Security", "https://example.com/security").
		Logo("/assets/logo.png", "Myrtle Security").
		AddHeading("Confirm account deletion", myrtle.HeadingLevel(2)).
		AddText(
			"We received a request to permanently delete your account and all related data.",
			"If you continue, this action cannot be undone.",
		).
		AddCallout(myrtle.CalloutTypeCritical, "Permanent action", "All projects, API keys, and audit history will be removed.", myrtle.CalloutStyle(myrtle.CalloutVariantSolid)).
		AddKeyValue("Request details", []myrtle.KeyValuePair{{Key: "Requested by", Value: "alex@example.com"}, {Key: "Requested at", Value: "2026-02-23 14:21 UTC"}, {Key: "IP", Value: "203.0.113.42"}}).
		AddText("To continue, confirm your decision below:").
		AddButton("Delete account permanently", "https://example.com/security/delete-account/confirm?token=demo-token", myrtle.ButtonTone(myrtle.ButtonToneDanger), myrtle.ButtonFullWidth(true)).
		AddButton("Keep my account", "https://example.com/security/delete-account/cancel", myrtle.ButtonStyle(myrtle.ButtonStyleGhost)).
		AddLegal("Myrtle Inc.", "123 Market St, San Francisco, CA", "https://example.com/preferences", "https://example.com/unsubscribe").
		Build(), nil
}
