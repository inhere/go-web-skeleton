package app

import (
	"github.com/garyburd/redigo/redis"
	"time"
)

// create new redis pool
func NewRedisPool(url, password string, redisDb int) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     100,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", url)
			if err != nil {
				return nil, err
			}

			if password != "" {
				_, err := c.Do("AUTH", password)
				if err != nil {
					c.Close()
					return nil, err
				}
			}
			c.Do("SELECT", redisDb)

			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

// @see https://git-books.github.io/books/go-web-programme/?p=05.6.md
func CloseRedisPool(pl *redis.Pool) {
	pl.Close()
}

// @see https://git-books.github.io/books/go-web-programme/?p=05.6.md
// func closePool() {
//    c := make(chan os.Signal, 1)
//    signal.Notify(c, os.Interrupt)
//    signal.Notify(c, syscall.SIGTERM)
//    signal.Notify(c, syscall.SIGKILL)
//    go func() {
//        <-c
//        Pool.Close()
//        os.Exit(0)
//    }()
// }
