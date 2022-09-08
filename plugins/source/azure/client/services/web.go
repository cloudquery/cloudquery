//go:generate mockgen -destination=./mocks/web.go -package=mocks . WebAppsClient,WebSiteAuthSettingsClient,WebVnetConnectionsClient,WebPublishingProfilesClient
package services

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/web/mgmt/2020-12-01/web"
	"github.com/Azure/go-autorest/autorest"
)

type WebClient struct {
	Apps               WebAppsClient
	SiteAuthSettings   WebSiteAuthSettingsClient
	VnetConnections    WebVnetConnectionsClient
	PublishingProfiles WebPublishingProfilesClient
}

type WebAppsClient interface {
	List(ctx context.Context) (result web.AppCollectionPage, err error)
}

type WebPublishingProfilesClient interface {
	ListPublishingProfileXMLWithSecrets(ctx context.Context, resourceGroupName string, name string, publishingProfileOptions web.CsmPublishingProfileOptions) (result web.ReadCloser, err error)
}

type WebVnetConnectionsClient interface {
	GetVnetConnection(ctx context.Context, resourceGroupName string, name string, vnetName string) (result web.VnetInfo, err error)
}

type WebSiteAuthSettingsClient interface {
	GetAuthSettings(ctx context.Context, resourceGroupName string, name string) (result web.SiteAuthSettings, err error)
}

func NewWebClient(subscriptionId string, auth autorest.Authorizer) WebClient {
	apps := web.NewAppsClient(subscriptionId)
	apps.Authorizer = auth
	return WebClient{
		Apps:               apps,
		SiteAuthSettings:   apps,
		VnetConnections:    apps,
		PublishingProfiles: apps,
	}
}
