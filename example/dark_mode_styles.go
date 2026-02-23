package example

import (
	"github.com/gzuidhof/myrtle"
	"github.com/gzuidhof/myrtle/theme"
	defaulttheme "github.com/gzuidhof/myrtle/theme/default"
)

func DarkModeStylesEmail() (*myrtle.Email, error) {
	return DarkModeStylesEmailWithTheme(defaulttheme.New())
}

func DarkModeStylesEmailWithTheme(selectedTheme theme.Theme) (*myrtle.Email, error) {
	if selectedTheme == nil {
		selectedTheme = defaulttheme.New()
	}

	return myrtle.NewBuilder(
		selectedTheme,
		myrtle.WithStyles(theme.Styles{
			ColorPrimary:        "#60a5fa",
			ColorSecondary:      "#34d399",
			ColorText:           "#e5e7eb",
			ColorTextMuted:      "#94a3b8",
			ColorBorder:         "#334155",
			ColorCodeBackground: "#0f172a",
			ColorPageBackground: "#020617",
			ColorMainBackground: "#0b1220",
			BorderMain:          "1px solid #334155",
			RadiusMain:          "16px",
		}),
	).
		Preheader("Dark mode style override demo").
		Product("Myrtle", "https://github.com/gzuidhof/myrtle").
		Logo("/assets/logo.png", "Myrtle").
		AddHeading("Dark mode via style overrides", myrtle.HeadingLevel(2)).
		AddText("This example uses style overrides only, without a dedicated dark theme implementation.").
		AddCallout(myrtle.CalloutTypeSuccess, "Looks good", "For many use cases, color and shell tokens are enough to produce a dark mode look.").
		AddVerificationCode("Verification code", "582941").
		AddMessage(myrtle.MessageBlock{
			Subject: "Deployment complete",
			Preview: "Production rollout succeeded with no alerts.",
			SentAt:  "5 minutes ago",
			URL:     "https://example.com/deployments/123",
		}).
		AddButton("Open dashboard", "https://example.com/dashboard").
		Build(), nil
}
