// Auto generated code - DO NOT EDIT.

package web

import (
	"bytes"
	"context"
	"encoding/xml"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/pkg/errors"

	"github.com/Azure/azure-sdk-for-go/services/web/mgmt/2020-12-01/web"
)

func publishingProfiles() *schema.Table {
	return &schema.Table{
		Name:     "azure_web_publishing_profiles",
		Resolver: fetchWebPublishingProfiles,
		Columns: []schema.Column{
			{
				Name:     "subscription_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAzureSubscription,
			},
			{
				Name:     "cq_id_parent",
				Type:     schema.TypeUUID,
				Resolver: schema.ParentIDResolver,
			},
			{
				Name:     "publish_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PublishUrl"),
			},
			{
				Name:     "user_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("UserName"),
			},
			{
				Name:     "user_pwd",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("UserPWD"),
			},
		},
	}
}

type PublishProfile struct {
	PublishUrl string `xml:"publishUrl,attr"`
	UserName   string `xml:"userName,attr"`
	UserPWD    string `xml:"userPWD,attr"`
}

type publishData struct {
	XMLName     xml.Name         `xml:"publishUrl,attr"`
	PublishData []PublishProfile `xml:"PublishProfile"`
}

func fetchWebPublishingProfiles(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().Web.PublishingProfiles

	site := parent.Item.(web.Site)
	response, err := svc.ListPublishingProfileXMLWithSecrets(ctx, *site.ResourceGroup, *site.Name, web.CsmPublishingProfileOptions{})
	if err != nil {
		return errors.WithStack(err)
	}

	buf := new(bytes.Buffer)
	if _, err = buf.ReadFrom(response.Body); err != nil {
		return errors.WithStack(err)
	}
	var profileData publishData
	if err = xml.Unmarshal(buf.Bytes(), &profileData); err != nil {
		return errors.WithStack(err)
	}

	res <- profileData.PublishData
	return nil
}
