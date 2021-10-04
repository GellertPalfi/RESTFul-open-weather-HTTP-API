package main

import (
	"encoding/json"
	"fmt"
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

func checkHealth(w http.ResponseWriter, r *http.Request){
	healtStatus := Health{Health: "healthy"}
	resp, err := http.Get("https://" + url)

	fmt.Println(url)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	fmt.Println("checkhealth endpoint hit")
	json.NewEncoder(w).Encode(healtStatus)
}

func handleRequests(){
	port := ":3000"

	r.HandleFunc("/api/health", checkHealth)
	fmt.Println("Server is up and running on port " + port)
	log.Fatal(http.ListenAndServe(port, r))
}

func main() {

	r.Get("/api/wesather", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("weather"))
	})

	handleRequests()
}