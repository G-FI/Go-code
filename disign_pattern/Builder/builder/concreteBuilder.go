package builder

type CarBuilder struct {
	car *Car
}

func (c *CarBuilder) SetSeats() {
	c.car.seats = 2
}
func (c *CarBuilder) SetEngin() {
	c.car.engine = "Car engine"
}

func (c *CarBuilder) Reset() {
	c.car = &Car{
		seats:  0,
		engine: "",
	}
}
func (c *CarBuilder) GetProduct() *Car {
	return c.car
}

func newCarBuilder() *CarBuilder {
	return &CarBuilder{
		car: nil,
	}
}

type CarManualBuilder struct {
	manual *CarManual
}

func (c *CarManualBuilder) SetSeats() {
	c.manual.seats = 2
}
func (c *CarManualBuilder) SetEngin() {
	c.manual.engine = "Car engine"
}

func (c *CarManualBuilder) Reset() {
	c.manual = &CarManual{
		seats:  0,
		engine: "",
	}
}
func (c *CarManualBuilder) GetProduct() *CarManual {
	return c.manual
}

func newCarManualBuilder() *CarManualBuilder {
	return &CarManualBuilder{
		manual: nil,
	}
}
