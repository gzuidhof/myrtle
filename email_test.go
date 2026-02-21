package myrtle_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/gzuidhof/myrtle"
	"github.com/gzuidhof/myrtle/theme"
	defaulttheme "github.com/gzuidhof/myrtle/theme/default"
	"github.com/gzuidhof/myrtle/theme/flat"
)

func TestBuildAndRender(t *testing.T) {
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
	defer func() {
		if recover() == nil {
			t.Fatalf("expected panic when theme is nil")
		}
	}()

	_ = myrtle.NewBuilder(nil)
}

func TestThemeBlockFallbackToDefault(t *testing.T) {
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

func TestButtonCalloutAndTableVariantsRender(t *testing.T) {
	builder := myrtle.NewBuilder(defaulttheme.New())

	email := builder.
		AddButton("Secondary", "https://example.com", myrtle.ButtonStyle(myrtle.ButtonVariantSecondary), myrtle.ButtonAlign(myrtle.ButtonAlignmentRight), myrtle.ButtonFullWidth(true)).
		AddButtonGroup([]myrtle.ButtonGroupButton{{Label: "Approve", URL: "https://example.com/approve", Variant: myrtle.ButtonVariantPrimary}, {Label: "Review", URL: "https://example.com/review", Variant: myrtle.ButtonVariantSecondary}}, myrtle.ButtonGroupAlign(myrtle.ButtonAlignmentCenter), myrtle.ButtonGroupJoined(true)).
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
	if !strings.Contains(html, "<p style=\"margin:20px 0;text-align:right;\">") {
		t.Fatalf("expected aligned button wrapper style")
	}
	if !strings.Contains(html, "https://example.com/approve") || !strings.Contains(html, "https://example.com/review") {
		t.Fatalf("expected button group links to render")
	}
	if !strings.Contains(html, "margin-left:-1px;") {
		t.Fatalf("expected joined button group to collapse adjacent borders")
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
	builder := myrtle.NewBuilder(
		defaulttheme.New(),
		myrtle.WithStyles(theme.Styles{PrimaryColor: "#ff00aa"}),
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

func TestNewBuilderWithHeaderOverrideLater(t *testing.T) {
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

type minimalTheme struct {
	fallback theme.Theme
}

func (themeImpl *minimalTheme) Name() string {
	return "minimal"
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
