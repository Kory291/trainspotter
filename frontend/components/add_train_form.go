package components

import (
	"github.com/maddalax/htmgo/framework/h"
)

func formField(label, name, placeholder string) *h.Element {
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

func AddTrainForm() *h.Element {
	return h.Div(
		h.Class("bg-white border border-saffron rounded-lg shadow-sm p-6"),
		h.H2(
			h.Class("text-lg font-bold text-prussian mb-4"),
			h.Text("Add a Train"),
		),
		h.Form(
			h.Method("post"),
			h.Action("#"),
			h.Attribute("id", "add-train-form"),
			h.Attribute("onSubmit", "event.preventDefault(); submitTrainForm(this);"),
			h.Class("flex flex-col gap-4"),
			formField("TZ Number", "tz", "e.g. 9001"),
			formField("Baureihe", "baureihe", "e.g. ICE 3"),
			formField("Name", "name", "e.g. München"),
			h.Button(
				h.Type("submit"),
				h.Class("mt-2 bg-prussian text-champagne font-semibold rounded px-4 py-2 hover:bg-vermilion transition-colors"),
				h.Text("Add Train"),
				
			),
		),
	)
}
