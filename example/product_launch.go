package example

import (
	"github.com/gzuidhof/myrtle"
	"github.com/gzuidhof/myrtle/theme"
	defaulttheme "github.com/gzuidhof/myrtle/theme/default"
)

func ProductLaunchEmail() (*myrtle.Email, error) {
	return ProductLaunchEmailWithTheme(defaulttheme.New())
}

func ProductLaunchEmailWithTheme(selectedTheme theme.Theme) (*myrtle.Email, error) {
	if selectedTheme == nil {
		selectedTheme = defaulttheme.New()
	}

	return myrtle.NewBuilder(
		selectedTheme,
		myrtle.WithStyles(theme.Styles{ColorPrimary: "#7c3aed"}),
		myrtle.WithHeaderOptions(
			myrtle.HeaderTitle("Product launch"),
			myrtle.HeaderProduct("Myrtle", "https://github.com/gzuidhof/myrtle"),
			myrtle.HeaderLogo("/assets/logo.png", "Myrtle"),
			myrtle.HeaderShowTextWithLogo(true),
		),
	).
		Preheader("Meet the new onboarding flows").
		Add(myrtle.HeroBlock{
			Eyebrow:  "New release",
			Title:    "Launch your emails faster",
			Body:     "We introduced reusable content blocks, stronger defaults, and cleaner theme fallback behavior.",
			CTALabel: "Read release notes",
			CTAURL:   "https://github.com/gzuidhof/myrtle",
			ImageURL: "/assets/hero.png",
			ImageAlt: "Myrtle launch hero image",
		}).
		AddCallout(myrtle.CalloutTypeInfo, "What changed", "Hero and summary blocks now make promotional emails easier to compose.", myrtle.CalloutStyle(myrtle.CalloutVariantSoft)).
		AddButton("Try the new examples", "https://github.com/gzuidhof/myrtle/tree/main/example", myrtle.ButtonTone(myrtle.ButtonTonePrimary)).
		AddFooterLinks([]myrtle.FooterLink{
			{Label: "Docs", URL: "https://github.com/gzuidhof/myrtle"},
			{Label: "Changelog", URL: "https://github.com/gzuidhof/myrtle"},
			{Label: "Support", URL: "https://github.com/gzuidhof/myrtle/discussions"},
		}, "You are receiving this update because your team uses Myrtle examples.").
		Build(), nil
}
