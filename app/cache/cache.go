package cache

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/inhere/go-web-skeleton/app/utils"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"time"
)

var debug bool
var cachePrefix string
var pool *redis.Pool
var logger *zap.Logger

// init cache
// ref package: github.com/astaxie/beego/cache/redis
// redigo doc https://godoc.org/github.com/gomodule/redigo/redis#pkg-examples
func Init(pl *redis.Pool, prefix string, l *zap.Logger, d bool) {
	// 建立连接池
	pool = pl
	debug = d
	logger = l
	cachePrefix = prefix
}

// GenKey
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
		return utils.JsonDecode(ret.([]byte), v)
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
	jsonBytes, _ := utils.JsonEncode(data)

	_, err := exec("setEx", key, int64(ttl), jsonBytes)

	return err
}

// Delete cache
func Delete(key string) error {
	_, err := exec("del", key)

	if err != nil {
		logger.Error("redis error: " + err.Error())
	}

	return err
}

// Has cache key
func Has(key string) bool {
	// 0 OR 1
	one, err := redis.Int(exec("exists", key))

	if err != nil {
		logger.Error("redis error: " + err.Error())
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

		logger.Debug("operate redis cache: "+commandName,
			zap.Namespace("context"),
			zap.String("cache_key", fullKey),
			zap.String("elapsed_time", utils.CalcElapsedTime(st)),
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
	logger.Info("get new redis connection from pool",
		zap.Namespace("context"),
		zap.Int("IdleCount", pool.IdleCount()),
		zap.Int("ActiveCount", pool.ActiveCount()),
	)

	// 记录操作日志
	if debug {
		return redis.NewLoggingConn(pool.Get(), zap.NewStdLog(logger), "rds")
	}

	return pool.Get()
}
