package api

import (
	"github.com/gookit/sux"
	"github.com/inhere/go-web-skeleton/app"
	"os"
)

func Home(c *sux.Context) {
	c.JSON(200, sux.M{"hello": "welcome"})
}

func SwagDoc(c *sux.Context) {
	fInfo, _ := os.Stat("static/swagger.json")

	data := map[string]string{
		"Env":        app.Env,
		"AppName":    app.Name,
		"JsonFile":   "/static/swagger.json",
		"SwgUIPath":  "/static/swagger-ui",
		"AssetPath":  "/static",
		"UpdateTime": fInfo.ModTime().Format(app.BaseDate),
	}

	c.HTML(200, "swagger.tpl", data)
}

// @Tags InternalApi
// @Summary 检测API
// @Description get app health
// @Success 201 {string} json data
// @Failure 403 body is empty
// @Router /health [get]
func AppHealth(c *sux.Context) {
	data := map[string]interface{}{
		"status": "UP",
		"info":   app.GitInfo,
	}

	c.JSON(200, data)
}

func AppStatus(c *sux.Context) {
	data := map[string]interface{}{
		"status": "UP",
		"info":   app.GitInfo,
	}

	c.JSON(200, data)
}
