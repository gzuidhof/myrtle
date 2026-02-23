package myrtle

import "github.com/gzuidhof/myrtle/theme"

func normalizeValues(values theme.Values, defaultStyles theme.Styles) theme.Values {
	normalized := values
	normalized.Styles = mergeStyles(defaultStyles, values.Styles)

	if normalized.Styles.ColorPrimary == "" {
		normalized.Styles.ColorPrimary = "#265cff"
	}
	if normalized.Styles.ColorSecondary == "" {
		normalized.Styles.ColorSecondary = "#10b981"
	}
	if normalized.Styles.ColorText == "" {
		normalized.Styles.ColorText = "#111827"
	}
	if normalized.Styles.ColorTextMuted == "" {
		normalized.Styles.ColorTextMuted = "#6b7280"
	}
	if normalized.Styles.ColorBorder == "" {
		normalized.Styles.ColorBorder = "#e5e7eb"
	}
	if normalized.Styles.ColorCodeBackground == "" {
		normalized.Styles.ColorCodeBackground = "#f8fafc"
	}
	if normalized.Styles.ColorPageBackground == "" {
		normalized.Styles.ColorPageBackground = "#f3f4f6"
	}
	if normalized.Styles.ColorMainBackground == "" {
		normalized.Styles.ColorMainBackground = "#ffffff"
	}
	if normalized.Styles.BorderMain == "" {
		normalized.Styles.BorderMain = "1px solid " + normalized.Styles.ColorBorder
	}
	if normalized.Styles.RadiusMain == "" {
		normalized.Styles.RadiusMain = "12px"
	}
	if normalized.Styles.FontFamilyBase == "" {
		normalized.Styles.FontFamilyBase = "system-ui,-apple-system,BlinkMacSystemFont,'Segoe UI',Roboto,'Helvetica Neue',Arial,sans-serif"
	}
	if normalized.Styles.FontFamilyMono == "" {
		normalized.Styles.FontFamilyMono = "ui-monospace,SFMono-Regular,Menlo,Monaco,Consolas,monospace"
	}
	if normalized.Styles.FontSizeBase == "" {
		normalized.Styles.FontSizeBase = "14px"
	}
	if normalized.Styles.LineHeightBase == "" {
		normalized.Styles.LineHeightBase = "1.6"
	}
	if normalized.Styles.FontWeightHeading == "" {
		normalized.Styles.FontWeightHeading = "700"
	}

	if normalized.LogoAlt == "" {
		switch {
		case normalized.ProductName != "":
			normalized.LogoAlt = normalized.ProductName
		default:
			normalized.LogoAlt = "Logo"
		}
	}

	return normalized
}

func mergeStyles(defaults, overrides theme.Styles) theme.Styles {
	merged := defaults

	if overrides.ColorPrimary != "" {
		merged.ColorPrimary = overrides.ColorPrimary
	}
	if overrides.ColorSecondary != "" {
		merged.ColorSecondary = overrides.ColorSecondary
	}
	if overrides.ColorText != "" {
		merged.ColorText = overrides.ColorText
	}
	if overrides.ColorTextMuted != "" {
		merged.ColorTextMuted = overrides.ColorTextMuted
	}
	if overrides.ColorBorder != "" {
		merged.ColorBorder = overrides.ColorBorder
	}
	if overrides.ColorCodeBackground != "" {
		merged.ColorCodeBackground = overrides.ColorCodeBackground
	}
	if overrides.ColorPageBackground != "" {
		merged.ColorPageBackground = overrides.ColorPageBackground
	}
	if overrides.ColorMainBackground != "" {
		merged.ColorMainBackground = overrides.ColorMainBackground
	}
	if overrides.BorderMain != "" {
		merged.BorderMain = overrides.BorderMain
	}
	if overrides.RadiusMain != "" {
		merged.RadiusMain = overrides.RadiusMain
	}
	if overrides.FontFamilyBase != "" {
		merged.FontFamilyBase = overrides.FontFamilyBase
	}
	if overrides.FontFamilyMono != "" {
		merged.FontFamilyMono = overrides.FontFamilyMono
	}
	if overrides.FontSizeBase != "" {
		merged.FontSizeBase = overrides.FontSizeBase
	}
	if overrides.LineHeightBase != "" {
		merged.LineHeightBase = overrides.LineHeightBase
	}
	if overrides.FontWeightHeading != "" {
		merged.FontWeightHeading = overrides.FontWeightHeading
	}

	return merged
}
