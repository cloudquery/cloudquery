package services

import (
	"fmt"
	"path"
	"reflect"
	"strconv"
	"strings"

	"github.com/cloudquery/plugin-sdk/caser"
)

type (
	TypeInfo struct {
		Str     string
		Imports []string
		Aliases []*TypeAlias
	}
	TypeAlias struct {
		Alias      string
		Definition string
	}
)

func getTypeInfo(typ reflect.Type) *TypeInfo {
	switch typ.Kind() {
	// Array, Chan, Map, Pointer, or Slice
	case reflect.Array:
		val := reflect.New(typ.Elem())
		info := getTypeInfo(val.Type())
		info.Str = fmt.Sprintf("[%d]", typ.Len()) + info.Str
		return info
	case reflect.Chan:
		val := reflect.New(typ.Elem())
		info := getTypeInfo(val.Type())
		info.Str = typ.ChanDir().String() + info.Str
		return info
	case reflect.Map:
		val := reflect.New(typ.Elem())
		valInfo := getTypeInfo(val.Type())
		key := reflect.New(typ.Key())
		keyInfo := getTypeInfo(key.Type())
		return &TypeInfo{
			Str:     "map[" + valInfo.Str + "]" + keyInfo.Str,
			Imports: append(keyInfo.Imports, valInfo.Imports...),
		}
	case reflect.Pointer:
		val := reflect.New(typ.Elem())
		info := getTypeInfo(val.Elem().Type())
		info.Str = "*" + info.Str
		return info
	case reflect.Slice:
		val := reflect.New(typ.Elem())
		info := getTypeInfo(val.Type())
		info.Str = "[]" + info.Str
		return info
	}

	name, imports, aliases := canonicalName(typ)
	return &TypeInfo{
		Str:     name,
		Imports: append(imports, typ.PkgPath()),
		Aliases: aliases,
	}
}

// canonicalName will check only for 1 level down, more complex cases will probably require refactoring
func canonicalName(typ reflect.Type) (name string, imports []string, aliases []*TypeAlias) {
	parts := strings.Split(typ.String(), "[")
	name = parts[0]
	if len(parts) < 2 {
		return
	}

	// Generics-infused type.
	// In this case we'll make an alias to be able to gen the mock
	csr := caser.New()
	normalizePart := func(s string) string {
		return csr.ToPascal(strings.ReplaceAll(s, ".", ""))
	}

	altImport := strings.TrimSuffix(parts[1], "]")
	if !strings.Contains(altImport, "/") {
		// built-in
		typeName := normalizePart(name) + normalizePart(parts[0])
		aliases = []*TypeAlias{{
			Alias:      typeName,
			Definition: name + "[" + parts[0] + "]",
		}}
		return typeName, []string{altImport}, aliases
	}
	// We will only handle the type itself, not its pointers here.
	parts = strings.Split(altImport, "/")

	last := parts[len(parts)-1]
	lastTypeParts := strings.Split(last, ".")
	if len(lastTypeParts) != 2 {
		panic(fmt.Sprintf("unhandled case for path %q (%q)", altImport, last))
	}

	// handle .../vXXX
	ver := lastTypeParts[0]
	if strings.HasPrefix(ver, "v") {
		_, err := strconv.ParseInt(strings.TrimPrefix(ver, "v"), 10, 32)
		if err == nil {
			// indeed, .../vXXX. import fully, but use prev part for package name
			last = parts[len(parts)-2] + "." + lastTypeParts[1]
		}
	}

	typeName := normalizePart(name) + normalizePart(last)
	aliases = []*TypeAlias{{
		Alias:      typeName,
		Definition: name + "[" + last + "]",
	}}
	return typeName, []string{path.Join(append(parts[:len(parts)-1], lastTypeParts[0])...)}, aliases
}
