package closure

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

var getSingleton = func() func() Singleton {
	var value int
	var impl Singleton
	return func() Singleton {
		if impl == nil {
			// Lazy init
			value = 42

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
}()

func GetSingleton() Singleton {
	return getSingleton()
}
