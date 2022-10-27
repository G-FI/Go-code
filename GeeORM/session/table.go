package session

import (
	"fmt"
	"geeorm/log"
	"geeorm/schema"
	"reflect"
	"strings"
)

//session根据传入的对象创建或者返回对应的table(根据当前是否有那个table引用)
func (s *Session) Model(obj interface{}) *Session {
	//如果当前obj类型的name == s.refTable.Name 直接返回不用解析了,否者调用schema.Parse获取解析的schema
	if s.refTable != nil && reflect.TypeOf(obj).Name() == s.refTable.Name {
		return s
	}
	s.refTable = schema.Parse(obj, s.dial)
	return s
}

//获取当前session引用的table, 使用RefTable之前一定是使用过Model的
func (s *Session) RefTable() *schema.Schema {
	if s.refTable == nil {
		log.Error("this session's RefTable is nil")
	}
	return s.refTable
}

//根据model 创建的table 执行sql语句创建数据库table
func (s *Session) CreateTable() error {
	t := s.RefTable()
	columns := make([]string, 0)
	for _, filed := range t.FieldNames {
		columns = append(columns, filed)
	}
	desc := strings.Join(columns, ",")
	sql := fmt.Sprintf("create table %s(%s)", t.Name, desc)
	_, err := s.Raw(sql).Exec()
	return err
}
func (s *Session) DropTable() error {
	t := s.RefTable()
	_, err := s.Raw("drop table if exists " + t.Name).Exec()
	s.refTable = nil
	return err
}
func (s *Session) HasTable() bool {
	t := s.RefTable()
	sql, vars := s.dial.TableExist(t.Name)
	row := s.Raw(sql, vars...).QueryRow()
	var name string
	row.Scan(&name)
	return name == t.Name
}
