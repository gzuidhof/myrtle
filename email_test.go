package myrtle_test

import (
	"fmt"
	"strings"
	"sync"
	"testing"

	"github.com/gzuidhof/myrtle"
	"github.com/gzuidhof/myrtle/theme"
	defaulttheme "github.com/gzuidhof/myrtle/theme/default"
	"github.com/gzuidhof/myrtle/theme/flat"
)

func TestBuildAndRender(t *testing.T) {
	t.Parallel()
	builder := myrtle.NewBuilder(defaulttheme.New())

	email := builder.
		Preheader("This is the preheader").
		Product("Myrtle", "https://myrtle.example").
		WithHeader(myrtle.HeaderRenderInMarkdown(true)).
		Logo("https://myrtle.example/logo.png", "").
		AddText("Hello there").
		AddButton("Open", "https://example.com").
		Build()

	html, err := email.HTML()
	if err != nil {
		t.Fatalf("HTML returned error: %v", err)
	}
	if !strings.Contains(html, "Hello there") {
		t.Fatalf("expected html to contain text block content")
	}
	if !strings.Contains(html, "https://example.com") {
		t.Fatalf("expected html to contain button url")
	}
	if !strings.Contains(html, "https://myrtle.example/logo.png") {
		t.Fatalf("expected html to contain logo url")
	}
	if !strings.Contains(html, `alt="Myrtle"`) {
		t.Fatalf("expected html to use product name as default logo alt")
	}

	markdown, err := email.Text()
	if err != nil {
		t.Fatalf("Text returned error: %v", err)
	}
	if !strings.Contains(markdown, "[Open](https://example.com)") {
		t.Fatalf("expected markdown to include button link")
	}
	if strings.Contains(markdown, "# Welcome") {
		t.Fatalf("expected markdown not to include subject heading when logo is present by default")
	}
	if strings.Contains(markdown, "[Myrtle](https://myrtle.example)") {
		t.Fatalf("expected markdown not to include product link when logo is present by default")
	}
}

func TestCustomBlockRegistry(t *testing.T) {
	t.Parallel()
	type Promo struct {
		Title string
	}

	registry := myrtle.NewRegistry()
	err := myrtle.Register(registry, "promo",
		func(value Promo, context myrtle.RenderContext) (string, error) {
			return "## " + value.Title + " for " + context.Values.ProductName, nil
		},
	)
	if err != nil {
		t.Fatalf("register returned error: %v", err)
	}

	builder := myrtle.NewBuilder(
		defaulttheme.New(),
		myrtle.WithRegistry(registry),
	)

	promoBlock, err := myrtle.Create(registry, "promo", Promo{Title: "Launch"})
	if err != nil {
		t.Fatalf("create returned error: %v", err)
	}

	email := builder.WithHeader(myrtle.HeaderTitle("Custom")).ProductName("Myrtle").Add(promoBlock).Build()

	if _, err := email.HTML(); err == nil {
		t.Fatalf("expected html render to fail for custom block without theme override")
	}

	markdown, err := email.Text()
	if err != nil {
		t.Fatalf("text returned error: %v", err)
	}
	if !strings.Contains(markdown, "## Launch for Myrtle") {
		t.Fatalf("expected markdown to contain custom rendered content")
	}
}

func TestFlatThemeBuildAndRender(t *testing.T) {
	t.Parallel()
	builder := myrtle.NewBuilder(flat.New())

	email := builder.WithHeader(myrtle.HeaderTitle("Flat style")).
		Preheader("Simple layout").
		Product("Myrtle", "https://myrtle.example").
		Logo("https://myrtle.example/logo.png", "Myrtle logo").
		AddText("Hello from flat").
		Build()

	html, err := email.HTML()
	if err != nil {
		t.Fatalf("html returned error: %v", err)
	}
	if !strings.Contains(html, "Hello from flat") {
		t.Fatalf("expected flat html to include block content")
	}
	if !strings.Contains(html, "Myrtle logo") {
		t.Fatalf("expected flat html to include custom logo alt text")
	}

	markdown, err := email.Text()
	if err != nil {
		t.Fatalf("text returned error: %v", err)
	}
	if strings.Contains(markdown, "[Myrtle](https://myrtle.example)") {
		t.Fatalf("expected flat markdown not to include product link when logo is present by default")
	}
}

func TestNewBuilderPanicsWithoutTheme(t *testing.T) {
	t.Parallel()
	defer func() {
		if recover() == nil {
			t.Fatalf("expected panic when theme is nil")
		}
	}()

	_ = myrtle.NewBuilder(nil)
}

func TestThemeBlockFallbackToDefault(t *testing.T) {
	t.Parallel()
	minimal := &minimalTheme{
		fallback: defaulttheme.New(),
	}

	builder := myrtle.NewBuilder(minimal)
	email := builder.WithHeader(myrtle.HeaderTitle("Fallback")).
		Preheader("delegation").
		AddText("Primary theme text").
		AddButton("Fallback button", "https://example.com/fallback").
		Build()

	html, err := email.HTML()
	if err != nil {
		t.Fatalf("html returned error: %v", err)
	}

	if !strings.Contains(html, "Primary theme text") {
		t.Fatalf("expected html to include primary theme text block")
	}
	if !strings.Contains(html, "https://example.com/fallback") {
		t.Fatalf("expected html to include fallback rendered button")
	}
}

