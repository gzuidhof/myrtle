package example

import (
	"github.com/gzuidhof/myrtle"
	"github.com/gzuidhof/myrtle/theme"
	defaulttheme "github.com/gzuidhof/myrtle/theme/default"
)

func FeatureDigestEmail() (*myrtle.Email, error) {
	return FeatureDigestEmailWithTheme(defaulttheme.New())
}

func FeatureDigestEmailWithTheme(selectedTheme theme.Theme) (*myrtle.Email, error) {
	if selectedTheme == nil {
		selectedTheme = defaulttheme.New()
	}

	return myrtle.NewBuilder(
		selectedTheme,
	).
		AddHeading("New this month in Myrtle").
		Preheader("Faster rendering, richer blocks, and better previews").
		Product("Myrtle", "https://github.com/gzuidhof/myrtle").
		Logo("/assets/logo.png", "Myrtle").
		AddHeading("February product updates", myrtle.HeadingLevel(2)).
		AddQuote("Our team shipped a smaller set of changes with clearer rollout notes.", "Platform Team").
		AddKeyValue("Highlights", []myrtle.KeyValuePair{
			{Key: "Rendering", Value: "Improved block consistency across themes"},
			{Key: "Previews", Value: "Cleaner example gallery and navigation"},
		}).
		AddText("Explore the examples and previews:").
		AddButton("Open example server", "http://localhost:8380").
		Build(), nil
}
