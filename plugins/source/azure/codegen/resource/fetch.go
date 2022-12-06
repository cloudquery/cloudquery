package resource

import (
	"bytes"
	"context"
	"fmt"
	"path"
	"reflect"
	"strings"
	"text/template"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/cloudquery/plugins/source/azure/codegen/util"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/gertd/go-pluralize"
	"github.com/iancoleman/strcase"
)

type FuncParams struct {
	Func        any
	Params      []string
	BasicValue  any
	ExtraValues []string
}

func paramsFromAny(fn any) *FuncParams {
	switch typed := fn.(type) {
	case *FuncParams:
		return typed
	case FuncParams:
		return &typed
	default:
		return &FuncParams{
			Func: fn,
		}
	}
}

func (f *FuncParams) getFetcher(parent *Resource) *Fetcher {
	if f == nil {
		return nil
	}

	refT := reflect.TypeOf(f.Func)
	fetcher := &Fetcher{
		BasicStructName: f.basicStructName(),
		ExtraValues:     f.ExtraValues,
		Params:          f.Params,
	}

	if len(fetcher.Params) == 0 {
		fetcher.Params = extractParams(parent, refT)
	}

	// we have the receiver interface as param[0]
	receiverInterface := reflect.TypeOf(f.Func).In(0)

	c := reflect.TypeOf(new(client.Services)).Elem()
	for _, field := range reflect.VisibleFields(c) {
		fieldType := field.Type.Elem()
		for _, interfaceField := range reflect.VisibleFields(fieldType) {
			if interfaceField.Type == receiverInterface {
				fetcher.Service = field.Name
				fetcher.Client = interfaceField.Name
				for m := 0; m < interfaceField.Type.NumMethod(); m++ {
					method := interfaceField.Type.Method(m)
					if methodMatch(method, refT) {
						fetcher.Method = method.Name
						ret := method.Type.Out(0)
						if ret.Kind() == reflect.Pointer {
							ret = ret.Elem()
						}
						fetcher.StructName = getStructName(ret)
						fetcher.StructPackageName = getStructPackageName(reflect.New(ret).Elem().Interface())
						fetcher.Import = ret.PkgPath()
						if fetcher.IsList() {
							// need to look up the value field
							fetcher.ValueField = getValueField(ret)
						}
					}
				}
				return fetcher
			}
		}
	}
	return fetcher
}

func (f *FuncParams) basicStructName() string {
	if f.BasicValue == nil {
		return ""
	}
	return getStructName(reflect.TypeOf(f.BasicValue).Elem())
}

func getValueField(typ reflect.Type) string {
	for _, fld := range reflect.VisibleFields(typ) {
		if fld.Type.Kind() == reflect.Slice {
			return fld.Name
		}
	}
	return ""
}

func valueVariables(parent *Resource) []string {
	if parent == nil {
		return nil
	}
	var reversed []string
	for ; parent != nil; parent = parent.parent {
		reversed = append(reversed, "*"+strcase.ToLowerCamel(pluralize.NewClient().Singular(parent.SubService))+".Name")
	}
	values := []string{"id.ResourceGroupName"}
	for i := len(reversed) - 1; i >= 0; i-- {
		values = append(values, reversed[i])
	}
	return values
}

func extractParams(parent *Resource, fn reflect.Type) (names []string) {
	if fn.NumIn() < 2 {
		// fn is with receiver
		// skip last as it's considered to be "nil"
		return nil
	}

	values := valueVariables(parent)
	ctx := reflect.TypeOf(new(context.Context)).Elem()
	for i := 1; i < fn.NumIn()-1; i++ {
		param := fn.In(i)
		switch param.Kind() {
		case reflect.String:
			// string -> param name
			names, values = append(names, values[0]), values[1:]
		case reflect.Interface:
			if param.Implements(ctx) {
				// ctx -> "ctx"
				names = append(names, "ctx")
			}
		case reflect.Pointer:
			// *struct -> "nil"
			names = append(names, "nil")
		}
	}
	return
}

func methodMatch(method reflect.Method, fn reflect.Type) bool {
	// assume fn is with receiver
	methodType := method.Type
	if methodType.NumIn() != fn.NumIn()-1 {
		return false
	}
	if methodType.NumOut() != fn.NumOut() {
		return false
	}
	for i := 0; i < methodType.NumIn(); i++ {
		if methodType.In(i) != fn.In(i+1) {
			return false
		}
	}
	for i := 1; i < methodType.NumOut(); i++ {
		if methodType.Out(i) != fn.Out(i) {
			if methodType.NumOut() != fn.NumOut() {
				return false
			}
		}
	}
	return true
}

func getStructName(typ reflect.Type) string {
	last := typ.String()
	parts := strings.Split(last, "[")
	switch len(parts) {
	case 1: // nop
	case 2:
		last = strings.TrimSuffix(parts[len(parts)-1], "]")
	default:
		panic("unexpected amount of [")
	}
	dotParts := strings.Split(last, ".")
	return dotParts[len(dotParts)-1]
}

func getStructPackageName(x any) string {
	varType := strings.TrimPrefix(fmt.Sprintf("%T", x), "*")
	return strings.TrimPrefix(strings.TrimSuffix(varType, path.Ext(varType)), "arm")
}

type Fetcher struct {
	Service string
	Client  string
	Method  string
	Params  []string

	ValueField string

	StructName      string
	BasicStructName string

	Import            string
	StructPackageName string

	PreResolver *Fetcher

	ExtraValues []string
}

func (f *Fetcher) IsPager() bool {
	return strings.HasPrefix(f.Method, "NewList")
}

func (f *Fetcher) IsList() bool {
	return strings.HasPrefix(f.Method, "List")
}

// Fetcher returns the fetcher to be used in templates
func (r *Resource) Fetcher() *Fetcher {
	return r.fetcher
}

func (r *Resource) propagateFetcher() {
	r.fetcher = paramsFromAny(r.Resolver).getFetcher(r.parent)
	if r.PreResolver != nil {
		r.fetcher.PreResolver = paramsFromAny(r.PreResolver).getFetcher(r.parent)
	}
}

func (r *Resource) generateFetch(dir string) error {
	if r.Resolver == nil {
		return nil
	}
	tpl, err := template.New("fetch.go.tpl").Funcs(template.FuncMap{
		"Singular":     pluralize.NewClient().Singular,
		"ToCamel":      strcase.ToCamel,
		"ToLower":      strings.ToLower,
		"ToLowerCamel": strcase.ToLowerCamel,
	}).ParseFS(templatesFS, "templates/*.go.tpl")
	if err != nil {
		return fmt.Errorf("failed to parse templates: %w", err)
	}
	tpl, err = tpl.ParseFS(codegen.TemplatesFS, "templates/*.go.tpl")
	if err != nil {
		return fmt.Errorf("failed to parse sdk template: %w", err)
	}

	var buff bytes.Buffer
	if err := tpl.Execute(&buff, r); err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}

	filePath := path.Join(dir, r.SubService+"_fetch.go")
	return util.WriteAndFormat(filePath, buff.Bytes())
}
