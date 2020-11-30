package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

// Weather -- This is a sample wather interface
type Weather struct {
	Title        string `json:"title"`
	LocationType string `json:"location_type"`
	Woeid        int    `json:"woeid"`
	LallLong     string `json:"latt_long"`
}

func (we *Weather) Write(w io.Writer) {
	b, _ := json.Marshal(*we)
	w.Write(b)
}

type Test []Weather

func weatherHandler(res http.ResponseWriter, req *http.Request) {

	// data := []Weather{}
	var data []Weather
	res.Header().Set("Content-Type", "application/json")

	url := "https://www.metaweather.com/api/location/search/?query=san"

	// Building the request
	request, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}

	client := &http.Client{}

	response, err := client.Do(request)

	if err != nil {
		log.Fatal("Do: ", err)
		return
	}
	body, readErr := ioutil.ReadAll(response.Body)
	if readErr != nil {
		log.Fatal("Read Error: ", readErr)
		return
	}
	defer response.Body.Close()
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println(err)
	}
	c, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	res.Write(c)
}

func weather() {
	http.HandleFunc("/", weatherHandler)
	fmt.Println("Server is starting..")
	http.ListenAndServe(":5000", nil)
}
