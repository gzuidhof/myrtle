package myrtle

import (
	"sync"

	"github.com/gzuidhof/myrtle/theme"
)

type Builder struct {
	mu         sync.Mutex
	header     *HeaderSection
	footer     *FooterSection
	preheader  string
	headerMode HeaderMode
	footerMode FooterMode
	values     theme.Values
	blocks     []Block
	theme      theme.Theme
}

type HeaderMode int

const (
	HeaderModeAuto HeaderMode = iota
	HeaderModeEnabled
	HeaderModeDisabled
)

type FooterMode int

const (
	FooterModeAuto FooterMode = iota
	FooterModeEnabled
	FooterModeDisabled
)

type BuilderOption func(*Builder)

func WithStyles(value theme.Styles) BuilderOption {
	return func(builder *Builder) {
		builder.values.Styles = value
	}
}

func WithDirection(value theme.Direction) BuilderOption {
	return func(builder *Builder) {
		if value == theme.DirectionRTL {
			builder.values.Direction = theme.DirectionRTL
			return
		}

		builder.values.Direction = theme.DirectionLTR
	}
}

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

func WithFooterOptions(block Block, options ...FooterOption) BuilderOption {
	return WithFooter(block, options...)
}

func WithHeaderOptions(block Block, options ...HeaderOption) BuilderOption {
	return WithHeader(block, options...)
}

func NewBuilder(themeImpl theme.Theme, options ...BuilderOption) *Builder {
	if themeImpl == nil {
		panic("myrtle: theme is required")
	}

	builder := &Builder{
		headerMode: HeaderModeAuto,
		footerMode: FooterModeAuto,
		theme:      themeImpl,
	}

	for _, option := range options {
		option(builder)
	}

	return builder
}

// Clone returns a new builder initialized with the current builder state.
//
// The returned builder can be mutated independently, making it suitable for
// per-goroutine customization based on a shared prototype builder.
func (builder *Builder) Clone() *Builder {
	builder.mu.Lock()
	defer builder.mu.Unlock()

	return &Builder{
		header:     cloneHeader(builder.header),
		footer:     cloneFooter(builder.footer),
		preheader:  builder.preheader,
		headerMode: builder.headerMode,
		footerMode: builder.footerMode,
		values:     builder.values,
		blocks:     append([]Block(nil), builder.blocks...),
		theme:      builder.theme,
	}
}
