package codegen

import (
	"google.golang.org/api/cloudkms/v1"
)

var kmsResources = []*Resource{
	{
		SubService:   "crypto_keys",
		Struct:       &cloudkms.CryptoKey{},
		ListFunction: "c.Services.Kms.Projects.Locations.KeyRings.CryptoKeys.List(r.Parent.Item.(*kms.KeyRing).Name)",
		Imports:      []string{"google.golang.org/api/cloudkms/v1"},
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
