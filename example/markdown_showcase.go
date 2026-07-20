package example

import (
	"github.com/gzuidhof/myrtle"
	"github.com/gzuidhof/myrtle/theme"
	defaulttheme "github.com/gzuidhof/myrtle/theme/default"
)

// MarkdownShowcaseEmail renders an email demonstrating inline markdown support
// (links, bold, italic, code) across every block that opts into it.
func MarkdownShowcaseEmail() (*myrtle.Email, error) {
	return MarkdownShowcaseEmailWithTheme(defaulttheme.New())
}

// MarkdownShowcaseEmailWithTheme renders the markdown showcase with the given theme.
func MarkdownShowcaseEmailWithTheme(selectedTheme theme.Theme) (*myrtle.Email, error) {
	if selectedTheme == nil {
		selectedTheme = defaulttheme.New()
	}

	const docsURL = "https://github.com/gzuidhof/myrtle"

	return myrtle.NewBuilder(selectedTheme).
		WithPreheader("Inline markdown — links, bold, italic, and code — across blocks").
		AddHeading("Inline markdown", myrtle.HeadingLevel(1)).
		AddText(
			"Every block below opts into inline markdown. Authored strings can carry "+
				"**bold**, *italic*, `code`, and [links]("+docsURL+") that render as real "+
				"anchors in HTML and degrade to \"label (url)\" in plaintext.",
			myrtle.TextMarkdown(),
		).

		// Callout
		AddHeading("Callout", myrtle.HeadingLevel(3)).
		AddCallout(
			myrtle.ToneInfo,
			"Heads up",
			"Rotate keys older than 90 days. See the [rotation guide]("+docsURL+") for the **exact** steps.",
			myrtle.CalloutMarkdown(),
		).

		// Summary card (body + footer)
		AddHeading("Summary card", myrtle.HeadingLevel(3)).
		AddSummaryCard(
			"Weekly report",
			"Delivery held at **99.8%** with no active incidents. Full breakdown in the [dashboard]("+docsURL+").",
			"Updated just now · [manage alerts]("+docsURL+")",
			myrtle.SummaryCardMarkdown(),
		).

		// List
		AddHeading("List", myrtle.HeadingLevel(3)).
		AddList(
			[]string{
				"Read the [quick-start](" + docsURL + "#quick-start)",
				"Compose with **reusable** blocks",
				"Render `HTML` and plaintext",
			},
			false,
			myrtle.ListMarkdown(),
		).

		// Hero
		AddHeading("Hero", myrtle.HeadingLevel(3)).
		AddHero(
			"Compose beautiful emails",
			"Start with the [quick-start guide]("+docsURL+") and ship your first **composable** email today.",
			"Open docs",
			docsURL,
			myrtle.HeroMarkdown(),
		).

		// Timeline
		AddHeading("Timeline", myrtle.HeadingLevel(3)).
		AddTimeline(
			"Incident recap",
			[]myrtle.TimelineItem{
				{Time: "09:07", Title: "Detected", Detail: "Queue latency spiked — see the [alert](" + docsURL + ")."},
				{Time: "09:18", Title: "Mitigated", Detail: "Scaled workers; **latency** returning to normal."},
				{Time: "09:42", Title: "Resolved", Detail: "Full [postmortem](" + docsURL + ") published."},
			},
			myrtle.TimelineMarkdown(),
			myrtle.TimelineCurrentIndex(2),
		).

		// Quote
		AddHeading("Quote", myrtle.HeadingLevel(3)).
		AddQuote(
			"Migrating to [composable blocks]("+docsURL+") cut our template maintenance time by *60%*.",
			"Email Engineering Team",
			myrtle.QuoteMarkdown(),
		).

		// Key/value
		AddHeading("Key / value", myrtle.HeadingLevel(3)).
		AddKeyValue(
			"Order summary",
			[]myrtle.KeyValuePair{
				{Key: "Plan", Value: "**Scale** Annual"},
				{Key: "Invoice", Value: "[#8241](" + docsURL + ")"},
				{Key: "Tracking", Value: "[1Z999AA10123456784](" + docsURL + ")"},
			},
			myrtle.KeyValueMarkdown(),
		).

		// Empty state
		AddHeading("Empty state", myrtle.HeadingLevel(3)).
		AddEmptyState(
			"No pending approvals",
			"Everything is up to date. You can [create a rule]("+docsURL+") to automate the next one.",
			"Create rule",
			docsURL,
			myrtle.EmptyStateMarkdown(),
		).
		AddDivider().
		AddText(
			"Markdown is strictly opt-in per block — omit the option and the same string renders literally.",
			myrtle.TextMarkdown(),
			myrtle.TextTone(myrtle.ToneMuted),
			myrtle.TextSize(myrtle.TextSizeSmall),
		).
		WithFooter(commonFooterGroup(), myrtle.FooterPlacement(myrtle.FooterPlacementOutside)).
		Build(), nil
}
