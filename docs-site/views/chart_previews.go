package views

import (
	"context"
	"encoding/json"
	"html"
	"io"

	"github.com/a-h/templ"
)

func rawComponent(htmlString string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		_, err := io.WriteString(w, htmlString)
		return err
	})
}

func chartFrame(chartType string, config any) templ.Component {
	payload, err := json.Marshal(config)
	if err != nil {
		payload = []byte(`{}`)
	}

	engine := ` data-engine="tanstack"`
	scripts := `<script type="module" src="/public/tanstack-runtime.js"></script>`
	srcdoc := `<!doctype html><html lang="en"><head><meta charset="utf-8"><meta name="viewport" content="width=device-width,initial-scale=1"><link rel="stylesheet" href="/public/globals.css">` + scripts + `</head><body class="bg-background text-foreground"><div class="w-full rounded-xl bg-background p-4"><div class="min-h-[288px] w-full"` + engine + ` data-chart="` + html.EscapeString(chartType) + `" data-chart-config="` + html.EscapeString(string(payload)) + `" aria-label="` + html.EscapeString(chartType+" chart") + `"></div></div></body></html>`

	return rawComponent(
		`<iframe class="h-[320px] w-full rounded-xl border border-border/70 bg-background" loading="lazy" title="` + html.EscapeString(chartType) + `" srcdoc="` + html.EscapeString(srcdoc) + `"></iframe>`,
	)
}

func areaChartBasicPreview() templ.Component {
	return chartFrame("area", map[string]any{
		"title":      "Revenue over time",
		"stacked":    false,
		"showPoints": true,
		"series": []map[string]any{
			{"key": "revenue", "label": "Revenue"},
		},
		"data": []map[string]any{
			{"label": "Jan", "revenue": 68},
			{"label": "Feb", "revenue": 84},
			{"label": "Mar", "revenue": 76},
			{"label": "Apr", "revenue": 122},
			{"label": "May", "revenue": 108},
			{"label": "Jun", "revenue": 142},
			{"label": "Jul", "revenue": 156},
		},
	})
}

func areaChartStackedPreview() templ.Component {
	return chartFrame("area", map[string]any{
		"title":   "Stacked revenue and subscriptions",
		"stacked": true,
		"series": []map[string]any{
			{"key": "desktop", "label": "Desktop"},
			{"key": "mobile", "label": "Mobile"},
		},
		"data": []map[string]any{
			{"label": "Jan", "desktop": 42, "mobile": 24},
			{"label": "Feb", "desktop": 55, "mobile": 26},
			{"label": "Mar", "desktop": 63, "mobile": 31},
			{"label": "Apr", "desktop": 74, "mobile": 39},
			{"label": "May", "desktop": 71, "mobile": 42},
			{"label": "Jun", "desktop": 88, "mobile": 49},
			{"label": "Jul", "desktop": 96, "mobile": 57},
		},
	})
}

func areaChartInteractivePreview() templ.Component {
	return chartFrame("area", map[string]any{
		"title":         "Traffic trend with focus state",
		"stacked":       false,
		"showPoints":    true,
		"showLastPoint": true,
		"series": []map[string]any{
			{"key": "traffic", "label": "Traffic"},
		},
		"data": []map[string]any{
			{"label": "Mon", "traffic": 92},
			{"label": "Tue", "traffic": 116},
			{"label": "Wed", "traffic": 112},
			{"label": "Thu", "traffic": 138},
			{"label": "Fri", "traffic": 146},
			{"label": "Sat", "traffic": 162},
			{"label": "Sun", "traffic": 154},
		},
	})
}

func barChartVerticalPreview() templ.Component {
	return chartFrame("bar", map[string]any{
		"orientation": "vertical",
		"series": []map[string]any{
			{"key": "sales", "label": "Sales"},
		},
		"data": []map[string]any{
			{"label": "Mon", "sales": 24},
			{"label": "Tue", "sales": 36},
			{"label": "Wed", "sales": 48},
			{"label": "Thu", "sales": 31},
			{"label": "Fri", "sales": 58},
			{"label": "Sat", "sales": 74},
		},
	})
}

func barChartHorizontalPreview() templ.Component {
	return chartFrame("bar", map[string]any{
		"orientation": "horizontal",
		"series": []map[string]any{
			{"key": "completion", "label": "Completion"},
		},
		"data": []map[string]any{
			{"label": "Marketing", "completion": 84},
			{"label": "Engineering", "completion": 92},
			{"label": "Sales", "completion": 66},
			{"label": "Support", "completion": 78},
			{"label": "Design", "completion": 58},
		},
	})
}

