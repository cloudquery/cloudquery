package recipes

import (
	pbv2 "cloud.google.com/go/run/apiv2/runpb"
	runv1 "google.golang.org/api/run/v1"
)

func init() {
	resources := []*Resource{
		{
			SubService:  "locations",
			Struct:      &runv1.Location{},
			SkipFetch:   true,
			SkipMock:    true,
			Description: "https://cloud.google.com/run/docs/reference/rest/v1/projects.locations#Location",
			Relations:   []string{"Services()"},
		},
		{
			SubService:  "services",
			Struct:      &pbv2.Service{},
			Description: "https://cloud.google.com/run/docs/reference/rest/v2/projects.locations.services#Service",
			ChildTable:  true,
			SkipFetch:   true,
			SkipMock:    true,
		},
	}

	for _, resource := range resources {
		resource.Service = "run"
		resource.Template = "newapi_list"
		resource.MockTemplate = "newapi_list_grpc_mock"
	}

	Resources = append(Resources, resources...)
}
