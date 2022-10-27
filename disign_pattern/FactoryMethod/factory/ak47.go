package factory

//concrete product
type ak47 struct {
	name  string
	power int
}

func (a *ak47) SetName(name string) {
	a.name = name
}

func (a *ak47) GetName() string {
	return a.name
}

func (a *ak47) SetPower(power int) {
	a.power = power
}

func (a *ak47) GetPower() int {
	return a.power
}

func NewAk47() IGun {
	return &ak47{
		name:  "ak47",
		power: 55,
	}
}
