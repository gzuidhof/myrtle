package themerender

import (
	"fmt"
	"html/template"
	"strings"

	"github.com/gzuidhof/myrtle"
	"github.com/gzuidhof/myrtle/theme"
)

func DefaultBlockRenderHandlers() map[theme.BlockKind]BlockRenderHandler {
	return map[theme.BlockKind]BlockRenderHandler{
		theme.BlockKindText:               renderTextBlock,
		theme.BlockKindHeading:            renderHeadingBlock,
		theme.BlockKindSpacer:             renderSpacerBlock,
		theme.BlockKindList:               renderListBlock,
		theme.BlockKindKeyValue:           renderKeyValueBlock,
		theme.BlockKindHorizontalBarChart: renderHorizontalBarChartBlock,
		theme.BlockKindVerticalBarChart:   renderVerticalBarChartBlock,
		theme.BlockKindSparkline:          renderSparklineBlock,
		theme.BlockKindStackedBar:         renderStackedBarBlock,
		theme.BlockKindProgress:           renderProgressBlock,
		theme.BlockKindDistribution:       renderDistributionBlock,
		theme.BlockKindTimeline:           renderTimelineBlock,
		theme.BlockKindStatsRow:           renderStatsRowBlock,
		theme.BlockKindBadge:              renderBadgeBlock,
		theme.BlockKindSummaryCard:        renderSummaryCardBlock,
		theme.BlockKindAttachment:         renderAttachmentBlock,
		theme.BlockKindHero:               renderHeroBlock,
		theme.BlockKindFooterLinks:        renderFooterLinksBlock,
		theme.BlockKindPriceSummary:       renderPriceSummaryBlock,
		theme.BlockKindEmptyState:         renderEmptyStateBlock,
		theme.BlockKindQuote:              renderQuoteBlock,
		theme.BlockKindCallout:            renderCalloutBlock,
		theme.BlockKindMessage:            renderMessageBlock,
		theme.BlockKindMessageDigest:      renderMessageDigestBlock,
		theme.BlockKindLegal:              renderLegalBlock,
		theme.BlockKindColumns:            renderColumnsBlock,
		theme.BlockKindPanel:              renderPanelBlock,
		theme.BlockKindGrid:               renderGridBlock,
		theme.BlockKindCardList:           renderCardListBlock,
		theme.BlockKindButton:             renderButtonBlock,
		theme.BlockKindButtonGroup:        renderButtonGroupBlock,
		theme.BlockKindDivider:            renderDividerBlock,
		theme.BlockKindImage:              renderImageBlock,
		theme.BlockKindTable:              renderTableBlock,
		theme.BlockKindVerificationCode:   renderVerificationCodeBlock,
		theme.BlockKindTiles:              renderTilesBlock,
		theme.BlockKindFreeMarkdown:       renderFreeMarkdownBlock,
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

	normalized := spacerBlock.TemplateData().(myrtle.SpacerBlock)

	result, err := ExecuteTemplate(templates, "block.spacer.html.tmpl", struct {
		Block myrtle.SpacerBlock
	}{Block: normalized})
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

func renderHorizontalBarChartBlock(templates *template.Template, view theme.BlockView) (string, bool, error) {
	horizontalBarChartBlock, ok := view.Data.(myrtle.HorizontalBarChartBlock)
	if !ok {
		return "", false, nil
	}

	result, err := ExecuteTemplate(templates, "block.bar_chart.html.tmpl", struct {
		Block  myrtle.HorizontalBarChartBlock
		Values theme.Values
	}{Block: horizontalBarChartBlock, Values: view.Values})
	if err != nil {
		return "", false, err
	}

	return result, true, nil
}

func renderVerticalBarChartBlock(templates *template.Template, view theme.BlockView) (string, bool, error) {
	verticalBarChartData, ok := view.Data.(myrtle.VerticalBarChartTemplateData)
	if !ok {
		verticalBarChartBlock, blockOK := view.Data.(myrtle.VerticalBarChartBlock)
		if !blockOK {
			return "", false, nil
		}

		verticalBarChartData = verticalBarChartBlock.TemplateData().(myrtle.VerticalBarChartTemplateData)
	}

	result, err := ExecuteTemplate(templates, "block.vertical_bar_chart.html.tmpl", struct {
		Block  myrtle.VerticalBarChartTemplateData
		Values theme.Values
	}{Block: verticalBarChartData, Values: view.Values})
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

func renderMessageBlock(templates *template.Template, view theme.BlockView) (string, bool, error) {
	messageBlock, ok := view.Data.(myrtle.MessageBlock)
	if !ok {
		return "", false, nil
	}
	normalized := messageBlock.TemplateData().(myrtle.MessageBlock)

	subjectHTML, err := renderMarkdownInline(normalized.Subject)
	if err != nil {
		return "", false, err
	}
	previewSource := normalized.Preview
	if normalized.PreviewMarkdown != "" {
		previewSource = normalized.PreviewMarkdown
	}

	previewHTML, err := renderMarkdownInline(previewSource)
	if err != nil {
		return "", false, err
	}

	result, err := ExecuteTemplate(templates, "block.message.html.tmpl", struct {
		Block       myrtle.MessageBlock
		SubjectHTML template.HTML
		PreviewHTML template.HTML
		MetaLine    string
		JumpURL     string
		JumpLabel   string
		Values      theme.Values
	}{
		Block:       normalized,
		SubjectHTML: subjectHTML,
		PreviewHTML: previewHTML,
		MetaLine:    messageMetaLine(normalized),
		JumpURL:     normalized.URL,
		JumpLabel:   "Jump to message",
		Values:      view.Values,
	})
	if err != nil {
		return "", false, err
	}

	return result, true, nil
}

func renderMessageDigestBlock(templates *template.Template, view theme.BlockView) (string, bool, error) {
	digestBlock, ok := view.Data.(myrtle.MessageDigestBlock)
	if !ok {
		return "", false, nil
	}

	normalized := digestBlock.TemplateData().(myrtle.MessageDigestBlock)
	type messageDigestItemView struct {
		Block       myrtle.MessageBlock
		SubjectHTML template.HTML
		PreviewHTML template.HTML
		MetaLine    string
	}

	items := make([]messageDigestItemView, 0, len(normalized.Messages))
	hasAvatar := false
	for _, message := range normalized.Messages {
		if message.AvatarURL != "" {
			hasAvatar = true
		}

		subjectHTML, err := renderMarkdownInline(message.Subject)
		if err != nil {
			return "", false, err
		}
		previewSource := message.Preview
		if message.PreviewMarkdown != "" {
			previewSource = message.PreviewMarkdown
		}

		previewHTML, err := renderMarkdownInline(previewSource)
		if err != nil {
			return "", false, err
		}

		items = append(items, messageDigestItemView{
			Block:       message,
			SubjectHTML: subjectHTML,
			PreviewHTML: previewHTML,
			MetaLine:    messageMetaLine(message),
		})
	}

	subtitleHTML, err := renderMarkdownHTML(normalized.Subtitle)
	if err != nil {
		return "", false, err
	}
	footerHTML, err := renderMarkdownHTML(normalized.Footer)
	if err != nil {
		return "", false, err
	}

	result, err := ExecuteTemplate(templates, "block.message_digest.html.tmpl", struct {
		Block        myrtle.MessageDigestBlock
		Items        []messageDigestItemView
		HasAvatar    bool
		SubtitleHTML template.HTML
		FooterHTML   template.HTML
		Values       theme.Values
	}{
		Block:        normalized,
		Items:        items,
		HasAvatar:    hasAvatar,
		SubtitleHTML: subtitleHTML,
		FooterHTML:   footerHTML,
		Values:       view.Values,
	})
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

func renderPanelBlock(templates *template.Template, view theme.BlockView) (string, bool, error) {
	sectionBlock, ok := view.Data.(myrtle.PanelBlock)
	if !ok {
		return "", false, nil
	}

	bodyHTML, err := renderNestedBlocksHTML(templates, sectionBlock.Blocks, view.Values)
	if err != nil {
		return "", false, err
	}

	result, err := ExecuteTemplate(templates, "block.panel.html.tmpl", struct {
		Block    myrtle.PanelBlock
		BodyHTML string
		Values   theme.Values
	}{Block: sectionBlock, BodyHTML: bodyHTML, Values: view.Values})
	if err != nil {
		return "", false, err
	}

	return result, true, nil
}

func renderGridBlock(templates *template.Template, view theme.BlockView) (string, bool, error) {
	gridBlock, ok := view.Data.(myrtle.GridBlock)
	if !ok {
		return "", false, nil
	}

	normalized := gridBlock.TemplateData().(myrtle.GridBlock)
	if len(normalized.Items) == 0 {
		return "", true, nil
	}

	cellHTML := make([]string, 0, len(normalized.Items))
	for _, item := range normalized.Items {
		if item.Content == nil {
			continue
		}

		html, err := renderNestedBlocksHTML(templates, []myrtle.Block{item.Content}, view.Values)
		if err != nil {
			return "", false, err
		}
		cellHTML = append(cellHTML, html)
	}
	if len(cellHTML) == 0 {
		return "", true, nil
	}

	rows := make([][]string, 0, (len(cellHTML)+normalized.Columns-1)/normalized.Columns)
	for index := 0; index < len(cellHTML); index += normalized.Columns {
		end := index + normalized.Columns
		if end > len(cellHTML) {
			end = len(cellHTML)
		}

		row := append([]string(nil), cellHTML[index:end]...)
		for len(row) < normalized.Columns {
			row = append(row, "")
		}
		rows = append(rows, row)
	}

	result, err := ExecuteTemplate(templates, "block.grid.html.tmpl", struct {
		Block       myrtle.GridBlock
		Rows        [][]string
		ColumnWidth int
		Values      theme.Values
	}{Block: normalized, Rows: rows, ColumnWidth: 100 / normalized.Columns, Values: view.Values})
	if err != nil {
		return "", false, err
	}

	return result, true, nil
}

func renderCardListBlock(templates *template.Template, view theme.BlockView) (string, bool, error) {
	cardListBlock, ok := view.Data.(myrtle.CardListBlock)
	if !ok {
		return "", false, nil
	}

	normalized := cardListBlock.TemplateData().(myrtle.CardListBlock)
	if len(normalized.Cards) == 0 {
		return "", true, nil
	}

	rows := make([][]myrtle.CardItem, 0, (len(normalized.Cards)+normalized.Columns-1)/normalized.Columns)
	for index := 0; index < len(normalized.Cards); index += normalized.Columns {
		end := index + normalized.Columns
		if end > len(normalized.Cards) {
			end = len(normalized.Cards)
		}

		row := append([]myrtle.CardItem(nil), normalized.Cards[index:end]...)
		rows = append(rows, row)
	}

	result, err := ExecuteTemplate(templates, "block.card_list.html.tmpl", struct {
		Block       myrtle.CardListBlock
		Rows        [][]myrtle.CardItem
		ColumnWidth int
		Values      theme.Values
	}{Block: normalized, Rows: rows, ColumnWidth: 100 / normalized.Columns, Values: view.Values})
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
	dividerBlock, ok := view.Data.(myrtle.DividerBlock)
	if !ok {
		return "", false, nil
	}

	normalized := dividerBlock.TemplateData().(myrtle.DividerBlock)

	result, err := ExecuteTemplate(templates, "block.divider.html.tmpl", struct {
		Block  myrtle.DividerBlock
		Values theme.Values
	}{Block: normalized, Values: view.Values})
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

	normalized := imageBlock.TemplateData().(myrtle.ImageBlock)

	result, err := ExecuteTemplate(templates, "block.image.html.tmpl", struct {
		Block  myrtle.ImageBlock
		Values theme.Values
	}{Block: normalized, Values: view.Values})
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

	columnCount := len(tableBlock.Columns)
	for _, row := range tableBlock.Rows {
		if len(row) > columnCount {
			columnCount = len(row)
		}
	}
	if tableBlock.HasLegendSwatches {
		columnCount++
	}
	if columnCount <= 0 {
		columnCount = 1
	}

	result, err := ExecuteTemplate(templates, "block.table.html.tmpl", struct {
		Block       myrtle.TableBlock
		ColumnCount int
		Values      theme.Values
	}{Block: tableBlock, ColumnCount: columnCount, Values: view.Values})
	if err != nil {
		return "", false, err
	}

	return result, true, nil
}

func renderVerificationCodeBlock(templates *template.Template, view theme.BlockView) (string, bool, error) {
	codeBlock, ok := view.Data.(myrtle.VerificationCodeBlock)
	if !ok {
		return "", false, nil
	}

	result, err := ExecuteTemplate(templates, "block.verification_code.html.tmpl", struct {
		Block  myrtle.VerificationCodeBlock
		Values theme.Values
	}{Block: codeBlock, Values: view.Values})
	if err != nil {
		return "", false, err
	}

	return result, true, nil
}

func renderTilesBlock(templates *template.Template, view theme.BlockView) (string, bool, error) {
	block, ok := view.Data.(myrtle.TilesBlock)
	if !ok {
		return "", false, nil
	}

	normalized := block.TemplateData().(myrtle.TilesBlock)
	if len(normalized.Entries) == 0 {
		return "", true, nil
	}

	rows := make([][]myrtle.TileEntry, 0, (len(normalized.Entries)+normalized.Columns-1)/normalized.Columns)
	for index := 0; index < len(normalized.Entries); index += normalized.Columns {
		end := index + normalized.Columns
		if end > len(normalized.Entries) {
			end = len(normalized.Entries)
		}

		row := append([]myrtle.TileEntry(nil), normalized.Entries[index:end]...)
		rows = append(rows, row)
	}

	columnWidth := 100 / normalized.Columns

	result, err := ExecuteTemplate(templates, "block.tiles.html.tmpl", struct {
		Block       myrtle.TilesBlock
		Rows        [][]myrtle.TileEntry
		ColumnWidth int
		Values      theme.Values
	}{
		Block:       normalized,
		Rows:        rows,
		ColumnWidth: columnWidth,
		Values:      view.Values,
	})
	if err != nil {
		return "", false, err
	}

	return result, true, nil
}
