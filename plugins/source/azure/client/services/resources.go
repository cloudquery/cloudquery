//go:generate mockgen -destination=./mocks/resources.go -package=mocks . ResourcesGroupsClient,ResourcesPolicyAssignmentsClient,ResourcesLinksClient
package services

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/preview/resources/mgmt/2020-03-01-preview/policy"
	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2016-09-01/links"
	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2020-10-01/resources"
	"github.com/Azure/go-autorest/autorest"
)

type ResourcesClient struct {
	Groups            ResourcesGroupsClient
	PolicyAssignments ResourcesPolicyAssignmentsClient
	Links             ResourcesLinksClient
}

type ResourcesGroupsClient interface {
	List(ctx context.Context, filter string, top *int32) (result resources.GroupListResultPage, err error)
}

type ResourcesPolicyAssignmentsClient interface {
	List(ctx context.Context, subscriptionID string, filter string, top *int32) (result policy.AssignmentListResultPage, err error)
}

type ResourcesLinksClient interface {
	ListAtSubscription(ctx context.Context, filter string) (result links.ResourceLinkResultPage, err error)
}

func NewResourcesClient(subscriptionId string, auth autorest.Authorizer) ResourcesClient {
	groups := resources.NewGroupsClient(subscriptionId)
	groups.Authorizer = auth
	assignments := policy.NewAssignmentsClient()
	assignments.Authorizer = auth
	ls := links.NewResourceLinksClient(subscriptionId)
	ls.Authorizer = auth
	return ResourcesClient{
		Groups:            groups,
		PolicyAssignments: assignments,
		Links:             ls,
	}
}
