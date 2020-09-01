package bootstrap

import (
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/gookit/color"
	"github.com/gookit/config/v2/dotnev"
	"github.com/gookit/config/v2/toml"
	"github.com/gookit/gcli/v2/show"
	"github.com/gookit/goutil/fsutil"
	"github.com/gookit/goutil/jsonutil"
	"github.com/gookit/i18n"
	"github.com/gookit/rux"
	"github.com/inhere/go-web-skeleton/app"
	"github.com/inhere/go-web-skeleton/app/clog"
	"github.com/inhere/go-web-skeleton/model/mongo"
	"github.com/inhere/go-web-skeleton/model/myrds"
	"github.com/inhere/go-web-skeleton/model/mysql"

	"github.com/gookit/config/v2"
	"github.com/inhere/go-web-skeleton/model"

	"github.com/gookit/view"
)

// components of the application

func initEnv() error {
	err := dotnev.LoadExists("./", ".env")
	if err != nil {
		return err
	}

	app.Hostname, _ = os.Hostname()
	if env := config.Getenv("APP_ENV"); env != "" {
		app.EnvName = env
	}

	if port := config.Getenv("APP_PORT"); port != "" {
		app.HttpPort, _ = strconv.Atoi(port)
	}

	// in dev, test
	if app.IsEnv(app.EnvDev) || app.IsEnv(app.EnvTest) {
		rux.Debug(true)
	} else {
		rux.Debug(false)
	}

	return nil
}

// initConfig load app config
func initConfig()error {
	baseFile := "config/app" + app.ConfigSuffix
	envFile := "config/app-" + app.EnvName + app.ConfigSuffix

	show.AList("project information", map[string]string{
		"Work directory": app.WorkDir,
		"Loaded config":  baseFile + ", " + envFile,
	}, nil)

	// fmt.Printf("- work directory: %s\n", WorkDir)
	// fmt.Printf("- loaded config: config/app.ini, %s\n", envFile)

	// add ini driver
	config.AddDriver(toml.Driver)
	config.WithOptions(config.Readonly)

	err := config.LoadFiles(baseFile, envFile)
	if err != nil {
		return err
	}

	// setting some info
	// _= config.LoadData(map[string]interface{}{
	// 	"env": EnvName,
	// 	"debug": debug,
	// })
	app.Name = config.String("name")
	app.Debug = config.Bool("debug")

	clog.Printf(
		"======================== Bootstrap (EnvName: %s, Debug: %v) ========================\n",
		app.EnvName, app.Debug,
	)

	clog.SetDebug(app.Debug)
	return nil
}

func initApp() error {
	// view templates
	view.Initialize(func(r *view.Renderer) {
		r.ViewsDir = "resource/views"
	})

	return nil
}

func initAppInfo() {
	// ensure http port
	if app.HttpPort == 0 {
		app.HttpPort = config.Int("httpPort")
	}

	// git repo info
	app.GitInfo = model.GitInfo{}
	infoFile := "static/app.json"

	if fsutil.IsFile(infoFile) {
		err := jsonutil.ReadFile(infoFile, &app.GitInfo)
		if err != nil {
			color.Error.Println(err.Error())
		}
	}
}

func initI18n() error {
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

	return nil
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

	err = myrds.ClosePool()
	if err != nil {
		clog.Errorf("Close redis error: %s", err.Error())
	}
}

// 获取某个文件夹下的配置文件列表
func getConfigFiles(confDir string) ([]string, error) {
	var files = make([]string, 0)

	fileInfoList, err := ioutil.ReadDir(confDir)
	if err != nil {
		return files, err
	}

	pathSep := string(os.PathSeparator)
	// app.toml is must exists
	baseFile := confDir + pathSep + "app" + app.ConfigSuffix
	files = append(files, baseFile)

	// _dev.toml
	suffix := "-" + app.EnvName + app.ConfigSuffix
	for _, f := range fileInfoList {
		// app_dev.toml
		if !f.IsDir() && strings.HasSuffix(f.Name(), suffix) {
			files = append(files, confDir+pathSep+f.Name())
		}
	}

	return files, err
}
