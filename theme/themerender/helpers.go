package themerender

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gzuidhof/myrtle"
	"github.com/gzuidhof/myrtle/theme"
)

func messageMetaLine(block myrtle.MessageBlock) string {
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

func isNumericLike(value string) bool {
	cleaned := strings.TrimSpace(value)
	if cleaned == "" {
		return false
	}

	cleaned = strings.ReplaceAll(cleaned, ",", "")
	cleaned = strings.TrimPrefix(cleaned, "$")
	cleaned = strings.TrimPrefix(cleaned, "€")
	cleaned = strings.TrimPrefix(cleaned, "£")
	cleaned = strings.TrimPrefix(cleaned, "¥")
	cleaned = strings.TrimSuffix(cleaned, "%")

	if strings.HasPrefix(cleaned, "(") && strings.HasSuffix(cleaned, ")") {
		cleaned = "-" + strings.TrimSuffix(strings.TrimPrefix(cleaned, "("), ")")
	}

	if cleaned == "" {
		return false
	}

	_, err := strconv.ParseFloat(cleaned, 64)
	return err == nil
}

func isDiscountLike(label, value string) bool {
	normalizedLabel := normalizeToken(label)
	if strings.Contains(normalizedLabel, "discount") || strings.Contains(normalizedLabel, "credit") {
		return true
	}

	normalizedValue := normalizeTokenPreserveCase(value)
	return strings.HasPrefix(normalizedValue, "-")
}

func isRTL(values theme.Values) bool {
	return values.Direction == theme.DirectionRTL
}

func physicalAlign(alignment any, values theme.Values) string {
	value := normalizeToken(fmt.Sprint(alignment))
	switch value {
	case "center":
		return "center"
	case "end":
		if isRTL(values) {
			return "left"
		}
		return "right"
	default:
		if isRTL(values) {
			return "right"
		}
		return "left"
	}
}

func physicalSide(side any, values theme.Values) string {
	value := normalizeToken(fmt.Sprint(side))
	switch value {
	case "end":
		if isRTL(values) {
			return "left"
		}
		return "right"
	default:
		if isRTL(values) {
			return "right"
		}
		return "left"
	}
}
