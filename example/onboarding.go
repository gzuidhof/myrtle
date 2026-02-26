package example

import (
	"github.com/gzuidhof/myrtle"
	"github.com/gzuidhof/myrtle/theme"
	defaulttheme "github.com/gzuidhof/myrtle/theme/default"
)

func OnboardingChecklistEmail() (*myrtle.Email, error) {
	return OnboardingChecklistEmailWithTheme(defaulttheme.New())
}

func OnboardingChecklistEmailWithTheme(selectedTheme theme.Theme) (*myrtle.Email, error) {
	if selectedTheme == nil {
		selectedTheme = defaulttheme.New()
	}

	return myrtle.NewBuilder(
		selectedTheme,
	).
		WithPreheader("Complete your onboarding checklist").
		AddHeading("Welcome!", myrtle.HeadingLevel(1)).
		AddText("Your workspace is ready. Complete these steps to get value quickly:").
		AddList([]string{"Invite your team", "Connect your domain", "Send your first transactional email"}, true).
		AddCallout(myrtle.CalloutTypeInfo, "Need help?", "The docs include copy-paste examples for all common flows.").
		AddText("Open the quickstart and complete setup:").
		AddButton("Start onboarding", "https://github.com/gzuidhof/myrtle").
		Build(), nil
}
