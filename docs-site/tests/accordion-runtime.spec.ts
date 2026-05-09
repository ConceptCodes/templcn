import { expect, test } from "@playwright/test"

test.beforeEach(async ({ page }) => {
  await page.goto("/tests/fixtures/accordion-runtime.html")
})

test("initializes default open item and content state", async ({ page }) => {
  await expect(page.locator("#billing-item")).toHaveAttribute("open", "")
  await expect(page.locator("#billing-item")).toHaveAttribute("data-state", "open")
  await expect(page.locator("#billing-content")).toHaveAttribute("data-state", "open")
  await expect(page.locator("#account-content")).toHaveAttribute("data-state", "closed")
})

test("single accordion closes previously open item", async ({ page }) => {
  await page.locator("#account-trigger").click()
  await expect(page.locator("#account-item")).toHaveAttribute("open", "")
  await expect(page.locator("#account-content")).toHaveAttribute("data-state", "open")
  await expect(page.locator("#billing-item")).not.toHaveAttribute("open", "")
  await expect(page.locator("#billing-content")).toHaveAttribute("data-state", "closed")
})

test("arrow keys move between triggers", async ({ page }) => {
  await page.locator("#account-trigger").focus()
  await page.keyboard.press("ArrowDown")
  await expect(page.locator("#billing-trigger")).toBeFocused()
  await page.keyboard.press("End")
  await expect(page.locator("#security-trigger")).toBeFocused()
  await page.keyboard.press("Home")
  await expect(page.locator("#account-trigger")).toBeFocused()
})

