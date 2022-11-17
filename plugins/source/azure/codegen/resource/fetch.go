package resource

import (
	"bytes"
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

func (f *FuncParams) getFetcher() *Fetcher {
	if f == nil {
		return nil
	}
	fetcher := &Fetcher{
		Params:          f.Params,
		BasicStructName: f.basicStructName(),
		ExtraValues:     f.ExtraValues,
	}

	refT := reflect.TypeOf(f.Func)
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
	r.fetcher = r.Resolver.getFetcher()
	if r.PreResolver != nil {
		r.fetcher.PreResolver = r.PreResolver.getFetcher()
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
