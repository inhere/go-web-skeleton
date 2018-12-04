package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"github.com/inhere/go-webx/app"
	"log"
	"os"
	"strconv"
	"strings"
)

var engine *xorm.Engine

func init() {
	var err error

	db, _ := app.Cfg.StringMap("db")
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8",
		db["user"], db["password"], db["host"], db["port"], db["name"],
	)

	fmt.Printf("mysql - %s\n", dsn)

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

	if app.Debug {
		engine.ShowSQL(true)
		engine.Logger().SetLevel(core.LOG_DEBUG)
	}

	// replace
	logFile, _ := app.Cfg.Get("log.sqlLog")
	logFile = strings.NewReplacer(
		"{date}", app.LocTime().Format("20060102"),
		"{hostname}", app.Hostname,
	).Replace(logFile)

	f, err := os.Create(logFile)
	if err != nil {
		log.Fatal(err.Error())
	}

	engine.SetLogger(xorm.NewSimpleLogger(f))
}

func Db() *xorm.Engine {
	return engine
}

// UpdateById
// usage:
// user := new(User)
// num, err := mysql.UpdateById(23, user, "name", "email")
func UpdateById(id int64, model interface{}, fields ...string) (affected int64, err error) {
	affected, err = engine.ID(id).Cols(fields...).Update(model)

	return
}

// DeleteById
// usage:
// user := new(User)
// num, err := mysql.DeleteById(23, user)
func DeleteById(id int64, model interface{}) (affected int64, err error) {
	affected, err = engine.ID(id).Delete(model)

	return
}
