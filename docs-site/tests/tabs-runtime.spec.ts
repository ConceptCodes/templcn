import { expect, test } from "@playwright/test"

test.beforeEach(async ({ page }) => {
  await page.goto("/tests/fixtures/tabs-runtime.html")
})

test("initializes default tab state", async ({ page }) => {
  await expect(page.locator("#account-tab")).toHaveAttribute("data-state", "active")
  await expect(page.locator("#account-tab")).toHaveAttribute("aria-selected", "true")
  await expect(page.locator("#account-tab")).toHaveAttribute("tabindex", "0")
  await expect(page.locator("#account-panel")).toHaveAttribute("data-state", "active")

  await expect(page.locator("#password-tab")).toHaveAttribute("data-state", "inactive")
  await expect(page.locator("#password-panel")).toHaveAttribute("hidden", "")
})

test("click activates tab and panel", async ({ page }) => {
  await page.locator("#password-tab").click()

  await expect(page.locator("#password-tab")).toHaveAttribute("data-state", "active")
  await expect(page.locator("#password-panel")).toHaveAttribute("data-state", "active")
  await expect(page.locator("#password-panel")).not.toHaveAttribute("hidden", "")
  await expect(page.locator("#account-panel")).toHaveAttribute("hidden", "")
})

test("arrow keys rove and activate enabled tabs", async ({ page }) => {
  await page.locator("#account-tab").focus()
  await page.keyboard.press("ArrowRight")

  await expect(page.locator("#password-tab")).toBeFocused()
  await expect(page.locator("#password-tab")).toHaveAttribute("data-state", "active")

  await page.keyboard.press("ArrowRight")
  await expect(page.locator("#account-tab")).toBeFocused()
  await expect(page.locator("#account-tab")).toHaveAttribute("data-state", "active")
})

