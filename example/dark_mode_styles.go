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
		myrtle.WithStyles(darkModeStylesOverrides()),
	).
		WithPreheader("Dark mode style override demo").
		WithHeader(commonHeaderGroup("Myrtle")).
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

func darkModeStylesOverrides() theme.Styles {
	styles := theme.DefaultDarkModeStyles()
	styles.ColorPrimary = "#297ce1"
	styles.ColorSecondary = "#34d399"

	return styles
}
