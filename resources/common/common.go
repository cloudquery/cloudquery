package common

import (
	"strings"
)

func StringListToString(arr *[]string) *string {
	s := strings.Builder{}
	isEmpty := true
	for _, val := range *arr {
		s.WriteString(val)
		s.WriteString(",")
		isEmpty = false
	}
	if isEmpty {
		return nil
	} else {
		res := s.String()
		return &res
	}
}