func TestAddColumnsFunctionalAPI(t *testing.T) {
	t.Parallel()
	builder := myrtle.NewBuilder(defaulttheme.New())

	email := builder.WithHeader(myrtle.HeaderTitle("Columns")).
		AddColumns(
			func(column *myrtle.ColumnBuilder) {
				column.AddHeading("Left", myrtle.HeadingLevel(3)).
					AddText("Left body")
			},
			func(column *myrtle.ColumnBuilder) {
				column.AddHeading("Right", myrtle.HeadingLevel(3)).
					AddList([]string{"One", "Two"}, false)
			},
			myrtle.ColumnsWidths(70, 30),
		).
		Build()

	html, err := email.HTML()
	if err != nil {
		t.Fatalf("html returned error: %v", err)
	}
	if !strings.Contains(html, "Left body") || !strings.Contains(html, "Right") {
		t.Fatalf("expected html to contain rendered column content")
	}

	markdown, err := email.Text()
	if err != nil {
		t.Fatalf("text returned error: %v", err)
	}
	if !strings.Contains(markdown, "### Column 1") || !strings.Contains(markdown, "### Column 2") {
		t.Fatalf("expected markdown to contain column sections")
	}
}

func TestAddTextVariadic(t *testing.T) {
	t.Parallel()
	builder := myrtle.NewBuilder(defaulttheme.New())

	email := builder.
		WithHeader(myrtle.HeaderTitle("Paragraphs")).
		AddText("First paragraph.", "Second paragraph.").
		Build()

	html, err := email.HTML()
	if err != nil {
		t.Fatalf("html returned error: %v", err)
	}

	if strings.Count(html, "<p style=\"margin:0 0 16px;line-height:1.6;color:#111827;\">") < 2 {
		t.Fatalf("expected html to contain one text block per AddText argument")
	}

	markdown, err := email.Text()
	if err != nil {
		t.Fatalf("text returned error: %v", err)
	}

	if !strings.Contains(markdown, "First paragraph.\n\nSecond paragraph.") {
		t.Fatalf("expected markdown to contain both text entries")
	}
}

func TestStatsRowDeltaSemanticNormalization(t *testing.T) {
	t.Parallel()
	builder := myrtle.NewBuilder(defaulttheme.New())

	email := builder.
		AddStatsRow("Stats", []myrtle.StatItem{{Label: "Delivery", Value: "99.8%", Delta: "+0.3%", DeltaSemantic: "unexpected"}}).
		Build()

	html, err := email.HTML()
	if err != nil {
		t.Fatalf("html returned error: %v", err)
	}

	if strings.Contains(html, "#15803d") || strings.Contains(html, "#b91c1c") {
		t.Fatalf("expected unknown semantic to fallback to neutral delta color")
	}
}

func TestHeaderModeDisabledAndOverride(t *testing.T) {
	t.Parallel()
	builder := myrtle.NewBuilder(
		defaulttheme.New(),
		myrtle.WithHeaderMode(myrtle.HeaderModeDisabled),
	)

	emailWithoutHeader := builder.
		Product("Myrtle", "https://myrtle.example").
		Logo("https://myrtle.example/logo.png", "Myrtle logo").
		AddText("Body").
		Build()

	htmlWithoutHeader, err := emailWithoutHeader.HTML()
	if err != nil {
		t.Fatalf("html returned error: %v", err)
	}
	if strings.Contains(htmlWithoutHeader, "https://myrtle.example/logo.png") {
		t.Fatalf("expected no header logo when header mode is disabled")
	}
}

func TestHeaderLogoCenteredOption(t *testing.T) {
	t.Parallel()
	builder := myrtle.NewBuilder(defaulttheme.New())

	email := builder.
		WithHeader(
			myrtle.HeaderTitle("Centered logo"),
			myrtle.HeaderLogo("https://myrtle.example/logo.png", "Myrtle logo"),
			myrtle.HeaderLogoCentered(true),
		).
		AddText("Body").
		Build()

	html, err := email.HTML()
	if err != nil {
		t.Fatalf("html returned error: %v", err)
	}

	if !strings.Contains(html, "text-align:center;") {
		t.Fatalf("expected centered logo style when HeaderLogoCentered is enabled")
	}
}

func TestHeaderAlignmentDefaultCentered(t *testing.T) {
	t.Parallel()
	builder := myrtle.NewBuilder(defaulttheme.New())

	email := builder.
		WithHeader(
			myrtle.HeaderTitle("Default aligned"),
			myrtle.HeaderLogo("https://myrtle.example/logo.png", "Myrtle logo"),
			myrtle.HeaderProduct("Myrtle", "https://myrtle.example"),
			myrtle.HeaderShowTextWithLogo(true),
		).
		AddText("Body").
		Build()

	html, err := email.HTML()
	if err != nil {
		t.Fatalf("html returned error: %v", err)
	}

	if strings.Count(html, "text-align:center;") < 2 {
		t.Fatalf("expected centered header alignment by default for logo and product")
	}
}

