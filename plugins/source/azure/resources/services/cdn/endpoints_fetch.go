// Code generated by codegen; DO NOT EDIT.

package cdn

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	cdn "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchEndpoints(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Cdn

	profile := parent.Item.(*cdn.Profile)
	id, err := arm.ParseResourceID(*profile.ID)
	if err != nil {
		return err
	}

	pager := svc.EndpointsClient.NewListByProfilePager(id.ResourceGroupName, *profile.Name, nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- page.Value
	}

	return nil
}
