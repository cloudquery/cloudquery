package client

type Storage interface {
	DSN() string
}

type SimpleStorage struct {
	dsn string
}

func NewStorage(dsn string) *SimpleStorage {
	return &SimpleStorage{dsn}
}

func (s SimpleStorage) DSN() string {
	return s.dsn
}
