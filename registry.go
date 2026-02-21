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
	renderMarkdown func(any, RenderContext) (string, error)
}

func NewRegistry() *Registry {
	return &Registry{
		renderers: map[theme.BlockKind]customRenderer{},
	}
}

func Register[T any](
	registry *Registry,
	kind theme.BlockKind,
	renderMarkdown func(T, RenderContext) (string, error),
) error {
	if registry == nil {
		return errors.New("myrtle: registry cannot be nil")
	}

	if _, exists := registry.renderers[kind]; exists {
		return fmt.Errorf("%w: %s", ErrRendererAlreadyRegistered, kind)
	}

	renderer := customRenderer{
		renderMarkdown: func(value any, context RenderContext) (string, error) {
			typed, ok := value.(T)
			if !ok {
				return "", fmt.Errorf("myrtle: renderer %q expected %T got %T", kind, *new(T), value)
			}
			return renderMarkdown(typed, context)
		},
	}

	registry.renderers[kind] = renderer
	return nil
}

func Create[T any](registry *Registry, kind string, data T) (Block, error) {
	if registry == nil {
		return nil, errors.New("myrtle: registry cannot be nil")
	}

	return registry.create(theme.BlockKind(kind), data)
}

func (registry *Registry) create(kind theme.BlockKind, data any) (Block, error) {
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

func (block customBlock[T]) RenderMarkdown(context RenderContext) (string, error) {
	return block.renderer.renderMarkdown(block.data, context)
}
