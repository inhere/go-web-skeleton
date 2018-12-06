package mongo

import (
	"errors"
	"fmt"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/inhere/go-wex-skeleton/app"
	"log"
	"reflect"
	"strings"
)

// Collection mongodb collection interface
type Collection interface {
	CollectionName() string
}

// DebugLogger for mongodb
type DebugLogger struct {
}

var (
	connection *mgo.Session

	auth, servers, mgoUri, database string
)

var (
	invalidObjectId = errors.New("mongo: must provide an valid document Id")
)

func init() {
	fmt.Printf("mongo - %s db=%s\n", servers, database)

	if app.Debug {
		// 设为 true 数据打印太多了
		mgo.SetDebug(false)
		// mgo.SetLogger(DebugLogger{})
	}

	// get config
	auth = app.Cfg.MustString("mgo.auth")
	mgoUri = app.Cfg.MustString("mgo.uri")
	servers = app.Cfg.MustString("mgo.servers")
	database = app.Cfg.MustString("mgo.database")

	// create connection
	createConnection()
}

func (d DebugLogger) Output(calldepth int, s string) error {
	log.Print("mongo: ", s, "\n")
	return nil
}

// Create connection
func createConnection() {
	var err error
	connection, err = mgo.Dial(auth + "@" + servers + mgoUri)

	if err != nil {
		panic(err) // 直接终止程序运行
	}

	// Optional. Switch the session to a monotonic behavior.
	// connection.SetMode(mgo.Monotonic, true)
	// 最大连接池默认为 4096
	connection.SetPoolLimit(1024)
}

// Connection return mongodb connection.
// usage:
//   conn := mongo.Connection()
//   defer conn.Close()
//   ... do something ...
func Connection() *mgo.Session {
	if connection == nil {
		createConnection()
	}

	return connection.Clone()
}

// WithCollection 公共方法，使用 collection 对象
// usage:
//   error = mongo.WithCollection("name", func (conn *mgo.Collection) error {
//       ... do something ...
//   })
func WithCollection(collection string, s func(*mgo.Collection) error) error {
	conn := Connection()
	defer conn.Close()
	c := conn.DB(database).C(collection)

	return s(c)
}

/**
========================= some command functions =========================
*/

// FindById Finding a record by primary key ID
// usage:
// m := &mongo.Moment{}  // NOTICE: please use ref
// mongo.FindById("collection name", "id", m)
func FindById(cName string, id string, model interface{}, fields string) (code int, err error) {
	if len(id) < 10 || !bson.IsObjectIdHex(id) {
		return app.ErrInvalidParam, invalidObjectId
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
			code = app.ErrNotFound
		} else {
			code = app.ErrDatabase
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
			code = app.ErrNotFound
		} else {
			code = app.ErrDatabase
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
		return app.ErrInvalidParam, invalidObjectId
	}

	conn := Connection()
	defer conn.Close()
	c := conn.DB(database).C(cName)

	oid := bson.ObjectIdHex(id)
	err = c.Update(
		bson.M{"_id": oid},
		// bson.M{"$set": bson.M{"password": newHash}, "$currentDate": bson.M{"lastModified": true}},
		bson.M{"$set": change},
	)

	if err != nil {
		if err == mgo.ErrNotFound {
			code = app.ErrNotFound
		} else {
			code = app.ErrUpdateFail
		}
	}

	return
}

// UpdateSome
func UpdateBy(cName string, query bson.M, change bson.M) (code int, err error) {
	conn := Connection()
	defer conn.Close()
	c := conn.DB(database).C(cName)

	err = c.Update(
		// bson.M{"_id": id, "password": oldHash},
		query,
		// bson.M{"$set": bson.M{"password": newHash, "salt": newSalt}},
		bson.M{"$set": change},
	)

	if err != nil {
		if err == mgo.ErrNotFound {
			return app.ErrNotFound, err
		}

		return app.ErrUpdateFail, err
	}

	return
}

// DeleteById Delete a record by primary key ID
func DeleteById(cName string, id string) (code int, err error) {
	if !bson.IsObjectIdHex(id) {
		return app.ErrInvalidParam, invalidObjectId
	}

	// do delete
	err = WithCollection(cName, func(c *mgo.Collection) error {
		return c.RemoveId(bson.ObjectIdHex(id))
	})

	if err != nil {
		if err == mgo.ErrNotFound {
			code = app.ErrNotFound
		} else {
			code = app.ErrDeleteFail
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
