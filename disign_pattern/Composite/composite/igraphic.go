package composite

type iGraphic interface {
	Move(int, int)
	Scale(float32)
	Draw()
}
