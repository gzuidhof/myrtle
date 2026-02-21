package myrtle

import (
	"strings"

	"github.com/gzuidhof/myrtle/theme"
)

type QuoteBlock struct {
	Text   string
	Author string
}

func (block QuoteBlock) Kind() theme.BlockKind {
	return theme.BlockKindQuote
}

func (block QuoteBlock) TemplateData() any {
	return block
}

func (block QuoteBlock) RenderMarkdown(_ RenderContext) (string, error) {
	text := strings.TrimSpace(block.Text)
	if text == "" {
		return "", nil
	}

	parts := []string{"> " + strings.ReplaceAll(text, "\n", "\n> ")}
	if strings.TrimSpace(block.Author) != "" {
		parts = append(parts, "> — "+strings.TrimSpace(block.Author))
	}

	return strings.Join(parts, "\n"), nil
}

type CalloutType string

const (
	CalloutTypeInfo     CalloutType = "info"
	CalloutTypeSuccess  CalloutType = "success"
	CalloutTypeWarning  CalloutType = "warning"
	CalloutTypeError    CalloutType = "error"
	CalloutTypeCritical CalloutType = "critical"
)

type CalloutVariant string

const (
	CalloutVariantSoft    CalloutVariant = "soft"
	CalloutVariantOutline CalloutVariant = "outline"
	CalloutVariantSolid   CalloutVariant = "solid"
)

type CalloutBlock struct {
	Type      CalloutType
	Variant   CalloutVariant
	Title     string
	Body      string
	LinkLabel string
	LinkURL   string
}

func (block CalloutBlock) Kind() theme.BlockKind {
	return theme.BlockKindCallout
}

func (block CalloutBlock) TemplateData() any {
	normalized := block
	normalized.Type = normalizedCalloutType(block.Type)
	normalized.Variant = normalizedCalloutVariant(block.Variant)
	return normalized
}

func (block CalloutBlock) RenderMarkdown(_ RenderContext) (string, error) {
	parts := make([]string, 0, 3)
	if strings.TrimSpace(block.Title) != "" {
		parts = append(parts, "**"+strings.TrimSpace(block.Title)+"**")
	}
	if strings.TrimSpace(block.Body) != "" {
		parts = append(parts, strings.TrimSpace(block.Body))
	}
	if strings.TrimSpace(block.LinkLabel) != "" && strings.TrimSpace(block.LinkURL) != "" {
		parts = append(parts, "["+strings.TrimSpace(block.LinkLabel)+"]("+strings.TrimSpace(block.LinkURL)+")")
	}
	if len(parts) == 0 {
		return "", nil
	}

	label := strings.ToUpper(string(normalizedCalloutType(block.Type)))
	return "> **" + label + "**\n> \n> " + strings.ReplaceAll(strings.Join(parts, "\n\n"), "\n", "\n> "), nil
}

func normalizedCalloutType(value CalloutType) CalloutType {
	switch value {
	case CalloutTypeSuccess, CalloutTypeWarning, CalloutTypeError, CalloutTypeCritical:
		return value
	default:
		return CalloutTypeInfo
	}
}

func normalizedCalloutVariant(value CalloutVariant) CalloutVariant {
	switch value {
	case CalloutVariantOutline, CalloutVariantSolid:
		return value
	default:
		return CalloutVariantSoft
	}
}
