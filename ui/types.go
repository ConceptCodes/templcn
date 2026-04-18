package ui

import "github.com/a-h/templ"

type DOMProps struct {
	ID      string
	Class   string
	Element string
	Attrs   templ.Attributes
}

type ButtonVariant string

const (
	ButtonVariantDefault     ButtonVariant = "default"
	ButtonVariantDestructive ButtonVariant = "destructive"
	ButtonVariantOutline     ButtonVariant = "outline"
	ButtonVariantSecondary   ButtonVariant = "secondary"
	ButtonVariantGhost       ButtonVariant = "ghost"
	ButtonVariantLink        ButtonVariant = "link"
)

type ButtonSize string

const (
	ButtonSizeDefault ButtonSize = "default"
	ButtonSizeXS      ButtonSize = "xs"
	ButtonSizeSM      ButtonSize = "sm"
	ButtonSizeLG      ButtonSize = "lg"
	ButtonSizeIcon    ButtonSize = "icon"
	ButtonSizeIconXS  ButtonSize = "icon-xs"
	ButtonSizeIconSM  ButtonSize = "icon-sm"
	ButtonSizeIconLG  ButtonSize = "icon-lg"
)

type ButtonProps struct {
	DOMProps
	Label    string
	Variant  ButtonVariant
	Size     ButtonSize
	Type     string
	Disabled bool
	Href     string
}
