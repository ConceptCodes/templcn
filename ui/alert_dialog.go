package ui

import (
	"github.com/a-h/templ"
)

type AlertDialogProps struct {
	DialogProps
}

func AlertDialog(props AlertDialogProps) templ.Component    { return Dialog(props.DialogProps) }
func AlertDialogTrigger(props DOMProps) templ.Component     { return DialogTrigger(props) }
func AlertDialogPortal(props DOMProps) templ.Component      { return DialogPortal(props) }
func AlertDialogOverlay(props DOMProps) templ.Component     { return DialogOverlay(props) }
func AlertDialogContent(props DOMProps) templ.Component     { return DialogContent(props) }
func AlertDialogHeader(props DOMProps) templ.Component      { return DialogHeader(props) }
func AlertDialogFooter(props DOMProps) templ.Component      { return DialogFooter(props) }
func AlertDialogTitle(props DOMProps) templ.Component       { return DialogTitle(props) }
func AlertDialogDescription(props DOMProps) templ.Component { return DialogDescription(props) }
func AlertDialogAction(props DOMProps) templ.Component      { return DialogClose(props) }
func AlertDialogCancel(props DOMProps) templ.Component      { return DialogClose(props) }
