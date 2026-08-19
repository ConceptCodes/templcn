import {
  createTable,
  getCoreRowModel,
  getSortedRowModel,
  type SortingState,
} from "@tanstack/table-core"

type TableRow = Record<string, string>

function readRows(table: HTMLTableElement, keys: string[]): TableRow[] {
  return Array.from(table.tBodies[0]?.rows || []).map((row) => {
    const cells = Array.from(row.cells)
    return Object.fromEntries(keys.map((key, index) => [key, cells[index]?.textContent?.trim() || ""]))
  })
}

export function initTanStackTables(root: ParentNode = document) {
  root.querySelectorAll<HTMLElement>('[data-engine="tanstack"]:not([data-tanstack-mounted])').forEach((container) => {
    const table = container.querySelector<HTMLTableElement>("table")
    if (!table || !table.tHead || !table.tBodies[0]) return

    const headers = Array.from(table.tHead.rows[0]?.cells || [])
    const keys = headers.map((header, index) => header.getAttribute("data-column-key") || `column-${index}`)
    const sortable = new Set(headers.filter((header) => header.hasAttribute("data-sortable")).map((header) => header.getAttribute("data-column-key")))
    const initialData = readRows(table, keys)
    let state: { sorting: SortingState } = { sorting: [] }

    const instance = createTable({
      data: initialData,
      columns: keys.map((key, index) => ({
        id: key,
        accessorKey: key,
        header: headers[index]?.textContent?.trim() || key,
        enableSorting: sortable.has(key),
      })),
      state,
      onStateChange: (updater) => {
        state = typeof updater === "function" ? updater(state) : updater
        instance.setOptions((options) => ({ ...options, state }))
        render()
      },
      getCoreRowModel: getCoreRowModel(),
      getSortedRowModel: getSortedRowModel(),
      renderFallbackValue: "",
    })

    function render() {
      const body = table!.tBodies[0]
      body.innerHTML = ""
      for (const row of instance.getRowModel().rows) {
        const tr = document.createElement("tr")
        tr.className = "border-b transition-colors hover:bg-muted/50"
        for (const key of keys) {
          const td = document.createElement("td")
          td.className = "p-4 align-middle"
          td.textContent = String(row.getValue(key) ?? "")
          tr.appendChild(td)
        }
        body.appendChild(tr)
      }
      headers.forEach((header) => {
        const key = header.getAttribute("data-column-key")
        const column = key ? instance.getColumn(key) : undefined
        if (column && sortable.has(key)) header.setAttribute("aria-sort", column.getIsSorted() || "none")
      })
    }

    headers.forEach((header) => {
      const key = header.getAttribute("data-column-key")
      const column = key ? instance.getColumn(key) : undefined
      if (!column || !sortable.has(key)) return
      header.tabIndex = 0
      header.setAttribute("role", "columnheader")
      const sort = () => column.toggleSorting()
      header.addEventListener("click", sort)
      header.addEventListener("keydown", (event) => {
        if (event.key === "Enter" || event.key === " ") {
          event.preventDefault()
          sort()
        }
      })
    })

    container.dataset.tanstackMounted = "true"
    render()
  })
}

