package main

import (
	"net/http"
	"fmt"
	"encoding/json"
	"io/ioutil"
)

type ReturnObj struct {
	Success 		string 	`json:"success"`
	Data 			Data 	`json:"data"`
	Message 		string 	`json:"message"`
}

type Data struct {
	Default_currency 	string 	`json:"default_currency"`
	Price 				[]Price `json:"price"`
}

type Price struct {
    Name 			string 			`json:"name"`
    Trading_price 	[]Trading_price `json:"trading_price"`
}

type Trading_price struct {
    Name       				string `json:"name"`
    Average_buy_price   	string `json:"average_buy_price"`
    Average_sell_price  	string `json:"average_sell_price"`
    Average_market_price  	string `json:"average_market_price"`
}

func updateDbFromMarket(p string, cur string) {

	//fetch data
	url := "http://market.v1.rockx.com/v1/markets/?pairs="+p+"&base_currency="+cur
	response, err := http.Get(url)
	defer response.Body.Close()

	//store
	if err != nil {
        fmt.Printf("The HTTP request failed with error %s\n", err)
    } else {
    	var data ReturnObj
    	body, readErr := ioutil.ReadAll(response.Body)
    	if readErr != nil {
			fmt.Println(readErr)
		}
		jsonErr := json.Unmarshal(body, &data)
		if jsonErr != nil {
			fmt.Println(jsonErr)
		}
    	err := set("Latest", data)
		if err != nil {
			fmt.Println(err)
		}
    }

}