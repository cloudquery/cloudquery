package services

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/preview/resources/mgmt/2020-03-01-preview/policy"
	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2016-09-01/links"
	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2020-10-01/resources"
	"github.com/Azure/go-autorest/autorest"
)

type ResourcesClient struct {
	Groups      GroupsClient
	Resources   ResClient
	Assignments AssignmentsClient
	Links       LinksClient
}

type GroupsClient interface {
	List(ctx context.Context, filter string, top *int32) (result resources.GroupListResultPage, err error)
}

type ResClient interface {
	List(ctx context.Context, filter string, expand string, top *int32) (result resources.ListResultPage, err error)
}

type AssignmentsClient interface {
	List(ctx context.Context, subscriptionID string, filter string, top *int32) (result policy.AssignmentListResultPage, err error)
}

type LinksClient interface {
	ListAtSubscription(ctx context.Context, filter string) (result links.ResourceLinkResultPage, err error)
}

func NewResourcesClient(subscriptionId string, auth autorest.Authorizer) ResourcesClient {
	groups := resources.NewGroupsClient(subscriptionId)
	groups.Authorizer = auth
	client := resources.NewClient(subscriptionId)
	client.Authorizer = auth
	assignments := policy.NewAssignmentsClient()
	assignments.Authorizer = auth
	ls := links.NewResourceLinksClient(subscriptionId)
	ls.Authorizer = auth
	return ResourcesClient{
		Groups:      groups,
		Resources:   client,
		Assignments: assignments,
		Links:       ls,
	}
}
