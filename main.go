package main

import (
	"fmt"
	// boot and init some services(log, cache, eureka)
	"github.com/inhere/go-web-skeleton/app"

	// init redis, mongo, mysql connection
	_ "github.com/inhere/go-web-skeleton/model/mongo"
	_ "github.com/inhere/go-web-skeleton/model/mysql"
	_ "github.com/inhere/go-web-skeleton/model/rds"

	"github.com/gookit/sux"
	"github.com/gookit/sux/handlers"
	"log"
	"os"
	"github.com/gookit/view"
)

var router *sux.Router

func init() {
	app.Boot()

	// view templates
	v := view.NewInitialized(func(r *view.Renderer) {
		r.ViewsDir = "resource/views"
	})

	// router and routes
	router = sux.New()
	if app.IsEnv(app.DEV) {
		sux.Debug(true)
	}

	// global middleware
	router.Use(handlers.RequestLogger())

	addRoutes(router)
}

func main() {
	log.Printf("======================== Begin Running(PID: %d) ========================", os.Getpid())

	// default is listen and serve on 0.0.0.0:8080
	router.Listen(fmt.Sprintf("0.0.0.0:%d", app.HttpPort))
}
