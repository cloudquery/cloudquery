package values

import (
	"fmt"
)

func unwrap[A any](value any) (A, bool) {
	var unwrapped A
	var ptr *A

	switch v := value.(type) {
	case *A:
		ptr = v

	case **A:
		if v == (**A)(nil) || v == nil {
			return unwrapped, false
		}
		ptr = *v

	default:
		panic(fmt.Sprintf("unwrapping %T to %T isn't supported", value, unwrapped))
	}

	if ptr == nil || ptr == (*A)(nil) {
		return unwrapped, false
	}
	return *ptr, true
}
