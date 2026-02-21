package themerender

import (
	"bytes"
	"fmt"
	"html/template"
	"io/fs"
	"strconv"
	"strings"
	texttemplate "text/template"

	"github.com/gzuidhof/myrtle"
	"github.com/gzuidhof/myrtle/theme"
	"github.com/yuin/goldmark"
)

type BlockRenderHandler func(templates *template.Template, view theme.BlockView) (string, bool, error)

var templateLessBlockKinds = map[theme.BlockKind]struct{}{
	theme.BlockKindFreeMarkdown: {},
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
	}).ParseFS(filesystem, files...))
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

func DefaultBlockRenderHandlers() map[theme.BlockKind]BlockRenderHandler {
	return map[theme.BlockKind]BlockRenderHandler{
		theme.BlockKindText:         renderTextBlock,
		theme.BlockKindHeading:      renderHeadingBlock,
		theme.BlockKindSpacer:       renderSpacerBlock,
		theme.BlockKindList:         renderListBlock,
		theme.BlockKindKeyValue:     renderKeyValueBlock,
		theme.BlockKindBarChart:     renderBarChartBlock,
		theme.BlockKindSparkline:    renderSparklineBlock,
		theme.BlockKindStackedBar:   renderStackedBarBlock,
		theme.BlockKindProgress:     renderProgressBlock,
		theme.BlockKindDistribution: renderDistributionBlock,
		theme.BlockKindTimeline:     renderTimelineBlock,
		theme.BlockKindStatsRow:     renderStatsRowBlock,
		theme.BlockKindBadge:        renderBadgeBlock,
		theme.BlockKindSummaryCard:  renderSummaryCardBlock,
		theme.BlockKindAttachment:   renderAttachmentBlock,
		theme.BlockKindHero:         renderHeroBlock,
		theme.BlockKindFooterLinks:  renderFooterLinksBlock,
		theme.BlockKindPriceSummary: renderPriceSummaryBlock,
		theme.BlockKindEmptyState:   renderEmptyStateBlock,
		theme.BlockKindQuote:        renderQuoteBlock,
		theme.BlockKindCallout:      renderCalloutBlock,
		theme.BlockKindLegal:        renderLegalBlock,
		theme.BlockKindColumns:      renderColumnsBlock,
		theme.BlockKindButton:       renderButtonBlock,
		theme.BlockKindButtonGroup:  renderButtonGroupBlock,
		theme.BlockKindDivider:      renderDividerBlock,
		theme.BlockKindImage:        renderImageBlock,
		theme.BlockKindTable:        renderTableBlock,
		theme.BlockKindAction:       renderActionBlock,
		theme.BlockKindCode:         renderCodeBlock,
		theme.BlockKindFreeMarkdown: renderFreeMarkdownBlock,
	}
}

func renderFallback(fallback theme.Theme, view theme.BlockView) (string, bool, error) {
	if fallback == nil {
		return "", false, nil
	}

	return fallback.RenderBlockHTML(view)
}

func renderTextBlock(templates *template.Template, view theme.BlockView) (string, bool, error) {
	textBlock, ok := view.Data.(myrtle.TextBlock)
	if !ok {
		return "", false, nil
	}

	result, err := ExecuteTemplate(templates, "block.text.html.tmpl", struct {
		Block  myrtle.TextBlock
		Values theme.Values
	}{Block: textBlock, Values: view.Values})
	if err != nil {
		return "", false, err
	}

	return result, true, nil
}

func renderHeadingBlock(templates *template.Template, view theme.BlockView) (string, bool, error) {
	headingBlock, ok := view.Data.(myrtle.HeadingBlock)
	if !ok {
		return "", false, nil
	}

	result, err := ExecuteTemplate(templates, "block.heading.html.tmpl", struct {
		Block  myrtle.HeadingBlock
		Values theme.Values
	}{Block: headingBlock, Values: view.Values})
	if err != nil {
		return "", false, err
	}

	return result, true, nil
}

