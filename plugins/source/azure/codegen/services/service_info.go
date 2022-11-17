package services

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/cloudquery/plugin-sdk/caser"
)

type serviceInfo struct {
	Imports     []string
	Name        string
	PackageName string
	SourceFile  string
	Signatures  []string
	ExtraTypes  []*TypeAlias

	ConstructorPackageName string
	ConstructorCall        string // this will be the call required to produce the needed type
}

func getServiceInfo(constructor any) *serviceInfo {
	constructorType := reflect.TypeOf(constructor)

	if constructorType.NumOut() != 2 {
		panic("unexpected ret amount")
	}
	switch tp := reflect.New(constructorType.Out(1)).Interface().(type) {
	case *error:
	default:
		panic(fmt.Sprintf("actual type %T", tp))
	}
	clientType := constructorType.Out(0)
	if clientType.Kind() != reflect.Pointer {
		panic("unexpected ret to be pointer")
	}
	clientVal := reflect.New(clientType)

	var imports []string
	var aliases []*TypeAlias
	pkgName := packageName(clientVal.Elem().Interface())
	csr := caser.New()
	name := strings.TrimSuffix(csr.ToPascal(clientType.Elem().Name()), "Client")

	signatures := make([]string, 0, clientVal.NumMethod())
	for i := 0; i < clientType.NumMethod(); i++ {
		method := clientType.Method(i)
		if isAllowedFunc(method.Name) {
			signature, neededImports, extraTypes := getSignature(method.Func.Interface())
			signatures = append(signatures, signature)
			imports = append(imports, neededImports...)
			aliases = append(aliases, extraTypes...)
		}
	}
	sourceFile := csr.ToSnake(name)
	if len(sourceFile) == 0 {
		sourceFile = "client"
	}
	return &serviceInfo{
		Imports:                uniqPackages(imports...),
		Name:                   name,
		PackageName:            strings.TrimPrefix(pkgName, "arm"),
		Signatures:             signatures,
		SourceFile:             sourceFile,
		ExtraTypes:             aliases,
		ConstructorPackageName: pkgName,
		ConstructorCall:        getConstructorCall(constructor),
	}
}
