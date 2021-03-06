package gredis

import (
	"encoding/json"
	"github.com/gomodule/redigo/redis"
	"github.com/stone955/my-gin-blog/pkg/setting"
	"log"
	"time"
)

var RedisConn *redis.Pool

func Setup() {
	RedisConn = &redis.Pool{
		Dial: func() (conn redis.Conn, err error) {
			c, err := redis.Dial("tcp", setting.RedisCfg.Host)
			if err != nil {
				log.Fatalf("gredis.Setup.Dial err: %v\n", err)
				return nil, err
			}
			if setting.RedisCfg.Password != "" {
				if _, err := c.Do("AUTH", setting.RedisCfg.Password); err != nil {
					_ = c.Close()
					log.Fatalf("gredis.Setup.AUTH err: %v\n", err)
					return nil, err
				}
			}
			return c, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
		MaxIdle:     setting.RedisCfg.MaxIdle,
		MaxActive:   setting.RedisCfg.MaxActive,
		IdleTimeout: setting.RedisCfg.IdleTimeout,
	}
}

func Close() {
	_ = RedisConn.Close()
}

func Set(key string, data interface{}, time int) error {
	conn := RedisConn.Get()
	defer conn.Close()

	value, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = conn.Do("SET", key, value)
	if err != nil {
		return err
	}

	_, err = conn.Do("EXPIRE", key, time)
	if err != nil {
		return err
	}

	return nil
}

func Exists(key string) bool {
	conn := RedisConn.Get()
	defer conn.Close()

	exists, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return false
	}

	return exists
}

func Get(key string) ([]byte, error) {
	conn := RedisConn.Get()
	defer conn.Close()

	reply, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return nil, err
	}

	return reply, nil
}

func Delete(key string) (bool, error) {
	conn := RedisConn.Get()
	defer conn.Close()

	return redis.Bool(conn.Do("DEL", key))
}

func LikeDeletes(key string) error {
	conn := RedisConn.Get()
	defer conn.Close()

	keys, err := redis.Strings(conn.Do("KEYS", "*"+key+"*"))
	if err != nil {
		return err
	}

	for _, key := range keys {
		_, err = Delete(key)
		if err != nil {
			return err
		}
	}

	return nil
}
