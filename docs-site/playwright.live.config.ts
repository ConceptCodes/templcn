import { defineConfig, devices } from "@playwright/test"

export default defineConfig({
  testDir: "./tests",
  reporter: "list",
  use: {
    ...devices["Desktop Chrome"],
    baseURL: "http://127.0.0.1:4173",
    browserName: "chromium",
  },
})
