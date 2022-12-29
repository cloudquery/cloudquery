package recipes

import (
	appengine "cloud.google.com/go/appengine/apiv1"
	pb "cloud.google.com/go/appengine/apiv1/appenginepb"
)

func init() {
	resources := []*Resource{
		{
			SubService:          "apps",
			Struct:              &pb.Application{},
			PrimaryKeys:         []string{ProjectIdColumn.Name, "name"},
			NewFunction:         appengine.NewApplicationsClient,
			ListFunction:        (&appengine.ApplicationsClient{}).GetApplication,
			RegisterServer:      pb.RegisterApplicationsServer,
			RequestStructFields: `Parent: "apps/" + c.ProjectId,`,
			Description:         "https://cloud.google.com/appengine/docs/admin-api/reference/rest/v1/apps#Application",
			SkipFetch:           true,
			SkipMock:            true,
		},
		{
			SubService:          "services",
			Struct:              &pb.Service{},
			PrimaryKeys:         []string{ProjectIdColumn.Name, "name"},
			NewFunction:         appengine.NewServicesClient,
			RegisterServer:      pb.RegisterServicesServer,
			RequestStructFields: `Parent: "apps/" + c.ProjectId,`,
			Description:         "https://cloud.google.com/appengine/docs/admin-api/reference/rest/v1/apps.services#Service",
			Relations:           []string{"Versions()"},
			SkipMock:            true,
		},
		{
			SubService:          "versions",
			Struct:              &pb.Version{},
			PrimaryKeys:         []string{ProjectIdColumn.Name, "name"},
			NewFunction:         appengine.NewVersionsClient,
			RegisterServer:      pb.RegisterVersionsServer,
			RequestStructFields: `Parent: parent.Item.(*pb.Service).Name,`,
			Description:         "https://cloud.google.com/appengine/docs/admin-api/reference/rest/v1/apps.services.versions#Version",
			ChildTable:          true,
			Relations:           []string{"Instances()"},
			SkipMock:            true,
		},
		{
			SubService:          "instances",
			Struct:              &pb.Instance{},
			PrimaryKeys:         []string{ProjectIdColumn.Name, "name"},
			NewFunction:         appengine.NewInstancesClient,
			RegisterServer:      pb.RegisterInstancesServer,
			RequestStructFields: `Parent: parent.Item.(*pb.Version).Name,`,
			ChildTable:          true,
			Description:         "https://cloud.google.com/appengine/docs/admin-api/reference/rest/v1/apps.services.versions.instances#Instance",
			SkipMock:            true,
		},
		{
			SubService:          "authorized_certificates",
			Struct:              &pb.AuthorizedCertificate{},
			PrimaryKeys:         []string{ProjectIdColumn.Name, "name"},
			NewFunction:         appengine.NewAuthorizedCertificatesClient,
			RegisterServer:      pb.RegisterAuthorizedCertificatesServer,
			RequestStructFields: `Parent: "apps/" + c.ProjectId,`,
			Description:         "https://cloud.google.com/appengine/docs/admin-api/reference/rest/v1/apps.authorizedCertificates#AuthorizedCertificate",
		},
		{
			SubService:          "authorized_domains",
			Struct:              &pb.AuthorizedDomain{},
			PrimaryKeys:         []string{ProjectIdColumn.Name, "name"},
			NewFunction:         appengine.NewAuthorizedDomainsClient,
			RegisterServer:      pb.RegisterAuthorizedDomainsServer,
			RequestStructFields: `Parent: "apps/" + c.ProjectId,`,
			Description:         "https://cloud.google.com/appengine/docs/admin-api/reference/rest/v1/apps.authorizedDomains#AuthorizedDomain",
		},
		{
			SubService:          "domain_mappings",
			Struct:              &pb.DomainMapping{},
			PrimaryKeys:         []string{ProjectIdColumn.Name, "name"},
			NewFunction:         appengine.NewDomainMappingsClient,
			RegisterServer:      pb.RegisterDomainMappingsServer,
			RequestStructFields: `Parent: "apps/" + c.ProjectId,`,
			Description:         "https://cloud.google.com/appengine/docs/admin-api/reference/rest/v1/apps.domainMappings#DomainMapping",
		},
		{
			SubService:          "firewall_ingress_rules",
			Struct:              &pb.FirewallRule{},
			PrimaryKeys:         []string{ProjectIdColumn.Name, "name"},
			NewFunction:         appengine.NewFirewallClient,
			RegisterServer:      pb.RegisterFirewallServer,
			RequestStructFields: `Parent: "apps/" + c.ProjectId,`,
			Description:         "https://cloud.google.com/appengine/docs/admin-api/reference/rest/v1/apps.firewall.ingressRules#FirewallRule",
		},
	}

	for _, resource := range resources {
		resource.Service = "appengine"
		resource.Template = "newapi_list"
		resource.MockTemplate = "newapi_list_grpc_mock"
		resource.ProtobufImport = "cloud.google.com/go/appengine/apiv1/appenginepb"
		resource.MockImports = []string{"cloud.google.com/go/appengine/apiv1"}
	}

	Resources = append(Resources, resources...)
}
