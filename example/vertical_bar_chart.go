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
		{Key: "base", Label: "Base", Values: []float64{95, 92, 94, 93}},
		{Key: "small", Label: "Small", Values: []float64{5, 8, 6, 7}},
	}
	axisLabels := []string{"Jan", "Feb", "Mar", "Apr"}

	monthNames := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	axisLabels24Months := make([]string, 0, 24)
	series24Months := []myrtle.VerticalBarChartSeries{
		{Key: "new", Label: "New", Color: "#2563eb", Values: make([]float64, 0, 24)},
		{Key: "expansion", Label: "Expansion", Color: "#16a34a", Values: make([]float64, 0, 24)},
		{Key: "churn", Label: "Churn", Color: "#dc2626", Values: make([]float64, 0, 24)},
	}
	for i := 0; i < 24; i++ {
		year := 2024
		if i >= 12 {
			year = 2025
		}
		month := monthNames[i%12]
		axisLabels24Months = append(axisLabels24Months, fmt.Sprintf("%s '%02d", month, year%100))
		series24Months[0].Values = append(series24Months[0].Values, 26+float64((i*5)%18))
		series24Months[1].Values = append(series24Months[1].Values, 10+float64((i*3)%11))
		series24Months[2].Values = append(series24Months[2].Values, -float64(6+(i*4)%10))
	}

	return myrtle.NewBuilder(
		selectedTheme,
		myrtle.WithStyles(theme.Styles{ColorPrimary: "#0ea5e9"}),
	).
		WithPreheader("Vertical stacked bars with legend and axis options").
		WithHeader(commonHeaderGroup("Myrtle")).
		AddHeading("Net revenue composition", myrtle.HeadingLevel(2)).
		AddText("Shows stacked gains and losses per month with a visible baseline for negative values.").
		AddVerticalBarChart(
			"MRR movement by month",
			axisLabels,
			series,
			myrtle.VerticalBarChartHeight(190),
			myrtle.VerticalBarChartLegendPlacement(myrtle.VerticalBarChartLegendBottom),
			myrtle.VerticalBarChartAxisShowYTicks(true),
			myrtle.VerticalBarChartAxisTickCount(5),
			myrtle.VerticalBarChartAxisShowBaseline(true),
		).
		AddHeading("Inline value labels", myrtle.HeadingLevel(3)).
		AddVerticalBarChart(
			"MRR movement by month",
			axisLabels,
			series,
			myrtle.VerticalBarChartValueFormatterOption(myrtle.VerticalBarChartValueFormatter{Prefix: "€"}),
			myrtle.VerticalBarChartLegendConfigOption(myrtle.VerticalBarChartLegendConfig{Placement: myrtle.VerticalBarChartLegendBottom}),
			myrtle.VerticalBarChartAxisConfig(myrtle.VerticalBarChartAxis{ShowYTicks: true, ShowBaseline: true, LabelFormat: myrtle.VerticalBarChartAxisLabelFormatNumber}),
			myrtle.VerticalBarChartValueLabelsOption(myrtle.VerticalBarChartValueLabels{Show: true, MinSegmentHeight: 14, Color: "#ffffff"}),
		).
		AddHeading("1px spacing between columns", myrtle.HeadingLevel(3)).
		AddVerticalBarChart(
			"Compact bars",
			axisLabels,
			series,
			myrtle.VerticalBarChartColumnGap(1),
			myrtle.VerticalBarChartLegendPlacement(myrtle.VerticalBarChartLegendBottom),
		).
		AddHeading("24 months (full width default)", myrtle.HeadingLevel(3)).
		AddVerticalBarChart(
			"MRR movement over 24 months",
			axisLabels24Months,
			series24Months,
			myrtle.VerticalBarChartHeight(170),
			myrtle.VerticalBarChartColumnGap(1),
			myrtle.VerticalBarChartAxisShowCategoryLabels(false),
			myrtle.VerticalBarChartAxisShowYTicks(true),
			myrtle.VerticalBarChartAxisTickCount(4),
		).
		AddHeading("Normalized columns (positive-only series)", myrtle.HeadingLevel(3)).
		AddVerticalBarChart(
			"MRR composition",
			axisLabels,
			normalizedSeries,
			myrtle.VerticalBarChartNormalize(true),
			myrtle.VerticalBarChartLegendPlacement(myrtle.VerticalBarChartLegendBottom),
			myrtle.VerticalBarChartColumnGap(1),
		).
		AddHeading("Normalized columns with 0%-100% axis", myrtle.HeadingLevel(3)).
		AddVerticalBarChart(
			"Composition share",
			axisLabels,
			normalizedPercentSeries,
			myrtle.VerticalBarChartNormalize(true),
			myrtle.VerticalBarChartColumnGap(1),
			myrtle.VerticalBarChartAxisConfig(myrtle.VerticalBarChartAxis{ShowYTicks: true, LabelFormat: myrtle.VerticalBarChartAxisLabelFormatPercent, HasMin: true, Min: 0, HasMax: true, Max: 100}),
			myrtle.VerticalBarChartLegendPlacement(myrtle.VerticalBarChartLegendBottom),
			myrtle.VerticalBarChartValueLabelsOption(myrtle.VerticalBarChartValueLabels{Show: true, MinSegmentHeight: 14, Color: "#ffffff"}),
		).
		AddHeading("Tiny segment label fallback (single series)", myrtle.HeadingLevel(3)).
		AddVerticalBarChart(
			"Small values above the bar",
			axisLabels,
			tooSmallSingleSeries,
			myrtle.VerticalBarChartHeight(180),
			myrtle.VerticalBarChartAxisConfig(myrtle.VerticalBarChartAxis{ShowYTicks: true, ShowBaseline: true, LabelFormat: myrtle.VerticalBarChartAxisLabelFormatNumber, HasMin: true, Min: 0, HasMax: true, Max: 100}),
			myrtle.VerticalBarChartValueLabelsOption(myrtle.VerticalBarChartValueLabels{Show: true, MinSegmentHeight: 12, Color: "#ffffff"}),
		).
		AddHeading("Tiny segment label fallback (multiple series)", myrtle.HeadingLevel(3)).
		AddVerticalBarChart(
			"Small top positive + small upper negative",
			axisLabels,
			tooSmallMultiSeries,
			myrtle.VerticalBarChartHeight(190),
			myrtle.VerticalBarChartLegendPlacement(myrtle.VerticalBarChartLegendBottom),
			myrtle.VerticalBarChartAxisConfig(myrtle.VerticalBarChartAxis{ShowYTicks: true, ShowBaseline: true, LabelFormat: myrtle.VerticalBarChartAxisLabelFormatNumber, HasMin: true, Min: -100, HasMax: true, Max: 140}),
			myrtle.VerticalBarChartValueLabelsOption(myrtle.VerticalBarChartValueLabels{Show: true, MinSegmentHeight: 12, Color: "#ffffff"}),
		).
		AddHeading("Tiny segment fallback boundary (no free space)", myrtle.HeadingLevel(3)).
		AddVerticalBarChart(
			"Small values stay hidden",
			axisLabels,
			tooSmallNoSpaceSeries,
			myrtle.VerticalBarChartHeight(180),
			myrtle.VerticalBarChartAxisConfig(myrtle.VerticalBarChartAxis{ShowYTicks: true, ShowBaseline: true, LabelFormat: myrtle.VerticalBarChartAxisLabelFormatNumber, HasMin: true, Min: 0, HasMax: true, Max: 100}),
			myrtle.VerticalBarChartValueLabelsOption(myrtle.VerticalBarChartValueLabels{Show: true, MinSegmentHeight: 12, Color: "#ffffff"}),
		).
		AddText("Expected: tiny labels remain hidden when the stack reaches the top and no above-label space is available.", myrtle.TextTone(myrtle.TextToneMuted), myrtle.TextSize(myrtle.TextSizeSmall)).
		AddCallout(myrtle.CalloutTypeInfo, "Tip", "Use axis min/max overrides when comparing multiple reports to preserve scale.").
		WithFooter(commonFooterGroup()).
		Build(), nil
}
