package bootstrap

import (
	"github.com/inhere/go-web-skeleton/app"
	"github.com/inhere/go-web-skeleton/app/cache"
	"github.com/inhere/go-web-skeleton/app/listener"
	"github.com/inhere/go-web-skeleton/model/mongo"
)

// Web Bootstrap web application
func Web() {
	l := &Launcher{}
	l.Add(
		BootFunc(initEnv),
		BootFunc(initConfig),
		BootFunc(initApp),
		BootFunc(func() error {
			initAppInfo()
			return nil
		}),
		BootFunc(app.InitLogger),
		BootFunc(initI18n),
		// init cache redis connection pool
		BootFunc(cache.InitCache),
		BootFunc(mongo.InitMongo),
	)

	l.Run()

	// initEurekaService()

	// listen exit signal
	listener.ListenSignals(onExit)
}
