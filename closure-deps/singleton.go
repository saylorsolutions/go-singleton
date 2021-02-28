package closure_deps

type Singleton interface {
	GetValue() int
	SetValue(newValue int)
}

type implProxy struct {
	getValue func() int
	setValue func(int)
}

func (p *implProxy) GetValue() int {
	if p.getValue != nil {
		return p.getValue()
	}
	return 0
}

func (p *implProxy) SetValue(newValue int) {
	if p.setValue != nil {
		p.setValue(newValue)
	}
}

var getSingleton func() Singleton

func InitSingleton(initParam int) {
	// Prevent double initialization
	if getSingleton == nil {
		var value int
		var impl Singleton
		// Assign the function
		getSingleton = func() Singleton {
			if impl == nil {
				// Lazy init from initialization value in outer context
				value = initParam

				impl = &implProxy{
					getValue: func() int {
						return value
					},
					setValue: func(newValue int) {
						value = newValue
					},
				}
			}
			return impl
		}
	}
}

func GetSingleton() Singleton {
	if getSingleton == nil {
		panic("Attempting to retrieve uninitialized singleton")
	}
	return getSingleton()
}
