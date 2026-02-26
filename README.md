<picture>
  <source media="(prefers-color-scheme: dark)" srcset="logo-light.png">
  <img alt="Myrtle logo" src="logo.png">
</picture>

# 📩 myrtle
Composable, strongly typed email content builder for Go.

Myrtle focuses on building email content only: subject, preheader, and composable blocks.
It renders both HTML and plain text output, and keeps delivery concerns (from/to/headers)
outside the library.

## Features

- Fluent builder pattern for email content.
- Strongly typed built-in blocks.
- Optional generic registry for typed custom blocks.
- Theme packages under `theme/<themename>`.
- Embedded templates (`embed`) stored as separate files next to theme code.
- Shared style values available to all blocks (`ColorPrimary`, `ColorSecondary`, text/border/code colors).
- Built-in advanced blocks: table, action, code, free markdown, bar chart.
- High-impact blocks: timeline, stats row, badge, summary card, attachment.
- Dual rendering APIs:
  - `HTML()` for final HTML output.
  - `Text()` for plain-text fallback output.

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
    WithPreheader("Composable email building blocks in Go").
    WithHeader(myrtle.HeadingBlock{Text: "Myrtle", Level: 1}).
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
  func(p Promo, values theme.Values) (string, error) {
    return "<section><h2 style=\"color:" + values.Styles.ColorPrimary + "\">" + p.Title + "</h2><p>" + p.Body + "</p></section>", nil
  },
  func(p Promo, context myrtle.RenderContext) (string, error) {
    return "## " + p.Title + "\n\n" + p.Body, nil
  },
)

b := myrtle.NewBuilder(
  defaulttheme.New(),
)

b.Add(myrtle.NewCustomBlock("promo", Promo{Title: "Launch", Body: "It works."},
  func(p Promo, values theme.Values) (string, error) {
    return "<section><h2 style=\"color:" + values.Styles.ColorPrimary + "\">" + p.Title + "</h2><p>" + p.Body + "</p></section>", nil
  },
  func(p Promo, context myrtle.RenderContext) (string, error) {
    return "## " + p.Title + "\n\n" + p.Body, nil
  },
))
```

## Notes

- Custom blocks can be added directly with `myrtle.NewCustomBlock(...)` and `Add(...)`.
- `Registry` remains available for registration/lookup workflows via `myrtle.Register` + `myrtle.CreateBlock`.
- Direction can be set with `myrtle.WithDirection(theme.DirectionRTL)` (or `builder.WithDirection(...)`) and themes render `dir="rtl"` when enabled.
- Alignment constants use logical values (`start`, `center`, `end`) and are mapped to physical left/right at render time for email-client compatibility.
- A theme can delegate missing block renderers to another theme (for example `flat.WithFallback(defaulttheme.New())`).
- Themes can optionally wrap text output globally.

### Direction example

```go
email := myrtle.NewBuilder(defaulttheme.New(), myrtle.WithDirection(theme.DirectionRTL)).
  WithHeader(myrtle.HeadingBlock{Text: "Aligned to logical start", Level: 1}).
  AddText("Aligned to logical start", myrtle.TextAlign(myrtle.TextAlignStart)).
  AddButton("Open", "https://example.com", myrtle.ButtonAlign(myrtle.ButtonAlignmentEnd)).
  Build()
```

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

#### Send test emails via SMTP (optional)

The example server can optionally show a **Send email** form above each example email preview.
This is enabled only when an SMTP config file is present.

1. Copy the example config:

```bash
cp example/server/smtp.config.example.json example/server/smtp.config.json
```

2. Fill in your SMTP settings and credentials in `example/server/smtp.config.json`.

JSON fields:

- `host`: SMTP host
- `port`: SMTP port (for example `587`)
- `username`: SMTP username (optional if your relay allows anonymous)
- `password`: SMTP password
- `from_name`: Display name for the sender
- `from_address`: Sender email address
- `default_to`: Default pre-filled recipient in the send form

3. Start the server:

```bash
go run ./example/server/cmd
```

When the config file exists and is valid, each example card on the index page shows a recipient field + send button.

Notes:

- The credentials file `example/server/smtp.config.json` is ignored by Git.
- You can override the config path with `MYRTLE_SMTP_CONFIG=/path/to/file.json`.

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

