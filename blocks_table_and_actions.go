package myrtle

import (
	"fmt"
	"strings"

	"github.com/gzuidhof/myrtle/theme"
)

type TableBlock struct {
	Header                   string
	Columns                  []string
	Rows                     [][]string
	ZebraRows                bool
	Compact                  bool
	RightAlignNumericColumns bool
	EmphasizeTotalRow        bool
	ColumnAlignments         map[int]TableColumnAlignment
}

type TableColumnAlignment string

const (
	TableColumnAlignmentLeft   TableColumnAlignment = "left"
	TableColumnAlignmentCenter TableColumnAlignment = "center"
	TableColumnAlignmentRight  TableColumnAlignment = "right"
)

func (block TableBlock) Kind() theme.BlockKind {
	return theme.BlockKindTable
}

func (block TableBlock) TemplateData() any {
	normalized := block
	if len(block.ColumnAlignments) > 0 {
		normalized.ColumnAlignments = make(map[int]TableColumnAlignment, len(block.ColumnAlignments))
		for index, alignment := range block.ColumnAlignments {
			normalized.ColumnAlignments[index] = normalizedTableColumnAlignment(alignment)
		}
	}

	return normalized
}

func (block TableBlock) RenderMarkdown(_ RenderContext) (string, error) {
	var parts []string
	if strings.TrimSpace(block.Header) != "" {
		parts = append(parts, "### "+strings.TrimSpace(block.Header))
	}

	if len(block.Columns) > 0 {
		parts = append(parts, "| "+strings.Join(block.Columns, " | ")+" |")

		separators := make([]string, len(block.Columns))
		for index := range separators {
			separators[index] = "---"
		}
		parts = append(parts, "| "+strings.Join(separators, " | ")+" |")
	}

	for _, row := range block.Rows {
		parts = append(parts, "| "+strings.Join(row, " | ")+" |")
	}

	return strings.Join(parts, "\n"), nil
}

type ActionBlock struct {
	Instructions string
	ButtonLabel  string
	ButtonURL    string
}

func (block ActionBlock) Kind() theme.BlockKind {
	return theme.BlockKindAction
}

func (block ActionBlock) TemplateData() any {
	return block
}

func (block ActionBlock) RenderMarkdown(_ RenderContext) (string, error) {
	parts := make([]string, 0, 2)
	if strings.TrimSpace(block.Instructions) != "" {
		parts = append(parts, strings.TrimSpace(block.Instructions))
	}
	parts = append(parts, fmt.Sprintf("[%s](%s)", block.ButtonLabel, block.ButtonURL))
	return strings.Join(parts, "\n\n"), nil
}

type CodeBlock struct {
	Code  string
	Label string
}

func (block CodeBlock) Kind() theme.BlockKind {
	return theme.BlockKindCode
}

func (block CodeBlock) TemplateData() any {
	return block
}

func (block CodeBlock) RenderMarkdown(_ RenderContext) (string, error) {
	value := strings.TrimSpace(block.Code)
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

func (block FreeMarkdownBlock) RenderMarkdown(_ RenderContext) (string, error) {
	return strings.TrimSpace(block.Markdown), nil
}

func normalizedTableColumnAlignment(value TableColumnAlignment) TableColumnAlignment {
	switch value {
	case TableColumnAlignmentCenter, TableColumnAlignmentRight:
		return value
	default:
		return TableColumnAlignmentLeft
	}
}
