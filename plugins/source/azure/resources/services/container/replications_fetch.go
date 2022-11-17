// Code generated by codegen; DO NOT EDIT.

package container

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	containerregistry "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistry"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchReplications(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Containerregistry

	registry := parent.Item.(*containerregistry.Registry)
	id, err := arm.ParseResourceID(*registry.ID)
	if err != nil {
		return err
	}

	pager := svc.ReplicationsClient.NewListPager(id.ResourceGroupName, *registry.Name, nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- page.Value
	}

	return nil
}