func TestHeaderAlignOptionLeft(t *testing.T) {
	t.Parallel()
	builder := myrtle.NewBuilder(defaulttheme.New())

	email := builder.
		WithHeader(
			myrtle.HeaderTitle("Left aligned"),
			myrtle.HeaderLogo("https://myrtle.example/logo.png", "Myrtle logo"),
			myrtle.HeaderProduct("Myrtle", "https://myrtle.example"),
			myrtle.HeaderShowTextWithLogo(true),
			myrtle.HeaderAlign(myrtle.HeaderAlignmentLeft),
		).
		AddText("Body").
		Build()

	html, err := email.HTML()
	if err != nil {
		t.Fatalf("html returned error: %v", err)
	}

	if strings.Count(html, "text-align:left;") < 2 {
		t.Fatalf("expected left header alignment when HeaderAlign(left) is set")
	}
}

func TestNewBuilderWithHeaderOptions(t *testing.T) {
	t.Parallel()
	builder := myrtle.NewBuilder(
		defaulttheme.New(),
		myrtle.WithHeaderOptions(
			myrtle.HeaderTitle("Configured in constructor"),
			myrtle.HeaderProduct("Myrtle", "https://myrtle.example"),
			myrtle.HeaderRenderInMarkdown(true),
		),
	)

	email := builder.Preheader("Set via NewBuilder option").AddText("Body").Build()

	html, err := email.HTML()
	if err != nil {
		t.Fatalf("html returned error: %v", err)
	}
	if !strings.Contains(html, "<title>Configured in constructor</title>") {
		t.Fatalf("expected header title from constructor header options")
	}

	markdown, err := email.Text()
	if err != nil {
		t.Fatalf("text returned error: %v", err)
	}
	if !strings.Contains(markdown, "_Set via NewBuilder option_") {
		t.Fatalf("expected preheader from constructor header options")
	}
}

func TestHeaderWithLogoHidesTitleAndProductByDefault(t *testing.T) {
	t.Parallel()
	builder := myrtle.NewBuilder(defaulttheme.New())

	email := builder.
		WithHeader(
			myrtle.HeaderTitle("Hidden by default"),
			myrtle.HeaderLogo("https://myrtle.example/logo.png", "Myrtle logo"),
			myrtle.HeaderProduct("Myrtle", "https://myrtle.example"),
		).
		AddText("Body").
		Build()

	html, err := email.HTML()
	if err != nil {
		t.Fatalf("html returned error: %v", err)
	}
	if strings.Contains(html, ">Myrtle<") {
		t.Fatalf("expected product name not to render in HTML header when logo is present by default")
	}

	markdown, err := email.Text()
	if err != nil {
		t.Fatalf("text returned error: %v", err)
	}
	if strings.Contains(markdown, "# Hidden by default") {
		t.Fatalf("expected title not to render in markdown header when logo is present by default")
	}
	if strings.Contains(markdown, "[Myrtle](https://myrtle.example)") {
		t.Fatalf("expected product link not to render in markdown header when logo is present by default")
	}
}

func TestHeaderWithLogoCanShowTitleAndProductWhenOptedIn(t *testing.T) {
	t.Parallel()
	builder := myrtle.NewBuilder(defaulttheme.New())

	email := builder.
		WithHeader(
			myrtle.HeaderTitle("Visible when opted in"),
			myrtle.HeaderLogo("https://myrtle.example/logo.png", "Myrtle logo"),
			myrtle.HeaderProduct("Myrtle", "https://myrtle.example"),
			myrtle.HeaderRenderInMarkdown(true),
			myrtle.HeaderShowTextWithLogo(true),
		).
		AddText("Body").
		Build()

	html, err := email.HTML()
	if err != nil {
		t.Fatalf("html returned error: %v", err)
	}
	if !strings.Contains(html, ">Myrtle<") {
		t.Fatalf("expected product name to render in HTML header when opted in")
	}

	markdown, err := email.Text()
	if err != nil {
		t.Fatalf("text returned error: %v", err)
	}
	if !strings.Contains(markdown, "# Visible when opted in") {
		t.Fatalf("expected title to render in markdown header when opted in")
	}
	if !strings.Contains(markdown, "[Myrtle](https://myrtle.example)") {
		t.Fatalf("expected product link to render in markdown header when opted in")
	}
}

