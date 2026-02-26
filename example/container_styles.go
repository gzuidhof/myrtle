package example

import (
	"github.com/gzuidhof/myrtle"
	"github.com/gzuidhof/myrtle/theme"
	defaulttheme "github.com/gzuidhof/myrtle/theme/default"
)

func ContainerStylesEmail() (*myrtle.Email, error) {
	return ContainerStylesEmailWithTheme(defaulttheme.New())
}

func ContainerStylesEmailWithTheme(selectedTheme theme.Theme) (*myrtle.Email, error) {
	if selectedTheme == nil {
		selectedTheme = defaulttheme.New()
	}

	return myrtle.NewBuilder(
		selectedTheme,
		myrtle.WithStyles(theme.Styles{
			ColorPageBackground: "#0f172a",
			ColorMainBackground: "#ffffff",
			BorderMain:          "1px solid #334155",
			RadiusMain:          "20px",
		}),
	).
		WithPreheader("Container style override demo").
		WithHeader(commonHeaderGroup("Myrtle")).
		AddHeading("Container style overrides", myrtle.HeadingLevel(2)).
		AddText("This example customizes only the outer shell/container styling.").
		AddList([]string{"Background comes from ColorPageBackground", "Card surface comes from ColorMainBackground", "Card border comes from BorderMain", "Card corner radius comes from RadiusMain"}, false).
		Build(), nil
}
