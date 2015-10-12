package db

import (
	"github.com/garyburd/redigo/redis"
)

var (
	Redis_conn redis.Conn
)

//
func RegisterRedis(redis_uri string) (err error) {
	Redis_conn, err = redis.Dial("tcp", redis_uri)
}

//
func IsExistKey(key string) bool {

}

//
func GetValueByKey(key string) bool {

}

//
func SetValueByKey(key string, value interface{}) {

}
