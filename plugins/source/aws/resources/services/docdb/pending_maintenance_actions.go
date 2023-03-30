package docdb

import (
	"github.com/aws/aws-sdk-go-v2/service/docdb/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func PendingMaintenanceActions() *schema.Table {
	tableName := "aws_docdb_pending_maintenance_actions"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/documentdb/latest/developerguide/API_PendingMaintenanceAction.html`,
		Resolver:    fetchDocdbPendingMaintenanceActions,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "docdb"),
		Transform:   client.TransformWithStruct(&types.ResourcePendingMaintenanceActions{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
		},
	}
}
