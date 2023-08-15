package cdn

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func EdgeNodes() *schema.Table {
	return &schema.Table{
		Name:                 "azure_cdn_edge_nodes",
		Resolver:             fetchEdgeNodes,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/cdn/edge-nodes/list?tabs=HTTP#edgenode",
		Transform:            transformers.TransformWithStruct(&armcdn.EdgeNode{}, transformers.WithPrimaryKeys("ID")),
	}
}

func fetchEdgeNodes(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armcdn.NewEdgeNodesClient(cl.Creds, cl.Options)
	if err != nil {
		return err
	}
	pager := svc.NewListPager(nil)
	for pager.More() {
		p, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- p.Value
	}
	return nil
}