func renderSpacerBlock(templates *template.Template, view theme.BlockView) (string, bool, error) {
	spacerBlock, ok := view.Data.(myrtle.SpacerBlock)
	if !ok {
		return "", false, nil
	}

	result, err := ExecuteTemplate(templates, "block.spacer.html.tmpl", struct {
		Block myrtle.SpacerBlock
	}{Block: spacerBlock})
	if err != nil {
		return "", false, err
	}

	return result, true, nil
}

func renderListBlock(templates *template.Template, view theme.BlockView) (string, bool, error) {
	listBlock, ok := view.Data.(myrtle.ListBlock)
	if !ok {
		return "", false, nil
	}

	result, err := ExecuteTemplate(templates, "block.list.html.tmpl", struct {
		Block  myrtle.ListBlock
		Values theme.Values
	}{Block: listBlock, Values: view.Values})
	if err != nil {
		return "", false, err
	}

	return result, true, nil
}

func renderKeyValueBlock(templates *template.Template, view theme.BlockView) (string, bool, error) {
	keyValueBlock, ok := view.Data.(myrtle.KeyValueBlock)
	if !ok {
		return "", false, nil
	}

	result, err := ExecuteTemplate(templates, "block.key_value.html.tmpl", struct {
		Block  myrtle.KeyValueBlock
		Values theme.Values
	}{Block: keyValueBlock, Values: view.Values})
	if err != nil {
		return "", false, err
	}

	return result, true, nil
}

func renderBarChartBlock(templates *template.Template, view theme.BlockView) (string, bool, error) {
	barChartBlock, ok := view.Data.(myrtle.BarChartBlock)
	if !ok {
		return "", false, nil
	}

	result, err := ExecuteTemplate(templates, "block.bar_chart.html.tmpl", struct {
		Block  myrtle.BarChartBlock
		Values theme.Values
	}{Block: barChartBlock, Values: view.Values})
	if err != nil {
		return "", false, err
	}

	return result, true, nil
}

func renderSparklineBlock(templates *template.Template, view theme.BlockView) (string, bool, error) {
	block, ok := view.Data.(myrtle.SparklineBlock)
	if !ok {
		return "", false, nil
	}

	result, err := ExecuteTemplate(templates, "block.sparkline.html.tmpl", struct {
		Block  myrtle.SparklineBlock
		Values theme.Values
	}{Block: block, Values: view.Values})
	if err != nil {
		return "", false, err
	}

	return result, true, nil
}

func renderStackedBarBlock(templates *template.Template, view theme.BlockView) (string, bool, error) {
	block, ok := view.Data.(myrtle.StackedBarBlock)
	if !ok {
		return "", false, nil
	}

	result, err := ExecuteTemplate(templates, "block.stacked_bar.html.tmpl", struct {
		Block  myrtle.StackedBarBlock
		Values theme.Values
	}{Block: block, Values: view.Values})
	if err != nil {
		return "", false, err
	}

	return result, true, nil
}

func renderProgressBlock(templates *template.Template, view theme.BlockView) (string, bool, error) {
	block, ok := view.Data.(myrtle.ProgressBlock)
	if !ok {
		return "", false, nil
	}

	result, err := ExecuteTemplate(templates, "block.progress.html.tmpl", struct {
		Block  myrtle.ProgressBlock
		Values theme.Values
	}{Block: block, Values: view.Values})
	if err != nil {
		return "", false, err
	}

	return result, true, nil
}

func renderDistributionBlock(templates *template.Template, view theme.BlockView) (string, bool, error) {
	block, ok := view.Data.(myrtle.DistributionBlock)
	if !ok {
		return "", false, nil
	}

	result, err := ExecuteTemplate(templates, "block.distribution.html.tmpl", struct {
		Block  myrtle.DistributionBlock
		Values theme.Values
	}{Block: block, Values: view.Values})
	if err != nil {
		return "", false, err
	}

	return result, true, nil
}

func renderTimelineBlock(templates *template.Template, view theme.BlockView) (string, bool, error) {
	block, ok := view.Data.(myrtle.TimelineBlock)
	if !ok {
		return "", false, nil
	}

	result, err := ExecuteTemplate(templates, "block.timeline.html.tmpl", struct {
		Block  myrtle.TimelineBlock
		Values theme.Values
	}{Block: block, Values: view.Values})
	if err != nil {
		return "", false, err
	}

	return result, true, nil
}