func TestMarkdownHeaderNotRenderedByDefault(t *testing.T) {
	t.Parallel()
	builder := myrtle.NewBuilder(defaulttheme.New())

	email := builder.
		Preheader("Header preheader").
		WithHeader(
			myrtle.HeaderTitle("Header title"),
			myrtle.HeaderProduct("Myrtle", "https://myrtle.example"),
		).
		AddText("Body text").
		Build()

	markdown, err := email.Text()
	if err != nil {
		t.Fatalf("text returned error: %v", err)
	}

	if strings.Contains(markdown, "# Header title") {
		t.Fatalf("expected title to be hidden in markdown by default")
	}
	if !strings.Contains(markdown, "_Header preheader_") {
		t.Fatalf("expected preheader to render in markdown independently of markdown header settings")
	}
	if strings.Contains(markdown, "[Myrtle](https://myrtle.example)") {
		t.Fatalf("expected product to be hidden in markdown by default")
	}
	if !strings.Contains(markdown, "Body text") {
		t.Fatalf("expected markdown body content to render")
	}
}

func TestMarkdownHeaderRenderedWhenEnabled(t *testing.T) {
	t.Parallel()
	builder := myrtle.NewBuilder(defaulttheme.New())

	email := builder.
		Preheader("Header preheader").
		WithHeader(
			myrtle.HeaderTitle("Header title"),
			myrtle.HeaderProduct("Myrtle", "https://myrtle.example"),
			myrtle.HeaderRenderInMarkdown(true),
		).
		AddText("Body text").
		Build()

	markdown, err := email.Text()
	if err != nil {
		t.Fatalf("text returned error: %v", err)
	}

	if !strings.Contains(markdown, "# Header title") {
		t.Fatalf("expected title to render in markdown when enabled")
	}
	if !strings.Contains(markdown, "_Header preheader_") {
		t.Fatalf("expected preheader to render in markdown when enabled")
	}
	if !strings.Contains(markdown, "[Myrtle](https://myrtle.example)") {
		t.Fatalf("expected product to render in markdown when enabled")
	}
}

func TestNewBlocksRender(t *testing.T) {
	t.Parallel()
	builder := myrtle.NewBuilder(defaulttheme.New())

	email := builder.
		Add(myrtle.HeroBlock{Title: "Hero title", Body: "Hero body", CTALabel: "Start", CTAURL: "https://example.com/start"}).
		AddFooterLinks([]myrtle.FooterLink{{Label: "Help", URL: "https://example.com/help"}}, "Footer note").
		AddPriceSummary("Summary", []myrtle.PriceLine{{Label: "Subtotal", Value: "$10.00"}}, "Total", "$10.00").
		AddEmptyState("Nothing here", "No items available", "Create", "https://example.com/create").
		Build()

	html, err := email.HTML()
	if err != nil {
		t.Fatalf("html returned error: %v", err)
	}

	for _, needle := range []string{"Hero title", "Footer note", "Summary", "Nothing here"} {
		if !strings.Contains(html, needle) {
			t.Fatalf("expected html to contain %q", needle)
		}
	}
}

func TestAdvancedLayoutAndPrimitiveVariantsRender(t *testing.T) {
	t.Parallel()
	builder := myrtle.NewBuilder(defaulttheme.New())

	email := builder.
		AddSection(
			[]myrtle.Block{
				myrtle.TextBlock{Text: "Section body"},
				myrtle.ButtonBlock{Label: "Open section", URL: "https://example.com/section", Style: myrtle.ButtonStyleOutline},
			},
			myrtle.SectionTitle("Section title"),
			myrtle.SectionSubtitle("Section subtitle"),
			myrtle.SectionPadding(20),
		).
		AddGrid(
			[]myrtle.GridItem{
				{Blocks: []myrtle.Block{myrtle.TextBlock{Text: "Grid 1"}}},
				{Blocks: []myrtle.Block{myrtle.TextBlock{Text: "Grid 2"}}},
				{Blocks: []myrtle.Block{myrtle.TextBlock{Text: "Grid 3"}}},
			},
			myrtle.GridColumns(2),
			myrtle.GridGap(14),
			myrtle.GridBorder(true),
		).
		AddCardList(
			[]myrtle.CardItem{{Title: "Card one", Body: "Body one", URL: "https://example.com/card1", CTALabel: "View"}, {Title: "Card two", Body: "Body two"}},
			myrtle.CardListColumns(2),
			myrtle.CardListGap(10),
		).
		AddColumns(
			func(column *myrtle.ColumnBuilder) {
				column.AddText("Left")
			},
			func(column *myrtle.ColumnBuilder) {
				column.AddText("Right")
			},
			myrtle.ColumnsGap(20),
			myrtle.ColumnsAlign(myrtle.ColumnsVerticalAlignMiddle),
		).
		AddDividerStyled(myrtle.DividerStyle(myrtle.DividerVariantDashed), myrtle.DividerThickness(2), myrtle.DividerInset(16)).
		AddSpacer(myrtle.SpacerSize(24)).
		Build()

	html, err := email.HTML()
	if err != nil {
		t.Fatalf("html returned error: %v", err)
	}

	if !strings.Contains(html, "Section title") || !strings.Contains(html, "Section subtitle") {
		t.Fatalf("expected section block heading and subtitle to render")
	}
	if !strings.Contains(html, "border-collapse:separate;border-spacing:14px") {
		t.Fatalf("expected grid gap configuration to render")
	}
	if !strings.Contains(html, "Grid 3") {
		t.Fatalf("expected wrapped grid item content to render")
	}
	if !strings.Contains(html, "Card one") || !strings.Contains(html, "https://example.com/card1") {
		t.Fatalf("expected card list item content and link to render")
	}
	if !strings.Contains(html, "valign=\"middle\"") || !strings.Contains(html, "padding:0 20px 0 0;") {
		t.Fatalf("expected columns alignment and gap styles to render")
	}
	if !strings.Contains(html, "border-top:2px dashed") || !strings.Contains(html, "margin:24px 16px;") {
		t.Fatalf("expected divider variant, thickness, and inset to render")
	}
	if !strings.Contains(html, "height:24px;") {
		t.Fatalf("expected spacer variant to resolve to configured size")
	}
}

