package server

import (
	"github.com/gzuidhof/myrtle"
	"github.com/gzuidhof/myrtle/example/server/customblock"
	"github.com/gzuidhof/myrtle/theme"
	defaulttheme "github.com/gzuidhof/myrtle/theme/default"
)

func CustomFeatureFlagRolloutEmailWithTheme(selectedTheme theme.Theme) (*myrtle.Email, error) {
	if selectedTheme == nil {
		selectedTheme = defaulttheme.New()
	}

	builder := myrtle.NewBuilder(selectedTheme)

	email := builder.
		WithPreheader("Custom block demo: feature flag rollout guardrail snapshot").
		WithHeader(
			myrtle.NewGroup().
				Add(myrtle.ImageBlock{Src: "/assets/logo.png", Alt: "Myrtle", Width: 140, Align: myrtle.ImageAlignmentCenter}).
				Add(myrtle.TextBlock{Text: "Myrtle Control Plane", Align: myrtle.TextAlignCenter, Weight: myrtle.TextWeightSemibold}),
		).
		AddHeading("Feature rollout guardrail", myrtle.HeadingLevel(1)).
		AddText("This email demonstrates a registry-backed custom block rendered without modifying any built-in themes.").
		Add(customblock.NewFeatureFlagRolloutBlock(customblock.FeatureFlagRollout{
			FlagName:         "checkout.v2",
			Environment:      "production",
			RolloutPercent:   35,
			ErrorBudgetUsed:  "14%",
			P95LatencyDelta:  "+23ms",
			AutoRollback:     "enabled at +30ms",
			Status:           "at-risk",
			Owner:            "growth-platform",
			ChangeID:         "chg_2026_02_23_1842",
			OpenFlagURL:      "https://example.com/flags/checkout.v2",
			RollbackNowURL:   "https://example.com/flags/checkout.v2/rollback",
			IncidentBoardURL: "https://example.com/incidents/active",
		})).
		AddLegal("Myrtle Inc.", "123 Market St, San Francisco, CA", "https://example.com/preferences", "https://example.com/unsubscribe").
		Build()

	return email, nil
}
