package services

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/containerservice/mgmt/2021-03-01/containerservice"
	"github.com/Azure/go-autorest/autorest"
)

type ContainerServiceClient struct {
	ManagedClusters ManagedClustersClient
}

type ManagedClustersClient interface {
	List(ctx context.Context) (result containerservice.ManagedClusterListResultPage, err error)
}

func NewContainerServiceClient(subscriptionID string, auth autorest.Authorizer) ContainerServiceClient {
	m := containerservice.NewManagedClustersClient(subscriptionID)
	m.Authorizer = auth
	return ContainerServiceClient{
		ManagedClusters: m,
	}
}
