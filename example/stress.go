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
		Preheader("Stress/overflow test cases for legacy and modern clients").
		AddHeading("Stress/Overflow Test Cases", myrtle.HeadingLevel(1)).
		AddText("Long unbroken token:", "AVeryLongUnbrokenIdentifierThatShouldTriggerWordWrappingBehaviorChecksInLegacyEmailClients_ABCDEFGHIJKLMNOPQRSTUVWXYZ_0123456789").
		AddTable(
			"Overflow table rows",
			[]string{"Type", "Payload", "Status"},
			[][]string{{"id", "msg_0000000000000000000000000000000000000000000000000000000000001", "ok"}, {"url", "https://example.com/really/long/path/with/many/segments/that/legacy/engines/may/not/wrap/properly?token=abcdefghijklmnopqrstuvwxyz0123456789", "queued"}},
			myrtle.TableCompact(true),
			myrtle.TableRightAlignNumericColumns(true),
		).
		AddMessage(myrtle.MessageBlock{SenderName: "Legacy QA Bot", SenderHandle: "@legacy-qa", Subject: "Rendering edge case candidate", Preview: "Inspect this unbroken payload: ZXCVBNMASDFGHJKLQWERTYUIOP0123456789ZXCVBNMASDFGHJKLQWERTYUIOP0123456789", SentAt: "now", Platform: "Myrtle QA", URL: "https://example.com/messages/stress-case"}).
		Build(), nil
}
