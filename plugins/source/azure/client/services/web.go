//go:generate mockgen -destination=./mocks/web.go -package=mocks . WebAppsClient,WebFunctionsClient,WebSiteAuthSettingsClient,WebSiteAuthSettingsV2Client,WebVnetConnectionsClient,WebPublishingProfilesClient
package services

import (
	"bytes"
	"context"
	"encoding/xml"

	"github.com/Azure/azure-sdk-for-go/services/web/mgmt/2020-12-01/web"
	"github.com/Azure/go-autorest/autorest"
)

type WebClient struct {
	Apps               WebAppsClient
	SiteAuthSettings   WebSiteAuthSettingsClient
	SiteAuthSettingsV2 WebSiteAuthSettingsV2Client
	VnetConnections    WebVnetConnectionsClient
	PublishingProfiles WebPublishingProfilesClient
	Functions          WebFunctionsClient
}

type WebAppsClient interface {
	List(ctx context.Context) (result web.AppCollectionPage, err error)
}

type WebPublishingProfilesClient interface {
	ListPublishingProfiles(ctx context.Context, resourceGroupName string, name string) (result PublishingProfiles, err error)
}

type WebVnetConnectionsClient interface {
	GetVnetConnection(ctx context.Context, resourceGroupName string, name string, vnetName string) (result web.VnetInfo, err error)
}

type WebSiteAuthSettingsClient interface {
	GetAuthSettings(ctx context.Context, resourceGroupName string, name string) (result web.SiteAuthSettings, err error)
}

type WebSiteAuthSettingsV2Client interface {
	GetAuthSettingsV2(ctx context.Context, resourceGroupName string, name string) (result web.SiteAuthSettingsV2, err error)
}

type WebFunctionsClient interface {
	ListFunctions(ctx context.Context, resourceGroupName string, name string) (result web.FunctionEnvelopeCollectionPage, err error)
}
type WebPublishingProfilesClientImpl struct {
	web.AppsClient
}

type PublishingProfile struct {
	PublishUrl string `xml:"publishUrl,attr"`
	UserName   string `xml:"userName,attr"`
	UserPWD    string `xml:"userPWD,attr"`
}
type PublishData struct {
	XMLName     xml.Name            `xml:"publishData"`
	PublishData []PublishingProfile `xml:"publishProfile"`
}

type PublishingProfiles []PublishingProfile

func (c WebPublishingProfilesClientImpl) ListPublishingProfiles(ctx context.Context, resourceGroupName string, siteName string) (result PublishingProfiles, err error) {
	response, err := c.ListPublishingProfileXMLWithSecrets(ctx, resourceGroupName, siteName, web.CsmPublishingProfileOptions{})
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	if _, err = buf.ReadFrom(response.Body); err != nil {
		return nil, err
	}
	var profileData PublishData
	if err = xml.Unmarshal(buf.Bytes(), &profileData); err != nil {
		return nil, err
	}

	return profileData.PublishData, nil
}

func NewWebClient(subscriptionId string, auth autorest.Authorizer) WebClient {
	apps := web.NewAppsClient(subscriptionId)
	apps.Authorizer = auth

	return WebClient{
		Apps:               apps,
		SiteAuthSettings:   apps,
		SiteAuthSettingsV2: apps,
		VnetConnections:    apps,
		PublishingProfiles: WebPublishingProfilesClientImpl{apps},
		Functions:          apps,
	}
}
