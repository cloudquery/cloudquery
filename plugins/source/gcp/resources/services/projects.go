package services

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
)

func Projects() *schema.Table {
	return &schema.Table{
		Name: "gcp_projects",
		Description: "This table contains the list of all project_id's synced by cloudquery. " +
			"It may contain projects missing from `gcp_resourcemanager_projects` (i.e. projects where the `resourcemanager` API is not enabled)",
		Resolver:  noopTableResolver,
		Multiplex: client.ProjectMultiplex,
		Columns: []schema.Column{
			{
				Name:       "project_id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   client.ResolveProject,
				PrimaryKey: true,
			},
		},
	}
}

func noopTableResolver(ctx context.Context, meta schema.ClientMeta, r *schema.Resource, res chan<- any) error {
	res <- struct{}{}

	return nil
}
