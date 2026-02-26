package themetest

import (
	"strings"
	"testing"

	"github.com/gzuidhof/myrtle"
	"github.com/gzuidhof/myrtle/theme"
)

type ThemeFactory func() theme.Theme

func RunSuite(t *testing.T, factory ThemeFactory) {
	t.Helper()

	themeImpl := factory()
	if themeImpl == nil {
		t.Fatalf("theme factory returned nil")
	}
	if strings.TrimSpace(themeImpl.Name()) == "" {
		t.Fatalf("theme name should not be empty")
	}

	for _, block := range sampleBlocks() {
		block := block
		t.Run("block-"+string(block.Kind()), func(t *testing.T) {
			html, ok, err := themeImpl.RenderBlockHTML(theme.BlockView{
				Kind:   block.Kind(),
				Data:   block.TemplateData(),
				Values: sampleValues(),
			})
			if err != nil {
				t.Fatalf("RenderBlockHTML(%s) returned error: %v", block.Kind(), err)
			}
			if !ok {
				t.Fatalf("RenderBlockHTML(%s) was not handled", block.Kind())
			}
			if strings.TrimSpace(html) == "" {
				t.Fatalf("RenderBlockHTML(%s) returned empty html", block.Kind())
			}
		})
	}

	t.Run("email-centered-default", func(t *testing.T) {
		email := buildFullEmail(themeImpl, nil)

		html, err := email.HTML()
		if err != nil {
			t.Fatalf("HTML returned error: %v", err)
		}
		if strings.TrimSpace(html) == "" {
			t.Fatalf("HTML should not be empty")
		}

		text, err := email.Text()
		if err != nil {
			t.Fatalf("Text returned error: %v", err)
		}
		if strings.TrimSpace(text) == "" {
			t.Fatalf("Text should not be empty")
		}
	})

	t.Run("email-text-header", func(t *testing.T) {
		email := buildFullEmail(themeImpl, []myrtle.HeaderOption{
			myrtle.HeaderRenderInText(true),
		})

		html, err := email.HTML()
		if err != nil {
			t.Fatalf("HTML returned error: %v", err)
		}
		if strings.TrimSpace(html) == "" {
			t.Fatalf("HTML should not be empty")
		}
	})
}

func buildFullEmail(themeImpl theme.Theme, extraHeaderOptions []myrtle.HeaderOption) *myrtle.Email {
	headerBlock := myrtle.NewGroup().
		AddHeading("Theme suite", myrtle.HeadingLevel(1)).
		AddText("Myrtle")

	headerOptions := []myrtle.HeaderOption{}
	headerOptions = append(headerOptions, extraHeaderOptions...)

	builder := myrtle.NewBuilder(
		themeImpl,
		myrtle.WithStyles(sampleValues().Styles),
		myrtle.WithHeader(headerBlock, headerOptions...),
	)
	builder.WithPreheader("Comprehensive rendering check")

	for _, block := range sampleBlocks() {
		builder.Add(block)
	}

	return builder.Build()
}

func sampleValues() theme.Values {
	return theme.Values{
		Styles: theme.Styles{
			ColorPrimary:        "#2563eb",
			ColorText:           "#111827",
			ColorTextMuted:      "#6b7280",
			ColorBorder:         "#e5e7eb",
			ColorCodeBackground: "#f8fafc",
		},
	}
}

