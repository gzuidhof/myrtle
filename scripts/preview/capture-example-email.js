#!/usr/bin/env node

import fs from "node:fs/promises";
import path from "node:path";
import { chromium } from "playwright";

function parseArgs(argv) {
  const args = {
    email: "welcome",
    emails: [],
    themes: ["default", "flat", "terminal", "editorial"],
    baseUrl: "http://localhost:8380",
    outDir: "../../example/screenshot",
    width: 720,
    height: 400,
  };

  for (let index = 2; index < argv.length; index++) {
    const value = argv[index];
    if (!value.startsWith("--")) {
      continue;
    }

    const key = value.slice(2);
    const next = argv[index + 1];
    const hasValue = typeof next === "string" && !next.startsWith("--");

    switch (key) {
      case "email":
        if (hasValue) {
          args.email = next;
          index++;
        }
        break;
      case "themes":
        if (hasValue) {
          args.themes = next
            .split(",")
            .map((themeName) => themeName.trim())
            .filter(Boolean);
          index++;
        }
        break;
      case "emails":
        if (hasValue) {
          args.emails = next
            .split(",")
            .map((emailName) => emailName.trim())
            .filter(Boolean);
          index++;
        }
        break;
      case "base-url":
        if (hasValue) {
          args.baseUrl = next;
          index++;
        }
        break;
      case "out-dir":
        if (hasValue) {
          args.outDir = next;
          index++;
        }
        break;
      case "width":
        if (hasValue) {
          args.width = Number.parseInt(next, 10);
          index++;
        }
        break;
      case "height":
        if (hasValue) {
          args.height = Number.parseInt(next, 10);
          index++;
        }
        break;
      default:
        break;
    }
  }

  return args;
}

function buildPreviewUrl(baseUrl, email, themeName) {
  const root = baseUrl.replace(/\/+$/, "");
  return `${root}/emails/${encodeURIComponent(email)}/html?theme=${encodeURIComponent(themeName)}`;
}

async function run() {
  const options = parseArgs(process.argv);
  const emailKeys = options.emails.length > 0 ? options.emails : [options.email];
  if (!Array.isArray(emailKeys) || emailKeys.length === 0 || !emailKeys[0]) {
    throw new Error("Missing email value. Use --email <key> or --emails a,b,c.");
  }
  if (!Array.isArray(options.themes) || options.themes.length === 0) {
    throw new Error("No themes selected. Use --themes default,flat,terminal,editorial");
  }

  const outputDirectory = path.resolve(process.cwd(), options.outDir);
  await fs.mkdir(outputDirectory, { recursive: true });

  const browser = await chromium.launch();
  const context = await browser.newContext({
    viewport: {
      width: options.width,
      height: options.height,
    },
    deviceScaleFactor: 2,
  });

  const page = await context.newPage();
  const screenshotPaths = [];

  try {
    for (const emailKey of emailKeys) {
      for (const themeName of options.themes) {
        const url = buildPreviewUrl(options.baseUrl, emailKey, themeName);
        const fileName = `${themeName}--${emailKey}.png`;
        const outputPath = path.join(outputDirectory, fileName);

        await page.goto(url, { waitUntil: "networkidle" });
        await page.screenshot({ path: outputPath, fullPage: true });

        screenshotPaths.push(outputPath);
        console.log(`captured ${emailKey} (${themeName}): ${outputPath}`);
      }
    }
  } finally {
    await context.close();
    await browser.close();
  }

  console.log(`done: ${screenshotPaths.length} screenshot(s)`);
}

run().catch((error) => {
  console.error(error instanceof Error ? error.message : String(error));
  process.exit(1);
});
