package ui

import (
	"github.com/a-h/templ"
)

type MenubarProps struct{ DOMProps }

func Menubar(props MenubarProps) templ.Component {
	return DropdownMenu(DropdownMenuProps{DOMProps: props.DOMProps})
}
func MenubarMenu(props DOMProps) templ.Component              { return DropdownMenuGroup(props) }
func MenubarTrigger(props DOMProps) templ.Component           { return DialogTrigger(props) }
func MenubarPortal(props DOMProps) templ.Component            { return DialogPortal(props) }
func MenubarContent(props DOMProps) templ.Component           { return DropdownMenuContent(props) }
func MenubarGroup(props DOMProps) templ.Component             { return DropdownMenuGroup(props) }
func MenubarItem(props DropdownMenuItemProps) templ.Component { return DropdownMenuItem(props) }
func MenubarCheckboxItem(props DropdownMenuItemProps) templ.Component {
	return DropdownMenuCheckboxItem(props)
}
func MenubarRadioGroup(props DropdownMenuRadioGroupProps) templ.Component {
	return DropdownMenuRadioGroup(props)
}
func MenubarRadioItem(props DropdownMenuItemProps) templ.Component {
	return DropdownMenuRadioItem(props)
}
func MenubarLabel(props DOMProps) templ.Component     { return DropdownMenuLabel(props) }
func MenubarSeparator(props DOMProps) templ.Component { return DropdownMenuSeparator(props) }
func MenubarShortcut(props DOMProps) templ.Component  { return DropdownMenuShortcut(props) }
func MenubarSub(props DOMProps) templ.Component       { return DropdownMenuSub(props) }
func MenubarSubTrigger(props DropdownMenuItemProps) templ.Component {
	return DropdownMenuSubTrigger(props)
}
func MenubarSubContent(props DOMProps) templ.Component { return DropdownMenuSubContent(props) }
