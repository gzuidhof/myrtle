package myrtle

import (
	"regexp"
	"strings"

	"github.com/gzuidhof/myrtle/theme"
)

// QuoteBlock renders quoted text with optional attribution.
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

func (block QuoteBlock) RenderText(_ RenderContext) (string, error) {
	text := strings.TrimSpace(block.Text)
	if text == "" {
		return "", nil
	}
	author := strings.TrimSpace(block.Author)

	parts := []string{"\"" + text + "\""}
	if author != "" {
		parts = append(parts, "- "+author)
	}

	return strings.Join(parts, "\n"), nil
}

// CalloutVariant defines the visual treatment used for a callout block.
type CalloutVariant string

const (
	// CalloutVariantSoft renders a subtle filled callout.
	CalloutVariantSoft CalloutVariant = "soft"
	// CalloutVariantOutline renders a bordered callout.
	CalloutVariantOutline CalloutVariant = "outline"
	// CalloutVariantSolid renders a strong filled callout.
	CalloutVariantSolid CalloutVariant = "solid"
)

// CalloutBlock renders an emphasized informational or alert callout.
type CalloutBlock struct {
	Tone      Tone
	Variant   CalloutVariant
	Title     string
	Body      string
	LinkLabel string
	LinkURL   string
	InsetMode InsetMode
}

// MessageBlock renders one message item in a conversational digest format.
type MessageBlock struct {
	SenderName      string
	SenderHandle    string
	AvatarURL       string
	LogoAlt         string
	LogoHref        string
	Subject         string
	Preview         string
	PreviewMarkdown string
	SentAt          string
	Platform        string
	URL             string
	ActionLabel     string
	ActionURL       string
	InsetMode       InsetMode
}

// MessageDigestBlock renders a grouped list of message items.
type MessageDigestBlock struct {
	Title     string
	Subtitle  string
	Messages  []MessageBlock
	EmptyText string
	Footer    string
	MaxItems  int
	InsetMode InsetMode
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
	normalized.Tone = normalizedCalloutTone(block.Tone)
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
		InsetMode: normalizedLayoutSpec(LayoutSpec{InsetMode: block.InsetMode}).InsetMode,
	}

	for _, message := range block.Messages {
		normalizedMessage := normalizeMessageBlock(message)
		if normalizedMessage.Subject == "" && normalizedMessage.Preview == "" && normalizedMessage.PreviewMarkdown == "" && normalizedMessage.SenderName == "" && normalizedMessage.SenderHandle == "" {
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
		SenderName:      strings.TrimSpace(block.SenderName),
		SenderHandle:    strings.TrimSpace(block.SenderHandle),
		AvatarURL:       strings.TrimSpace(block.AvatarURL),
		LogoAlt:         strings.TrimSpace(block.LogoAlt),
		LogoHref:        strings.TrimSpace(block.LogoHref),
		Subject:         strings.TrimSpace(block.Subject),
		Preview:         strings.TrimSpace(block.Preview),
		PreviewMarkdown: strings.TrimSpace(block.PreviewMarkdown),
		SentAt:          strings.TrimSpace(block.SentAt),
		Platform:        strings.TrimSpace(block.Platform),
		URL:             strings.TrimSpace(block.URL),
		ActionLabel:     strings.TrimSpace(block.ActionLabel),
		ActionURL:       strings.TrimSpace(block.ActionURL),
		InsetMode:       normalizedLayoutSpec(LayoutSpec{InsetMode: block.InsetMode}).InsetMode,
	}
}

func (block CalloutBlock) RenderText(_ RenderContext) (string, error) {
	title := strings.TrimSpace(block.Title)
	body := strings.TrimSpace(block.Body)
	linkLabel := strings.TrimSpace(block.LinkLabel)
	linkURL := strings.TrimSpace(block.LinkURL)

	parts := make([]string, 0, 3)
	if title != "" {
		parts = append(parts, "**"+title+"**")
	}
	if body != "" {
		parts = append(parts, body)
	}
	if linkLabel != "" && linkURL != "" {
		parts = append(parts, linkLabel+" ("+linkURL+")")
	}
	if len(parts) == 0 {
		return "", nil
	}

	label := strings.ToUpper(string(normalizedCalloutTone(block.Tone)))
	return "[ " + label + " ]\n--------------------\n" + strings.Join(parts, "\n\n"), nil
}

func (block MessageBlock) RenderText(_ RenderContext) (string, error) {
	normalized := block.TemplateData().(MessageBlock)
	parts := make([]string, 0, 3)

	if normalized.Subject != "" {
		parts = append(parts, normalized.Subject)
	}
	previewText := messagePreviewText(normalized)
	if previewText != "" {
		parts = append(parts, previewText)
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
		parts = append(parts, meta)
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
		return body + "\n\n" + linkLabel + " (" + normalized.ActionURL + ")", nil
	}
	if normalized.URL != "" {
		return body + "\n\n" + linkLabel + " (" + normalized.URL + ")", nil
	}

	return body, nil
}

func (block MessageDigestBlock) RenderText(_ RenderContext) (string, error) {
	normalized := block.TemplateData().(MessageDigestBlock)
	parts := make([]string, 0, len(normalized.Messages)+1)

	if normalized.Title != "" {
		parts = append(parts, normalized.Title, strings.Repeat("-", min(48, max(8, len(normalized.Title)))))
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
			line += subject + " (" + message.URL + ")"
		} else {
			line += subject
		}

		previewText := messagePreviewText(message)
		if previewText != "" {
			line += " — " + previewText
		}

		meta := messageMetadataLine(message)
		if meta != "" {
			line += " (" + meta + ")"
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

var previewMarkdownLinkPattern = regexp.MustCompile(`\[([^\]]+)\]\(([^)]+)\)`)

func messagePreviewText(block MessageBlock) string {
	if block.Preview != "" {
		return block.Preview
	}
	if block.PreviewMarkdown == "" {
		return ""
	}

	result := previewMarkdownLinkPattern.ReplaceAllString(block.PreviewMarkdown, `$1 ($2)`)
	replacer := strings.NewReplacer("**", "", "__", "", "*", "", "_", "", "`", "")
	result = replacer.Replace(result)
	result = strings.ReplaceAll(result, "\n", " ")
	result = strings.TrimSpace(strings.Join(strings.Fields(result), " "))

	return result
}

func normalizedCalloutTone(value Tone) Tone {
	switch value {
	case ToneSuccess, ToneWarning, ToneDanger, ToneDark:
		return value
	default:
		return ToneInfo
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

func (block QuoteBlock) LayoutSpec() LayoutSpec { return defaultLayoutSpec() }

func (block CalloutBlock) LayoutSpec() LayoutSpec {
	return normalizedLayoutSpec(LayoutSpec{InsetMode: block.InsetMode})
}

func (block MessageBlock) LayoutSpec() LayoutSpec {
	return normalizedLayoutSpec(LayoutSpec{InsetMode: block.InsetMode})
}

func (block MessageDigestBlock) LayoutSpec() LayoutSpec {
	return normalizedLayoutSpec(LayoutSpec{InsetMode: block.InsetMode})
}
