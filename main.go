package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	//"github.com/go-chi/chi/v5/middleware"
)


func main() {
	//apikey = "a7886d54f2fbe455ffb285d182c8e2db"
	//city := "Budapest"
	port := ":3000"
	r := chi.NewRouter()
	//toDo dont hardcode api key
	//url := "api.openweathermap.org/data/2.5/weather?q=" + city + "&appid=a7886d54f2fbe455ffb285d182c8e2db"

	// Get api health
	r.Get("/api/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("healt"))
	})

	r.Get("/api/weather", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("weather"))
	})
	//toDo write something with actual information
	fmt.Println("somethingplaceholer")
	http.ListenAndServe(port, r)
}