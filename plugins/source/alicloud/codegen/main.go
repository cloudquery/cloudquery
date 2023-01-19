package main

import (
	"path"
	"reflect"
	"runtime"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/bssopenapi"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/cloudquery/codegen/interfaces"
	"github.com/thoas/go-funk"
)

var clients = []any{
	&oss.Client{},
	&bssopenapi.Client{},
	&ecs.Client{},
}

// Generate the service interfaces under in client/services for use with mockgen.
func main() {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("failed to get caller information")
	}
	err := interfaces.Generate(
		clients,
		path.Join(path.Dir(filename), "../client/services"),
		interfaces.WithIncludeFunc(include),
		interfaces.WithExtraImports(extraImports),
	)
	if err != nil {
		panic(err)
	}
}

func include(m reflect.Method) bool {
	// these functions require extra imports, so skipping them until they are needed
	exclude := []string{
		"GetConfig",
		"GetConnectTimeout",
		"GetLogger",
		"GetReadTimeout",
		"GetSigner",
		"GetTracerRootSpan",
	}
	if funk.ContainsString(exclude, m.Name) {
		return false
	}
	return interfaces.MethodHasAnyPrefix(m, []string{"List", "Get", "Describe", "Query"})
}

func extraImports(_ reflect.Method) []string {
	return []string{}
}
