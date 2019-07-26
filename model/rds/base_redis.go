package rds

import (
	"fmt"
	"strconv"

	"github.com/gomodule/redigo/redis"
	"github.com/gookit/config/v2"
	"github.com/inhere/go-web-skeleton/app"
	"go.uber.org/zap"
)

var pool *redis.Pool

// redisPrefix
const redisPrefix = "feeds:"

// GenRedisKey Gen redis key
func GenRedisKey(tpl string, keys ...interface{}) string {
	if len(keys) == 0 {
		return redisPrefix + tpl
	}
	
	return redisPrefix + fmt.Sprintf(tpl, keys...)
}

// init redis connection pool
// redigo doc https://godoc.org/github.com/gomodule/redigo/redis#pkg-examples
func init() {
	if !config.Bool("redis.enable") {
		app.Printf("redis is disabled, skip init redis connection")
		return
	}
	
	conf := config.StringMap("redis")
	
	// 从配置文件获取redis的ip以及db
	redisUrl := conf["server"]
	password := conf["auth"]
	redisDb, _ := strconv.Atoi(conf["db"])
	
	fmt.Printf("redis - server=%s db=%d auth=%s\n", redisUrl, redisDb, password)
	
	// 建立连接池
	pool = app.NewRedisPool(redisUrl, password, redisDb)
	// closePool()
}

// Connection return redis connection.
// usage:
//   conn := redis.Connection()
//   defer conn.Close()
//   ... do something ...
func Connection() redis.Conn {
	app.Logger.Info("get new redis connection from pool",
		zap.Namespace("context"),
		zap.Int("IdleCount", pool.IdleCount()),
		zap.Int("ActiveCount", pool.ActiveCount()),
	)

	// 记录操作日志
	if app.IsDebug() {
		return redis.NewLoggingConn(pool.Get(), zap.NewStdLog(app.Logger), "rds")
	}
	
	return pool.Get()
}

// WithConnection 公共方法，使用 collection 对象
// usage:
//   error = redis.WithConnection(func (c redis.Conn) error {
//       ... do something ...
//   })
func WithConnection(fn func(c redis.Conn) (interface{}, error)) (interface{}, error) {
	conn := Connection()
	defer conn.Close()
	
	return fn(conn)
}

// HasZSet
func HasZSet(key string) bool {
	count, _ := redis.Int(WithConnection(func(c redis.Conn) (interface{}, error) {
		return c.Do("zCard", key)
	}))
	
	return count > 0
}
