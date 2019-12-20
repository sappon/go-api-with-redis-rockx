package main

import (
	"net/http"
	"fmt"
	"encoding/json"
	"io/ioutil"
)

type Param struct {
	Pairs 		string `json:"pairs"`
	Currency 	string `json:"currency"`
}

type Route struct {
	Name        string
	Method      string
	Pattern     string
	Handle	    func(http.ResponseWriter, *http.Request)
}

type Routes []Route

var routes = Routes{
	Route{
		"index",
		"GET",
		"/",
		index,
	},

	Route{
		"get latest data",
		"GET",
		"/getLatest",
		getLatest,
	},

	Route{
		"update data",
		"POST",
		"/updateData",
		updateData,
	},
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "The services are ready!")
}

func getLatest(w http.ResponseWriter, r *http.Request){

	// step 1: update redis from market API
	//updateDbFromMarket()

	// step 2: get from redis
	key := "Latest"
    data, err := get(key)

	// step 3: create response
    w.Header().Set("Content-Type", "application/json")
    if err != nil {
    	w.Write([]byte(`{"message": "`+err.Error()+`"}`))
		fmt.Println(err)
	}
    w.Write(data)
}

func updateData(w http.ResponseWriter, r *http.Request){

	// step 1: update redis from market API
	var p Param
	body, err := ioutil.ReadAll(r.Body)
	pErr := json.Unmarshal(body, &p)
	if pErr != nil {
		fmt.Println(pErr)
	}
	updateDbFromMarket(p.Pairs, p.Currency)

    // step 2: get from redis
	key := "Latest"
    data, err := get(key)

	// step 3: create response
    w.Header().Set("Content-Type", "application/json")
    if err != nil {
    	w.Write([]byte(`{"message": "`+err.Error()+`"}`))
		fmt.Println(err)
	}
    w.Write(data)

}