package client

import "github.com/cloudquery/cloudquery/pkg/client/database"

type Storage interface {
	DSN() string

	DialectExecutor() database.DialectExecutor
}

type SimpleStorage struct {
	dsn     string
	dialect database.DialectExecutor
}

func NewStorage(dsn string, dialect database.DialectExecutor) *SimpleStorage {
	if dialect == nil {
		_, dialect, _ = database.GetExecutor(dsn, nil)
	}
	return &SimpleStorage{dsn, dialect}
}

func (s SimpleStorage) DSN() string {
	return s.dsn
}

func (s SimpleStorage) DialectExecutor() database.DialectExecutor {
	return s.dialect
}
