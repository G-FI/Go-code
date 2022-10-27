package visitor

import (
	"fmt"
)

///exportVisitor
type exportVisitor struct {
}

func NewExportVisitor() *exportVisitor {
	return &exportVisitor{}
}

//实现不同的观察者方法之后，为观察对象添加accept方法

func (e *exportVisitor) AcceptCircle(circle *Circle) {
	center := circle.GetCenter()
	radius := circle.GetRadius()
	fmt.Printf("%%circle%%<center:(%d, %d), radius:%d>\n", center[0], center[1], radius)
}

func (e *exportVisitor) AcceptRectangle(rectangle *Rectangle) {
	start := rectangle.GetStart()
	length := rectangle.GetLength()
	width := rectangle.GetWidth()
	fmt.Printf("%%rectangle%%<start:(%d, %d), lenght:%d, width:%d>\n",
		start[0], start[1], length, width)
}

func (e *exportVisitor) AcceptCompoundGraphic(graphic *CompoundGraphic) {
	fmt.Println("%compound graphic%")
	for _, child := range graphic.GetChildren() {
		child.Accept(e)
	}
	fmt.Println("%compound graphic%")
}

type dyeingVisitor struct {
	color string
}

func NewDyeingVisitor(color string) *dyeingVisitor {
	return &dyeingVisitor{color: color}
}
func (d *dyeingVisitor) AcceptCircle(circle *Circle) {
	fmt.Printf("dyeing a circle: %s\n", circle.GetDetails())
}

func (d *dyeingVisitor) AcceptRectangle(rectangle *Rectangle) {
	fmt.Printf("dyeing a rectange: %s\n", rectangle.GetDetails())
}

func (d *dyeingVisitor) AcceptCompoundGraphic(graphic *CompoundGraphic) {
	panic("implement me")
}
