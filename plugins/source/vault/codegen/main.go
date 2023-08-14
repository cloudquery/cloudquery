package main

import (
	"github.com/cloudquery/codegen/interfaces"
	_ "github.com/cloudquery/codegen/interfaces"
	"github.com/hashicorp/vault/api"
	"log"
	"path"
	"reflect"
	"runtime"
)

func main() {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("failed to get caller information")
	}
	client := api.Client{}
	err := interfaces.Generate(
		[]any{client.Sys()},
		path.Join(path.Dir(filename), "../client/services"),
		interfaces.WithIncludeFunc(include),
		interfaces.WithExtraImports(extraImports),
	)
	if err != nil {
		log.Fatal(err)
	}
}

func include(m reflect.Method) bool {
	return interfaces.MethodHasAnyPrefix(m, []string{"List", "Get", "Read"})
}

func extraImports(_ reflect.Method) []string {
	return []string{"context"}
}
