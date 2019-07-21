package app

import (
	"fmt"
	"github.com/gookit/i18n"
	"github.com/gookit/rux"
	"os"
	"strings"

	"github.com/gookit/ini/v2"
	"github.com/inhere/go-web-skeleton/app/utils"
	"github.com/inhere/go-web-skeleton/model"

	"github.com/gookit/view"
	"github.com/inhere/go-web-skeleton/app/cache"
	"log"
	"strconv"
)

// components of the application
var (
	Cfg  *ini.Ini
	View *view.Renderer
)

// Boot app
func Boot() {
	initApp()

	initAppEnv()

	loadAppConfig()

	log.Printf(
		"======================== Bootstrap (Env: %s, Debug: %v) ========================",
		Env, Debug,
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

	envFile := "conf/app-" + Env + ".ini"

	fmt.Printf("- work dir: %s\n", WorkDir)
	fmt.Printf("- load config: conf/app.ini, %s\n", envFile)

	Cfg, err = ini.LoadFiles("conf/app.ini", envFile)
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	// Cfg.WriteTo(os.Stdout)

	// setting some info
	Name = Cfg.MustString("name")
	Debug = Cfg.MustBool("debug")
}

func initAppInfo() {
	// ensure http port
	if HttpPort == 0 {
		HttpPort = Cfg.MustInt("httpPort")
	}

	// git repo info
	GitInfo = model.GitInfoData{}
	infoFile := "static/app.json"

	if utils.FileExists(infoFile) {
		utils.ReadJsonFile(infoFile, &GitInfo)
	}
}

// init redis connection pool
func initCache() {
	conf, _ := Cfg.StringMap("cache")

	// 从配置文件获取redis的ip以及db
	prefix := conf["prefix"]
	server := conf["server"]
	password := conf["auth"]
	redisDb, _ := strconv.Atoi(conf["db"])

	fmt.Printf("cache - server=%s db=%d auth=%s\n", server, redisDb, password)

	// 建立连接池
	// closePool()
	cache.Init(NewRedisPool(server, password, redisDb), prefix, Logger, Debug)
}

func initLanguage() {
	// conf := map[string]string{
	// 	"langDir": "res/lang",
	// 	"allowed": "en:English|zh-CN:简体中文",
	// 	"default": "en",
	// }
	conf, _ := Cfg.StringMap("lang")
	fmt.Printf("language - %v\n", conf)

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
