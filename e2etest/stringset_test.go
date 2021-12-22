package e2etest

type stringSet map[string]struct{}

func stringSetFromList(list []string) stringSet {
	m := make(stringSet)
	for _, item := range list {
		m[item] = struct{}{}
	}
	return m
}

func (s stringSet) Sub(other stringSet) stringSet {
	d := make(stringSet)
	for k := range other {
		if _, ok := s[k]; !ok {
			d[k] = struct{}{}
		}
	}
	return d
}

func (s stringSet) Eq(other stringSet) bool {
	if s.Len() != other.Len() {
		return false
	}
	for k := range other {
		if _, ok := s[k]; !ok {
			return false
		}
	}
	return true
}

func (s stringSet) Len() int {
	return len(s)
}

func (s stringSet) ToList() []string {
	r := make([]string, 0, s.Len())
	for k := range s {
		r = append(r, k)
	}
	return r
}