func renderStatsRowBlock(templates *template.Template, view theme.BlockView) (string, bool, error) {
	block, ok := view.Data.(myrtle.StatsRowBlock)
	if !ok {
		return "", false, nil
	}

	result, err := ExecuteTemplate(templates, "block.stats_row.html.tmpl", struct {
		Block  myrtle.StatsRowBlock
		Values theme.Values
	}{Block: block, Values: view.Values})
	if err != nil {
		return "", false, err
	}

	return result, true, nil
}

func renderBadgeBlock(templates *template.Template, view theme.BlockView) (string, bool, error) {
	block, ok := view.Data.(myrtle.BadgeBlock)
	if !ok {
		return "", false, nil
	}

	result, err := ExecuteTemplate(templates, "block.badge.html.tmpl", struct {
		Block  myrtle.BadgeBlock
		Values theme.Values
	}{Block: block, Values: view.Values})
	if err != nil {
		return "", false, err
	}

	return result, true, nil
}

func renderSummaryCardBlock(templates *template.Template, view theme.BlockView) (string, bool, error) {
	block, ok := view.Data.(myrtle.SummaryCardBlock)
	if !ok {
		return "", false, nil
	}

	result, err := ExecuteTemplate(templates, "block.summary_card.html.tmpl", struct {
		Block  myrtle.SummaryCardBlock
		Values theme.Values
	}{Block: block, Values: view.Values})
	if err != nil {
		return "", false, err
	}

	return result, true, nil
}

func renderAttachmentBlock(templates *template.Template, view theme.BlockView) (string, bool, error) {
	block, ok := view.Data.(myrtle.AttachmentBlock)
	if !ok {
		return "", false, nil
	}

	result, err := ExecuteTemplate(templates, "block.attachment.html.tmpl", struct {
		Block  myrtle.AttachmentBlock
		Values theme.Values
	}{Block: block, Values: view.Values})
	if err != nil {
		return "", false, err
	}

	return result, true, nil
}

func renderHeroBlock(templates *template.Template, view theme.BlockView) (string, bool, error) {
	block, ok := view.Data.(myrtle.HeroBlock)
	if !ok {
		return "", false, nil
	}

	result, err := ExecuteTemplate(templates, "block.hero.html.tmpl", struct {
		Block  myrtle.HeroBlock
		Values theme.Values
	}{Block: block, Values: view.Values})
	if err != nil {
		return "", false, err
	}

	return result, true, nil
}

func renderFooterLinksBlock(templates *template.Template, view theme.BlockView) (string, bool, error) {
	block, ok := view.Data.(myrtle.FooterLinksBlock)
	if !ok {
		return "", false, nil
	}

	result, err := ExecuteTemplate(templates, "block.footer_links.html.tmpl", struct {
		Block  myrtle.FooterLinksBlock
		Values theme.Values
	}{Block: block, Values: view.Values})
	if err != nil {
		return "", false, err
	}

	return result, true, nil
}

func renderPriceSummaryBlock(templates *template.Template, view theme.BlockView) (string, bool, error) {
	block, ok := view.Data.(myrtle.PriceSummaryBlock)
	if !ok {
		return "", false, nil
	}

	result, err := ExecuteTemplate(templates, "block.price_summary.html.tmpl", struct {
		Block  myrtle.PriceSummaryBlock
		Values theme.Values
	}{Block: block, Values: view.Values})
	if err != nil {
		return "", false, err
	}

	return result, true, nil
}

func renderEmptyStateBlock(templates *template.Template, view theme.BlockView) (string, bool, error) {
	block, ok := view.Data.(myrtle.EmptyStateBlock)
	if !ok {
		return "", false, nil
	}

	result, err := ExecuteTemplate(templates, "block.empty_state.html.tmpl", struct {
		Block  myrtle.EmptyStateBlock
		Values theme.Values
	}{Block: block, Values: view.Values})
	if err != nil {
		return "", false, err
	}

	return result, true, nil
}

