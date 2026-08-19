import { areaY, barY, defineChart, lineY } from "@tanstack/charts"
import { mountChart } from "@tanstack/charts/dom"
import { pie, polar, radialArc, radialText } from "@tanstack/charts/polar"
import { initTanStackTables } from "./tanstack-table"

type Series = { key: string; label?: string; color?: string }
type ChartConfig = {
  type?: "area" | "bar" | "line" | "pie" | "radial"
  data?: Record<string, unknown>[]
  series?: Series[]
  orientation?: "vertical" | "horizontal"
  ariaLabel?: string
  height?: number
  innerRadius?: number
  centerLabel?: string
  centerCaption?: string
  value?: number
  max?: number
  segments?: { label: string; value: number }[]
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

function renderChart(element: HTMLElement) {
  const config = readConfig(element)
  if (!config?.data?.length || !config.series?.length) return

  const series = config.series
  const type = config.type || element.dataset.chart || "line"

  if (type === "pie" || type === "radial") {
    return renderPolarChart(element, config, type)
  }

  const marks = series.map((entry) => {
    const markOptions = {
      x: "label",
      y: entry.key,
      ...(entry.color ? { stroke: entry.color, fill: entry.color } : {}),
    } as const
    return type === "bar"
      ? barY(config.data!, markOptions)
      : type === "area"
        ? areaY(config.data!, markOptions)
        : lineY(config.data!, markOptions)
  })

  const definition = defineChart({
    marks,
  })

  const host = mountChart(element, {
    definition,
    height: config.height || 288,
    ariaLabel: config.ariaLabel || element.getAttribute("aria-label") || `${type} chart`,
  })

  return host
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
    }),
  ]

  if (config.centerLabel) {
    marks.push(radialText([{ label: "center", text: config.centerLabel }], {
      angle: 0,
      radius: 0,
      text: "text",
      key: "label",
      fill: "currentColor",
      fontSize: 22,
      fontWeight: 700,
    }))
  }

  const definition = defineChart({
    marks: [polar({ inset: 8, marks })],
    color: {
      domain: labels,
      range: ["#0ea5e9", "#6366f1", "#a855f7", "#ec4899", "#f97316", "#94a3b8"],
    },
  })

  return mountChart(element, {
    definition,
    height: config.height || 288,
    ariaLabel: config.ariaLabel || element.getAttribute("aria-label") || `${type} chart`,
  })
}

export function initTanStackCharts(root: ParentNode = document) {
  root.querySelectorAll<HTMLElement>('[data-engine="tanstack"]:not([data-tanstack-mounted])').forEach((element) => {
    const host = renderChart(element)
    if (!host) return
    element.dataset.tanstackMounted = "true"
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
