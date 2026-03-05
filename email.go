package myrtle

import (
	"errors"
	"strings"

	"github.com/gzuidhof/myrtle/theme"
)

var ErrThemeCannotRenderBlock = errors.New("myrtle: theme cannot render block")

type htmlBlock interface {
	RenderHTML(values theme.Values) (string, error)
}

// Email is an immutable rendered message composed from header metadata, values, and blocks.
type Email struct {
	header                 *HeaderSection
	footer                 *FooterSection
	preheader              string
	preheaderPaddingRepeat int
	values                 theme.Values
	blocks                 []Block
	theme                  theme.Theme
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
	header, err := email.headerHTMLView()
	if err != nil {
		return "", err
	}
	footer, err := email.footerHTMLView()
	if err != nil {
		return "", err
	}

	fragments := make([]string, 0, len(email.blocks))
	blockViews := make([]theme.EmailBlockView, 0, len(email.blocks))

	for _, block := range email.blocks {
		if block == nil {
			continue
		}

		customHTML, ok, err := email.renderBlockHTML(block)
		if err != nil {
			return "", err
		}

		if ok {
			spec := normalizedLayoutSpec(block.LayoutSpec())
			fragments = append(fragments, customHTML)
			blockViews = append(blockViews, theme.EmailBlockView{
				Kind:        block.Kind(),
				HTML:        customHTML,
				InsetMode:   string(spec.InsetMode),
				CustomInset: spec.CustomInset,
			})
			continue
		}

		return "", errors.Join(ErrThemeCannotRenderBlock, errors.New(string(block.Kind())))
	}

	return email.theme.RenderHTML(theme.EmailView{
		Header:                 header,
		Footer:                 footer,
		Preheader:              email.preheader,
		PreheaderPaddingRepeat: email.preheaderPaddingRepeat,
		Values:                 email.values,
		Blocks:                 fragments,
		BlockViews:             blockViews,
	})
}

func (email *Email) renderBlockHTML(block Block) (string, bool, error) {
	data := block.TemplateData()

	if htmlRenderer, ok := block.(htmlBlock); ok {
		customHTML, err := htmlRenderer.RenderHTML(email.values)
		if err != nil {
			return "", false, err
		}

		return customHTML, true, nil
	}

	return email.theme.RenderBlockHTML(theme.BlockView{
		Kind:   block.Kind(),
		Data:   data,
		Values: email.values,
	})
}

// Text renders the email into a plain-text fallback representation.
func (email *Email) Text() (string, error) {
	textParts := make([]string, 0, len(email.blocks)+2)

	for _, block := range email.blocks {
		if block == nil {
			continue
		}

		text, err := block.RenderText(RenderContext{
			Preheader: email.Preheader(),
			Values:    email.values,
		})
		if err != nil {
			return "", err
		}

		if strings.TrimSpace(text) == "" {
			continue
		}

		textParts = append(textParts, text)
	}

	body := strings.Join(textParts, "\n\n")
	textHeader, err := email.textHeaderView()
	if err != nil {
		return "", err
	}
	textFooter, err := email.textFooterView()
	if err != nil {
		return "", err
	}

	if wrapper, ok := email.theme.(theme.TextWrapper); ok {
		return wrapper.WrapText(theme.TextView{
			Header:    textHeader,
			Footer:    textFooter,
			Preheader: email.preheader,
			Values:    email.values,
			Body:      body,
		})
	}

	var outputParts []string
	if strings.TrimSpace(email.Preheader()) != "" {
		outputParts = append(outputParts, email.Preheader())
	}
	if strings.TrimSpace(body) != "" {
		outputParts = append(outputParts, body)
	}
	if textFooter != nil && strings.TrimSpace(textFooter.Text) != "" {
		outputParts = append(outputParts, textFooter.Text)
	}

	return strings.Join(outputParts, "\n\n"), nil
}

func (email *Email) textHeaderView() (*theme.HeaderView, error) {
	if email.header == nil || !email.header.RenderInText {
		return nil, nil
	}

	if email.header.Block == nil {
		return nil, nil
	}

	text, err := email.header.Block.RenderText(RenderContext{
		Preheader: email.Preheader(),
		Values:    email.values,
	})
	if err != nil {
		return nil, err
	}

	text = strings.TrimSpace(text)
	if text == "" {
		return nil, nil
	}

	return &theme.HeaderView{Text: text}, nil
}

func (email *Email) headerHTMLView() (*theme.HeaderView, error) {
	if email.header == nil {
		return nil, nil
	}

	if email.header.Block == nil {
		return nil, nil
	}

	headerHTML, ok, err := email.renderBlockHTML(email.header.Block)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, errors.Join(ErrThemeCannotRenderBlock, errors.New(string(email.header.Block.Kind())))
	}
	spec := normalizedLayoutSpec(email.header.Block.LayoutSpec())

	return &theme.HeaderView{
		HTML:        headerHTML,
		Placement:   string(normalizedHeaderPlacement(email.header.Placement)),
		InsetMode:   string(spec.InsetMode),
		CustomInset: spec.CustomInset,
	}, nil
}

func (email *Email) textFooterView() (*theme.FooterView, error) {
	if email.footer == nil || !email.footer.RenderInText {
		return nil, nil
	}

	if email.footer.Block == nil {
		return nil, nil
	}

	text, err := email.footer.Block.RenderText(RenderContext{
		Preheader: email.Preheader(),
		Values:    email.values,
	})
	if err != nil {
		return nil, err
	}

	text = strings.TrimSpace(text)
	if text == "" {
		return nil, nil
	}

	return &theme.FooterView{Text: text}, nil
}

func (email *Email) footerHTMLView() (*theme.FooterView, error) {
	if email.footer == nil {
		return nil, nil
	}

	if email.footer.Block == nil {
		return nil, nil
	}

	footerHTML, ok, err := email.renderBlockHTML(email.footer.Block)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, errors.Join(ErrThemeCannotRenderBlock, errors.New(string(email.footer.Block.Kind())))
	}
	spec := normalizedLayoutSpec(email.footer.Block.LayoutSpec())

	return &theme.FooterView{
		HTML:        footerHTML,
		Placement:   string(normalizedFooterPlacement(email.footer.Placement)),
		InsetMode:   string(spec.InsetMode),
		CustomInset: spec.CustomInset,
	}, nil
}
