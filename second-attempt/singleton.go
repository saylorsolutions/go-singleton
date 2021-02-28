package second_attempt

type Singleton struct {
	value int
}

func (s *Singleton) GetValue() int {
	return s.value
}

func (s *Singleton) SetValue(newValue int) {
	s.value = newValue
}

var state = Singleton{}

func GetSingleton() *Singleton {
	if state.value == 0 {
		state.value = 42
	}
	return &state
}
