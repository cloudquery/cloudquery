package recipes

import (
	containeranalysis "cloud.google.com/go/containeranalysis/apiv1beta1"
	grafeaspb "cloud.google.com/go/containeranalysis/apiv1beta1/grafeas/grafeaspb"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func init(){
	resources := []*Resource{
		{
			SubService: "occurrences",
			Struct:     &grafeaspb.Occurrence{},
			SkipFields: []string{"Name"},
			ExtraColumns: []codegen.ColumnDefinition{
				ProjectIdColumn,
				{
					Name:     "name",
					Type:     schema.TypeString,
					Resolver: `schema.PathResolver("Name")`,
					Options:  schema.ColumnCreationOptions{PrimaryKey: true},
				},
			},
			Template:            "newapi_list",
			ListFunction:        (&containeranalysis.GrafeasV1Beta1Client{}).ListOccurrences,
			RequestStruct:       &grafeaspb.ListOccurrencesRequest{},
			ResponseStruct:      &grafeaspb.ListOccurrencesResponse{},
			RequestStructFields: `Parent: "projects/" + c.ProjectId,`,
		},
	}

	for _, resource := range resources {
		resource.Service = "containeranalysis"

		resource.MockImports = []string{"cloud.google.com/go/containeranalysis/apiv1beta1"}
		resource.ProtobufImport = "cloud.google.com/go/containeranalysis/apiv1beta1/grafeas/grafeaspb"
		resource.NewFunction = containeranalysis.NewGrafeasV1Beta1Client

		resource.MockTemplate = "newapi_list_grpc_mock"
		resource.RegisterServer = grafeaspb.RegisterGrafeasV1Beta1Server
		resource.UnimplementedServer = &grafeaspb.UnimplementedGrafeasV1Beta1Server{}
	}

	Resources = append(Resources, resources...)
}
