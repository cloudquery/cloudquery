package database

type Storage interface {
	DSN() string
	DialectExecutor() DialectExecutor
}

type SimpleStorage struct {
	dsn     string
	dialect DialectExecutor
}

func NewStorage(dsn string, dialect DialectExecutor) *SimpleStorage {
	if dialect == nil {
		_, dialect, _ = GetExecutor(dsn, nil)
	}
	return &SimpleStorage{dsn, dialect}
}

func (s SimpleStorage) DSN() string {
	return s.dsn
}

func (s SimpleStorage) DialectExecutor() DialectExecutor {
	return s.dialect
}
