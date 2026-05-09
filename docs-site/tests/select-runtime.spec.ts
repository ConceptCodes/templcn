import { expect, test } from "@playwright/test"

test.beforeEach(async ({ page }) => {
  await page.goto("/tests/fixtures/select-runtime.html")
})

test("initializes default value, visible label, selected state, and hidden input", async ({ page }) => {
  await expect(page.locator("#plan-select")).toHaveAttribute("data-value", "starter")
  await expect(page.locator('[data-slot="select-value"]')).toHaveText("Starter")
  await expect(page.locator("#starter-option")).toHaveAttribute("aria-selected", "true")
  await expect(page.locator('input[type="hidden"][name="plan"]')).toHaveValue("starter")
})

test("opens from trigger, selects an item, syncs form value, and restores focus", async ({ page }) => {
  const trigger = page.locator("#plan-trigger")
  const root = page.locator("#plan-select")

  await trigger.click()
  await expect(root).toHaveAttribute("data-state", "open")
  await expect(trigger).toHaveAttribute("aria-expanded", "true")
  await expect(page.locator('[data-slot="select-content"]')).not.toHaveAttribute("hidden", "")

  await page.locator("#pro-option").click()

  await expect(root).toHaveAttribute("data-state", "closed")
  await expect(root).toHaveAttribute("data-value", "pro")
  await expect(trigger).toHaveAttribute("aria-expanded", "false")
  await expect(page.locator('[data-slot="select-value"]')).toHaveText("Pro")
  await expect(page.locator("#pro-option")).toHaveAttribute("aria-selected", "true")
  await expect(page.locator("#starter-option")).toHaveAttribute("aria-selected", "false")
  await expect(page.locator('input[type="hidden"][name="plan"]')).toHaveValue("pro")
  await expect(trigger).toBeFocused()
})

test("keyboard opens, moves focus, selects, and Escape closes", async ({ page }) => {
  const trigger = page.locator("#plan-trigger")
  const root = page.locator("#plan-select")

  await trigger.focus()
  await page.keyboard.press("ArrowDown")
  await expect(root).toHaveAttribute("data-state", "open")
  await expect(page.locator("#starter-option")).toBeFocused()

  await page.keyboard.press("ArrowDown")
  await expect(page.locator("#pro-option")).toBeFocused()

  await page.keyboard.press("Enter")
  await expect(root).toHaveAttribute("data-value", "pro")
  await expect(root).toHaveAttribute("data-state", "closed")
  await expect(trigger).toBeFocused()

  await page.keyboard.press("ArrowDown")
  await page.keyboard.press("Escape")
  await expect(root).toHaveAttribute("data-state", "closed")
  await expect(trigger).toBeFocused()
})

test("typeahead focuses matching option", async ({ page }) => {
  await page.locator("#plan-trigger").focus()
  await page.keyboard.press("ArrowDown")
  await page.keyboard.press("e")
  await expect(page.locator("#enterprise-option")).toBeFocused()
})