func renderQuoteBlock(templates *template.Template, view theme.BlockView) (string, bool, error) {
	quoteBlock, ok := view.Data.(myrtle.QuoteBlock)
	if !ok {
		return "", false, nil
	}

	result, err := ExecuteTemplate(templates, "block.quote.html.tmpl", struct {
		Block  myrtle.QuoteBlock
		Values theme.Values
	}{Block: quoteBlock, Values: view.Values})
	if err != nil {
		return "", false, err
	}

	return result, true, nil
}

func renderCalloutBlock(templates *template.Template, view theme.BlockView) (string, bool, error) {
	calloutBlock, ok := view.Data.(myrtle.CalloutBlock)
	if !ok {
		return "", false, nil
	}

	result, err := ExecuteTemplate(templates, "block.callout.html.tmpl", struct {
		Block  myrtle.CalloutBlock
		Values theme.Values
	}{Block: calloutBlock, Values: view.Values})
	if err != nil {
		return "", false, err
	}

	return result, true, nil
}

func renderLegalBlock(templates *template.Template, view theme.BlockView) (string, bool, error) {
	legalBlock, ok := view.Data.(myrtle.LegalBlock)
	if !ok {
		return "", false, nil
	}

	result, err := ExecuteTemplate(templates, "block.legal.html.tmpl", struct {
		Block  myrtle.LegalBlock
		Values theme.Values
	}{Block: legalBlock, Values: view.Values})
	if err != nil {
		return "", false, err
	}

	return result, true, nil
}

func renderColumnsBlock(templates *template.Template, view theme.BlockView) (string, bool, error) {
	columnsBlock, ok := view.Data.(myrtle.ColumnsBlock)
	if !ok {
		return "", false, nil
	}

	leftHTML, err := renderNestedBlocksHTML(templates, columnsBlock.Left, view.Values)
	if err != nil {
		return "", false, err
	}

	rightHTML, err := renderNestedBlocksHTML(templates, columnsBlock.Right, view.Values)
	if err != nil {
		return "", false, err
	}

	result, err := ExecuteTemplate(templates, "block.columns.html.tmpl", struct {
		Block     myrtle.ColumnsBlock
		LeftHTML  string
		RightHTML string
		Values    theme.Values
	}{Block: columnsBlock, LeftHTML: leftHTML, RightHTML: rightHTML, Values: view.Values})
	if err != nil {
		return "", false, err
	}

	return result, true, nil
}

func renderNestedBlocksHTML(templates *template.Template, blocks []myrtle.Block, values theme.Values) (string, error) {
	handlers := DefaultBlockRenderHandlers()
	parts := make([]string, 0, len(blocks))

	for _, block := range blocks {
		if block == nil {
			continue
		}

		html, ok, err := RenderBlockHTMLWithHandlers(templates, theme.BlockView{
			Kind:   block.Kind(),
			Data:   block.TemplateData(),
			Values: values,
		}, handlers, nil)
		if err != nil {
			return "", err
		}
		if !ok {
			return "", fmt.Errorf("myrtle: nested columns block cannot render kind %s", block.Kind())
		}

		parts = append(parts, html)
	}

	return strings.Join(parts, ""), nil
}

func renderButtonBlock(templates *template.Template, view theme.BlockView) (string, bool, error) {
	buttonBlock, ok := view.Data.(myrtle.ButtonBlock)
	if !ok {
		return "", false, nil
	}

	result, err := ExecuteTemplate(templates, "block.button.html.tmpl", struct {
		Block  myrtle.ButtonBlock
		Values theme.Values
	}{Block: buttonBlock, Values: view.Values})
	if err != nil {
		return "", false, err
	}

	return result, true, nil
}

func renderButtonGroupBlock(templates *template.Template, view theme.BlockView) (string, bool, error) {
	block, ok := view.Data.(myrtle.ButtonGroupBlock)
	if !ok {
		return "", false, nil
	}

	result, err := ExecuteTemplate(templates, "block.button_group.html.tmpl", struct {
		Block  myrtle.ButtonGroupBlock
		Values theme.Values
	}{Block: block, Values: view.Values})
	if err != nil {
		return "", false, err
	}

	return result, true, nil
}

