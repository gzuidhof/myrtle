package myrtle

import (
	"errors"
	"strings"

	"github.com/gzuidhof/myrtle/theme"
)

var ErrThemeCannotRenderBlock = errors.New("myrtle: theme cannot render block")

// Email is an immutable rendered message composed from header metadata, values, and blocks.
type Email struct {
	header    *HeaderSection
	preheader string
	values    theme.Values
	blocks    []Block
	theme     theme.Theme
}

// Preheader returns the preheader text.
func (email *Email) Preheader() string {
	return email.preheader
}

// Values returns the resolved theme values used while rendering this email.
func (email *Email) Values() theme.Values {
	return email.values
}

// HTML renders the email into themed HTML by rendering each block then composing the full layout.
func (email *Email) HTML() (string, error) {
	fragments := make([]string, 0, len(email.blocks))

	for _, block := range email.blocks {
		if block == nil {
			continue
		}

		if customHTML, ok, err := email.theme.RenderBlockHTML(theme.BlockView{
			Kind:   block.Kind(),
			Data:   block.TemplateData(),
			Values: email.values,
		}); err != nil {
			return "", err
		} else if ok {
			fragments = append(fragments, customHTML)
			continue
		}

		return "", errors.Join(ErrThemeCannotRenderBlock, errors.New(string(block.Kind())))
	}

	return email.theme.RenderHTML(theme.EmailView{
		Header:    email.headerView(),
		Preheader: email.preheader,
		Values:    email.values,
		Blocks:    fragments,
	})
}

// Text renders the email into a markdown/text fallback representation.
func (email *Email) Text() (string, error) {
	markdownParts := make([]string, 0, len(email.blocks)+2)

	for _, block := range email.blocks {
		if block == nil {
			continue
		}

		markdown, err := block.RenderMarkdown(RenderContext{
			Preheader: email.Preheader(),
			Values:    email.values,
		})
		if err != nil {
			return "", err
		}

		if strings.TrimSpace(markdown) == "" {
			continue
		}

		markdownParts = append(markdownParts, markdown)
	}

	body := strings.Join(markdownParts, "\n\n")
	markdownHeader := email.markdownHeaderView()

	if wrapper, ok := email.theme.(theme.MarkdownWrapper); ok {
		return wrapper.WrapMarkdown(theme.TextView{
			Header:    markdownHeader,
			Preheader: email.preheader,
			Values:    email.values,
			Body:      body,
		})
	}

	var outputParts []string
	if strings.TrimSpace(email.Preheader()) != "" {
		outputParts = append(outputParts, "_"+email.Preheader()+"_")
	}
	if strings.TrimSpace(body) != "" {
		outputParts = append(outputParts, body)
	}

	return strings.Join(outputParts, "\n\n"), nil
}

func (email *Email) markdownHeaderView() *theme.HeaderView {
	if email.header == nil || !email.header.RenderInMarkdown {
		return nil
	}

	return email.headerView()
}

func (email *Email) headerView() *theme.HeaderView {
	if email.header == nil {
		return nil
	}

	alignment := normalizedHeaderAlignment(email.header.Alignment)
	hasLogo := strings.TrimSpace(email.header.LogoURL) != ""
	showTextWithLogo := email.header.ShowTextWithLogo
	showTitle := strings.TrimSpace(email.header.Title) != "" && (!hasLogo || showTextWithLogo)
	showProductName := strings.TrimSpace(email.header.ProductName) != "" && (!hasLogo || showTextWithLogo)

	return &theme.HeaderView{
		Title:           email.header.Title,
		ShowTitle:       showTitle,
		ProductName:     email.header.ProductName,
		ShowProductName: showProductName,
		ProductLink:     email.header.ProductLink,
		LogoURL:         email.header.LogoURL,
		LogoAlt:         email.header.LogoAlt,
		LogoCentered:    alignment == HeaderAlignmentCenter,
		Alignment:       string(alignment),
	}
}
