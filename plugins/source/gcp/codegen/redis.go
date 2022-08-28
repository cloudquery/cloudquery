package codegen

import (
	"fmt"

	"github.com/iancoleman/strcase"
	"google.golang.org/api/redis/v1"
)

var redisResources = []*Resource{
	{
		SubService: "instances",
		Struct:     &redis.Instance{},
	},
}

func RedisResources() []*Resource {
	var resources []*Resource
	resources = append(resources, redisResources...)

	for _, resource := range resources {
		resource.Service = "redis"
		resource.Template = "resource_list"
		resource.ListFunction = fmt.Sprintf(`c.Services.Redis.Projects.Locations.%s.List("projects/" + c.ProjectId + "/locations/-").PageToken(nextPageToken).Do()`, strcase.ToCamel(resource.SubService))
		resource.OutputField = strcase.ToCamel(resource.SubService)
	}

	return resources
}
