import { defineConfig, devices } from "@playwright/test"

export default defineConfig({
  testDir: "./tests",
  testIgnore: "**/component-pages.spec.ts",
  reporter: "list",
  webServer: {
    command: "PORT=4174 node ./tests/static-server.mjs",
    url: "http://127.0.0.1:4174/tests/fixtures/dialog-runtime.html",
    reuseExistingServer: false,
  },
  use: {
    ...devices["Desktop Chrome"],
    baseURL: "http://127.0.0.1:4174",
    browserName: "chromium",
    trace: "on-first-retry",
  },
})
