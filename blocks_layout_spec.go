package myrtle

import "strings"

func normalizedLayoutSpec(spec LayoutSpec) LayoutSpec {
	if spec.InsetMode != InsetModeNone && spec.InsetMode != InsetModeCustom {
		spec.InsetMode = InsetModeDefault
	}

	spec.CustomInset = strings.TrimSpace(spec.CustomInset)
	if spec.InsetMode == InsetModeCustom && spec.CustomInset == "" {
		spec.InsetMode = InsetModeDefault
	}

	return spec
}

func defaultLayoutSpec() LayoutSpec {
	return LayoutSpec{InsetMode: InsetModeDefault}
}
