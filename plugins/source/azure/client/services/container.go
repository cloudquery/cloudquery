//go:generate mockgen -destination=./mocks/containerregistry.go -package=mocks . ContainerRegistriesClient,ContainerReplicationsClient,ContainerManagedClustersClient
package services

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/containerregistry/mgmt/2019-05-01/containerregistry"
	"github.com/Azure/azure-sdk-for-go/services/containerservice/mgmt/2021-03-01/containerservice"
	"github.com/Azure/go-autorest/autorest"
)

type ContainerClient struct {
	Registries      ContainerRegistriesClient
	Replications    ContainerReplicationsClient
	ManagedClusters ContainerManagedClustersClient
}

type ContainerRegistriesClient interface {
	List(ctx context.Context) (result containerregistry.RegistryListResultPage, err error)
}

type ContainerReplicationsClient interface {
	List(ctx context.Context, resourceGroupName string, registryName string) (result containerregistry.ReplicationListResultPage, err error)
}

type ContainerManagedClustersClient interface {
	List(ctx context.Context) (result containerservice.ManagedClusterListResultPage, err error)
}

func NewContainerClient(subscriptionID string, auth autorest.Authorizer) ContainerClient {
	reg := containerregistry.NewRegistriesClient(subscriptionID)
	reg.Authorizer = auth

	rep := containerregistry.NewReplicationsClient(subscriptionID)
	rep.Authorizer = auth

	m := containerservice.NewManagedClustersClient(subscriptionID)
	m.Authorizer = auth

	return ContainerClient{
		Registries:      reg,
		Replications:    rep,
		ManagedClusters: m,
	}
}
