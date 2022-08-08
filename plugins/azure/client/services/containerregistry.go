//go:generate mockgen -destination=./mocks/containerregistry.go -package=mocks . ContainerRegistriesClient,ContainerReplicationsClient
package services

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/containerregistry/mgmt/2019-05-01/containerregistry"
	"github.com/Azure/go-autorest/autorest"
)

type ContainerRegistryClient struct {
	Registries   ContainerRegistriesClient
	Replications ContainerReplicationsClient
}

type ContainerRegistriesClient interface {
	List(ctx context.Context) (result containerregistry.RegistryListResultPage, err error)
}

type ContainerReplicationsClient interface {
	List(ctx context.Context, resourceGroupName string, registryName string) (result containerregistry.ReplicationListResultPage, err error)
}

func NewContainerRegistryClient(subscriptionID string, auth autorest.Authorizer) ContainerRegistryClient {
	reg := containerregistry.NewRegistriesClient(subscriptionID)
	reg.Authorizer = auth

	rep := containerregistry.NewReplicationsClient(subscriptionID)
	rep.Authorizer = auth
	return ContainerRegistryClient{
		Registries:   reg,
		Replications: rep,
	}
}
