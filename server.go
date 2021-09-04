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

	workoutinfo, err := createTestData()
	if err != nil {
		log.Fatal(err)
	}
	context := map[string]interface{}{
		"workouts": workoutinfo[:4],
		"startDay": workoutinfo[0].Workout.Date,
		"endDay":   workoutinfo[3].Workout.Date,
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
		log.Fatal(err)
	}

	// sets := data.GetSets(workout.id)

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

	t.ExecuteTemplate(w, "layout", workout)

}

func New(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

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

// combines an array of Lift, Set and SetQuantity to create an array of LiftInfo
// make sure Set is all from the same workout.
// this is test code for displaying data on screen.
// The variables are
func combineInfo(lifts []data.Lift, sets []data.Set, sqs []data.SetQuantity) []LiftInfo {
	setinfos := make([]SetInfo, 0)

	for idx, val := range sets {

		sq := sqs[idx]

		setinfo := SetInfo{
			Set:      val,
			Quantity: sq,
		}

		setinfos = append(setinfos, setinfo)
	}

	infoIdx := 0
	liftInfos := make([]LiftInfo, 0)
	for _, val := range lifts {
		infos := make([]SetInfo, 0)

		for infoIdx < len(setinfos) {
			info := setinfos[infoIdx]
			if info.Set.LiftId != val.Id {
				break
			}

			infos = append(infos, info)
			infoIdx++
		}

		liftInfos = append(liftInfos, LiftInfo{
			Lift:     val,
			Setinfos: infos,
		})

	}

	return liftInfos
}
