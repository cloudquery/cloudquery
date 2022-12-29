package recipes

import (
	"github.com/fastly/go-fastly/v7/fastly"
)

func ServiceResources() []*Resource {
	resources := []*Resource{
		{
			DataStruct:  &fastly.Service{},
			Description: "https://developer.fastly.com/reference/api/services/service/",
			PKColumns:   []string{"id"},
			Relations: []string{
				"ServiceVersions()",
			},
		},
		{
			TableName:   "service_versions",
			DataStruct:  &fastly.Version{},
			Description: "https://developer.fastly.com/reference/api/services/version/",
			PKColumns:   []string{"service_id", "number"},
			Relations: []string{
				"ServiceHealthChecks()",
				"ServiceDomains()",
				"ServiceBackends()",
			},
		},
		{
			TableName:   "service_backends",
			DataStruct:  &fastly.Backend{},
			Description: "https://developer.fastly.com/reference/api/services/backend/",
			PKColumns:   []string{"service_id", "service_version", "name"},
		},
		{
			TableName:   "service_health_checks",
			DataStruct:  &fastly.HealthCheck{},
			Description: "https://developer.fastly.com/reference/api/services/healthcheck/",
			PKColumns:   []string{"service_id", "service_version", "name"},
		},
		{
			TableName:   "service_domains",
			DataStruct:  &fastly.Domain{},
			Description: "https://developer.fastly.com/reference/api/services/domain/",
			PKColumns:   []string{"service_id", "service_version", "name"},
		},
	}
	for _, r := range resources {
		r.Service = "services"
	}
	return resources
}
