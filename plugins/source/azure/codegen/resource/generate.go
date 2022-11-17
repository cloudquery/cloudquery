package resource

import (
	"embed"
	"fmt"
	"os"
	"path"
	"reflect"
	"runtime"
	"strconv"
	"strings"

	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/iancoleman/strcase"
	"golang.org/x/exp/slices"
)

//go:embed templates/*.go.tpl
var templatesFS embed.FS

func (r *Resource) Generate() error {
	r.sanitize()
	fmt.Println("Generating", r.Name)
	if err := r.propagateTable(); err != nil {
		return err
	}

	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return fmt.Errorf("failed to get caller information")
	}

	dir := path.Dir(filename)
	dir = path.Join(dir, "../../resources/services", r.Service)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create directory %s: %w", dir, err)
	}

	if err := r.generateSchema(dir); err != nil {
		return err
	}

	if r.fetcher != nil {
		if err := r.generateFetch(dir); err != nil {
			return err
		}

		if !r.SkipMock {
			if err := r.generateMockTest(dir); err != nil {
				return err
			}
		}
	}

	for _, child := range r.Children {
		if err := child.Generate(); err != nil {
			return err
		}
	}

	return nil
}

func (r *Resource) propagateTable() error {
	if !r.hasField("SubscriptionID") {
		r.ExtraColumns = append(r.ExtraColumns, SubscriptionIDCol)
	}
	opts := []codegen.TableOption{
		codegen.WithExtraColumns(r.ExtraColumns),
		codegen.WithPKColumns(r.PKColumns...),
		codegen.WithUnwrapStructFields(append(r.UnwrapStructFields, "Properties")),
		codegen.WithUnwrapAllEmbeddedStructs(),
		codegen.WithTypeTransformer(fixStringArray),
	}

	if err := r.checkName(); err != nil {
		return err
	}

	var err error
	r.Table, err = codegen.NewTableFromStruct(
		r.Name,
		r.Struct,
		opts...,
	)
	if err != nil {
		return fmt.Errorf("error generating %s: %w", r.Name, err)
	}
	azureType := reflect.TypeOf(r.Struct).Elem()
	r.Table.Description = fmt.Sprintf("https://pkg.go.dev/%s#%s", azureType.PkgPath(), azureType.Name())
	r.Table.Resolver = "fetch" + strcase.ToCamel(r.SubService)

	r.propagateFetcher()
	if r.fetcher != nil && r.fetcher.PreResolver != nil {
		r.Table.PreResourceResolver = "get" + strcase.ToCamel(r.SubService)
	}

	// set up multiplex
	switch {
	case r.parent != nil:
	// nop, no need for multiplex
	case len(r.Multiplex) > 0:
		// custom multiplex already specified
		r.Table.Multiplex = r.Multiplex
	default:
		r.Table.Multiplex = "client.SubscriptionMultiplex"
	}

	for _, child := range r.Children {
		r.Table.Relations = append(r.Table.Relations, child.SchemaFuncName()+"()")
	}
	slices.Sort(r.Table.Relations)

	r.propagateColumns()

	return nil
}

func (r *Resource) checkName() error {
	// table name must be plural
	if strings.HasSuffix(r.Name, "s") {
		return nil
	}

	// actually, check if last part is _vXXX
	parts := strings.Split(r.Name, "_")
	if len(parts) > 1 {
		if strings.HasSuffix(parts[len(parts)-2], "s") {
			// check the _vXXX
			last := parts[len(parts)-1]
			if strings.HasPrefix(last, "v") {
				_, err := strconv.ParseInt(strings.TrimPrefix(last, "v"), 10, 32)
				if err == nil {
					return nil
				}
			}
		}
	}

	return fmt.Errorf("invalid table name: %s. must be plural", r.Name)
}