func TestBlockGroupHelperAcrossMultiBlockAPIs(t *testing.T) {
	t.Parallel()
	left := myrtle.NewGroup().
		AddHeading("Left group", myrtle.HeadingLevel(3)).
		AddText("Left text")
	right := myrtle.NewGroup().
		AddHeading("Right group", myrtle.HeadingLevel(3)).
		AddText("Right text")

	section := myrtle.NewGroup().
		AddText("Section grouped content").
		AddButton("Open grouped section", "https://example.com/section", myrtle.ButtonStyle(myrtle.ButtonStyleOutline))

	gridOne := myrtle.NewGroup().AddText("Grid group 1")
	gridTwo := myrtle.NewGroup().AddText("Grid group 2")

	email := myrtle.NewBuilder(defaulttheme.New()).
		AddColumnsGroups(left, right, myrtle.ColumnsWidths(55, 45)).
		AddSectionGroup(section, myrtle.SectionTitle("Grouped section")).
		AddGridGroups([]*myrtle.Group{gridOne, gridTwo}, myrtle.GridColumns(2)).
		Build()

	html, err := email.HTML()
	if err != nil {
		t.Fatalf("html returned error: %v", err)
	}

	for _, needle := range []string{"Left group", "Right group", "Section grouped content", "Grid group 1", "Grid group 2"} {
		if !strings.Contains(html, needle) {
			t.Fatalf("expected html to contain %q", needle)
		}
	}
}

func TestGroupCanRenderAsBlock(t *testing.T) {
	t.Parallel()
	group := myrtle.NewGroup().
		AddHeading("Grouped heading", myrtle.HeadingLevel(3)).
		AddText("Grouped text").
		AddButton("Grouped CTA", "https://example.com/group", myrtle.ButtonStyle(myrtle.ButtonStyleOutline))

	email := myrtle.NewBuilder(defaulttheme.New()).
		Add(group).
		Build()

	html, err := email.HTML()
	if err != nil {
		t.Fatalf("html returned error: %v", err)
	}

	if !strings.Contains(html, "Grouped heading") || !strings.Contains(html, "Grouped text") || !strings.Contains(html, "https://example.com/group") {
		t.Fatalf("expected grouped block content to render in html")
	}

	markdown, err := email.Text()
	if err != nil {
		t.Fatalf("text returned error: %v", err)
	}

	if !strings.Contains(markdown, "### Grouped heading") || !strings.Contains(markdown, "Grouped text") || !strings.Contains(markdown, "[Grouped CTA](https://example.com/group)") {
		t.Fatalf("expected grouped block content to render in markdown")
	}
}

