package myrtle

import "github.com/gzuidhof/myrtle/theme"

func normalizeValues(values theme.Values, defaultStyles theme.Styles) theme.Values {
	normalized := values
	if normalized.Direction != theme.DirectionRTL {
		normalized.Direction = theme.DirectionLTR
	}
	normalized.Styles = mergeStyles(defaultStylesStruct(), defaultStyles)
	normalized.Styles = mergeStyles(normalized.Styles, values.Styles)
	if normalized.Styles.BorderMain == "" {
		normalized.Styles.BorderMain = "1px solid " + normalized.Styles.ColorBorder
	}

	return normalized
}

func defaultStylesStruct() theme.Styles {
	return theme.Styles{
		ColorPrimary:           "#265cff",
		ColorSecondary:         "#10b981",
		ColorText:              "#111827",
		ColorTextMuted:         "#6b7280",
		ColorBorder:            "#e5e7eb",
		ColorCodeBackground:    "#f8fafc",
		ColorPageBackground:    "#f3f4f6",
		ColorMainBackground:    "#ffffff",
		ColorSurface:           "#ffffff",
		ColorSurfaceMuted:      "#f8fafc",
		ColorTextOnSolid:       "#ffffff",
		ColorInfo:              "#2563eb",
		ColorInfoBorder:        "#93c5fd",
		ColorInfoBackground:    "#eff6ff",
		ColorInfoText:          "#1d4ed8",
		ColorSuccess:           "#16a34a",
		ColorSuccessBorder:     "#86efac",
		ColorSuccessBackground: "#f0fdf4",
		ColorSuccessText:       "#15803d",
		ColorWarning:           "#ca8a04",
		ColorWarningBorder:     "#fcd34d",
		ColorWarningBackground: "#fffbeb",
		ColorWarningText:       "#92400e",
		ColorDanger:            "#dc2626",
		ColorDangerBorder:      "#fca5a5",
		ColorDangerBackground:  "#fef2f2",
		ColorDangerText:        "#b91c1c",
		WidthMain:              "100%",
		MaxWidthMain:           "640px",
		OuterPadding:           "24px",
		OutsideContentInset:    "24px",
		RadiusMain:             "12px",
		RadiusElement:          "10px",
		RadiusButton:           "8px",
		RadiusPill:             "999px",
		FontFamilyBase:         "system-ui,-apple-system,BlinkMacSystemFont,'Segoe UI',Roboto,'Helvetica Neue',Arial,sans-serif",
		FontFamilyMono:         "ui-monospace,SFMono-Regular,Menlo,Monaco,Consolas,monospace",
		FontSizeBase:           "14px",
		LineHeightBase:         "1.6",
		FontWeightHeading:      "700",
	}
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
	if overrides.ColorSurface != "" {
		merged.ColorSurface = overrides.ColorSurface
	}
	if overrides.ColorSurfaceMuted != "" {
		merged.ColorSurfaceMuted = overrides.ColorSurfaceMuted
	}
	if overrides.ColorTextOnSolid != "" {
		merged.ColorTextOnSolid = overrides.ColorTextOnSolid
	}
	if overrides.ColorInfo != "" {
		merged.ColorInfo = overrides.ColorInfo
	}
	if overrides.ColorInfoBorder != "" {
		merged.ColorInfoBorder = overrides.ColorInfoBorder
	}
	if overrides.ColorInfoBackground != "" {
		merged.ColorInfoBackground = overrides.ColorInfoBackground
	}
	if overrides.ColorInfoText != "" {
		merged.ColorInfoText = overrides.ColorInfoText
	}
	if overrides.ColorSuccess != "" {
		merged.ColorSuccess = overrides.ColorSuccess
	}
	if overrides.ColorSuccessBorder != "" {
		merged.ColorSuccessBorder = overrides.ColorSuccessBorder
	}
	if overrides.ColorSuccessBackground != "" {
		merged.ColorSuccessBackground = overrides.ColorSuccessBackground
	}
	if overrides.ColorSuccessText != "" {
		merged.ColorSuccessText = overrides.ColorSuccessText
	}
	if overrides.ColorWarning != "" {
		merged.ColorWarning = overrides.ColorWarning
	}
	if overrides.ColorWarningBorder != "" {
		merged.ColorWarningBorder = overrides.ColorWarningBorder
	}
	if overrides.ColorWarningBackground != "" {
		merged.ColorWarningBackground = overrides.ColorWarningBackground
	}
	if overrides.ColorWarningText != "" {
		merged.ColorWarningText = overrides.ColorWarningText
	}
	if overrides.ColorDanger != "" {
		merged.ColorDanger = overrides.ColorDanger
	}
	if overrides.ColorDangerBorder != "" {
		merged.ColorDangerBorder = overrides.ColorDangerBorder
	}
	if overrides.ColorDangerBackground != "" {
		merged.ColorDangerBackground = overrides.ColorDangerBackground
	}
	if overrides.ColorDangerText != "" {
		merged.ColorDangerText = overrides.ColorDangerText
	}
	if overrides.BorderMain != "" {
		merged.BorderMain = overrides.BorderMain
	}
	if overrides.WidthMain != "" {
		merged.WidthMain = overrides.WidthMain
	}
	if overrides.MaxWidthMain != "" {
		merged.MaxWidthMain = overrides.MaxWidthMain
	}
	if overrides.OuterPadding != "" {
		merged.OuterPadding = overrides.OuterPadding
	}
	if overrides.OutsideContentInset != "" {
		merged.OutsideContentInset = overrides.OutsideContentInset
	}
	if overrides.RadiusMain != "" {
		merged.RadiusMain = overrides.RadiusMain
	}
	if overrides.RadiusElement != "" {
		merged.RadiusElement = overrides.RadiusElement
	}
	if overrides.RadiusButton != "" {
		merged.RadiusButton = overrides.RadiusButton
	}
	if overrides.RadiusPill != "" {
		merged.RadiusPill = overrides.RadiusPill
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
