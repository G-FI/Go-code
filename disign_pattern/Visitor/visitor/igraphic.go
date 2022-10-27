package visitor

type IGraphic interface {
	Move(int, int)
	Scale(float32)
	GetDetails() string
	Draw()
	Accept(IVisitor) //表示这个接口的对象可以被观察到->通过接口进行访问具体类
}
