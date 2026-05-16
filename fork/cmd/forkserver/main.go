package main

import (
	"net/http"

	"miaomiaowu/extensions/basepath"
	"miaomiaowu/internal/app"
	"miaomiaowu/internal/publicpath"
)

func main() {
	cfg := basepath.FromEnv()
	publicpath.SetProvider(basepath.NewProvider(cfg))
	app.Run(app.WithHandlerWrap(func(h http.Handler) http.Handler {
		return basepath.Wrap(h, cfg)
	}))
}
