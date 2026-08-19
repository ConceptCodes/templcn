# v1 interaction matrix

This is the browser-coverage contract for v1. Every interactive item has at least one rendered HTML assertion and one browser assertion. A behavior is marked `n/a` only when the primitive does not expose it.

| Family | Keyboard | Focus / dismissal | State / forms | Responsive / theme |
|---|---|---|---|---|
| Accordion / collapsible | covered | n/a / covered | controlled + default | covered |
| Calendar / date-picker | covered | outside dismissal | hidden form value | dark mode |
| Checkbox / radio / switch / toggle | covered | focus-visible | checked + disabled + invalid | dark mode |
| Combobox / command | covered | escape + outside | selected + hidden value | dark mode |
| Dialog / alert-dialog / sheet / drawer | covered | trap + restore + escape | controlled + default | mobile + dark mode |
| Dropdown-menu / context-menu / menubar | covered | escape + outside | open/closed state | dark mode |
| Hover-card / popover / tooltip | covered | escape + outside | open/closed state | responsive |
| Select | covered | escape + outside | selected + disabled | dark mode |
| Sidebar | shortcut + trigger | restore after collapse | persistence + state | mobile + dark mode |
| Slider / resizable | arrows | focus-visible | multiple values + orientation | responsive |
| Tabs / navigation-menu | arrows | focus-visible | active value | responsive + dark mode |
| Input OTP / form controls | arrows | focus-visible | submission + invalid | responsive |

The source of truth for the concrete spec files is `docs-site/tests`. The component scope matrix records whether each row has evidence and whether it is intentionally adapted for Go-native rendering.
