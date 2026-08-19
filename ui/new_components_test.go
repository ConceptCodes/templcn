package ui

import (
	"strings"
	"testing"
)

func TestNewShadcnComponentsRenderExpectedSlots(t *testing.T) {
	tests := []struct {
		name  string
		html  string
		slots []string
	}{
		{name: "attachment", html: mustRender(t, Attachment(AttachmentProps{})), slots: []string{`data-slot="attachment"`}},
		{name: "bubble", html: mustRender(t, Bubble(BubbleProps{})), slots: []string{`data-slot="bubble"`}},
		{name: "marker", html: mustRender(t, Marker(MarkerProps{})), slots: []string{`data-slot="marker"`}},
		{name: "message", html: mustRender(t, Message(MessageProps{})), slots: []string{`data-slot="message"`}},
		{name: "message scroller", html: mustRender(t, MessageScroller(MessageScrollerProps{})), slots: []string{`data-slot="message-scroller"`}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, slot := range tt.slots {
				if !strings.Contains(tt.html, slot) {
					t.Fatalf("expected %s in %q", slot, tt.html)
				}
			}
		})
	}
}
