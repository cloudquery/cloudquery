package resource

import (
	"reflect"
	"strings"

	"github.com/cloudquery/plugin-sdk/caser"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/gertd/go-pluralize"
	"github.com/iancoleman/strcase"
	"golang.org/x/exp/slices"
)

type Resource struct {
	Service             string
	SubService          string
	Name                string // rare case
	Struct              any
	Multiplex           string
	ExtraColumns        []codegen.ColumnDefinition
	PKColumns           []string
	IgnoreInTestColumns []string
	Table               *codegen.TableDefinition
	UnwrapStructFields  []string

	Resolver    *FuncParams
	PreResolver *FuncParams
	fetcher     *Fetcher // will be used in the fetcher gen
	SkipMock    bool     // don't generate mock test

	// used for generating resolver and mock tests, but set automatically
	parent   *Resource
	Children []*Resource
}

// StructName returns the name of the resource's Struct field
func (r *Resource) StructName() string {
	// because usually the 'Struct' field contains a pointer, we need to dereference with '.Elem()'.
	return reflect.TypeOf(r.Struct).Elem().Name()
}

// StructPackage returns the package of the resource's Struct field
func (r *Resource) StructPackage() string {
	// because usually the 'Struct' field contains a pointer, we need to dereference with '.Elem()'.
	return reflect.TypeOf(r.Struct).Elem().PkgPath()
}

// StructPackageName returns the package name for the resource's Struct field
func (r *Resource) StructPackageName() string {
	return getStructPackageName(r.Struct)
}

// Import return the import to be added to use for the r.Struct
func (r *Resource) Import() string {
	return reflect.TypeOf(r.Struct).Elem().PkgPath()
}

func (r *Resource) SchemaFuncName() string {
	if r.parent != nil {
		return strcase.ToLowerCamel(r.SubService)
	}
	return strcase.ToCamel(r.SubService)
}

// CloudQueryServiceName is used for accessing 'client.Services().{{.CloudqueryServiceName}}' in templates
func (r *Resource) CloudQueryServiceName() string {
	csr := caser.New()
	return csr.ToPascal(r.Service)
}

func (r *Resource) hasField(fieldName string) bool {
	for _, field := range reflect.VisibleFields(reflect.TypeOf(r.Struct).Elem()) {
		if field.Name == fieldName {
			return true
		}
	}
	return false
}

func (r *Resource) sanitize() {
	// set service
	switch {
	case len(r.Service) > 0:
	// nop
	case r.parent != nil:
		r.Service = r.parent.Service
	default:
		r.Service = strings.TrimPrefix(r.StructPackageName(), "arm")
	}

	if r.hasField("ID") && !slices.Contains(r.PKColumns, "id") {
		r.PKColumns = append(r.PKColumns, "id")
	}

	if len(r.SubService) == 0 {
		parts := strings.Split(strcase.ToSnake(reflect.TypeOf(r.Struct).Elem().Name()), "_")

		r.SubService = strings.Join(append(parts[:len(parts)-1], pluralize.NewClient().Plural(parts[len(parts)-1])), "_")
	}

	// set name
	if len(r.Name) == 0 {
		r.Name = "azure_" + r.Service + "_" + r.SubService
	}

	for _, child := range r.Children {
		child.parent = r
		child.sanitize()
	}
}
