package sql

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
	sql "google.golang.org/api/sqladmin/v1beta4"
)

func fetchUsers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	sqlClient, err := sql.NewService(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}
	instance := parent.Item.(*sql.DatabaseInstance)
	output, err := sqlClient.Users.List(c.ProjectId, instance.Name).Do()
	if err != nil {
		return err
	}
	res <- output.Items
	return nil
}
