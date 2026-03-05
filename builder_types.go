package myrtle

import (
	"sync"

	"github.com/gzuidhof/myrtle/theme"
)

// Builder composes email content and rendering configuration before Build is called.
type Builder struct {
	mu                     sync.Mutex
	header                 *HeaderSection
	footer                 *FooterSection
	preheader              string
	preheaderPaddingRepeat int
	headerMode             HeaderMode
	footerMode             FooterMode
	values                 theme.Values
	blocks                 []Block
	theme                  theme.Theme
}

// HeaderMode controls whether an email header is auto-rendered, forced, or disabled.
type HeaderMode int

const (
	// HeaderModeAuto includes a header only when one is configured.
	HeaderModeAuto HeaderMode = iota
	// HeaderModeEnabled forces header rendering even when empty defaults apply.
	HeaderModeEnabled
	// HeaderModeDisabled suppresses header rendering.
	HeaderModeDisabled
)

// FooterMode controls whether an email footer is auto-rendered, forced, or disabled.
type FooterMode int

const (
	// FooterModeAuto includes a footer only when one is configured.
	FooterModeAuto FooterMode = iota
	// FooterModeEnabled forces footer rendering even when empty defaults apply.
	FooterModeEnabled
	// FooterModeDisabled suppresses footer rendering.
	FooterModeDisabled
)

// BuilderOption configures a Builder instance at creation time.
type BuilderOption func(*Builder)

// WithStyles overrides theme style tokens for the builder.
func WithStyles(value theme.Styles) BuilderOption {
	return func(builder *Builder) {
		builder.values.Styles = value
	}
}

// WithDirection sets the email direction (LTR or RTL) for builder output.
func WithDirection(value theme.Direction) BuilderOption {
	return func(builder *Builder) {
		if value == theme.DirectionRTL {
			builder.values.Direction = theme.DirectionRTL
			return
		}

		builder.values.Direction = theme.DirectionLTR
	}
}

// WithMSOCompatibility sets Outlook-specific compatibility fallback behavior for rendered HTML.
func WithMSOCompatibility(value theme.MSOCompatibilityMode) BuilderOption {
	return func(builder *Builder) {
		if value == theme.MSOCompatibilityModeOff {
			builder.values.Styles.MSOCompatibility = theme.MSOCompatibilityModeOff
			return
		}

		builder.values.Styles.MSOCompatibility = theme.MSOCompatibilityModeOn
	}
}

// WithOutlookCompatibility enables or disables Outlook-specific compatibility fallbacks.
func WithOutlookCompatibility(enabled bool) BuilderOption {
	if enabled {
		return WithMSOCompatibility(theme.MSOCompatibilityModeOn)
	}

	return WithMSOCompatibility(theme.MSOCompatibilityModeOff)
}

// WithHeaderMode controls automatic header behavior for the builder.
func WithHeaderMode(mode HeaderMode) BuilderOption {
	return func(builder *Builder) {
		switch mode {
		case HeaderModeEnabled:
			builder.headerMode = HeaderModeEnabled
		case HeaderModeDisabled:
			builder.headerMode = HeaderModeDisabled
			builder.header = nil
		default:
			builder.headerMode = HeaderModeAuto
		}
	}
}

// WithFooterMode controls automatic footer behavior for the builder.
func WithFooterMode(mode FooterMode) BuilderOption {
	return func(builder *Builder) {
		switch mode {
		case FooterModeEnabled:
			builder.footerMode = FooterModeEnabled
		case FooterModeDisabled:
			builder.footerMode = FooterModeDisabled
			builder.footer = nil
		default:
			builder.footerMode = FooterModeAuto
		}
	}
}

// WithHeader sets a header block and applies header options.
func WithHeader(block Block, options ...HeaderOption) BuilderOption {
	return func(builder *Builder) {
		if block == nil {
			builder.header = nil
			builder.headerMode = HeaderModeDisabled
			return
		}

		header := &HeaderSection{Block: block}
		for _, option := range options {
			if option == nil {
				continue
			}

			option(header)
		}

		builder.headerMode = HeaderModeEnabled
		builder.header = header
	}
}

// WithFooter sets a footer block and applies footer options.
func WithFooter(block Block, options ...FooterOption) BuilderOption {
	return func(builder *Builder) {
		if block == nil {
			builder.footer = nil
			builder.footerMode = FooterModeDisabled
			return
		}

		footer := &FooterSection{Block: block, Placement: FooterPlacementInside}
		for _, option := range options {
			if option == nil {
				continue
			}

			option(footer)
		}

		builder.footerMode = FooterModeEnabled
		builder.footer = footer
	}
}

