package first_attempt

var State int

func GetSingleton() int {
	if State == 0 {
		State = 42
	}
	return State
}
