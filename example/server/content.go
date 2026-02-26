package server

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gzuidhof/myrtle"
	"github.com/gzuidhof/myrtle/theme"
	defaulttheme "github.com/gzuidhof/myrtle/theme/default"
	"github.com/gzuidhof/myrtle/theme/flat"
	"github.com/gzuidhof/myrtle/theme/terminal"
)

func buildExampleEmail(name string, selectedTheme theme.Theme) (*myrtle.Email, error) {
	for _, emailBuilder := range exampleEmails {
		if emailBuilder.Name == name {
			return emailBuilder.Build(selectedTheme)
		}
	}

	return nil, fmt.Errorf("unknown example email: %s", name)
}

func buildBlockEmail(name string, selectedTheme theme.Theme) (*myrtle.Email, error) {
	builder := myrtle.NewBuilder(selectedTheme).
		WithoutHeader()

	switch name {
	case "text":
		builder.AddText("This is a text block.")

		builder.AddHeading("Tone variants", myrtle.HeadingLevel(3))
		builder.AddText("Tone default", myrtle.TextTone(myrtle.TextToneDefault))
		builder.AddText("Tone muted", myrtle.TextTone(myrtle.TextToneMuted))
		builder.AddText("Tone info", myrtle.TextTone(myrtle.TextToneInfo))
		builder.AddText("Tone success", myrtle.TextTone(myrtle.TextToneSuccess))
		builder.AddText("Tone warning", myrtle.TextTone(myrtle.TextToneWarning))
		builder.AddText("Tone danger", myrtle.TextTone(myrtle.TextToneDanger))

		builder.AddHeading("Size variants", myrtle.HeadingLevel(3))
		builder.AddText("Size small", myrtle.TextSize(myrtle.TextSizeSmall))
		builder.AddText("Size base", myrtle.TextSize(myrtle.TextSizeBase))
		builder.AddText("Size large", myrtle.TextSize(myrtle.TextSizeLarge))

		builder.AddHeading("Align variants", myrtle.HeadingLevel(3))
		builder.AddText("Align start", myrtle.TextAlign(myrtle.TextAlignStart))
		builder.AddText("Align center", myrtle.TextAlign(myrtle.TextAlignCenter))
		builder.AddText("Align end", myrtle.TextAlign(myrtle.TextAlignEnd))

		builder.AddHeading("Weight variants", myrtle.HeadingLevel(3))
		builder.AddText("Weight normal", myrtle.TextWeight(myrtle.TextWeightNormal))
		builder.AddText("Weight medium", myrtle.TextWeight(myrtle.TextWeightMedium))
		builder.AddText("Weight semibold", myrtle.TextWeight(myrtle.TextWeightSemibold))
		builder.AddText("Weight bold", myrtle.TextWeight(myrtle.TextWeightBold))

		builder.AddHeading("Spacing variants", myrtle.HeadingLevel(3))
		builder.AddText("Spacing compact: This paragraph intentionally contains enough words to wrap over multiple lines in the preview so the compact line-height is easy to compare.", myrtle.TextSpacing(myrtle.TextSpacingCompact))
		builder.AddText("Spacing normal: This paragraph intentionally contains enough words to wrap over multiple lines in the preview so the normal line-height is easy to compare.", myrtle.TextSpacing(myrtle.TextSpacingNormal))
		builder.AddText("Spacing relaxed: This paragraph intentionally contains enough words to wrap over multiple lines in the preview so the relaxed line-height is easy to compare.", myrtle.TextSpacing(myrtle.TextSpacingRelaxed))

		builder.AddHeading("Transform variants", myrtle.HeadingLevel(3))
		builder.AddText("Transform none", myrtle.TextTransform(myrtle.TextTransformNone))
		builder.AddText("Transform uppercase", myrtle.TextTransform(myrtle.TextTransformUppercase))
		builder.AddText("TRANSFORM LOWERCASE", myrtle.TextTransform(myrtle.TextTransformLowercase))
		builder.AddText("transform capitalize variant", myrtle.TextTransform(myrtle.TextTransformCapitalize))

		builder.AddHeading("No margin", myrtle.HeadingLevel(3))
		builder.AddText("This line has no bottom margin.", myrtle.TextNoMargin(true))
		builder.AddText("This line has default spacing after it.")

		builder.AddHeading("Combined styles", myrtle.HeadingLevel(3))
		builder.AddText("Centered semibold info", myrtle.TextTone(myrtle.TextToneInfo), myrtle.TextAlign(myrtle.TextAlignCenter), myrtle.TextWeight(myrtle.TextWeightSemibold))
		builder.AddText("Compact uppercase warning", myrtle.TextTone(myrtle.TextToneWarning), myrtle.TextSize(myrtle.TextSizeSmall), myrtle.TextSpacing(myrtle.TextSpacingCompact), myrtle.TextTransform(myrtle.TextTransformUppercase), myrtle.TextNoMargin(true))
		builder.AddText("Relaxed large success", myrtle.TextTone(myrtle.TextToneSuccess), myrtle.TextSize(myrtle.TextSizeLarge), myrtle.TextSpacing(myrtle.TextSpacingRelaxed))
	case "heading":
		builder.AddHeading("This is a heading", myrtle.HeadingLevel(2))
	case "spacer":
		builder.AddHeading("Fixed size", myrtle.HeadingLevel(3))
		builder.AddText("Content above spacer").AddSpacer(myrtle.SpacerSize(24)).AddText("Content below spacer")
		builder.AddHeading("Custom sizes", myrtle.HeadingLevel(3))
		builder.AddText("8px").AddSpacer(myrtle.SpacerSize(8)).AddText("12px").AddSpacer(myrtle.SpacerSize(12)).AddText("16px default").AddSpacer().AddText("24px").AddSpacer(myrtle.SpacerSize(24))
	case "list":
		builder.AddHeading("Unordered", myrtle.HeadingLevel(3))
		builder.AddList([]string{"First item", "Second item", "Third item"}, false)
		builder.AddHeading("Ordered", myrtle.HeadingLevel(3))
		builder.AddList([]string{"First item", "Second item", "Third item"}, true)
	case "key-value":
		builder.AddKeyValue("Order details", []myrtle.KeyValuePair{{Key: "Order", Value: "#4821"}, {Key: "Total", Value: "$129.00"}, {Key: "Status", Value: "Shipped"}})
	case "bar-chart":
		builder.AddHeading("Default", myrtle.HeadingLevel(3))
		builder.AddHorizontalBarChart("Delivery by region", []myrtle.HorizontalBarChartItem{
			{Label: "US", Value: "52%", Percent: 52},
			{Label: "EMEA", Value: "31%", Percent: 31},
			{Label: "APAC", Value: "17%", Percent: 17},
		})
		builder.AddHeading("Thicker bars", myrtle.HeadingLevel(3))
		builder.AddHorizontalBarChart("Delivery by region", []myrtle.HorizontalBarChartItem{
			{Label: "US", Value: "52%", Percent: 52},
			{Label: "EMEA", Value: "31%", Percent: 31},
			{Label: "APAC", Value: "17%", Percent: 17},
		}, myrtle.HorizontalBarChartThickness(14))
		builder.AddHeading("Transparent background", myrtle.HeadingLevel(3))
		builder.AddHorizontalBarChart("Delivery by region", []myrtle.HorizontalBarChartItem{
			{Label: "US", Value: "52%", Percent: 52},
			{Label: "EMEA", Value: "31%", Percent: 31},
			{Label: "APAC", Value: "17%", Percent: 17},
		}, myrtle.HorizontalBarChartTransparentBackground(true))
		builder.AddHeading("Tone: success", myrtle.HeadingLevel(3))
		builder.AddHorizontalBarChart("Delivery by region", []myrtle.HorizontalBarChartItem{
			{Label: "US", Value: "52%", Percent: 52},
			{Label: "EMEA", Value: "31%", Percent: 31},
			{Label: "APAC", Value: "17%", Percent: 17},
		}, myrtle.HorizontalBarChartTone(myrtle.ChartToneSuccess))
		builder.AddHeading("Tone: warning", myrtle.HeadingLevel(3))
		builder.AddHorizontalBarChart("Delivery by region", []myrtle.HorizontalBarChartItem{
			{Label: "US", Value: "52%", Percent: 52},
			{Label: "EMEA", Value: "31%", Percent: 31},
			{Label: "APAC", Value: "17%", Percent: 17},
		}, myrtle.HorizontalBarChartTone(myrtle.ChartToneWarning))
	case "vertical-bar-chart":
		axisLabels := []string{"Jan", "Feb", "Mar", "Apr"}
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
		monthNames := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
		axisLabels24Months := make([]string, 0, 24)
		series24Months := []myrtle.VerticalBarChartSeries{
			{Key: "new", Label: "New", Color: "#1d4ed8", ValueLabelColor: "#ffffff", Values: make([]float64, 0, 24)},
			{Key: "expansion", Label: "Expansion", Color: "#3b82f6", ValueLabelColor: "#ffffff", Values: make([]float64, 0, 24)},
			{Key: "churn", Label: "Churn", Color: "#ef4444", ValueLabelColor: "#ffffff", Values: make([]float64, 0, 24)},
		}
		for i := 0; i < 24; i++ {
			year := 2024
			if i >= 12 {
				year = 2025
			}
			axisLabels24Months = append(axisLabels24Months, fmt.Sprintf("%s '%02d", monthNames[i%12], year%100))
			series24Months[0].Values = append(series24Months[0].Values, 26+float64((i*5)%18))
			series24Months[1].Values = append(series24Months[1].Values, 10+float64((i*3)%11))
			series24Months[2].Values = append(series24Months[2].Values, -float64(6+(i*4)%10))
		}

		builder.AddHeading("Default mixed values", myrtle.HeadingLevel(3))
		builder.AddVerticalBarChart("MRR movement", axisLabels, series)

		builder.AddHeading("Legend + ticks + baseline", myrtle.HeadingLevel(3))
		builder.AddVerticalBarChart(
			"MRR movement",
			axisLabels,
			series,
			myrtle.VerticalBarChartHeight(190),
			myrtle.VerticalBarChartLegendPlacement(myrtle.VerticalBarChartLegendBottom),
			myrtle.VerticalBarChartAxisShowYTicks(true),
			myrtle.VerticalBarChartAxisTickCount(5),
			myrtle.VerticalBarChartAxisShowBaseline(true),
		)

		builder.AddHeading("Fixed axis range + no category labels", myrtle.HeadingLevel(3))
		builder.AddVerticalBarChart(
			"MRR movement",
			axisLabels,
			series,
			myrtle.VerticalBarChartAxisMin(-30),
			myrtle.VerticalBarChartAxisMax(90),
			myrtle.VerticalBarChartAxisShowYTicks(true),
			myrtle.VerticalBarChartAxisLabelFormat(myrtle.VerticalBarChartAxisLabelFormatNumber),
			myrtle.VerticalBarChartAxisShowCategoryLabels(false),
			myrtle.VerticalBarChartColumnGap(8),
		)

		builder.AddHeading("Column gap comparison (0, 4, 12)", myrtle.HeadingLevel(3))
		builder.AddGrid(
			[]myrtle.GridItem{
				myrtle.GridItemGroup(myrtle.NewGroup().AddVerticalBarChart("Gap 0", axisLabels, series, myrtle.VerticalBarChartColumnGap(0))),
				myrtle.GridItemGroup(myrtle.NewGroup().AddVerticalBarChart("Gap 4", axisLabels, series, myrtle.VerticalBarChartColumnGap(4))),
				myrtle.GridItemGroup(myrtle.NewGroup().AddVerticalBarChart("Gap 12", axisLabels, series, myrtle.VerticalBarChartColumnGap(12))),
			},
			myrtle.GridColumns(3),
		)

		builder.AddHeading("1px spacing between columns", myrtle.HeadingLevel(3))
		builder.AddVerticalBarChart(
			"MRR movement",
			axisLabels,
			series,
			myrtle.VerticalBarChartColumnGap(1),
			myrtle.VerticalBarChartLegendPlacement(myrtle.VerticalBarChartLegendBottom),
		)

		builder.AddHeading("Normalized columns (positive-only series)", myrtle.HeadingLevel(3))
		builder.AddVerticalBarChart(
			"MRR composition",
			axisLabels,
			normalizedSeries,
			myrtle.VerticalBarChartColumnGap(1),
			myrtle.VerticalBarChartNormalize(true),
			myrtle.VerticalBarChartLegendPlacement(myrtle.VerticalBarChartLegendBottom),
		)

		builder.AddHeading("Normalized with 0%-100% axis", myrtle.HeadingLevel(3))
		builder.AddVerticalBarChart(
			"Composition share",
			axisLabels,
			normalizedPercentSeries,
			myrtle.VerticalBarChartNormalize(true),
			myrtle.VerticalBarChartColumnGap(1),
			myrtle.VerticalBarChartAxisConfig(myrtle.VerticalBarChartAxis{ShowYTicks: true, LabelFormat: myrtle.VerticalBarChartAxisLabelFormatPercent, HasMin: true, Min: 0, HasMax: true, Max: 100}),
			myrtle.VerticalBarChartLegendPlacement(myrtle.VerticalBarChartLegendBottom),
			myrtle.VerticalBarChartValueLabelsOption(myrtle.VerticalBarChartValueLabels{Show: true, MinSegmentHeight: 14, Color: "#ffffff"}),
		)

		builder.AddHeading("24 months (full width default)", myrtle.HeadingLevel(3))
		builder.AddVerticalBarChart(
			"MRR movement",
			axisLabels24Months,
			series24Months,
			myrtle.VerticalBarChartHeight(170),
			myrtle.VerticalBarChartColumnGap(1),
			myrtle.VerticalBarChartAxisShowCategoryLabels(false),
			myrtle.VerticalBarChartAxisShowYTicks(true),
			myrtle.VerticalBarChartAxisTickCount(4),
		)

		builder.AddHeading("Struct config + in-bar value labels", myrtle.HeadingLevel(3))
		builder.AddVerticalBarChart(
			"MRR movement",
			axisLabels,
			series,
			myrtle.VerticalBarChartValueFormatterOption(myrtle.VerticalBarChartValueFormatter{Prefix: "€"}),
			myrtle.VerticalBarChartAxisConfig(myrtle.VerticalBarChartAxis{ShowYTicks: true, ShowBaseline: true, LabelFormat: myrtle.VerticalBarChartAxisLabelFormatNumber}),
			myrtle.VerticalBarChartLegendConfigOption(myrtle.VerticalBarChartLegendConfig{Placement: myrtle.VerticalBarChartLegendBottom}),
			myrtle.VerticalBarChartValueLabelsOption(myrtle.VerticalBarChartValueLabels{Show: true, MinSegmentHeight: 14, Color: "#ffffff"}),
		)

		builder.AddHeading("Tiny segment label fallback (single series)", myrtle.HeadingLevel(3))
		builder.AddVerticalBarChart(
			"Small values above the bar",
			axisLabels,
			tooSmallSingleSeries,
			myrtle.VerticalBarChartHeight(180),
			myrtle.VerticalBarChartAxisConfig(myrtle.VerticalBarChartAxis{ShowYTicks: true, ShowBaseline: true, LabelFormat: myrtle.VerticalBarChartAxisLabelFormatNumber, HasMin: true, Min: 0, HasMax: true, Max: 100}),
			myrtle.VerticalBarChartValueLabelsOption(myrtle.VerticalBarChartValueLabels{Show: true, MinSegmentHeight: 12, Color: "#ffffff"}),
		)

		builder.AddHeading("Tiny segment label fallback (multiple series)", myrtle.HeadingLevel(3))
		builder.AddVerticalBarChart(
			"Small top positive + small upper negative",
			axisLabels,
			tooSmallMultiSeries,
			myrtle.VerticalBarChartHeight(190),
			myrtle.VerticalBarChartLegendPlacement(myrtle.VerticalBarChartLegendBottom),
			myrtle.VerticalBarChartAxisConfig(myrtle.VerticalBarChartAxis{ShowYTicks: true, ShowBaseline: true, LabelFormat: myrtle.VerticalBarChartAxisLabelFormatNumber, HasMin: true, Min: -100, HasMax: true, Max: 140}),
			myrtle.VerticalBarChartValueLabelsOption(myrtle.VerticalBarChartValueLabels{Show: true, MinSegmentHeight: 12, Color: "#ffffff"}),
		)

		builder.AddHeading("Tiny segment fallback boundary (no free space)", myrtle.HeadingLevel(3))
		builder.AddVerticalBarChart(
			"Small values stay hidden",
			axisLabels,
			tooSmallNoSpaceSeries,
			myrtle.VerticalBarChartHeight(180),
			myrtle.VerticalBarChartAxisConfig(myrtle.VerticalBarChartAxis{ShowYTicks: true, ShowBaseline: true, LabelFormat: myrtle.VerticalBarChartAxisLabelFormatNumber, HasMin: true, Min: 0, HasMax: true, Max: 100}),
			myrtle.VerticalBarChartValueLabelsOption(myrtle.VerticalBarChartValueLabels{Show: true, MinSegmentHeight: 12, Color: "#ffffff"}),
		)
		builder.AddText("Expected: tiny labels remain hidden when the stack reaches the top and no above-label space is available.", myrtle.TextTone(myrtle.TextToneMuted), myrtle.TextSize(myrtle.TextSizeSmall))

		builder.AddHeading("Large values with compact formatter", myrtle.HeadingLevel(3))
		builder.AddVerticalBarChart(
			"Volume by quarter",
			[]string{"Q1", "Q2", "Q3", "Q4"},
			[]myrtle.VerticalBarChartSeries{{Key: "volume", Label: "Volume", Values: []float64{100000, 1300000, 10200, -1000}}},
			myrtle.VerticalBarChartAxisConfig(myrtle.VerticalBarChartAxis{ShowYTicks: true, ShowBaseline: true, LabelFormat: myrtle.VerticalBarChartAxisLabelFormatNumber, HasMin: true, Min: -1000, HasMax: true, Max: 1300000}),
			myrtle.VerticalBarChartValueFormatterOption(myrtle.VerticalBarChartValueFormatter{MagnitudeSuffix: myrtle.VerticalBarChartMagnitudeSuffixShort}),
			myrtle.VerticalBarChartValueLabelsOption(myrtle.VerticalBarChartValueLabels{Show: true, MinSegmentHeight: 10, Color: "#ffffff"}),
		)
	case "sparkline":
		builder.AddHeading("Signups", myrtle.HeadingLevel(3))
		builder.AddSparkline("Weekly trend", "Signups", "1,204", []int{8, 12, 9, 14, 18, 16, 20}, myrtle.SparklineDelta("+8%"), myrtle.SparklineDeltaSemantic(myrtle.StatDeltaSemanticPositive))
		builder.AddHeading("Revenue", myrtle.HeadingLevel(3))
		builder.AddSparkline("Weekly trend", "Revenue", "$18.2k", []int{15, 14, 17, 16, 19, 21, 20})
		builder.AddHeading("Incidents", myrtle.HeadingLevel(3))
		builder.AddSparkline("Weekly trend", "Incidents", "12", []int{20, 18, 16, 15, 13, 12, 10}, myrtle.SparklineDelta("-2"), myrtle.SparklineDeltaSemantic(myrtle.StatDeltaSemanticNegative))
		builder.AddHeading("Tone: info", myrtle.HeadingLevel(3))
		builder.AddSparkline("Weekly trend", "Queue depth", "312", []int{10, 11, 13, 12, 14, 16, 15}, myrtle.SparklineTone(myrtle.ChartToneInfo))
		builder.AddHeading("Tone: danger", myrtle.HeadingLevel(3))
		builder.AddSparkline("Weekly trend", "Error rate", "2.1%", []int{8, 9, 10, 12, 15, 17, 16}, myrtle.SparklineTone(myrtle.ChartToneDanger))
	case "stacked-bar":
		builder.AddStackedBar("Channel mix", []myrtle.StackedBarRow{
			{Label: "Acquisition", Segments: []myrtle.StackedBarSegment{{Label: "Email", Percent: 58, Value: "58%"}, {Label: "SMS", Percent: 24, Value: "24%"}, {Label: "Push", Percent: 18, Value: "18%"}}},
			{Label: "Activation", Segments: []myrtle.StackedBarSegment{{Label: "Email", Percent: 42, Value: "42%"}, {Label: "SMS", Percent: 33, Value: "33%"}, {Label: "Push", Percent: 25, Value: "25%"}}},
		}, myrtle.StackedBarTotal("Total", "120k"))
		builder.AddHeading("Tone: info", myrtle.HeadingLevel(3))
		builder.AddStackedBar("Channel mix", []myrtle.StackedBarRow{
			{Label: "Acquisition", Segments: []myrtle.StackedBarSegment{{Label: "Email", Percent: 58, Value: "58%"}, {Label: "SMS", Percent: 24, Value: "24%"}, {Label: "Push", Percent: 18, Value: "18%"}}},
			{Label: "Activation", Segments: []myrtle.StackedBarSegment{{Label: "Email", Percent: 42, Value: "42%"}, {Label: "SMS", Percent: 33, Value: "33%"}, {Label: "Push", Percent: 25, Value: "25%"}}},
		}, myrtle.StackedBarTotal("Total", "120k"), myrtle.StackedBarTone(myrtle.ChartToneInfo))
		builder.AddHeading("Tone: muted", myrtle.HeadingLevel(3))
		builder.AddStackedBar("Channel mix", []myrtle.StackedBarRow{
			{Label: "Acquisition", Segments: []myrtle.StackedBarSegment{{Label: "Email", Percent: 58, Value: "58%"}, {Label: "SMS", Percent: 24, Value: "24%"}, {Label: "Push", Percent: 18, Value: "18%"}}},
			{Label: "Activation", Segments: []myrtle.StackedBarSegment{{Label: "Email", Percent: 42, Value: "42%"}, {Label: "SMS", Percent: 33, Value: "33%"}, {Label: "Push", Percent: 25, Value: "25%"}}},
		}, myrtle.StackedBarTotal("Total", "120k"), myrtle.StackedBarTone(myrtle.ChartToneMuted))
	case "progress":
		builder.AddProgress("Rollout progress", []myrtle.ProgressItem{{Label: "Schema migration", Percent: 100}, {Label: "API deploy", Percent: 76}, {Label: "Client rollout", Percent: 48}})
	case "distribution":
		builder.AddDistribution("Latency distribution (ms)", []myrtle.DistributionBucket{{Label: "0-50", Count: 62}, {Label: "51-100", Count: 44}, {Label: "101-200", Count: 21}, {Label: "200+", Count: 8}})
	case "timeline":
		builder.AddTimeline("Incident timeline", []myrtle.TimelineItem{
			{Time: "09:07", Title: "Detected", Detail: "Elevated webhook latency"},
			{Time: "09:18", Title: "Mitigation", Detail: "Queue workers scaled to 2x"},
			{Time: "09:42", Title: "Resolved", Detail: "Latency returned to baseline"},
		}, myrtle.TimelineAggregateHeader("3 events · currently mitigating"), myrtle.TimelineCurrentIndex(1))
	case "stats-row":
		builder.AddStatsRow("Weekly KPIs", []myrtle.StatItem{
			{Label: "Delivery", Value: "99.8%", Delta: "+0.3%", DeltaSemantic: myrtle.StatDeltaSemanticPositive},
			{Label: "CTR", Value: "4.1%", Delta: "+0.4%", DeltaSemantic: myrtle.StatDeltaSemanticPositive},
			{Label: "Bounces", Value: "0.9%", Delta: "-0.1%", DeltaSemantic: myrtle.StatDeltaSemanticNegative},
		})
	case "badge":
		builder.AddHeading("Info", myrtle.HeadingLevel(3))
		builder.AddBadge(myrtle.BadgeToneInfo, "Informational")
		builder.AddHeading("Success", myrtle.HeadingLevel(3))
		builder.AddBadge(myrtle.BadgeToneSuccess, "Operational")
		builder.AddHeading("Warning", myrtle.HeadingLevel(3))
		builder.AddBadge(myrtle.BadgeToneWarning, "Needs review")
		builder.AddHeading("Error", myrtle.HeadingLevel(3))
		builder.AddBadge(myrtle.BadgeToneError, "Action required")
	case "summary-card":
		builder.AddSummaryCard("Deployment complete", "The rollout to production finished successfully with no customer-facing impact.", "Updated 5 minutes ago")
	case "attachment":
		builder.AddAttachment("invoice-Feb-2026.pdf", "PDF · 284 KB", "https://example.com/invoices/feb-2026.pdf", "Download")
	case "hero":
		builder.AddHeading("Default", myrtle.HeadingLevel(3))
		builder.Add(myrtle.HeroBlock{
			Eyebrow:  "New",
			Title:    "Faster sends with Myrtle",
			Body:     "Compose transactional emails with reusable blocks and themeable output.",
			CTALabel: "Read docs",
			CTAURL:   "https://github.com/gzuidhof/myrtle",
		})
		builder.AddHeading("With image", myrtle.HeadingLevel(3))
		builder.Add(myrtle.HeroBlock{
			Eyebrow:  "New",
			Title:    "Faster sends with Myrtle",
			Body:     "Compose transactional emails with reusable blocks and themeable output.",
			CTALabel: "Read docs",
			CTAURL:   "https://github.com/gzuidhof/myrtle",
			ImageURL: "/assets/hero.svg",
			ImageAlt: "Hero image",
		})
		builder.AddHeading("Minimal", myrtle.HeadingLevel(3))
		builder.Add(myrtle.HeroBlock{
			Title: "A compact hero",
			Body:  "Useful for announcement emails that do not need an image or CTA.",
		})
	case "footer-links":
		builder.AddHeading("Default", myrtle.HeadingLevel(3))
		builder.AddFooterLinks([]myrtle.FooterLink{
			{Label: "Help", URL: "https://example.com/help"},
			{Label: "Privacy", URL: "https://example.com/privacy"},
			{Label: "Terms", URL: "https://example.com/terms"},
		}, "You’re receiving this because you have an active account.")
		builder.AddHeading("Note only", myrtle.HeadingLevel(3))
		builder.AddFooterLinks(nil, "You’re receiving this because you have an active account.")
		builder.AddHeading("Links only", myrtle.HeadingLevel(3))
		builder.AddFooterLinks([]myrtle.FooterLink{{Label: "Status", URL: "https://example.com/status"}, {Label: "Help", URL: "https://example.com/help"}}, "")
	case "price-summary":
		builder.AddHeading("Default", myrtle.HeadingLevel(3))
		builder.AddPriceSummary("Order summary", []myrtle.PriceLine{
			{Label: "Subtotal", Value: "$89.00"},
			{Label: "Tax", Value: "$7.12"},
			{Label: "Discount", Value: "-$5.00"},
		}, "Total", "$91.12")
		builder.AddHeading("Minimal", myrtle.HeadingLevel(3))
		builder.AddPriceSummary("Order summary", []myrtle.PriceLine{{Label: "Subtotal", Value: "$89.00"}}, "", "")
		builder.AddHeading("With discount", myrtle.HeadingLevel(3))
		builder.AddPriceSummary("Order summary", []myrtle.PriceLine{{Label: "Subtotal", Value: "$100.00"}, {Label: "Discount", Value: "-$12.00"}, {Label: "Tax", Value: "$7.04"}}, "Total", "$95.04")
	case "empty-state":
		builder.AddHeading("Default", myrtle.HeadingLevel(3))
		builder.AddEmptyState("No incidents", "Everything looks healthy right now.", "", "")
		builder.AddHeading("With action", myrtle.HeadingLevel(3))
		builder.AddEmptyState("No incidents", "Everything looks healthy right now.", "View dashboard", "https://example.com/dashboard")
		builder.AddHeading("Long copy", myrtle.HeadingLevel(3))
		builder.AddEmptyState("No open tasks", "You’re all caught up for now. We’ll notify you when something needs your attention.", "Create workflow", "https://example.com/workflows/new")
	case "quote":
		builder.AddQuote("Myrtle made our transactional emails easy to maintain.", "Engineering Team")
	case "callout":
		builder.AddHeading("Info · soft", myrtle.HeadingLevel(3))
		builder.AddCallout(myrtle.CalloutTypeInfo, "FYI", "Routine update with no action required.", myrtle.CalloutStyle(myrtle.CalloutVariantSoft))
		builder.AddHeading("Info · with link", myrtle.HeadingLevel(3))
		builder.AddCallout(myrtle.CalloutTypeInfo, "FYI", "Routine update with no action required.", myrtle.CalloutStyle(myrtle.CalloutVariantSoft), myrtle.CalloutLink("View details", "https://example.com/details"))
		builder.AddHeading("Success · soft", myrtle.HeadingLevel(3))
		builder.AddCallout(myrtle.CalloutTypeSuccess, "Done", "Deployment completed successfully.", myrtle.CalloutStyle(myrtle.CalloutVariantSoft))
		builder.AddHeading("Warning · outline", myrtle.HeadingLevel(3))
		builder.AddCallout(myrtle.CalloutTypeWarning, "Action needed", "Please verify your billing details before March 1.", myrtle.CalloutStyle(myrtle.CalloutVariantOutline))
		builder.AddHeading("Critical · solid", myrtle.HeadingLevel(3))
		builder.AddCallout(myrtle.CalloutTypeCritical, "Action needed", "Please verify your billing details before March 1.", myrtle.CalloutStyle(myrtle.CalloutVariantSolid))
	case "legal":
		builder.AddLegal("Myrtle Inc.", "123 Market St, SF, CA", "https://example.com/preferences", "https://example.com/unsubscribe")
	case "columns":
		builder.AddHeading("Custom widths + gap + middle align", myrtle.HeadingLevel(3))
		builder.AddColumns(
			myrtle.NewGroup().
				AddHeading("Start column", myrtle.HeadingLevel(3)).
				AddText("Summary and quick context.").
				AddText("Additional details to make this column taller."),
			myrtle.NewGroup().
				AddHeading("End column", myrtle.HeadingLevel(3)).
				AddList([]string{"Point one", "Point two"}, false),
			myrtle.ColumnsWidths(60, 40),
			myrtle.ColumnsGap(24),
			myrtle.ColumnsAlign(myrtle.ColumnsVerticalAlignMiddle),
		)
	case "button":
		builder.AddHeading("Primary", myrtle.HeadingLevel(3))
		builder.AddButton("Open docs", "https://github.com/gzuidhof/myrtle", myrtle.ButtonTone(myrtle.ButtonTonePrimary))
		builder.AddHeading("Primary · centered", myrtle.HeadingLevel(3))
		builder.AddButton("Open docs", "https://github.com/gzuidhof/myrtle", myrtle.ButtonTone(myrtle.ButtonTonePrimary), myrtle.ButtonAlign(myrtle.ButtonAlignmentCenter))
		builder.AddHeading("Primary · end", myrtle.HeadingLevel(3))
		builder.AddButton("Open docs", "https://github.com/gzuidhof/myrtle", myrtle.ButtonTone(myrtle.ButtonTonePrimary), myrtle.ButtonAlign(myrtle.ButtonAlignmentEnd))
		builder.AddHeading("Secondary", myrtle.HeadingLevel(3))
		builder.AddButton("Open docs", "https://github.com/gzuidhof/myrtle", myrtle.ButtonTone(myrtle.ButtonToneSecondary))
		builder.AddHeading("Danger", myrtle.HeadingLevel(3))
		builder.AddButton("Delete workspace", "https://github.com/gzuidhof/myrtle", myrtle.ButtonTone(myrtle.ButtonToneDanger))
		builder.AddHeading("Outline", myrtle.HeadingLevel(3))
		builder.AddButton("Open docs", "https://github.com/gzuidhof/myrtle", myrtle.ButtonStyle(myrtle.ButtonStyleOutline))
		builder.AddHeading("Ghost", myrtle.HeadingLevel(3))
		builder.AddButton("Open docs", "https://github.com/gzuidhof/myrtle", myrtle.ButtonStyle(myrtle.ButtonStyleGhost))
		builder.AddHeading("Danger · ghost", myrtle.HeadingLevel(3))
		builder.AddButton("Delete workspace", "https://github.com/gzuidhof/myrtle", myrtle.ButtonTone(myrtle.ButtonToneDanger), myrtle.ButtonStyle(myrtle.ButtonStyleGhost))
		builder.AddHeading("Primary · full width", myrtle.HeadingLevel(3))
		builder.AddButton("Open docs", "https://github.com/gzuidhof/myrtle", myrtle.ButtonTone(myrtle.ButtonTonePrimary), myrtle.ButtonFullWidth(true))
		builder.AddHeading("Secondary · full width", myrtle.HeadingLevel(3))
		builder.AddButton("Open docs", "https://github.com/gzuidhof/myrtle", myrtle.ButtonTone(myrtle.ButtonToneSecondary), myrtle.ButtonFullWidth(true))
		builder.AddHeading("Outline · small · no-wrap", myrtle.HeadingLevel(3))
		builder.AddButton("A longer CTA label", "https://github.com/gzuidhof/myrtle", myrtle.ButtonStyle(myrtle.ButtonStyleOutline), myrtle.ButtonSize(myrtle.ButtonSizeSmall), myrtle.ButtonNoWrap(true))
	case "button-group":
		builder.AddHeading("Centered", myrtle.HeadingLevel(3))
		builder.AddButtonGroup([]myrtle.ButtonGroupButton{{Label: "Approve", URL: "https://example.com/approve", Tone: myrtle.ButtonTonePrimary}, {Label: "Review", URL: "https://example.com/review", Tone: myrtle.ButtonToneSecondary}, {Label: "Delete", URL: "https://example.com/delete", Tone: myrtle.ButtonToneDanger}}, myrtle.ButtonGroupAlign(myrtle.ButtonAlignmentCenter))
		builder.AddHeading("Centered · joined", myrtle.HeadingLevel(3))
		builder.AddButtonGroup([]myrtle.ButtonGroupButton{{Label: "Approve", URL: "https://example.com/approve", Tone: myrtle.ButtonTonePrimary}, {Label: "Review", URL: "https://example.com/review", Tone: myrtle.ButtonToneSecondary}, {Label: "Delete", URL: "https://example.com/delete", Tone: myrtle.ButtonToneDanger}}, myrtle.ButtonGroupAlign(myrtle.ButtonAlignmentCenter), myrtle.ButtonGroupJoined(true))
		builder.AddHeading("End", myrtle.HeadingLevel(3))
		builder.AddButtonGroup([]myrtle.ButtonGroupButton{{Label: "Retry", URL: "https://example.com/retry", Tone: myrtle.ButtonTonePrimary}, {Label: "Details", URL: "https://example.com/details", Tone: myrtle.ButtonToneSecondary}, {Label: "Delete", URL: "https://example.com/delete", Tone: myrtle.ButtonToneDanger}}, myrtle.ButtonGroupAlign(myrtle.ButtonAlignmentEnd))
		builder.AddHeading("Full width on mobile", myrtle.HeadingLevel(3))
		builder.AddButtonGroup([]myrtle.ButtonGroupButton{{Label: "Approve", URL: "https://example.com/approve", Tone: myrtle.ButtonTonePrimary}, {Label: "Review", URL: "https://example.com/review", Tone: myrtle.ButtonToneSecondary}}, myrtle.ButtonGroupFullWidthOnMobile(true))
		builder.AddHeading("Stack on mobile · custom gap", myrtle.HeadingLevel(3))
		builder.AddButtonGroup([]myrtle.ButtonGroupButton{{Label: "Approve", URL: "https://example.com/approve", Tone: myrtle.ButtonTonePrimary}, {Label: "Review", URL: "https://example.com/review", Style: myrtle.ButtonStyleOutline}, {Label: "Delete", URL: "https://example.com/delete", Tone: myrtle.ButtonToneDanger, Style: myrtle.ButtonStyleGhost}}, myrtle.ButtonGroupGap(14), myrtle.ButtonGroupStackOnMobile(true))
	case "divider":
		builder.AddHeading("Default", myrtle.HeadingLevel(3))
		builder.AddDivider()
		builder.AddHeading("Dashed", myrtle.HeadingLevel(3))
		builder.AddDividerStyled(myrtle.DividerStyle(myrtle.DividerVariantDashed), myrtle.DividerThickness(2))
		builder.AddHeading("Dotted + inset", myrtle.HeadingLevel(3))
		builder.AddDividerStyled(myrtle.DividerStyle(myrtle.DividerVariantDotted), myrtle.DividerThickness(2), myrtle.DividerInset(32))
		builder.AddHeading("With label", myrtle.HeadingLevel(3))
		builder.AddDividerStyled(myrtle.DividerLabel("OR"))
		builder.AddHeading("Dashed + label", myrtle.HeadingLevel(3))
		builder.AddDividerStyled(myrtle.DividerStyle(myrtle.DividerVariantDashed), myrtle.DividerThickness(2), myrtle.DividerLabel("AND"), myrtle.DividerInset(16))
	case "image":
		builder.AddHeading("Default", myrtle.HeadingLevel(3))
		builder.AddImage("/assets/myrtle-placeholder.png", "Myrtle placeholder")
		builder.AddHeading("Full width", myrtle.HeadingLevel(3))
		builder.AddImage("/assets/myrtle-placeholder.png", "Myrtle placeholder", myrtle.ImageFullWidth())
		builder.AddHeading("Centered, 320px wide", myrtle.HeadingLevel(3))
		builder.AddImage("/assets/myrtle-placeholder.png", "Myrtle placeholder", myrtle.ImageWidth(320), myrtle.ImageAlign(myrtle.ImageAlignmentCenter))
	case "table":
		builder.AddHeading("Default", myrtle.HeadingLevel(3))
		builder.AddTable("Quarterly numbers", []string{"Metric", "Q1", "Q2"}, [][]string{{"Active Users", "8922", "10452"}, {"Conversion", "4.1%", "4.6%"}, {"Churn", "2.8%", "2.3%"}})
		builder.AddHeading("Compact only", myrtle.HeadingLevel(3))
		builder.AddTable("Quarterly numbers", []string{"Metric", "Q1", "Q2"}, [][]string{{"Active Users", "8922", "10452"}, {"Conversion", "4.1%", "4.6%"}, {"Churn", "2.8%", "2.3%"}}, myrtle.TableCompact(true))
		builder.AddHeading("Compact · zebra · end-aligned numeric", myrtle.HeadingLevel(3))
		builder.AddTable("Quarterly numbers", []string{"Metric", "Q1", "Q2"}, [][]string{{"Active Users", "8922", "10452"}, {"Conversion", "4.1%", "4.6%"}, {"Churn", "2.8%", "2.3%"}}, myrtle.TableZebraRows(true), myrtle.TableCompact(true), myrtle.TableRightAlignNumericColumns(true))
		builder.AddHeading("Relaxed + muted header", myrtle.HeadingLevel(3))
		builder.AddTable("Quarterly numbers", []string{"Metric", "Q1", "Q2"}, [][]string{{"Active Users", "8922", "10452"}, {"Conversion", "4.1%", "4.6%"}, {"Churn", "2.8%", "2.3%"}}, myrtle.TableDensity(myrtle.TableDensityRelaxed), myrtle.TableHeaderTone(myrtle.TableHeaderToneMuted))
		builder.AddHeading("Plain header + dashed borders", myrtle.HeadingLevel(3))
		builder.AddTable("Quarterly numbers", []string{"Metric", "Q1", "Q2"}, [][]string{{"Active Users", "8922", "10452"}, {"Conversion", "4.1%", "4.6%"}, {"Churn", "2.8%", "2.3%"}}, myrtle.TableHeaderTone(myrtle.TableHeaderTonePlain), myrtle.TableBorderStyle(myrtle.TableBorderStyleDashed))
	case "verification_code":
		builder.Add(myrtle.VerificationCodeBlock{Label: "Verification code", Value: "493817"})
	case "message":
		builder.AddMessage(myrtle.MessageBlock{
			SenderName:   "Alex Johnson",
			SenderHandle: "@alex",
			AvatarURL:    "/assets/avatar1.png",
			LogoAlt:      "Alex Johnson avatar",
			LogoHref:     "https://example.com/messages/42",
			Subject:      "New private message",
			Preview:      "Can you review the release notes before 3 PM?",
			SentAt:       "2m ago",
			Platform:     "Myrtle Chat",
			URL:          "https://example.com/messages/42",
			ActionLabel:  "Open thread",
			ActionURL:    "https://example.com/messages/42",
		})
	case "message-digest":
		builder.AddMessageDigest([]myrtle.MessageBlock{
			{SenderName: "Maya", SenderHandle: "@maya", AvatarURL: "/assets/avatar2.png", LogoAlt: "Maya avatar", LogoHref: "https://example.com/messages/43", Subject: "**Launch update**", Preview: "Can you check [the draft](https://example.com/draft)?", SentAt: "5m ago", Platform: "Myrtle Chat", URL: "https://example.com/messages/43"},
			{SenderName: "Nina", SenderHandle: "@nina", AvatarURL: "/assets/avatar3.png", LogoAlt: "Nina avatar", LogoHref: "https://example.com/messages/46", Preview: "Quick ping: can we move this to tomorrow morning?", SentAt: "11m ago", Platform: "Myrtle Chat", URL: "https://example.com/messages/46"},
			{SenderName: "Ben", SenderHandle: "@ben", AvatarURL: "/assets/avatar1.png", LogoAlt: "Ben avatar", LogoHref: "https://example.com/messages/44", Subject: "Design feedback", Preview: "Looks good overall.", SentAt: "20m ago", Platform: "Myrtle Chat", URL: "https://example.com/messages/44"},
			{SenderName: "Ari", SenderHandle: "@ari", Subject: "Follow-up", Preview: "Can we sync tomorrow?", SentAt: "1h ago", Platform: "Myrtle Chat", URL: "https://example.com/messages/45"},
		},
			myrtle.MessageDigestTitle("Inbox"),
			myrtle.MessageDigestSubtitle("Recent direct messages from Myrtle Chat"),
			myrtle.MessageDigestFooter("[Open inbox](https://example.com/messages)"),
		)
	case "tiles":
		builder.AddHeading("Default (3 columns)", myrtle.HeadingLevel(3))
		builder.AddTiles([]myrtle.TileEntry{{Content: "🚀", Title: "Launch", Subtitle: "Ready", URL: "https://example.com/launch", Variant: myrtle.TileVariantHighlight}, {Content: "12", Title: "Queued", Subtitle: "Jobs", Variant: myrtle.TileVariantWarning}, {Content: "✅", Title: "Healthy", Subtitle: "All systems", Variant: myrtle.TileVariantSuccess}})
		builder.AddHeading("Start aligned", myrtle.HeadingLevel(3))
		builder.AddTiles([]myrtle.TileEntry{{Content: "📦", Title: "Shipped", Subtitle: "2h ago"}, {Content: "9", Title: "Pending", Subtitle: "Orders"}, {Content: "⚠️", Title: "Delayed", Subtitle: "1 route", Variant: myrtle.TileVariantWarning}}, myrtle.TilesAlign(myrtle.TileAlignmentStart))
		builder.AddHeading("End aligned", myrtle.HeadingLevel(3))
		builder.AddTiles([]myrtle.TileEntry{{Content: "🧾", Title: "Invoices"}, {Content: "3", Title: "Overdue", Variant: myrtle.TileVariantCritical}, {Content: "✅", Title: "Paid", Subtitle: "This week"}}, myrtle.TilesAlign(myrtle.TileAlignmentEnd))
		builder.AddHeading("No content examples", myrtle.HeadingLevel(3))
		builder.AddTiles([]myrtle.TileEntry{{Title: "No icon tile", Subtitle: "Title + subtitle only"}, {Title: "Title only"}, {Subtitle: "Subtitle only"}}, myrtle.TilesBorder(true))
		builder.AddHeading("4 columns with border", myrtle.HeadingLevel(3))
		builder.AddTiles([]myrtle.TileEntry{{Content: "1", Title: "One"}, {Content: "2", Title: "Two"}, {Content: "3", Title: "Three", Variant: myrtle.TileVariantCritical}, {Content: "4", Title: "Four"}}, myrtle.TilesColumns(4), myrtle.TilesBorder(true))
		builder.AddHeading("Wraps when items exceed columns", myrtle.HeadingLevel(3))
		builder.AddTiles([]myrtle.TileEntry{{Content: "📥", Title: "Inbox"}, {Content: "📤", Title: "Outbox"}, {Content: "⚙️", Title: "Settings"}, {Content: "🔔", Title: "Alerts"}, {Content: "🧪", Title: "Experiments"}, {Content: "🗂️", Title: "Archive"}}, myrtle.TilesColumns(4), myrtle.TilesBorder(true))
	case "section":
		builder.AddHeading("Full header (category + title + subtitle)", myrtle.HeadingLevel(3))
		builder.AddSection(
			[]myrtle.Block{
				myrtle.TextBlock{Text: "Section body content with grouped context."},
				myrtle.ButtonBlock{Label: "Open section", URL: "https://example.com/section", Style: myrtle.ButtonStyleOutline},
			},
			myrtle.SectionCategory("Operational Summary"),
			myrtle.SectionTitle("Section block"),
			myrtle.SectionSubtitle("Optional subtitle for context"),
			myrtle.SectionPadding(18),
			myrtle.SectionBorder(true),
		)
		builder.AddHeading("No header fields", myrtle.HeadingLevel(3))
		builder.AddSection(
			[]myrtle.Block{
				myrtle.TextBlock{Text: "This section has no category, title, or subtitle."},
				myrtle.ButtonBlock{Label: "View details", URL: "https://example.com/section/no-header", Style: myrtle.ButtonStyleOutline},
			},
			myrtle.SectionPadding(18),
			myrtle.SectionBorder(true),
		)
		builder.AddHeading("Title only", myrtle.HeadingLevel(3))
		builder.AddSection(
			[]myrtle.Block{
				myrtle.TextBlock{Text: "This section has only a title."},
			},
			myrtle.SectionTitle("Title only"),
			myrtle.SectionPadding(18),
			myrtle.SectionBorder(true),
		)
		builder.AddHeading("Category only", myrtle.HeadingLevel(3))
		builder.AddSection(
			[]myrtle.Block{
				myrtle.TextBlock{Text: "This section has only a category value."},
			},
			myrtle.SectionCategory("Category only"),
			myrtle.SectionPadding(18),
			myrtle.SectionBorder(true),
		)
	case "grid":
		builder.AddGrid(
			[]myrtle.GridItem{
				{Blocks: []myrtle.Block{myrtle.HeadingBlock{Text: "Item 1", Level: 4}, myrtle.TextBlock{Text: "Grid content one."}}},
				{Blocks: []myrtle.Block{myrtle.HeadingBlock{Text: "Item 2", Level: 4}, myrtle.TextBlock{Text: "Grid content two."}}},
				{Blocks: []myrtle.Block{myrtle.HeadingBlock{Text: "Item 3", Level: 4}, myrtle.TextBlock{Text: "Grid content three."}}},
				{Blocks: []myrtle.Block{myrtle.HeadingBlock{Text: "Item 4", Level: 4}, myrtle.TextBlock{Text: "Grid content four."}}},
				{Blocks: []myrtle.Block{myrtle.HeadingBlock{Text: "Item 5", Level: 4}, myrtle.TextBlock{Text: "Wraps to new row."}}},
			},
			myrtle.GridColumns(3),
			myrtle.GridGap(12),
			myrtle.GridBorder(true),
		)
	case "card-list":
		builder.AddCardList(
			[]myrtle.CardItem{
				{Title: "Deploy complete", Subtitle: "Production", Body: "No customer impact detected.", URL: "https://example.com/deploy/1", CTALabel: "View"},
				{Title: "Billing updated", Subtitle: "Finance", Body: "Invoice #8241 has been paid.", URL: "https://example.com/billing/8241", CTALabel: "Open"},
				{Title: "Weekly report", Body: "Read the latest metrics and highlights.", URL: "https://example.com/reports/weekly"},
			},
			myrtle.CardListColumns(2),
			myrtle.CardListGap(12),
			myrtle.CardListBorder(true),
		)
	case "free-markdown":
		builder.AddFreeMarkdown("### Custom Markdown\n\nYou can use **bold**, lists, and links.\n\n- First\n- Second")
	default:
		return nil, errors.New("unknown block")
	}

	return builder.Build(), nil
}

