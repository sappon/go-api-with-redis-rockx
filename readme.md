RESTful API to call rockx market API and stores in Redis


## Run the application
*Ensure that Redis is running on port 6379*
```sh
go run .
```

or

Run `go-api-with-redis-rockx` in bin folder


## RESTFUL APIs

These are the links:
```
GET localhost:8080/
GET localhost:8080/getLatest
POST localhost:8080/updateData
```

### GET localhost:8080/

just the default to show services are ready.

### GET localhost:8080/getLatest

get the most recent record from Redis without calling the market API

### POST localhost:8080/updateData

get data from the market API with request parameters, update Redis, and return the data

example for `POST localhost:8080/updateData` request body in JSON
```
{
    "pairs":"ETH,ATOM",
    "currency":"BTC"
}
```

## Run Test

*Ensure the APIs are running by `go run .` or run bin/go-api-with-redis-rockx*

1. direct to src/ folder
2. `go test` or `go test route_test.go -v` to see more details 

## Dependencies:

```sh
go get github.com/gomodule/redigo/redis
go get -u github.com/gorilla/mux
```

## References:

create go project: https://golang.org/doc/code.html

redis with go: https://medium.com/@gilcrest_65433/basic-redis-examples-with-go-a3348a12878e, https://godoc.org/github.com/garyburd/redigo/redis

buid restful api: https://medium.com/codezillas/building-a-restful-api-with-go-part-1-9e234774b14d

consume restful api: https://www.thepolyglotdeveloper.com/2017/07/consume-restful-api-endpoints-golang-application/

test: https://medium.com/@kiransp0007/test-cases-for-rest-services-in-golang-beb6699c75a0