func renderDividerBlock(templates *template.Template, view theme.BlockView) (string, bool, error) {
	result, err := ExecuteTemplate(templates, "block.divider.html.tmpl", struct {
		Values theme.Values
	}{Values: view.Values})
	if err != nil {
		return "", false, err
	}

	return result, true, nil
}

func renderImageBlock(templates *template.Template, view theme.BlockView) (string, bool, error) {
	imageBlock, ok := view.Data.(myrtle.ImageBlock)
	if !ok {
		return "", false, nil
	}

	result, err := ExecuteTemplate(templates, "block.image.html.tmpl", struct {
		Block  myrtle.ImageBlock
		Values theme.Values
	}{Block: imageBlock, Values: view.Values})
	if err != nil {
		return "", false, err
	}

	return result, true, nil
}

func renderTableBlock(templates *template.Template, view theme.BlockView) (string, bool, error) {
	tableBlock, ok := view.Data.(myrtle.TableBlock)
	if !ok {
		return "", false, nil
	}

	result, err := ExecuteTemplate(templates, "block.table.html.tmpl", struct {
		Block  myrtle.TableBlock
		Values theme.Values
	}{Block: tableBlock, Values: view.Values})
	if err != nil {
		return "", false, err
	}

	return result, true, nil
}

func renderActionBlock(templates *template.Template, view theme.BlockView) (string, bool, error) {
	actionBlock, ok := view.Data.(myrtle.ActionBlock)
	if !ok {
		return "", false, nil
	}

	result, err := ExecuteTemplate(templates, "block.action.html.tmpl", struct {
		Block  myrtle.ActionBlock
		Values theme.Values
	}{Block: actionBlock, Values: view.Values})
	if err != nil {
		return "", false, err
	}

	return result, true, nil
}

func renderCodeBlock(templates *template.Template, view theme.BlockView) (string, bool, error) {
	codeBlock, ok := view.Data.(myrtle.CodeBlock)
	if !ok {
		return "", false, nil
	}

	result, err := ExecuteTemplate(templates, "block.code.html.tmpl", struct {
		Block  myrtle.CodeBlock
		Values theme.Values
	}{Block: codeBlock, Values: view.Values})
	if err != nil {
		return "", false, err
	}

	return result, true, nil
}

func renderFreeMarkdownBlock(_ *template.Template, view theme.BlockView) (string, bool, error) {
	markdownBlock, ok := view.Data.(myrtle.FreeMarkdownBlock)
	if !ok {
		return "", false, nil
	}

	var markdownOutput bytes.Buffer
	if err := goldmark.Convert([]byte(markdownBlock.Markdown), &markdownOutput); err != nil {
		return "", false, err
	}

	return markdownOutput.String(), true, nil
}

func isNumericLike(value string) bool {
	cleaned := strings.TrimSpace(value)
	if cleaned == "" {
		return false
	}

	cleaned = strings.ReplaceAll(cleaned, ",", "")
	cleaned = strings.TrimPrefix(cleaned, "$")
	cleaned = strings.TrimPrefix(cleaned, "€")
	cleaned = strings.TrimPrefix(cleaned, "£")
	cleaned = strings.TrimPrefix(cleaned, "¥")
	cleaned = strings.TrimSuffix(cleaned, "%")

	if strings.HasPrefix(cleaned, "(") && strings.HasSuffix(cleaned, ")") {
		cleaned = "-" + strings.TrimSuffix(strings.TrimPrefix(cleaned, "("), ")")
	}

	if cleaned == "" {
		return false
	}

	_, err := strconv.ParseFloat(cleaned, 64)
	return err == nil
}

func isDiscountLike(label, value string) bool {
	normalizedLabel := strings.ToLower(strings.TrimSpace(label))
	if strings.Contains(normalizedLabel, "discount") || strings.Contains(normalizedLabel, "credit") {
		return true
	}

	normalizedValue := strings.TrimSpace(value)
	return strings.HasPrefix(normalizedValue, "-")
}
