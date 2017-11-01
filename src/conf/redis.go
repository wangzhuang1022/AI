package conf

import (
	"github.com/garyburd/redigo/redis"
)

func SetRedis(key string, value string) string{
	conn := Pool.Get()
	defer conn.Close()
	reply, _ := redis.String(conn.Do("set", key, value))
	return reply
}
func GetRedis(key string) string {
	conn := Pool.Get()
	defer conn.Close()
	value, _ := redis.String(conn.Do("get", key))
	return value
}
