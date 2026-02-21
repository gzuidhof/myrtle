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
		myrtle.WithStyles(theme.Styles{PrimaryColor: "#0ea5e9"}),
	).
		Preheader("Compose beautiful transactional emails").
		Product("Myrtle", "https://github.com/gzuidhof/myrtle").
		Logo("/assets/logo.png", "Myrtle Logo").
		AddHeading("Welcome aboard", myrtle.HeadingLevel(1)).
		AddText("Hi there,").
		AddText("Thanks for joining us. You can now build composable email content in Go.").
		AddList([]string{"Choose a theme", "Compose with blocks", "Render HTML and Markdown"}, false).
		AddCallout(myrtle.CalloutTypeInfo, "Tip", "You can preview all built-in blocks in the example server.").
		AddAction("Start with the quick-start docs:", "Open docs", "https://github.com/gzuidhof/myrtle").
		AddDivider().
		AddLegal("Myrtle Inc.", "123 Market St, San Francisco, CA", "https://example.com/preferences", "https://example.com/unsubscribe").
		AddFreeMarkdown("Need help? Reach out in **GitHub Discussions**.").
		Build(), nil
}