func TestButtonCalloutAndTableVariantsRender(t *testing.T) {
	t.Parallel()
	builder := myrtle.NewBuilder(
		defaulttheme.New(),
		myrtle.WithStyles(theme.Styles{ColorPrimary: "#2563eb", ColorSecondary: "#a855f7"}),
	)

	email := builder.
		AddButton("Secondary", "https://example.com", myrtle.ButtonTone(myrtle.ButtonToneSecondary), myrtle.ButtonAlign(myrtle.ButtonAlignmentRight), myrtle.ButtonFullWidth(true)).
		AddButton("Compact nowrap", "https://example.com/compact", myrtle.ButtonStyle(myrtle.ButtonStyleOutline), myrtle.ButtonSize(myrtle.ButtonSizeSmall), myrtle.ButtonNoWrap(true)).
		AddButton("Outline", "https://example.com/outline", myrtle.ButtonStyle(myrtle.ButtonStyleOutline)).
		AddButtonGroup([]myrtle.ButtonGroupButton{{Label: "Approve", URL: "https://example.com/approve", Tone: myrtle.ButtonTonePrimary}, {Label: "Review", URL: "https://example.com/review", Tone: myrtle.ButtonToneSecondary}}, myrtle.ButtonGroupAlign(myrtle.ButtonAlignmentCenter), myrtle.ButtonGroupJoined(true), myrtle.ButtonGroupFullWidthOnMobile(true)).
		AddButtonGroup([]myrtle.ButtonGroupButton{{Label: "Outline", URL: "https://example.com/outline-group", Style: myrtle.ButtonStyleOutline}, {Label: "Ghost", URL: "https://example.com/ghost-group", Style: myrtle.ButtonStyleGhost}}, myrtle.ButtonGroupAlign(myrtle.ButtonAlignmentLeft), myrtle.ButtonGroupGap(14), myrtle.ButtonGroupStackOnMobile(true)).
		AddCallout(myrtle.CalloutTypeCritical, "Critical", "Body", myrtle.CalloutStyle(myrtle.CalloutVariantSolid), myrtle.CalloutLink("Investigate", "https://example.com/investigate")).
		AddTable("Metrics", []string{"Name", "Value"}, [][]string{{"Users", "1200"}, {"Rate", "4.2%"}}, myrtle.TableZebraRows(true), myrtle.TableCompact(true), myrtle.TableRightAlignNumericColumns(true)).
		Build()

	html, err := email.HTML()
	if err != nil {
		t.Fatalf("html returned error: %v", err)
	}

	if !strings.Contains(html, "width:100%;max-width:100%;box-sizing:border-box;text-align:center;") {
		t.Fatalf("expected full-width secondary button styles")
	}
	if !strings.Contains(html, "background:#a855f7;color:#ffffff;") {
		t.Fatalf("expected secondary button variant to render as filled secondary color")
	}
	if !strings.Contains(html, "https://example.com/outline") || !strings.Contains(html, "border:1px solid #2563eb;color:#2563eb;background:#ffffff;") {
		t.Fatalf("expected outline button variant to render with outlined primary color style")
	}
	if !strings.Contains(html, "https://example.com/compact") || !strings.Contains(html, "padding:8px 14px;font-size:13px;") || !strings.Contains(html, "white-space:nowrap;") {
		t.Fatalf("expected compact button with no-wrap styling to render")
	}
	if !strings.Contains(html, "<p style=\"margin:20px 0;text-align:right;\">") {
		t.Fatalf("expected aligned button wrapper style")
	}
	if !strings.Contains(html, "https://example.com/approve") || !strings.Contains(html, "https://example.com/review") {
		t.Fatalf("expected button group links to render")
	}
	if !strings.Contains(html, "margin-left:-1px;") {
		t.Fatalf("expected joined button group to collapse adjacent borders")
	}
	if !strings.Contains(html, "https://example.com/review") || !strings.Contains(html, "background:#a855f7;color:#ffffff;") {
		t.Fatalf("expected secondary button group item to render as filled secondary color")
	}
	if !strings.Contains(html, "https://example.com/outline-group") || !strings.Contains(html, "border:1px solid #2563eb;color:#2563eb;background:#ffffff;") {
		t.Fatalf("expected outline button group item to render with outlined primary color style")
	}
	if !strings.Contains(html, "padding-left:14px;") {
		t.Fatalf("expected configured button group gap to render")
	}
	if !strings.Contains(html, ".myrtle-btn-group-mobile-full td a") {
		t.Fatalf("expected full-width-on-mobile button group styles to render")
	}
	if !strings.Contains(html, ".myrtle-btn-group-stack td") || !strings.Contains(html, "padding-top:14px") {
		t.Fatalf("expected stack-on-mobile button group styles to render")
	}
	if !strings.Contains(html, "border-top-left-radius:8px;border-bottom-left-radius:8px;") || !strings.Contains(html, "border-top-right-radius:8px;border-bottom-right-radius:8px;") {
		t.Fatalf("expected joined button group to keep only outer corner rounding")
	}
	if !strings.Contains(html, "background:#dc2626") {
		t.Fatalf("expected critical solid callout background")
	}
	if !strings.Contains(html, "https://example.com/investigate") {
		t.Fatalf("expected callout link to render")
	}
	if !strings.Contains(html, "text-align:right;") {
		t.Fatalf("expected numeric table cells to be right aligned")
	}
	if !strings.Contains(html, "background:#f9fafb;") {
		t.Fatalf("expected zebra row styling in table")
	}
}

func TestBarChartThicknessAndTransparentBackgroundRender(t *testing.T) {
	t.Parallel()
	builder := myrtle.NewBuilder(defaulttheme.New())

	email := builder.
		AddBarChart("Delivery", []myrtle.BarChartItem{{Label: "US", Percent: 52}}, myrtle.BarChartThickness(14), myrtle.BarChartTransparentBackground(true)).
		Build()

	html, err := email.HTML()
	if err != nil {
		t.Fatalf("html returned error: %v", err)
	}

	if !strings.Contains(html, "height:14px;line-height:14px") {
		t.Fatalf("expected configured bar thickness in html")
	}
	if !strings.Contains(html, "background:transparent;") {
		t.Fatalf("expected transparent bar chart background in html")
	}
}

