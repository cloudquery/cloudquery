package recipes

import (
	certificatemanager "cloud.google.com/go/certificatemanager/apiv1"
	pb "cloud.google.com/go/certificatemanager/apiv1/certificatemanagerpb"
)

func init() {
	resources := []*Resource{
		{
			SubService:          "certificate_issuance_configs",
			Struct:              &pb.CertificateIssuanceConfig{},
			PrimaryKeys:         []string{ProjectIdColumn.Name, "name"},
			RequestStructFields: `Parent: "projects/" + c.ProjectId + "/locations/-",`,
			Description:         "https://cloud.google.com/certificate-manager/docs/reference/rest/v1/projects.locations.certificateIssuanceConfigs#CertificateIssuanceConfig",
		},
		{
			SubService:          "certificate_maps",
			Struct:              &pb.CertificateMap{},
			PrimaryKeys:         []string{ProjectIdColumn.Name, "name"},
			RequestStructFields: `Parent: "projects/" + c.ProjectId + "/locations/-",`,
			Description:         "https://cloud.google.com/certificate-manager/docs/reference/rest/v1/projects.locations.certificateMaps#CertificateMap",
			Relations:           []string{"CertificateMapEntries()"},
			SkipMock:            true,
		},
		{
			SubService:          "certificates",
			Struct:              &pb.Certificate{},
			PrimaryKeys:         []string{ProjectIdColumn.Name, "name"},
			RequestStructFields: `Parent: "projects/" + c.ProjectId + "/locations/-",`,
			Description:         "https://cloud.google.com/certificate-manager/docs/reference/rest/v1/projects.locations.certificates#Certificate",
		},
		{
			SubService:          "dns_authorizations",
			Struct:              &pb.DnsAuthorization{},
			PrimaryKeys:         []string{ProjectIdColumn.Name, "name"},
			RequestStructFields: `Parent: "projects/" + c.ProjectId + "/locations/-",`,
			Description:         "https://cloud.google.com/certificate-manager/docs/reference/rest/v1/projects.locations.dnsAuthorizations#DnsAuthorization",
		},
		{
			SubService:          "certificate_map_entries",
			Struct:              &pb.CertificateMapEntry{},
			PrimaryKeys:         []string{ProjectIdColumn.Name, "name"},
			RequestStructFields: `Parent: parent.Item.(*pb.CertificateMap).Name,`,
			Description:         "https://cloud.google.com/certificate-manager/docs/reference/rest/v1/projects.locations.certificateMaps.certificateMapEntries#CertificateMapEntry",
			ChildTable:          true,
			SkipMock:            true,
		},
	}

	for _, resource := range resources {
		resource.Service = "certificatemanager"
		resource.Template = "newapi_list"
		resource.MockTemplate = "newapi_list_grpc_mock"
		resource.ProtobufImport = "cloud.google.com/go/certificatemanager/apiv1/certificatemanagerpb"
		resource.MockImports = []string{"cloud.google.com/go/certificatemanager/apiv1"}
		resource.NewFunction = certificatemanager.NewClient
		resource.RegisterServer = pb.RegisterCertificateManagerServer
	}

	Resources = append(Resources, resources...)
}
