package route

import (
	"github.com/gookit/sux"
	"github.com/inhere/go-web-skeleton/api"
)

// AddRoutes
func AddRoutes(r *gin.Engine) {
	r.GET("/", api.Home)

	r.LoadHTMLFiles("res/views/swagger.tpl")
	r.GET("/api-docs", api.SwagDoc)

	// status
	r.GET("/health", api.AppHealth)
	r.GET("/status", api.AppStatus)

	r.GET("/ping", func(c *sux.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	v1 := r.Group("/v1")
	{
		v1.GET("/health", api.AppHealth)

		internal := new(api.InternalApi)
		v1.GET("/config", internal.Config)
	}

	// static assets
	r.Static("/static", "./static")

	// not found routes
	r.NoRoute(func(c *sux.Context) {
		c.JSON(
			404,
			api.JsonMapData{0, "page not found", map[string]string{}},
		)
	})
}
