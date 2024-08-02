package values

import (
	"fmt"
)

func unwrap[A any](value any) (A, bool) {
	var unwrapped A
	var ptr *A

	switch v := value.(type) {
	case **A:
		//nolint:govet
		if v == (**A)(nil) || v == nil {
			return unwrapped, false
		}
		ptr = *v

	case *A:
		ptr = v

	case A:
		return v, true

	default:
		panic(fmt.Sprintf("unwrapping %T to %T isn't supported", value, unwrapped))
	}

	//nolint:govet
	if ptr == nil || ptr == (*A)(nil) {
		return unwrapped, false
	}
	return *ptr, true
}
