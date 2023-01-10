package costmanagement

import (
	"context"
	"encoding/json"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/costmanagement/armcostmanagement"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchViewQueries(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	item := parent.Item.(*armcostmanagement.View)
	if item.Properties == nil {
		return nil
	}

	svc, err := armcostmanagement.NewQueryClient(cl.Creds, cl.Options)
	if err != nil {
		return err
	}

	b, err := json.Marshal(item.Properties.Query)
	if err != nil {
		return err
	}
	var qd armcostmanagement.QueryDefinition
	if err := json.Unmarshal(b, &qd); err != nil {
		return err
	}

	data, err := svc.Usage(ctx, *item.Properties.Scope, qd, nil)
	if err != nil {
		return err
	}

	res <- data
	return nil
}
