package example

import (
	"github.com/gzuidhof/myrtle"
	"github.com/gzuidhof/myrtle/theme"
	"github.com/gzuidhof/myrtle/theme/flat"
)

func SecurityCodeEmail() (*myrtle.Email, error) {
	return myrtle.NewBuilder(
		flat.New(),
		myrtle.WithStyles(myrtleStyles()),
	).
		AddHeading("Your verification code").
		Preheader("Use this one-time code to sign in").
		Product("Myrtle Security", "https://example.com/security").
		Logo("https://example.com/security-logo.png", "").
		AddText("Use the code below to complete your sign-in. This code expires in 10 minutes.").
		Add(myrtle.CodeBlock{Label: "Verification code", Code: "493817"}).
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
