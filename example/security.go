package example

import (
	"github.com/gzuidhof/myrtle"
	"github.com/gzuidhof/myrtle/theme"
	"github.com/gzuidhof/myrtle/theme/flat"
)

func SecurityCodeEmail() (*myrtle.Email, error) {
	return SecurityCodeEmailWithTheme(flat.New())
}

func SecurityCodeEmailWithTheme(selectedTheme theme.Theme) (*myrtle.Email, error) {
	if selectedTheme == nil {
		selectedTheme = flat.New()
	}

	return myrtle.NewBuilder(
		selectedTheme,
	).
		WithPreheader("Use this one-time code to sign in").
		AddHeading("Your verification code").
		WithHeader(commonHeaderGroup("Myrtle Security", selectedTheme)).
		AddText("Use the code below to complete your sign-in. This code expires in 10 minutes.").
		Add(myrtle.VerificationCodeBlock{Label: "Verification code", Value: "493817"}).
		AddKeyValue("Request details", []myrtle.KeyValuePair{{Key: "IP", Value: "203.0.113.5"}, {Key: "Location", Value: "Amsterdam, NL"}}).
		AddText("If you did not request this code, secure your account immediately.").
		AddButton("Review account", "https://example.com/account/security").
		Build(), nil
}
