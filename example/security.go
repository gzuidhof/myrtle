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
		myrtle.WithStyles(myrtleStyles()),
	).
		AddHeading("Your verification code").
		Preheader("Use this one-time code to sign in").
		Product("Myrtle Security", "https://example.com/security").
		Logo("/assets/logo.png", "").
		AddCallout(myrtle.CalloutTypeWarning, "Security notice", "Never share this code with anyone.").
		AddText("Use the code below to complete your sign-in. This code expires in 10 minutes.").
		AddSpacer(10).
		Add(myrtle.CodeBlock{Label: "Verification code", Code: "493817"}).
		AddKeyValue("Request details", []myrtle.KeyValuePair{{Key: "IP", Value: "203.0.113.5"}, {Key: "Location", Value: "Amsterdam, NL"}}).
		AddAction("If you did not request this code, secure your account immediately.", "Review account", "https://example.com/account/security").
		Build(), nil
}

func myrtleStyles() theme.Styles {
	return theme.Styles{
		PrimaryColor:        "#7c3aed",
		TextColor:           "#111827",
		MutedTextColor:      "#4b5563",
		BorderColor:         "#ddd6fe",
		CodeBackgroundColor: "#f5f3ff",
	}
}
