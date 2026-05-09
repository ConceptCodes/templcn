package ui

import "context"

type floatingRenderStateKey struct{}

type floatingRenderState struct {
	open bool
}

func floatingStateFromContext(ctx context.Context) string {
	if state, ok := ctx.Value(floatingRenderStateKey{}).(floatingRenderState); ok && state.open {
		return "open"
	}
	return "closed"
}

func floatingOpenFromContext(ctx context.Context) bool {
	if state, ok := ctx.Value(floatingRenderStateKey{}).(floatingRenderState); ok {
		return state.open
	}
	return false
}
