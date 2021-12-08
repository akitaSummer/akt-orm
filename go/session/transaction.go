package session

import "aktorm/log"

func (s *Session) Begin() error {
	log.Info("transaction begin")
	var err error
	if s.tx, err = s.db.Begin(); err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (s *Session) Commit() error {
	log.Info("transaction commit")
	var err error
	if err = s.tx.Commit(); err != nil {
		log.Error(err)
	}
	return err
}

func (s *Session) Rollback() error {
	log.Info("transaction rollback")
	var err error
	if err = s.tx.Rollback(); err != nil {
		log.Error(err)
	}
	return err
}
