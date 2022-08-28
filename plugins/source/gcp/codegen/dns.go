package codegen

import (
	"fmt"

	"github.com/iancoleman/strcase"
	"google.golang.org/api/dns/v1"
)

var dnsResources = []*Resource{
	{
		SubService: "policies",
		Struct:     &dns.Policy{},
	},
	{
		SubService: "managed_zones",
		Struct:     &dns.ManagedZone{},
	},
}

func DnsResources() []*Resource {
	var resources []*Resource
	resources = append(resources, dnsResources...)

	for _, resource := range resources {
		resource.Service = "dns"
		resource.MockImports = []string{"google.golang.org/api/dns/v1"}
		resource.ListFunction = fmt.Sprintf(`c.Services.Dns.%s.List(c.ProjectId).PageToken(nextPageToken).Do()`, strcase.ToCamel(resource.SubService))
		resource.Template = "resource_list"
		resource.OutputField = strcase.ToCamel(resource.SubService)
	}

	return resources
}
