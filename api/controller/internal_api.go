package controller

import (
	"github.com/gookit/config/v2"
	"github.com/gookit/ini/v2"
	"github.com/gookit/rux"
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
func (a *InternalApi) Config(c *rux.Context) {
	key := c.Query("key")
	if key == "" {
		a.JSON(c, 200, config.Data())
	}

	val := ini.StringMap(key)

	a.JSON(c, 200, val)
}
