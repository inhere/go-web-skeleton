package api

import (
	"github.com/gookit/rux"
	"github.com/inhere/go-web-skeleton/api/controller"
)

// AddRoutes add http routes
func AddRoutes(r *rux.Router) {
	// static assets
	r.StaticDir("/static", "static")

	r.GET("/", controller.Home)
	r.GET("/api-docs", controller.SwagDoc)

	// status
	r.GET("/health", controller.AppHealth)
	r.GET("/status", controller.AppStatus)

	r.GET("/ping", func(c *rux.Context) {
		c.Text(200, "pong")
	})

	r.Group("/v1", func() {
		r.GET("/health", controller.AppHealth)

		internal := new(controller.InternalApi)
		r.GET("/config", internal.Config)
	})

	// not found routes
	r.NotFound(func(c *rux.Context) {
		c.JSONBytes(404, []byte(`{"code": 0, "message": "page not found", data: {}}`))
	})
}
