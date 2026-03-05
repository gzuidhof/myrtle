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
		WithPreheader("Use the secure link below to choose a new password").
		AddHeading("Password reset request", myrtle.HeadingLevel(2)).
		AddText("We received a request to reset the password for your account.").
		AddText("If this was you, continue below. This link expires in 30 minutes.").
		AddText("Confirm your identity and set a new password:").
		AddButton("Reset password", "https://example.com/security/reset?token=demo-token").
		AddCallout(myrtle.ToneWarning, "Didn't request this?", "You can ignore this email. Your password will stay unchanged.").
		AddKeyValue("Request details", []myrtle.KeyValuePair{{Key: "Time", Value: "2026-02-21 10:14 UTC"}, {Key: "IP", Value: "203.0.113.5"}}).
		AddLegal("Myrtle Inc.", "Dam Square 1, 1012 JS Amsterdam, Netherlands", "https://example.com/preferences", "https://example.com/unsubscribe").
		Build(), nil
}
