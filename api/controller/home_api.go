package controller

import (
	"os"

	"github.com/gookit/rux"
	"github.com/inhere/go-web-skeleton/app"
)

func Home(c *rux.Context) {
	c.JSON(200, rux.M{"hello": "welcome"})
}

func SwagDoc(c *rux.Context) {
	fInfo, _ := os.Stat("static/swagger.json")

	data := map[string]string{
		"Env":        app.Env,
		"AppName":    app.Name,
		"JsonFile":   "/static/swagger.json",
		"SwgUIPath":  "/static/swagger-ui",
		"AssetPath":  "/static",
		"UpdateTime": fInfo.ModTime().Format(app.BaseDate),
	}

	// c.HTML(200, nil)
	app.View.Partial(c.Resp, "swagger.tpl", data)
}

// @Tags InternalApi
// @Summary 检测API
// @Description get app health
// @Success 201 {string} json data
// @Failure 403 body is empty
// @Router /health [get]
func AppHealth(c *rux.Context) {
	data := map[string]interface{}{
		"status": "UP",
		"info":   app.GitInfo,
	}

	c.JSON(200, data)
}

func AppStatus(c *rux.Context) {
	data := map[string]interface{}{
		"status": "UP",
		"info":   app.GitInfo,
	}

	c.JSON(200, data)
}
