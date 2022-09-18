// Auto generated code - DO NOT EDIT.

package web

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"

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
				Name:     "web_app_id",
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

func fetchWebPublishingProfiles(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().Web.PublishingProfiles

	site := parent.Item.(web.Site)
	response, err := svc.ListPublishingProfiles(ctx, *site.ResourceGroup, *site.Name)
	if err != nil {
		return err
	}

	res <- response
	return nil
}
