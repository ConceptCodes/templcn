import { expect, test } from "@playwright/test"

test.beforeEach(async ({ page }) => {
  await page.goto("/tests/fixtures/dropdown-menu-runtime.html")
})

test("opens and closes from trigger", async ({ page }) => {
  const trigger = page.locator("#menu-trigger")
  const root = page.locator("#account-menu")
  const content = page.locator('[data-slot="dropdown-menu-content"]')

  await trigger.click()
  await expect(root).toHaveAttribute("data-state", "open")
  await expect(trigger).toHaveAttribute("aria-expanded", "true")
  await expect(content).not.toHaveAttribute("hidden", "")

  await trigger.click()
  await expect(root).toHaveAttribute("data-state", "closed")
  await expect(trigger).toHaveAttribute("aria-expanded", "false")
  await expect(content).toHaveAttribute("hidden", "")
})

test("keyboard opens, roves focus, selects item, and restores focus", async ({ page }) => {
  const trigger = page.locator("#menu-trigger")
  const root = page.locator("#account-menu")

  await trigger.focus()
  await page.keyboard.press("ArrowDown")
  await expect(root).toHaveAttribute("data-state", "open")
  await expect(page.locator("#profile-item")).toBeFocused()

  await page.keyboard.press("ArrowDown")
  await expect(page.locator("#billing-item")).toBeFocused()

  await page.keyboard.press("Enter")
  await expect(root).toHaveAttribute("data-state", "closed")
  await expect(trigger).toBeFocused()
})

test("Escape and outside click dismiss the menu", async ({ page }) => {
  const trigger = page.locator("#menu-trigger")
  const root = page.locator("#account-menu")

  await trigger.click()
  await page.keyboard.press("Escape")
  await expect(root).toHaveAttribute("data-state", "closed")
  await expect(trigger).toBeFocused()

  await trigger.click()
  await page.locator("#outside-menu").click()
  await expect(root).toHaveAttribute("data-state", "closed")
})

