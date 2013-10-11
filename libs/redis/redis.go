package redis

import (
	"github.com/marconi/devfeed/core"
	rds "github.com/garyburd/redigo/redis"
)

type RedisClient struct {
	rds.PubSubConn
}

func NewRedisClient() *RedisClient {
	conn, err := rds.Dial("tcp", core.Config.Redis.Url)
	if err != nil {
		panic("error connecting to redis: " + err.Error())
	}
	return &RedisClient{rds.PubSubConn{conn}}
}
