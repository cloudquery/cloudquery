package codegen

import (
	"google.golang.org/api/cloudkms/v1"
)

var kmsResources = []*Resource{
	{
		SubService: "keyrings",
		Struct:     &cloudkms.KeyRing{},
	},
}

func KmsResources() []*Resource {
	var resources []*Resource
	resources = append(resources, kmsResources...)

	for _, resource := range resources {
		resource.Service = "kms"
		resource.Template = "resource_list"
	}

	return resources
}
