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
	"github.com/gzuidhof/myrtle/theme/terminal"
)

func TestBuildAndRender(t *testing.T) {
	t.Parallel()
	builder := myrtle.NewBuilder(defaulttheme.New())

	email := builder.
		WithPreheader("This is the preheader").
		WithHeader(
			myrtle.NewGroup().
				Add(myrtle.ImageBlock{Src: "https://myrtle.example/logo.png", Alt: "Myrtle"}).
				Add(myrtle.HeadingBlock{Text: "Welcome", Level: 1}),
			myrtle.HeaderRenderInText(true),
		).
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
	if !strings.Contains(html, "Welcome") {
		t.Fatalf("expected html to contain header heading")
	}

	text, err := email.Text()
	if err != nil {
		t.Fatalf("Text returned error: %v", err)
	}
	if !strings.Contains(text, "Open (https://example.com)") {
		t.Fatalf("expected text fallback to include button link")
	}
	if !strings.Contains(text, "Welcome") {
		t.Fatalf("expected text fallback to include header when enabled")
	}
}

func TestImageBlockHrefRendersLinkedImage(t *testing.T) {
	t.Parallel()
	builder := myrtle.NewBuilder(defaulttheme.New())

	email := builder.
		Add(myrtle.ImageBlock{Src: "https://myrtle.example/logo.png", Alt: "Myrtle", Href: "https://example.com"}).
		Build()

	html, err := email.HTML()
	if err != nil {
		t.Fatalf("HTML returned error: %v", err)
	}

	if !strings.Contains(html, `<a href="https://example.com"`) {
		t.Fatalf("expected image to be wrapped in link")
	}
	if !strings.Contains(html, `src="https://myrtle.example/logo.png"`) {
		t.Fatalf("expected image src to render inside linked image")
	}
}

func TestImageBlockStyleOverridesRender(t *testing.T) {
	t.Parallel()
	builder := myrtle.NewBuilder(defaulttheme.New())

	email := builder.
		AddImage(
			"https://myrtle.example/logo.png",
			"Myrtle",
			myrtle.ImageTopSpacing(8),
			myrtle.ImageBottomSpacing(0),
			myrtle.ImageInsetMode(myrtle.InsetModeNone),
			myrtle.ImageTopCorners(),
		).
		Build()

	html, err := email.HTML()
	if err != nil {
		t.Fatalf("HTML returned error: %v", err)
	}

	if !strings.Contains(html, `margin:8px 0 0px;`) {
		t.Fatalf("expected custom top and bottom spacing to render")
	}
	if !strings.Contains(html, `border-radius:10px 10px 0 0;`) {
		t.Fatalf("expected top-corner image radius to render")
	}
}

func TestCustomContainerWidthAndPaddingStylesRender(t *testing.T) {
	t.Parallel()
	builder := myrtle.NewBuilder(
		flat.New(),
		myrtle.WithStyles(theme.Styles{
			WidthMain:           "92%",
			MaxWidthMain:        "720px",
			OuterPadding:        "40px",
			OutsideContentInset: "30px",
		}),
	)

	email := builder.
		WithHeader(myrtle.TextBlock{Text: "Outside header"}, myrtle.HeaderPlacement(myrtle.HeaderPlacementOutside)).
		AddText("Body").
		WithFooter(myrtle.TextBlock{Text: "Outside footer"}, myrtle.FooterPlacement(myrtle.FooterPlacementOutside)).
		Build()

	html, err := email.HTML()
	if err != nil {
		t.Fatalf("html returned error: %v", err)
	}

	if !strings.Contains(html, "padding:40px;") {
		t.Fatalf("expected custom outer padding to render in html")
	}
	if !strings.Contains(html, "width:92%;max-width:720px;") {
		t.Fatalf("expected custom width and max-width to render in html")
	}
	if !strings.Contains(html, "padding:0 30px;") {
		t.Fatalf("expected custom outside content inset to render in html")
	}
}

func TestMSOCompatibilityStyleToggleControlsOutlookSpacerFallback(t *testing.T) {
	t.Parallel()

	builderDefault := myrtle.NewBuilder(defaulttheme.New())
	emailDefault := builderDefault.AddText("First").AddText("Second").Build()
	htmlDefault, err := emailDefault.HTML()
	if err != nil {
		t.Fatalf("html returned error: %v", err)
	}

	if !strings.Contains(htmlDefault, "mso-padding-alt:16px 0 0 0;line-height:0;font-size:0;height:0;") {
		t.Fatalf("expected mso compatibility spacer fallback to render by default")
	}

	builderDisabled := myrtle.NewBuilder(
		defaulttheme.New(),
		myrtle.WithStyles(theme.Styles{MSOCompatibility: theme.MSOCompatibilityModeOff}),
	)
	emailDisabled := builderDisabled.AddText("First").AddText("Second").Build()
	htmlDisabled, err := emailDisabled.HTML()
	if err != nil {
		t.Fatalf("html returned error: %v", err)
	}

	if strings.Contains(htmlDisabled, "mso-padding-alt:16px 0 0 0;line-height:0;font-size:0;height:0;") {
		t.Fatalf("expected mso compatibility spacer fallback to be disabled when MSOCompatibility is off")
	}
}

func TestAddWithNewCustomBlock(t *testing.T) {
	t.Parallel()
	type Promo struct {
		Title string
	}

	builder := myrtle.NewBuilder(defaulttheme.New())
	block := myrtle.NewCustomBlock(
		"promo_direct",
		Promo{Title: "Launch"},
		func(value Promo, values theme.Values) (string, error) {
			_ = values
			return "<section><h2>" + value.Title + " for Myrtle</h2></section>", nil
		},
		func(value Promo, context myrtle.RenderContext) (string, error) {
			_ = context
			return "## " + value.Title + " for Myrtle", nil
		},
	)

	email := builder.Add(block).Build()

	html, err := email.HTML()
	if err != nil {
		t.Fatalf("html returned error: %v", err)
	}
	if !strings.Contains(html, "Launch for Myrtle") {
		t.Fatalf("expected html to contain custom rendered content")
	}

	text, err := email.Text()
	if err != nil {
		t.Fatalf("text returned error: %v", err)
	}
	if !strings.Contains(text, "## Launch for Myrtle") {
		t.Fatalf("expected text to contain custom rendered content")
	}
}

func TestNewCustomBlockWithLayoutSpecUsesProvidedLayout(t *testing.T) {
	t.Parallel()

	block := myrtle.NewCustomBlockWithLayoutSpec(
		"promo_direct",
		struct{ Title string }{Title: "Launch"},
		myrtle.LayoutSpec{InsetMode: myrtle.InsetModeCustom, CustomInset: "14px"},
		func(value struct{ Title string }, values theme.Values) (string, error) {
			_ = value
			_ = values
			return "<p>ok</p>", nil
		},
		func(value struct{ Title string }, context myrtle.RenderContext) (string, error) {
			_ = value
			_ = context
			return "ok", nil
		},
	)

	spec := block.LayoutSpec()
	if spec.InsetMode != myrtle.InsetModeCustom {
		t.Fatalf("expected inset mode custom, got %q", spec.InsetMode)
	}
	if spec.CustomInset != "14px" {
		t.Fatalf("expected custom inset 14px, got %q", spec.CustomInset)
	}
}

func TestNewCustomBlockDefaultsToDefaultLayoutSpec(t *testing.T) {
	t.Parallel()

	block := myrtle.NewCustomBlock(
		"promo_direct",
		struct{ Title string }{Title: "Launch"},
		func(value struct{ Title string }, values theme.Values) (string, error) {
			_ = value
			_ = values
			return "<p>ok</p>", nil
		},
		func(value struct{ Title string }, context myrtle.RenderContext) (string, error) {
			_ = value
			_ = context
			return "ok", nil
		},
	)

	spec := block.LayoutSpec()
	if spec.InsetMode != myrtle.InsetModeDefault {
		t.Fatalf("expected default inset mode, got %q", spec.InsetMode)
	}
	if spec.CustomInset != "" {
		t.Fatalf("expected empty custom inset for default layout, got %q", spec.CustomInset)
	}
}

func TestFlatThemeBuildAndRender(t *testing.T) {
	t.Parallel()
	builder := myrtle.NewBuilder(flat.New())
	email := builder.WithHeader(myrtle.HeadingBlock{Text: "Flat style", Level: 1}).
		WithPreheader("Simple layout").
		AddText("Hello from flat").
		Build()

	html, err := email.HTML()
	if err != nil {
		t.Fatalf("html returned error: %v", err)
	}
	if !strings.Contains(html, "Hello from flat") {
		t.Fatalf("expected flat html to include block content")
	}
	if !strings.Contains(html, "Flat style") {
		t.Fatalf("expected flat html to include header content")
	}

	text, err := email.Text()
	if err != nil {
		t.Fatalf("text returned error: %v", err)
	}
	if strings.Contains(text, "# Flat style") {
		t.Fatalf("expected flat text not to include header by default")
	}
}

