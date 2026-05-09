import { expect, test } from "@playwright/test"

test.beforeEach(async ({ page }) => {
  await page.goto("/tests/fixtures/menu-family-runtime.html")
})

test("context menu opens at pointer position, roves focus, typeaheads, and closes on escape", async ({ page }) => {
  const root = page.locator("#canvas-menu")
  const content = page.locator("#context-content")

  await root.click({ button: "right", position: { x: 40, y: 30 } })
  await expect(root).toHaveAttribute("data-state", "open")
  await expect(content).not.toHaveAttribute("hidden", "")
  await expect(page.locator("#copy-item")).toBeFocused()

  await page.keyboard.press("ArrowDown")
  await expect(page.locator("#paste-item")).toBeFocused()

  await page.keyboard.press("c")
  await expect(page.locator("#copy-item")).toBeFocused()

  await page.keyboard.press("Escape")
  await expect(root).toHaveAttribute("data-state", "closed")
  await expect(content).toHaveAttribute("hidden", "")
})

test("menubar opens from keyboard and moves between top-level triggers", async ({ page }) => {
  const root = page.locator("#app-menubar")
  const trigger = page.locator("#file-trigger")
  const content = page.locator("#file-menu")

  await trigger.focus()
  await page.keyboard.press("Enter")
  await expect(root).toHaveAttribute("data-state", "open")
  await expect(trigger).toHaveAttribute("aria-expanded", "true")
  await expect(content).not.toHaveAttribute("hidden", "")
  await expect(page.locator("#new-item")).toBeFocused()

  await page.keyboard.press("Escape")
  await expect(root).toHaveAttribute("data-state", "closed")
  await expect(trigger).toBeFocused()

  await page.keyboard.press("ArrowRight")
  await expect(page.locator("#edit-trigger")).toBeFocused()
})

test("navigation menu trigger toggles content and outside click closes it", async ({ page }) => {
  const root = page.locator("#site-nav")
  const trigger = page.locator("#products-trigger")
  const content = page.locator("#products-content")

  await trigger.click()
  await expect(root).toHaveAttribute("data-state", "open")
  await expect(trigger).toHaveAttribute("aria-expanded", "true")
  await expect(content).not.toHaveAttribute("hidden", "")

  await page.locator("#outside").click()
  await expect(root).toHaveAttribute("data-state", "closed")
  await expect(trigger).toHaveAttribute("aria-expanded", "false")
  await expect(content).toHaveAttribute("hidden", "")
})
