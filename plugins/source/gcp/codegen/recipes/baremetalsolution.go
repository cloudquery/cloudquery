package recipes

import (
	baremetalsolution "cloud.google.com/go/baremetalsolution/apiv2"
	pb "cloud.google.com/go/baremetalsolution/apiv2/baremetalsolutionpb"
)

func init() {
	resources := []*Resource{
		{
			SubService:          "instances",
			Struct:              &pb.Instance{},
			PrimaryKeys:         []string{ProjectIdColumn.Name, "name"},
			RequestStructFields: `Parent: "projects/" + c.ProjectId + "/locations/-",`,
			Description:         "https://cloud.google.com/bare-metal/docs/reference/rest/v2/projects.locations.instances#Instance",
		},
		{
			SubService:          "networks",
			Struct:              &pb.Network{},
			PrimaryKeys:         []string{ProjectIdColumn.Name, "name"},
			RequestStructFields: `Parent: "projects/" + c.ProjectId + "/locations/-",`,
			Description:         "https://cloud.google.com/bare-metal/docs/reference/rest/v2/projects.locations.networks#Network",
		},
		{
			SubService:          "nfs_shares",
			Struct:              &pb.NfsShare{},
			PrimaryKeys:         []string{ProjectIdColumn.Name, "name"},
			RequestStructFields: `Parent: "projects/" + c.ProjectId + "/locations/-",`,
			Description:         "https://cloud.google.com/bare-metal/docs/reference/rest/v2/projects.locations.nfsShares#NfsShare",
		},
		{
			SubService:          "volumes",
			Struct:              &pb.Volume{},
			PrimaryKeys:         []string{ProjectIdColumn.Name, "name"},
			RequestStructFields: `Parent: "projects/" + c.ProjectId + "/locations/-",`,
			Description:         "https://cloud.google.com/bare-metal/docs/reference/rest/v2/projects.locations.volumes#Volume",
			Relations:           []string{"VolumeLuns()"},
			SkipMock:            true,
		},
		{
			SubService:          "volume_luns",
			Struct:              &pb.Lun{},
			PrimaryKeys:         []string{ProjectIdColumn.Name, "name"},
			RequestStructFields: `Parent: parent.Item.(*pb.Volume).Name,`,
			Description:         "https://cloud.google.com/bare-metal/docs/reference/rest/v2/projects.locations.volumes.luns#Lun",
			ChildTable:          true,
			SkipMock:            true,
		},
	}

	for _, resource := range resources {
		resource.Service = "baremetalsolution"
		resource.Template = "newapi_list"
		resource.MockTemplate = "newapi_list_grpc_mock"
		resource.ProtobufImport = "cloud.google.com/go/baremetalsolution/apiv2/baremetalsolutionpb"
		resource.MockImports = []string{"cloud.google.com/go/baremetalsolution/apiv2"}
		resource.NewFunction = baremetalsolution.NewClient
		resource.RegisterServer = pb.RegisterBareMetalSolutionServer
	}

	Resources = append(Resources, resources...)
}
