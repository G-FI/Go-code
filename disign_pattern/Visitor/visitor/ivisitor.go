package visitor

type IVisitor interface {
	AcceptCircle(*Circle)
	AcceptRectangle(*Rectangle)
	AcceptCompoundGraphic(*CompoundGraphic)
}
