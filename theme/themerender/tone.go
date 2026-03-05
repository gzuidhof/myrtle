package themerender

import (
	"fmt"
	"strings"

	"github.com/gzuidhof/myrtle/theme"
)

func toneColor(tone any, values theme.Values, role string) string {
	styles := values.Styles
	normalizedTone := normalizeToneToken(tone)
	normalizedRole := normalizeToken(role)

	switch normalizedRole {
	case "strong":
		return toneStrongColor(normalizedTone, styles)
	case "border":
		switch normalizedTone {
		case "primary":
			return styles.ColorPrimary
		case "dark":
			return styles.ColorText
		case "info":
			return styles.ColorInfoBorder
		case "success":
			return styles.ColorSuccessBorder
		case "warning":
			return styles.ColorWarningBorder
		case "danger":
			return styles.ColorDangerBorder
		default:
			return styles.ColorBorder
		}
	case "background":
		switch normalizedTone {
		case "primary":
			return styles.ColorPrimary
		case "dark":
			return styles.ColorText
		case "info":
			return styles.ColorInfoBackground
		case "success":
			return styles.ColorSuccessBackground
		case "warning":
			return styles.ColorWarningBackground
		case "danger":
			return styles.ColorDangerBackground
		default:
			return styles.ColorSurface
		}
	case "text":
		return toneTextColor(normalizedTone, styles)
	case "muted-text":
		switch normalizedTone {
		case "primary":
			return styles.ColorTextOnSolid
		case "dark":
			return styles.ColorTextOnSolid
		case "info":
			return styles.ColorInfoText
		case "success":
			return styles.ColorSuccessText
		case "warning":
			return styles.ColorWarningText
		case "danger":
			return styles.ColorDangerText
		default:
			return styles.ColorTextMuted
		}
	case "link":
		switch normalizedTone {
		case "default", "", "muted":
			return styles.ColorPrimary
		default:
			return toneTextColor(normalizedTone, styles)
		}
	case "cta-background":
		switch normalizedTone {
		case "primary":
			return styles.ColorMainBackground
		case "dark":
			return styles.ColorTextOnSolid
		case "info":
			return styles.ColorInfoText
		case "success":
			return styles.ColorSuccessText
		case "warning":
			return styles.ColorWarningText
		case "danger":
			return styles.ColorDangerText
		default:
			return styles.ColorPrimary
		}
	case "cta-text":
		if normalizedTone == "primary" || normalizedTone == "dark" {
			return styles.ColorPrimary
		}

		return styles.ColorTextOnSolid
	default:
		return styles.ColorText
	}
}

func normalizeToken(value string) string {
	return strings.ToLower(strings.TrimSpace(value))
}

func normalizeToneToken(value any) string {
	normalized := normalizeToken(fmt.Sprint(value))

	switch normalized {
	case "error", "critical":
		return "danger"
	case "black":
		return "dark"
	default:
		return normalized
	}
}

func toneTextColor(normalizedTone string, styles theme.Styles) string {
	switch normalizedTone {
	case "muted":
		return styles.ColorTextMuted
	case "primary":
		return styles.ColorTextOnSolid
	case "dark":
		return styles.ColorTextOnSolid
	case "info":
		return styles.ColorInfoText
	case "success":
		return styles.ColorSuccessText
	case "warning":
		return styles.ColorWarningText
	case "danger":
		return styles.ColorDangerText
	default:
		return styles.ColorText
	}
}

func toneStrongColor(normalizedTone string, styles theme.Styles) string {
	switch normalizedTone {
	case "muted":
		return styles.ColorTextMuted
	case "dark":
		return styles.ColorText
	case "info":
		return styles.ColorInfo
	case "success":
		return styles.ColorSuccess
	case "warning":
		return styles.ColorWarning
	case "danger":
		return styles.ColorDanger
	default:
		return styles.ColorPrimary
	}
}
