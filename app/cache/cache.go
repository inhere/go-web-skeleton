package cache

import (
	"fmt"
	"log"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/gookit/goutil/calc"
	"github.com/gookit/goutil/jsonutil"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

var debug bool
var cachePrefix string
var pool *redis.Pool

// init cache
// ref package: github.com/astaxie/beego/cache/redis
// redigo doc https://godoc.org/github.com/gomodule/redigo/redis#pkg-examples
func Init(pl *redis.Pool, prefix string, d bool) {
	// 建立连接池
	pool = pl
	debug = d
	cachePrefix = prefix
}

// GenKey gen cache key
func GenKey(tpl string, keys ...interface{}) string {
	// return cachePrefix + fmt.Sprintf(tpl, keys...)
	// 初始化缓存时已经设置了前缀了
	return fmt.Sprintf(tpl, keys...)
}

// Get cache and map to a struct
// usage:
// 	cache.GetAndMapTo("key", &User{})
func GetAndMapTo(key string, v interface{}) (err error) {
	var ret interface{}

	ret, err = exec("get", key)
	if err == nil {
		// data must convert to byte
		return jsonutil.Decode(ret.([]byte), v)
	}

	return
}

// Get cache from redis.
func Get(key string) interface{} {
	if v, err := exec("get", key); err == nil {
		return v
	}

	return nil
}

// Set cache
func Set(key string, data interface{}, ttl int) error {
	jsonBytes, _ := jsonutil.Encode(data)

	_, err := exec("setEx", key, int64(ttl), jsonBytes)

	return err
}

// Delete cache
func Delete(key string) error {
	_, err := exec("del", key)
	if err != nil {
		logrus.Error("redis error: ", err.Error())
	}

	return err
}

// Has cache key
func Has(key string) bool {
	// 0 OR 1
	one, err := redis.Int(exec("exists", key))
	if err != nil {
		logrus.Error("redis error: ", err.Error())
	}

	return one == 1
}

// actually do the redis cmds, args[0] must be the key name.
func exec(commandName string, args ...interface{}) (reply interface{}, err error) {
	if len(args) < 1 {
		return nil, errors.New("missing required arguments")
	}

	var fullKey string
	if cachePrefix != "" {
		fullKey = fmt.Sprintf("%s:%s", cachePrefix, args[0])
	} else {

	}

	args[0] = fullKey

	if debug {
		st := time.Now()
		c := Connection()
		defer c.Close()
		reply, err = c.Do(commandName, args...)

		logrus.Debug(
			"operate redis cache: ", commandName,
			"cache_key", fullKey,
			"elapsed_time", calc.ElapsedTime(st),
		)

		return
	}

	c := Connection()
	defer c.Close()
	return c.Do(commandName, args...)
}

// Connection return redis connection.
// usage:
//   conn := redis.Connection()
//   defer conn.Close()
//   ... do something ...
func Connection() redis.Conn {
	logrus.Info("get new redis connection from pool",
		// zap.Namespace("context"),
		// zap.Int("IdleCount", pool.IdleCount()),
		// zap.Int("ActiveCount", pool.ActiveCount()),
	)

	// 记录操作日志
	if debug {
		w := logrus.StandardLogger().Writer()
		return redis.NewLoggingConn(pool.Get(), log.New(w, "", 0), "rds")
	}

	return pool.Get()
}
