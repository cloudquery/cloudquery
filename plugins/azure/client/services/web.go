package services

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/web/mgmt/2020-12-01/web"
	"github.com/Azure/go-autorest/autorest"
)

type WebClient struct {
	Apps AppsClient
}

func NewWebClient(subscriptionId string, auth autorest.Authorizer) WebClient {
	apps := web.NewAppsClient(subscriptionId)
	apps.Authorizer = auth
	return WebClient{
		Apps: apps,
	}
}

type AppsClient interface {
	List(ctx context.Context) (result web.AppCollectionPage, err error)
	ListPublishingProfileXMLWithSecrets(ctx context.Context, resourceGroupName string, name string, publishingProfileOptions web.CsmPublishingProfileOptions) (result web.ReadCloser, err error)
	GetAuthSettings(ctx context.Context, resourceGroupName string, name string) (result web.SiteAuthSettings, err error)
	GetVnetConnection(ctx context.Context, resourceGroupName string, name string, vnetName string) (result web.VnetInfo, err error)
}
