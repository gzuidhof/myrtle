package myrtle

import (
	"strings"

	"github.com/gzuidhof/myrtle/theme"
)

type TableBlock struct {
	Header                   string
	Columns                  []string
	Rows                     [][]string
	ZebraRows                bool
	Compact                  bool
	Density                  TableDensityValue
	HeaderTone               TableHeaderToneValue
	BorderStyle              TableBorderStyleValue
	RightAlignNumericColumns bool
	EmphasizeTotalRow        bool
	ColumnAlignments         map[int]TableColumnAlignmentValue
}

type TableColumnAlignmentValue string

type TableDensityValue string

type TableHeaderToneValue string

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
	var parts []string
	if strings.TrimSpace(block.Header) != "" {
		header := strings.TrimSpace(block.Header)
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

type VerificationCodeBlock struct {
	Value string
	Label string
}

func (block VerificationCodeBlock) Kind() theme.BlockKind {
	return theme.BlockKindVerificationCode
}

func (block VerificationCodeBlock) TemplateData() any {
	return block
}

func (block VerificationCodeBlock) RenderText(_ RenderContext) (string, error) {
	value := strings.TrimSpace(block.Value)
	if value == "" {
		return "", nil
	}

	if strings.TrimSpace(block.Label) == "" {
		return value, nil
	}

	return strings.TrimSpace(block.Label) + "\n\n" + value, nil
}

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
