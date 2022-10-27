package session

import "geeorm/log"

//封装标准库提供的transaction操做

func (s *Session) Begin() (err error) {
	log.Info("transaction start")
	if s.tx, err = s.db.Begin(); err != nil {
		log.Error(err)
	}
	return
}

func (s *Session) Commit() (err error) {
	log.Info("transaction commit")
	if err = s.tx.Commit(); err != nil {
		log.Error(err)
	}
	return
}
func (s *Session) RollBack() (err error) {
	log.Info("transaction rollback")
	if err = s.tx.Rollback(); err != nil {
		log.Error(err)
	}
	return
}
