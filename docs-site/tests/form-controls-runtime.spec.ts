import { expect, test } from "@playwright/test"

test.beforeEach(async ({ page }) => {
  await page.goto("/tests/fixtures/form-controls-runtime.html")
})

test("radio group syncs checked state and arrow key selection", async ({ page }) => {
  const root = page.locator("#density-radio")

  await page.locator("#compact-item input").check()
  await expect(root).toHaveAttribute("data-value", "compact")
  await expect(page.locator("#compact-item")).toHaveAttribute("data-state", "checked")
  await expect(page.locator("#comfortable-item")).toHaveAttribute("data-state", "unchecked")

  await page.locator("#compact-item input").focus()
  await page.keyboard.press("ArrowDown")
  await expect(root).toHaveAttribute("data-value", "comfortable")
  await expect(page.locator("#comfortable-item input")).toBeFocused()
})

test("single toggle group keeps one active item and roves focus", async ({ page }) => {
  const root = page.locator("#align-toggle")

  await page.locator("#center-toggle").click()
  await expect(root).toHaveAttribute("data-value", "center")
  await expect(page.locator("#center-toggle")).toHaveAttribute("aria-pressed", "true")
  await expect(page.locator("#left-toggle")).toHaveAttribute("aria-pressed", "false")

  await page.keyboard.press("ArrowRight")
  await expect(page.locator("#right-toggle")).toBeFocused()
})

test("slider updates data value and aria-valuenow on input", async ({ page }) => {
  const slider = page.locator("#volume-slider")
  const input = page.locator("#volume-input")

  await input.fill("65")
  await input.dispatchEvent("input")

  await expect(slider).toHaveAttribute("data-value", "65")
  await expect(input).toHaveAttribute("aria-valuenow", "65")
  await expect(input).toHaveValue("65")
})
