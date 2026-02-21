package myrtle

import (
	"strings"
	"testing"
)

func TestNormalizedIntPoints(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		output []int
	}{
		{
			name:   "empty",
			input:  []int{},
			output: []int{},
		},
		{
			name:   "clamps negatives",
			input:  []int{-10, -1, 0, 5},
			output: []int{0, 0, 0, 5},
		},
		{
			name:   "keeps positives",
			input:  []int{1, 2, 3},
			output: []int{1, 2, 3},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := normalizedIntPoints(test.input)
			if len(actual) != len(test.output) {
				t.Fatalf("expected len %d, got %d", len(test.output), len(actual))
			}
			for i := range actual {
				if actual[i] != test.output[i] {
					t.Fatalf("at index %d expected %d, got %d", i, test.output[i], actual[i])
				}
			}
		})
	}
}

func TestSparklineGlyphs(t *testing.T) {
	tests := []struct {
		name           string
		points         []int
		expectedLength int
		expectedRunes  string
	}{
		{
			name:           "empty",
			points:         []int{},
			expectedLength: 0,
			expectedRunes:  "",
		},
		{
			name:           "single point",
			points:         []int{7},
			expectedLength: 1,
			expectedRunes:  "▅",
		},
		{
			name:           "flat line",
			points:         []int{5, 5, 5, 5},
			expectedLength: 4,
			expectedRunes:  "▅",
		},
		{
			name:           "increasing range",
			points:         []int{0, 1, 2, 3, 4, 5, 6, 7},
			expectedLength: 8,
			expectedRunes:  "▁▂▃▄▅▆▇█",
		},
		{
			name:           "large range",
			points:         []int{0, 100, 1000, 10000},
			expectedLength: 4,
			expectedRunes:  "▁",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := sparklineGlyphs(test.points)
			if len([]rune(actual)) != test.expectedLength {
				t.Fatalf("expected glyph length %d, got %d (%q)", test.expectedLength, len([]rune(actual)), actual)
			}

			if test.expectedRunes != "" {
				for _, glyph := range []rune(test.expectedRunes) {
					if !strings.ContainsRune(actual, glyph) {
						t.Fatalf("expected %q to contain glyph %q", actual, glyph)
					}
				}
			}
		})
	}
}

func TestSparklineBlockTemplateDataAndMarkdown(t *testing.T) {
	block := SparklineBlock{
		Header:        " Trend ",
		Label:         "Users",
		Value:         "1,240",
		Delta:         "+12%",
		DeltaSemantic: "unexpected",
		Points:        []int{-2, 0, 4, 9},
	}

	normalized := block.TemplateData().(SparklineBlock)
	if len(normalized.Points) != 4 {
		t.Fatalf("expected 4 points, got %d", len(normalized.Points))
	}
	if normalized.Points[0] != 0 {
		t.Fatalf("expected negative values to clamp to 0, got %d", normalized.Points[0])
	}
	if normalized.DeltaSemantic != StatDeltaSemanticNone {
		t.Fatalf("expected unknown sparkline delta semantic to normalize to none, got %q", normalized.DeltaSemantic)
	}

	markdown, err := block.RenderMarkdown(RenderContext{})
	if err != nil {
		t.Fatalf("RenderMarkdown returned error: %v", err)
	}
	if !strings.Contains(markdown, "### Trend") {
		t.Fatalf("expected header in markdown, got: %q", markdown)
	}
	if !strings.Contains(markdown, "Users: 1,240 (+12%)") {
		t.Fatalf("expected summary line in markdown, got: %q", markdown)
	}
	if !strings.Contains(markdown, "▁") || !strings.Contains(markdown, "█") {
		t.Fatalf("expected sparkline glyphs in markdown, got: %q", markdown)
	}

	withoutDelta := SparklineBlock{Header: "Trend", Label: "Users", Value: "1,240", Points: []int{1, 2, 3}}
	withoutDeltaMarkdown, err := withoutDelta.RenderMarkdown(RenderContext{})
	if err != nil {
		t.Fatalf("RenderMarkdown without delta returned error: %v", err)
	}
	if strings.Contains(withoutDeltaMarkdown, "(") {
		t.Fatalf("expected no delta section in markdown when unset, got: %q", withoutDeltaMarkdown)
	}
}

