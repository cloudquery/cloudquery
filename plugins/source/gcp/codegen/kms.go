package codegen

import (
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
	"google.golang.org/api/cloudkms/v1"
)

var emptyString = ""

var kmsResources = []*Resource{
	{
		SubService:   "crypto_keys",
		Struct:       &cloudkms.CryptoKey{},
		ListFunction: "c.Services.Kms.Projects.Locations.KeyRings.CryptoKeys.List(r.Parent.Item.(*cloudkms.KeyRing).Name).PageToken(nextPageToken).Do()",
		Imports:      []string{"google.golang.org/api/cloudkms/v1"},
		DefaultColumns: []codegen.ColumnDefinition{
			ProjectIdColumn,
			{
				Name:     "policy",
				Type:     schema.TypeJSON,
				Resolver: "resolveKmsKeyringCryptoKeyPolicy",
			},
		},
		Multiplex:  &emptyString,
		ChildTable: true,
		SkipMock:   true,
	},
	{
		SubService: "keyrings",
		Struct:     &cloudkms.KeyRing{},
		DefaultColumns: []codegen.ColumnDefinition{
			ProjectIdColumn,
			{
				Name: "location",
				Type: schema.TypeString,
			},
		},
		SkipMock:  true,
		Relations: []string{"CryptoKeys()"},
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
