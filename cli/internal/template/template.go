package template

import (
	"fmt"
	"regexp"

	"github.com/thoas/go-funk"
)

var reVariables = regexp.MustCompile(`@@([a-zA-Z0-9_\.-]+)`)

func ReplaceVariables(src []byte, values any) ([]byte, error) {
	var lastErr error
	result := reVariables.ReplaceAllFunc(src, func(s []byte) []byte {
		variablePath := s[2:]
		res := funk.Get(values, string(variablePath))
		if res == nil {
			lastErr = fmt.Errorf("variable %s not found", variablePath)
			return s
		}
		resString, ok := res.(string)
		if !ok {
			lastErr = fmt.Errorf("variable %s is not a string", variablePath)
			return s
		}
		return []byte(resString)
	})
	return result, lastErr
}