package themerender

import (
	"bytes"
	"html/template"
	"strings"

	"github.com/gzuidhof/myrtle"
	"github.com/gzuidhof/myrtle/theme"
	"github.com/yuin/goldmark"
)

func renderFreeMarkdownBlock(_ *template.Template, view theme.BlockView) (string, bool, error) {
	markdownBlock, ok := view.Data.(myrtle.FreeMarkdownBlock)
	if !ok {
		return "", false, nil
	}

	var markdownOutput bytes.Buffer
	if err := goldmark.Convert([]byte(markdownBlock.Markdown), &markdownOutput); err != nil {
		return "", false, err
	}

	content := strings.TrimSpace(markdownOutput.String())
	if content == "" {
		return "", true, nil
	}

	return `<div style="color:` + view.Values.Styles.ColorText + `;">` + content + `</div>`, true, nil
}

func renderMarkdownHTML(value string) (template.HTML, error) {
	value = strings.TrimSpace(value)
	if value == "" {
		return "", nil
	}

	var output bytes.Buffer
	if err := goldmark.Convert([]byte(value), &output); err != nil {
		return "", err
	}

	return template.HTML(strings.TrimSpace(output.String())), nil
}

func renderMarkdownInline(value string) (template.HTML, error) {
	htmlValue, err := renderMarkdownHTML(value)
	if err != nil {
		return "", err
	}

	normalized := normalizeTokenPreserveCase(string(htmlValue))
	if strings.HasPrefix(normalized, "<p>") && strings.HasSuffix(normalized, "</p>") && strings.Count(normalized, "<p>") == 1 && strings.Count(normalized, "</p>") == 1 {
		normalized = strings.TrimPrefix(normalized, "<p>")
		normalized = strings.TrimSuffix(normalized, "</p>")
		normalized = normalizeTokenPreserveCase(normalized)
	}

	return template.HTML(normalized), nil
}

func normalizeTokenPreserveCase(value string) string {
	return strings.TrimSpace(value)
}
