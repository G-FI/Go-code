package composite

import "fmt"

type circle struct {
	radius int
	center [2]int
}

func NewCircle(x int, y int, r int) *circle {
	return &circle{radius: r,
		center: [2]int{x, y},
	}
}

func (c *circle) Move(x int, y int) {
	c.center[0] += x
	c.center[1] += y
}
func (c *circle) Scale(scale float32) {
	c.radius = int(float32(c.radius) * scale)
}
func (c *circle) Draw() {
	fmt.Printf("draw a circle at (%d, %d) with radius %d\n", c.center[0], c.center[1], c.radius)
}

type rectangle struct {
	start  [2]int
	length int
	width  int
}

func NewRectangle(x, y, length, width int) *rectangle {
	return &rectangle{start: [2]int{x, y},
		length: length,
		width:  width,
	}
}
func (r *rectangle) Move(x int, y int) {
	r.start[0] += x
	r.start[1] += y
}
func (r *rectangle) Scale(scale float32) {
	r.length = int(float32(r.length) * scale)
	r.width = int(float32(r.width) * scale)
}
func (r *rectangle) Draw() {
	fmt.Printf("draw a rectangle at (%d, %d) with length %d and width %d\n",
		r.start[0], r.start[1], r.length, r.width)
}
