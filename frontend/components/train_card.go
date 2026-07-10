package components

import (
	"github.com/maddalax/htmgo/framework/h"
)

type Train struct {
	TZNumber string
	Baureihe string
	Name     string
}

func TrainCard(t Train) *h.Element {
	return h.Div(
		h.Class("bg-white border border-saffron rounded-lg shadow-sm p-4 flex flex-col gap-1"),
		h.Div(
			h.Class("flex items-center justify-between"),
			h.Span(
				h.Class("text-lg font-bold text-prussian"),
				h.TextF("TZ %s", t.TZNumber),
			),
			h.Span(
				h.Class("text-sm font-medium text-ochre"),
				h.Text(t.Baureihe),
			),
		),
		h.Span(
			h.Class("text-sm text-prussian"),
			h.Text(t.Name),
		),
	)
}
