package armcdn

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchEndpoints(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().ArmcdnEndpoints
	p := parent.Item.(*armcdn.Profile)
	resource, err := client.ParseResourceID(*p.ID)
	if err != nil {
		return err
	}
	pager := svc.NewListByProfilePager(resource.ResourceGroup, resource.ResourceName, nil)
	for pager.More() {
		p, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- p.Value
	}
	return nil
}
