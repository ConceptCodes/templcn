import { expect, test } from "@playwright/test"

test.beforeEach(async ({ page }) => {
  await page.goto("/tests/fixtures/floating-runtime.html")
})

test("popover opens, positions, closes outside, and restores focus on escape", async ({ page }) => {
  const root = page.locator("#profile-popover")
  const trigger = page.locator("#popover-trigger")
  const content = page.locator("#popover-content")

  await trigger.click()
  await expect(root).toHaveAttribute("data-state", "open")
  await expect(trigger).toHaveAttribute("aria-expanded", "true")
  await expect(content).not.toHaveAttribute("hidden", "")
  await expect(content).toHaveAttribute("data-side", /bottom|top/)

  const box = await content.boundingBox()
  expect(box?.x).toBeGreaterThan(0)
  expect(box?.y).toBeGreaterThan(0)

  await page.keyboard.press("Escape")
  await expect(root).toHaveAttribute("data-state", "closed")
  await expect(trigger).toHaveAttribute("aria-expanded", "false")
  await expect(content).toHaveAttribute("hidden", "")
  await expect(trigger).toBeFocused()

  await trigger.click()
  await page.locator("#outside").click()
  await expect(root).toHaveAttribute("data-state", "closed")
})

test("tooltip opens on hover and focus, then closes on pointer leave and escape", async ({ page }) => {
  const root = page.locator("#info-tooltip")
  const trigger = page.locator("#tooltip-trigger")
  const content = page.locator("#tooltip-content")

  await trigger.hover()
  await expect(root).toHaveAttribute("data-state", "open")
  await expect(content).not.toHaveAttribute("hidden", "")
  await expect(content).toHaveAttribute("role", "tooltip")

  await page.locator("#outside").hover()
  await expect(root).toHaveAttribute("data-state", "closed")
  await expect(content).toHaveAttribute("hidden", "")

  await trigger.focus()
  await expect(root).toHaveAttribute("data-state", "open")
  await page.keyboard.press("Escape")
  await expect(root).toHaveAttribute("data-state", "closed")
})

test("hover card opens on hover and closes when pointer leaves the root", async ({ page }) => {
  const root = page.locator("#user-hover-card")
  const content = page.locator("#hover-content")

  await page.locator("#hover-trigger").hover()
  await expect(root).toHaveAttribute("data-state", "open")
  await expect(content).not.toHaveAttribute("hidden", "")

  await page.locator("#outside").hover()
  await expect(root).toHaveAttribute("data-state", "closed")
  await expect(content).toHaveAttribute("hidden", "")
})
