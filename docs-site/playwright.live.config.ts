import { defineConfig } from "@playwright/test"

export default defineConfig({
  testDir: "./tests",
  reporter: "list",
  workers: process.env.CI ? 2 : 1,
  use: {
    viewport: { width: 1280, height: 720 },
    baseURL: "http://127.0.0.1:4173",
    channel: process.env.CI ? undefined : "chrome",
  },
})
