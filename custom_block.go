package myrtle

import (
	"fmt"

	"github.com/gzuidhof/myrtle/theme"
)

// customRenderer stores normalized renderer callbacks for dynamic custom blocks.
type customRenderer struct {
	renderHTML func(any, theme.Values) (string, error)
	renderText func(any, RenderContext) (string, error)
}

// NewCustomBlock creates a block backed by caller-provided HTML and plain-text renderers.
//
// Use this when you need to add a one-off or app-specific block without modifying Myrtle's
// built-in block types or theme template sets. The same typed `data` value is passed to both
// renderers, so you can keep custom rendering logic cohesive while still supporting text-only
// clients.
//
// Type parameter `T` is the custom payload type used by both callbacks. Myrtle wraps the
// callbacks in a runtime type-checking adapter so the block fails with a clear error if a
// renderer ever receives an unexpected payload type.
//
// Both `renderHTML` and `renderText` are required and this function panics if either callback
// is nil. Errors returned by either callback are propagated during email rendering.
//
// The resulting block uses the default layout spec (normal content inset behavior). See
// NewCustomBlockWithLayoutSpec if you need to customize layout metadata.
func NewCustomBlock[T any](
	kind theme.BlockKind,
	data T,
	renderHTML func(T, theme.Values) (string, error),
	renderText func(T, RenderContext) (string, error),
) Block {
	return NewCustomBlockWithLayoutSpec(kind, data, defaultLayoutSpec(), renderHTML, renderText)
}

// NewCustomBlockWithLayoutSpec creates a custom block with explicit layout metadata.
//
// It behaves like NewCustomBlock, but lets callers control inset behavior used by theme
// layouts (for example InsetModeNone for full-width content or InsetModeCustom with
// CustomInset for per-block horizontal padding).
func NewCustomBlockWithLayoutSpec[T any](
	kind theme.BlockKind,
	data T,
	layoutSpec LayoutSpec,
	renderHTML func(T, theme.Values) (string, error),
	renderText func(T, RenderContext) (string, error),
) Block {
	renderer := buildCustomRenderer(kind, renderHTML, renderText)

	return customBlock[T]{
		kind:       kind,
		data:       data,
		renderer:   renderer,
		layoutSpec: normalizedLayoutSpec(layoutSpec),
	}
}

func buildCustomRenderer[T any](
	kind theme.BlockKind,
	renderHTML func(T, theme.Values) (string, error),
	renderText func(T, RenderContext) (string, error),
) customRenderer {
	if renderHTML == nil {
		panic("myrtle: renderHTML cannot be nil")
	}

	if renderText == nil {
		panic("myrtle: renderText cannot be nil")
	}

	expected := *new(T)
	convert := func(value any) (T, error) {
		typed, ok := value.(T)
		if !ok {
			return typed, fmt.Errorf("myrtle: renderer %q expected %T got %T", kind, expected, value)
		}

		return typed, nil
	}

	return customRenderer{
		renderHTML: func(value any, values theme.Values) (string, error) {
			typed, err := convert(value)
			if err != nil {
				return "", err
			}
			return renderHTML(typed, values)
		},
		renderText: func(value any, context RenderContext) (string, error) {
			typed, err := convert(value)
			if err != nil {
				return "", err
			}
			return renderText(typed, context)
		},
	}
}

// customBlock is a runtime block wrapper backed by a registered custom renderer.
type customBlock[T any] struct {
	kind       theme.BlockKind
	data       T
	renderer   customRenderer
	layoutSpec LayoutSpec
}

func (block customBlock[T]) Kind() theme.BlockKind {
	return block.kind
}

func (block customBlock[T]) TemplateData() any {
	return block.data
}

func (block customBlock[T]) RenderText(context RenderContext) (string, error) {
	return block.renderer.renderText(block.data, context)
}

func (block customBlock[T]) RenderHTML(values theme.Values) (string, error) {
	return block.renderer.renderHTML(block.data, values)
}

func (block customBlock[T]) LayoutSpec() LayoutSpec { return normalizedLayoutSpec(block.layoutSpec) }
