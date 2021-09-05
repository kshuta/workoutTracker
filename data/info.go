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
