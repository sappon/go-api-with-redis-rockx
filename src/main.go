package main

import (
    "log"
    "net/http"
	"fmt"
	"io/ioutil"

	"github.com/gomodule/redigo/redis"
	"github.com/gorilla/mux"
)


func test(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
    w.Header().Set("Content-Type", "application/json")
    w.Write([]byte(`{"message": "hello world"}`))
}

func main() {

	// Connect to Redis (DB)
	pool := newPool()
	conn := pool.Get()
	defer conn.Close()  // close the connection when the function completes
	
	// call Redis PING command to test connectivity
	err := ping(conn)
	if err != nil {
		fmt.Println(err)
	}

	//Step 2: Get data from Market API
	response, err := http.Get("http://market.v1.rockx.com/v1/markets/?pair=ETH&base_currency=CNY")
	if err != nil {
        fmt.Printf("The HTTP request failed with error %s\n", err)
    } else {
        data, _ := ioutil.ReadAll(response.Body)
        fmt.Println(string(data))
    }

	// init router
	router := mux.NewRouter()
	// endpoints
	router.HandleFunc("/fetchData", test).Methods("GET")

    log.Fatal(http.ListenAndServe(":8080", router))

}

func newPool() *redis.Pool {
	return &redis.Pool{
		// Maximum number of idle connections in the pool.
		MaxIdle: 80,
		// max number of connections
		MaxActive: 12000,
		// Dial is an application supplied function for creating and
		// configuring a connection.
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", ":6379")
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}
}

// ping tests connectivity for redis (PONG should be returned)
func ping(c redis.Conn) error {
	// Send PING command to Redis
	// PING command returns a Redis "Simple String"
	// Use redis.String to convert the interface type to string
	s, err := redis.String(c.Do("PING"))
	if err != nil {
		return err
	}

	fmt.Printf("PING Response = %s\n", s)
	// Output: PONG

	return nil
}