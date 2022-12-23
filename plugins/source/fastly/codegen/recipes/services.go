package recipes

import "github.com/fastly/go-fastly/v7/fastly"

func ServiceResources() []*Resource {
	resources := []*Resource{
		{
			DataStruct:  &fastly.Service{},
			Description: "https://developer.fastly.com/reference/api/services/service/",
			PKColumns:   []string{"id"},
			SkipFields:  []string{},
			Relations:   []string{},
		},
	}
	for _, r := range resources {
		r.Service = "services"
	}
	return resources
}
