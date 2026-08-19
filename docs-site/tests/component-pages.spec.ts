import { expect, test } from "@playwright/test"

test("every component documentation page renders without browser errors", async ({ page }) => {
	test.setTimeout(120_000)
	page.setDefaultNavigationTimeout(5_000)
  const browserErrors: string[] = []
  page.on("pageerror", (error) => browserErrors.push(error.message))
  page.on("console", (message) => {
    if (message.type() === "error") browserErrors.push(message.text())
  })

  await page.goto("/docs/components", { waitUntil: "domcontentloaded" })
  const hrefs = await page.locator('a[href^="/docs/components/"]').evaluateAll((links) =>
    [...new Set(links
      .map((link) => link.getAttribute("href"))
      .filter((href): href is string => Boolean(href) && href.split("/").length === 4))],
  )

  expect(hrefs.length).toBeGreaterThan(0)
  const failures: string[] = []
  for (const href of hrefs) {
    browserErrors.length = 0
    try {
      const response = await page.goto(href, { waitUntil: "domcontentloaded" })
      if (response?.status() !== 200) failures.push(`${href}: HTTP ${response?.status()}`)
      if (!(await page.locator("h1").first().isVisible())) failures.push(`${href}: missing h1`)
      if (!(await page.locator("[data-slot], button, input, select, textarea").first().isVisible())) failures.push(`${href}: missing rendered component`)
      if (browserErrors.length) failures.push(`${href}: ${browserErrors.join(" | ")}`)
    } catch (error) {
      failures.push(`${href}: ${error instanceof Error ? error.message : String(error)}`)
    }
  }
  expect(failures).toEqual([])
})
