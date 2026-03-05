<picture>
  <source media="(prefers-color-scheme: dark)" srcset="logo-light.png">
  <img alt="Myrtle logo" src="logo.png">
</picture>

# 🌸 myrtle
Myrtle is a composable, strongly typed email content builder for Go.

## Quick preview

Security example email side-by-side in two themes:

| Default | Terminal |
| --- | --- |
| ![Security example (default)](screenshots/default--security.png) | ![Security example (terminal)](screenshots/terminal--security.png) |

## Features

- Fluent builder pattern for email content.
- Strongly typed library of blocks.
- Modern built-in themes: `default`, `flat`, `terminal`, `editorial`.
- Shared customizable styles.
- Built-in advanced blocks such as tables, charts, grids.
- High-impact blocks: timelines, standout stats rows, badges, attachments.
- Dual rendering APIs:
  - `HTML()` for final HTML output.
  - `Text()` for plain-text fallback output.
- Customizable: bring your own theme, styles or custom blocks.
- Left-to-right and right-to-left direction support (e.g. for Arabic/Hebrew).
- Renders OK in Outlook Classic and other notoriously difficult email clients.
- Dependency-free aside from [`goldmark`](https://github.com/yuin/goldmark) for Markdown rendering.

## Installation

```bash
go get github.com/gzuidhof/myrtle
```

## Quick start (security email)

```go
package main

import (
  "github.com/gzuidhof/myrtle"
  defaulttheme "github.com/gzuidhof/myrtle/theme/default"
)

func main() {
  email := myrtle.NewBuilder(defaulttheme.New()).
    WithPreheader("Use this one-time code to sign in").
    AddHeading("Your verification code").
    AddText("Use the code below to complete your sign-in. This code expires in 10 minutes.").
    Add(myrtle.VerificationCodeBlock{Label: "Verification code", Value: "493817"}).
    AddKeyValue("Request details", []myrtle.KeyValuePair{
      {Key: "IP", Value: "203.0.113.5"},
      {Key: "Location", Value: "Amsterdam, NL"},
    }).
    AddText("If you did not request this code, secure your account immediately.").
    AddButton("Review account", "https://example.com/account/security").
    Build()

  html, err := email.HTML()
  if err != nil {
    panic(err)
  }

  md, err := email.Text()
  if err != nil {
    panic(err)
  }

  // Use your favorite e-mail sending library to send the email with the generated HTML and text content.
  // ...

  _ = html
  _ = md
}
```

Make use of auto-complete/Intellisense in your IDE to explore the rich library of blocks and customization options.

## Examples

- [example/weekly_operations_brief.go](example/high_impact.go)
- [example/account_deletion_confirmation.go](example/account_deletion_confirmation.go)
- [example/security.go](example/security.go)
- [example/monster.go](example/monster.go)

### Rendered examples

#### Weekly operations brief

| Theme | Preview |
| --- | --- |
| Default | ![Weekly operations brief (default)](screenshots/default--weekly-operations-brief.png) |
| Flat | ![Weekly operations brief (flat)](screenshots/flat--weekly-operations-brief.png) |
| Terminal | ![Weekly operations brief (terminal)](screenshots/terminal--weekly-operations-brief.png) |
| Editorial | ![Weekly operations brief (editorial)](screenshots/editorial--weekly-operations-brief.png) |

#### Account deletion confirmation

| Theme | Preview |
| --- | --- |
| Default | ![Account deletion confirmation (default)](screenshots/default--account-deletion-confirmation.png) |
| Flat | ![Account deletion confirmation (flat)](screenshots/flat--account-deletion-confirmation.png) |
| Terminal | ![Account deletion confirmation (terminal)](screenshots/terminal--account-deletion-confirmation.png) |
| Editorial | ![Account deletion confirmation (editorial)](screenshots/editorial--account-deletion-confirmation.png) |

#### Security confirmation

| Theme | Preview |
| --- | --- |
| Default | ![Security confirmation (default)](screenshots/default--security.png) |
| Flat | ![Security confirmation (flat)](screenshots/flat--security.png) |
| Terminal | ![Security confirmation (terminal)](screenshots/terminal--security.png) |
| Editorial | ![Security confirmation (editorial)](screenshots/editorial--security.png) |

#### Monster

The monster example is a fun showcase of many blocks and styles together. It intentionally has a lot of content to demonstrate how the builder and themes handle it.

- [screenshots/default--monster.png](screenshots/default--monster.png)

## Example server

The [example/server](example/server) package serves a directory of all example emails and block previews.

Clone this repository and run the server to preview example emails in the browser at `http://localhost:8380/`.

```bash
go run ./example/server/cmd
```

![Example server preview](screenshots/example-server.png)

## Development
The code for this repository is repetitive and verbose, I recommend you use AI-assisted code generation to speed up development. Writing inlined CSS manually is particularly painful.

> Myrtle she wrote.
