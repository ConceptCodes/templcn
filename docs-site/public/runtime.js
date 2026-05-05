import { initDialogs } from "./runtime/dialog.js";
import { initMenus } from "./runtime/menu.js";
import { initPopovers } from "./runtime/popover.js";
import { initTabsAndDisclosures } from "./runtime/tabs.js";

function initCharts() {
  document.querySelectorAll("[data-chart-container]").forEach((chart) => {
    if (chart.getAttribute("data-config")) {
      console.warn("ChartContainer found: inject your preferred JS charting library here to render.");
    }
  });
}

function initRuntime() {
  initDialogs();
  initPopovers();
  initMenus();
  initTabsAndDisclosures();
  initCharts();
}

if (document.readyState === "loading") {
  document.addEventListener("DOMContentLoaded", initRuntime, { once: true });
} else {
  initRuntime();
}
