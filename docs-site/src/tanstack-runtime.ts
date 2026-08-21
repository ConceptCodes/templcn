import { areaY, barX, barY, defineChart, lineY, stack } from "@tanstack/charts"
import { mountChart } from "@tanstack/charts/dom"
import { d3Curve } from "@tanstack/charts/d3/shape"
import { focusGroupAngle, pie, polar, radialArc } from "@tanstack/charts/polar"
import { scaleBand } from "@tanstack/charts/scales/band"
import { scaleLinear } from "@tanstack/charts/scales/linear"
import { tooltip } from "@tanstack/charts/tooltip"
import { curveStep } from "d3-shape"
import { initTanStackTables } from "./tanstack-table"

type Series = { key: string; label?: string; color?: string }
type ChartConfig = {
  type?: "area" | "bar" | "line" | "pie" | "radial"
  data?: Record<string, unknown>[]
  series?: Series[]
  orientation?: "vertical" | "horizontal"
  stacked?: boolean
  step?: boolean
  ariaLabel?: string
  height?: number
  innerRadius?: number
  centerLabel?: string
  centerCaption?: string
  value?: number
  max?: number
  segments?: { label: string; value: number }[]
}

const CHART_TOKEN_COUNT = 5

const THEME_TOKENS = [
  "background", "foreground",
  "card", "card-foreground",
  "popover", "popover-foreground",
  "primary", "primary-foreground",
  "secondary", "secondary-foreground",
  "muted", "muted-foreground",
  "accent", "accent-foreground",
  "destructive", "destructive-foreground",
  "border", "input", "ring", "radius",
  "chart-1", "chart-2", "chart-3", "chart-4", "chart-5",
  "sidebar", "sidebar-foreground",
  "sidebar-primary", "sidebar-primary-foreground",
  "sidebar-accent", "sidebar-accent-foreground",
  "sidebar-border", "sidebar-ring",
]

let parentThemeRoot: HTMLElement | null = null

/**
 * Charts live in srcdoc iframes that load globals.css independently, so the
 * parent document's dark class, preset attribute, and generated preset style
 * never reach them. Mirror the parent's resolved theme tokens into the iframe
 * and keep them in sync while the page is open.
 */
function collectParentTokens(parentRoot: HTMLElement): string {
  const computed = getComputedStyle(parentRoot)
  const declarations: string[] = []
  for (const token of THEME_TOKENS) {
    const value = computed.getPropertyValue(`--${token}`).trim()
    if (value) declarations.push(`  --${token}: ${value};`)
  }
  return `:root {\n${declarations.join("\n")}\n}`
}

function applyThemeTokens(css: string) {
  let style = document.getElementById("templcn-theme-sync")
  if (!style) {
    style = document.createElement("style")
    style.id = "templcn-theme-sync"
    document.head.appendChild(style)
  }
  style.textContent = css
}

function syncIframeThemeNow() {
  if (!parentThemeRoot) return
  const root = document.documentElement
  root.classList.toggle("dark", parentThemeRoot.classList.contains("dark"))
  const preset = parentThemeRoot.getAttribute("data-theme-preset")
  if (preset) root.setAttribute("data-theme-preset", preset)
  applyThemeTokens(collectParentTokens(parentThemeRoot))
}

/**
 * The parent applies its saved theme during its own DOMContentLoaded, which
 * can happen after the iframe boots. data-theme-mode is only present once the
 * parent theme runtime has initialized, so wait for it before snapshotting.
 */
function waitForParentTheme(timeoutMs = 3000): Promise<HTMLElement | null> {
  return new Promise((resolve) => {
    if (window.parent === window) return resolve(null)
    let root: HTMLElement
    try {
      root = window.parent.document.documentElement
    } catch {
      return resolve(null)
    }
    const started = performance.now()
    const poll = () => {
      if (root.hasAttribute("data-theme-mode")) return resolve(root)
      if (performance.now() - started >= timeoutMs) return resolve(root)
      setTimeout(poll, 50)
    }
    poll()
  })
}

function observeParentTheme(root: HTMLElement) {
  const observer = new MutationObserver(syncIframeThemeNow)
  observer.observe(root, {
    attributes: true,
    attributeFilter: ["class", "data-theme-preset", "data-theme-mode"],
  })
}

function initIframeThemeSync() {
  if (window.parent === window) return
  void waitForParentTheme().then((root) => {
    if (!root) return
    parentThemeRoot = root
    syncIframeThemeNow()
    observeParentTheme(root)
  })
}

/**
 * TanStack's default categorical scheme paints marks with var(--ts-chart-N, …)
 * fallbacks. Redefine those custom properties on the chart host so every mark
 * resolves through the site's --chart-N tokens (light, dark, and presets).
 */
function applyChartTokenVars(element: HTMLElement) {
  for (let i = 1; i <= CHART_TOKEN_COUNT; i++) {
    element.style.setProperty(`--ts-chart-${i}`, `var(--chart-${i})`)
  }
  element.style.setProperty("--ts-chart-tooltip-background", "var(--popover)")
  element.style.setProperty("--ts-chart-tooltip-color", "var(--popover-foreground)")
  element.style.setProperty("--ts-chart-tooltip-border", "1px solid var(--border)")
  element.style.setProperty("--ts-chart-tooltip-border-radius", "calc(var(--radius) * 0.8)")
  element.style.setProperty("--ts-chart-tooltip-shadow", "0 4px 6px -1px rgb(0 0 0 / 0.1), 0 2px 4px -2px rgb(0 0 0 / 0.1)")
}

function chartColorRange(): string[] {
  return Array.from({ length: CHART_TOKEN_COUNT }, (_, i) => `var(--chart-${i + 1})`)
}

