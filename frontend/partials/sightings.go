package partials

import (
	"io"
	"slices"
	"trainspotter-frontend/components"

	"encoding/json"
	"net/http"
	"net/url"

	"github.com/maddalax/htmgo/framework/h"
)

var knownSightings = []components.Sight{}

func GetSightings() (sightings []components.Sight) {
	resp, err := http.Get("http://localhost:8080/sightings")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(body, &knownSightings)

	slices.SortStableFunc(knownSightings, func(sightA components.Sight, sightB components.Sight) int {
		if sightA.Date.Before(sightB.Date) {
			return 1
		} else {
			return -1
		}
	})

	return knownSightings
}

func SightList(sightings []components.Sight) *h.Element {
	if len(sightings) == 0 {
		return h.Div(
			h.Id("sightings-list"),
			h.Class("text-prussian text-sm"),
			h.Text("No sightings added."),
		)
	}
	return h.Div(
		h.Id("sightings-list"),
		h.Class("flex flex-col gap-3"),
		h.List(sightings, func(s components.Sight, _ int) *h.Element {
			return components.SightCard(s)
		}),
	)
}

func PostSightingPartial(ctx *h.RequestContext) *h.Partial {
	place := ctx.FormValue("place")
	date := ctx.FormValue("date")
	tz := ctx.FormValue("tz")

	errorPartial := h.NewPartial(h.Div(
		h.Id("sighting-list"),
		h.Class("text-vermilion text-sm"),
		h.Text("Invalid TZ number. Please enter a valid integer"),
	))

	resp, err := http.PostForm("http://localhost:8080/sightings", url.Values{
		"place": {place},
		"date":  {date},
		"tz":    {tz},
	})
	if err != nil || resp.StatusCode >= 400 {
		return errorPartial
	}

	GetSightings()

	return h.NewPartial(SightList(knownSightings))
}

func AddSightingForm() *h.Element {
	return h.Div(
		h.Class("bg-white border border-saffron rounded-lg shadow-sm p-6"),
		h.H2(
			h.Class("text-lg font-bold text-prussian mb-4"),
			h.Text("Add a sighting"),
		),
		h.Form(
			h.Attribute("id", "add-sighting-form"),
			h.PostPartial(PostSightingPartial),
			h.HxTarget("#sightings-list"),
			h.Class("flex flex-col gap-4"),
			components.FormField("Place", "place", "e.g. Berlin"),
			components.FormField("Date", "date", "e.g. 15.01.2026"),
			components.FormField("TZ", "tz", "e.g. 9001"),
			h.Button(
				h.Type("submit"),
				h.Class("mt-2 bg-prussian text-champagne font-semibold rounded px-4 py-2 hover:bg-vermilion transition-colors"),
				h.Text("Add sighting"),
			),
		),
	)
}