func barChartMultiplePreview() templ.Component {
	return chartFrame("bar", map[string]any{
		"orientation": "vertical",
		"series": []map[string]any{
			{"key": "desktop", "label": "Desktop"},
			{"key": "mobile", "label": "Mobile"},
		},
		"data": []map[string]any{
			{"label": "Jan", "desktop": 24, "mobile": 14},
			{"label": "Feb", "desktop": 32, "mobile": 18},
			{"label": "Mar", "desktop": 40, "mobile": 22},
			{"label": "Apr", "desktop": 38, "mobile": 26},
			{"label": "May", "desktop": 50, "mobile": 30},
			{"label": "Jun", "desktop": 56, "mobile": 34},
		},
	})
}

func lineChartLinearPreview() templ.Component {
	return chartFrame("line", map[string]any{
		"series": []map[string]any{
			{"key": "visits", "label": "Visits"},
		},
		"showPoints": true,
		"data": []map[string]any{
			{"label": "Jan", "visits": 92},
			{"label": "Feb", "visits": 114},
			{"label": "Mar", "visits": 104},
			{"label": "Apr", "visits": 136},
			{"label": "May", "visits": 122},
			{"label": "Jun", "visits": 158},
			{"label": "Jul", "visits": 174},
		},
	})
}

func lineChartStepPreview() templ.Component {
	return chartFrame("line", map[string]any{
		"series": []map[string]any{
			{"key": "sessions", "label": "Sessions"},
		},
		"step":       true,
		"showPoints": true,
		"data": []map[string]any{
			{"label": "Mon", "sessions": 8},
			{"label": "Tue", "sessions": 13},
			{"label": "Wed", "sessions": 13},
			{"label": "Thu", "sessions": 19},
			{"label": "Fri", "sessions": 19},
			{"label": "Sat", "sessions": 26},
			{"label": "Sun", "sessions": 31},
		},
	})
}

func lineChartMultiplePreview() templ.Component {
	return chartFrame("line", map[string]any{
		"series": []map[string]any{
			{"key": "desktop", "label": "Desktop"},
			{"key": "mobile", "label": "Mobile"},
		},
		"showPoints": true,
		"data": []map[string]any{
			{"label": "Jan", "desktop": 88, "mobile": 64},
			{"label": "Feb", "desktop": 96, "mobile": 70},
			{"label": "Mar", "desktop": 104, "mobile": 78},
			{"label": "Apr", "desktop": 118, "mobile": 82},
			{"label": "May", "desktop": 126, "mobile": 94},
			{"label": "Jun", "desktop": 144, "mobile": 104},
		},
	})
}

func pieChartBasicPreview() templ.Component {
	return chartFrame("pie", map[string]any{
		"innerRadius": 0.0,
		"data": []map[string]any{
			{"label": "Desktop", "value": 46},
			{"label": "Mobile", "value": 32},
			{"label": "Tablet", "value": 22},
		},
	})
}

func pieChartDonutPreview() templ.Component {
	return chartFrame("pie", map[string]any{
		"innerRadius": 0.62,
		"data": []map[string]any{
			{"label": "Desktop", "value": 54},
			{"label": "Mobile", "value": 28},
			{"label": "Tablet", "value": 18},
		},
	})
}

func pieChartLabelPreview() templ.Component {
	return chartFrame("pie", map[string]any{
		"innerRadius":   0.68,
		"centerLabel":   "72%",
		"centerCaption": "Completion",
		"data": []map[string]any{
			{"label": "Complete", "value": 72},
			{"label": "Remaining", "value": 28},
		},
	})
}

func radialChartBasicPreview() templ.Component {
	return chartFrame("radial", map[string]any{
		"value":         64,
		"max":           100,
		"centerLabel":   "64%",
		"centerCaption": "Progress",
	})
}

func radialChartLabelPreview() templ.Component {
	return chartFrame("radial", map[string]any{
		"value":         84,
		"max":           100,
		"centerLabel":   "84",
		"centerCaption": "Score",
	})
}

func radialChartShapePreview() templ.Component {
	return chartFrame("radial", map[string]any{
		"segments": []map[string]any{
			{"label": "Core", "value": 42},
			{"label": "Growth", "value": 26},
			{"label": "Ops", "value": 18},
			{"label": "R&D", "value": 14},
		},
		"centerLabel":   "100",
		"centerCaption": "Units",
	})
}
