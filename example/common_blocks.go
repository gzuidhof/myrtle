package example

import (
	"github.com/gzuidhof/myrtle"
	"github.com/gzuidhof/myrtle/theme"
	defaulttheme "github.com/gzuidhof/myrtle/theme/default"
)

func CommonBlocksEmail() (*myrtle.Email, error) {
	return CommonBlocksEmailWithTheme(defaulttheme.New())
}

func CommonBlocksEmailWithTheme(selectedTheme theme.Theme) (*myrtle.Email, error) {
	if selectedTheme == nil {
		selectedTheme = defaulttheme.New()
	}

	return myrtle.NewBuilder(
		selectedTheme,
		myrtle.WithStyles(theme.Styles{ColorPrimary: "#2563eb"}),
	).
		AddHeading("Common block showcase").
		Preheader("Preview all common content blocks in one email").
		Product("Myrtle", "https://github.com/gzuidhof/myrtle").
		Logo("/assets/logo.png", "Myrtle").
		AddHeading("Common blocks", myrtle.HeadingLevel(1)).
		AddText("A compact order-status update using common block patterns.").
		AddSpacer(myrtle.SpacerSize(12)).
		AddList([]string{"Track your order", "Update delivery preferences", "Contact support"}, false).
		AddColumns(
			func(column *myrtle.ColumnBuilder) {
				column.AddHeading("Column A", myrtle.HeadingLevel(3)).
					AddText("Useful summary text for the first column.")
			},
			func(column *myrtle.ColumnBuilder) {
				column.AddHeading("Column B", myrtle.HeadingLevel(3)).
					AddList([]string{"Compact info", "Secondary details"}, false)
			},
			myrtle.ColumnsWidths(60, 40),
		).
		AddKeyValue("Order Summary", []myrtle.KeyValuePair{{Key: "Order", Value: "#94721"}, {Key: "Status", Value: "Shipped"}}).
		AddQuote("Keeping updates short improves completion rates.", "CX Team").
		AddCallout(myrtle.CalloutTypeSuccess, "On schedule", "Your package is expected to arrive tomorrow.").
		AddLegal("Myrtle Inc.", "123 Market St, San Francisco, CA", "https://example.com/preferences", "https://example.com/unsubscribe").
		Build(), nil
}