func TestHeadingLevelsRenderExpectedTags(t *testing.T) {
	t.Parallel()

	email := myrtle.NewBuilder(defaulttheme.New()).
		AddHeading("Level one", myrtle.HeadingLevel(1)).
		AddHeading("Level two", myrtle.HeadingLevel(2)).
		AddHeading("Level three", myrtle.HeadingLevel(3)).
		AddHeading("Level four", myrtle.HeadingLevel(4)).
		AddHeading("Level five", myrtle.HeadingLevel(5)).
		AddHeading("Level six", myrtle.HeadingLevel(6)).
		Build()

	html, err := email.HTML()
	if err != nil {
		t.Fatalf("html returned error: %v", err)
	}

	if !strings.Contains(html, "<h1") || !strings.Contains(html, ">Level one</h1>") {
		t.Fatalf("expected level 1 heading to render as h1")
	}
	if !strings.Contains(html, "<h2") || !strings.Contains(html, ">Level two</h2>") {
		t.Fatalf("expected level 2 heading to render as h2")
	}
	if !strings.Contains(html, "<h3") || !strings.Contains(html, ">Level three</h3>") {
		t.Fatalf("expected level 3 heading to render as h3")
	}
	if !strings.Contains(html, "<h4") || !strings.Contains(html, ">Level four</h4>") {
		t.Fatalf("expected level 4 heading to render as h4")
	}
	if !strings.Contains(html, "<h5") || !strings.Contains(html, ">Level five</h5>") {
		t.Fatalf("expected level 5 heading to render as h5")
	}
	if !strings.Contains(html, "<h6") || !strings.Contains(html, ">Level six</h6>") {
		t.Fatalf("expected level 6 heading to render as h6")
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
	email := builder.WithHeader(myrtle.HeadingBlock{Text: "Fallback", Level: 1}).
		WithPreheader("delegation").
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

func TestAddColumnsGroupAPI(t *testing.T) {
	t.Parallel()
	builder := myrtle.NewBuilder(defaulttheme.New())
	left := myrtle.NewGroup().
		AddHeading("Left", myrtle.HeadingLevel(3)).
		AddText("Left body")
	right := myrtle.NewGroup().
		AddHeading("Right", myrtle.HeadingLevel(3)).
		AddList([]string{"One", "Two"}, false)

	email := builder.WithHeader(myrtle.HeadingBlock{Text: "Columns", Level: 1}).
		AddColumns(
			left,
			right,
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

	text, err := email.Text()
	if err != nil {
		t.Fatalf("text returned error: %v", err)
	}
	if !strings.Contains(text, "[ Column 1 ]") || !strings.Contains(text, "[ Column 2 ]") {
		t.Fatalf("expected text fallback to contain column sections")
	}
}

func TestAddTextAndOptions(t *testing.T) {
	t.Parallel()
	builder := myrtle.NewBuilder(defaulttheme.New())

	email := builder.
		WithHeader(myrtle.HeadingBlock{Text: "Paragraphs", Level: 1}).
		AddText("First paragraph.", myrtle.TextTone(myrtle.ToneMuted)).
		AddText("Second paragraph.", myrtle.TextSize(myrtle.TextSizeSmall)).
		AddText("Third paragraph.", myrtle.TextTone(myrtle.ToneDanger), myrtle.TextAlign(myrtle.TextAlignEnd), myrtle.TextWeight(myrtle.TextWeightSemibold)).
		AddText("Fourth paragraph.", myrtle.TextNoMargin(true), myrtle.TextSpacing(myrtle.TextSpacingCompact), myrtle.TextTransform(myrtle.TextTransformUppercase)).
		Build()

	html, err := email.HTML()
	if err != nil {
		t.Fatalf("html returned error: %v", err)
	}

	if strings.Count(html, "<p style=\"margin:0 0 16px;line-height:1.6;") < 2 {
		t.Fatalf("expected html to contain one text block per AddText argument")
	}
	if !strings.Contains(html, "color:#6b7280;") {
		t.Fatalf("expected muted text style to be rendered")
	}
	if !strings.Contains(html, "font-size:13px;") {
		t.Fatalf("expected small text style to be rendered")
	}
	if !strings.Contains(html, "color:#b91c1c;") {
		t.Fatalf("expected danger tone text style to be rendered")
	}
	if !strings.Contains(html, "text-align:right;") {
		t.Fatalf("expected text alignment style to be rendered")
	}
	if !strings.Contains(html, "font-weight:600;") {
		t.Fatalf("expected text weight style to be rendered")
	}
	if !strings.Contains(html, "margin:0;line-height:1.4;") {
		t.Fatalf("expected no-margin and compact spacing styles to be rendered")
	}
	if !strings.Contains(html, "text-transform:uppercase;") {
		t.Fatalf("expected text transform style to be rendered")
	}

	text, err := email.Text()
	if err != nil {
		t.Fatalf("text returned error: %v", err)
	}

	if !strings.Contains(text, "First paragraph.\n\nSecond paragraph.\n\nThird paragraph.\n\nFourth paragraph.") {
		t.Fatalf("expected text to contain all text entries")
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
		WithHeader(myrtle.ImageBlock{Src: "https://myrtle.example/logo.png", Alt: "Myrtle logo"}).
		WithoutHeader().
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

func TestWithHeaderNilDisablesHeader(t *testing.T) {
	t.Parallel()
	builder := myrtle.NewBuilder(defaulttheme.New())

	email := builder.
		WithHeader(
			myrtle.ImageBlock{Src: "https://myrtle.example/logo.png", Alt: "Myrtle logo"},
		).
		WithHeader(nil).
		AddText("Body").
		Build()

	html, err := email.HTML()
	if err != nil {
		t.Fatalf("html returned error: %v", err)
	}

	if strings.Contains(html, "https://myrtle.example/logo.png") {
		t.Fatalf("expected header logo not to render when WithHeader(nil) is used")
	}
}

func TestNewBuilderWithHeaderOptionsNilDisablesHeader(t *testing.T) {
	t.Parallel()
	builder := myrtle.NewBuilder(
		defaulttheme.New(),
		myrtle.WithHeader(myrtle.HeadingBlock{Text: "Should not render", Level: 1}),
		myrtle.WithHeader(nil),
	)

	email := builder.AddText("Body").Build()

	html, err := email.HTML()
	if err != nil {
		t.Fatalf("html returned error: %v", err)
	}

	if strings.Contains(html, "Should not render") {
		t.Fatalf("expected header title not to render when WithHeaderOptions(nil) is used")
	}
}

func TestHeaderBlockCanControlAlignment(t *testing.T) {
	t.Parallel()
	builder := myrtle.NewBuilder(defaulttheme.New())

	email := builder.
		WithHeader(
			myrtle.TextBlock{Text: "Centered header text", Align: myrtle.TextAlignCenter},
		).
		AddText("Body").
		Build()

	html, err := email.HTML()
	if err != nil {
		t.Fatalf("html returned error: %v", err)
	}

	if !strings.Contains(html, "Centered header text") || !strings.Contains(html, "text-align:center;") {
		t.Fatalf("expected centered header block content in html")
	}
}

func TestHeaderBlockRendersInHTML(t *testing.T) {
	t.Parallel()
	builder := myrtle.NewBuilder(defaulttheme.New())

	email := builder.WithHeader(myrtle.HeadingBlock{Text: "Default aligned", Level: 1}).
		AddText("Body").
		Build()

	html, err := email.HTML()
	if err != nil {
		t.Fatalf("html returned error: %v", err)
	}

	if !strings.Contains(html, "Default aligned") {
		t.Fatalf("expected header block to render in html")
	}
}

func TestHeaderCanRenderOutsideMainContentBox(t *testing.T) {
	t.Parallel()
	builder := myrtle.NewBuilder(defaulttheme.New())

	email := builder.
		WithHeader(myrtle.HeadingBlock{Text: "Outside header", Level: 1}, myrtle.HeaderPlacement(myrtle.HeaderPlacementOutside)).
		AddText("Body").
		Build()

	html, err := email.HTML()
	if err != nil {
		t.Fatalf("html returned error: %v", err)
	}

	containerIndex := strings.Index(html, "max-width:640px;margin:0 auto;mso-table-lspace:0pt;mso-table-rspace:0pt;border-collapse:separate;border-spacing:0;")
	headerIndex := strings.Index(html, "Outside header")
	if containerIndex == -1 || headerIndex == -1 {
		t.Fatalf("expected both main container and header to be rendered")
	}
	if headerIndex > containerIndex {
		t.Fatalf("expected outside header to render before the main content container")
	}
}

func TestHeaderDefaultInsetMatchesMainContentLogic(t *testing.T) {
	t.Parallel()

	builder := myrtle.NewBuilder(defaulttheme.New())

	insideEmail := builder.
		WithHeader(myrtle.MessageBlock{Subject: "Inside header subject", Preview: "inside"}).
		AddText("Body").
		Build()

	insideHTML, err := insideEmail.HTML()
	if err != nil {
		t.Fatalf("inside html returned error: %v", err)
	}
	insideHeaderIndex := strings.Index(insideHTML, "Inside header subject")
	if insideHeaderIndex == -1 {
		t.Fatalf("expected inside header to render")
	}
	insideStart := insideHeaderIndex - 1200
	if insideStart < 0 {
		insideStart = 0
	}
	insideSnippet := insideHTML[insideStart:insideHeaderIndex]
	if !strings.Contains(insideSnippet, "padding-top:24px;padding-right:24px;padding-bottom:0;padding-left:24px;") {
		t.Fatalf("expected inside header default inset to use main content side inset")
	}

	outsideEmail := myrtle.NewBuilder(defaulttheme.New()).
		WithHeader(
			myrtle.MessageBlock{Subject: "Outside header subject", Preview: "outside"},
			myrtle.HeaderPlacement(myrtle.HeaderPlacementOutside),
		).
		AddText("Body").
		Build()

	outsideHTML, err := outsideEmail.HTML()
	if err != nil {
		t.Fatalf("outside html returned error: %v", err)
	}
	outsideHeaderIndex := strings.Index(outsideHTML, "Outside header subject")
	if outsideHeaderIndex == -1 {
		t.Fatalf("expected outside header to render")
	}
	outsideStart := outsideHeaderIndex - 1200
	if outsideStart < 0 {
		outsideStart = 0
	}
	outsideSnippet := outsideHTML[outsideStart:outsideHeaderIndex]
	if !strings.Contains(outsideSnippet, "padding:0 24px;") {
		t.Fatalf("expected outside header default inset to use outside content inset")
	}
}

func TestHeaderBlockStartAlignment(t *testing.T) {
	t.Parallel()
	builder := myrtle.NewBuilder(defaulttheme.New())

	email := builder.
		WithHeader(myrtle.TextBlock{Text: "Left aligned", Align: myrtle.TextAlignStart}).
		AddText("Body").
		Build()

	html, err := email.HTML()
	if err != nil {
		t.Fatalf("html returned error: %v", err)
	}

	if !strings.Contains(html, "Left aligned") || !strings.Contains(html, "text-align:left;") {
		t.Fatalf("expected start-aligned header block in html")
	}
}

func TestRTLDirectionUsesLogicalStartEnd(t *testing.T) {
	t.Parallel()
	builder := myrtle.NewBuilder(defaulttheme.New(), myrtle.WithDirection(theme.DirectionRTL))

	email := builder.
		WithHeader(myrtle.TextBlock{Text: "RTL aligned", Align: myrtle.TextAlignStart}).
		AddText("Logical start text", myrtle.TextAlign(myrtle.TextAlignStart)).
		AddText("Logical end text", myrtle.TextAlign(myrtle.TextAlignEnd)).
		AddButton("Start button", "https://example.com/start", myrtle.ButtonAlign(myrtle.ButtonAlignmentStart)).
		AddButton("End button", "https://example.com/end", myrtle.ButtonAlign(myrtle.ButtonAlignmentEnd)).
		Build()

	html, err := email.HTML()
	if err != nil {
		t.Fatalf("html returned error: %v", err)
	}

	if !strings.Contains(html, `dir="rtl"`) {
		t.Fatalf("expected rtl direction attributes in rendered html")
	}
	if !strings.Contains(html, "RTL aligned") || !strings.Contains(html, "text-align:right;") {
		t.Fatalf("expected header logical start alignment to map to right in rtl")
	}
	if !strings.Contains(html, "text-align:right;") {
		t.Fatalf("expected logical start alignment to map to right in rtl")
	}
	if !strings.Contains(html, "text-align:left;") {
		t.Fatalf("expected logical end alignment to map to left in rtl")
	}
}

func TestRTLTableNumericAlignsToVisualEnd(t *testing.T) {
	t.Parallel()
	builder := myrtle.NewBuilder(defaulttheme.New(), myrtle.WithDirection(theme.DirectionRTL))

	email := builder.
		AddTable(
			[]string{"Name", "Value"},
			[][]string{{"Users", "1200"}},
			myrtle.TableRightAlignNumericColumns(true),
		).
		Build()

	html, err := email.HTML()
	if err != nil {
		t.Fatalf("html returned error: %v", err)
	}

	if !strings.Contains(html, `dir="rtl"`) {
		t.Fatalf("expected rtl direction attributes in rendered html")
	}
	if !strings.Contains(html, "text-align:left;\">1200</td>") {
		t.Fatalf("expected numeric column alignment to map to visual end (left) in rtl")
	}
}

func TestRTLProgressValueAlignsToVisualEnd(t *testing.T) {
	t.Parallel()
	builder := myrtle.NewBuilder(defaulttheme.New(), myrtle.WithDirection(theme.DirectionRTL))

	email := builder.
		AddProgress("Completion", []myrtle.ProgressItem{{Label: "Adoption", Value: "75%", Percent: 75}}).
		Build()

	html, err := email.HTML()
	if err != nil {
		t.Fatalf("html returned error: %v", err)
	}

	if !strings.Contains(html, `dir="rtl"`) {
		t.Fatalf("expected rtl direction attributes in rendered html")
	}
	if !strings.Contains(html, "text-align:left;white-space:nowrap;vertical-align:top;\">75%</td>") {
		t.Fatalf("expected progress value alignment to map to visual end (left) in rtl")
	}
}

func TestRTLStackedBarLegendAlignsToVisualEnd(t *testing.T) {
	t.Parallel()
	builder := myrtle.NewBuilder(defaulttheme.New(), myrtle.WithDirection(theme.DirectionRTL))

	email := builder.
		AddStackedBar("Pipeline", []myrtle.StackedBarRow{{
			Label:    "Channel Mix",
			Segments: []myrtle.StackedBarSegment{{Label: "Email", Value: "40%", Percent: 40}},
		}}).
		Build()

	html, err := email.HTML()
	if err != nil {
		t.Fatalf("html returned error: %v", err)
	}

	if !strings.Contains(html, `dir="rtl"`) {
		t.Fatalf("expected rtl direction attributes in rendered html")
	}
	if !strings.Contains(html, "font-size:12px;text-align:left;") {
		t.Fatalf("expected stacked bar legend alignment to map to visual end (left) in rtl")
	}
}

func TestNewBuilderWithHeaderOptions(t *testing.T) {
	t.Parallel()
	builder := myrtle.NewBuilder(
		defaulttheme.New(),
		myrtle.WithHeader(myrtle.HeadingBlock{Text: "Configured in constructor", Level: 1}, myrtle.HeaderRenderInText(true)),
	)

	email := builder.WithPreheader("Set via NewBuilder option").AddText("Body").Build()

	html, err := email.HTML()
	if err != nil {
		t.Fatalf("html returned error: %v", err)
	}
	if !strings.Contains(html, "Configured in constructor") {
		t.Fatalf("expected header block from constructor option")
	}

	text, err := email.Text()
	if err != nil {
		t.Fatalf("text returned error: %v", err)
	}
	if !strings.Contains(text, "Set via NewBuilder option") {
		t.Fatalf("expected preheader from constructor header options")
	}
}

func TestWithPreheaderPaddingRepeatOption(t *testing.T) {
	t.Parallel()
	builder := myrtle.NewBuilder(defaulttheme.New())

	email := builder.
		WithPreheader("Preview", myrtle.PreheaderPaddingRepeat(3)).
		AddText("Body").
		Build()

	html, err := email.HTML()
	if err != nil {
		t.Fatalf("html returned error: %v", err)
	}

	if count := strings.Count(html, "&nbsp;&zwnj;"); count != 3 {
		t.Fatalf("expected exactly 3 hidden preheader filler pairs, got %d", count)
	}

	text, err := email.Text()
	if err != nil {
		t.Fatalf("text returned error: %v", err)
	}

	if !strings.Contains(text, "Preview") {
		t.Fatalf("expected preheader text in plain-text output")
	}
	if strings.Contains(text, "&nbsp;&zwnj;") {
		t.Fatalf("expected hidden html preheader filler not to leak into plain-text output")
	}
}

func TestHeaderBlockNotRenderedInTextByDefault(t *testing.T) {
	t.Parallel()
	builder := myrtle.NewBuilder(defaulttheme.New())

	email := builder.
		WithHeader(myrtle.HeadingBlock{Text: "Hidden by default", Level: 1}).
		AddText("Body").
		Build()

	html, err := email.HTML()
	if err != nil {
		t.Fatalf("html returned error: %v", err)
	}
	if !strings.Contains(html, "Hidden by default") {
		t.Fatalf("expected header block to render in html")
	}

	text, err := email.Text()
	if err != nil {
		t.Fatalf("text returned error: %v", err)
	}
	if strings.Contains(text, "Hidden by default\n") {
		t.Fatalf("expected header block not to render in text by default")
	}
}

func TestHeaderBlockRenderedInTextWhenOptedIn(t *testing.T) {
	t.Parallel()
	builder := myrtle.NewBuilder(defaulttheme.New())

	email := builder.
		WithHeader(
			myrtle.HeadingBlock{Text: "Visible when opted in", Level: 1},
			myrtle.HeaderRenderInText(true),
		).
		AddText("Body").
		Build()

	html, err := email.HTML()
	if err != nil {
		t.Fatalf("html returned error: %v", err)
	}
	if !strings.Contains(html, "Visible when opted in") {
		t.Fatalf("expected header block to render in html")
	}

	text, err := email.Text()
	if err != nil {
		t.Fatalf("text returned error: %v", err)
	}
	if !strings.Contains(text, "Visible when opted in") {
		t.Fatalf("expected title to render in text header when opted in")
	}
	if !strings.Contains(text, "Body") {
		t.Fatalf("expected header block to render in text when enabled")
	}
}

func TestTextHeaderNotRenderedByDefault(t *testing.T) {
	t.Parallel()
	builder := myrtle.NewBuilder(defaulttheme.New())

	email := builder.
		WithPreheader("Header preheader").
		WithHeader(myrtle.HeadingBlock{Text: "Header title", Level: 1}).
		AddText("Body text").
		Build()

	text, err := email.Text()
	if err != nil {
		t.Fatalf("text returned error: %v", err)
	}

	if strings.Contains(text, "Header title\n") {
		t.Fatalf("expected title to be hidden in text by default")
	}
	if !strings.Contains(text, "Header preheader") {
		t.Fatalf("expected preheader to render in text independently of header settings")
	}
	if !strings.Contains(text, "Body text") {
		t.Fatalf("expected text body content to render")
	}
}

func TestTextHeaderRenderedWhenEnabled(t *testing.T) {
	t.Parallel()
	builder := myrtle.NewBuilder(defaulttheme.New())

	email := builder.
		WithPreheader("Header preheader").
		WithHeader(
			myrtle.HeadingBlock{Text: "Header title", Level: 1},
			myrtle.HeaderRenderInText(true),
		).
		AddText("Body text").
		Build()

	text, err := email.Text()
	if err != nil {
		t.Fatalf("text returned error: %v", err)
	}

	if !strings.Contains(text, "Header title") {
		t.Fatalf("expected title to render in text when enabled")
	}
	if !strings.Contains(text, "Header preheader") {
		t.Fatalf("expected preheader to render in text when enabled")
	}
}

func TestFooterBlockRendersInHTML(t *testing.T) {
	t.Parallel()
	builder := myrtle.NewBuilder(defaulttheme.New())

	email := builder.
		AddText("Body").
		WithFooter(myrtle.HeadingBlock{Text: "Footer heading", Level: 2}).
		Build()

	html, err := email.HTML()
	if err != nil {
		t.Fatalf("html returned error: %v", err)
	}

	if !strings.Contains(html, "Footer heading") {
		t.Fatalf("expected footer block to render in html")
	}
}

func TestFooterBlockNotRenderedInTextByDefault(t *testing.T) {
	t.Parallel()
	builder := myrtle.NewBuilder(defaulttheme.New())

	email := builder.
		AddText("Body").
		WithFooter(myrtle.TextBlock{Text: "Hidden footer"}).
		Build()

	text, err := email.Text()
	if err != nil {
		t.Fatalf("text returned error: %v", err)
	}

	if strings.Contains(text, "Hidden footer") {
		t.Fatalf("expected footer block not to render in text by default")
	}
}

func TestFooterBlockRenderedInTextWhenOptedIn(t *testing.T) {
	t.Parallel()
	builder := myrtle.NewBuilder(defaulttheme.New())

	email := builder.
		AddText("Body").
		WithFooter(myrtle.TextBlock{Text: "Visible footer"}, myrtle.FooterRenderInText(true)).
		Build()

	text, err := email.Text()
	if err != nil {
		t.Fatalf("text returned error: %v", err)
	}

	if !strings.Contains(text, "Visible footer") {
		t.Fatalf("expected footer block to render in text when enabled")
	}
}

func TestFooterCanRenderOutsideMainContentBox(t *testing.T) {
	t.Parallel()
	builder := myrtle.NewBuilder(defaulttheme.New())

	email := builder.
		AddText("Body").
		WithFooter(myrtle.TextBlock{Text: "Outside footer"}, myrtle.FooterPlacement(myrtle.FooterPlacementOutside)).
		Build()

	html, err := email.HTML()
	if err != nil {
		t.Fatalf("html returned error: %v", err)
	}

	containerIndex := strings.Index(html, "max-width:640px;margin:0 auto;mso-table-lspace:0pt;mso-table-rspace:0pt;border-collapse:separate;border-spacing:0;")
	footerIndex := strings.LastIndex(html, "Outside footer")
	if containerIndex == -1 || footerIndex == -1 {
		t.Fatalf("expected both main container and footer to be rendered")
	}
	if footerIndex < containerIndex {
		t.Fatalf("expected outside footer to render after the main content container")
	}
}

func TestFlatAndTerminalSupportOutsideHeaderAndFooter(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		theme theme.Theme
	}{
		{name: "flat", theme: flat.New()},
		{name: "terminal", theme: terminal.New()},
	}

	for _, testCase := range tests {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			builder := myrtle.NewBuilder(testCase.theme)
			email := builder.
				WithHeader(myrtle.TextBlock{Text: "Outside header", Align: myrtle.TextAlignCenter}, myrtle.HeaderPlacement(myrtle.HeaderPlacementOutside)).
				AddText("Body").
				WithFooter(myrtle.TextBlock{Text: "Outside footer", Align: myrtle.TextAlignCenter}, myrtle.FooterPlacement(myrtle.FooterPlacementOutside)).
				Build()

			html, err := email.HTML()
			if err != nil {
				t.Fatalf("html returned error: %v", err)
			}

			containerNeedle := "max-width:640px;margin:0 auto;word-break:break-word;word-wrap:break-word;background:"
			if testCase.name == "terminal" {
				containerNeedle = "max-width:640px;margin:0 auto;mso-table-lspace:0pt;mso-table-rspace:0pt;border-collapse:separate;border-spacing:0;word-break:break-word;word-wrap:break-word;"
			}
			containerIndex := strings.Index(html, containerNeedle)
			headerIndex := strings.Index(html, "Outside header")
			footerIndex := strings.LastIndex(html, "Outside footer")
			if containerIndex == -1 || headerIndex == -1 || footerIndex == -1 {
				t.Fatalf("expected container, outside header, and outside footer to render")
			}
			if headerIndex > containerIndex {
				t.Fatalf("expected outside header to render before the main content container")
			}
			if footerIndex < containerIndex {
				t.Fatalf("expected outside footer to render after the main content container")
			}
		})
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

func TestHeroPrimaryToneRendersInvertedStyles(t *testing.T) {
	t.Parallel()

	email := myrtle.NewBuilder(defaulttheme.New()).
		AddHero(
			"Hero title",
			"Hero body",
			"Start",
			"https://example.com/start",
			myrtle.HeroTone(myrtle.TonePrimary),
		).
		Build()

	html, err := email.HTML()
	if err != nil {
		t.Fatalf("html returned error: %v", err)
	}

	for _, needle := range []string{
		"background:#265cff;",
		"color:#ffffff;",
		"background:#ffffff;color:#265cff;",
	} {
		if !strings.Contains(html, needle) {
			t.Fatalf("expected html to contain %q", needle)
		}
	}
}

func TestSummaryCardAndEmptyStatePrimaryToneRenderInvertedStyles(t *testing.T) {
	t.Parallel()

	email := myrtle.NewBuilder(defaulttheme.New()).
		AddSummaryCard(
			"Deployment complete",
			"No customer impact detected.",
			"Updated 5 minutes ago",
			myrtle.SummaryCardTone(myrtle.TonePrimary),
		).
		AddEmptyState(
			"No incidents",
			"Everything looks healthy right now.",
			"View dashboard",
			"https://example.com/dashboard",
			myrtle.EmptyStateTone(myrtle.TonePrimary),
		).
		Build()

	html, err := email.HTML()
	if err != nil {
		t.Fatalf("html returned error: %v", err)
	}

	for _, needle := range []string{
		"background:#265cff;",
		"color:#ffffff;",
		"No incidents",
		"Deployment complete",
	} {
		if !strings.Contains(html, needle) {
			t.Fatalf("expected html to contain %q", needle)
		}
	}
}

func TestAdvancedLayoutAndPrimitiveVariantsRender(t *testing.T) {
	t.Parallel()
	builder := myrtle.NewBuilder(defaulttheme.New())

	email := builder.
		AddPanel(
			myrtle.NewGroup().
				AddText("Section body").
				AddButton("Open section", "https://example.com/section", myrtle.ButtonStyle(myrtle.ButtonStyleOutline)),
			myrtle.PanelTitle("Section title"),
			myrtle.PanelSubtitle("Section subtitle"),
			myrtle.PanelPadding(20),
		).
		AddGrid(
			[]myrtle.GridItem{
				{Content: myrtle.TextBlock{Text: "Grid 1"}},
				{Content: myrtle.TextBlock{Text: "Grid 2"}},
				{Content: myrtle.TextBlock{Text: "Grid 3"}},
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
			myrtle.NewGroup().AddText("Left"),
			myrtle.NewGroup().AddText("Right"),
			myrtle.ColumnsGap(20),
			myrtle.ColumnsAlign(myrtle.ColumnsVerticalAlignMiddle),
		).
		AddDivider(myrtle.DividerStyle(myrtle.DividerVariantDashed), myrtle.DividerThickness(2), myrtle.DividerInset(16)).
		AddDivider(myrtle.DividerLabel("OR"), myrtle.DividerStyle(myrtle.DividerVariantDotted)).
		AddSpacer(myrtle.SpacerSize(24)).
		Build()

	html, err := email.HTML()
	if err != nil {
		t.Fatalf("html returned error: %v", err)
	}

	if !strings.Contains(html, "Section title") || !strings.Contains(html, "Section subtitle") {
		t.Fatalf("expected section block heading and subtitle to render")
	}
	if !strings.Contains(html, "padding-right:14px;") || !strings.Contains(html, "padding-bottom:14px;") {
		t.Fatalf("expected grid gap configuration to render")
	}
	if !strings.Contains(html, "Grid 3") {
		t.Fatalf("expected wrapped grid item content to render")
	}
	if !strings.Contains(html, "Card one") || !strings.Contains(html, "https://example.com/card1") {
		t.Fatalf("expected card list item content and link to render")
	}
	if !strings.Contains(html, "padding:0;padding-right:5px;") || !strings.Contains(html, "padding:0;padding-left:5px;") {
		t.Fatalf("expected card list gap to be split across adjacent cards")
	}
	if !strings.Contains(html, "valign=\"middle\"") || !strings.Contains(html, "padding:0;padding-right:10px;") || !strings.Contains(html, "padding:0;padding-left:10px;") {
		t.Fatalf("expected columns alignment and gap styles to render")
	}
	if !strings.Contains(html, "border-top:2px dashed") || !strings.Contains(html, "padding:0 16px;") {
		t.Fatalf("expected divider variant, thickness, and inset to render")
	}
	if !strings.Contains(html, ">OR<") || !strings.Contains(html, "text-transform:uppercase") {
		t.Fatalf("expected labeled divider to render centered label")
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
		AddColumns(left, right, myrtle.ColumnsWidths(55, 45)).
		AddPanel(section, myrtle.PanelTitle("Grouped section")).
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

	text, err := email.Text()
	if err != nil {
		t.Fatalf("text returned error: %v", err)
	}

	if !strings.Contains(text, "Grouped heading") || !strings.Contains(text, "Grouped text") || !strings.Contains(text, "Grouped CTA (https://example.com/group)") {
		t.Fatalf("expected grouped block content to render in text fallback")
	}
}

func TestButtonCalloutAndTableVariantsRender(t *testing.T) {
	t.Parallel()
	builder := myrtle.NewBuilder(
		defaulttheme.New(),
		myrtle.WithStyles(theme.Styles{ColorPrimary: "#2563eb", ColorSecondary: "#a855f7"}),
	)

	email := builder.
		AddButton("Secondary", "https://example.com", myrtle.ButtonTone(myrtle.ToneSecondary), myrtle.ButtonAlign(myrtle.ButtonAlignmentEnd), myrtle.ButtonFullWidth(true)).
		AddButton("Compact nowrap", "https://example.com/compact", myrtle.ButtonStyle(myrtle.ButtonStyleOutline), myrtle.ButtonSize(myrtle.ButtonSizeSmall), myrtle.ButtonNoWrap(true)).
		AddButton("Outline", "https://example.com/outline", myrtle.ButtonStyle(myrtle.ButtonStyleOutline)).
		AddButtonGroup([]myrtle.ButtonGroupButton{{Label: "Approve", URL: "https://example.com/approve", Tone: myrtle.TonePrimary}, {Label: "Review", URL: "https://example.com/review", Tone: myrtle.ToneSecondary}}, myrtle.ButtonGroupAlign(myrtle.ButtonAlignmentCenter), myrtle.ButtonGroupJoined(true), myrtle.ButtonGroupFullWidthOnMobile(true)).
		AddButtonGroup([]myrtle.ButtonGroupButton{{Label: "Outline", URL: "https://example.com/outline-group", Style: myrtle.ButtonStyleOutline}, {Label: "Ghost", URL: "https://example.com/ghost-group", Style: myrtle.ButtonStyleGhost}}, myrtle.ButtonGroupAlign(myrtle.ButtonAlignmentStart), myrtle.ButtonGroupGap(14), myrtle.ButtonGroupStackOnMobile(true)).
		AddCallout(myrtle.ToneDanger, "Critical", "Body", myrtle.CalloutStyle(myrtle.CalloutVariantSolid), myrtle.CalloutLink("Investigate", "https://example.com/investigate")).
		AddTable([]string{"Name", "Value"}, [][]string{{"Users", "1200"}, {"Rate", "4.2%"}}, myrtle.TableTitle("Metrics"), myrtle.TableZebraRows(true), myrtle.TableCompact(true), myrtle.TableRightAlignNumericColumns(true)).
		AddTable([]string{"Name", "Value"}, [][]string{{"Users", "1200"}, {"Rate", "4.2%"}}, myrtle.TableDensity(myrtle.TableDensityRelaxed), myrtle.TableHeaderTone(myrtle.TableHeaderToneMuted), myrtle.TableBorderStyle(myrtle.TableBorderStyleDashed)).
		AddTable([]string{"Name", "Value"}, [][]string{{"Users", "1200"}, {"Rate", "4.2%"}}, myrtle.TableHeaderTone(myrtle.TableHeaderTonePlain), myrtle.TableBorderStyle(myrtle.TableBorderStyleDotted)).
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
	if !strings.Contains(html, "<div style=\"margin:0 0 20px 0;text-align:right;\">") {
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
	if !strings.Contains(html, "padding-left:7px;") || !strings.Contains(html, "padding-right:7px;") {
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
	if !strings.Contains(html, "text-align:right;\">Value</th>") {
		t.Fatalf("expected numeric column header to align to visual end when numeric alignment is enabled")
	}
	if !strings.Contains(html, "background:#f8fafc;") {
		t.Fatalf("expected zebra row styling in table")
	}
	if !strings.Contains(html, "<th colspan=\"2\" align=\"center\"") || !strings.Contains(html, ">Metrics</th>") {
		t.Fatalf("expected table title row to render in a centered top bar")
	}
	if !strings.Contains(html, "padding:12px 14px;") {
		t.Fatalf("expected relaxed table density to render")
	}
	if !strings.Contains(html, "background:#f8fafc;color:#111827;") {
		t.Fatalf("expected muted table header tone to render")
	}
	if !strings.Contains(html, "border-bottom:1px dashed") || !strings.Contains(html, "border-bottom:2px dotted") {
		t.Fatalf("expected customized table border styles to render")
	}

	dotsEmail := myrtle.NewBuilder(defaulttheme.New()).
		AddTable(
			[]string{"Series", "Value"},
			[][]string{{"Email", "68%"}, {"SMS", "21%"}, {"Push", "11%"}},
			myrtle.TableLegendSwatches([]string{"#2563eb", "#7c3aed"}),
		).
		Build()

	dotsHTML, err := dotsEmail.HTML()
	if err != nil {
		t.Fatalf("html returned error for row color dots table: %v", err)
	}
	if !strings.Contains(dotsHTML, "width:11px;height:11px;border-radius:3px;background:#2563eb;") || !strings.Contains(dotsHTML, "width:11px;height:11px;border-radius:3px;background:#7c3aed;") {
		t.Fatalf("expected row color dots to render for configured rows")
	}
	if strings.Contains(dotsHTML, "width:11px;height:11px;border-radius:3px;background:#0ea5e9;") {
		t.Fatalf("expected missing row color dot entries to render no dot")
	}
}

func TestHorizontalBarChartThicknessAndTransparentBackgroundRender(t *testing.T) {
	t.Parallel()
	builder := myrtle.NewBuilder(defaulttheme.New())

	email := builder.
		AddHorizontalBarChart("Delivery", []myrtle.HorizontalBarChartItem{{Label: "US", Percent: 52}}, myrtle.HorizontalBarChartThickness(14), myrtle.HorizontalBarChartTransparentBackground(true), myrtle.HorizontalBarChartTone(myrtle.ToneWarning)).
		AddSparkline("Trend", "Rate", "4.1%", []int{8, 12, 9, 14, 18, 16, 20}, myrtle.SparklineTone(myrtle.ToneSuccess)).
		AddStackedBar("Stage mix", []myrtle.StackedBarRow{{Label: "Acquisition", Segments: []myrtle.StackedBarSegment{{Label: "Email", Percent: 58, Value: "58%"}, {Label: "SMS", Percent: 24, Value: "24%"}, {Label: "Push", Percent: 18, Value: "18%"}}}}).
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
	if !strings.Contains(html, "background:#eca40f;") {
		t.Fatalf("expected warning tone to color bar chart")
	}
	if !strings.Contains(html, "background:#16a34a;") {
		t.Fatalf("expected success tone to color sparkline")
	}
	if !strings.Contains(html, "background:#265cff;") || !strings.Contains(html, "background:#10b981;") {
		t.Fatalf("expected stacked bar defaults to include primary and secondary segment colors")
	}
	if strings.Contains(html, "opacity:") {
		t.Fatalf("expected stacked bar to render without opacity-based styling")
	}
}

func TestHorizontalBarChartCanRenderLabelsInsideBars(t *testing.T) {
	t.Parallel()
	builder := myrtle.NewBuilder(defaulttheme.New())

	email := builder.
		AddHorizontalBarChart(
			"Delivery",
			[]myrtle.HorizontalBarChartItem{{Label: "US", Value: "52%", Percent: 52}},
			myrtle.HorizontalBarChartLabelsInsideBars(true),
		).
		Build()

	html, err := email.HTML()
	if err != nil {
		t.Fatalf("html returned error: %v", err)
	}

	if !strings.Contains(html, "font-size:11px;line-height:1.2;font-weight:600;white-space:nowrap;overflow:hidden;text-overflow:ellipsis;\">US<") {
		t.Fatalf("expected item label to render inside bar")
	}
	if !strings.Contains(html, "font-size:11px;line-height:1.2;font-weight:700;color:") || !strings.Contains(html, ">52%<") {
		t.Fatalf("expected item value to render next to the bar")
	}
	if strings.Contains(html, "ColorTextOnSolid }};font-size:11px;line-height:1.2;font-weight:700") {
		t.Fatalf("expected item value to no longer render inside the filled bar")
	}
	if strings.Contains(html, "padding:0 0 4px;font-size:13px") {
		t.Fatalf("expected top label/value row to be omitted when rendering labels inside bars")
	}
	if !strings.Contains(html, "height:18px;line-height:18px") {
		t.Fatalf("expected thickness to be increased to minimum readable size for inside labels")
	}
}

func TestDatavizPerItemColorsRender(t *testing.T) {
	t.Parallel()
	builder := myrtle.NewBuilder(defaulttheme.New())

	email := builder.
		AddHorizontalBarChart("Delivery", []myrtle.HorizontalBarChartItem{{Label: "US", Percent: 52, Color: "#123456"}}).
		AddProgress("Rollout", []myrtle.ProgressItem{{Label: "Deploy", Percent: 80, Color: "#234567"}}).
		AddStackedBar("Funnel", []myrtle.StackedBarRow{{Label: "Q1", Segments: []myrtle.StackedBarSegment{{Label: "Won", Percent: 60, Color: "#345678"}}}}).
		AddDistribution("Latency", []myrtle.DistributionBucket{{Label: "0-50", Count: 10, Color: "#456789"}}).
		Build()

	html, err := email.HTML()
	if err != nil {
		t.Fatalf("html returned error: %v", err)
	}

	for _, color := range []string{"#123456", "#234567", "#345678", "#456789"} {
		if !strings.Contains(html, "background:"+color) {
			t.Fatalf("expected custom color %q to render in html", color)
		}
	}
}

func TestDistributionUsesFixedCountColumnWidth(t *testing.T) {
	t.Parallel()
	builder := myrtle.NewBuilder(defaulttheme.New())

	email := builder.
		AddDistribution("Latency", []myrtle.DistributionBucket{
			{Label: "0-50", Count: 8},
			{Label: "51-100", Count: 1200},
		}).
		Build()

	html, err := email.HTML()
	if err != nil {
		t.Fatalf("html returned error: %v", err)
	}

	if strings.Count(html, "width:4ch;") != 2 {
		t.Fatalf("expected fixed count-column width derived from max count digits across all rows")
	}
}

func TestVerticalBarChartLegendAxisAndNegativeRender(t *testing.T) {
	t.Parallel()
	builder := myrtle.NewBuilder(defaulttheme.New())

	series := []myrtle.VerticalBarChartSeries{
		{Key: "new", Label: "New", Color: "#2563eb", Values: []float64{42, 36}},
		{Key: "churn", Label: "Churn", Color: "#dc2626", Values: []float64{-11, -8}},
	}

	email := builder.
		AddVerticalBarChart(
			[]string{"Jan", "Feb"},
			series,
			myrtle.VerticalBarChartTitle("MRR movement"),
			myrtle.VerticalBarChartSubtitle("Net new and churn"),
			myrtle.VerticalBarChartLegendPlacement(myrtle.VerticalBarChartLegendBottom),
			myrtle.VerticalBarChartAxisShowYTicks(true),
			myrtle.VerticalBarChartAxisShowBaseline(true),
		).
		Build()

	html, err := email.HTML()
	if err != nil {
		t.Fatalf("html returned error: %v", err)
	}

	if !strings.Contains(html, `data-myrtle-vertical-bar-chart="1"`) {
		t.Fatalf("expected vertical bar chart container marker")
	}
	if !strings.Contains(html, "MRR movement") || !strings.Contains(html, "Net new and churn") || !strings.Contains(html, "text-align:center;") {
		t.Fatalf("expected centered title and subtitle to render")
	}
	if !strings.Contains(html, "Legend") && !strings.Contains(html, "New") {
		t.Fatalf("expected legend labels to render")
	}
	if !strings.Contains(html, "border-right:1px solid") {
		t.Fatalf("expected y-axis line to render beside the chart")
	}
	if !strings.Contains(html, ">42<") || !strings.Contains(html, ">0<") {
		t.Fatalf("expected y-axis to show max and zero labels")
	}
	if !strings.Contains(html, "title=\"Churn: -11\"") {
		t.Fatalf("expected negative segment title to render")
	}
	if !strings.Contains(html, "height:1px;line-height:1px") {
		t.Fatalf("expected baseline row to render")
	}
}

func TestVerticalBarChartCanHideCategoryLabels(t *testing.T) {
	t.Parallel()
	builder := myrtle.NewBuilder(defaulttheme.New())

	email := builder.
		AddVerticalBarChart(
			[]string{"Jan"},
			[]myrtle.VerticalBarChartSeries{{Key: "new", Label: "New", Values: []float64{20}}},
			myrtle.VerticalBarChartAxisShowCategoryLabels(false),
		).
		Build()

	html, err := email.HTML()
	if err != nil {
		t.Fatalf("html returned error: %v", err)
	}

	if strings.Contains(html, ">Jan<") {
		t.Fatalf("expected category labels to be hidden when disabled")
	}
}

func TestVerticalBarChartDrawsYAxisLineInInsetNoneByDefaultWhenYTicksEnabled(t *testing.T) {
	t.Parallel()
	builder := myrtle.NewBuilder(defaulttheme.New())

	email := builder.
		AddVerticalBarChart(
			[]string{"Jan", "Feb"},
			[]myrtle.VerticalBarChartSeries{{Key: "new", Label: "New", Values: []float64{20, 24}}},
			myrtle.VerticalBarChartInsetMode(myrtle.InsetModeNone),
			myrtle.VerticalBarChartAxisShowYTicks(true),
		).
		Build()

	html, err := email.HTML()
	if err != nil {
		t.Fatalf("html returned error: %v", err)
	}

	if !strings.Contains(html, "border-right:1px solid") {
		t.Fatalf("expected y-axis line to render in inset none mode when y-ticks are enabled")
	}
}

func TestVerticalBarChartCanDisableYAxisLineExplicitly(t *testing.T) {
	t.Parallel()
	builder := myrtle.NewBuilder(defaulttheme.New())

	email := builder.
		AddVerticalBarChart(
			[]string{"Jan", "Feb"},
			[]myrtle.VerticalBarChartSeries{{Key: "new", Label: "New", Values: []float64{20, 24}}},
			myrtle.VerticalBarChartAxisShowYTicks(true),
			myrtle.VerticalBarChartAxisDrawYAxisLine(false),
		).
		Build()

	html, err := email.HTML()
	if err != nil {
		t.Fatalf("html returned error: %v", err)
	}

	if strings.Contains(html, "border-right:1px solid") {
		t.Fatalf("expected y-axis line to be hidden when explicitly disabled")
	}
}

func TestVerticalBarChartYAxisLineUsesEndBorderInLTR(t *testing.T) {
	t.Parallel()
	builder := myrtle.NewBuilder(defaulttheme.New())

	email := builder.
		AddVerticalBarChart(
			[]string{"Jan", "Feb"},
			[]myrtle.VerticalBarChartSeries{{Key: "new", Label: "New", Values: []float64{20, 24}}},
			myrtle.VerticalBarChartAxisShowYTicks(true),
			myrtle.VerticalBarChartAxisDrawYAxisLine(true),
		).
		Build()

	html, err := email.HTML()
	if err != nil {
		t.Fatalf("html returned error: %v", err)
	}

	if !strings.Contains(html, "border-right:1px solid") {
		t.Fatalf("expected y-axis line to render on visual end side in ltr")
	}
}

func TestVerticalBarChartYAxisLineUsesEndBorderInRTL(t *testing.T) {
	t.Parallel()
	builder := myrtle.NewBuilder(defaulttheme.New(), myrtle.WithDirection(theme.DirectionRTL))

	email := builder.
		AddVerticalBarChart(
			[]string{"Jan", "Feb"},
			[]myrtle.VerticalBarChartSeries{{Key: "new", Label: "New", Values: []float64{20, 24}}},
			myrtle.VerticalBarChartAxisShowYTicks(true),
			myrtle.VerticalBarChartAxisDrawYAxisLine(true),
		).
		Build()

	html, err := email.HTML()
	if err != nil {
		t.Fatalf("html returned error: %v", err)
	}

	if !strings.Contains(html, "border-left:1px solid") {
		t.Fatalf("expected y-axis line to render on visual end side in rtl")
	}
}

func TestVerticalBarChartLegendSwatchUsesStyleTokensAndLogicalSpacing(t *testing.T) {
	t.Parallel()
	builderLTR := myrtle.NewBuilder(defaulttheme.New(), myrtle.WithStyles(theme.Styles{
		TableLegendSwatchSize:   "13px",
		TableLegendSwatchRadius: "3px",
		TableLegendSwatchBorder: "1px solid #222222",
	}))

	emailLTR := builderLTR.
		AddVerticalBarChart(
			[]string{"Jan", "Feb"},
			[]myrtle.VerticalBarChartSeries{{Key: "new", Label: "New", Color: "#2563eb", Values: []float64{20, 24}}},
			myrtle.VerticalBarChartLegendPlacement(myrtle.VerticalBarChartLegendBottom),
		).
		Build()

	htmlLTR, err := emailLTR.HTML()
	if err != nil {
		t.Fatalf("html returned error: %v", err)
	}

	if !strings.Contains(htmlLTR, "width:13px;height:13px;border-radius:3px;") {
		t.Fatalf("expected legend swatch to use style token size and radius")
	}
	if !strings.Contains(htmlLTR, "border:1px solid #222222;") {
		t.Fatalf("expected legend swatch to use style token border")
	}
	if !strings.Contains(htmlLTR, "margin-right:4px;") {
		t.Fatalf("expected legend swatch spacing to render on logical end side in ltr")
	}

	builderRTL := myrtle.NewBuilder(defaulttheme.New(),
		myrtle.WithDirection(theme.DirectionRTL),
		myrtle.WithStyles(theme.Styles{
			TableLegendSwatchSize:   "13px",
			TableLegendSwatchRadius: "3px",
			TableLegendSwatchBorder: "1px solid #222222",
		}),
	)

	emailRTL := builderRTL.
		AddVerticalBarChart(
			[]string{"Jan", "Feb"},
			[]myrtle.VerticalBarChartSeries{{Key: "new", Label: "New", Color: "#2563eb", Values: []float64{20, 24}}},
			myrtle.VerticalBarChartLegendPlacement(myrtle.VerticalBarChartLegendBottom),
		).
		Build()

	htmlRTL, err := emailRTL.HTML()
	if err != nil {
		t.Fatalf("html returned error: %v", err)
	}

	if !strings.Contains(htmlRTL, "margin-left:4px;") {
		t.Fatalf("expected legend swatch spacing to render on logical end side in rtl")
	}
}

func TestVerticalBarChartLegendStartOffsetIncludesYAxisWidthAndOuterGap(t *testing.T) {
	t.Parallel()
	builder := myrtle.NewBuilder(defaulttheme.New())

	email := builder.
		AddVerticalBarChart(
			[]string{"Jan", "Feb"},
			[]myrtle.VerticalBarChartSeries{{Key: "new", Label: "New", Values: []float64{42, 36}}},
			myrtle.VerticalBarChartLegendPlacement(myrtle.VerticalBarChartLegendBottom),
			myrtle.VerticalBarChartAxisShowYTicks(true),
			myrtle.VerticalBarChartOuterGap(4),
		).
		Build()

	html, err := email.HTML()
	if err != nil {
		t.Fatalf("html returned error: %v", err)
	}

	if !strings.Contains(html, "width:36px;padding:0;vertical-align:top;") {
		t.Fatalf("expected y-axis region width to render")
	}
	if !strings.Contains(html, "margin-left:40px;") {
		t.Fatalf("expected legend start offset to include y-axis width and outer gap")
	}
}

func TestVerticalBarChartDefaultsToFullWidthAndAdaptiveGap(t *testing.T) {
	t.Parallel()
	builder := myrtle.NewBuilder(defaulttheme.New())

	email := builder.
		AddVerticalBarChart(
			[]string{"Jan", "Feb"},
			[]myrtle.VerticalBarChartSeries{{Key: "new", Label: "New", Values: []float64{20, 24}}},
		).
		Build()

	html, err := email.HTML()
	if err != nil {
		t.Fatalf("html returned error: %v", err)
	}

	if strings.Count(html, "width:8px;padding:0;") < 3 {
		t.Fatalf("expected reduced adaptive default column gap to render between and around columns")
	}
	if !strings.Contains(html, "table-layout:fixed") {
		t.Fatalf("expected vertical chart columns to use fixed full-width layout")
	}
}

func TestVerticalBarChartCanDisableOuterGap(t *testing.T) {
	t.Parallel()
	builder := myrtle.NewBuilder(defaulttheme.New())

	email := builder.
		AddVerticalBarChart(
			[]string{"Jan", "Feb"},
			[]myrtle.VerticalBarChartSeries{{Key: "new", Label: "New", Values: []float64{20, 24}}},
			myrtle.VerticalBarChartOuterGap(0),
		).
		Build()

	html, err := email.HTML()
	if err != nil {
		t.Fatalf("html returned error: %v", err)
	}

	if strings.Count(html, "width:8px;padding:0;") != 1 {
		t.Fatalf("expected only inner column gap when outer gap is disabled")
	}
}

func TestVerticalBarChartCanSetCustomOuterGapPixels(t *testing.T) {
	t.Parallel()
	builder := myrtle.NewBuilder(defaulttheme.New())

	email := builder.
		AddVerticalBarChart(
			[]string{"Jan", "Feb"},
			[]myrtle.VerticalBarChartSeries{{Key: "new", Label: "New", Values: []float64{20, 24}}},
			myrtle.VerticalBarChartColumnGap(8),
			myrtle.VerticalBarChartOuterGap(4),
		).
		Build()

	html, err := email.HTML()
	if err != nil {
		t.Fatalf("html returned error: %v", err)
	}

	if strings.Count(html, "width:8px;padding:0;") != 1 {
		t.Fatalf("expected inner gap to stay at configured column gap")
	}
	if strings.Count(html, "width:4px;padding:0;") != 2 {
		t.Fatalf("expected outer gap to use configured pixel value on both sides")
	}
}

func TestVerticalBarChartDoesNotNormalizeByDefault(t *testing.T) {
	t.Parallel()
	builder := myrtle.NewBuilder(defaulttheme.New())

	email := builder.
		AddVerticalBarChart(
			[]string{"Jan", "Feb"},
			[]myrtle.VerticalBarChartSeries{{Key: "new", Label: "New", Values: []float64{100, 50}}},
			myrtle.VerticalBarChartHeight(100),
			myrtle.VerticalBarChartAxisMin(0),
			myrtle.VerticalBarChartAxisMax(100),
		).
		Build()

	html, err := email.HTML()
	if err != nil {
		t.Fatalf("html returned error: %v", err)
	}

	if !strings.Contains(html, "title=\"New: 50\"") {
		t.Fatalf("expected second column segment to render")
	}
	if !strings.Contains(html, "title=\"New: 50\" valign=\"middle\" style=\"height:50px;") {
		t.Fatalf("expected second column segment to keep proportional height without normalization")
	}
}

func TestVerticalBarChartCanNormalizeHeights(t *testing.T) {
	t.Parallel()
	builder := myrtle.NewBuilder(defaulttheme.New())

	email := builder.
		AddVerticalBarChart(
			[]string{"Jan", "Feb"},
			[]myrtle.VerticalBarChartSeries{{Key: "new", Label: "New", Values: []float64{100, 50}}},
			myrtle.VerticalBarChartHeight(100),
			myrtle.VerticalBarChartAxisMin(0),
			myrtle.VerticalBarChartAxisMax(100),
			myrtle.VerticalBarChartNormalize(true),
		).
		Build()

	html, err := email.HTML()
	if err != nil {
		t.Fatalf("html returned error: %v", err)
	}

	if !strings.Contains(html, "title=\"New: 50\" valign=\"middle\" style=\"height:100px;") {
		t.Fatalf("expected normalized second column segment to fill the positive region")
	}
}

func TestVerticalBarChartNormalizeFallsBackWhenNegativeValuesExist(t *testing.T) {
	t.Parallel()
	builder := myrtle.NewBuilder(defaulttheme.New())

	email := builder.
		AddVerticalBarChart(
			[]string{"Jan", "Feb"},
			[]myrtle.VerticalBarChartSeries{{Key: "new", Label: "New", Values: []float64{100, 50}}, {Key: "churn", Label: "Churn", Values: []float64{-20, -10}}},
			myrtle.VerticalBarChartHeight(100),
			myrtle.VerticalBarChartAxisMin(-30),
			myrtle.VerticalBarChartAxisMax(100),
			myrtle.VerticalBarChartNormalize(true),
		).
		Build()

	html, err := email.HTML()
	if err != nil {
		t.Fatalf("html returned error: %v", err)
	}

	if !strings.Contains(html, "title=\"New: 50\" valign=\"middle\" style=\"height:") {
		t.Fatalf("expected normalized mode to fall back to magnitude scaling for positive segments when negatives exist")
	}
	if !strings.Contains(html, "title=\"Churn: -10\" valign=\"middle\" style=\"height:") {
		t.Fatalf("expected normalized mode to fall back to magnitude scaling for negative segments when negatives exist")
	}
	if strings.Contains(html, "title=\"New: 50\" valign=\"middle\" style=\"height:77px;") {
		t.Fatalf("expected normalized fill behavior to be disabled when negative values are present")
	}
}

func TestVerticalBarChartAxisMaxDoesNotClampBelowData(t *testing.T) {
	t.Parallel()
	builder := myrtle.NewBuilder(defaulttheme.New())

	email := builder.
		AddVerticalBarChart(
			[]string{"Jan", "Feb"},
			[]myrtle.VerticalBarChartSeries{{Key: "new", Label: "New", Values: []float64{100, 50}}},
			myrtle.VerticalBarChartHeight(100),
			myrtle.VerticalBarChartAxisMin(0),
			myrtle.VerticalBarChartAxisMax(60),
		).
		Build()

	html, err := email.HTML()
	if err != nil {
		t.Fatalf("html returned error: %v", err)
	}

	if !strings.Contains(html, "title=\"New: 50\" valign=\"middle\" style=\"height:50px;") {
		t.Fatalf("expected axis max below observed values to be ignored for scaling")
	}
}

func TestVerticalBarChartCanRenderValueLabels(t *testing.T) {
	t.Parallel()
	builder := myrtle.NewBuilder(defaulttheme.New())

	email := builder.
		AddVerticalBarChart(
			[]string{"Jan", "Feb"},
			[]myrtle.VerticalBarChartSeries{{Key: "new", Label: "New", Values: []float64{100, 50}}},
			myrtle.VerticalBarChartHeight(100),
			myrtle.VerticalBarChartAxisMin(0),
			myrtle.VerticalBarChartAxisMax(100),
			myrtle.VerticalBarChartValueLabelsOption(myrtle.VerticalBarChartValueLabels{Show: true, MinSegmentHeight: 12}),
		).
		Build()

	html, err := email.HTML()
	if err != nil {
		t.Fatalf("html returned error: %v", err)
	}

	if !strings.Contains(html, ">50<") {
		t.Fatalf("expected value label to render inside sufficiently tall segment")
	}
}

func TestVerticalBarChartRendersValueLabelsAboveThinPositiveSegmentsWhenSpaceAllows(t *testing.T) {
	t.Parallel()
	builder := myrtle.NewBuilder(defaulttheme.New())

	email := builder.
		AddVerticalBarChart(
			[]string{"Jan", "Feb"},
			[]myrtle.VerticalBarChartSeries{{Key: "new", Label: "New", Values: []float64{100, 50}}},
			myrtle.VerticalBarChartHeight(100),
			myrtle.VerticalBarChartAxisMin(0),
			myrtle.VerticalBarChartAxisMax(100),
			myrtle.VerticalBarChartValueLabelsOption(myrtle.VerticalBarChartValueLabels{Show: true, MinSegmentHeight: 60}),
		).
		Build()

	html, err := email.HTML()
	if err != nil {
		t.Fatalf("html returned error: %v", err)
	}

	if !strings.Contains(html, ">50<") {
		t.Fatalf("expected thin positive value label to render above the segment when there is free space")
	}
}

func TestVerticalBarChartDoesNotRenderAboveLabelWithoutFreeSpace(t *testing.T) {
	t.Parallel()
	builder := myrtle.NewBuilder(defaulttheme.New())

	email := builder.
		AddVerticalBarChart(
			[]string{"Jan"},
			[]myrtle.VerticalBarChartSeries{
				{Key: "new", Label: "New", Values: []float64{5}},
				{Key: "expansion", Label: "Expansion", Values: []float64{95}},
			},
			myrtle.VerticalBarChartHeight(100),
			myrtle.VerticalBarChartAxisMin(0),
			myrtle.VerticalBarChartAxisMax(100),
			myrtle.VerticalBarChartValueLabelsOption(myrtle.VerticalBarChartValueLabels{Show: true, MinSegmentHeight: 60}),
		).
		Build()

	html, err := email.HTML()
	if err != nil {
		t.Fatalf("html returned error: %v", err)
	}

	if strings.Contains(html, ">5<") {
		t.Fatalf("expected thin positive label to remain hidden when no space exists above stack")
	}
}

func TestVerticalBarChartDoesNotRenderAboveLabelForNegativeSegments(t *testing.T) {
	t.Parallel()
	builder := myrtle.NewBuilder(defaulttheme.New())

	email := builder.
		AddVerticalBarChart(
			[]string{"Jan", "Feb"},
			[]myrtle.VerticalBarChartSeries{{Key: "churn", Label: "Churn", Values: []float64{-5, -4}}},
			myrtle.VerticalBarChartHeight(100),
			myrtle.VerticalBarChartAxisMin(-100),
			myrtle.VerticalBarChartAxisMax(0),
			myrtle.VerticalBarChartValueLabelsOption(myrtle.VerticalBarChartValueLabels{Show: true, MinSegmentHeight: 60}),
		).
		Build()

	html, err := email.HTML()
	if err != nil {
		t.Fatalf("html returned error: %v", err)
	}

	if strings.Contains(html, ">-5<") || strings.Contains(html, ">-4<") {
		t.Fatalf("expected thin negative value labels to stay hidden")
	}
}

func TestVerticalBarChartRendersAboveLabelForUpperThinNegativeSegmentWhenNoPositiveValues(t *testing.T) {
	t.Parallel()
	builder := myrtle.NewBuilder(defaulttheme.New())

	email := builder.
		AddVerticalBarChart(
			[]string{"Jan", "Feb"},
			[]myrtle.VerticalBarChartSeries{{Key: "churn", Label: "Churn", Values: []float64{-5, -80}}},
			myrtle.VerticalBarChartHeight(100),
			myrtle.VerticalBarChartAxisMin(-100),
			myrtle.VerticalBarChartAxisMax(100),
			myrtle.VerticalBarChartValueLabelsOption(myrtle.VerticalBarChartValueLabels{Show: true, MinSegmentHeight: 12}),
		).
		Build()

	html, err := email.HTML()
	if err != nil {
		t.Fatalf("html returned error: %v", err)
	}

	if !strings.Contains(html, ">-5<") {
		t.Fatalf("expected thin upper negative label to render above baseline when no positive values exist")
	}
}

func TestVerticalBarChartSupportsStructConfigOptions(t *testing.T) {
	t.Parallel()
	builder := myrtle.NewBuilder(defaulttheme.New())

	email := builder.
		AddVerticalBarChart(
			[]string{"Jan", "Feb"},
			[]myrtle.VerticalBarChartSeries{
				{Key: "new", Label: "New", Values: []float64{42, 36}},
				{Key: "churn", Label: "Churn", Values: []float64{-11, -8}},
			},
			myrtle.VerticalBarChartAxisConfig(myrtle.VerticalBarChartAxis{ShowYTicks: true, ShowBaseline: true, LabelFormat: myrtle.VerticalBarChartAxisLabelFormatNumber}),
			myrtle.VerticalBarChartLegendConfigOption(myrtle.VerticalBarChartLegendConfig{Placement: myrtle.VerticalBarChartLegendBottom}),
			myrtle.VerticalBarChartValueLabelsOption(myrtle.VerticalBarChartValueLabels{Show: true, MinSegmentHeight: 10}),
		).
		Build()

	html, err := email.HTML()
	if err != nil {
		t.Fatalf("html returned error: %v", err)
	}

	if !strings.Contains(html, ">42<") {
		t.Fatalf("expected value labels to render with struct config option")
	}
	if !strings.Contains(html, "New") || !strings.Contains(html, "Churn") {
		t.Fatalf("expected legend labels to render with struct config option")
	}
}

func TestVerticalBarChartValueFormatterCanRenderEuro(t *testing.T) {
	t.Parallel()
	builder := myrtle.NewBuilder(defaulttheme.New())

	email := builder.
		AddVerticalBarChart(
			[]string{"Jan", "Feb"},
			[]myrtle.VerticalBarChartSeries{{Key: "new", Label: "New", Values: []float64{42, 36}}},
			myrtle.VerticalBarChartAxisConfig(myrtle.VerticalBarChartAxis{ShowYTicks: true, LabelFormat: myrtle.VerticalBarChartAxisLabelFormatNumber}),
			myrtle.VerticalBarChartValueFormatterOption(myrtle.VerticalBarChartValueFormatter{Prefix: "€"}),
			myrtle.VerticalBarChartValueLabelsOption(myrtle.VerticalBarChartValueLabels{Show: true, MinSegmentHeight: 10}),
		).
		Build()

	html, err := email.HTML()
	if err != nil {
		t.Fatalf("html returned error: %v", err)
	}

	if !strings.Contains(html, "€42") {
		t.Fatalf("expected euro formatter to render value labels and/or axis labels with euro sign")
	}
}

func TestVerticalBarChartValueFormatterCanRenderUSDAndGBP(t *testing.T) {
	t.Parallel()
	builder := myrtle.NewBuilder(defaulttheme.New())

	emailUSD := builder.
		AddVerticalBarChart(
			[]string{"Jan", "Feb"},
			[]myrtle.VerticalBarChartSeries{{Key: "new", Label: "New", Values: []float64{42, 36}}},
			myrtle.VerticalBarChartAxisConfig(myrtle.VerticalBarChartAxis{ShowYTicks: true, LabelFormat: myrtle.VerticalBarChartAxisLabelFormatNumber}),
			myrtle.VerticalBarChartValueFormatterOption(myrtle.VerticalBarChartValueFormatter{Prefix: "$"}),
			myrtle.VerticalBarChartValueLabelsOption(myrtle.VerticalBarChartValueLabels{Show: true, MinSegmentHeight: 10}),
		).
		Build()

	htmlUSD, err := emailUSD.HTML()
	if err != nil {
		t.Fatalf("html returned error: %v", err)
	}
	if !strings.Contains(htmlUSD, "$42") {
		t.Fatalf("expected usd formatter to render dollar values")
	}

	emailGBP := builder.
		AddVerticalBarChart(
			[]string{"Jan", "Feb"},
			[]myrtle.VerticalBarChartSeries{{Key: "new", Label: "New", Values: []float64{42, 36}}},
			myrtle.VerticalBarChartAxisConfig(myrtle.VerticalBarChartAxis{ShowYTicks: true, LabelFormat: myrtle.VerticalBarChartAxisLabelFormatNumber}),
			myrtle.VerticalBarChartValueFormatterOption(myrtle.VerticalBarChartValueFormatter{Prefix: "£"}),
			myrtle.VerticalBarChartValueLabelsOption(myrtle.VerticalBarChartValueLabels{Show: true, MinSegmentHeight: 10}),
		).
		Build()

	htmlGBP, err := emailGBP.HTML()
	if err != nil {
		t.Fatalf("html returned error: %v", err)
	}
	if !strings.Contains(htmlGBP, "£42") {
		t.Fatalf("expected gbp formatter to render pound values")
	}
}

func TestVerticalBarChartValueFormatterCanRenderCompact(t *testing.T) {
	t.Parallel()
	builder := myrtle.NewBuilder(defaulttheme.New())

	emailK := builder.
		AddVerticalBarChart(
			[]string{"Jan", "Feb"},
			[]myrtle.VerticalBarChartSeries{{Key: "new", Label: "New", Values: []float64{1200, 1500}}},
			myrtle.VerticalBarChartAxisConfig(myrtle.VerticalBarChartAxis{ShowYTicks: true, LabelFormat: myrtle.VerticalBarChartAxisLabelFormatNumber}),
			myrtle.VerticalBarChartValueFormatterOption(myrtle.VerticalBarChartValueFormatter{MagnitudeSuffix: myrtle.VerticalBarChartMagnitudeSuffixShort}),
			myrtle.VerticalBarChartValueLabelsOption(myrtle.VerticalBarChartValueLabels{Show: true, MinSegmentHeight: 10}),
		).
		Build()

	htmlK, err := emailK.HTML()
	if err != nil {
		t.Fatalf("html returned error: %v", err)
	}
	if !strings.Contains(htmlK, "1.2K") {
		t.Fatalf("expected compact formatter to render K values")
	}

	emailM := builder.
		AddVerticalBarChart(
			[]string{"Jan", "Feb"},
			[]myrtle.VerticalBarChartSeries{{Key: "new", Label: "New", Values: []float64{1200000, 1500000}}},
			myrtle.VerticalBarChartAxisConfig(myrtle.VerticalBarChartAxis{ShowYTicks: true, LabelFormat: myrtle.VerticalBarChartAxisLabelFormatNumber}),
			myrtle.VerticalBarChartValueFormatterOption(myrtle.VerticalBarChartValueFormatter{MagnitudeSuffix: myrtle.VerticalBarChartMagnitudeSuffixShort}),
			myrtle.VerticalBarChartValueLabelsOption(myrtle.VerticalBarChartValueLabels{Show: true, MinSegmentHeight: 10}),
		).
		Build()

	htmlM, err := emailM.HTML()
	if err != nil {
		t.Fatalf("html returned error: %v", err)
	}
	if !strings.Contains(htmlM, "1.5M") {
		t.Fatalf("expected compact formatter to render M values")
	}
}

func TestVerticalBarChartValueFormatterCanRenderNegativeParentheses(t *testing.T) {
	t.Parallel()
	builder := myrtle.NewBuilder(defaulttheme.New())

	email := builder.
		AddVerticalBarChart(
			[]string{"Jan", "Feb"},
			[]myrtle.VerticalBarChartSeries{{Key: "churn", Label: "Churn", Values: []float64{-40, -60}}},
			myrtle.VerticalBarChartHeight(120),
			myrtle.VerticalBarChartAxisMin(-100),
			myrtle.VerticalBarChartAxisMax(0),
			myrtle.VerticalBarChartValueFormatterOption(myrtle.VerticalBarChartValueFormatter{NegativeFormat: myrtle.VerticalBarChartNegativeFormatParentheses}),
			myrtle.VerticalBarChartValueLabelsOption(myrtle.VerticalBarChartValueLabels{Show: true, MinSegmentHeight: 10}),
		).
		Build()

	html, err := email.HTML()
	if err != nil {
		t.Fatalf("html returned error: %v", err)
	}

	if !strings.Contains(html, ">(40)<") {
		t.Fatalf("expected negative labels to use parentheses format")
	}
	if strings.Contains(html, ">-40<") {
		t.Fatalf("expected negative labels to avoid minus format when parentheses are configured")
	}
	if !strings.Contains(html, "title=\"Churn: (40)\"") {
		t.Fatalf("expected negative title values to use parentheses format")
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
			myrtle.TilesAlign(myrtle.TileAlignmentStart),
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
			myrtle.TilesAlign(myrtle.TileAlignmentEnd),
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

func TestCalloutSolidSuccessUsesStrongBorderColor(t *testing.T) {
	t.Parallel()
	builder := myrtle.NewBuilder(defaulttheme.New())

	email := builder.
		AddCallout(myrtle.ToneSuccess, "Success", "Everything is healthy", myrtle.CalloutStyle(myrtle.CalloutVariantSolid)).
		Build()

	html, err := email.HTML()
	if err != nil {
		t.Fatalf("html returned error: %v", err)
	}

	if !strings.Contains(html, "background:#16a34a;") {
		t.Fatalf("expected solid success callout background to use strong success color")
	}
	if !strings.Contains(html, "border:1px solid #16a34a;") {
		t.Fatalf("expected solid success callout border to use strong success color")
	}
	if strings.Contains(html, "border:1px solid #86efac;") {
		t.Fatalf("expected solid success callout to avoid light success border color")
	}
}

func TestNewBuilderWithHeaderOverrideLater(t *testing.T) {
	t.Parallel()
	builder := myrtle.NewBuilder(
		defaulttheme.New(),
		myrtle.WithHeader(myrtle.HeadingBlock{Text: "Initial", Level: 1}),
	)

	email := builder.WithPreheader("Initial preheader").WithHeader(myrtle.HeadingBlock{Text: "Overridden", Level: 1}).AddText("Body").Build()

	html, err := email.HTML()
	if err != nil {
		t.Fatalf("html returned error: %v", err)
	}
	if !strings.Contains(html, "Overridden") || strings.Contains(html, "Initial</h1>") {
		t.Fatalf("expected chained WithHeader to override constructor header block")
	}
}

func TestBuilderCloneConcurrentUsage(t *testing.T) {
	t.Parallel()
	baseBuilder := myrtle.NewBuilder(defaulttheme.New()).
		WithHeader(myrtle.HeadingBlock{Text: "Template", Level: 1}).
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
				WithPreheader(fmt.Sprintf("preheader-%d", index)).
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

func TestMessagePreviewMarkdownUsedForHTML(t *testing.T) {
	t.Parallel()

	email := myrtle.NewBuilder(defaulttheme.New()).
		AddMessage(myrtle.MessageBlock{
			Subject:         "Subject",
			Preview:         "Plain preview text",
			PreviewMarkdown: "Markdown preview with [docs](https://example.com/docs)",
			URL:             "https://example.com/message",
		}).
		Build()

	html, err := email.HTML()
	if err != nil {
		t.Fatalf("html returned error: %v", err)
	}

	if !strings.Contains(html, "Markdown preview with") || !strings.Contains(html, "https://example.com/docs") {
		t.Fatalf("expected html to render PreviewMarkdown when provided")
	}
	if strings.Contains(html, "Plain preview text") {
		t.Fatalf("expected html preview to prioritize PreviewMarkdown over Preview")
	}
}

func TestMessagePreviewMarkdownTextFallback(t *testing.T) {
	t.Parallel()

	email := myrtle.NewBuilder(defaulttheme.New()).
		AddMessageDigest([]myrtle.MessageBlock{
			{
				Subject:         "Only markdown preview",
				PreviewMarkdown: "Can you check [the draft](https://example.com/draft)?",
				URL:             "https://example.com/message/1",
			},
			{
				Subject:         "Both preview forms",
				Preview:         "Use plain preview here",
				PreviewMarkdown: "Use [markdown](https://example.com/markdown) instead",
				URL:             "https://example.com/message/2",
			},
		}).
		Build()

	text, err := email.Text()
	if err != nil {
		t.Fatalf("text returned error: %v", err)
	}

	if !strings.Contains(text, "Can you check the draft (https://example.com/draft)?") {
		t.Fatalf("expected text fallback to normalize markdown links from PreviewMarkdown")
	}
	if !strings.Contains(text, "Use plain preview here") {
		t.Fatalf("expected text fallback to prefer Preview over PreviewMarkdown when both are set")
	}
	if strings.Contains(text, "[the draft](https://example.com/draft)") {
		t.Fatalf("expected text fallback not to include raw markdown link syntax")
	}
}

func TestMessageDigestInsetNoneRemovesSideBordersAndCornersInTerminalTheme(t *testing.T) {
	t.Parallel()

	email := myrtle.NewBuilder(terminal.New()).
		AddMessageDigest(
			[]myrtle.MessageBlock{{Subject: "Digest item", Preview: "Preview", URL: "https://example.com/message"}},
			myrtle.MessageDigestInsetMode(myrtle.InsetModeNone),
		).
		Build()

	html, err := email.HTML()
	if err != nil {
		t.Fatalf("html returned error: %v", err)
	}

	if !strings.Contains(html, "border-radius:0;") {
		t.Fatalf("expected message digest to remove rounded corners in inset none mode")
	}
	if !strings.Contains(html, "border-left:0;border-right:0;") {
		t.Fatalf("expected message digest to remove side borders in inset none mode")
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
