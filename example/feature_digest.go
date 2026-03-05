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
		myrtle.WithStyles(theme.Styles{MaxWidthMain: "560px"}),
	).
		AddHeading("New this month in Myrtle", myrtle.HeadingLevel(1)).
		WithPreheader("Faster rendering, richer blocks, and better previews").
		WithHeader(commonHeaderGroup("Myrtle", selectedTheme)).
		AddHeading("February product updates", myrtle.HeadingLevel(2)).
		AddText("This example uses a narrower container (max-width: 560px).", myrtle.TextTone(myrtle.ToneMuted), myrtle.TextSize(myrtle.TextSizeSmall)).
		AddQuote("Our team shipped a smaller set of changes with clearer rollout notes.", "Platform Team").
		AddKeyValue("Highlights", []myrtle.KeyValuePair{
			{Key: "Rendering", Value: "Improved block consistency across themes"},
			{Key: "Previews", Value: "Cleaner example gallery and navigation"},
		}).
		AddText("Explore the examples and previews:").
		AddButton("Open example server", "http://localhost:8380").
		Build(), nil
}
