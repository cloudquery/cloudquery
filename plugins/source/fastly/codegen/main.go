package main

import (
	"path"
	"reflect"
	"runtime"

	"github.com/cloudquery/codegen/interfaces"
	"github.com/fastly/go-fastly/v7/fastly"
)

// Generate the service interfaces under in client/services for use with mockgen.
func main() {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("failed to get caller information")
	}
	err := interfaces.Generate(
		[]any{&fastly.Client{}},
		path.Join(path.Dir(filename), "../client/services"),
		interfaces.WithIncludeFunc(include),
		interfaces.WithExtraImports(extraImports),
	)
	if err != nil {
		panic(err)
	}
}

func include(m reflect.Method) bool {
	return interfaces.MethodHasAnyPrefix(m, []string{"Get", "List", "NewGet", "NewList"})
}

func extraImports(m reflect.Method) []string {
	return []string{"net/http"}
}
