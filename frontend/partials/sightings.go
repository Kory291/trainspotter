package partials

import (
	"io"
	"slices"
	"trainspotter-frontend/components"

	"encoding/json"
	"net/http"

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

func PostSighting(ctx *h.RequestContext) {
	
}