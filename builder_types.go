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
