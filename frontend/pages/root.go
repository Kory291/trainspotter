package pages

import (
	"github.com/maddalax/htmgo/framework/h"
	"trainspotter-frontend/__htmgo/assets"
)

func RootPage(children ...h.Ren) *h.Page {
	title := "Trainspotter Frontend"
	description := "A simple frontend for my trainspotter project, built with htmgo."
	author := "Kory291"
	url := "https://github.com/Kory291/trainspotter"

	return h.NewPage(
		h.Html(
			h.HxExtensions(
				h.BaseExtensions(),
			),
			h.Head(
				h.Title(
					h.Text(title),
				),
				h.Meta("viewport", "width=device-width, initial-scale=1"),
				h.Link(assets.FaviconIco, "icon"),
				h.Link(assets.AppleTouchIconPng, "apple-touch-icon"),
				h.Meta("title", title),
				h.Meta("charset", "utf-8"),
				h.Meta("author", author),
				h.Meta("description", description),
				h.Meta("og:title", title),
				h.Meta("og:url", url),
				h.Link("canonical", url),
				h.Meta("og:description", description),
				h.Link(assets.MainCss, "stylesheet"),
				h.Script(assets.HtmgoJs),
				h.Script(`
					function submitTrainForm(form) {
						const formData = new FormData(form);
						fetch("http://localhost:8080/trains", {
							method: "POST",
							headers: { "Content-Type": "application/x-www-form-urlencoded" },
							body: new URLSearchParams(formData).toString()
						})
							.then((response) => {
								if (!response.ok) {
									throw new Error("Request failed");
								}
								form.reset();
							})
							.catch((error) => {
								console.error(error);
							});
					}
				`),
			),
			h.Body(
				h.Div(
					h.Fragment(children...),
				),
			),
		),
	)
}
