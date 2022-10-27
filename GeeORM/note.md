1. **对象关系模型就是依赖于方法和对象类型就能构造出相应的sql语句以及参数，参数依赖对象字段值，sql语句依赖于对象字段名，因此关键就是反射(根据对象获取字段Name，以及相反)**
2. log包，实现log等级输出信息，封装log方法
3. session包，封装对sql语句的调用接口, 进行同一日志打印，并且重用一个session查询
4. TODO: ...运算符将参数算作一个数组，所以参数在session中为一个切片，然后将切片传给sql.db.Exec时，会被当作一整个参数，导致不支持,还有query没结果
5. golang 反射：
   1. ```go
      接口类型：type Type interface{...}
              type Value interface{...}
              type StructField struct
      通过对象获取reflect.Type/.Value方法：
            reflect.TypeOf(obj)
            reflect.ValueOf(obj)
      常用方法：
            type常用方法：
                TypeName() string,
                TypeKind() Kind,
      访问字段方法：
            NumField() int,
            Field(i int) StructField,
   2. ```go
      type StructField struct {
      Name string          // 字段名
      PkgPath string       // 字段路径
      Type      Type       // 字段反射类型对象
      Tag       StructTag  // 字段的结构体标签
      Offset    uintptr    // 字段在结构体中的相对偏移
      Index     []int      // Type.FieldByIndex中的返回的索引值
      Anonymous bool       // 是否为匿名字段
      }

6. 对象模型的接口如Find(), Insert(),Update(),Delete()等都是通过传入的interface,先反射得到对象类型，然后用Model()构建出对应的table,然后构造Clause,最后通过返回的sql语句和参数执行Raw()+Exec()
