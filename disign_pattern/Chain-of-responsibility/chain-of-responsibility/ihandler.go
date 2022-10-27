package chain_of_responsibility

type iHandler interface {
	SetNext(iHandler)
	Handle()
}
