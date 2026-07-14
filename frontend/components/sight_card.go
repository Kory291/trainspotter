package components

import (
	"time"

	"github.com/maddalax/htmgo/framework/h"
)

type Sight struct {
	Train    int `json:"train"`
	Place string `json:"place"`
	Date     time.Time `json:"date"`
}

func SightCard(s Sight) *h.Element {
	return h.Div(
		h.Class("bg-white border border-saffron rounded-lg shadow-sm p-4 flex flex-col gap-2"),
		h.Div(
			h.Class("flex items-center justify-between"),
			h.Span(
				h.Class("text-lg font-bold text-prussian"),
				h.TextF("TZ %d", s.Train),
			),
			h.Span(
				h.Class("text-sm text-gray-500"),
				h.Text(s.Date.Format("02.01.2006")),
			),
		),
		h.Div(
			h.Class("flex items-center gap-1 text-sm text-prussian"),
			h.Span(
				h.Class("font-medium"),
				h.Text(s.Place),
			),
		),
	)
}