func sampleBlocks() []myrtle.Block {
	return []myrtle.Block{
		myrtle.TextBlock{Text: "Sample text"},
		myrtle.HeadingBlock{Text: "Sample heading", Level: 2},
		myrtle.SpacerBlock{Size: 12},
		myrtle.ListBlock{Items: []string{"One", "Two"}, Ordered: false},
		myrtle.KeyValueBlock{Header: "Details", Pairs: []myrtle.KeyValuePair{{Key: "A", Value: "1"}, {Key: "B", Value: "2"}}},
		myrtle.HorizontalBarChartBlock{Header: "Chart", Items: []myrtle.HorizontalBarChartItem{{Label: "US", Value: "60%", Percent: 60}, {Label: "EU", Value: "40%", Percent: 40}}},
		myrtle.SparklineBlock{Header: "Sparkline", Label: "Signups", Value: "1,204", Delta: "+8%", Points: []int{8, 12, 9, 14, 18, 16, 20}},
		myrtle.StackedBarBlock{Header: "Stacked", TotalLabel: "Total", TotalValue: "120k", Rows: []myrtle.StackedBarRow{{Label: "Channels", Segments: []myrtle.StackedBarSegment{{Label: "Email", Percent: 60, Value: "60%"}, {Label: "SMS", Percent: 25, Value: "25%"}, {Label: "Push", Percent: 15, Value: "15%"}}}}},
		myrtle.ProgressBlock{Header: "Progress", Items: []myrtle.ProgressItem{{Label: "Onboarding", Percent: 72}, {Label: "Verification", Percent: 45}}},
		myrtle.DistributionBlock{Header: "Distribution", Buckets: []myrtle.DistributionBucket{{Label: "0-10", Count: 12}, {Label: "11-20", Count: 34}, {Label: "21-30", Count: 18}}},
		myrtle.TimelineBlock{Header: "Timeline", Items: []myrtle.TimelineItem{{Time: "09:00", Title: "Start", Detail: "Initialization"}, {Time: "09:30", Title: "Done", Detail: "Completed"}}},
		myrtle.StatsRowBlock{Header: "KPIs", Stats: []myrtle.StatItem{{Label: "Delivery", Value: "99.8%", Delta: "+0.2%", DeltaSemantic: myrtle.StatDeltaSemanticPositive}, {Label: "Bounces", Value: "0.9%", Delta: "-0.1%", DeltaSemantic: myrtle.StatDeltaSemanticNegative}}},
		myrtle.BadgeBlock{Tone: myrtle.BadgeToneInfo, Text: "Info"},
		myrtle.SummaryCardBlock{Title: "Summary", Body: "All systems operational", Footer: "Updated now"},
		myrtle.AttachmentBlock{Filename: "report.pdf", Meta: "PDF · 123 KB", URL: "https://example.com/report.pdf", CTA: "Download"},
		myrtle.HeroBlock{Eyebrow: "New", Title: "Faster sends", Body: "Compose emails quickly", CTALabel: "Open docs", CTAURL: "https://example.com/docs", ImageURL: "https://example.com/hero.png", ImageAlt: "Hero"},
		myrtle.FooterLinksBlock{Links: []myrtle.FooterLink{{Label: "Help", URL: "https://example.com/help"}, {Label: "Privacy", URL: "https://example.com/privacy"}}, Note: "You are receiving this email because you are subscribed."},
		myrtle.PriceSummaryBlock{Header: "Order summary", Items: []myrtle.PriceLine{{Label: "Subtotal", Value: "$89.00"}, {Label: "Tax", Value: "$7.12"}}, TotalLabel: "Total", TotalValue: "$96.12"},
		myrtle.EmptyStateBlock{Title: "No activity", Body: "Everything is up to date.", ActionLabel: "Open dashboard", ActionURL: "https://example.com/dashboard"},
		myrtle.QuoteBlock{Text: "Great product", Author: "User"},
		myrtle.CalloutBlock{Type: myrtle.CalloutTypeWarning, Variant: myrtle.CalloutVariantSolid, Title: "Attention", Body: "Action required"},
		myrtle.MessageBlock{SenderName: "Alex Johnson", SenderHandle: "@alex", AvatarURL: "https://i.pravatar.cc/80?img=12", LogoAlt: "Alex Johnson avatar", LogoHref: "https://example.com/messages/42", Subject: "New private message", Preview: "Can you review the release notes before 3 PM?", SentAt: "2m ago", Platform: "Myrtle Chat", URL: "https://example.com/messages/42", ActionLabel: "Open thread", ActionURL: "https://example.com/messages/42"},
		myrtle.MessageDigestBlock{Title: "Inbox", Subtitle: "Recent direct messages", Footer: "[Open inbox](https://example.com/messages)", Messages: []myrtle.MessageBlock{{SenderName: "Maya", SenderHandle: "@maya", AvatarURL: "https://i.pravatar.cc/80?img=5", LogoAlt: "Maya avatar", LogoHref: "https://example.com/messages/43", Subject: "**Launch update**", Preview: "Can you check [the draft](https://example.com/draft)?", SentAt: "5m ago", Platform: "Myrtle Chat", URL: "https://example.com/messages/43"}, {SenderName: "Nina", SenderHandle: "@nina", Preview: "Quick ping: can we move this to tomorrow morning?", SentAt: "11m ago", Platform: "Myrtle Chat", URL: "https://example.com/messages/46"}, {SenderName: "Ben", SenderHandle: "@ben", Subject: "Design feedback", Preview: "Looks good overall.", SentAt: "20m ago", Platform: "Myrtle Chat", URL: "https://example.com/messages/44"}}},
		myrtle.LegalBlock{CompanyName: "Myrtle Inc.", Address: "123 Market St", ManageURL: "https://example.com/manage", UnsubscribeURL: "https://example.com/unsub"},
		myrtle.ColumnsBlock{
			Left:       []myrtle.Block{myrtle.TextBlock{Text: "Left column"}},
			Right:      []myrtle.Block{myrtle.TextBlock{Text: "Right column"}},
			LeftWidth:  50,
			RightWidth: 50,
		},
		myrtle.ButtonBlock{Label: "Open", URL: "https://example.com/open", Tone: myrtle.ButtonToneSecondary, Alignment: myrtle.ButtonAlignmentCenter, FullWidth: true},
		myrtle.ButtonGroupBlock{Buttons: []myrtle.ButtonGroupButton{{Label: "Approve", URL: "https://example.com/approve", Tone: myrtle.ButtonTonePrimary}, {Label: "Review", URL: "https://example.com/review", Tone: myrtle.ButtonToneSecondary}, {Label: "Later", URL: "https://example.com/later", Style: myrtle.ButtonStyleGhost}}, Alignment: myrtle.ButtonAlignmentCenter},
		myrtle.DividerBlock{},
		myrtle.ImageBlock{Src: "https://example.com/image.png", Alt: "Preview"},
		myrtle.TableBlock{Header: "Table", Columns: []string{"A", "B"}, Rows: [][]string{{"1", "2"}, {"3", "4"}}, ZebraRows: true, Compact: true, RightAlignNumericColumns: true},
		myrtle.VerificationCodeBlock{Label: "Verification code", Value: "123456"},
		myrtle.TilesBlock{Columns: 3, Border: true, Entries: []myrtle.TileEntry{{Content: "🚀", Title: "Launch", Subtitle: "Active", Variant: myrtle.TileVariantHighlight}, {Content: "42", Title: "Queued", Subtitle: "Pending", Variant: myrtle.TileVariantWarning}, {Content: "✅", Title: "Done", Variant: myrtle.TileVariantSuccess}}},
		myrtle.FreeMarkdownBlock{Markdown: "### Custom markdown\n\n- one\n- two"},
	}
}
