package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	//"github.com/go-chi/chi/v5/middleware"
)

const (
	api = "a7886d54f2fbe455ffb285d182c8e2db"
)

var (
	r = chi.NewRouter()
	url = "api.openweathermap.org/data/2.5/weather?q=London&appid=" + api
)

type Health struct {
	Health string `json:"health"`
}

type Weather struct {
	Temp      float32 `json:"temp"`
	WindSpeed float32 `json:"wind_speed"`
	Sunrise   string `json:"sunrise"`
	Sunset    string `json:"sunset"`
	Pressure  int `json:"pressure"`
	Humidity  int `json:"humidity"`
}

func checkHealth(w http.ResponseWriter, r *http.Request){
	healthStatus := Health{Health: "healthy"}
	resp, err := http.Get("https://" + url)

	//fmt.Println(url)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		fmt.Println("checkhealth endpoint hit")
		json.NewEncoder(w).Encode(healthStatus)
	}

}

func fetchWeather(w http.ResponseWriter, r *http.Request){
	//toDo make the api call only once
	resp, err  := http.Get("https://" + url)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//weather := Weather{resp}
	body, err := ioutil.ReadAll(resp.Body)
	w.Write(body)
	fmt.Println(string(body))

}

func handleRequests(){
	port := ":3000"

	r.HandleFunc("/api/health", checkHealth)
	r.HandleFunc("/api/weather", fetchWeather)
	fmt.Println("Server is up and running on port " + port)
	log.Fatal(http.ListenAndServe(port, r))
}

func main() {
	handleRequests()
}