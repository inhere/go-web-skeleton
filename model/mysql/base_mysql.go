package mysql

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/gookit/config/v2"
	"github.com/inhere/go-web-skeleton/app/clog"
	"github.com/inhere/go-web-skeleton/app/helper"
)

var (
	debug bool
	enable bool
	engine *xorm.Engine
)

func init() {
	debug = config.Bool("debug")
	enable = config.Bool("db.enable")
	if !enable {
		clog.Debugf("mysql is disabled, skip init mysql database connection")
		return
	}

	db := config.StringMap("db")
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8",
		db["user"], db["password"], db["host"], db["port"], db["name"],
	)

	clog.Printf("mysql config - %s\n", dsn)
	var err error

	maxIdleConn, _ := strconv.Atoi(db["maxIdleConn"])
	maxOpenConn, _ := strconv.Atoi(db["MaxOpenConn"])

	// create engine
	engine, err = xorm.NewEngine("mysql", dsn)
	if err != nil {
		log.Fatalf("Init mysql DB Failure! Error: %s\n", err.Error())
	}

	engine.SetMaxIdleConns(maxIdleConn)
	engine.SetMaxOpenConns(maxOpenConn)

	// core.NewCacheMapper(core.SnakeMapper{})
	// engine.SetDefaultCacher()

	if debug {
		engine.ShowSQL(true)
		engine.Logger().SetLevel(xorm.DEFAULT_LOG_LEVEL)
	}

	// replace
	logFile := config.String("log.sqlLog")
	logFile = strings.NewReplacer(
		"{date}", helper.LocTime().Format("20060102"),
	).Replace(logFile)

	f, err := os.Create(logFile)
	if err != nil {
		clog.Fatalf("create db log file error: ", err.Error())
	}

	engine.SetLogger(xorm.NewSimpleLogger(f))
}

// Db get db connection
func Db() *xorm.Engine {
	return engine
}

// CloseEngine Close mysql engine
func CloseEngine() error {
	if enable {
		return engine.Close()
	}
	return nil
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
