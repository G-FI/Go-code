package visitor

import (
	"fmt"
)

type Circle struct {
	radius int
	center [2]int
}

func NewCircle(x int, y int, r int) *Circle {
	return &Circle{radius: r,
		center: [2]int{x, y},
	}
}

func (c *Circle) Accept(visitor IVisitor) {
	visitor.AcceptCircle(c)
}

func (c *Circle) Move(x int, y int) {
	c.center[0] += x
	c.center[1] += y
}
func (c *Circle) Scale(scale float32) {
	c.radius = int(float32(c.radius) * scale)
}
func (c *Circle) Draw() {
	fmt.Printf("draw a Circle at (%d, %d) with radius %d\n", c.center[0], c.center[1], c.radius)
}
func (c *Circle) GetDetails() string {
	return fmt.Sprintf("<center:(%d, %d), radius:%d>\n", c.center[0], c.center[1], c.radius)
}
func (c *Circle) GetCenter() [2]int {
	var ret [2]int
	ret[0], ret[1] = c.center[0], c.center[1]
	return ret
}
func (c *Circle) GetRadius() int {
	return c.radius
}

type Rectangle struct {
	start  [2]int
	length int
	width  int
}

func NewRectangle(x, y, length, width int) *Rectangle {
	return &Rectangle{start: [2]int{x, y},
		length: length,
		width:  width,
	}
}

func (r *Rectangle) Accept(visitor IVisitor) {
	visitor.AcceptRectangle(r)
}

func (r *Rectangle) Move(x int, y int) {
	r.start[0] += x
	r.start[1] += y
}
func (r *Rectangle) Scale(scale float32) {
	r.length = int(float32(r.length) * scale)
	r.width = int(float32(r.width) * scale)
}
func (r *Rectangle) Draw() {
	fmt.Printf("draw a Rectangle at (%d, %d) with length %d and width %d\n",
		r.start[0], r.start[1], r.length, r.width)
}
func (r *Rectangle) GetDetails() string {
	return fmt.Sprintf("<start:(%d, %d), lenght:%d, width:%d>\n",
		r.start[0], r.start[1], r.length, r.width)
}

func (r *Rectangle) GetStart() [2]int {
	var ret [2]int
	ret[0], ret[1] = r.start[0], r.start[1]
	return ret
}

func (r *Rectangle) GetLength() int {
	return r.length
}
func (r *Rectangle) GetWidth() int {
	return r.width
}
