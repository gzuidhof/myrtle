package example

import (
	"github.com/gzuidhof/myrtle"
	"github.com/gzuidhof/myrtle/theme"
	defaulttheme "github.com/gzuidhof/myrtle/theme/default"
)

func WeeklyReportEmail() (*myrtle.Email, error) {
	rows := [][]string{
		{"Signups", "1,204", "+18%"},
		{"Active Users", "8,922", "+6%"},
		{"MRR", "$42,310", "+4.1%"},
	}

	return myrtle.NewBuilder(
		defaulttheme.New(),
		myrtle.WithStyles(theme.Styles{PrimaryColor: "#059669"}),
	).
		Preheader("Your key numbers for this week").
		Product("Myrtle Analytics", "https://example.com/analytics").
		Logo("https://example.com/analytics-logo.png", "Analytics logo").
		AddText("Here is your weekly snapshot:").
		AddTable("Highlights", []string{"Metric", "Value", "Delta"}, rows).
		AddAction("See the full dashboard for breakdowns and cohorts.", "Open dashboard", "https://example.com/dashboard").
		AddFreeMarkdown("_Tip:_ You can customize this email with additional blocks and custom themes.").
		Build(), nil
}
