package example

import (
	"github.com/gzuidhof/myrtle"
	"github.com/gzuidhof/myrtle/theme"
	defaulttheme "github.com/gzuidhof/myrtle/theme/default"
)

func ColumnsComplexEmail() (*myrtle.Email, error) {
	return ColumnsComplexEmailWithTheme(defaulttheme.New())
}

func ColumnsComplexEmailWithTheme(selectedTheme theme.Theme) (*myrtle.Email, error) {
	if selectedTheme == nil {
		selectedTheme = defaulttheme.New()
	}

	return myrtle.NewBuilder(
		selectedTheme,
	).
		WithoutHeader().
		Preheader("A multi-column layout with actionable sections").
		WithHeader(
			myrtle.HeaderTitle("Product updates and account insights"),
			myrtle.HeaderProduct("Myrtle Weekly", "https://example.com/weekly"),
			myrtle.HeaderLogo("/assets/logo.png", "Myrtle Weekly"),
		).
		AddHeading("Your weekly digest", myrtle.HeadingLevel(1)).
		AddText("A concise operational digest with one metrics column and one actions column.").
		AddColumns(
			func(column *myrtle.ColumnBuilder) {
				column.AddHeading("Team activity", myrtle.HeadingLevel(3)).
					AddKeyValue("Highlights", []myrtle.KeyValuePair{{Key: "New users", Value: "184"}, {Key: "Delivery", Value: "99.8%"}}).
					AddCallout(myrtle.CalloutTypeSuccess, "Great performance", "Delivery rate improved by 0.4% week over week.")
			},
			func(column *myrtle.ColumnBuilder) {
				column.AddHeading("Next actions", myrtle.HeadingLevel(3)).
					AddList([]string{"Review failed webhooks", "Confirm SLA report"}, false).
					AddText("Open the operations dashboard for full context.").
					AddButton("Open dashboard", "https://example.com/ops")
			},
			myrtle.ColumnsWidths(60, 40),
		).
		AddSpacer(myrtle.SpacerSize(10)).
		AddColumns(
			func(column *myrtle.ColumnBuilder) {
				column.AddQuote("This structure made our transactional templates cleaner and easier to maintain.", "Platform Team")
			},
			func(column *myrtle.ColumnBuilder) {
				column.AddCallout(myrtle.CalloutTypeInfo, "Next step", "Audit one of your old templates and migrate it to block composition.")
			},
			myrtle.ColumnsWidths(50, 50),
		).
		AddDivider().
		AddLegal("Myrtle Inc.", "123 Market St, San Francisco, CA", "https://example.com/preferences", "https://example.com/unsubscribe").
		Build(), nil
}
