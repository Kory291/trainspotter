package partials

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
	"trainspotter-frontend/components"
	"strconv"
	"slices"
	"github.com/maddalax/htmgo/framework/h"
)

// KnownTrains holds placeholder trains – replace with real data once the backend is connected.
var knownTrains = []components.Train{}

func TrainListPartial(ctx *h.RequestContext) *h.Partial {
	query := strings.TrimSpace(ctx.FormValue("tz_filter"))

	var filtered []components.Train
	for _, t := range knownTrains {
		tzStr := strconv.Itoa(t.Tz)
		if query == "" || strings.Contains(tzStr, query) {
			filtered = append(filtered, t)
		}
	}

	return h.NewPartial(TrainList(filtered))
}

// GetKnownTrains returns the current list of trains.
func GetKnownTrains() []components.Train {
	// return knownTrains
	resp, err := http.Get("http://localhost:8080/trains")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Parse the response body and return the list of trains
	// This is a placeholder – replace with actual parsing logic
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(body, &knownTrains)
	slices.SortStableFunc(knownTrains, func(trainA components.Train, trainB components.Train) int {
		return trainA.Tz - trainB.Tz
	})  
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

func AddTrainForm() *h.Element {
	return h.Div(
		h.Class("bg-white border border-saffron rounded-lg shadow-sm p-6"),
		h.H2(
			h.Class("text-lg font-bold text-prussian mb-4"),
			h.Text("Add a Train"),
		),
		h.Form(
			h.Attribute("id", "add-train-form"),
			h.PostPartial(AddTrainPartial),
			h.HxTarget("#train-list"),
			h.Class("flex flex-col gap-4"),
			components.FormField("TZ Number", "tz", "e.g. 9001"),
			components.FormField("Baureihe", "baureihe", "e.g. 407"),
			components.FormField("Name", "name", "e.g. München"),
			h.Button(
				h.Type("submit"),
				h.Class("mt-2 bg-prussian text-champagne font-semibold rounded px-4 py-2 hover:bg-vermilion transition-colors"),
				h.Text("Add Train"),
			),
		),
	)
}

func AddTrainPartial(ctx *h.RequestContext) *h.Partial {
	tz := ctx.FormValue("tz")
	baureihe := ctx.FormValue("baureihe")
	name := ctx.FormValue("name")
	
	tzInt, err := strconv.Atoi(tz)
	if err != nil {
		return h.NewPartial(h.Div(
			h.Id("train-list"),
			h.Class("text-vermilion text-sm"),
			h.Text("Invalid TZ number. Please enter a valid integer."),
		))
	}
	resp, err := http.PostForm("http://localhost:8080/trains", url.Values{
		"tz":       {tz},
		"baureihe": {baureihe},
		"name":     {name},
	})
	if err != nil || resp.StatusCode >= 400 {
		return h.NewPartial(h.Div(
			h.Id("train-list"),
			h.Class("text-vermilion text-sm"),
			h.Text("Error adding train. Please try again."),
		))
	}
	knownTrains = append(knownTrains, components.Train{
		Tz:       tzInt,
		Baureihe: baureihe,
		Name:     name,
	})

	return h.NewPartial(TrainList(knownTrains))
}
