package example

import (
	"github.com/gzuidhof/myrtle"
	"github.com/gzuidhof/myrtle/theme"
	defaulttheme "github.com/gzuidhof/myrtle/theme/default"
)

func PasswordResetEmail() (*myrtle.Email, error) {
	return PasswordResetEmailWithTheme(defaulttheme.New())
}

func PasswordResetEmailWithTheme(selectedTheme theme.Theme) (*myrtle.Email, error) {
	if selectedTheme == nil {
		selectedTheme = defaulttheme.New()
	}

	return myrtle.NewBuilder(
		selectedTheme,
		myrtle.WithStyles(theme.Styles{ColorPrimary: "#2563eb"}),
	).
		Preheader("Use the secure link below to choose a new password").
		Product("Myrtle Security", "https://example.com/security").
		Logo("/assets/logo.png", "Myrtle Security").
		AddHeading("Password reset request", myrtle.HeadingLevel(2)).
		AddText(
			"We received a request to reset the password for your account.",
			"If this was you, continue below. This link expires in 30 minutes.",
		).
		AddText("Confirm your identity and set a new password:").
		AddButton("Reset password", "https://example.com/security/reset?token=demo-token").
		AddCallout(myrtle.CalloutTypeWarning, "Didn't request this?", "You can ignore this email. Your password will stay unchanged.").
		AddKeyValue("Request details", []myrtle.KeyValuePair{{Key: "Time", Value: "2026-02-21 10:14 UTC"}, {Key: "IP", Value: "203.0.113.5"}}).
		AddLegal("Myrtle Inc.", "123 Market St, San Francisco, CA", "https://example.com/preferences", "https://example.com/unsubscribe").
		Build(), nil
}
