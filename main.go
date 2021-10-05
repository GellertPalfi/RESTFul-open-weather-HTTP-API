package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	//"github.com/go-chi/chi/v5/middleware"
)

const (
	api = "a7886d54f2fbe455ffb285d182c8e2db"
)

var (
	r = chi.NewRouter()
	city = "Budapest" //default city
	url = "https://api.openweathermap.org/data/2.5/weather?q=" + city + "&appid=" + api
)

type Health struct {
	Health string `json:"health"`
}

// weatherAll contains the whole api return body, weatherImportant only the needed parts
type weatherAll struct {
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"-"`
	Weather []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		TempMin   float64 `json:"temp_min"`
		TempMax   float64 `json:"temp_max"`
		Pressure  int     `json:"pressure"`
		Humidity  int     `json:"humidity"`
	} `json:"main"`
	Visibility int `json:"visibility"`
	Wind       struct {
		Speed float64 `json:"speed"`
		Deg   int     `json:"deg"`
		Gust  float64 `json:"gust"`
	} `json:"wind"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Dt  int `json:"dt"`
	Sys struct {
		Type    int    `json:"type"`
		ID      int    `json:"id"`
		Country string `json:"country"`
		Sunrise int64    `json:"sunrise"`
		Sunset  int64    `json:"sunset"`
	} `json:"sys"`
	Timezone int    `json:"timezone"`
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Cod      int    `json:"cod"`
}

type weatherImporant struct{
	Temp      float64 `json:"temp"`
	WindSpeed float64 `json:"wind_speed"`
	Sunrise   string  `json:"sunrise"`
	Sunset    string  `json:"sunset"`
	Pressure  int     `json:"pressure"`
	Humidity  int     `json:"humidity"`
}

func checkHealth(w http.ResponseWriter, r *http.Request){
	healthStatus := Health{Health: "healthy"}
	resp, err := http.Get(url)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	//if status code 200 is received the server health is OK
	if resp.StatusCode == 200 {
		json.NewEncoder(w).Encode(healthStatus)
	}
}

func fetchWeather(w http.ResponseWriter, r *http.Request){
	queriedCity := r.URL.Query().Get("city")
	url = "https://api.openweathermap.org/data/2.5/weather?q=" + queriedCity + "&appid=" + api
	resp, err  := http.Get(url)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	var weather weatherAll
	json.Unmarshal(body, &weather)

	weatherImp := weatherImporant{
		//2 digit flote precision, converting to celsius sometimes resulted in a lot of trailing zeros
		Temp: float64(int((weather.Main.Temp - 273.15) * 100)) /100,
		WindSpeed: weather.Wind.Speed,
		Sunrise: convertTime(weather.Sys.Sunrise),
		Sunset: convertTime(weather.Sys.Sunset),
		Pressure: weather.Main.Pressure,
		Humidity: weather.Main.Humidity,
	}
	json.NewEncoder(w).Encode(weatherImp)

}

//convert time to the needed format and only keep relevant information
func convertTime(timeToConvert int64) string{
	t :=time.Unix(timeToConvert, 0)
	return t.Format("15:02")
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