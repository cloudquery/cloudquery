package services

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2020-10-01/resources"
	"github.com/Azure/go-autorest/autorest"
)

type ResourcesClient struct {
	Groups GroupsClient
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
