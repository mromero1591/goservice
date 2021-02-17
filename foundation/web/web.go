// Package web contains a small web framework extension.
package web

import "github.com/dimfeld/httptreemux/v5"

type App struct {
	//this makes it so app is everything a context mux - this is calleld embedding
	*httptreemux.ContextMux
}

func NewApp() *App {
	app := App{
		ContextMux: httptreemux.NewContextMux(),
	}

	return &app
}
