package recipes

import (
	"cloud.google.com/go/kms/apiv1/kmspb"
)

var emptyString = ""

var kmsResources = []*Resource{
	{
		SubService: "crypto_keys",
		Struct:     &kmspb.CryptoKey{},
		Multiplex:  &emptyString,
		ChildTable: true,
		SkipMock:   true,
		SkipFetch:  true,
		SkipFields: []string{"RotationSchedule"},
	},
	{
		SubService: "keyrings",
		Struct:     &kmspb.KeyRing{},
		Relations:  []string{"CryptoKeys()"},
		SkipFetch:  true,
		SkipMock:   true,
	},
}

func KmsResources() []*Resource {
	var resources []*Resource
	resources = append(resources, kmsResources...)

	for _, resource := range resources {
		resource.Service = "kms"
		resource.MockImports = []string{"cloud.google.com/go/kms/apiv1"}
		resource.ProtobufImport = "cloud.google.com/go/kms/apiv1/kmspb"
		resource.Template = "newapi_list"
		resource.MockTemplate = "newapi_list_grpc_mock"
	}

	return resources
}