func TestPriceSummaryDiscountLineEmphasis(t *testing.T) {
	t.Parallel()
	builder := myrtle.NewBuilder(
		defaulttheme.New(),
		myrtle.WithStyles(theme.Styles{ColorPrimary: "#ff00aa"}),
	)

	email := builder.
		AddPriceSummary("Invoice", []myrtle.PriceLine{{Label: "Subtotal", Value: "$100.00"}, {Label: "Discount", Value: "-$5.00"}}, "Total", "$95.00").
		Build()

	html, err := email.HTML()
	if err != nil {
		t.Fatalf("html returned error: %v", err)
	}

	if !strings.Contains(html, "Discount") {
		t.Fatalf("expected html to include discount line")
	}
	if !strings.Contains(html, "color:#ff00aa;font-weight:600;") {
		t.Fatalf("expected discount line to render with emphasized discount styling")
	}
}

func TestTilesRenderOptionsAndVariants(t *testing.T) {
	t.Parallel()
	builder := myrtle.NewBuilder(defaulttheme.New())

	email := builder.
		AddTiles(
			[]myrtle.TileEntry{
				{Content: "🚀", Title: "Launch", Subtitle: "Ready", Variant: myrtle.TileVariantHighlight},
				{Content: "7", Title: "Queued"},
				{Content: "OK", Variant: myrtle.TileVariantSuccess},
				{Content: "!", Title: "Attention", Variant: myrtle.TileVariantCritical},
			},
			myrtle.TilesColumns(4),
			myrtle.TilesBorder(true),
		).
		Build()

	html, err := email.HTML()
	if err != nil {
		t.Fatalf("html returned error: %v", err)
	}

	if !strings.Contains(html, "width=\"25%\"") {
		t.Fatalf("expected 4-column tiles width in html")
	}
	if !strings.Contains(html, "border:1px solid") {
		t.Fatalf("expected optional tile border in html")
	}
	if !strings.Contains(html, "background:#f8fafc;") {
		t.Fatalf("expected default subtle tile card background in html")
	}
	if strings.Contains(html, "background:#2563eb") || strings.Contains(html, "background:#dcfce7") || strings.Contains(html, "background:#fef9c3") || strings.Contains(html, "background:#fee2e2") {
		t.Fatalf("expected tile content area to have no background fill")
	}
	if !strings.Contains(html, "background:#eff6ff;") || !strings.Contains(html, "border:1px solid #fca5a5;") {
		t.Fatalf("expected variant-specific tile card background and border in html")
	}
	if !strings.Contains(html, "font-size:40px;") {
		t.Fatalf("expected larger tile content size in html")
	}
	if !strings.Contains(html, "height:100%;box-sizing:border-box;") {
		t.Fatalf("expected tile cards to fill row height for consistent sizing")
	}
	if strings.Contains(html, "color:#2563eb;") || strings.Contains(html, "color:#166534;") || strings.Contains(html, "color:#854d0e;") || strings.Contains(html, "color:#991b1b;") {
		t.Fatalf("expected variants not to change tile content text color")
	}
	if !strings.Contains(html, "Launch") || !strings.Contains(html, "Ready") {
		t.Fatalf("expected title and subtitle to render in html")
	}

	linkedTitleEmail := myrtle.NewBuilder(defaulttheme.New()).
		AddTiles(
			[]myrtle.TileEntry{{Content: "🔗", Title: "Open details", URL: "https://example.com/details"}, {Content: "📝", Title: "No link title"}},
		).
		Build()

	linkedTitleHTML, err := linkedTitleEmail.HTML()
	if err != nil {
		t.Fatalf("linked title html returned error: %v", err)
	}
	if !strings.Contains(linkedTitleHTML, "<a href=\"https://example.com/details\"") {
		t.Fatalf("expected tile title to render as link when url is provided")
	}
	if strings.Contains(linkedTitleHTML, "<a href=\"\"") {
		t.Fatalf("expected empty tile title url not to render as link")
	}

	leftAlignedEmail := myrtle.NewBuilder(defaulttheme.New()).
		AddTiles(
			[]myrtle.TileEntry{{Content: "📦", Title: "Shipped"}},
			myrtle.TilesAlign(myrtle.TileAlignmentLeft),
		).
		Build()

	leftAlignedHTML, err := leftAlignedEmail.HTML()
	if err != nil {
		t.Fatalf("left aligned html returned error: %v", err)
	}
	if !strings.Contains(leftAlignedHTML, "text-align:left;") {
		t.Fatalf("expected left aligned tiles text in html")
	}
	if !strings.Contains(leftAlignedHTML, "margin:0 0 6px;") {
		t.Fatalf("expected left aligned tile content margin in html")
	}

	rightAlignedEmail := myrtle.NewBuilder(defaulttheme.New()).
		AddTiles(
			[]myrtle.TileEntry{{Content: "🧾", Title: "Invoices"}},
			myrtle.TilesAlign(myrtle.TileAlignmentRight),
		).
		Build()

	rightAlignedHTML, err := rightAlignedEmail.HTML()
	if err != nil {
		t.Fatalf("right aligned html returned error: %v", err)
	}
	if !strings.Contains(rightAlignedHTML, "text-align:right;") {
		t.Fatalf("expected right aligned tiles text in html")
	}
	if !strings.Contains(rightAlignedHTML, "margin:0 0 6px auto;") {
		t.Fatalf("expected right aligned tile content margin in html")
	}

	noContentEmail := myrtle.NewBuilder(defaulttheme.New()).
		AddTiles(
			[]myrtle.TileEntry{{Title: "No icon", Subtitle: "Text only"}},
		).
		Build()

	noContentHTML, err := noContentEmail.HTML()
	if err != nil {
		t.Fatalf("no content html returned error: %v", err)
	}
	if strings.Contains(noContentHTML, "font-size:40px;") {
		t.Fatalf("expected no tile content container when content is omitted")
	}
	if !strings.Contains(noContentHTML, "padding:8px;") {
		t.Fatalf("expected no-content tile card to render with smaller padding")
	}
	if !strings.Contains(noContentHTML, "border:1px solid transparent;") {
		t.Fatalf("expected unbordered default tiles to reserve border space for consistent sizing")
	}
	if !strings.Contains(noContentHTML, "No icon") {
		t.Fatalf("expected no-content tile title to render in html")
	}

	transparentEmail := myrtle.NewBuilder(defaulttheme.New()).
		AddTiles(
			[]myrtle.TileEntry{{Content: "1", Title: "One"}},
			myrtle.TilesTransparentBackground(true),
		).
		Build()

	transparentHTML, err := transparentEmail.HTML()
	if err != nil {
		t.Fatalf("transparent html returned error: %v", err)
	}
	if !strings.Contains(transparentHTML, "background:transparent;") {
		t.Fatalf("expected transparent tile card background when enabled")
	}
}

