package views

import "github.com/a-h/templ"

var blockPreviews = map[string]func() templ.Component{
	"dashboard-01": Dashboard01Preview,
	"sidebar-07":   Sidebar07Preview,
	"sidebar-03":   Sidebar03Preview,
	"login-01":     Login01Preview,
	"login-03":     Login03Preview,
	"login-04":     Login04Preview,
	"signup-01":     Signup01Preview,
	"signup-02":     Signup02Preview,
}

func BlockPreviewForSlug(slug string) (templ.Component, bool) {
	preview, ok := blockPreviews[slug]
	if !ok {
		return nil, false
	}
	return preview(), true
}
