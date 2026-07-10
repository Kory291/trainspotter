package components

import (
	"github.com/maddalax/htmgo/framework/h"
)

func NavItem(label string, href string, active bool) *h.Element {
	class := "pb-1"
	if active {
		class = "border-b-2 border-vermilion pb-1"
	}
	return h.Li(
		h.Class(class),
		h.A(
			h.Href(href),
			h.Text(label),
		),
	)
}

func Nav(ctx *h.RequestContext) *h.Element {
	path := ctx.Request.URL.Path
	return h.Nav(
		h.Class("flex justify-center w-full bg-ochre"),
		h.Ol(
			h.Id("nav"),
			h.Class("flex gap-6 px-4 py-2"),
			NavItem("Sights", "/", path == "/"),
			NavItem("Trains", "/trains", path == "/trains"),
		),
	)
}
