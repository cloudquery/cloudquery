package backup

import (
	"github.com/aws/aws-sdk-go-v2/service/backup"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func PlanSelections() *schema.Table {
	tableName := "aws_backup_plan_selections"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/aws-backup/latest/devguide/API_GetBackupSelection.html`,
		Resolver:    fetchBackupPlanSelections,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "backup"),
		Transform:   client.TransformWithStruct(&backup.GetBackupSelectionOutput{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "plan_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
			},
		},
	}
}
