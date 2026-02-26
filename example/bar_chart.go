package example

import (
	"github.com/gzuidhof/myrtle"
	"github.com/gzuidhof/myrtle/theme"
	defaulttheme "github.com/gzuidhof/myrtle/theme/default"
)

func HorizontalBarChartEmail() (*myrtle.Email, error) {
	return HorizontalBarChartEmailWithTheme(defaulttheme.New())
}

func HorizontalBarChartEmailWithTheme(selectedTheme theme.Theme) (*myrtle.Email, error) {
	if selectedTheme == nil {
		selectedTheme = defaulttheme.New()
	}

	return myrtle.NewBuilder(
		selectedTheme,
		myrtle.WithStyles(theme.Styles{ColorPrimary: "#0ea5e9"}),
	).
		AddHeading("Delivery analytics snapshot").
		WithPreheader("Simple email-safe bar chart for regional delivery share").
		WithHeader(commonHeaderGroup("Myrtle")).
		AddHeading("Regional message distribution", myrtle.HeadingLevel(1)).
		AddText("This example shows a compact, email-client-safe horizontal bar chart rendered with tables.").
		AddHorizontalBarChart("Share of delivered messages", []myrtle.HorizontalBarChartItem{
			{Label: "US", Value: "52%", Percent: 52},
			{Label: "EMEA", Value: "31%", Percent: 31},
			{Label: "APAC", Value: "17%", Percent: 17},
		}).
		AddCallout(myrtle.CalloutTypeInfo, "Note", "Use percentages for consistent scale across all rows.").
		AddLegal("Myrtle Inc.", "123 Market St, San Francisco, CA", "https://example.com/preferences", "https://example.com/unsubscribe").
		Build(), nil
}
