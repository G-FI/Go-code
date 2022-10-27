package decorator

type fileDataSource struct {
	data string
}

func newFileDataSource() *fileDataSource {
	return &fileDataSource{data: ""}
}

func (f *fileDataSource) Write(data string) {
	f.data = data
}
func (f *fileDataSource) Read() string {
	return f.data
}