function readConfig(element: HTMLElement): ChartConfig | null {
  const raw = element.dataset.chartConfig || element.dataset.data
  if (!raw) return null
  try {
    return JSON.parse(raw) as ChartConfig
  } catch {
    console.warn("templcn: invalid TanStack chart configuration", element)
    return null
  }
}

function renderCartesianChart(element: HTMLElement, config: ChartConfig, type: "area" | "bar" | "line") {
  const data = config.data!
  const series = config.series!
  const horizontal = config.orientation === "horizontal"

  const marks = series.map((entry, index) => {
    const channel = horizontal
      ? { x: entry.key, y: "label" }
      : { x: "label", y: entry.key }
    const color = entry.color || `var(--chart-${(index % CHART_TOKEN_COUNT) + 1})`
    const markOptions = {
      ...channel,
      stroke: color,
      fill: color,
      z: () => entry.label || entry.key,
      ...(type === "line" && config.step ? { curve: d3Curve(curveStep) } : {}),
    } as const

    if (type === "bar") {
      return horizontal
        ? barX(data, markOptions)
        : barY(data, markOptions)
    }
    if (type === "area") {
      if (horizontal) throw new Error("horizontal area charts are not supported")
      return areaY(data, { ...markOptions, ...(config.stacked && series.length > 1 ? { layout: stack(), fillOpacity: 0.35 } : {}) })
    }
    if (horizontal) throw new Error("horizontal line charts are not supported")
    return lineY(data, markOptions)
  })

  const categoryAxis = { scale: () => scaleBand<string>().padding(0.1), grid: false }
  const valueAxis = { scale: scaleLinear, nice: true, grid: true }

  const definition = defineChart({
    marks,
    x: horizontal ? valueAxis : categoryAxis,
    y: horizontal ? categoryAxis : valueAxis,
    focus: horizontal ? (series.length > 1 ? "group-y" : "nearest-y") : (series.length > 1 ? "group-x" : "nearest-x"),
    tooltip,
  })

  return mountChart(element, {
    definition,
    height: config.height || 288,
    ariaLabel: config.ariaLabel || element.getAttribute("aria-label") || `${type} chart`,
  })
}

function renderChart(element: HTMLElement) {
  const config = readConfig(element)
  if (!config) return

  const type = config.type || element.dataset.chart || "line"

  if (type === "pie" || type === "radial") {
    return renderPolarChart(element, config, type)
  }

  if (!config.data?.length || !config.series?.length) return

  return renderCartesianChart(element, config, type as "area" | "bar" | "line")
}

function renderPolarChart(element: HTMLElement, config: ChartConfig, type: "pie" | "radial") {
  const source = type === "pie"
    ? (config.data || []).map((row) => ({
        label: String(row.label || ""),
        value: Number(row.value || 0),
      }))
    : config.segments?.length
      ? config.segments.map((row) => ({ label: row.label, value: Number(row.value || 0) }))
      : [
          { label: "Complete", value: Math.max(0, Math.min(Number(config.max || 100), Number(config.value || 0))) },
          { label: "Remaining", value: Math.max(0, Number(config.max || 100) - Number(config.value || 0)) },
        ]

  if (!source.length || !source.some((row) => row.value > 0)) return

  const slices = pie(source, { value: "value", gapAngle: 0.025 })
  const labels = source.map((row) => row.label)
  const innerRadius = Math.max(0, Math.min(0.9, Number(config.innerRadius || (type === "radial" ? 0.68 : 0))))
  const marks = [
    radialArc(slices, {
      innerRadius: ({ radius }) => radius * innerRadius,
      outerRadius: ({ radius }) => radius * 0.9,
      cornerRadius: 4,
      color: "label",
      key: "label",
      z: "label",
    }),
  ]

  const definition = defineChart({
    marks: [polar({ inset: 8, marks })],
    color: {
      domain: labels,
      range: chartColorRange(),
    },
    focus: focusGroupAngle,
    tooltip,
  })

  const host = mountChart(element, {
    definition,
    height: config.height || 288,
    ariaLabel: config.ariaLabel || element.getAttribute("aria-label") || `${type} chart`,
  })

  if (config.centerLabel) {
    element.style.position = "relative"
    const overlay = document.createElement("div")
    overlay.setAttribute("aria-hidden", "true")
    overlay.style.cssText = "position:absolute;inset:0;display:flex;flex-direction:column;align-items:center;justify-content:center;pointer-events:none;text-align:center"
    const label = document.createElement("span")
    label.textContent = config.centerLabel
    label.style.cssText = "font-size:22px;font-weight:700;line-height:1.2"
    overlay.appendChild(label)
    if (config.centerCaption) {
      const caption = document.createElement("span")
      caption.textContent = config.centerCaption
      caption.style.cssText = "font-size:11px;opacity:0.6"
      overlay.appendChild(caption)
    }
    element.appendChild(overlay)
  }

  return host
}

export function initTanStackCharts(root: ParentNode = document) {
  initIframeThemeSync()
  root.querySelectorAll<HTMLElement>('[data-engine="tanstack"]:not([data-tanstack-mounted])').forEach((element) => {
    try {
      applyChartTokenVars(element)
      const host = renderChart(element)
      if (!host) return
      element.dataset.tanstackMounted = "true"
    } catch (error) {
      console.error("templcn: failed to render chart", element, error)
    }
  })
}

export function initTanStackRuntime(root: ParentNode = document) {
  initTanStackCharts(root)
  initTanStackTables(root)
}

if (document.readyState === "loading") {
  document.addEventListener("DOMContentLoaded", () => initTanStackRuntime(), { once: true })
} else {
  initTanStackRuntime()
}
