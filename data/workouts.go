package data

import (
	"time"
)

type Workout struct {
	Id        int
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
	if workout.Name == "" || workout.WeekNo == 0 || workout.Date.IsZero() || workout.CreatedAt.IsZero() {
		err = ErrWorkoutMissingField
		return
	}

	statement := "insert into workouts (name, week_no, date, created_at) values ($1, $2, $3, $4) returning id"
	stmt, err := db.Prepare(statement)
	if err != nil {
		return
	}

	defer stmt.Close()
	err = stmt.QueryRow(workout.Name, workout.WeekNo, workout.Date, workout.CreatedAt).Scan(&workout.Id)

	return
}

func GetWorkout(id int) (workout Workout, err error) {
	err = db.QueryRowx("select * from workouts where id=$1", id).StructScan(&workout)
	return
}

func (workout *Workout) Update() (err error) {
	_, err = db.Exec("update workouts set name = $2, week_no= $3, date = $4 where id = $1", workout.Id, workout.Name, workout.WeekNo, workout.Date)
	return
}

func (workout *Workout) Delete() (err error) {
	_, err = db.Exec("delete from workouts where id = $1", workout.Id)
	return
}
