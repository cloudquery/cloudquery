package resources

import (
	"reflect"
)

func (r *Resource) description() string {
	typ := reflect.TypeOf(r.Struct).Elem()
	return "https://pkg.go.dev/" + typ.PkgPath() + "#" + typ.Name()
}
