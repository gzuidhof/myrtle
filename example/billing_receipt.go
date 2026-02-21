package example

import (
	"github.com/gzuidhof/myrtle"
	"github.com/gzuidhof/myrtle/theme"
	defaulttheme "github.com/gzuidhof/myrtle/theme/default"
)

func BillingReceiptEmail() (*myrtle.Email, error) {
	return BillingReceiptEmailWithTheme(defaulttheme.New())
}

func BillingReceiptEmailWithTheme(selectedTheme theme.Theme) (*myrtle.Email, error) {
	if selectedTheme == nil {
		selectedTheme = defaulttheme.New()
	}

	return myrtle.NewBuilder(
		selectedTheme,
	).
		AddHeading("Your payment receipt").
		Preheader("Invoice INV-2026-021 has been paid").
		Product("Myrtle Billing", "https://example.com/billing").
		Logo("/assets/logo.png", "Myrtle Billing").
		AddHeading("Payment confirmed", myrtle.HeadingLevel(2)).
		AddKeyValue("Receipt details", []myrtle.KeyValuePair{
			{Key: "Invoice", Value: "INV-2026-021"},
			{Key: "Date", Value: "2026-02-21"},
			{Key: "Amount", Value: "$249.00"},
			{Key: "Method", Value: "Visa •••• 4242"},
		}).
		AddAction("Download your PDF receipt:", "View receipt", "https://example.com/billing/receipt/INV-2026-021").
		AddLegal("Myrtle Inc.", "123 Market St, San Francisco, CA", "https://example.com/preferences", "https://example.com/unsubscribe").
		Build(), nil
}
