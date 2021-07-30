package data

import "time"

type Workout struct {
	Id        int
	PlanId    int `db:"plan_id"`
	Name      string
	WeekNo    int `db:"week_no"`
	Date      time.Time
	CreatedAt time.Time `db:"created_at"`
}

const (
	ErrWorkoutMissingField = WorkoutErr("Workout struct is missing field")
)

type WorkoutErr string

func (e WorkoutErr) Error() string {
	return string(e)
}

func (workout *Workout) Create() (err error) {
	if workout.PlanId == 0 || workout.Name == "" || workout.WeekNo == 0 || workout.Date.IsZero() || workout.CreatedAt.IsZero() {
		err = ErrWorkoutMissingField
		return
	}

	statement := "insert into workouts (plan_id, name, week_no, date, created_at) values ($1, $2, $3, $4, $5) returning id"
	stmt, err := db.Prepare(statement)
	if err != nil {
		return
	}

	defer stmt.Close()
	err = stmt.QueryRow(workout.PlanId, workout.Name, workout.WeekNo, workout.Date, workout.CreatedAt).Scan(&workout.Id)

	return
}

func GetWorkout(id int) (workout Workout, err error) {
	err = db.QueryRowx("select * from workouts where id=$1", id).StructScan(&workout)
	return
}

func (workout *Workout) Update() (err error) {
	_, err = db.Exec("update workouts set name = $2, week_no= $3, date = $4, plan_id = $5 where id = $1", workout.Id, workout.Name, workout.WeekNo, workout.Date, workout.PlanId)
	return
}
