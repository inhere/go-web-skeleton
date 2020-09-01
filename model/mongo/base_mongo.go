package mongo

import (
	"errors"
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/gookit/config/v2"
	"github.com/inhere/go-web-skeleton/app"
	"github.com/inhere/go-web-skeleton/app/clog"
	"github.com/inhere/go-web-skeleton/app/errcode"
)

// Collection mongodb collection interface
type Collection interface {
	CollectionName() string
}

// DebugLogger for mongodb
type DebugLogger struct {
}

type mgoConfig struct {
	Auth string
	Uri string
	Servers string
	Database string

	Disable bool
}

var (
	cfg mgoConfig
	session *mgo.Session
)

var (
	invalidObjectId = errors.New("mongo: must provide an valid document Id")
)

func InitMongo() (err error) {
	// get config
	err = config.MapStruct("mgo", &cfg)
	if err != nil {
		return
	}

	if cfg.Disable {
		clog.Debugf("mongo is disabled, skip init mongo connection")
		return
	}

	clog.Printf("mongo - %s db=%s\n", cfg.Servers, cfg.Database)

	if app.Debug {
		// 设为 true 数据打印太多了
		mgo.SetDebug(false)
		// mgo.SetLogger(DebugLogger{})
	}

	// create session
	return createSession()
}

func (d DebugLogger) Output(callDepth int, s string) error {
	log.Print("mongo: ", s, "\n")
	return nil
}

// Create session
func createSession() (err error) {
	session, err = mgo.Dial(cfg.Auth + "@" + cfg.Servers + cfg.Uri)
	if err != nil {
		return
	}

	// Optional. Switch the session to a monotonic behavior.
	// session.SetMode(mgo.Monotonic, true)
	// 最大连接池默认为 4096
	session.SetPoolLimit(1024)
	return
}

// Connection return new mongodb connection.
// Usage:
//   conn := mongo.Connection()
//   defer conn.Close()
//   ... do something ...
func Connection() *mgo.Session {
	return session.Clone()
}

// WithCollection 公共方法，使用 collection 对象
// Usage:
//   error = mongo.WithCollection("name", func (conn *mgo.Collection) error {
//       ... do something ...
//   })
func WithCollection(collection string, s func(*mgo.Collection) error) error {
	conn := Connection()
	defer conn.Close()

	c := conn.DB(cfg.Database).C(collection)
	return s(c)
}

// CloseSession close mgo connection session
func CloseSession() {
	if !cfg.Disable {
		session.Close()
	}
}

/**
========================= some command functions =========================
*/

// FindById Finding a record by primary key ID
// Usage:
// m := &mongo.Moment{}  // NOTICE: please use ref
// mongo.FindById("collection name", "id", m)
func FindById(cName string, id string, model interface{}, fields string) (code int, err error) {
	if len(id) < 10 || !bson.IsObjectIdHex(id) {
		return errcode.ErrParam, invalidObjectId
	}

	// "col1,col2" => bson.M{"col1": 1,"col1": 2}
	fm := fieldsString2BsonM(fields)

	// do query
	err = WithCollection(cName, func(c *mgo.Collection) error {
		oid := bson.ObjectIdHex(id)

		return c.FindId(oid).Select(fm).One(model)
	})

	if err != nil {
		if err == mgo.ErrNotFound {
			code = errcode.ErrNotFound
		} else {
			code = errcode.ErrDatabase
		}
	}

	return
}

// FindOne
func FindOne(cName string, query bson.M, model interface{}, fields string) (code int, err error) {
	// "col1,col2" => bson.M{"col1": 1,"col1": 2}
	fm := fieldsString2BsonM(fields)

	// do query
	err = WithCollection(cName, func(c *mgo.Collection) error {
		return c.Find(query).Select(fm).One(model)
	})

	if err != nil {
		if err == mgo.ErrNotFound {
			code = errcode.ErrNotFound
		} else {
			code = errcode.ErrDatabase
		}
	}

	return
}

