import { expect, test } from "@playwright/test"

test.beforeEach(async ({ page }) => {
  await page.goto("/tests/fixtures/calendar-runtime.html")
})

test("calendar syncs selected day and moves focus with arrow keys", async ({ page }) => {
  const root = page.locator("#calendar")

  await page.locator("#may-08").click()
  await expect(root).toHaveAttribute("data-selected", "2026-05-08")
  await expect(page.locator("#may-08")).toHaveAttribute("aria-selected", "true")
  await expect(page.locator("#may-07")).toHaveAttribute("aria-selected", "false")

  await page.locator("#may-07").focus()
  await page.keyboard.press("ArrowRight")
  await expect(page.locator("#may-08")).toBeFocused()
})

test("calendar month controls update current month state", async ({ page }) => {
  const root = page.locator("#calendar")

  await page.locator('[data-slot="calendar-next"]').click()
  await expect(root).toHaveAttribute("data-current-month", "2026-06")

  await page.locator('[data-slot="calendar-prev"]').click()
  await expect(root).toHaveAttribute("data-current-month", "2026-05")
})
