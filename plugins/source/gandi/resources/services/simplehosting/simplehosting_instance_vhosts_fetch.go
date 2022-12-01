package simplehosting

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/gandi/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/go-gandi/go-gandi/simplehosting"
)

func fetchSimplehostingInstanceVhosts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	ins := parent.Item.(simplehosting.Instance)

	list, err := cl.Services.SimpleHostingClient.ListVhosts(ins.ID)
	if err != nil {
		return err
	}
	res <- list
	return nil
}
