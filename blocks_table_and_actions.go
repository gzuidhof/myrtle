package myrtle

import (
	"strings"

	"github.com/gzuidhof/myrtle/theme"
)

// TableBlock renders tabular data with configurable density and styling.
type TableBlock struct {
	Header                   string
	Columns                  []string
	Rows                     [][]string
	LegendSwatches           []string
	HasLegendSwatches        bool
	ZebraRows                bool
	Compact                  bool
	Density                  TableDensityValue
	HeaderTone               TableHeaderToneValue
	BorderStyle              TableBorderStyleValue
	RightAlignNumericColumns bool
	EmphasizeTotalRow        bool
	ColumnAlignments         map[int]TableColumnAlignmentValue
	InsetMode                InsetMode
}

// TableColumnAlignmentValue defines horizontal alignment for a table column.
type TableColumnAlignmentValue string

// TableDensityValue defines compactness for table row spacing.
type TableDensityValue string

// TableHeaderToneValue defines the visual tone used by the table header row.
type TableHeaderToneValue string

// TableBorderStyleValue defines the border line style used by table separators.
type TableBorderStyleValue string

const (
	TableColumnAlignmentStart  TableColumnAlignmentValue = "start"
	TableColumnAlignmentCenter TableColumnAlignmentValue = "center"
	TableColumnAlignmentEnd    TableColumnAlignmentValue = "end"

	TableDensityCompact TableDensityValue = "compact"
	TableDensityNormal  TableDensityValue = "normal"
	TableDensityRelaxed TableDensityValue = "relaxed"

	TableHeaderTonePrimary TableHeaderToneValue = "primary"
	TableHeaderToneMuted   TableHeaderToneValue = "muted"
	TableHeaderTonePlain   TableHeaderToneValue = "plain"

	TableBorderStyleSolid  TableBorderStyleValue = "solid"
	TableBorderStyleDashed TableBorderStyleValue = "dashed"
	TableBorderStyleDotted TableBorderStyleValue = "dotted"
)

func (block TableBlock) Kind() theme.BlockKind {
	return theme.BlockKindTable
}

func (block TableBlock) TemplateData() any {
	normalized := block
	normalized.Header = strings.TrimSpace(normalized.Header)
	normalized.HasLegendSwatches = normalized.HasLegendSwatches || len(normalized.LegendSwatches) > 0
	if normalized.HasLegendSwatches {
		swatches := make([]string, 0, len(normalized.LegendSwatches))
		for _, color := range normalized.LegendSwatches {
			swatches = append(swatches, strings.TrimSpace(color))
		}
		normalized.LegendSwatches = swatches
	}
	if normalized.Density == "" {
		if normalized.Compact {
			normalized.Density = TableDensityCompact
		} else {
			normalized.Density = TableDensityNormal
		}
	} else {
		normalized.Density = normalizedTableDensity(normalized.Density)
	}
	normalized.HeaderTone = normalizedTableHeaderTone(normalized.HeaderTone)
	normalized.BorderStyle = normalizedTableBorderStyle(normalized.BorderStyle)
	if len(block.ColumnAlignments) > 0 {
		normalized.ColumnAlignments = make(map[int]TableColumnAlignmentValue, len(block.ColumnAlignments))
		for index, alignment := range block.ColumnAlignments {
			normalized.ColumnAlignments[index] = normalizedTableColumnAlignment(alignment)
		}
	}

	return normalized
}

func (block TableBlock) RenderText(_ RenderContext) (string, error) {
	header := strings.TrimSpace(block.Header)

	var parts []string
	if header != "" {
		parts = append(parts, "[ "+header+" ]")
		parts = append(parts, strings.Repeat("-", max(8, min(48, len(header)+4))))
	}

	if len(block.Columns) > 0 {
		headerLine := strings.Join(block.Columns, " - ")
		parts = append(parts, headerLine)
		parts = append(parts, strings.Repeat("-", max(8, min(72, len(headerLine)))))
	}

	for _, row := range block.Rows {
		parts = append(parts, strings.Join(row, " - "))
	}

	return strings.Join(parts, "\n"), nil
}

func (block TableBlock) LayoutSpec() LayoutSpec {
	return normalizedLayoutSpec(LayoutSpec{InsetMode: block.InsetMode})
}

// VerificationCodeBlock renders a labeled one-time verification code.
type VerificationCodeBlock struct {
	Value              string
	Label              string
	Tone               Tone
	UseMonospace       bool
	CharacterSpacingEm float64
	InsetMode          InsetMode

	useMonospaceSet       bool
	characterSpacingEmSet bool
}

func (block VerificationCodeBlock) Kind() theme.BlockKind {
	return theme.BlockKindVerificationCode
}

func (block VerificationCodeBlock) TemplateData() any {
	normalized := block
	normalized.Tone = normalizedTone(normalized.Tone)
	if !normalized.useMonospaceSet {
		normalized.UseMonospace = true
	}
	if !normalized.characterSpacingEmSet {
		normalized.CharacterSpacingEm = 0.16
	}
	if normalized.CharacterSpacingEm < 0 {
		normalized.CharacterSpacingEm = 0
	}

	return normalized
}

func (block VerificationCodeBlock) RenderText(_ RenderContext) (string, error) {
	value := strings.TrimSpace(block.Value)
	if value == "" {
		return "", nil
	}
	label := strings.TrimSpace(block.Label)

	if label == "" {
		return value, nil
	}

	return label + "\n\n" + value, nil
}

func (block VerificationCodeBlock) LayoutSpec() LayoutSpec {
	return normalizedLayoutSpec(LayoutSpec{InsetMode: block.InsetMode})
}

// FreeMarkdownBlock renders raw markdown content.
type FreeMarkdownBlock struct {
	Markdown string
}

func (block FreeMarkdownBlock) Kind() theme.BlockKind {
	return theme.BlockKindFreeMarkdown
}

func (block FreeMarkdownBlock) TemplateData() any {
	return block
}

func (block FreeMarkdownBlock) RenderText(_ RenderContext) (string, error) {
	return strings.TrimSpace(block.Markdown), nil
}

func (block FreeMarkdownBlock) LayoutSpec() LayoutSpec { return defaultLayoutSpec() }

func normalizedTableColumnAlignment(value TableColumnAlignmentValue) TableColumnAlignmentValue {
	switch value {
	case TableColumnAlignmentCenter, TableColumnAlignmentEnd:
		return value
	default:
		return TableColumnAlignmentStart
	}
}

func normalizedTableDensity(value TableDensityValue) TableDensityValue {
	switch value {
	case TableDensityCompact, TableDensityRelaxed:
		return value
	default:
		return TableDensityNormal
	}
}

func normalizedTableHeaderTone(value TableHeaderToneValue) TableHeaderToneValue {
	switch value {
	case TableHeaderToneMuted, TableHeaderTonePlain:
		return value
	default:
		return TableHeaderTonePrimary
	}
}

func normalizedTableBorderStyle(value TableBorderStyleValue) TableBorderStyleValue {
	switch value {
	case TableBorderStyleDashed, TableBorderStyleDotted:
		return value
	default:
		return TableBorderStyleSolid
	}
}
