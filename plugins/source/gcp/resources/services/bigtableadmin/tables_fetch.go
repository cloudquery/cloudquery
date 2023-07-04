package bigtableadmin

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugins/source/gcp/client"

	"cloud.google.com/go/bigtable"
)

func fetchTables(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	instance := parent.Item.(*bigtable.InstanceInfo)
	gcpClient, err := bigtable.NewAdminClient(ctx, c.ProjectId, instance.Name, c.ClientOptions...)
	if err != nil {
		return err
	}
	resp, err := gcpClient.Tables(ctx)
	if err != nil {
		return err
	}

	res <- resp
	return nil
}

type tableInfoWithName struct {
	*bigtable.TableInfo
	Name string
}

func getTableInfo(ctx context.Context, meta schema.ClientMeta, r *schema.Resource) error {
	c := meta.(*client.Client)
	instance := r.Parent.Item.(*bigtable.InstanceInfo)
	table := r.Item.(string)
	gcpClient, err := bigtable.NewAdminClient(ctx, c.ProjectId, instance.Name, c.ClientOptions...)
	if err != nil {
		return err
	}
	tableInfo, err := gcpClient.TableInfo(ctx, table)
	if err != nil {
		return err
	}

	r.SetItem(tableInfoWithName{TableInfo: tableInfo, Name: table})
	return nil
}