// WithFooterOptions is an alias for WithFooter.
func WithFooterOptions(block Block, options ...FooterOption) BuilderOption {
	return WithFooter(block, options...)
}

// WithHeaderOptions is an alias for WithHeader.
func WithHeaderOptions(block Block, options ...HeaderOption) BuilderOption {
	return WithHeader(block, options...)
}

// NewBuilder creates a new email builder for the given theme.
// Optional BuilderOption values can set initial header, footer, and styling.
func NewBuilder(themeImpl theme.Theme, options ...BuilderOption) *Builder {
	if themeImpl == nil {
		panic("myrtle: theme is required")
	}

	builder := &Builder{
		headerMode:             HeaderModeAuto,
		footerMode:             FooterModeAuto,
		preheaderPaddingRepeat: defaultPreheaderPaddingRepeat,
		theme:                  themeImpl,
	}

	for _, option := range options {
		option(builder)
	}

	return builder
}

// Clone returns a new builder initialized with the current builder state.
// It is useful for deriving variants from a shared baseline configuration.
//
// The returned builder can be mutated independently, making it suitable for
// per-goroutine customization based on a shared prototype builder.
func (builder *Builder) Clone() *Builder {
	builder.mu.Lock()
	defer builder.mu.Unlock()

	return &Builder{
		header:                 cloneHeader(builder.header),
		footer:                 cloneFooter(builder.footer),
		preheader:              builder.preheader,
		preheaderPaddingRepeat: builder.preheaderPaddingRepeat,
		headerMode:             builder.headerMode,
		footerMode:             builder.footerMode,
		values:                 builder.values,
		blocks:                 append([]Block(nil), builder.blocks...),
		theme:                  builder.theme,
	}
}

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
		ColorPrimary:              "#265cff",
		ColorSecondary:            "#10b981",
		ColorText:                 "#111827",
		ColorTextMuted:            "#6b7280",
		ColorBorder:               "#e5e7eb",
		ColorCodeBackground:       "#f8fafc",
		ColorPageBackground:       "#f3f4f6",
		ColorMainBackground:       "#ffffff",
		ColorSurface:              "#ffffff",
		ColorSurfaceMuted:         "#f8fafc",
		ColorTextOnSolid:          "#ffffff",
		ColorInfo:                 "#2563eb",
		ColorInfoBorder:           "#93c5fd",
		ColorInfoBackground:       "#eff6ff",
		ColorInfoText:             "#1d4ed8",
		ColorSuccess:              "#16a34a",
		ColorSuccessBorder:        "#86efac",
		ColorSuccessBackground:    "#f0fdf4",
		ColorSuccessText:          "#15803d",
		ColorWarning:              "#ca8a04",
		ColorWarningBorder:        "#fcd34d",
		ColorWarningBackground:    "#fffbeb",
		ColorWarningText:          "#92400e",
		ColorDanger:               "#dc2626",
		ColorDangerBorder:         "#fca5a5",
		ColorDangerBackground:     "#fef2f2",
		ColorDangerText:           "#b91c1c",
		WidthMain:                 "100%",
		MaxWidthMain:              "640px",
		OuterPadding:              "24px",
		OutsideContentInset:       "24px",
		MainContentBodyTopSpacing: "24px",
		MSOCompatibility:          theme.MSOCompatibilityModeOn,
		RadiusMain:                "12px",
		RadiusElement:             "10px",
		RadiusButton:              "8px",
		RadiusPill:                "999px",
		TableLegendSwatchSize:     "12px",
		TableLegendSwatchRadius:   "4px",
		TableLegendSwatchBorder:   "",
		FontFamilyBase:            "system-ui,-apple-system,BlinkMacSystemFont,'Segoe UI',Roboto,'Helvetica Neue',Arial,sans-serif",
		FontFamilyMono:            "ui-monospace,SFMono-Regular,Menlo,Monaco,Consolas,monospace",
		FontSizeBase:              "14px",
		LineHeightBase:            "1.6",
		FontWeightHeading:         "700",
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
	if overrides.MainContentBodyTopSpacing != "" {
		merged.MainContentBodyTopSpacing = overrides.MainContentBodyTopSpacing
	}
	if overrides.MSOCompatibility != "" {
		merged.MSOCompatibility = overrides.MSOCompatibility
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
	if overrides.TableLegendSwatchSize != "" {
		merged.TableLegendSwatchSize = overrides.TableLegendSwatchSize
	}
	if overrides.TableLegendSwatchRadius != "" {
		merged.TableLegendSwatchRadius = overrides.TableLegendSwatchRadius
	}
	if overrides.TableLegendSwatchBorder != "" {
		merged.TableLegendSwatchBorder = overrides.TableLegendSwatchBorder
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
