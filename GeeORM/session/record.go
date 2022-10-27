package session

import (
	"errors"
	"geeorm/clause"
	"reflect"
)

func (s *Session) Insert(values ...interface{}) (int64, error) {
	/*Insert依赖于insert和values子句*/
	//循环获取每一个obj的vars, 组成一个vars的切片组合，
	//配置clause,build得到sql语句和flatten参数
	//调用s.Raw().Exec()进行操做
	//返回number of rows of affected

	recordVars := make([]interface{}, 0)
	table := s.Model(values[0]).RefTable()
	for _, value := range values {
		/*对每个传入对象调用beforeInsertHook*/
		s.CallMethod(BeforeInsert, value)

		recordVars = append(recordVars, table.RecordValues(value))
	}
	s.clause.Set(clause.INSERT, table.Name, table.FieldNames)
	s.clause.Set(clause.VALUES, recordVars...)
	sql, vars := s.clause.Build(clause.INSERT, clause.VALUES)
	result, err := s.Raw(sql, vars...).Exec()
	if err != nil {
		return 0, err
	}

	s.CallMethod(AfterInsert, nil)
	return result.RowsAffected()
}

//传入一个切片指针，然后将查询到的值保存到切片中
//根据切片的*元素类型*找到反射得到字段名称对应数据库列
//构造clause, build
//执行s.QueryRows
//将返回的结果写到dest指向的切片中
func (s *Session) Find(values interface{}) error {
	/*Find依赖于SELECT WEHRE ORDER BY LIMIT子句*/
	//反射得到切片引用
	//根据切片引用反射得到element类型
	//根据element类型创建对象，转化为接口，再获取对应的table
	//现在有了table就可以构建sql语句了，就能得到查询结果rows
	//循环读取每一个row，并将row读到一各element类型的对象中
	//然后将每一个创建的对象加入到传入的slice中

	/*添加hook调用*/
	s.CallMethod(BeforeQuery, nil)

	destSlice := reflect.Indirect(reflect.ValueOf(values))
	destType := destSlice.Type().Elem()                                   //Type.Element返回符合类型的元素类型
	table := s.Model(reflect.New(destType).Elem().Interface()).RefTable() //Value.Elem()返回指针v指向的元素，而reflect.New创建的是指针型的Value
	s.clause.Set(clause.SELECT, table.Name, table.FieldNames)
	sql, vars := s.clause.Build(clause.SELECT, clause.WHERE, clause.ORDERBY, clause.LIMIT)

	rows, _ := s.Raw(sql, vars...).QueryRows()
	for rows.Next() {
		dest := reflect.New(destType).Elem()
		//rows.Scan(attrs...)是将查询结果的属性序列化顺序赋值到一组接收值中，所以
		//需要将对象的字段铺平，而我有不知道dest的具体类型，无法判断字段，
		//此时有table，就知道attribute的name，也就是dest的各个field的名字
		//通过反射可以获取字段，然后将他们*的地址*保存在一个interface slice中，传给scan就行了
		var fields []interface{}
		for _, name := range table.FieldNames {
			fields = append(fields, dest.FieldByName(name).Addr().Interface())
		}
		if err := rows.Scan(fields...); err != nil {
			return err
		}
		/*添加hook after query，传入每个对象的指针*/
		s.CallMethod(AfterQuery, dest.Addr().Interface())

		destSlice.Set(reflect.Append(destSlice, dest)) // <=>destSlice = append(destSlice, dest)
	}
	return rows.Close()
}

func (s *Session) Count() (int64, error) {
	s.clause.Set(clause.COUNT, s.RefTable().Name)
	sql, vars := s.clause.Build(clause.COUNT, clause.WHERE)
	row := s.Raw(sql, vars...).QueryRow()
	var count int64
	if err := row.Scan(&count); err != nil {
		return 0, err
	}
	return count, nil
}

//接受参数，1:tableName, 2:map[string]interface{}
//或者，1:tableName 2: []interface{}
func (s *Session) Update(value interface{}) (int64, error) {
	tableName := s.RefTable().Name
	kv, ok := value.(map[string]interface{})
	if !ok {
		//此时传入的参数是第二种类型，就是k v按顺序放再slice中，而不是放到一个map中
		kv = make(map[string]interface{})
		m := value.([]interface{})
		for i := 0; i < len(m); i += 2 {
			kv[m[i].(string)] = m[i+1]
		}
	}
	s.clause.Set(clause.UPDATE, tableName, kv)
	sql, vars := s.clause.Build(clause.UPDATE, clause.WHERE)
	result, err := s.Raw(sql, vars...).Exec()
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (s *Session) Delete() (int64, error) {
	tableName := s.RefTable().Name

	s.clause.Set(clause.DELETE, tableName)
	sql, vars := s.clause.Build(clause.DELETE, clause.WHERE)

	result, err := s.Raw(sql, vars...).Exec()
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

//Where "predict", args
func (s *Session) Where(pred string, args ...interface{}) *Session {
	var values []interface{}
	s.clause.Set(clause.WHERE, append(append(values, pred), args))
	return s
}

//OrderBy "description"
func (s *Session) OrderBy(description string) *Session {
	s.clause.Set(clause.ORDERBY, description)
	return s
}

//LIMIT num
func (s *Session) Limit(num int) *Session {
	s.clause.Set(clause.LIMIT, num)
	return s
}

func (s *Session) First(value interface{}) error {
	//反射获得原对象dest
	//使用find需要原类型的slice，故创建一个
	//链式查找，limit为一
	//给原对象dest赋值为存放结果的slice的第一个元素
	dest := reflect.Indirect(reflect.ValueOf(value))
	destSlice := reflect.New(reflect.SliceOf(dest.Type())).Elem() //New创建的是指针

	if err := s.Limit(1).Find(destSlice.Addr().Interface()); err != nil {
		return err
	}

	if destSlice.Len() == 0 {
		return errors.New("NOT FOUND")
	}
	dest.Set(destSlice.Index(0))
	return nil
}
