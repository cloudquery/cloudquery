package services

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
)

// Adapted from https://stackoverflow.com/a/54129236
func getSignature(f any) (str string, imports []string, extraTypes []*TypeAlias) {
	t := reflect.TypeOf(f)
	if t.Kind() != reflect.Func {
		panic("<not a function>")
	}

	buf := new(strings.Builder)
	buf.WriteString(getFuncName(f) + "(")
	// don't forget that the arg[0] is the type itself
	for i := 1; i < t.NumIn(); i++ {
		typInfo := getTypeInfo(t.In(i))
		imports = append(imports, typInfo.Imports...)
		extraTypes = append(extraTypes, typInfo.Aliases...)

		if i > 1 {
			buf.WriteString(", ")
		}
		if t.IsVariadic() && i == t.NumIn()-1 {
			buf.WriteString("..." + strings.TrimPrefix(typInfo.Str, "[]"))
		} else {
			buf.WriteString(typInfo.Str)
		}
	}
	buf.WriteString(")")
	if numOut := t.NumOut(); numOut > 0 {
		if numOut > 1 {
			buf.WriteString(" (")
		} else {
			buf.WriteString(" ")
		}
		for i := 0; i < t.NumOut(); i++ {
			typInfo := getTypeInfo(t.Out(i))
			imports = append(imports, typInfo.Imports...)
			extraTypes = append(extraTypes, typInfo.Aliases...)

			if i > 0 {
				buf.WriteString(", ")
			}
			buf.WriteString(typInfo.Str)
		}
		if numOut > 1 {
			buf.WriteString(")")
		}
	}

	return buf.String(), imports, extraTypes
}

var allowedFuncPrefixes = []string{"Get", "List", "NewList"}

func isAllowedFunc(sig string) bool {
	for _, pfx := range allowedFuncPrefixes {
		if strings.HasPrefix(sig, pfx) {
			return true
		}
	}
	return false
}

func getConstructorCall(constructor any) string {
	fn := reflect.TypeOf(constructor)
	// 3 possible params
	//subscriptionID string
	//credentials    azcore.TokenCredential
	//options        *arm.ClientOptions
	var params []string
	for i := 0; i < fn.NumIn(); i++ {
		switch tp := reflect.New(fn.In(i)).Interface().(type) {
		case *string:
			params = append(params, "subscriptionID")
		case *azcore.TokenCredential:
			params = append(params, "credentials")
		case **arm.ClientOptions:
			params = append(params, "options")
		default:
			panic(fmt.Sprintf("unsupported type %T", tp))
		}
	}

	return getFuncName(constructor) + "(" + strings.Join(params, ", ") + ")"
}

func getFuncName(f any) string {
	parts := strings.Split(runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name(), ".")
	return parts[len(parts)-1]
}
