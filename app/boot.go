package app

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/gookit/config/v2/ini"
	"github.com/gookit/i18n"
	"github.com/gookit/rux"

	"github.com/gookit/config/v2"
	"github.com/inhere/go-web-skeleton/app/helper"
	"github.com/inhere/go-web-skeleton/model"

	"github.com/gookit/view"
	"github.com/inhere/go-web-skeleton/app/cache"
)

// components of the application
var (
	View *view.Renderer
)

// BootWeb Bootstrap web application
func BootWeb() {
	initApp()

	initAppEnv()

	loadAppConfig()

	log.Printf(
		"======================== Bootstrap (Env: %s, Debug: %v) ========================",
		Env, debug,
	)

	initAppInfo()

	initLogger()

	initLanguage()

	initCache()

	// initEurekaService()

	listenSignals()
}

func initApp() {
	// view templates
	View = view.NewInitialized(func(r *view.Renderer) {
		r.ViewsDir = "resource/views"
	})
}

func initAppEnv() {
	Hostname, _ = os.Hostname()
	if env := os.Getenv("APP_ENV"); env != "" {
		Env = env
	}

	if port := os.Getenv("APP_PORT"); port != "" {
		HttpPort, _ = strconv.Atoi(port)
	}

	// in dev, test
	if IsEnv(DEV) || IsEnv(TEST) {
		rux.Debug(true)
	} else {
		rux.Debug(false)
	}
}

// loadAppConfig
func loadAppConfig() {
	var err error

	// add ini driver
	config.AddDriver(ini.Driver)

	envFile := "config/app-" + Env + ".ini"

	fmt.Printf("- work dir: %s\n", WorkDir)
	fmt.Printf("- load config: config/app.ini, %s\n", envFile)

	err = config.LoadFiles("config/app.ini", envFile)
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	// ini.WriteTo(os.Stdout)

	// setting some info
	Name = config.String("name")
	debug = config.Bool("debug")
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
		helper.ReadJsonFile(infoFile, &GitInfo)
	}
}

// init redis connection pool
func initCache() {
	conf := config.StringMap("cache")
	if conf["enable"] == "0" {
		Printf("cache is disabled, skip init it")
		return
	}

	// 从配置文件获取redis的ip以及db
	prefix := conf["prefix"]
	server := conf["server"]
	password := conf["auth"]
	redisDb, _ := strconv.Atoi(conf["db"])

	fmt.Printf("cache config - server=%s db=%d auth=%s\n", server, redisDb, password)

	// 建立连接池
	// closePool()
	cache.Init(NewRedisPool(server, password, redisDb), prefix, debug)
}

func initLanguage() {
	// conf := map[string]string{
	// 	"langDir": "res/lang",
	// 	"allowed": "en:English|zh-CN:简体中文",
	// 	"default": "en",
	// }
	conf := config.StringMap("lang")
	fmt.Printf("language config - %v\n", conf)

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
