package myrtle

import "github.com/gzuidhof/myrtle/theme"

func (group *Group) Kind() theme.BlockKind {
	return theme.BlockKind("group")
}

func (group *Group) TemplateData() any {
	return group
}

func (group *Group) RenderText(context RenderContext) (string, error) {
	if group == nil {
		return "", nil
	}

	return renderColumnText(group.Blocks(), context)
}
