package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	workouts := getTestWorkouts()
	fmt.Fprint(w, workouts[0].Name)
}

func main() {
	fmt.Println("hello from server")

	mux := httprouter.New()
	mux.GET("/", Index)
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}
	log.Fatal(server.ListenAndServe())
}
