package resources

import (
	"reflect"
	"strings"

	"github.com/cloudquery/plugin-sdk/codegen"
)

type Resource struct {
	Name       string
	Struct     any
	Service    string
	SubService string

	PKColumns    []string
	ExtraColumns codegen.ColumnDefinitions

	Children []*Resource
	parent   *Resource

	PreResolver  string
	PostResolver string

	table *codegen.TableDefinition
}

func (r *Resource) sanitize() {
	for _, child := range r.Children {
		child.parent = r
		if len(child.Service) == 0 {
			child.Service = r.Service
		}
	}

	if len(r.SubService) == 0 {
		r.SubService = csr.ToSnake(reflect.TypeOf(r.Struct).Elem().Name())
		if !strings.HasSuffix(r.SubService, "s") {
			r.SubService += "s"
		}
	}
}

func (r *Resource) Table() *codegen.TableDefinition {
	return r.table
}

func (r *Resource) Parent() *Resource {
	return r.parent
}
