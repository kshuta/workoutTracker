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
	funcs := template.FuncMap{"fdate": fDate, "calcWeight": calcWeight}
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

	liftinfos := combineInfo(lifts, sets, sqs)
	setWeight(liftinfos)

	workoutinfo, err := createTestData()
	if err != nil {
		log.Fatal(err)
	}

	log.Print(workoutinfo)

	t.ExecuteTemplate(w, "layout", workoutinfo)
}

func fDate(t time.Time) string {
	layout := "2006-01-02"
	return t.Format(layout)
}

func setWeight(liftinfos []LiftInfo) {
	for liftIdx, lift := range liftinfos {
		for setIdx, setinfo := range lift.Setinfos {
			weight := calcWeight(lift.Lift, setinfo.Quantity)
			liftinfos[liftIdx].Setinfos[setIdx].Quantity.Weight = weight
		}
	}
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
