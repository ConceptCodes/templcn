import { expect, test } from "@playwright/test"

test.beforeEach(async ({ page }) => {
  await page.goto("/tests/fixtures/dialog-runtime.html")
})

test("opens from trigger, focuses first control, and restores focus on Escape", async ({ page }) => {
  const trigger = page.locator("#open-dialog")
  const dialog = page.locator("#settings-dialog")
  const content = page.locator('[data-slot="dialog-content"]')

  await trigger.focus()
  await trigger.click()

  await expect(dialog).toHaveJSProperty("open", true)
  await expect(dialog).toHaveAttribute("data-state", "open")
  await expect(trigger).toHaveAttribute("aria-expanded", "true")
  await expect(content).toHaveAttribute("data-state", "open")
  await expect(page.locator("#first-field")).toBeFocused()

  await page.keyboard.press("Escape")

  await expect(dialog).toHaveJSProperty("open", false)
  await expect(dialog).toHaveAttribute("data-state", "closed")
  await expect(trigger).toHaveAttribute("aria-expanded", "false")
  await expect(trigger).toBeFocused()
})

test("close button closes the dialog and restores trigger focus", async ({ page }) => {
  const trigger = page.locator("#open-dialog")
  const dialog = page.locator("#settings-dialog")

  await trigger.click()
  await page.locator("#close-dialog").click()

  await expect(dialog).toHaveJSProperty("open", false)
  await expect(dialog).toHaveAttribute("data-state", "closed")
  await expect(trigger).toBeFocused()
})

test("modal tab loop wraps within the dialog", async ({ page }) => {
  await page.locator("#open-dialog").click()

  await page.locator("#first-field").focus()
  await page.keyboard.press("Shift+Tab")
  await expect(page.locator("#close-dialog")).toBeFocused()

  await page.keyboard.press("Tab")
  await expect(page.locator("#first-field")).toBeFocused()
})
