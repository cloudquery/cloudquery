package codegen

import (
	kms "cloud.google.com/go/kms/apiv1"
	pb "google.golang.org/genproto/googleapis/cloud/kms/v1"
)

var emptyString = ""

var kmsResources = []*Resource{
	{
		SubService:          "crypto_keys",
		Struct:              &pb.CryptoKey{},
		NewFunction:         kms.NewKeyManagementClient,
		RequestStruct:       &pb.ListCryptoKeysRequest{},
		ResponseStruct:      &pb.ListCryptoKeysResponse{},
		RegisterServer:      pb.RegisterKeyManagementServiceServer,
		ListFunction:        (&pb.UnimplementedKeyManagementServiceServer{}).ListCryptoKeys,
		UnimplementedServer: &pb.UnimplementedKeyManagementServiceServer{},
		Multiplex:           &emptyString,
		ChildTable:          true,
		SkipMock:            true,
		SkipFetch:           true,
	},
	{
		SubService:          "keyrings",
		Struct:              &pb.KeyRing{},
		NewFunction:         kms.NewKeyManagementClient,
		RequestStruct:       &pb.ListKeyRingsRequest{},
		ResponseStruct:      &pb.ListKeyRingsResponse{},
		RegisterServer:      pb.RegisterKeyManagementServiceServer,
		ListFunction:        (&pb.UnimplementedKeyManagementServiceServer{}).ListKeyRings,
		UnimplementedServer: &pb.UnimplementedKeyManagementServiceServer{},
		Relations:           []string{"CryptoKeys()"},
		SkipFetch:           true,
		SkipMock:            true,
	},
}

func KmsResources() []*Resource {
	var resources []*Resource
	resources = append(resources, kmsResources...)

	for _, resource := range resources {
		resource.Service = "kms"
		resource.MockImports = []string{"cloud.google.com/go/kms/apiv1"}
		resource.ProtobufImport = "google.golang.org/genproto/googleapis/cloud/kms/v1"
		resource.Template = "newapi_list"
		resource.MockTemplate = "newapi_list_grpc_mock"
	}

	return resources
}
