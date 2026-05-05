import { contentFor, px, triggerFor } from "./core.js";

const viewportPadding = 8;

export function positionFloating(root) {
  const trigger = triggerFor(root);
  const content = contentFor(root);
  if (!root || !trigger || !content) return;

  const preferredSide = root.getAttribute("data-side") || content.getAttribute("data-side") || "bottom";
  const align = root.getAttribute("data-align") || content.getAttribute("data-align") || "center";
  const sideOffset = px(root.getAttribute("data-side-offset") || content.getAttribute("data-side-offset"), 4);
  const alignOffset = px(root.getAttribute("data-align-offset") || content.getAttribute("data-align-offset"), 0);

  content.style.position = "fixed";
  content.style.margin = "0";
  content.style.zIndex ||= "50";
  content.style.left = "0px";
  content.style.top = "0px";

  const anchor = trigger.getBoundingClientRect();
  const floating = content.getBoundingClientRect();
  let side = preferredSide;

  if (side === "bottom" && anchor.bottom + sideOffset + floating.height > window.innerHeight - viewportPadding) side = "top";
  if (side === "top" && anchor.top - sideOffset - floating.height < viewportPadding) side = "bottom";
  if (side === "right" && anchor.right + sideOffset + floating.width > window.innerWidth - viewportPadding) side = "left";
  if (side === "left" && anchor.left - sideOffset - floating.width < viewportPadding) side = "right";

  let top = 0;
  let left = 0;

  if (side === "top" || side === "bottom") {
    top = side === "bottom" ? anchor.bottom + sideOffset : anchor.top - floating.height - sideOffset;
    if (align === "start") left = anchor.left + alignOffset;
    else if (align === "end") left = anchor.right - floating.width + alignOffset;
    else left = anchor.left + (anchor.width - floating.width) / 2 + alignOffset;
  } else {
    left = side === "right" ? anchor.right + sideOffset : anchor.left - floating.width - sideOffset;
    if (align === "start") top = anchor.top + alignOffset;
    else if (align === "end") top = anchor.bottom - floating.height + alignOffset;
    else top = anchor.top + (anchor.height - floating.height) / 2 + alignOffset;
  }

  left = Math.max(viewportPadding, Math.min(left, window.innerWidth - floating.width - viewportPadding));
  top = Math.max(viewportPadding, Math.min(top, window.innerHeight - floating.height - viewportPadding));

  content.style.left = `${Math.round(left)}px`;
  content.style.top = `${Math.round(top)}px`;
  content.setAttribute("data-side", side);
  content.setAttribute("data-align", align);

  const availableWidth = Math.max(0, window.innerWidth - viewportPadding - left);
  const availableHeight = Math.max(0, window.innerHeight - viewportPadding - top);
  content.style.setProperty("--radix-popover-trigger-width", `${anchor.width}px`);
  content.style.setProperty("--radix-popover-trigger-height", `${anchor.height}px`);
  content.style.setProperty("--radix-dropdown-menu-trigger-width", `${anchor.width}px`);
  content.style.setProperty("--radix-dropdown-menu-trigger-height", `${anchor.height}px`);
  content.style.setProperty("--anchor-width", `${anchor.width}px`);
  content.style.setProperty("--anchor-height", `${anchor.height}px`);
  content.style.setProperty("--available-width", `${availableWidth}px`);
  content.style.setProperty("--available-height", `${availableHeight}px`);
  content.style.setProperty("--transform-origin", transformOrigin(side, align));
}

export function positionAllOpenFloating() {
  document.querySelectorAll('[data-open="true"]').forEach(positionFloating);
}

function transformOrigin(side, align) {
  const cross = align === "start" ? "left" : align === "end" ? "right" : "center";
  if (side === "top") return `${cross} bottom`;
  if (side === "bottom") return `${cross} top`;
  if (side === "left") return `right ${cross}`;
  return `left ${cross}`;
}

window.addEventListener("resize", positionAllOpenFloating);
window.addEventListener("scroll", positionAllOpenFloating, true);

