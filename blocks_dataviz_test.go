package myrtle

import (
	"strings"
	"testing"
)

func TestNormalizedIntPoints(t *testing.T) {
	t.Parallel()

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
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

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
	t.Parallel()

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
			expectedRunes:  "+",
		},
		{
			name:           "flat line",
			points:         []int{5, 5, 5, 5},
			expectedLength: 4,
			expectedRunes:  "+",
		},
		{
			name:           "increasing range",
			points:         []int{0, 1, 2, 3, 4, 5, 6, 7},
			expectedLength: 8,
			expectedRunes:  ".:-=+*#@",
		},
		{
			name:           "large range",
			points:         []int{0, 100, 1000, 10000},
			expectedLength: 4,
			expectedRunes:  ".",
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			actual := sparklineGlyphs(test.points)
			if len([]rune(actual)) != test.expectedLength {
				t.Fatalf("expected glyph length %d, got %d (%q)", test.expectedLength, len([]rune(actual)), actual)
			}

			if test.expectedRunes != "" {
				for _, glyph := range test.expectedRunes {
					if !strings.ContainsRune(actual, glyph) {
						t.Fatalf("expected %q to contain glyph %q", actual, glyph)
					}
				}
			}
		})
	}
}

func TestSparklineBlockTemplateDataAndText(t *testing.T) {
	t.Parallel()

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

	text, err := block.RenderText(RenderContext{})
	if err != nil {
		t.Fatalf("RenderText returned error: %v", err)
	}
	if !strings.Contains(text, "Trend\n") {
		t.Fatalf("expected header in text output, got: %q", text)
	}
	if !strings.Contains(text, "Users: 1,240 (+12%)") {
		t.Fatalf("expected summary line in text, got: %q", text)
	}
	if !strings.Contains(text, ".") || !strings.Contains(text, "@") {
		t.Fatalf("expected sparkline glyphs in text output, got: %q", text)
	}

	withoutDelta := SparklineBlock{Header: "Trend", Label: "Users", Value: "1,240", Points: []int{1, 2, 3}}
	withoutDeltaText, err := withoutDelta.RenderText(RenderContext{})
	if err != nil {
		t.Fatalf("RenderText without delta returned error: %v", err)
	}
	if strings.Contains(withoutDeltaText, "(") {
		t.Fatalf("expected no delta section in text when unset, got: %q", withoutDeltaText)
	}
}

func TestStackedBarBlockTemplateDataAndText(t *testing.T) {
	t.Parallel()

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

	text, err := block.RenderText(RenderContext{})
	if err != nil {
		t.Fatalf("RenderText returned error: %v", err)
	}
	if !strings.Contains(text, "Funnel") || !strings.Contains(text, "Q1:") {
		t.Fatalf("expected header and row in text output, got: %q", text)
	}
	if !strings.Contains(text, "Total: 120k") {
		t.Fatalf("expected total line in text output, got: %q", text)
	}
	if !strings.Contains(text, "Won 100%") || !strings.Contains(text, "Lost 12") {
		t.Fatalf("expected segment values in text, got: %q", text)
	}
}

func TestProgressBlockTemplateDataAndText(t *testing.T) {
	t.Parallel()

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

	text, err := block.RenderText(RenderContext{})
	if err != nil {
		t.Fatalf("RenderText returned error: %v", err)
	}
	if !strings.Contains(text, "Delivery") || !strings.Contains(text, "Build") {
		t.Fatalf("expected progress text content, got: %q", text)
	}
	if !strings.Contains(text, "#") {
		t.Fatalf("expected progress bar glyphs, got: %q", text)
	}
}

func TestDistributionBlockTemplateDataAndText(t *testing.T) {
	t.Parallel()

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

	text, err := block.RenderText(RenderContext{})
	if err != nil {
		t.Fatalf("RenderText returned error: %v", err)
	}
	if !strings.Contains(text, "Latency") || !strings.Contains(text, "(10)") {
		t.Fatalf("expected distribution text content, got: %q", text)
	}
	if !strings.Contains(text, "(0)") {
		t.Fatalf("expected clamped zero-count bucket in text output, got: %q", text)
	}
}
