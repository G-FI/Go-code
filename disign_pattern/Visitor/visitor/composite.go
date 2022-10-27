package visitor

type CompoundGraphic struct {
	graphics []IGraphic
}

func NewCompoundGraphic() *CompoundGraphic {
	return &CompoundGraphic{
		graphics: nil,
	}
}

func (c *CompoundGraphic) Accept(visitor IVisitor) {
	visitor.AcceptCompoundGraphic(c)
}

func (c *CompoundGraphic) Move(x int, y int) {
	for _, child := range c.graphics {
		child.Move(x, y)
	}
}

func (c *CompoundGraphic) Draw() {
	for _, child := range c.graphics {
		child.Draw()
	}
}
func (c *CompoundGraphic) GetDetails() string {
	details := "_%_CompoundGraphic:"
	for _, child := range c.graphics {
		details = details + child.GetDetails()
	}
	details += "_%_"
	return details
}

func (c *CompoundGraphic) Scale(scale float32) {
	for _, child := range c.graphics {
		child.Scale(scale)
	}
}

func (c *CompoundGraphic) Add(child IGraphic) {
	c.graphics = append(c.graphics, child)
}

func (c *CompoundGraphic) Remove(child IGraphic) {
	i := 0
	for _, graphic := range c.graphics {
		if graphic != child {
			c.graphics[i] = graphic
			i++
		}
	}
	c.graphics = c.graphics[0 : len(c.graphics)-1]
}

func (c *CompoundGraphic) GetChildren() []IGraphic {
	return c.graphics
}
