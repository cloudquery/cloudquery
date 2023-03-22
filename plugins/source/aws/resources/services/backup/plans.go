package backup

import (
	"github.com/aws/aws-sdk-go-v2/service/backup"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Plans() *schema.Table {
	tableName := "aws_backup_plans"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/aws-backup/latest/devguide/API_GetBackupPlan.html`,
		Resolver:            fetchBackupPlans,
		PreResourceResolver: getPlan,
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "backup"),
		Transform:           transformers.TransformWithStruct(&backup.GetBackupPlanOutput{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("BackupPlanArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolvePlanTags,
			},
		},

		Relations: []*schema.Table{
			PlanSelections(),
		},
	}
}
