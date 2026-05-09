import { expect, test } from "@playwright/test"

test.beforeEach(async ({ page }) => {
  await page.goto("/tests/fixtures/dialog-family-runtime.html")
})

for (const spec of [
  { name: "alert dialog", trigger: "#alert-trigger", dialog: "#delete-alert", close: "#alert-cancel", role: "alertdialog" },
  { name: "sheet", trigger: "#sheet-trigger", dialog: "#settings-sheet", close: "#sheet-close", role: "dialog" },
  { name: "drawer", trigger: "#drawer-trigger", dialog: "#nav-drawer", close: "#drawer-close", role: "dialog" },
]) {
  test(`${spec.name} opens, focuses content, closes, and restores focus`, async ({ page }) => {
    const trigger = page.locator(spec.trigger)
    const dialog = page.locator(spec.dialog)

    await expect(dialog).toHaveAttribute("role", spec.role)
    await trigger.click()
    await expect(dialog).toHaveAttribute("data-state", "open")
    await expect(trigger).toHaveAttribute("aria-expanded", "true")
    await expect(page.locator(`${spec.dialog} [data-slot$="-content"]`)).toHaveAttribute("data-state", "open")

    await page.locator(spec.close).click()
    await expect(dialog).toHaveAttribute("data-state", "closed")
    await expect(trigger).toHaveAttribute("aria-expanded", "false")
    await expect(trigger).toBeFocused()
  })
}
