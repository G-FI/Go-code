package decorator

import "strings"

// dataSourceDecorator base decorator
type dataSourceDecorator struct {
	source iDataSource
}

func NewDataSourceDecorator(source iDataSource) *dataSourceDecorator {
	return &dataSourceDecorator{source: source}
}

func (d *dataSourceDecorator) Write(data string) {
	d.source.Write(data)
}
func (d *dataSourceDecorator) Read() string {
	return d.source.Read()
}

// encryptionDecorator extended Decorator
type encryptionDecorator struct {
	*dataSourceDecorator
}

func NewEncryptionDecorator(source iDataSource) *encryptionDecorator {
	return &encryptionDecorator{
		NewDataSourceDecorator(source),
	}
}

func (e *encryptionDecorator) Write(data string) {
	encrypted := "<encrypted>" + data + "<encrypted>"
	e.dataSourceDecorator.source.Write(encrypted)
}
func (e *encryptionDecorator) Read() string {
	encrypted := e.dataSourceDecorator.source.Read()
	start := strings.Index(encrypted, ">") + 1
	end := strings.LastIndex(encrypted, "<")
	return encrypted[start:end]
}

//extend Decorator
type compressDecorator struct {
	*dataSourceDecorator
}

func NewCompressDecorator(source iDataSource) *compressDecorator {
	return &compressDecorator{
		NewDataSourceDecorator(source),
	}
}

func (c *compressDecorator) Write(data string) {
	c.dataSourceDecorator.Write(data)
}
func (c *compressDecorator) Read() string {
	return c.dataSourceDecorator.Read()
}
