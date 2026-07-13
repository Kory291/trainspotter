package components

import (
	"github.com/maddalax/htmgo/framework/h"
)

func FormField(label, name, placeholder string) *h.Element {
	return h.Div(
		h.Class("flex flex-col gap-1"),
		h.Label(
			h.For(name),
			h.Class("text-sm font-medium text-prussian"),
			h.Text(label),
		),
		h.Input(
			"text",
			h.Name(name),
			h.Id(name),
			h.Attribute("placeholder", placeholder),
			h.Class("border border-saffron rounded px-3 py-2 text-prussian bg-white focus:outline-none focus:ring-2 focus:ring-ochre"),
		),
	)
}
