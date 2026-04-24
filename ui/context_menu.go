package ui

import (
	"github.com/a-h/templ"
)

type ContextMenuProps struct{ DOMProps }

func ContextMenu(props ContextMenuProps) templ.Component {
	return DropdownMenu(DropdownMenuProps{DOMProps: props.DOMProps})
}
func ContextMenuTrigger(props DOMProps) templ.Component           { return DialogTrigger(props) }
func ContextMenuPortal(props DOMProps) templ.Component            { return DialogPortal(props) }
func ContextMenuContent(props DOMProps) templ.Component           { return DropdownMenuContent(props) }
func ContextMenuItem(props DropdownMenuItemProps) templ.Component { return DropdownMenuItem(props) }
func ContextMenuCheckboxItem(props DropdownMenuItemProps) templ.Component {
	return DropdownMenuCheckboxItem(props)
}
func ContextMenuRadioItem(props DropdownMenuItemProps) templ.Component {
	return DropdownMenuRadioItem(props)
}
func ContextMenuRadioGroup(props DropdownMenuRadioGroupProps) templ.Component {
	return DropdownMenuRadioGroup(props)
}
func ContextMenuLabel(props DOMProps) templ.Component      { return DropdownMenuLabel(props) }
func ContextMenuSeparator(props DOMProps) templ.Component  { return DropdownMenuSeparator(props) }
func ContextMenuShortcut(props DOMProps) templ.Component   { return DropdownMenuShortcut(props) }
func ContextMenuGroup(props DOMProps) templ.Component      { return DropdownMenuGroup(props) }
func ContextMenuSub(props DOMProps) templ.Component        { return DropdownMenuSub(props) }
func ContextMenuSubContent(props DOMProps) templ.Component { return DropdownMenuSubContent(props) }
func ContextMenuSubTrigger(props DropdownMenuItemProps) templ.Component {
	return DropdownMenuSubTrigger(props)
}
