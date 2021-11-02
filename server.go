package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"

	"github.com/kshuta/workoutTracker/data"
)

func main() {
	mux := httprouter.New()
	mux.GET("/", Index)
	mux.GET("/workouts/:workoutId", Detail)
	mux.GET("/new", New)
	mux.POST("/new", New)

	mux.ServeFiles("/static/*filepath", http.Dir("public"))

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Fatal(server.ListenAndServe())

}

var funcs = template.FuncMap{
	"fdate": fDate,
}

// Index
// Detail
// NewWorkout
// NewLift
var isFirst bool = true
var workoutinfo []data.WorkoutInfo

func Index(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
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

	if isFirst {
		workoutinfo, err = createTestData() // for developing templates
		isFirst = false
	}
	if err != nil {
		log.Fatal(err)
	}

	lifts, err := data.GetLifts()
	if err != nil {
		log.Fatal(err)
	}

	context := map[string]interface{}{
		"workouts": workoutinfo[:4],
		"startDay": workoutinfo[0].Workout.Date,
		"endDay":   workoutinfo[3].Workout.Date,
		"lifts":    lifts,
	}

	t.ExecuteTemplate(w, "layout", context)
}

func Detail(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, err := strconv.Atoi(p.ByName("workoutId"))
	if err != nil {
		log.Fatal(err)
	}

	// load data
	workout, err := data.GetWorkout(id)
	if err != nil {
		w.WriteHeader(400)
		log.Fatalf("error in Detail Handler: %s", err)
	}

	lifts, err := data.GetWorkoutLifts(workout)
	if err != nil {
		log.Fatal(err)
	}

	liftinfos := make([]data.LiftInfo, 0)

	for _, lift := range lifts {
		setinfos, err := data.GetSetInfos(workout.Id, lift.Id)
		if err != nil {
			log.Fatal(err)
		}

		liftinfos = append(liftinfos, data.LiftInfo{
			Lift:     lift,
			Setinfos: setinfos,
		})

	}

	templ_files := []string{
		"templates/layout.html",
		"templates/detail.html",
	}
	t := template.New("layout").Funcs(funcs)
	t, err = t.ParseFiles(templ_files...)

	if err != nil {
		w.WriteHeader(400)
		log.Fatal(err)
	}

	t.ExecuteTemplate(w, "layout", liftinfos)

}

func New(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	templ_files := []string{
		"templates/layout.html",
		"templates/new.html",
	}
	t := template.New("layout").Funcs(funcs)
	t, err := t.ParseFiles(templ_files...)
	if err != nil {
		log.Fatal(err)
	}

	if r.Method != http.MethodPost {
		t.ExecuteTemplate(w, "layout", nil)
		return
	}

	// if it is a post request
	r.ParseForm()

	lift := data.Lift{
		Name: r.Form["liftName"][0],
	}
	err = lift.Create()
	if err != nil {
		log.Fatalln("error when creating lift" + err.Error())
	}

	http.Redirect(w, r, "/", http.StatusFound)
}

func fDate(t time.Time) string {
	layout := "2006/01/02"
	return t.Format(layout)
}

func calcWeight(lift data.Lift, sq data.SetQuantity) float64 {
	var finalWeight float64
	var weight int
	if sq.Ratiotype == data.Percentage {
		weight = int(lift.Max * float64(sq.PlannedRatio) / 10)
		det := weight % 100
		if det < 13 {

		} else if det < 38 {
			weight = weight - det + 25
		} else if det < 63 {
			weight = weight - det + 50
		} else if det < 88 {
			weight = weight - det + 75
		} else {
			weight = weight - det + 100
		}
	} else {
		// TODO: for when weight is rem
	}

	finalWeight = float64(weight) / 10.0
	return finalWeight
}
