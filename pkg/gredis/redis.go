package gredis

import (
	"encoding/json"
	"ginDemo/pkg/logging"
	"ginDemo/pkg/setting"
	"github.com/gomodule/redigo/redis"
	"time"
)

var RedisConn *redis.Pool

func Setup() error {
	//	也就是创建RedisConn
	RedisConn = &redis.Pool{
		MaxIdle:         setting.RedisSetting.MaxIdle,
		MaxConnLifetime: time.Duration(setting.RedisSetting.MaxActive),
		IdleTimeout:     time.Duration(setting.RedisSetting.IdleTimeout),
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", setting.RedisSetting.Host)
			if err != nil {
				return nil, err
			}
			if setting.RedisSetting.Password != "" {
				if _, err := c.Do("AUTH", setting.RedisSetting.Password); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
	return nil

}

func Set(key string, data interface{}, time int) error {
	conn := RedisConn.Get()
	defer conn.Close()

	jsonData, err := json.Marshal(data)
	if err != nil {
		logging.Error("Redis.Set marshal data failed")
		return err
	}
	// 硬编码了
	_, err = conn.Do("SET", key, jsonData)
	if err != nil {
		logging.Error("Redis.Set SET data failed")
		return err
	}
	_, err = conn.Do("EXPIRE", key, time)
	if err != nil {
		logging.Error("Redis.Set EXPIRE data failed")
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
