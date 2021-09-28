package data

import (
	"errors"

	"github.com/jmoiron/sqlx"
)

type LiftInfo struct {
	Lift     Lift
	Setinfos []SetInfo
}
type SetInfo struct {
	Set      Set
	Quantity SetQuantity
}

type WorkoutInfo struct {
	Workout   Workout
	Liftinfos []LiftInfo
}

func CreateLiftWorkout(workout *Workout, lift *Lift) (err error) {
	if workout.Id == 0 || lift.Id == 0 {
		return errors.New("Lift or Workout does not exist in database")
	}

	_, err = db.Exec("insert into workout_lifts (workout_id, lift_id) values ($1, $2)", workout.Id, lift.Id)
	return
}

func GetWorkoutLifts(workout Workout) (lifts []Lift, err error) {

	if workout.Id == 0 {
		err = errors.New("workout does not exist in database")
		return
	}

	rows, err := db.Queryx("select lift_id from workout_lifts where workout_id=$1", workout.Id)
	if err != nil {
		rows.Close()
		return
	}

	liftids := make([]int, 0)
	for rows.Next() {
		var id int
		rows.Scan(&id)
		liftids = append(liftids, id)
	}

	rows.Close()
	if rows.Err() != nil {
		err = rows.Err()
		return
	}
	query, args, err := sqlx.In("select * from lifts where id in (?)", liftids)
	if err != nil {
		return
	}
	query = db.Rebind(query)

	rows, err = db.Queryx(query, args...)
	if err != nil {
		return
	}

	for rows.Next() {
		var lift Lift
		rows.StructScan(&lift)
		lifts = append(lifts, lift)
	}

	rows.Close()
	if rows.Err() != nil {
		err = rows.Err()
		return
	}

	return
}

// どういう流れ？
// -> GetWorkout -> Get SetInfo -> GetLift // N would be SetInfo
// -> GetWorkout -> GetLift -> GetSetInfo -> // N would bet GetLift
// But, Matching Workout and Lift would would also take quite a lot of time if there is many pairs
// What should we do?
// It's good experience to try to implement a many to many relationship, so lets go with the second one.
