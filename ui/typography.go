package ui

import (
	"context"
	"io"

	"github.com/a-h/templ"
)

func typographyComponent(tag, slot, className string, props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, tag, attrsFromDOMProps(props, slot, className), templ.GetChildren(ctx))
	})
}

func H1(props DOMProps) templ.Component {
	return typographyComponent("h1", "h1", "scroll-m-20 text-4xl font-bold tracking-tight lg:text-5xl", props)
}
func H2(props DOMProps) templ.Component {
	return typographyComponent("h2", "h2", "scroll-m-20 border-b pb-2 text-3xl font-semibold tracking-tight first:mt-0", props)
}
func H3(props DOMProps) templ.Component {
	return typographyComponent("h3", "h3", "scroll-m-20 text-2xl font-semibold tracking-tight", props)
}
func H4(props DOMProps) templ.Component {
	return typographyComponent("h4", "h4", "scroll-m-20 text-xl font-semibold tracking-tight", props)
}
func P(props DOMProps) templ.Component {
	return typographyComponent("p", "p", "leading-7 [&:not(:first-child)]:mt-6", props)
}
func Blockquote(props DOMProps) templ.Component {
	return typographyComponent("blockquote", "blockquote", "mt-6 border-l-2 border-border pl-6 italic", props)
}
func InlineCode(props DOMProps) templ.Component {
	return typographyComponent("code", "inline-code", "relative rounded bg-muted px-[0.3rem] py-[0.2rem] font-mono text-sm font-semibold", props)
}
func Lead(props DOMProps) templ.Component {
	return typographyComponent("p", "lead", "text-xl text-muted-foreground", props)
}
func Large(props DOMProps) templ.Component {
	return typographyComponent("div", "large", "text-lg font-semibold", props)
}
func Small(props DOMProps) templ.Component {
	return typographyComponent("small", "small", "text-sm font-medium leading-none", props)
}
func Muted(props DOMProps) templ.Component {
	return typographyComponent("p", "muted", "text-sm text-muted-foreground", props)
}
func List(props DOMProps) templ.Component {
	return typographyComponent("ul", "list", "my-6 ml-6 list-disc [&>li]:mt-2", props)
}
func TableProse(props DOMProps) templ.Component {
	return typographyComponent("div", "table-prose", "my-6 w-full overflow-y-auto", props)
}
