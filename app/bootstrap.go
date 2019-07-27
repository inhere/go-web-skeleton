package app

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/gookit/config/v2/dotnev"
	"github.com/gookit/config/v2/ini"
	"github.com/gookit/gcli/v2/show"
	"github.com/gookit/goutil/jsonutil"
	"github.com/gookit/i18n"
	"github.com/gookit/rux"
	"github.com/inhere/go-web-skeleton/app/clog"
	"github.com/inhere/go-web-skeleton/app/listener"
	"github.com/inhere/go-web-skeleton/model/mongo"
	"github.com/inhere/go-web-skeleton/model/mysql"
	"github.com/inhere/go-web-skeleton/model/rds"

	"github.com/gookit/config/v2"
	"github.com/inhere/go-web-skeleton/app/helper"
	"github.com/inhere/go-web-skeleton/model"

	"github.com/gookit/view"
	"github.com/inhere/go-web-skeleton/app/cache"
)

// components of the application

// BootWeb Bootstrap web application
func BootWeb() {
	initEnv()

	initConfig()

	initApp()

	initAppInfo()

	initLogger()

	initI18n()

	// init cache redis connection pool
	cache.Init(debug)

	// initEurekaService()

	// listen exit signal
	listener.ListenSignals(onExit)
}

func initEnv() {
	err := dotnev.LoadExists("./", ".env")
	if err != nil {
		clog.Fatalf("Fail to load env file: %v", err)
	}

	Hostname, _ = os.Hostname()
	if env := config.GetEnv("APP_ENV"); env != "" {
		Env = env
	}

	if port := config.GetEnv("APP_PORT"); port != "" {
		HttpPort, _ = strconv.Atoi(port)
	}

	// in dev, test
	if IsEnv(DEV) || IsEnv(TEST) {
		rux.Debug(true)
	} else {
		rux.Debug(false)
	}
}

// initConfig load app config
func initConfig() {
	envFile := "config/app-" + Env + ".ini"

	show.AList("project information", map[string]string{
		"Work directory": WorkDir,
		"Loaded config":  "config/app.ini, " + envFile,
	}, nil)

	// fmt.Printf("- work directory: %s\n", WorkDir)
	// fmt.Printf("- loaded config: config/app.ini, %s\n", envFile)

	// add ini driver
	config.AddDriver(ini.Driver)
	config.WithOptions(config.Readonly)

	err := config.LoadFiles("config/app.ini", envFile)
	if err != nil {
		clog.Fatalf("Fail to read file: %v", err)
	}

	// setting some info
	// _= config.LoadData(map[string]interface{}{
	// 	"env": Env,
	// 	"debug": debug,
	// })
	Name = config.String("name")
	debug = config.Bool("debug")

	fmt.Printf(
		"======================== Bootstrap (Env: %s, Debug: %v) ========================\n",
		Env, debug,
	)

	clog.SetDebug(debug)
}

func initApp() {
	// view templates
	view.Initialize(func(r *view.Renderer) {
		r.ViewsDir = "resource/views"
	})
}

func initAppInfo() {
	// ensure http port
	if HttpPort == 0 {
		HttpPort = config.Int("httpPort")
	}

	// git repo info
	GitInfo = model.GitInfoData{}
	infoFile := "static/app.json"

	if helper.FileExists(infoFile) {
		_ = jsonutil.ReadFile(infoFile, &GitInfo)
	}
}

func initI18n() {
	// conf := map[string]string{
	// 	"langDir": "res/lang",
	// 	"allowed": "en:English|zh-CN:简体中文",
	// 	"default": "en",
	// }
	conf := config.StringMap("lang")
	clog.Printf("language config - %v", conf)

	// en:English|zh-CN:简体中文
	langList := strings.Split(conf["allowed"], "|")
	languages := make(map[string]string, len(langList))

	for _, str := range langList {
		item := strings.Split(str, ":")
		languages[item[0]] = item[1]
	}

	// init and load data
	i18n.Init(conf["langDir"], conf["default"], languages)
}

func onExit() {
	var err error
	// sync logs
	// logrus.

	// unregister from eureka
	// erkServer.Unregister()

	// close db,redis connection
	mongo.CloseSession()

	err = mysql.CloseEngine()
	if err != nil {
		clog.Errorf("Close mysql error: %s", err.Error())
	}

	err = rds.ClosePool()
	if err != nil {
		clog.Errorf("Close redis error: %s", err.Error())
	}

}
