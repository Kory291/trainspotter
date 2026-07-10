package pages

import (
	"trainspotter-frontend/components"
	"trainspotter-frontend/partials"

	"github.com/maddalax/htmgo/framework/h"
)

func TrainsPage(ctx *h.RequestContext) *h.Page {
	return RootPage(
		components.Nav(ctx),
		h.Div(
			h.Class("p-6 max-w-2xl mx-auto flex flex-col gap-8"),
			components.AddTrainForm(),
			h.Div(
				h.Class("flex flex-col gap-4"),
				h.H2(
					h.Class("text-xl font-bold text-prussian"),
					h.Text("My Trains"),
				),
				h.Input(
					"text",
					h.Name("tz_filter"),
					h.Id("tz_filter"),
					h.Attribute("placeholder", "Filter by TZ number..."),
					h.Class("border border-saffron rounded px-3 py-2 text-prussian bg-white focus:outline-none focus:ring-2 focus:ring-ochre"),
					h.HxTarget("#train-list"),
					h.PostPartial(partials.TrainListPartial, "input changed delay:200ms"),
				),
				partials.TrainList(partials.GetKnownTrains()),
			),
		),
	)
}
