package data

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

// どういう流れ？
// -> GetWorkout -> Get SetInfo -> GetLift // N would be SetInfo
// -> GetWorkout -> GetLift -> GetSetInfo -> // N would bet GetLift
// But, Matching Workout and Lift would would also take quite a lot of time if there is many pairs
// What should we do?
// It's good experience to try to implement a many to many relationship, so lets go with the second one.
