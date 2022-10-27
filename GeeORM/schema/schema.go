package schema

import (
	"geeorm/dialect"
	"reflect"
)

//table 类型，=>tableName, 属性
type Schema struct {
	Model      interface{}
	Name       string
	FieldNames []string
	FieldMap   map[string]*Field
}

//table中属性类型，名称，以及限制
type Field struct {
	Name string
	Type string
	Tag  string
}

//通过传入golang 对象，创建对应的table
func Parse(obj interface{}, d dialect.Dialect) *Schema {
	//传入的obj是指针类型，需要先得到真实值再获取类型
	objType := reflect.Indirect(reflect.ValueOf(obj)).Type()
	table := &Schema{
		Model:    obj,
		Name:     objType.Name(),
		FieldMap: make(map[string]*Field),
	}
	//添加obj的字段到表中
	for i := 0; i < objType.NumField(); i++ {
		p := objType.Field(i)
		//填入table中的字段必须不是匿名的，并且是导出的，以字段名作为属性名
		if !p.Anonymous && p.IsExported() {
			f := &Field{
				Name: p.Name,
				//通过创建一个字段类型的对象Value，然后让数据库dialect返回对应类型字符串
				Type: d.DataTypeOf(reflect.Indirect(reflect.New(p.Type))),
			}
			if tag, ok := p.Tag.Lookup("geeorm"); ok {
				f.Tag = tag
			}
			table.FieldNames = append(table.FieldNames, p.Name)
			table.FieldMap[p.Name] = f
		}
	}
	return table
}

//传入对象，返回对象字段切片
//根据数据库中列的顺序，从对象中找到对应的值，按顺序平铺。
func (s *Schema) RecordValues(dest interface{}) []interface{} {
	destValue := reflect.Indirect(reflect.ValueOf(dest))
	var vars []interface{}
	var field interface{}
	for _, name := range s.FieldNames {
		field = destValue.FieldByName(name).Interface()
		vars = append(vars, field)
	}
	return vars
}
