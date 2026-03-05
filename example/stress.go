package example

import (
	"github.com/gzuidhof/myrtle"
	"github.com/gzuidhof/myrtle/theme"
	defaulttheme "github.com/gzuidhof/myrtle/theme/default"
)

// StressTestEmailWithTheme returns an email with only the stress/overflow test blocks.
func StressTestEmailWithTheme(selectedTheme theme.Theme) (*myrtle.Email, error) {
	if selectedTheme == nil {
		selectedTheme = defaulttheme.New()
	}
	return myrtle.NewBuilder(selectedTheme).
		WithoutHeader().
		WithPreheader("Rendering QA sample").
		AddHeading("Client rendering QA", myrtle.HeadingLevel(1)).
		AddText("This internal sample verifies long-token and overflow behavior in legacy clients.").
		AddText("AVeryLongUnbrokenIdentifierThatShouldTriggerWordWrappingBehaviorChecksInLegacyEmailClients_ABCDEFGHIJKLMNOPQRSTUVWXYZ_0123456789", myrtle.TextSize(myrtle.TextSizeSmall)).
		AddHeading("Payload rows", myrtle.HeadingLevel(3)).
		AddTable(
			[]string{"Type", "Payload", "Status"},
			[][]string{{"id", "msg_0000000000000000000000000000000000000000000000000000000000001", "ok"}, {"url", "https://example.com/really/long/path/with/many/segments/that/legacy/engines/may/not/wrap/properly?token=abcdefghijklmnopqrstuvwxyz0123456789", "queued"}},
			myrtle.TableCompact(true),
			myrtle.TableRightAlignNumericColumns(true),
		).
		AddMessage(myrtle.MessageBlock{SenderName: "Legacy QA Bot", SenderHandle: "@legacy-qa", Subject: "Rendering edge case candidate", Preview: "Inspect this unbroken payload: ZXCVBNMASDFGHJKLQWERTYUIOP0123456789ZXCVBNMASDFGHJKLQWERTYUIOP0123456789", SentAt: "now", Platform: "Myrtle QA", URL: "https://example.com/messages/stress-case"}).
		Build(), nil
}
