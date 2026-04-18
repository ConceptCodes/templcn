package ui

import (
	"context"
	"io"

	"github.com/a-h/templ"
)

func attrsFromDOMProps(props DOMProps, slot string, className string) templ.Attributes {
	attrs := cloneAttributes(props.Attrs)
	if props.ID != "" {
		attrs["id"] = props.ID
	}
	if slot != "" {
		attrs["data-slot"] = slot
	}

	combined := className
	if props.Class != "" {
		combined = cn(combined, props.Class)
	}
	if existing, ok := attrs["class"]; ok {
		combined = cn(combined, templ.Classes(existing).String())
	}
	if combined != "" {
		attrs["class"] = combined
	}
	return attrs
}

func renderElement(ctx context.Context, w io.Writer, tag string, attrs templ.Attributes, children templ.Component) error {
	buf, isBuffer := withBuffer(w)
	if !isBuffer {
		defer templ.ReleaseBuffer(buf)
	}

	ctx = templ.InitializeContext(ctx)
	if _, err := buf.WriteString("<" + tag); err != nil {
		return err
	}
	if err := templ.RenderAttributes(ctx, buf, attrs); err != nil {
		return err
	}
	if _, err := buf.WriteString(">"); err != nil {
		return err
	}
	if err := renderChildren(ctx, buf, children); err != nil {
		return err
	}
	if _, err := buf.WriteString("</" + tag + ">"); err != nil {
		return err
	}
	return finishBuffer(buf, isBuffer, w)
}

func renderVoidElement(ctx context.Context, w io.Writer, tag string, attrs templ.Attributes) error {
	buf, isBuffer := withBuffer(w)
	if !isBuffer {
		defer templ.ReleaseBuffer(buf)
	}

	ctx = templ.InitializeContext(ctx)
	if _, err := buf.WriteString("<" + tag); err != nil {
		return err
	}
	if err := templ.RenderAttributes(ctx, buf, attrs); err != nil {
		return err
	}
	if _, err := buf.WriteString(" />"); err != nil {
		return err
	}
	return finishBuffer(buf, isBuffer, w)
}

func renderTextElement(ctx context.Context, w io.Writer, tag string, attrs templ.Attributes, text string) error {
	return renderElement(ctx, w, tag, attrs, templ.ComponentFunc(func(_ context.Context, ww io.Writer) error {
		_, err := io.WriteString(ww, templ.EscapeString(text))
		return err
	}))
}

func blockComponent(tag, slot string, props DOMProps, children templ.Component) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, tag, attrsFromDOMProps(props, slot, ""), children)
	})
}

func leafComponent(tag, slot string, props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderVoidElement(ctx, w, tag, attrsFromDOMProps(props, slot, ""))
	})
}

func textComponent(tag, slot string, props DOMProps, text string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderTextElement(ctx, w, tag, attrsFromDOMProps(props, slot, ""), text)
	})
}
