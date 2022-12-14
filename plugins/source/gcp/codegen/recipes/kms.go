package recipes

import (
	"cloud.google.com/go/kms/apiv1/kmspb"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

var emptyString = ""

func init() {
	resources := []*Resource{
		{
			SubService: "crypto_keys",
			Struct:     &kmspb.CryptoKey{},
			Multiplex:  &emptyString,
			ChildTable: true,
			SkipMock:   true,
			SkipFetch:  true,
			ExtraColumns: codegen.ColumnDefinitions{
				{
					Name:     "rotation_period",
					Type:     schema.TypeInt,
					Resolver: "resolveRotationPeriod",
				},
			},
			Description: "https://cloud.google.com/kms/docs/reference/rest/v1/projects.locations.keyRings.cryptoKeys#CryptoKey",
		},
		{
			SubService:  "keyrings",
			Struct:      &kmspb.KeyRing{},
			Relations:   []string{"CryptoKeys()"},
			SkipFetch:   true,
			SkipMock:    true,
			Description: "https://cloud.google.com/kms/docs/reference/rest/v1/projects.locations.keyRings#KeyRing",
		},
	}

	for _, resource := range resources {
		resource.Service = "kms"
		resource.MockImports = []string{"cloud.google.com/go/kms/apiv1"}
		resource.ProtobufImport = "cloud.google.com/go/kms/apiv1/kmspb"
		resource.Template = "newapi_list"
		resource.MockTemplate = "newapi_list_grpc_mock"
		resource.ServiceDNS = "cloudkms.googleapis.com"
	}

	Resources = append(Resources, resources...)
}
