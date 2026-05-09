import { expect, test } from "@playwright/test"

test.beforeEach(async ({ page }) => {
  await page.goto("/tests/fixtures/command-combobox-runtime.html")
})

test("command filters items, shows empty state, and selects highlighted item", async ({ page }) => {
  const input = page.locator("#command-input")
  let selected = ""
  await page.locator("#beta-command").evaluate((node) => {
    node.addEventListener("click", () => document.body.setAttribute("data-selected-command", "beta"))
  })

  await input.fill("be")
  await expect(page.locator("#alpha-command")).toHaveAttribute("hidden", "")
  await expect(page.locator("#beta-command")).not.toHaveAttribute("hidden", "")
  await expect(page.locator("#command-empty")).toHaveAttribute("hidden", "")

  await page.keyboard.press("ArrowDown")
  await expect(page.locator("#beta-command")).toBeFocused()
  await page.keyboard.press("Enter")
  selected = await page.locator("body").getAttribute("data-selected-command") ?? ""
  expect(selected).toBe("beta")

  await input.fill("zzz")
  await expect(page.locator("#command-empty")).not.toHaveAttribute("hidden", "")
})

test("combobox initializes value and supports keyboard selection through shared selectable runtime", async ({ page }) => {
  const root = page.locator("#framework-combobox")
  const trigger = page.locator("#combobox-trigger")

  await expect(page.locator('[data-slot="combobox-value"]')).toHaveText("Go")
  await expect(page.locator('input[type="hidden"][name="framework"]')).toHaveValue("go")

  await trigger.focus()
  await page.keyboard.press("ArrowDown")
  await expect(root).toHaveAttribute("data-state", "open")
  await page.keyboard.press("ArrowDown")
  await expect(page.locator("#templ-option")).toBeFocused()
  await page.keyboard.press("Enter")

  await expect(root).toHaveAttribute("data-state", "closed")
  await expect(page.locator('[data-slot="combobox-value"]')).toHaveText("templ")
  await expect(page.locator('input[type="hidden"][name="framework"]')).toHaveValue("templ")
})
