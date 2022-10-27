package decorator

type iDataSource interface {
	Write(data string)
	Read() string
}

func NewDataSource() iDataSource {
	return newFileDataSource()
}
