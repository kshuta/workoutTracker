package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	mux := httprouter.New()
	mux.GET("/", Index)

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Fatal(server.ListenAndServe())

}

// Index
// Detail
// NewWorkout
// NewLift

func Index(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	t := template.New("index.html")
	t, err := t.ParseFiles("templates/index.html")
	if err != nil {
		// change to redirect to error page
		w.WriteHeader(400)
		log.Fatal(err)
	}

	t.Execute(w, nil)
}
