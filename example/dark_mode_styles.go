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
		WithPreheader("Confirm it was you").
		WithHeader(commonHeaderGroupWithLogo("Myrtle", "Myrtle", commonHeaderLogoLightSrc)).
		AddHeading("New sign-in detected", myrtle.HeadingLevel(2)).
		AddText("We noticed a sign-in attempt from a new browser on macOS.").
		AddCallout(myrtle.ToneInfo, "If this was you", "Use the code below to finish signing in.").
		AddVerificationCode("Security code", "582941").
		AddMessage(myrtle.MessageBlock{
			Subject: "Sign-in request",
			Preview: "Chrome on macOS · Amsterdam, NL",
			SentAt:  "5 minutes ago",
			URL:     "https://example.com/security/activity",
		}).
		AddButton("Review activity", "https://example.com/security/activity").
		Build(), nil
}

func darkModeStylesOverrides() theme.Styles {
	styles := theme.DefaultDarkModeStyles()
	styles.ColorPrimary = "#297ce1"
	styles.ColorSecondary = "#34d399"

	return styles
}
