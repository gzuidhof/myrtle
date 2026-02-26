package example

import (
	"github.com/gzuidhof/myrtle"
	"github.com/gzuidhof/myrtle/theme"
	defaulttheme "github.com/gzuidhof/myrtle/theme/default"
)

func InvoiceSummaryEmail() (*myrtle.Email, error) {
	return InvoiceSummaryEmailWithTheme(defaulttheme.New())
}

func InvoiceSummaryEmailWithTheme(selectedTheme theme.Theme) (*myrtle.Email, error) {
	if selectedTheme == nil {
		selectedTheme = defaulttheme.New()
	}

	return myrtle.NewBuilder(
		selectedTheme,
		myrtle.WithStyles(theme.Styles{ColorPrimary: "#0ea5e9"}),
	).
		WithPreheader("Line items and totals for invoice #INV-0242").
		WithHeader(commonHeaderGroup("Myrtle Billing")).
		AddPriceSummary("Invoice #INV-0242", []myrtle.PriceLine{
			{Label: "Platform", Value: "$79.00"},
			{Label: "Email volume", Value: "$18.40"},
			{Label: "Discount", Value: "-$5.00"},
			{Label: "Tax", Value: "$7.79"},
		}, "Total due", "$100.19").
		AddTable(
			"Usage details",
			[]string{"Item", "Quantity", "Unit", "Total"},
			[][]string{
				{"Transactional sends", "92000", "$0.0002", "$18.40"},
				{"Seats", "5", "$15.80", "$79.00"},
				{"Subtotal", "", "", "$97.40"},
				{"Tax", "", "", "$7.79"},
			},
			myrtle.TableZebraRows(true),
			myrtle.TableCompact(true),
			myrtle.TableRightAlignNumericColumns(true),
		).
		AddButton("Download invoice", "https://example.com/billing/invoices/inv-0242", myrtle.ButtonTone(myrtle.ButtonToneSecondary), myrtle.ButtonFullWidth(true)).
		AddFooterLinks(
			[]myrtle.FooterLink{{Label: "Billing portal", URL: "https://example.com/billing"}, {Label: "Contact support", URL: "https://example.com/support"}},
			"Questions? Reach out to billing support anytime.",
		).
		Build(), nil
}
