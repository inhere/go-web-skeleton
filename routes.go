package main

import (
	"github.com/gookit/rux"
	"github.com/inhere/go-web-skeleton/controller/api"
)

func addRoutes(r *rux.Router) {
	// static assets
	r.StaticDir("/static", "static")

	r.GET("/", api.Home)
	r.GET("/api-docs", api.SwagDoc)

	// status
	r.GET("/health", api.AppHealth)
	r.GET("/status", api.AppStatus)

	r.GET("/ping", func(c *rux.Context) {
		c.Text(200, "pong")
	})

	r.Group("/v1", func() {
		r.GET("/health", api.AppHealth)

		internal := new(api.InternalApi)
		r.GET("/config", internal.Config)
	})

	// not found routes
	r.NotFound(func(c *rux.Context) {
		c.JSONBytes(404, []byte(`{"code": 0, "message": "page not found", data: {}}`))
	})
}
