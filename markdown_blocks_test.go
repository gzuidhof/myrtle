package myrtle_test

import (
	"strings"
	"testing"

	"github.com/gzuidhof/myrtle"
	defaulttheme "github.com/gzuidhof/myrtle/theme/default"
)

const mdURL = "https://example.com/docs"

// buildOne renders a single-block email to HTML and plaintext.
func buildOne(t *testing.T, build func(*myrtle.Builder) *myrtle.Builder) (string, string) {
	t.Helper()
	email := build(myrtle.NewBuilder(defaulttheme.New())).Build()

	html, err := email.HTML()
	if err != nil {
		t.Fatalf("HTML error: %v", err)
	}
	text, err := email.Text()
	if err != nil {
		t.Fatalf("Text error: %v", err)
	}
	return html, text
}

// assertInlineLink checks that an inline markdown link rendered to a real HTML
// anchor and degraded to "label (url)" in plaintext, with no raw markdown left.
func assertInlineLink(t *testing.T, html, text, label string) {
	t.Helper()
	if !strings.Contains(html, `href="`+mdURL+`"`) {
		t.Errorf("HTML missing anchor href %q:\n%s", mdURL, html)
	}
	if !strings.Contains(html, ">"+label+"</a>") {
		t.Errorf("HTML missing anchor label %q:\n%s", label, html)
	}
	if strings.Contains(html, "["+label+"]") {
		t.Errorf("raw markdown link leaked into HTML:\n%s", html)
	}
	if !strings.Contains(text, label+" ("+mdURL+")") {
		t.Errorf("plaintext did not degrade to %q:\n%s", label+" ("+mdURL+")", text)
	}
	if strings.Contains(text, "["+label+"]") {
		t.Errorf("raw markdown link leaked into plaintext:\n%s", text)
	}
}

func TestBlockMarkdownInlineLinks(t *testing.T) {
	t.Parallel()

	link := "[Docs](" + mdURL + ")"

	cases := []struct {
		name  string
		label string
		build func(*myrtle.Builder) *myrtle.Builder
	}{
		{"callout", "Docs", func(b *myrtle.Builder) *myrtle.Builder {
			return b.AddCallout(myrtle.ToneInfo, "Heads up", "Read the "+link+" first.", myrtle.CalloutMarkdown())
		}},
		{"summary-card-body", "Docs", func(b *myrtle.Builder) *myrtle.Builder {
			return b.AddSummaryCard("Report", "See the "+link+".", "", myrtle.SummaryCardMarkdown())
		}},
		{"summary-card-footer", "Docs", func(b *myrtle.Builder) *myrtle.Builder {
			return b.AddSummaryCard("Report", "Body.", "More in the "+link+".", myrtle.SummaryCardMarkdown())
		}},
		{"empty-state", "import", func(b *myrtle.Builder) *myrtle.Builder {
			return b.AddEmptyState("Nothing here", "Try an [import]("+mdURL+").", "", "", myrtle.EmptyStateMarkdown())
		}},
		{"hero", "Docs", func(b *myrtle.Builder) *myrtle.Builder {
			return b.AddHero("Title", "Start with the "+link+".", "", "", myrtle.HeroMarkdown())
		}},
		{"timeline", "logs", func(b *myrtle.Builder) *myrtle.Builder {
			return b.AddTimeline("History", []myrtle.TimelineItem{
				{Time: "09:00", Title: "Deploy", Detail: "Check the [logs](" + mdURL + ")."},
			}, myrtle.TimelineMarkdown())
		}},
		{"quote", "product", func(b *myrtle.Builder) *myrtle.Builder {
			return b.AddQuote("I love the [product]("+mdURL+").", "Sam", myrtle.QuoteMarkdown())
		}},
		{"key-value", "1Z999", func(b *myrtle.Builder) *myrtle.Builder {
			return b.AddKeyValue("Order", []myrtle.KeyValuePair{
				{Key: "Tracking", Value: "[1Z999](" + mdURL + ")"},
			}, myrtle.KeyValueMarkdown())
		}},
		{"list", "Docs", func(b *myrtle.Builder) *myrtle.Builder {
			return b.AddList([]string{"Read the " + link, "Then relax"}, false, myrtle.ListMarkdown())
		}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			html, text := buildOne(t, tc.build)
			assertInlineLink(t, html, text, tc.label)
		})
	}
}

func TestBlockMarkdownEmphasis(t *testing.T) {
	t.Parallel()

	// Hero's plaintext adds no decorative markers, so a global "no ** / no `"
	// assertion cleanly proves the body's emphasis was unwrapped.
	html, text := buildOne(t, func(b *myrtle.Builder) *myrtle.Builder {
		return b.AddHero("Title", "This is **very** important and `precise`.", "", "", myrtle.HeroMarkdown())
	})

	if !strings.Contains(html, "<strong>very</strong>") {
		t.Errorf("HTML missing bold rendering:\n%s", html)
	}
	if !strings.Contains(html, "<code>precise</code>") {
		t.Errorf("HTML missing code rendering:\n%s", html)
	}
	if strings.Contains(text, "**") || strings.Contains(text, "`") {
		t.Errorf("emphasis markers leaked into plaintext:\n%s", text)
	}
	if !strings.Contains(text, "very important and precise") {
		t.Errorf("plaintext did not unwrap emphasis:\n%s", text)
	}
}

// TestBlockMarkdownIsOptIn confirms that without the markdown option, inline
// markdown is left literal in both HTML and plaintext.
func TestBlockMarkdownIsOptIn(t *testing.T) {
	t.Parallel()

	html, text := buildOne(t, func(b *myrtle.Builder) *myrtle.Builder {
		return b.AddCallout(myrtle.ToneInfo, "Heads up", "Read the [Docs]("+mdURL+") first.")
	})

	if !strings.Contains(html, "[Docs](") {
		t.Errorf("expected literal markdown in HTML when opt-in is off:\n%s", html)
	}
	if !strings.Contains(text, "[Docs](") {
		t.Errorf("expected literal markdown in plaintext when opt-in is off:\n%s", text)
	}
}