func TestNewBuilderWithHeaderOverrideLater(t *testing.T) {
	t.Parallel()
	builder := myrtle.NewBuilder(
		defaulttheme.New(),
		myrtle.WithHeader(myrtle.BuildHeader(
			myrtle.HeaderTitle("Initial"),
		)),
	)

	email := builder.Preheader("Initial preheader").WithHeader(myrtle.HeaderTitle("Overridden")).AddText("Body").Build()

	html, err := email.HTML()
	if err != nil {
		t.Fatalf("html returned error: %v", err)
	}
	if !strings.Contains(html, "<title>Overridden</title>") {
		t.Fatalf("expected chained HeaderTitle to override constructor header")
	}
}

func TestBuilderCloneConcurrentUsage(t *testing.T) {
	t.Parallel()
	baseBuilder := myrtle.NewBuilder(defaulttheme.New()).
		WithHeader(myrtle.HeaderTitle("Template")).
		AddText("base")

	baseEmail := baseBuilder.Build()
	baseHTML, err := baseEmail.HTML()
	if err != nil {
		t.Fatalf("base html returned error: %v", err)
	}
	if strings.Contains(baseHTML, "thread-line-") {
		t.Fatalf("expected base builder output not to contain per-thread content")
	}

	const workerCount = 24
	results := make(chan string, workerCount)

	var waitGroup sync.WaitGroup
	waitGroup.Add(workerCount)

	for index := 0; index < workerCount; index++ {
		index := index
		go func() {
			defer waitGroup.Done()

			email := baseBuilder.
				Clone().
				Preheader(fmt.Sprintf("preheader-%d", index)).
				AddText(fmt.Sprintf("thread-line-%d", index)).
				Build()

			html, err := email.HTML()
			if err != nil {
				results <- "__error__:" + err.Error()
				return
			}

			results <- html
		}()
	}

	waitGroup.Wait()
	close(results)

	for html := range results {
		if strings.HasPrefix(html, "__error__:") {
			t.Fatalf("clone html returned error: %s", html)
		}
		if !strings.Contains(html, "base") {
			t.Fatalf("expected cloned builder output to retain base content")
		}
		if !strings.Contains(html, "thread-line-") {
			t.Fatalf("expected cloned builder output to include thread-specific content")
		}
	}
}

type minimalTheme struct {
	fallback theme.Theme
}

func (themeImpl *minimalTheme) Name() string {
	return "minimal"
}

func (themeImpl *minimalTheme) DefaultStyles() theme.Styles {
	if themeImpl.fallback == nil {
		return theme.Styles{}
	}

	return themeImpl.fallback.DefaultStyles()
}

func (themeImpl *minimalTheme) RenderHTML(view theme.EmailView) (string, error) {
	return "<html><body>" + strings.Join(view.Blocks, "") + "</body></html>", nil
}

func (themeImpl *minimalTheme) RenderBlockHTML(view theme.BlockView) (string, bool, error) {
	if textBlock, ok := view.Data.(myrtle.TextBlock); ok {
		return fmt.Sprintf("<p>%s</p>", textBlock.Text), true, nil
	}

	if themeImpl.fallback == nil {
		return "", false, nil
	}

	return themeImpl.fallback.RenderBlockHTML(view)
}
