package pages

import (
	"time"
	"trainspotter-frontend/components"

	"github.com/maddalax/htmgo/framework/h"
)

// placeholder sightings – replace with real data once the backend is connected
var recentSights = []components.Sight{
	{Train: "TZ 9001", Location: "München Hbf", Date: time.Date(2026, 7, 10, 9, 12, 0, 0, time.Local)},
	{Train: "TZ 9023", Location: "Nürnberg Hbf", Date: time.Date(2026, 7, 9, 14, 3, 0, 0, time.Local)},
	{Train: "TZ 9105", Location: "Frankfurt Hbf", Date: time.Date(2026, 7, 8, 11, 45, 0, 0, time.Local)},
	{Train: "TZ 9012", Location: "Hamburg Hbf", Date: time.Date(2026, 7, 7, 8, 22, 0, 0, time.Local)},
	{Train: "TZ 9047", Location: "Berlin Hbf", Date: time.Date(2026, 7, 6, 16, 55, 0, 0, time.Local)},
	{Train: "TZ 9088", Location: "Köln Hbf", Date: time.Date(2026, 7, 5, 10, 30, 0, 0, time.Local)},
	{Train: "TZ 9034", Location: "Stuttgart Hbf", Date: time.Date(2026, 7, 4, 13, 18, 0, 0, time.Local)},
	{Train: "TZ 9071", Location: "Mannheim Hbf", Date: time.Date(2026, 7, 3, 7, 50, 0, 0, time.Local)},
	{Train: "TZ 9056", Location: "Leipzig Hbf", Date: time.Date(2026, 7, 2, 18, 5, 0, 0, time.Local)},
	{Train: "TZ 9019", Location: "Hannover Hbf", Date: time.Date(2026, 7, 1, 12, 40, 0, 0, time.Local)},
}

func IndexPage(ctx *h.RequestContext) *h.Page {
	return RootPage(
		components.Nav(ctx),
		h.Div(
			h.Class("p-6 max-w-2xl mx-auto"),
			h.H1(
				h.Class("text-2xl font-bold text-prussian mb-6"),
				h.Text("Recent Sightings"),
			),
			h.Div(
				h.Class("flex flex-col gap-4"),
				h.List(recentSights, func(s components.Sight, _ int) *h.Element {
					return components.SightCard(s)
				}),
			),
		),
	)
}
