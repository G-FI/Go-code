1. dialect.go 设置存储支持的数据库类型的接口，采用键值对存取key:name, value:数据库类型
2. 其他文件表示具体的数据库，实现Dialect接口