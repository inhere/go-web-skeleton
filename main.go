package main

import (
	"fmt"
	"log"
	"os"

	"github.com/inhere/go-web-skeleton/api"
	// boot and init some services(log, cache, eureka)
	"github.com/inhere/go-web-skeleton/app"

	// init redis, mongo, mysql connection
	_ "github.com/inhere/go-web-skeleton/model/mongo"
	_ "github.com/inhere/go-web-skeleton/model/mysql"
	_ "github.com/inhere/go-web-skeleton/model/rds"

	"github.com/gookit/rux"
	"github.com/gookit/rux/handlers"
)

var router *rux.Router

func init() {
	app.BootWeb()

	// router and routes
	router = rux.New()
	// global middleware
	router.Use(handlers.RequestLogger())

	api.AddRoutes(router)
}

func main() {
	log.Printf("======================== Begin Running(PID: %d) ========================", os.Getpid())

	// default is listen and serve on 0.0.0.0:8080
	router.Listen(fmt.Sprintf("0.0.0.0:%d", app.HttpPort))
}
