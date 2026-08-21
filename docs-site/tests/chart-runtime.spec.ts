import { expect, test } from "@playwright/test"

test.describe("Chart component and tooltips", () => {
  test("area chart renders and displays tooltip on hover", async ({ page }) => {
    await page.goto("/charts/area", { waitUntil: "domcontentloaded" })
    
    // Check if iframe is rendered
    const iframeElement = page.locator("iframe").first()
    await expect(iframeElement).toBeVisible()

    const frame = page.frameLocator("iframe").first()
    const chartContainer = frame.locator('[data-chart="area"]').first()
    await expect(chartContainer).toBeVisible()

    // Hover over the chart svg
    const svg = chartContainer.locator("svg").first()
    await expect(svg).toBeVisible()

    const box = await svg.boundingBox()
    expect(box).not.toBeNull()
    if (box) {
      await page.mouse.move(box.x + box.width / 2, box.y + box.height / 2)
    }

    // Tooltip should be visible in the frame
    const tooltip = frame.locator(".ts-chart-tooltip")
    await expect(tooltip).toBeVisible({ timeout: 5000 })
  })

  test("bar chart renders and displays tooltip on hover", async ({ page }) => {
    await page.goto("/charts/bar", { waitUntil: "domcontentloaded" })

    const frame = page.frameLocator("iframe").first()
    const chartContainer = frame.locator('[data-chart="bar"]').first()
    await expect(chartContainer).toBeVisible()

    const svg = chartContainer.locator("svg").first()
    await expect(svg).toBeVisible()

    const box = await svg.boundingBox()
    expect(box).not.toBeNull()
    if (box) {
      await page.mouse.move(box.x + box.width / 2, box.y + box.height / 2)
    }

    const tooltip = frame.locator(".ts-chart-tooltip")
    await expect(tooltip).toBeVisible({ timeout: 5000 })
  })

  test("line chart renders and displays tooltip on hover", async ({ page }) => {
    await page.goto("/charts/line", { waitUntil: "domcontentloaded" })

    const frame = page.frameLocator("iframe").first()
    const chartContainer = frame.locator('[data-chart="line"]').first()
    await expect(chartContainer).toBeVisible()

    const svg = chartContainer.locator("svg").first()
    await expect(svg).toBeVisible()

    const box = await svg.boundingBox()
    expect(box).not.toBeNull()
    if (box) {
      await page.mouse.move(box.x + box.width / 2, box.y + box.height / 2)
    }

    const tooltip = frame.locator(".ts-chart-tooltip")
    await expect(tooltip).toBeVisible({ timeout: 5000 })
  })

  test("pie chart renders and displays tooltip on hover", async ({ page }) => {
    await page.goto("/charts/pie", { waitUntil: "domcontentloaded" })

    const frame = page.frameLocator("iframe").first()
    const chartContainer = frame.locator('[data-chart="pie"]').first()
    await expect(chartContainer).toBeVisible()

    const svg = chartContainer.locator("svg").first()
    await expect(svg).toBeVisible()

    const box = await svg.boundingBox()
    expect(box).not.toBeNull()
    if (box) {
      await page.mouse.move(box.x + box.width * 0.6, box.y + box.height * 0.4)
    }

    const tooltip = frame.locator(".ts-chart-tooltip")
    await expect(tooltip).toBeVisible({ timeout: 5000 })
  })

  test("radial chart renders and displays tooltip on hover", async ({ page }) => {
    await page.goto("/charts/radial", { waitUntil: "domcontentloaded" })

    const frame = page.frameLocator("iframe").first()
    const chartContainer = frame.locator('[data-chart="radial"]').first()
    await expect(chartContainer).toBeVisible()

    const arc = chartContainer.locator("path").first()
    await expect(arc).toBeVisible()
    await arc.hover({ force: true })

    const tooltip = frame.locator(".ts-chart-tooltip")
    await expect(tooltip).toBeVisible({ timeout: 5000 })
  })
})