func TestStackedBarBlockTemplateDataAndMarkdown(t *testing.T) {
	block := StackedBarBlock{
		Header:     " Funnel ",
		TotalLabel: "Total",
		TotalValue: "120k",
		Rows: []StackedBarRow{{
			Label: " Q1 ",
			Segments: []StackedBarSegment{
				{Label: " Won ", Percent: 140, Value: " "},
				{Label: " Lost ", Percent: -5, Value: "12"},
			},
		}},
	}

	normalized := block.TemplateData().(StackedBarBlock)
	if normalized.Rows[0].Label != "Q1" {
		t.Fatalf("expected trimmed row label, got %q", normalized.Rows[0].Label)
	}
	if normalized.Rows[0].Segments[0].Percent != 100 {
		t.Fatalf("expected upper clamp to 100, got %d", normalized.Rows[0].Segments[0].Percent)
	}
	if normalized.Rows[0].Segments[1].Percent != 0 {
		t.Fatalf("expected lower clamp to 0, got %d", normalized.Rows[0].Segments[1].Percent)
	}

	markdown, err := block.RenderMarkdown(RenderContext{})
	if err != nil {
		t.Fatalf("RenderMarkdown returned error: %v", err)
	}
	if !strings.Contains(markdown, "### Funnel") || !strings.Contains(markdown, "**Q1:**") {
		t.Fatalf("expected header and row in markdown, got: %q", markdown)
	}
	if !strings.Contains(markdown, "**Total:** 120k") {
		t.Fatalf("expected total line in markdown, got: %q", markdown)
	}
	if !strings.Contains(markdown, "Won 100%") || !strings.Contains(markdown, "Lost 12") {
		t.Fatalf("expected segment values in markdown, got: %q", markdown)
	}
}

func TestProgressBlockTemplateDataAndMarkdown(t *testing.T) {
	block := ProgressBlock{
		Header: " Delivery ",
		Items: []ProgressItem{
			{Label: "Build", Percent: 33, Value: ""},
			{Label: "Deploy", Percent: 120, Value: "done"},
		},
	}

	normalized := block.TemplateData().(ProgressBlock)
	if normalized.Items[0].Value != "33%" {
		t.Fatalf("expected default value from percent, got %q", normalized.Items[0].Value)
	}
	if normalized.Items[1].Percent != 100 {
		t.Fatalf("expected upper clamp to 100, got %d", normalized.Items[1].Percent)
	}

	markdown, err := block.RenderMarkdown(RenderContext{})
	if err != nil {
		t.Fatalf("RenderMarkdown returned error: %v", err)
	}
	if !strings.Contains(markdown, "### Delivery") || !strings.Contains(markdown, "Build") {
		t.Fatalf("expected progress markdown content, got: %q", markdown)
	}
	if !strings.Contains(markdown, "█") {
		t.Fatalf("expected progress bar glyphs, got: %q", markdown)
	}
}

func TestDistributionBlockTemplateDataAndMarkdown(t *testing.T) {
	block := DistributionBlock{
		Header: "Latency",
		Buckets: []DistributionBucket{
			{Label: "0-50ms", Count: -2},
			{Label: "50-100ms", Count: 10},
			{Label: "100-200ms", Count: 5},
		},
	}

	normalized := block.TemplateData().(DistributionBlock)
	if normalized.Buckets[0].Count != 0 {
		t.Fatalf("expected lower clamp on count, got %d", normalized.Buckets[0].Count)
	}
	if normalized.Buckets[1].WidthPercent != 100 {
		t.Fatalf("expected max bucket width 100, got %d", normalized.Buckets[1].WidthPercent)
	}
	if normalized.Buckets[2].WidthPercent != 50 {
		t.Fatalf("expected proportional bucket width 50, got %d", normalized.Buckets[2].WidthPercent)
	}

	markdown, err := block.RenderMarkdown(RenderContext{})
	if err != nil {
		t.Fatalf("RenderMarkdown returned error: %v", err)
	}
	if !strings.Contains(markdown, "### Latency") || !strings.Contains(markdown, "(10)") {
		t.Fatalf("expected distribution markdown content, got: %q", markdown)
	}
	if !strings.Contains(markdown, "(0)") {
		t.Fatalf("expected clamped zero-count bucket in markdown, got: %q", markdown)
	}
}
