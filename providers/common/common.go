package common

import (
	"strings"
)

func StringListToString(arr []*string) *string {
	s := strings.Builder{}
	isEmpty := true
	for _, val := range arr {
		if val != nil {
			s.WriteString(*val)
			s.WriteString(",")
			isEmpty = false
		}
	}
	if isEmpty {
		return nil
	} else {
		res := s.String()
		return &res
	}
}

type ClientInterface interface {
	CollectResource(resource string, config interface{}) error
}
