package main

import (
	"time"
	"github.com/gomodule/redigo/redis"
	"encoding/json"
)

type Object struct {
	Data 		ReturnObj
	CreatedAt 	time.Time
}

func redisConn() redis.Conn {
	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		panic(err.Error())
	}
	return c
}

// set executes the redis SET command
func set(key string, data ReturnObj) error {

	// Connect to Redis (DB)
	conn := redisConn()
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
	conn := redisConn()
	defer conn.Close()  // close the connection when the function completes

	// Simple GET example with String helper
	s, err := redis.String(conn.Do("GET", key))
	return []byte(s), err
}