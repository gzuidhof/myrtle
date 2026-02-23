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

type MessageBlock struct {
	SenderName   string
	SenderHandle string
	AvatarURL    string
	LogoAlt      string
	LogoHref     string
	Subject      string
	Preview      string
	SentAt       string
	Platform     string
	URL          string
	ActionLabel  string
	ActionURL    string
}

type MessageDigestBlock struct {
	Title     string
	Subtitle  string
	Messages  []MessageBlock
	EmptyText string
	Footer    string
	MaxItems  int
}

func (block CalloutBlock) Kind() theme.BlockKind {
	return theme.BlockKindCallout
}

func (block MessageBlock) Kind() theme.BlockKind {
	return theme.BlockKindMessage
}

func (block MessageDigestBlock) Kind() theme.BlockKind {
	return theme.BlockKindMessageDigest
}

func (block CalloutBlock) TemplateData() any {
	normalized := block
	normalized.Type = normalizedCalloutType(block.Type)
	normalized.Variant = normalizedCalloutVariant(block.Variant)
	return normalized
}

func (block MessageBlock) TemplateData() any {
	return normalizeMessageBlock(block)
}

func (block MessageDigestBlock) TemplateData() any {
	normalized := MessageDigestBlock{
		Title:     strings.TrimSpace(block.Title),
		Subtitle:  strings.TrimSpace(block.Subtitle),
		Messages:  make([]MessageBlock, 0, len(block.Messages)),
		EmptyText: strings.TrimSpace(block.EmptyText),
		Footer:    strings.TrimSpace(block.Footer),
		MaxItems:  block.MaxItems,
	}

	for _, message := range block.Messages {
		normalizedMessage := normalizeMessageBlock(message)
		if normalizedMessage.Subject == "" && normalizedMessage.Preview == "" && normalizedMessage.SenderName == "" && normalizedMessage.SenderHandle == "" {
			continue
		}
		normalized.Messages = append(normalized.Messages, normalizedMessage)
	}

	if normalized.MaxItems > 0 && len(normalized.Messages) > normalized.MaxItems {
		normalized.Messages = normalized.Messages[:normalized.MaxItems]
	}

	return normalized
}

func normalizeMessageBlock(block MessageBlock) MessageBlock {
	return MessageBlock{
		SenderName:   strings.TrimSpace(block.SenderName),
		SenderHandle: strings.TrimSpace(block.SenderHandle),
		AvatarURL:    strings.TrimSpace(block.AvatarURL),
		LogoAlt:      strings.TrimSpace(block.LogoAlt),
		LogoHref:     strings.TrimSpace(block.LogoHref),
		Subject:      strings.TrimSpace(block.Subject),
		Preview:      strings.TrimSpace(block.Preview),
		SentAt:       strings.TrimSpace(block.SentAt),
		Platform:     strings.TrimSpace(block.Platform),
		URL:          strings.TrimSpace(block.URL),
		ActionLabel:  strings.TrimSpace(block.ActionLabel),
		ActionURL:    strings.TrimSpace(block.ActionURL),
	}
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

func (block MessageBlock) RenderMarkdown(_ RenderContext) (string, error) {
	normalized := block.TemplateData().(MessageBlock)
	parts := make([]string, 0, 3)

	if normalized.Subject != "" {
		parts = append(parts, "**"+normalized.Subject+"**")
	}
	if normalized.Preview != "" {
		parts = append(parts, normalized.Preview)
	}

	meta := ""
	if normalized.SenderName != "" {
		meta = normalized.SenderName
	} else if normalized.SenderHandle != "" {
		meta = normalized.SenderHandle
	}
	if normalized.Platform != "" {
		if meta != "" {
			meta += " · "
		}
		meta += normalized.Platform
	}
	if normalized.SentAt != "" {
		if meta != "" {
			meta += " · "
		}
		meta += normalized.SentAt
	}
	if meta != "" {
		parts = append(parts, "_"+meta+"_")
	}

	if len(parts) == 0 {
		return "", nil
	}

	body := strings.Join(parts, "\n\n")

	linkLabel := normalized.ActionLabel
	if linkLabel == "" {
		linkLabel = "Open message"
	}

	if normalized.ActionURL != "" {
		return body + "\n\n[" + linkLabel + "](" + normalized.ActionURL + ")", nil
	}
	if normalized.URL != "" {
		return body + "\n\n[" + linkLabel + "](" + normalized.URL + ")", nil
	}

	return body, nil
}

func (block MessageDigestBlock) RenderMarkdown(_ RenderContext) (string, error) {
	normalized := block.TemplateData().(MessageDigestBlock)
	parts := make([]string, 0, len(normalized.Messages)+1)

	if normalized.Title != "" {
		parts = append(parts, "## "+normalized.Title)
	}
	if normalized.Subtitle != "" {
		parts = append(parts, normalized.Subtitle)
	}

	if len(normalized.Messages) == 0 {
		if normalized.EmptyText != "" {
			parts = append(parts, normalized.EmptyText)
		}
		return strings.Join(parts, "\n\n"), nil
	}

	for _, message := range normalized.Messages {
		subject := message.Subject
		if subject == "" {
			subject = "(no subject)"
		}

		line := "- "
		if message.URL != "" {
			line += "[" + subject + "](" + message.URL + ")"
		} else {
			line += subject
		}

		if message.Preview != "" {
			line += " — " + message.Preview
		}

		meta := messageMetadataLine(message)
		if meta != "" {
			line += " _(" + meta + ")_"
		}

		parts = append(parts, line)
	}

	if normalized.Footer != "" {
		parts = append(parts, normalized.Footer)
	}

	return strings.Join(parts, "\n\n"), nil
}

func messageMetadataLine(block MessageBlock) string {
	meta := ""
	if block.SenderName != "" {
		meta = block.SenderName
	} else if block.SenderHandle != "" {
		meta = block.SenderHandle
	}
	if block.Platform != "" {
		if meta != "" {
			meta += " · "
		}
		meta += block.Platform
	}
	if block.SentAt != "" {
		if meta != "" {
			meta += " · "
		}
		meta += block.SentAt
	}

	return meta
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
