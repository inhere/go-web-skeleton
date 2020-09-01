package main

import (
	"fmt"
	"os"

	"github.com/inhere/go-web-skeleton/app/bootstrap"
	"github.com/inhere/go-web-skeleton/app/clog"
	"github.com/inhere/go-web-skeleton/web"
	// boot and init some services(log, cache, eureka)
	"github.com/inhere/go-web-skeleton/app"

	// init redis, mongo, mysql connection

	"github.com/gookit/rux"
	"github.com/gookit/rux/handlers"
)

var router *rux.Router

func init() {
	bootstrap.Web()

	// router and routes
	router = rux.New()
	// global middleware
	router.Use(handlers.RequestLogger())

	web.AddRoutes(router)
}

// @title My Project API
// @version 1.0
// @description My Project API
// @termsOfService https://github.com/inhere

// @contact.name API Support
// @contact.url https://github.com/inhere
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /v1
func main() {
	clog.Printf(
		"======================== Begin Running(PID: %d) ========================\n",
		os.Getpid(),
	)

	// default is listen and serve on 0.0.0.0:8080
	router.Listen(fmt.Sprintf("0.0.0.0:%d", app.HttpPort))
}
