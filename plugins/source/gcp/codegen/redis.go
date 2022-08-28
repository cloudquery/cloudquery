package codegen

import (
	"reflect"

	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/iancoleman/strcase"
	"google.golang.org/api/redis/v1"
)

var redisResources = []Resource{
	{
		SubService: "instances",
		Struct:     &redis.Instance{},
	},
}

func RedisResources() []Resource {
	var resources []Resource
	resources = append(resources, redisResources...)

	for i := range resources {
		resources[i].Service = "redis"
		resources[i].DefaultColumns = []codegen.ColumnDefinition{ProjectIdColumn}
		resources[i].StructName = reflect.TypeOf(resources[i].Struct).Elem().Name()
		if resources[i].Template == "" {
			resources[i].Template = "resource_list"
		}
		if resources[i].SkipFields == nil {
			resources[i].SkipFields = []string{"ServerResponse", "NullFields", "ForceSendFields"}
		}
		resources[i].MockImports = []string{"google.golang.org/api/redis/v1"}
		if resources[i].MockListStruct == "" {
			resources[i].MockListStruct = strcase.ToCamel(resources[i].StructName)
		}
	}

	return resources
}
