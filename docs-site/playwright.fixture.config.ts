import { defineConfig } from "@playwright/test"

export default defineConfig({
  testDir: "./tests",
  testIgnore: "**/component-pages.spec.ts",
  reporter: "list",
  workers: process.env.CI ? 2 : 1,
  webServer: {
    command: "PORT=4174 node ./tests/static-server.mjs",
    url: "http://127.0.0.1:4174/tests/fixtures/dialog-runtime.html",
    reuseExistingServer: false,
  },
  use: {
    viewport: { width: 1280, height: 720 },
    baseURL: "http://127.0.0.1:4174",
    channel: process.env.CI ? undefined : "chrome",
    trace: "on-first-retry",
  },
})
