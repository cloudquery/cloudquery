package recipes

import (
	containeranalysis "cloud.google.com/go/containeranalysis/apiv1beta1"
	grafeaspb "cloud.google.com/go/containeranalysis/apiv1beta1/grafeas/grafeaspb"
)

func init() {
	resources := []*Resource{
		{
			SubService:          "occurrences",
			Struct:              &grafeaspb.Occurrence{},
			PrimaryKeys:         []string{"name"},
			Template:            "newapi_list",
			ListFunction:        (&containeranalysis.GrafeasV1Beta1Client{}).ListOccurrences,
			RequestStructFields: `Parent: "projects/" + c.ProjectId,`,
			Description:         "https://cloud.google.com/container-analysis/docs/reference/rest/v1beta1/projects.occurrences#Occurrence",
		},
	}

	for _, resource := range resources {
		resource.Service = "containeranalysis"

		resource.MockImports = []string{"cloud.google.com/go/containeranalysis/apiv1beta1"}
		resource.ProtobufImport = "cloud.google.com/go/containeranalysis/apiv1beta1/grafeas/grafeaspb"
		resource.NewFunction = containeranalysis.NewGrafeasV1Beta1Client

		resource.MockTemplate = "newapi_list_grpc_mock"
		resource.RegisterServer = grafeaspb.RegisterGrafeasV1Beta1Server
	}

	Resources = append(Resources, resources...)
}
