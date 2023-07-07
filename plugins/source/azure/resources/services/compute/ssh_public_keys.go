package compute

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v4"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func SSHPublicKeys() *schema.Table {
	return &schema.Table{
		Name:                 "azure_compute_ssh_public_keys",
		Resolver:             fetchSSHPublicKeys,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/compute/ssh-public-keys/list-by-subscription?tabs=HTTP#sshpublickeyresource",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_compute_ssh_public_keys", client.Namespacemicrosoft_compute),
		Transform:            transformers.TransformWithStruct(&armcompute.SSHPublicKeyResource{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
	}
}

func fetchSSHPublicKeys(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armcompute.NewSSHPublicKeysClient(cl.SubscriptionId, cl.Creds, cl.Options)
	if err != nil {
		return err
	}
	pager := svc.NewListBySubscriptionPager(nil)
	for pager.More() {
		p, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- p.Value
	}
	return nil
}
