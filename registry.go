package myrtle

import (
	"errors"
	"fmt"

	"github.com/gzuidhof/myrtle/theme"
)

var (
	ErrRendererAlreadyRegistered = errors.New("myrtle: renderer already registered")
	ErrRendererNotFound          = errors.New("myrtle: renderer not found")
)

type Registry struct {
	renderers map[theme.BlockKind]customRenderer
}

type customRenderer struct {
	renderHTML func(any, theme.Values) (string, error)
	renderText func(any, RenderContext) (string, error)
}

func NewRegistry() *Registry {
	return &Registry{
		renderers: map[theme.BlockKind]customRenderer{},
	}
}

func Register[T any](
	registry *Registry,
	kind theme.BlockKind,
	renderHTML func(T, theme.Values) (string, error),
	renderText func(T, RenderContext) (string, error),
) error {
	if registry == nil {
		return errors.New("myrtle: registry cannot be nil")
	}

	if renderHTML == nil {
		panic("myrtle: renderHTML cannot be nil")
	}

	if renderText == nil {
		panic("myrtle: renderText cannot be nil")
	}

	if _, exists := registry.renderers[kind]; exists {
		return fmt.Errorf("%w: %s", ErrRendererAlreadyRegistered, kind)
	}

	renderer := buildCustomRenderer(kind, renderHTML, renderText)

	registry.renderers[kind] = renderer
	return nil
}

func NewCustomBlock[T any](
	kind theme.BlockKind,
	data T,
	renderHTML func(T, theme.Values) (string, error),
	renderText func(T, RenderContext) (string, error),
) Block {
	renderer := buildCustomRenderer(kind, renderHTML, renderText)

	return customBlock[T]{
		kind:     kind,
		data:     data,
		renderer: renderer,
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

func CreateBlock[T any](registry *Registry, kind string, data T) (Block, error) {
	if registry == nil {
		return nil, errors.New("myrtle: registry cannot be nil")
	}

	return registry.createBlock(theme.BlockKind(kind), data)
}

func (registry *Registry) createBlock(kind theme.BlockKind, data any) (Block, error) {
	renderer, exists := registry.renderers[kind]
	if !exists {
		return nil, fmt.Errorf("%w: %s", ErrRendererNotFound, kind)
	}

	return customBlock[any]{
		kind:     kind,
		data:     data,
		renderer: renderer,
	}, nil
}

type customBlock[T any] struct {
	kind     theme.BlockKind
	data     T
	renderer customRenderer
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
