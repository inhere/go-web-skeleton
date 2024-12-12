package myrds

import (
	"fmt"
	"log"

	"github.com/gomodule/redigo/redis"
	"github.com/gookit/config/v2"
	"github.com/inhere/go-web-skeleton/app"
	"github.com/inhere/go-web-skeleton/app/clog"
	"github.com/inhere/go-web-skeleton/app/helper"
	"github.com/gookit/slog"
)

type rdsConfig struct {
	Server string
	Auth   string
	Db     int

	Disable bool
}

var (
	cfg  rdsConfig
	pool *redis.Pool
)

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
func InitRedis() (err error) {
	// 从配置文件获取redis的ip以及db
	err = config.MapStruct("redis", &cfg)
	if err != nil {
		return
	}

	if cfg.Disable {
		clog.Debugf("redis is disabled, skip init redis connection")
		return
	}

	fmt.Printf("redis - server=%s db=%d auth=%s\n", cfg.Server, cfg.Db, cfg.Auth)

	// 建立连接池
	pool = helper.NewRedisPool(cfg.Server, cfg.Auth, cfg.Db)
	return
}

// ClosePool close redis pool
func ClosePool() error {
	if cfg.Disable {
		return nil
	}

	return pool.Close()
}

// Connection return redis connection.
// Usage:
//   conn := redis.Connection()
//   defer conn.Close()
//   ... do something ...
func Connection() redis.Conn {
	slog.Info("get new redis connection from pool") // zap.Namespace("context"),
	// zap.Int("IdleCount", pool.IdleCount()),
	// zap.Int("ActiveCount", pool.ActiveCount()),

	// 记录操作日志
	if app.Debug {
		w := slog.Std()
		return redis.NewLoggingConn(pool.Get(), log.New(w, "", 0), "rds")
	}

	return pool.Get()
}

// WithConnection 公共方法，使用 collection 对象
// Usage:
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
