// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/attestation/armattestation"

func Armattestation() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armattestation.NewOperationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/attestation/armattestation",
		},
		{
			NewFunc: armattestation.NewProvidersClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/attestation/armattestation",
		},
		{
			NewFunc: armattestation.NewPrivateEndpointConnectionsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/attestation/armattestation",
		},
	}
	return resources
}

func init() {
	Resources = append(Resources, Armattestation())
}