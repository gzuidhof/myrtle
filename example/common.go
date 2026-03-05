package example

import (
	"strconv"
	"strings"

	"github.com/gzuidhof/myrtle"
	"github.com/gzuidhof/myrtle/theme"
)

const (
	commonHeaderLogoWidth     = 120
	commonHeaderLogoHref      = "https://example.com"
	commonHeaderLogoSrc       = "/assets/logo.png"
	commonHeaderLogoLightSrc  = "/assets/logo-light.png"
	commonLegalCompany        = "Myrtle Inc."
	commonLegalAddress        = "Dam Square 1, 1012 JS Amsterdam, Netherlands"
	commonLegalManageURL      = "https://example.com/preferences"
	commonLegalUnsubscribeURL = "https://example.com/unsubscribe"
)

func commonHeaderGroup(title string, selectedTheme ...theme.Theme) myrtle.Block {
	return commonHeaderGroupWithAlt(title, title, selectedTheme...)
}

func commonHeaderGroupWithAlt(title, logoAlt string, selectedTheme ...theme.Theme) myrtle.Block {
	logoSrc := commonHeaderLogoSrc
	if len(selectedTheme) > 0 && usesDarkBackground(selectedTheme[0]) {
		logoSrc = commonHeaderLogoLightSrc
	}

	return commonHeaderGroupWithLogo(title, logoAlt, logoSrc)
}

func commonHeaderGroupWithLogo(title, logoAlt, logoSrc string) myrtle.Block {
	_ = title
	return myrtle.NewGroup().
		AddImage(
			logoSrc,
			logoAlt,
			myrtle.ImageHref(commonHeaderLogoHref),
			myrtle.ImageWidth(commonHeaderLogoWidth),
			myrtle.ImageAlign(myrtle.ImageAlignmentCenter),
		)
}

func usesDarkBackground(selectedTheme theme.Theme) bool {
	if selectedTheme == nil {
		return false
	}

	styles := selectedTheme.DefaultStyles()
	return isDarkHexColor(styles.ColorMainBackground)
}

func isDarkHexColor(value string) bool {
	color := strings.TrimSpace(value)
	color = strings.TrimPrefix(color, "#")

	if len(color) == 3 {
		color = strings.Repeat(string(color[0]), 2) + strings.Repeat(string(color[1]), 2) + strings.Repeat(string(color[2]), 2)
	}
	if len(color) != 6 {
		return false
	}

	r, errR := strconv.ParseInt(color[0:2], 16, 64)
	g, errG := strconv.ParseInt(color[2:4], 16, 64)
	b, errB := strconv.ParseInt(color[4:6], 16, 64)
	if errR != nil || errG != nil || errB != nil {
		return false
	}

	luma := (299*r + 587*g + 114*b) / 1000
	return luma < 128
}

func commonFooterGroup() myrtle.Block {
	return myrtle.NewGroup().
		AddLegal(commonLegalCompany, commonLegalAddress, commonLegalManageURL, commonLegalUnsubscribeURL)
}
