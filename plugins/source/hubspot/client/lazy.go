package client

type lazy[CLIENT, CONFIG any] struct {
	value *CLIENT
	cfg   *CONFIG
	init  func(*CONFIG) *CLIENT
}

func (l *lazy[CLIENT, CONFIG]) Get() *CLIENT {
	if l.value == nil {
		l.value = l.init(l.cfg)
	}

	return l.value
}
