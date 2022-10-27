package factory

//concrete product
type m416 struct {
	name  string
	power int
}

func (m *m416) SetName(name string) {
	m.name = name
}

func (m *m416) GetName() string {
	return m.name
}

func (m *m416) SetPower(power int) {
	m.power = power
}

func (m *m416) GetPower() int {
	return m.power
}

func NewM416() IGun {
	return &m416{
		name:  "m416",
		power: 45,
	}
}
