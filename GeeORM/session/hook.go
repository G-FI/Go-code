package session

import "reflect"

const (
	BeforeInsert = "BeforeInsert"
	AfterInsert  = "AfterInsert"
	BeforeQuery  = "BeforeQuery"
	AfterQuery   = "AfterQuery"
	BeforeUpdate = "BeforeUpdate"
	AfterUpdate  = "AfterUpdate"
	BeforeDelete = "BeforeDelete"
	AfterDelete  = "AfterDelete"
)

//CallMethod 根据传入的对象和方法，类型和方法进行调用, 如果值传入方法则对对象不修改，如果传入对象指针，则需要调用
//该对象的方法，可能会修改对应对象
func (s *Session) CallMethod(method string, value interface{}) {
	//调用通过session中的model，反射获取对应方法，传入此session作为参数进行调用
	fm := reflect.ValueOf(s.RefTable().Model).MethodByName(method)
	if value != nil {
		fm = reflect.ValueOf(value).MethodByName(method)
	}

	prams := []reflect.Value{reflect.ValueOf(s)}
	if fm.IsValid() {
		fm.Call(prams)
	}
	return
}
