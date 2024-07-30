package main

import (
	"net/http"
)

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	app.render(w, "home.page.gohtml", nil)
}
