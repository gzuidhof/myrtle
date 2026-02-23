package myrtle

import (
	"sync"

	"github.com/gzuidhof/myrtle/theme"
)

type Builder struct {
	mu         sync.Mutex
	header     *HeaderSection
	preheader  string
	headerMode HeaderMode
	values     theme.Values
	blocks     []Block
	theme      theme.Theme
	registry   *Registry
}

type HeaderMode int

const (
	HeaderModeAuto HeaderMode = iota
	HeaderModeEnabled
	HeaderModeDisabled
)

type BuilderOption func(*Builder)

func WithStyles(value theme.Styles) BuilderOption {
	return func(builder *Builder) {
		builder.values.Styles = value
	}
}

func WithRegistry(value *Registry) BuilderOption {
	return func(builder *Builder) {
		if value != nil {
			builder.registry = value
		}
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

func WithHeader(value HeaderSection) BuilderOption {
	return func(builder *Builder) {
		builder.headerMode = HeaderModeEnabled
		headerCopy := value
		builder.header = &headerCopy
		builder.syncValuesFromHeader()
	}
}

func WithHeaderOptions(options ...HeaderOption) BuilderOption {
	return func(builder *Builder) {
		header := builder.ensureHeaderExplicit()
		for _, option := range options {
			option(header)
		}
		builder.syncValuesFromHeader()
	}
}

func NewBuilder(themeImpl theme.Theme, options ...BuilderOption) *Builder {
	if themeImpl == nil {
		panic("myrtle: theme is required")
	}

	builder := &Builder{
		headerMode: HeaderModeAuto,
		theme:      themeImpl,
		registry:   NewRegistry(),
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
		preheader:  builder.preheader,
		headerMode: builder.headerMode,
		values:     builder.values,
		blocks:     append([]Block(nil), builder.blocks...),
		theme:      builder.theme,
		registry:   builder.registry,
	}
}
