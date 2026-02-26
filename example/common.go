package example

import "github.com/gzuidhof/myrtle"

const commonHeaderLogoWidth = 140
const commonHeaderLogoHref = "https://example.com"
const commonLegalCompany = "Myrtle Inc."
const commonLegalAddress = "123 Market St, San Francisco, CA"
const commonLegalManageURL = "https://example.com/preferences"
const commonLegalUnsubscribeURL = "https://example.com/unsubscribe"

func commonHeaderGroup(title string) myrtle.Block {
	return commonHeaderGroupWithAlt(title, title)
}

func commonHeaderGroupWithAlt(title, logoAlt string) myrtle.Block {
	_ = title
	return myrtle.NewGroup().
		AddImage(
			"/assets/logo.png",
			logoAlt,
			myrtle.ImageHref(commonHeaderLogoHref),
			myrtle.ImageWidth(commonHeaderLogoWidth),
			myrtle.ImageAlign(myrtle.ImageAlignmentCenter),
		)
}

func commonFooterGroup() myrtle.Block {
	return myrtle.NewGroup().
		AddLegal(commonLegalCompany, commonLegalAddress, commonLegalManageURL, commonLegalUnsubscribeURL)
}
