package main

import (
	"visitor/visitor"
)

func main() {
	iGraph1 := visitor.NewCircle(0, 0, 0)
	iGraph2 := visitor.NewRectangle(1, 1, 1, 1)
	iGraph3 := visitor.NewCircle(2, 2, 2)
	iGraph4 := visitor.NewRectangle(3, 3, 3, 3)
	iGraph5 := visitor.NewCompoundGraphic()
	iGraph5.Add(iGraph1)
	iGraph5.Add(iGraph3)
	exportVisitor := visitor.NewExportVisitor()

	iGraph2.Accept(exportVisitor)
	iGraph4.Accept(exportVisitor)
	iGraph5.Accept(exportVisitor)

	dyeingVisitor := visitor.NewDyeingVisitor("green")
	iGraph1.Accept(dyeingVisitor)
	iGraph2.Accept(dyeingVisitor)
	iGraph3.Accept(dyeingVisitor)
	iGraph4.Accept(dyeingVisitor)
	iGraph5.Accept(dyeingVisitor) //will panic
}
