package app

import (
	"github.com/gookit/sux"
	"github.com/inhere/go-web-skeleton/api"
)

// AddRoutes
func AddRoutes(r *sux.Router) {
	r.GET("/", api.Home)

	r.LoadHTMLFiles("res/views/swagger.tpl")
	r.GET("/api-docs", api.SwagDoc)

	// status
	r.GET("/health", api.AppHealth)
	r.GET("/status", api.AppStatus)

	r.GET("/ping", func(c *sux.Context) {
		c.Text(200, "pong")
	})

	r.Group("/v1", func() {
		r.GET("/health", api.AppHealth)

		internal := new(api.InternalApi)
		r.GET("/config", internal.Config)
	})

	// static assets
	r.StaticDir("/static", "./static")

	// not found routes
	r.NotFound(func(c *sux.Context) {
		c.JSONBytes(404, []byte(`{"code": 0, "message": "page not found", data: {}}`))
	})
}
