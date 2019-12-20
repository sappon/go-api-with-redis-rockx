package main

import (
	"time"
	"encoding/json"
	"flag"

	"github.com/gomodule/redigo/redis"
)

type Object struct {
	Data 		ReturnObj
	CreatedAt 	time.Time
}

var (
  pool *redis.Pool
  redisServer = flag.String("redisServer", ":6379", "")
)

func newPool(addr string) *redis.Pool {
  return &redis.Pool{
    MaxIdle: 3,
    IdleTimeout: 240 * time.Second,
    Dial: func () (redis.Conn, error) { return redis.Dial("tcp", addr) },
  }
}

// set executes the redis SET command
func set(key string, data ReturnObj) error {

	// Connect to Redis (DB)
	conn := pool.Get()
	defer conn.Close()  // close the connection when the function completes

	obj := Object{}
	obj.Data = data
	obj.CreatedAt = time.Now().Local()

	json, json_err := json.Marshal(obj)
	if json_err != nil {
		return json_err
	}
    _, err := conn.Do("SET", key, json)
    return err
}

// get executes the redis GET command
func get(key string) ([]byte, error) {

	// Connect to Redis (DB)
	conn := pool.Get()
	defer conn.Close()  // close the connection when the function completes

	// Simple GET example with String helper
	s, err := redis.String(conn.Do("GET", key))
	return []byte(s), err
}