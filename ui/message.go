package ui

import (
	"context"
	"github.com/a-h/templ"
	"io"
)

type MessageProps struct {
	DOMProps
	Align string
}

func MessageGroup(props DOMProps) templ.Component {
	return messageBlock("message-group", "flex min-w-0 flex-col gap-2", "div", props)
}
func Message(props MessageProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		align := props.Align
		if align == "" {
			align = "start"
		}
		attrs := attrsFromDOMProps(props.DOMProps, "message", "group/message relative flex w-full min-w-0 gap-2 text-sm data-[align=end]:flex-row-reverse")
		attrs["data-align"] = align
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}
func MessageAvatar(props DOMProps) templ.Component {
	return messageBlock("message-avatar", "flex w-fit min-w-8 shrink-0 items-center justify-center self-end overflow-hidden rounded-full bg-muted", "div", props)
}
func MessageContent(props DOMProps) templ.Component {
	return messageBlock("message-content", "flex w-full min-w-0 flex-col gap-2.5 wrap-break-word", "div", props)
}
func MessageHeader(props DOMProps) templ.Component {
	return messageBlock("message-header", "flex max-w-full min-w-0 items-center px-3 text-xs font-medium text-muted-foreground", "div", props)
}
func MessageFooter(props DOMProps) templ.Component {
	return messageBlock("message-footer", "flex max-w-full min-w-0 items-center px-3 text-xs font-medium text-muted-foreground", "div", props)
}
func messageBlock(slot, className, tag string, props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, tag, attrsFromDOMProps(props, slot, className), templ.GetChildren(ctx))
	})
}
