package ui

import (
	"context"
	"io"
	"strconv"

	"github.com/a-h/templ"
)

type DataTableColumn struct {
	Key      string
	Header   string
	Sortable bool
	Width    string
}

type DataTableCell struct {
	Value string
	Class string
}

type DataTableRow struct {
	ID    string
	Cells []DataTableCell
}

type DataTableProps struct {
	DOMProps
	Columns  []DataTableColumn
	Rows     []DataTableRow
	Page     int
	PageSize int
	Sort     string
	Filters  map[string]string
	Empty    string
}

func DataTable(props DataTableProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props.DOMProps, "data-table", "grid gap-4")
		if props.Page > 0 {
			attrs["data-page"] = props.Page
		}
		if props.PageSize > 0 {
			attrs["data-page-size"] = props.PageSize
		}
		if props.Sort != "" {
			attrs["data-sort"] = props.Sort
		}
		if len(props.Filters) > 0 {
			attrs["data-filters"] = props.Filters
		}
		if props.Empty != "" {
			attrs["data-empty"] = props.Empty
		}

		return renderElement(ctx, w, "section", attrs, templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
			if err := renderElement(ctx, w, "div", attrsFromDOMProps(DOMProps{}, "data-table-toolbar", "flex flex-col gap-3 sm:flex-row sm:items-center sm:justify-between"), templ.GetChildren(ctx)); err != nil {
				return err
			}

			if _, err := io.WriteString(w, "<div class=\"rounded-md border\"><table"); err != nil {
				return err
			}
			if err := templ.RenderAttributes(ctx, w, attrsFromDOMProps(props.DOMProps, "data-table-table", "w-full caption-bottom text-sm")); err != nil {
				return err
			}
			if err := renderTableContents(ctx, w, props); err != nil {
				return err
			}
			return nil
		}))
	})
}

func renderTableContents(ctx context.Context, w io.Writer, props DataTableProps) error {
	if _, err := io.WriteString(w, ">"); err != nil {
		return err
	}

	if _, err := io.WriteString(w, "<thead class=\"[&_tr]:border-b\"><tr>"); err != nil {
		return err
	}
	for _, column := range props.Columns {
		thAttrs := templ.Attributes{"class": "h-12 px-4 text-left align-middle font-medium text-muted-foreground"}
		if column.Width != "" {
			thAttrs["style"] = "width: " + column.Width + ";"
		}
		if column.Key != "" {
			thAttrs["data-column-key"] = column.Key
		}
		if column.Sortable {
			thAttrs["data-sortable"] = "true"
		}
		if err := renderTextElement(ctx, w, "th", thAttrs, column.Header); err != nil {
			return err
		}
	}
	if _, err := io.WriteString(w, "</tr></thead><tbody class=\"[&_tr:last-child]:border-0\">"); err != nil {
		return err
	}

	if len(props.Rows) == 0 {
		if _, err := io.WriteString(w, `<tr><td class="p-6 text-center text-sm text-muted-foreground" colspan="`+strconv.Itoa(max(len(props.Columns), 1))+`">`+templ.EscapeString(props.Empty)+`</td></tr>`); err != nil {
			return err
		}
	} else {
		for _, row := range props.Rows {
			rowAttrs := templ.Attributes{"class": "border-b transition-colors hover:bg-muted/50"}
			if row.ID != "" {
				rowAttrs["data-row-id"] = row.ID
			}
			if _, err := io.WriteString(w, "<tr"); err != nil {
				return err
			}
			if err := templ.RenderAttributes(ctx, w, rowAttrs); err != nil {
				return err
			}
			if _, err := io.WriteString(w, ">"); err != nil {
				return err
			}
			for _, cell := range row.Cells {
				tdAttrs := templ.Attributes{"class": "p-4 align-middle"}
				if cell.Class != "" {
					tdAttrs["class"] = cn(tdAttrs["class"].(string), cell.Class)
				}
				if err := renderTextElement(ctx, w, "td", tdAttrs, cell.Value); err != nil {
					return err
				}
			}
			if _, err := io.WriteString(w, "</tr>"); err != nil {
				return err
			}
		}
	}

	if _, err := io.WriteString(w, "</tbody></table></div>"); err != nil {
		return err
	}

	if len(props.Rows) > 0 && props.PageSize > 0 {
		return DataTablePagination(DataTablePaginationProps{
			Page:     props.Page,
			PageSize: props.PageSize,
			Total:    len(props.Rows),
		}).Render(ctx, w)
	}

	return nil
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func DataTableToolbar(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "data-table-toolbar", "flex flex-col gap-3 sm:flex-row sm:items-center sm:justify-between"), templ.GetChildren(ctx))
	})
}

func DataTableFilters(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "data-table-filters", "flex flex-wrap gap-2"), templ.GetChildren(ctx))
	})
}

type DataTablePaginationProps struct {
	DOMProps
	Page     int
	PageSize int
	Total    int
}

func DataTablePagination(props DataTablePaginationProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props.DOMProps, "data-table-pagination", "flex items-center justify-between gap-2")
		if props.Page > 0 {
			attrs["data-page"] = props.Page
		}
		if props.PageSize > 0 {
			attrs["data-page-size"] = props.PageSize
		}
		if props.Total > 0 {
			attrs["data-total"] = props.Total
		}

		children := Pagination(DOMProps{})
		return renderElement(ctx, w, "div", attrs, children)
	})
}

type DataTableColumnHeaderProps struct {
	DOMProps
	Title    string
	Sortable bool
	Sorted   string
}

func DataTableColumnHeader(props DataTableColumnHeaderProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		className := "inline-flex items-center gap-2"
		if props.Sortable {
			className = cn(className, "cursor-pointer select-none")
		}
		attrs := attrsFromDOMProps(props.DOMProps, "data-table-column-header", className)
		if props.Sorted != "" {
			attrs["data-sorted"] = props.Sorted
		}
		return renderTextElement(ctx, w, "button", attrs, props.Title)
	})
}
