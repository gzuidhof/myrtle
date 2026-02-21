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
		myrtle.WithStyles(theme.Styles{PrimaryColor: "#059669"}),
	).
		AddHeading("Weekly metrics report").
		Preheader("Your key numbers for this week").
		Product("Myrtle Analytics", "https://example.com/analytics").
		Logo("/assets/logo.png", "Analytics logo").
		AddHeading("Weekly highlights", myrtle.HeadingLevel(2)).
		AddText("Here are the two KPIs we track most closely this week.").
		AddTable("Highlights", []string{"Metric", "Value", "Delta"}, rows).
		AddQuote("Activation improved after simplifying first-run setup.", "Growth Team").
		AddAction("See the full dashboard for breakdowns and cohorts.", "Open dashboard", "https://example.com/dashboard").
		AddFreeMarkdown("_Tip:_ You can customize this email with additional blocks and custom themes.").
		Build(), nil
}
