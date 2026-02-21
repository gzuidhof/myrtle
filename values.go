package myrtle

import "github.com/gzuidhof/myrtle/theme"

func normalizeValues(values theme.Values) theme.Values {
	normalized := values

	if normalized.Styles.PrimaryColor == "" {
		normalized.Styles.PrimaryColor = "#2563eb"
	}
	if normalized.Styles.TextColor == "" {
		normalized.Styles.TextColor = "#111827"
	}
	if normalized.Styles.MutedTextColor == "" {
		normalized.Styles.MutedTextColor = "#6b7280"
	}
	if normalized.Styles.BorderColor == "" {
		normalized.Styles.BorderColor = "#e5e7eb"
	}
	if normalized.Styles.CodeBackgroundColor == "" {
		normalized.Styles.CodeBackgroundColor = "#f8fafc"
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
