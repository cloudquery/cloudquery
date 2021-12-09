package resources

import "fmt"

func unexpectedResourceType(want, got interface{}) error {
	return fmt.Errorf("expected resource of type %T, but got %T", want, got)
}
