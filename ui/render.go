package ui

import (
	"bytes"
	"context"
	"io"

	"github.com/a-h/templ"
)

func withBuffer(w io.Writer) (*bytes.Buffer, bool) {
	buf, ok := w.(*bytes.Buffer)
	if ok {
		return buf, true
	}

	buf = templ.GetBuffer()
	return buf, false
}

func finishBuffer(buf *bytes.Buffer, isBuffer bool, w io.Writer) error {
	if !isBuffer {
		_, err := buf.WriteTo(w)
		return err
	}
	return nil
}

func renderChildren(ctx context.Context, w io.Writer, children templ.Component) error {
	if children == nil {
		return nil
	}
	return children.Render(ctx, w)
}

func childrenFromContext(ctx context.Context) (context.Context, templ.Component) {
	children := templ.GetChildren(ctx)
	if children == nil {
		children = templ.NopComponent
	}
	return templ.ClearChildren(ctx), children
}