func buildEmailItems(themeName string, selectedTheme theme.Theme) ([]pageItem, error) {
	items := make([]pageItem, 0, len(exampleEmails))
	for _, emailBuilder := range exampleEmails {
		email, err := emailBuilder.Build(selectedTheme)
		if err != nil {
			return nil, err
		}

		text, err := email.Text()
		if err != nil {
			return nil, err
		}

		items = append(items, pageItem{
			Key:     emailBuilder.Name,
			Name:    goEmailName(emailBuilder.Name),
			HTMLURL: "/emails/" + emailBuilder.Name + "/html?theme=" + themeName,
			Text:    text,
		})
	}

	return items, nil
}

func buildBlockItems(themeName string, selectedTheme theme.Theme) ([]groupedPageItems, error) {
	groups := make([]groupedPageItems, 0, len(blockGroups))
	for _, blockGroup := range blockGroups {
		items := make([]pageItem, 0, len(blockGroup.Items))
		for _, blockName := range blockGroup.Items {
			email, err := buildBlockEmail(blockName, selectedTheme)
			if err != nil {
				return nil, err
			}

			text, err := email.Text()
			if err != nil {
				return nil, err
			}

			items = append(items, pageItem{
				Key:     blockName,
				Name:    goBlockName(blockName),
				HTMLURL: "/blocks/" + blockName + "/html?theme=" + themeName,
				Text:    text,
			})
		}

		groups = append(groups, groupedPageItems{
			Name:  blockGroup.Name,
			Items: items,
		})
	}

	return groups, nil
}

func goEmailName(key string) string {
	return toSpacedTitle(key)
}

func goBlockName(key string) string {
	return toSpacedTitle(key)
}

func toSpacedTitle(value string) string {
	if value == "" {
		return value
	}

	parts := strings.FieldsFunc(value, func(r rune) bool {
		return r == '-' || r == '_'
	})
	for index, part := range parts {
		if part == "" {
			continue
		}

		parts[index] = strings.ToUpper(part[:1]) + part[1:]
	}

	return strings.Join(parts, " ")
}

func selectedThemeFromRequest(queryValue string) (string, theme.Theme) {
	switch queryValue {
	case "flat":
		return "flat", flat.New(flat.WithFallback(defaulttheme.New()))
	case "terminal":
		return "terminal", terminal.New(terminal.WithFallback(defaulttheme.New()))
	default:
		return "default", defaulttheme.New()
	}
}

func themeOptions(selected string) []themeOption {
	return []themeOption{
		{Name: "default", Selected: selected == "default"},
		{Name: "flat", Selected: selected == "flat"},
		{Name: "terminal", Selected: selected == "terminal"},
	}
}
