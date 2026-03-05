# Preview screenshots (Playwright)

Capture screenshots for one or more example emails across multiple themes.

Assumes the example server is already running (default: `http://localhost:8380`).

## Setup

```bash
cd scripts/preview
npm install
npx playwright install chromium
```

## Usage

```bash
npm run capture -- \
  --email welcome \
  --themes default,flat,terminal \
  --base-url http://localhost:8380
```

Batch capture multiple emails:

```bash
npm run capture -- \
  --emails welcome,security,vertical-bar-chart \
  --themes default,flat,terminal
```

Output files are named like:

- `default--welcome.png`
- `flat--welcome.png`
- `terminal--welcome.png`

By default, screenshots are written to the top-level `screenshots/` folder.

## Options

- `--email <key>`: Example email key (for `/emails/<key>/html`)
- `--emails <a,b,c>`: Comma-separated list of email keys
- `--themes <a,b,c>`: Comma-separated theme list
- `--base-url <url>`: Example server base URL
- `--out-dir <path>`: Output directory for screenshots (default `../../screenshots`)
- `--width <px>`: Viewport width (default `720`)
- `--height <px>`: Viewport height (default `400`) (this is intentionally small to reduce empty space below the email content as we capture the full page anyhow).
