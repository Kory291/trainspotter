package partials

import (
	"strings"
	"trainspotter/components"

	"github.com/maddalax/htmgo/framework/h"
)

// KnownTrains holds placeholder trains – replace with real data once the backend is connected.
var knownTrains = []components.Train{
	{TZNumber: "9001", Baureihe: "ICE 3", Name: "München"},
	{TZNumber: "9002", Baureihe: "ICE 3", Name: "Hamburg"},
	{TZNumber: "9003", Baureihe: "ICE 3", Name: "Berlin"},
	{TZNumber: "9023", Baureihe: "ICE 3neo", Name: "Frankfurt"},
	{TZNumber: "9047", Baureihe: "ICE 3neo", Name: "Köln"},
	{TZNumber: "9071", Baureihe: "ICE 3neo", Name: "Stuttgart"},
	{TZNumber: "9105", Baureihe: "ICE 4", Name: "Nürnberg"},
	{TZNumber: "9012", Baureihe: "ICE 4", Name: "Dresden"},
	{TZNumber: "9034", Baureihe: "ICE 1", Name: "Hannover"},
	{TZNumber: "9088", Baureihe: "ICE 1", Name: "Mannheim"},
}

func TrainListPartial(ctx *h.RequestContext) *h.Partial {
	query := strings.TrimSpace(ctx.FormValue("tz_filter"))

	var filtered []components.Train
	for _, t := range knownTrains {
		if query == "" || strings.Contains(t.TZNumber, query) {
			filtered = append(filtered, t)
		}
	}

	return h.NewPartial(TrainList(filtered))
}

// GetKnownTrains returns the current list of trains.
func GetKnownTrains() []components.Train {
	return knownTrains
}

func TrainList(trains []components.Train) *h.Element {
	if len(trains) == 0 {
		return h.Div(
			h.Id("train-list"),
			h.Class("text-prussian text-sm"),
			h.Text("No trains found."),
		)
	}
	return h.Div(
		h.Id("train-list"),
		h.Class("flex flex-col gap-3"),
		h.List(trains, func(t components.Train, _ int) *h.Element {
			return components.TrainCard(t)
		}),
	)
}
