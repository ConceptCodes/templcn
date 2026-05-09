import { expect, test } from "@playwright/test"

test.beforeEach(async ({ page }) => {
  await page.goto("/tests/fixtures/longtail-runtime.html")
})

test("checkbox and switch sync checked aria state", async ({ page }) => {
  await page.locator("#newsletter").check()
  await expect(page.locator("#newsletter")).toHaveAttribute("data-state", "checked")
  await expect(page.locator("#newsletter")).toHaveAttribute("aria-checked", "true")

  await page.locator("#dark-mode").check()
  await expect(page.locator("#dark-mode")).toHaveAttribute("data-state", "checked")
  await expect(page.locator("#dark-mode")).toHaveAttribute("aria-checked", "true")
})

test("input otp syncs hidden value from slots", async ({ page }) => {
  await expect(page.locator('input[type="hidden"][name="pin"]')).toHaveValue("12")
  await page.locator("#pin-2").fill("9")
  await page.locator("#pin-2").dispatchEvent("input")
  await expect(page.locator("#pin")).toHaveAttribute("data-value", "19")
  await expect(page.locator('input[type="hidden"][name="pin"]')).toHaveValue("19")
})

test("carousel next and previous update active slide", async ({ page }) => {
  await expect(page.locator("#slide-2")).toHaveAttribute("hidden", "")
  await page.locator('[data-slot="carousel-next"]').click()
  await expect(page.locator("#carousel")).toHaveAttribute("data-index", "1")
  await expect(page.locator("#slide-1")).toHaveAttribute("hidden", "")
  await expect(page.locator("#slide-2")).not.toHaveAttribute("hidden", "")

  await page.locator('[data-slot="carousel-previous"]').click()
  await expect(page.locator("#carousel")).toHaveAttribute("data-index", "0")
})

test("resizable handle updates adjacent panel sizes", async ({ page }) => {
  const before = Number(await page.locator("#left-panel").getAttribute("data-size"))
  await page.locator("#handle").dragTo(page.locator("#right-panel"), { targetPosition: { x: 100, y: 20 } })
  const after = Number(await page.locator("#left-panel").getAttribute("data-size"))
  expect(after).toBeGreaterThan(before)
})

test("toast close and sidebar trigger update state", async ({ page }) => {
  await page.locator('[data-slot="toast-close"]').click()
  await expect(page.locator("#toast")).toHaveAttribute("data-state", "closed")
  await expect(page.locator("#toast")).toHaveAttribute("hidden", "")

  await expect(page.locator("#sidebar-shell")).toHaveAttribute("data-state", "open")
  await page.locator('[data-slot="sidebar-trigger"]').click()
  await expect(page.locator("#sidebar-shell")).toHaveAttribute("data-state", "closed")
  await expect(page.locator('[data-slot="sidebar-trigger"]')).toHaveAttribute("aria-expanded", "false")
})

test("date picker day selection syncs root and hidden input", async ({ page }) => {
  await page.locator('[data-date="2026-05-08"]').click()
  await expect(page.locator("#date-picker")).toHaveAttribute("data-value", "2026-05-08")
  await expect(page.locator('input[type="hidden"][name="date"]')).toHaveValue("2026-05-08")
})
