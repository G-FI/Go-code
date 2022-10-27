package geeorm

import (
	"database/sql"
	"geeorm/dialect"
	"geeorm/log"
	"geeorm/session"
	_ "github.com/mattn/go-sqlite3"
)

// Engine 负责与数据库的连接以及收尾工作
type Engine struct {
	db *sql.DB
	//新加
	dial dialect.Dialect
}

func NewEngine(driver, name string) (e *Engine, err error) {
	db, err := sql.Open(driver, name)
	if err != nil {
		log.Error(err)
		return
	}
	if err = db.Ping(); err != nil {
		log.Error(err)
		return
	}
	dial, ok := dialect.GetDialect(driver)
	if !ok {
		log.Errorf("%s dialect not support", driver)
		return
	}
	e = &Engine{db: db, dial: dial}
	log.Info("Database connected")
	return
}
func (e *Engine) Close() (err error) {
	err = e.db.Close()
	if err != nil {
		log.Error("Fail to close database")
		return
	}
	log.Info("Database closed")
	return
}

func (e *Engine) NewSession() *session.Session {
	return session.New(e.db, e.dial)
}

//添加事务处理
//事务处理函数，客户端只需要将事务操做放入TxFunc回调函数中，然后执行Transaction即可保证实务操作
type TxFunc func(s *session.Session) (result interface{}, err error)

func (e *Engine) Transaction(tf TxFunc) (result interface{}, err error) {
	//一次transaction使用一个session，需要创建session，客户端也会在TxFunc使用这个session
	//相当于一个hook，再执行客户段事务之前开启事务，执行完之后提交事务，或者出错rollback
	s := e.NewSession()
	if err = s.Begin(); err != nil {
		return nil, err
	}

	defer func() {
		if p := recover(); p != nil {
			//需要recover,执行rollback
			_ = s.RollBack()
			panic(p)
		} else if err != nil {
			_ = s.RollBack()
		} else {
			err = s.Commit()
		}
	}()

	return tf(s)
}
