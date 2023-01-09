package main

import (
	"path"
	"reflect"
	"runtime"

	"github.com/cloudquery/plugin-sdk/mockgen"
	"github.com/fastly/go-fastly/v7/fastly"
)

// Generate the service interfaces under in client/services for use with mockgen.
func main() {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("failed to get caller information")
	}
	err := mockgen.GenerateInterfaces(
		[]any{&fastly.Client{}},
		path.Join(path.Dir(filename), "../client/services"),
		mockgen.WithIncludeFunc(include),
	)
	if err != nil {
		panic(err)
	}
}

func include(m reflect.Method) bool {
	return mockgen.MethodHasAnyPrefix(m, []string{"Get", "List", "NewGet", "NewList"})
}
