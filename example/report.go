package example

import (
	"github.com/gzuidhof/myrtle"
	"github.com/gzuidhof/myrtle/theme"
	defaulttheme "github.com/gzuidhof/myrtle/theme/default"
)

func WeeklyReportEmail() (*myrtle.Email, error) {
	return WeeklyReportEmailWithTheme(defaulttheme.New())
}

func WeeklyReportEmailWithTheme(selectedTheme theme.Theme) (*myrtle.Email, error) {
	if selectedTheme == nil {
		selectedTheme = defaulttheme.New()
	}

	rows := [][]string{
		{"Signups", "1,204", "+18%"},
		{"Activation", "42.6%", "+2.3%"},
	}

	return myrtle.NewBuilder(
		selectedTheme,
		myrtle.WithStyles(theme.Styles{ColorPrimary: "#059669", MaxWidthMain: "760px"}),
	).
		AddHeading("Weekly metrics report").
		WithPreheader("Your key numbers for this week").
		WithHeader(commonHeaderGroupWithAlt("Myrtle Analytics", "Analytics logo")).
		AddHeading("Weekly highlights", myrtle.HeadingLevel(2)).
		AddText("This example uses a wider container (max-width: 760px).", myrtle.TextTone(myrtle.TextToneMuted), myrtle.TextSize(myrtle.TextSizeSmall)).
		AddText("Here are the two KPIs we track most closely this week.").
		AddTable("Highlights", []string{"Metric", "Value", "Delta"}, rows).
		AddQuote("Activation improved after simplifying first-run setup.", "Growth Team").
		AddText("See the full dashboard for breakdowns and cohorts.").
		AddButton("Open dashboard", "https://example.com/dashboard").
		AddFreeMarkdown("_Tip:_ You can customize this email with additional blocks and custom themes.").
		Build(), nil
}
