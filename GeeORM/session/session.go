package session

import (
	"database/sql"
	"geeorm/clause"
	"geeorm/dialect"
	"geeorm/log"
	"geeorm/schema"
	"strings"
)

//Session 负责数据库交互
type Session struct {
	db      *sql.DB         //数据库实例
	sql     strings.Builder //sql语句
	sqlVars []interface{}   //sql语句参数
	//添加支持数据库方言，以及操作表
	dial     dialect.Dialect //Session 用于和数据库交互，因此需要数据库方言接口
	refTable *schema.Schema  //Session 和对应表交互时需要表的信息
	//添加clause, 实现object oriented model
	clause clause.Clause
	//添加transaction实现事务操做
	tx *sql.Tx
}

func New(db *sql.DB, dial dialect.Dialect) *Session {
	return &Session{db: db, dial: dial}
}

func (s *Session) Reset() {
	s.sql.Reset()
	s.sqlVars = nil
	s.clause = clause.Clause{}
}

//sql.Tx 和 sql.DB的共同接口
type CommonDB interface {
	Query(query string, args ...any) (*sql.Rows, error)
	QueryRow(query string, args ...any) *sql.Row
	Exec(query string, args ...any) (sql.Result, error)
}

func (s *Session) DB() CommonDB {
	//Exec,Query,QueryRow都会使用DB来获取于数据库交互的对象，可以指sql.DB 也可以是sql.Tx取决于是否开启了事务
	//当前tx若不为空，则操做需要用事务实现，所以返回事务
	if s.tx != nil {
		return s.tx
	}
	return s.db
}

//Raw 用户调用Raw方法，传入原始查询语句，为Session赋值
func (s *Session) Raw(sql string, values ...interface{}) *Session {
	s.sql.WriteString(sql)
	s.sql.WriteString(" ")
	s.sqlVars = append(s.sqlVars, values...)
	return s
}

func (s *Session) Exec() (result sql.Result, err error) {
	//重置当前session，以便重用
	defer s.Reset()

	log.Info(s.sql.String(), s.sqlVars)
	//更改为s.BD()因为事务操做需要tx执行
	if result, err = s.DB().Exec(s.sql.String(), s.sqlVars...); err != nil {
		log.Error(err)
	}
	return
}

// QueryRow 返回一条记录
func (s *Session) QueryRow() *sql.Row {
	defer s.Reset()
	log.Info(s.sql.String(), s.sqlVars)
	row := s.DB().QueryRow(s.sql.String(), s.sqlVars...)
	return row
}

//QueryRows 返回一个记录list
func (s *Session) QueryRows() (rows *sql.Rows, err error) {
	defer s.Reset()
	log.Info(s.sql.String(), s.sqlVars)
	if rows, err = s.DB().Query(s.sql.String(), s.sqlVars...); err != nil {
		log.Error(err)
	}
	return
}
