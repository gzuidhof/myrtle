package example

import (
	"fmt"

	"github.com/gzuidhof/myrtle"
	"github.com/gzuidhof/myrtle/theme"
	defaulttheme "github.com/gzuidhof/myrtle/theme/default"
)

func VerticalBarChartEmail() (*myrtle.Email, error) {
	return VerticalBarChartEmailWithTheme(defaulttheme.New())
}

func VerticalBarChartEmailWithTheme(selectedTheme theme.Theme) (*myrtle.Email, error) {
	if selectedTheme == nil {
		selectedTheme = defaulttheme.New()
	}

	series := []myrtle.VerticalBarChartSeries{
		{Key: "new", Label: "New", Color: "#1d4ed8", Values: []float64{42, 38, 49, 45}},
		{Key: "expansion", Label: "Expansion", Color: "#3b82f6", Values: []float64{18, 22, 26, 19}},
		{Key: "churn", Label: "Churn", Color: "#ef4444", Values: []float64{-11, -9, -14, -12}},
	}
	normalizedSeries := []myrtle.VerticalBarChartSeries{
		{Key: "new", Label: "New", Color: "#1d4ed8", Values: []float64{42, 38, 49, 45}},
		{Key: "expansion", Label: "Expansion", Color: "#3b82f6", Values: []float64{18, 22, 26, 19}},
		{Key: "renewal", Label: "Renewal", Color: "#60a5fa", Values: []float64{15, 17, 16, 18}},
	}
	normalizedPercentSeries := []myrtle.VerticalBarChartSeries{
		{Key: "new", Label: "New", Color: "#1d4ed8", Values: []float64{52, 48, 55, 50}},
		{Key: "expansion", Label: "Expansion", Color: "#3b82f6", Values: []float64{28, 31, 24, 30}},
		{Key: "renewal", Label: "Renewal", Color: "#60a5fa", Values: []float64{20, 21, 21, 20}},
	}
	tooSmallSingleSeries := []myrtle.VerticalBarChartSeries{
		{Key: "new", Label: "New", Values: []float64{2, 38, 3, 34}},
	}
	tooSmallMultiSeries := []myrtle.VerticalBarChartSeries{
		{Key: "base", Label: "Base", Color: "#1d4ed8", Values: []float64{70, 0, 65, 0}},
		{Key: "small", Label: "Small", Color: "#60a5fa", Values: []float64{3, 0, 4, 0}},
		{Key: "churn", Label: "Churn", Color: "#ef4444", Values: []float64{0, -3, 0, -4}},
		{Key: "refund", Label: "Refund", Color: "#f87171", Values: []float64{0, -30, 0, -22}},
	}
	tooSmallNoSpaceSeries := []myrtle.VerticalBarChartSeries{
		{Key: "base", Label: "Base", Color: "#60a5fa", Values: []float64{95, 92, 94, 93}},
		{Key: "small", Label: "Small", Color: "#3b82f6", Values: []float64{5, 8, 6, 7}},
	}
	axisLabels := []string{"Jan", "Feb", "Mar", "Apr"}

	monthNames := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	axisLabels24Months := make([]string, 0, 24)
	series24Months := []myrtle.VerticalBarChartSeries{
		{Key: "new", Label: "New", Color: "#1ebd59", Values: make([]float64, 0, 24)},
		{Key: "expansion", Label: "Expansion", Color: "#36de73", Values: make([]float64, 0, 24)},
		{Key: "churn", Label: "Churn", Color: "#ec581d", Values: make([]float64, 0, 24)},
	}
	for i := 0; i < 24; i++ {
		year := 2024
		if i >= 12 {
			year = 2025
		}
		month := monthNames[i%12]
		axisLabels24Months = append(axisLabels24Months, fmt.Sprintf("%s '%02d", month, year%100))
		series24Months[0].Values = append(series24Months[0].Values, 26+float64((i*6)%36))
		series24Months[1].Values = append(series24Months[1].Values, 10+float64((i*3)%22))
		series24Months[2].Values = append(series24Months[2].Values, -float64(1+i+(i%3)*2))
	}

	return myrtle.NewBuilder(
		selectedTheme,
		myrtle.WithStyles(theme.Styles{ColorPrimary: "#0ea5e9"}),
	).
		WithPreheader("Vertical stacked bars with legend and axis options").
		WithHeader(commonHeaderGroup("Myrtle", selectedTheme)).
		AddHeading("Net revenue composition", myrtle.HeadingLevel(2)).
		AddText("Shows stacked gains and losses per month with a visible baseline for negative values.").
		AddVerticalBarChart(
			axisLabels,
			series,
			myrtle.VerticalBarChartTitle("MRR movement by month"),
			myrtle.VerticalBarChartHeight(190),
			myrtle.VerticalBarChartLegendPlacement(myrtle.VerticalBarChartLegendBottom),
			myrtle.VerticalBarChartAxisShowYTicks(true),
			myrtle.VerticalBarChartAxisShowBaseline(true),
		).
		AddHeading("Inline value labels", myrtle.HeadingLevel(3)).
		AddVerticalBarChart(
			axisLabels,
			series,
			myrtle.VerticalBarChartTitle("MRR movement by month"),
			myrtle.VerticalBarChartValueFormatterOption(myrtle.VerticalBarChartValueFormatter{Prefix: "€"}),
			myrtle.VerticalBarChartLegendConfigOption(myrtle.VerticalBarChartLegendConfig{Placement: myrtle.VerticalBarChartLegendBottom}),
			myrtle.VerticalBarChartAxisConfig(myrtle.VerticalBarChartAxis{ShowYTicks: true, ShowBaseline: true, LabelFormat: myrtle.VerticalBarChartAxisLabelFormatNumber}),
			myrtle.VerticalBarChartValueLabelsOption(myrtle.VerticalBarChartValueLabels{Show: true, MinSegmentHeight: 14, Color: "#ffffff"}),
		).
		AddHeading("24 months (full width default)", myrtle.HeadingLevel(3)).
		AddVerticalBarChart(
			axisLabels24Months,
			series24Months,
			myrtle.VerticalBarChartTitle("MRR movement over 24 months"),
			myrtle.VerticalBarChartHeight(170),
			myrtle.VerticalBarChartColumnGap(1),
			myrtle.VerticalBarChartAxisShowCategoryLabels(false),
			myrtle.VerticalBarChartAxisShowYTicks(true),
		).
		AddHeading("Normalized columns (positive-only series)", myrtle.HeadingLevel(3)).
		AddVerticalBarChart(
			axisLabels,
			normalizedSeries,
			myrtle.VerticalBarChartTitle("MRR composition"),
			myrtle.VerticalBarChartNormalize(true),
			myrtle.VerticalBarChartLegendPlacement(myrtle.VerticalBarChartLegendBottom),
			myrtle.VerticalBarChartColumnGap(1),
		).
		AddHeading("Normalized columns with 0%-100% axis", myrtle.HeadingLevel(3)).
		AddVerticalBarChart(
			axisLabels,
			normalizedPercentSeries,
			myrtle.VerticalBarChartTitle("Composition share"),
			myrtle.VerticalBarChartNormalize(true),
			myrtle.VerticalBarChartColumnGap(1),
			myrtle.VerticalBarChartAxisConfig(myrtle.VerticalBarChartAxis{ShowYTicks: true, LabelFormat: myrtle.VerticalBarChartAxisLabelFormatPercent, HasMin: true, Min: 0, HasMax: true, Max: 100}),
			myrtle.VerticalBarChartLegendPlacement(myrtle.VerticalBarChartLegendBottom),
			myrtle.VerticalBarChartValueLabelsOption(myrtle.VerticalBarChartValueLabels{Show: true, MinSegmentHeight: 14, Color: "#ffffff"}),
		).
		AddHeading("Tiny segment label fallback (single series)", myrtle.HeadingLevel(3)).
		AddVerticalBarChart(
			axisLabels,
			tooSmallSingleSeries,
			myrtle.VerticalBarChartTitle("Small values above the bar"),
			myrtle.VerticalBarChartHeight(180),
			myrtle.VerticalBarChartAxisConfig(myrtle.VerticalBarChartAxis{ShowYTicks: true, ShowBaseline: true, LabelFormat: myrtle.VerticalBarChartAxisLabelFormatNumber, HasMin: true, Min: 0, HasMax: true, Max: 100}),
			myrtle.VerticalBarChartValueLabelsOption(myrtle.VerticalBarChartValueLabels{Show: true, MinSegmentHeight: 12, Color: "#ffffff"}),
		).
		AddHeading("Tiny segment label fallback (multiple series)", myrtle.HeadingLevel(3)).
		AddVerticalBarChart(
			axisLabels,
			tooSmallMultiSeries,
			myrtle.VerticalBarChartTitle("Small top positive + small upper negative"),
			myrtle.VerticalBarChartHeight(190),
			myrtle.VerticalBarChartLegendPlacement(myrtle.VerticalBarChartLegendBottom),
			myrtle.VerticalBarChartAxisConfig(myrtle.VerticalBarChartAxis{ShowYTicks: true, ShowBaseline: true, LabelFormat: myrtle.VerticalBarChartAxisLabelFormatNumber, HasMin: true, Min: -100, HasMax: true, Max: 140}),
			myrtle.VerticalBarChartValueLabelsOption(myrtle.VerticalBarChartValueLabels{Show: true, MinSegmentHeight: 12, Color: "#ffffff"}),
		).
		AddHeading("Tiny segment fallback boundary (no free space)", myrtle.HeadingLevel(3)).
		AddVerticalBarChart(
			axisLabels,
			tooSmallNoSpaceSeries,
			myrtle.VerticalBarChartTitle("Small values stay hidden"),
			myrtle.VerticalBarChartHeight(180),
			myrtle.VerticalBarChartAxisConfig(myrtle.VerticalBarChartAxis{ShowYTicks: true, ShowBaseline: true, LabelFormat: myrtle.VerticalBarChartAxisLabelFormatNumber, HasMin: true, Min: 0, HasMax: true, Max: 100}),
			myrtle.VerticalBarChartValueLabelsOption(myrtle.VerticalBarChartValueLabels{Show: true, MinSegmentHeight: 12, Color: "#ffffff"}),
		).
		AddText("Expected: tiny labels remain hidden when the stack reaches the top and no above-label space is available.", myrtle.TextTone(myrtle.ToneMuted), myrtle.TextSize(myrtle.TextSizeSmall)).
		AddCallout(myrtle.ToneInfo, "Tip", "Use axis min/max overrides when comparing multiple reports to preserve scale.").
		WithFooter(commonFooterGroup()).
		Build(), nil
}
