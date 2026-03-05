package example

import (
	"github.com/gzuidhof/myrtle"
	"github.com/gzuidhof/myrtle/theme"
	defaulttheme "github.com/gzuidhof/myrtle/theme/default"
)

func WelcomeEmail() (*myrtle.Email, error) {
	return WelcomeEmailWithTheme(defaulttheme.New())
}

func WelcomeEmailWithTheme(selectedTheme theme.Theme) (*myrtle.Email, error) {
	if selectedTheme == nil {
		selectedTheme = defaulttheme.New()
	}

	return myrtle.NewBuilder(
		selectedTheme,
		myrtle.WithStyles(theme.Styles{ColorPrimary: "#6d28d9", MainContentBodyTopSpacing: "0"}),
	).
		WithPreheader("Compose beautiful transactional emails").
		AddImage(
			"/assets/dark-gradient-hero.png",
			"Dark gradient hero",
			myrtle.ImageFullWidth(),
			myrtle.ImageInsetMode(myrtle.InsetModeNone),
			myrtle.ImageTopSpacing(0),
			myrtle.ImageTopCorners(),
		).
		AddHeading("Welcome aboard", myrtle.HeadingLevel(1)).
		AddText("Hi there,").
		AddText("Thanks for joining us. You can now build composable email content in Go.").
		AddList([]string{"Choose a theme", "Compose with blocks", "Render HTML and Markdown"}, false).
		AddCallout(myrtle.ToneInfo, "Tip", "You can preview all built-in blocks in the example server.").
		AddText("Start with the quick-start docs:").
		AddButton("Open docs", "https://github.com/gzuidhof/myrtle").
		AddDivider().
		AddFreeMarkdown("Need help? Reach out in **GitHub Discussions**.").
		WithFooter(commonFooterGroup(), myrtle.FooterPlacement(myrtle.FooterPlacementOutside)).
		Build(), nil
}
