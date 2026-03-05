package example

import (
	"github.com/gzuidhof/myrtle"
	"github.com/gzuidhof/myrtle/theme"
	defaulttheme "github.com/gzuidhof/myrtle/theme/default"
)

func VerticalBarChartTicksEmail() (*myrtle.Email, error) {
	return VerticalBarChartTicksEmailWithTheme(defaulttheme.New())
}

func VerticalBarChartTicksEmailWithTheme(selectedTheme theme.Theme) (*myrtle.Email, error) {
	if selectedTheme == nil {
		selectedTheme = defaulttheme.New()
	}

	axisLabels := []string{"Mon", "Tue", "Wed", "Thu"}
	series := []myrtle.VerticalBarChartSeries{
		{Key: "new", Label: "New", Color: "#2563eb", Values: []float64{24, 18, 28, 20}},
		{Key: "expansion", Label: "Expansion", Color: "#60a5fa", Values: []float64{8, 11, 6, 9}},
		{Key: "churn", Label: "Churn", Color: "#ef4444", Values: []float64{-6, -5, -8, -7}},
	}

	return myrtle.NewBuilder(
		selectedTheme,
		myrtle.WithStyles(theme.Styles{ColorPrimary: "#2563eb"}),
	).
		WithPreheader("ShowYTicks on and off for a vertical bar chart").
		WithHeader(commonHeaderGroup("Myrtle", selectedTheme)).
		AddHeading("Vertical bar chart: ShowYTicks", myrtle.HeadingLevel(2)).
		AddText("Compare the same data with Y-axis ticks enabled vs. disabled.").
		AddHeading("ShowYTicks = true", myrtle.HeadingLevel(3)).
		AddVerticalBarChart(
			axisLabels,
			series,
			myrtle.VerticalBarChartTitle("Weekly movement"),
			myrtle.VerticalBarChartHeight(180),
			myrtle.VerticalBarChartLegendPlacement(myrtle.VerticalBarChartLegendBottom),
			myrtle.VerticalBarChartAxisShowYTicks(true),
			myrtle.VerticalBarChartAxisShowBaseline(true),
		).
		AddHeading("ShowYTicks = false", myrtle.HeadingLevel(3)).
		AddVerticalBarChart(
			axisLabels,
			series,
			myrtle.VerticalBarChartTitle("Weekly movement"),
			myrtle.VerticalBarChartHeight(180),
			myrtle.VerticalBarChartLegendPlacement(myrtle.VerticalBarChartLegendBottom),
			myrtle.VerticalBarChartAxisShowYTicks(false),
			myrtle.VerticalBarChartAxisShowBaseline(true),
		).
		WithFooter(commonFooterGroup()).
		Build(), nil
}
