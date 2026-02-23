package example

import (
	"github.com/gzuidhof/myrtle"
	"github.com/gzuidhof/myrtle/theme"
	defaulttheme "github.com/gzuidhof/myrtle/theme/default"
)

func BarChartEmail() (*myrtle.Email, error) {
	return BarChartEmailWithTheme(defaulttheme.New())
}

func BarChartEmailWithTheme(selectedTheme theme.Theme) (*myrtle.Email, error) {
	if selectedTheme == nil {
		selectedTheme = defaulttheme.New()
	}

	return myrtle.NewBuilder(
		selectedTheme,
		myrtle.WithStyles(theme.Styles{ColorPrimary: "#0ea5e9"}),
	).
		AddHeading("Delivery analytics snapshot").
		Preheader("Simple email-safe bar chart for regional delivery share").
		Product("Myrtle", "https://github.com/gzuidhof/myrtle").
		Logo("/assets/logo.png", "Myrtle").
		AddHeading("Regional message distribution", myrtle.HeadingLevel(1)).
		AddText("This example shows a compact, email-client-safe horizontal bar chart rendered with tables.").
		AddBarChart("Share of delivered messages", []myrtle.BarChartItem{
			{Label: "US", Value: "52%", Percent: 52},
			{Label: "EMEA", Value: "31%", Percent: 31},
			{Label: "APAC", Value: "17%", Percent: 17},
		}).
		AddCallout(myrtle.CalloutTypeInfo, "Note", "Use percentages for consistent scale across all rows.").
		AddLegal("Myrtle Inc.", "123 Market St, San Francisco, CA", "https://example.com/preferences", "https://example.com/unsubscribe").
		Build(), nil
}
