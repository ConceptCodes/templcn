import { defineConfig } from "@playwright/test"

export default defineConfig({
  testDir: "./tests",
  reporter: "list",
  workers: process.env.CI ? 2 : 1,
  webServer: {
    command: "node ./tests/static-server.mjs",
    url: "http://127.0.0.1:4173/tests/fixtures/dialog-runtime.html",
    reuseExistingServer: !process.env.CI,
  },
  use: {
    baseURL: "http://127.0.0.1:4173",
    // Use the managed Playwright browser in CI; prefer installed Chrome locally.
    channel: process.env.CI ? undefined : "chrome",
    trace: "on-first-retry",
  },
  projects: [
    {
      name: "chromium",
      use: { viewport: { width: 1280, height: 720 } },
    },
  ],
})
