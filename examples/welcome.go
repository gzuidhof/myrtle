package example

import (
	"github.com/gzuidhof/myrtle"
	"github.com/gzuidhof/myrtle/theme"
	defaulttheme "github.com/gzuidhof/myrtle/theme/default"
)

func WelcomeEmail() (*myrtle.Email, error) {
	return myrtle.NewBuilder(
		defaulttheme.New(),
		myrtle.WithStyles(theme.Styles{ColorPrimary: "#0ea5e9"}),
	).
		WithHeader(myrtle.HeaderTitle("Welcome to Myrtle")).
		Preheader("Compose beautiful transactional emails").
		Product("Myrtle", "https://github.com/gzuidhof/myrtle").
		Logo("https://example.com/logo.png", "Myrtle Logo").
		AddText("Hi there,").
		AddText("Thanks for joining us. You can now build composable email content in Go.").
		AddText("Start with the quick-start docs:").
		AddButton("Open docs", "https://github.com/gzuidhof/myrtle").
		AddDivider().
		AddFreeMarkdown("Need help? Reach out in **GitHub Discussions**.").
		Build(), nil
}
