package composite

type compoundGraphic struct {
	graphics []iGraphic
}

func NewCompoundGraphic() *compoundGraphic {
	return &compoundGraphic{
		graphics: nil,
	}
}

func (c *compoundGraphic) Move(x int, y int) {
	for _, child := range c.graphics {
		child.Move(x, y)
	}
}

func (c *compoundGraphic) Draw() {
	for _, child := range c.graphics {
		child.Draw()
	}
}

func (c *compoundGraphic) Scale(scale float32) {
	for _, child := range c.graphics {
		child.Scale(scale)
	}
}

func (c *compoundGraphic) Add(child iGraphic) {
	c.graphics = append(c.graphics, child)
}

func (c *compoundGraphic) Remove(child iGraphic) {
	i := 0
	for _, graphic := range c.graphics {
		if graphic != child {
			c.graphics[i] = graphic
			i++
		}
	}
	c.graphics = c.graphics[0 : len(c.graphics)-1]
}