// FindAllByPage Execute the query, Paginate to get data
func FindAllByPage(cName string, query bson.M, sort string, fields string, page int, limit int, r interface{}) (err error) {
	if page < 1 {
		page = 1
	}

	if limit < 1 {
		limit = 1
	}

	// "col1,col2" => bson.M{"col1": 1,"col1": 2}
	fm := fieldsString2BsonM(fields)
	af := func(c *mgo.Collection) error {
		skip := (page - 1) * limit

		return c.Find(query).Sort(sort).Select(fm).Skip(skip).Limit(limit).All(r)
	}

	err = WithCollection(cName, af)

	return
}

// UpdateById
func UpdateById(cName string, id string, change bson.M) (code int, err error) {
	if !bson.IsObjectIdHex(id) {
		return errcode.ErrParam, invalidObjectId
	}

	conn := Connection()
	defer conn.Close()
	c := conn.DB(cfg.Database).C(cName)

	oid := bson.ObjectIdHex(id)
	err = c.Update(
		bson.M{"_id": oid},
		// bson.M{"$set": bson.M{"password": newHash}, "$currentDate": bson.M{"lastModified": true}},
		bson.M{"$set": change},
	)

	if err != nil {
		if err == mgo.ErrNotFound {
			code = errcode.ErrNotFound
		} else {
			code = errcode.ErrUpdateFail
		}
	}

	return
}

// UpdateSome
func UpdateBy(cName string, query bson.M, change bson.M) (code int, err error) {
	conn := Connection()
	defer conn.Close()
	c := conn.DB(cfg.Database).C(cName)

	err = c.Update(
		// bson.M{"_id": id, "password": oldHash},
		query,
		// bson.M{"$set": bson.M{"password": newHash, "salt": newSalt}},
		bson.M{"$set": change},
	)

	if err != nil {
		if err == mgo.ErrNotFound {
			return errcode.ErrNotFound, err
		}

		return errcode.ErrUpdateFail, err
	}

	return
}

// DeleteById Delete a record by primary key ID
func DeleteById(cName string, id string) (code int, err error) {
	if !bson.IsObjectIdHex(id) {
		return errcode.ErrParam, invalidObjectId
	}

	// do delete
	err = WithCollection(cName, func(c *mgo.Collection) error {
		return c.RemoveId(bson.ObjectIdHex(id))
	})

	if err != nil {
		if err == mgo.ErrNotFound {
			code = errcode.ErrNotFound
		} else {
			code = errcode.ErrDeleteFail
		}
	}

	return
}

// TransMapToBsonM
func TransMap2BsonM(mp map[string]interface{}) bson.M {
	// 类型转换 mp 和 bson.M 的本质类型是一样的，所以可以直接这样用
	return bson.M(mp)
}

// TransMapToBsonM
func TransList2BsonM(ls []string) (bm bson.M) {
	bm = bson.M{}
	for _, v := range ls {
		bm[v] = 1
	}

	return
}

// fieldsString2BsonM trans "col1,col2" to bson.M{"col1": 1,"col1": 1}
// fields string eg "col1,col2,...."
func fieldsString2BsonM(fields string) bson.M {
	var bm = bson.M{}

	if len(fields) == 0 || fields == "*" {
		return bm
	}

	fl := strings.Split(fields, ",")

	for _, n := range fl {
		// skip empty
		n = strings.Trim(n, " ")
		if n == "" {
			continue
		}

		// auto trans
		if n == "id" {
			n = "_id"
		}

		bm[n] = 1
	}

	return bm
}

// TransStruct2BsonM translate structure to map
func TransStruct2BsonM(obj interface{}) bson.M {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	bm := bson.M{}

	for i := 0; i < t.NumField(); i++ {
		fmt.Sprintf("%d %+v\n", i, t.Field(i))
		bm[t.Field(i).Name] = v.Field(i).Interface()
	}

	return bm
}
