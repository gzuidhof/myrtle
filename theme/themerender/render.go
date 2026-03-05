package themerender

import (
	"bytes"
	"embed"
	"fmt"
	"html/template"
	"io/fs"
	"strings"
	texttemplate "text/template"

	"github.com/gzuidhof/myrtle"
	"github.com/gzuidhof/myrtle/theme"
)

// BlockRenderHandler renders one block view and reports whether the handler accepted the input.
type BlockRenderHandler func(templates *template.Template, view theme.BlockView) (string, bool, error)

var templateLessBlockKinds = map[theme.BlockKind]struct{}{
	theme.BlockKindFreeMarkdown: {},
}

//go:embed shared/*.tmpl
var sharedTemplatesFS embed.FS

func SharedLayoutTemplateFiles() []string {
	return []string{"shared/header_footer.html.tmpl"}
}

func blockTemplateFileForKind(kind theme.BlockKind) string {
	return fmt.Sprintf("block.%s.html.tmpl", string(kind))
}

func SharedBlockTemplateFiles() []string {
	files := make([]string, 0, len(theme.AllBlockKinds))
	for _, kind := range theme.AllBlockKinds {
		if _, skip := templateLessBlockKinds[kind]; skip {
			continue
		}

		files = append(files, blockTemplateFileForKind(kind))
	}

	return files
}

func SharedBlockTemplateFilesAvailableInFS(filesystem fs.FS) []string {
	allFiles := SharedBlockTemplateFiles()
	files := make([]string, 0, len(allFiles))
	for _, file := range allFiles {
		if _, err := fs.Stat(filesystem, file); err != nil {
			continue
		}

		files = append(files, file)
	}

	return files
}

func SharedBlockTemplateFilesExcludingKinds(kinds []theme.BlockKind) []string {
	excludedFiles := make(map[string]struct{}, len(kinds))
	for _, kind := range kinds {
		excludedFiles[blockTemplateFileForKind(kind)] = struct{}{}
	}

	allFiles := SharedBlockTemplateFiles()
	files := make([]string, 0, len(allFiles))
	for _, file := range allFiles {
		if _, excluded := excludedFiles[file]; excluded {
			continue
		}

		files = append(files, file)
	}

	return files
}

func DefaultBlockRenderHandlersExcludingKinds(kinds []theme.BlockKind) map[theme.BlockKind]BlockRenderHandler {
	handlers := DefaultBlockRenderHandlers()
	for _, kind := range kinds {
		delete(handlers, kind)
	}

	return handlers
}

func DefaultBlockRenderHandlersForTemplateFiles(templateFiles []string) map[theme.BlockKind]BlockRenderHandler {
	handlers := DefaultBlockRenderHandlers()
	available := make(map[string]struct{}, len(templateFiles))
	for _, file := range templateFiles {
		available[file] = struct{}{}
	}

	for kind := range handlers {
		file := blockTemplateFileForKind(kind)
		if _, ok := available[file]; ok {
			continue
		}

		delete(handlers, kind)
	}

	return handlers
}

func ParseHTMLTemplates(name string, filesystem fs.FS, files ...string) *template.Template {
	return template.Must(template.New(name).Funcs(template.FuncMap{
		"safe": func(value string) template.HTML {
			return template.HTML(value)
		},
		"safeCSS": func(value string) template.CSS {
			return template.CSS(value)
		},
		"isNumericLike": func(value string) bool {
			return isNumericLike(value)
		},
		"isOdd": func(value int) bool {
			return value%2 == 1
		},
		"isLastRow": func(index, length int) bool {
			return index == length-1
		},
		"isDiscountLike": func(label, value string) bool {
			return isDiscountLike(label, value)
		},
		"physicalAlign": func(alignment any, values theme.Values) string {
			return physicalAlign(alignment, values)
		},
		"physicalSide": func(side any, values theme.Values) string {
			return physicalSide(side, values)
		},
		"toneColor": func(tone any, values theme.Values, role string) string {
			return toneColor(tone, values, role)
		},
		"isRTL": func(values theme.Values) bool {
			return isRTL(values)
		},
		"halfFloor": func(value int) int {
			if value <= 0 {
				return 0
			}

			return value / 2
		},
		"halfCeil": func(value int) int {
			if value <= 0 {
				return 0
			}

			return value - (value / 2)
		},
		"preheaderFiller": func(repeat int) template.HTML {
			if repeat <= 0 {
				return ""
			}

			return template.HTML(strings.Repeat("&nbsp;&zwnj;", repeat))
		},
		"tableColumnIsNumeric": func(rows [][]string, columnIndex int) bool {
			if columnIndex < 0 {
				return false
			}

			hasNumeric := false
			for _, row := range rows {
				if columnIndex >= len(row) {
					continue
				}

				cell := strings.TrimSpace(row[columnIndex])
				if cell == "" {
					continue
				}
				if !isNumericLike(cell) {
					return false
				}

				hasNumeric = true
			}

			return hasNumeric
		},
		"keepRightBorderForGaplessNone": func(row []string, index int) bool {
			if index < 0 || index >= len(row) {
				return false
			}

			trimmedRow := make([]string, len(row))
			for i, cell := range row {
				trimmedRow[i] = strings.TrimSpace(cell)
			}

			if trimmedRow[index] == "" {
				return false
			}

			hasEmptyCell := false
			for _, cell := range trimmedRow {
				if cell == "" {
					hasEmptyCell = true
					break
				}
			}
			if !hasEmptyCell {
				return false
			}

			for next := index + 1; next < len(trimmedRow); next++ {
				if trimmedRow[next] != "" {
					return false
				}
			}

			return true
		},
	}).ParseFS(filesystem, files...))
}

