<picture>
  <source media="(prefers-color-scheme: dark)" srcset="logo-light.png">
  <img alt="Myrtle logo" src="logo.png">
</picture>

# 📩 myrtle
Composable, strongly typed email content builder for Go.

Myrtle focuses on building email content only: subject, preheader, and composable blocks.
It renders both HTML and Markdown text output, and keeps delivery concerns (from/to/headers)
outside the library.

## Features

- Fluent builder pattern for email content.
- Strongly typed built-in blocks.
- Optional generic registry for typed custom blocks.
- Theme packages under `theme/<themename>`.
- Embedded templates (`embed`) stored as separate files next to theme code.
- High-level values available to wrappers and blocks (product name/link, logo URL/alt).
- Shared style values available to all blocks (`ColorPrimary`, `ColorSecondary`, text/border/code colors).
- Built-in advanced blocks: table, action, code, free markdown, bar chart.
- High-impact blocks: timeline, stats row, badge, summary card, attachment.
- Dual rendering APIs:
  - `HTML()` for final HTML output.
  - `Text()` for Markdown fallback output.

## Installation

```bash
go get github.com/gzuidhof/myrtle
```

## Usage

```go
package main

import (
  "fmt"

  "github.com/gzuidhof/myrtle"
  "github.com/gzuidhof/myrtle/theme"
  defaulttheme "github.com/gzuidhof/myrtle/theme/default"
  "github.com/gzuidhof/myrtle/theme/flat"
)

func main() {
  b := myrtle.NewBuilder(defaulttheme.New(), myrtle.WithStyles(theme.Styles{ColorPrimary: "#0ea5e9"}))
  _ = flat.New(flat.WithFallback(defaulttheme.New()))

  email := b.
    Subject("Welcome to Myrtle").
    Preheader("Composable email building blocks in Go").
    Product("Myrtle", "https://github.com/gzuidhof/myrtle").
    Logo("https://example.com/logo.png", "").
    AddText("Hi there,").
    AddText("Thanks for trying Myrtle.").
    AddText("Start with the quick-start docs:").
    AddButton("Open docs", "https://github.com/gzuidhof/myrtle").
    AddVerificationCode("493817").
    Build()

  html, err := email.HTML()
  if err != nil {
    panic(err)
  }

  md, err := email.Text()
  if err != nil {
    panic(err)
  }

  fmt.Println(html)
  fmt.Println(md)
}
```

## Custom blocks with strong typing

```go
type Promo struct {
  Title string
  Body  string
}

registry := myrtle.NewRegistry()

_ = myrtle.Register(registry, "promo",
  func(p Promo, context myrtle.RenderContext) (string, error) {
    return "## " + p.Title + "\n\n" + p.Body + " at " + context.Values.ProductName, nil
  },
)

b := myrtle.NewBuilder(
  defaulttheme.New(),
  myrtle.WithRegistry(registry),
)
  b.ProductName("Myrtle")

promoBlock, _ := myrtle.Create(registry, "promo", Promo{Title: "Launch", Body: "It works."})
b.Add(promoBlock)
```

## Notes

- Markdown rendering is block-owned by default.
- Themes fully own HTML rendering (base package does not generate HTML).
- A theme can delegate missing block renderers to another theme (for example `flat.WithFallback(defaulttheme.New())`).
- Themes can optionally wrap Markdown output globally.
- If logo alt is empty, Myrtle defaults it to product name (or `Logo` when product name is empty).

## Examples

- [example/welcome.go](example/welcome.go)
- [example/security.go](example/security.go)
- [example/password_reset.go](example/password_reset.go)
- [example/account_deletion_confirmation.go](example/account_deletion_confirmation.go)
- [example/report.go](example/report.go)
- [example/common_blocks.go](example/common_blocks.go)
- [example/onboarding.go](example/onboarding.go)
- [example/billing_receipt.go](example/billing_receipt.go)
- [example/incident_notice.go](example/incident_notice.go)
- [example/feature_digest.go](example/feature_digest.go)
- [example/high_impact.go](example/high_impact.go)
- [example/columns_complex.go](example/columns_complex.go)
- [example/bar_chart.go](example/bar_chart.go)
- [example/monster.go](example/monster.go)
- Rendering/sending test harness: [example/examples_test.go](example/examples_test.go)

### Example server

The [example/server](example/server) package serves a directory of all example emails and block previews.
Run it via the cmd entrypoint:

```bash
go run ./example/server/cmd
```

```go
package main

import (
  "log"

  "github.com/gzuidhof/myrtle/example/server"
)

func main() {
  srv, err := server.New()
  if err != nil {
    log.Fatal(err)
  }

  log.Fatal(srv.ListenAndServe(":8380"))
}
```

