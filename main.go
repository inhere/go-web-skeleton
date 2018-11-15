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
	"github.com/gookit/view"
	"log"
	"os"
)

var router *rux.Router

func init() {
	app.Boot()

	// router and routes
	router = rux.New()
	// global middleware
	router.Use(handlers.RequestLogger())

	addRoutes(router)
}

func main() {
	log.Printf("======================== Begin Running(PID: %d) ========================", os.Getpid())

	// default is listen and serve on 0.0.0.0:8080
	router.Listen(fmt.Sprintf("0.0.0.0:%d", app.HttpPort))
}
