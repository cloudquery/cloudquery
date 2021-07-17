package services

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2020-10-01/resources"
	"github.com/Azure/go-autorest/autorest"
)

type ResourcesClient struct {
	Groups    GroupsClient
	Resources ResClient
}

func NewResourcesClient(subscriptionId string, auth autorest.Authorizer) ResourcesClient {
	groups := resources.NewGroupsClient(subscriptionId)
	groups.Authorizer = auth
	return ResourcesClient{
		Groups: groups,
	}
}

type GroupsClient interface {
	List(ctx context.Context, filter string, top *int32) (result resources.GroupListResultPage, err error)
}

type ResClient interface {
	List(ctx context.Context, filter string, expand string, top *int32) (result resources.ListResultPage, err error)
}