func ParseHTMLTemplatesWithShared(name string, filesystem fs.FS, files ...string) *template.Template {
	templates := ParseHTMLTemplates(name, sharedTemplatesFS, SharedLayoutTemplateFiles()...)
	return template.Must(templates.ParseFS(filesystem, files...))
}

func ExecuteTemplate(templates *template.Template, name string, data any) (string, error) {
	var output bytes.Buffer
	if err := templates.ExecuteTemplate(&output, name, data); err != nil {
		return "", err
	}

	return output.String(), nil
}

func ExecuteTextTemplate(templates *texttemplate.Template, name string, data any) (string, error) {
	var output bytes.Buffer
	if err := templates.ExecuteTemplate(&output, name, data); err != nil {
		return "", err
	}

	return strings.TrimSpace(output.String()), nil
}

func RenderBlockHTML(
	templates *template.Template,
	view theme.BlockView,
	fallback theme.Theme,
) (string, bool, error) {
	return RenderBlockHTMLWithHandlers(templates, view, DefaultBlockRenderHandlers(), fallback)
}

func RenderBlockHTMLWithHandlers(
	templates *template.Template,
	view theme.BlockView,
	handlers map[theme.BlockKind]BlockRenderHandler,
	fallback theme.Theme,
) (string, bool, error) {
	switch data := view.Data.(type) {
	case *myrtle.Group:
		html, err := renderGroupDataHTML(templates, data, view.Values, handlers, fallback, map[*myrtle.Group]struct{}{})
		if err != nil {
			return "", false, err
		}

		return html, true, nil
	case myrtle.Group:
		copyData := data
		html, err := renderGroupDataHTML(templates, &copyData, view.Values, handlers, fallback, map[*myrtle.Group]struct{}{})
		if err != nil {
			return "", false, err
		}

		return html, true, nil
	}

	handler, ok := handlers[view.Kind]
	if !ok {
		return renderFallback(fallback, view)
	}

	result, handled, err := handler(templates, view)
	if err != nil {
		return "", false, err
	}
	if handled {
		return result, true, nil
	}

	return renderFallback(fallback, view)
}

func renderGroupDataHTML(
	templates *template.Template,
	group *myrtle.Group,
	values theme.Values,
	handlers map[theme.BlockKind]BlockRenderHandler,
	fallback theme.Theme,
	seen map[*myrtle.Group]struct{},
) (string, error) {
	if group == nil {
		return "", nil
	}

	if _, exists := seen[group]; exists {
		return "", fmt.Errorf("myrtle: group contains cyclic reference")
	}
	seen[group] = struct{}{}
	defer delete(seen, group)

	parts := make([]string, 0, len(group.Blocks()))
	for _, block := range group.Blocks() {
		if block == nil {
			continue
		}

		html, ok, err := RenderBlockHTMLWithHandlers(templates, theme.BlockView{
			Kind:   block.Kind(),
			Data:   block.TemplateData(),
			Values: values,
		}, handlers, fallback)
		if err != nil {
			return "", err
		}
		if !ok {
			return "", fmt.Errorf("myrtle: group child cannot render kind %s", block.Kind())
		}

		parts = append(parts, html)
	}

	return strings.Join(parts, ""), nil
}
