package api

import (
	"github.com/gookit/sux"
	"github.com/inhere/go-web-skeleton/app"
)

// InternalApi
type InternalApi struct {
	BaseApi
}

// @Tags InternalApi
// @Summary Get app config
// @Param	key		query 	string	 false		"config key string"
// @Success 201 {string} json data
// @Failure 403 body is empty
// @router /config [get]
func (a *InternalApi) Config(c *sux.Context) {
	key := c.Query("key")
	if key == "" {
		key = app.Cfg.DefSection()
	}

	val, _ := app.Cfg.StringMap(key)

	a.JSON(c, 200, val)
}
