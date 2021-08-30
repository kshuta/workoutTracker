package main

import (
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/kshuta/workoutTracker/data"
)

func main() {
	mux := httprouter.New()
	mux.GET("/", Index)

	mux.ServeFiles("/static/*filepath", http.Dir("public"))

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
	funcs := template.FuncMap{"fdate": fDate}
	t := template.New("layout").Funcs(funcs)

	templ_files := []string{
		"templates/layout.html",
		"templates/index.html",
	}

	t, err := t.ParseFiles(templ_files...)
	if err != nil {
		// change to redirect to error page
		w.WriteHeader(400)
		log.Fatal(err)
	}

	lifts := getTestLifts()
	sets := getTestSets()
	sqs := getTestSetQuantity()

	setinfos := combineInfo(sets, sqs)

	t.ExecuteTemplate(w, "layout", lifts)
}

func fDate(t time.Time) string {
	layout := "2006-01-02"
	return t.Format(layout)
}

type SetInfo struct {
	Set       data.Set
	Quantites []data.SetQuantity
}

func combineInfo(sets []data.Set, sqs []data.SetQuantity) []SetInfo {
	sqIdx := 0
	setinfos := make([]SetInfo, len(sets))

	for idx, val := range sets {
		sqset := make([]data.SetQuantity, 0)
		for {

			sq := sqs[sqIdx]
			if sq.SetId != val.Id {
				break
			}

			sqset = append(sqset, sq)
		}

		setinfo := SetInfo{
			Set:       val,
			Quantites: sqset,
		}
		setinfos[idx] = setinfo
	}

	return setinfos
}
