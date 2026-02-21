package myrtle

import (
	"fmt"
	"strings"

	"github.com/gzuidhof/myrtle/theme"
)

type HeadingBlock struct {
	Text  string
	Level int
}

func (block HeadingBlock) Kind() theme.BlockKind {
	return theme.BlockKindHeading
}

func (block HeadingBlock) TemplateData() any {
	return block
}

func (block HeadingBlock) RenderMarkdown(_ RenderContext) (string, error) {
	text := strings.TrimSpace(block.Text)
	if text == "" {
		return "", nil
	}

	level := block.Level
	if level < 1 {
		level = 1
	}
	if level > 6 {
		level = 6
	}

	return strings.Repeat("#", level) + " " + text, nil
}

type SpacerBlock struct {
	Size int
}

func (block SpacerBlock) Kind() theme.BlockKind {
	return theme.BlockKindSpacer
}

func (block SpacerBlock) TemplateData() any {
	return block
}

func (block SpacerBlock) RenderMarkdown(_ RenderContext) (string, error) {
	return "", nil
}

type ListBlock struct {
	Items   []string
	Ordered bool
}

func (block ListBlock) Kind() theme.BlockKind {
	return theme.BlockKindList
}

func (block ListBlock) TemplateData() any {
	return block
}

func (block ListBlock) RenderMarkdown(_ RenderContext) (string, error) {
	parts := make([]string, 0, len(block.Items))
	for _, item := range block.Items {
		value := strings.TrimSpace(item)
		if value == "" {
			continue
		}
		if block.Ordered {
			parts = append(parts, fmt.Sprintf("%d. %s", len(parts)+1, value))
			continue
		}
		parts = append(parts, "- "+value)
	}

	return strings.Join(parts, "\n"), nil
}
