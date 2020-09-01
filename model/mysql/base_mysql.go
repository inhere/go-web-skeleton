package mysql

import (
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/gookit/config/v2"
	"github.com/inhere/go-web-skeleton/app"
	"github.com/inhere/go-web-skeleton/app/clog"
	"github.com/inhere/go-web-skeleton/app/helper"
	"xorm.io/core"
)

const DSNTemplate = "%s:%s@tcp(%s:%d)/%s?charset=utf8"

type dbConfig struct {
	Host string
	Port int
	User string
	Name string
	Password string

	Disable bool
	MaxIdleConn int
	MaxOpenConn int
}

var (
	cfg dbConfig
	engine *xorm.Engine
)

func InitMysql() (err error) {
	err = config.MapStruct("db", &cfg)
	if err != nil {
		return
	}

	if cfg.Disable {
		clog.Debugf("mysql is disabled, skip init mysql database connection")
		return
	}

	dsn := fmt.Sprintf(DSNTemplate, cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name, )
	fmt.Printf("mysql - %s\n", dsn)

	// create engine
	engine, err = xorm.NewEngine("mysql", dsn)
	if err != nil {
		// log.Fatalf("Init mysql DB Failure! Error: %s\n", err.Error())
		return
	}

	engine.SetMaxIdleConns(cfg.MaxIdleConn)
	engine.SetMaxOpenConns(cfg.MaxOpenConn)

	// core.NewCacheMapper(core.SnakeMapper{})
	// engine.SetDefaultCacher()
	if app.Debug {
		engine.ShowSQL(true)
		engine.Logger().SetLevel(core.LOG_DEBUG)
	}

	// replace
	logFile := config.String("log.sqlLog")
	logFile = strings.NewReplacer(
		"{date}", helper.LocTime().Format("20060102"),
		"{hostname}", app.Hostname,
	).Replace(logFile)

	f, err := os.Create(logFile)
	if err != nil {
		log.Fatal(err.Error())
	}

	engine.SetLogger(xorm.NewSimpleLogger(f))
	return
}

// Db get db connection
func Db() *xorm.Engine {
	return engine
}

// CloseEngine Close mysql engine
func CloseEngine() error {
	if cfg.Disable {
		return nil
	}

	return engine.Close()
}

// UpdateById Update by ID
// Usage:
// user := new(User)
// num, err := mysql.UpdateById(23, user, "name", "email")
func UpdateById(id int64, model interface{}, fields ...string) (affected int64, err error) {
	affected, err = engine.ID(id).Cols(fields...).Update(model)
	return
}

// DeleteById
// Usage:
// user := new(User)
// num, err := mysql.DeleteById(23, user)
func DeleteById(id int64, model interface{}) (affected int64, err error) {
	affected, err = engine.ID(id).Delete(model)
	return
}
