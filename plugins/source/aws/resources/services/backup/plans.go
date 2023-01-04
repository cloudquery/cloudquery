package backup

import (
	"github.com/aws/aws-sdk-go-v2/service/backup"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Plans() *schema.Table {
	return &schema.Table{
		Name:                "aws_backup_plans",
		Description:         `https://docs.aws.amazon.com/aws-backup/latest/devguide/API_GetBackupPlan.html`,
		Resolver:            fetchBackupPlans,
		PreResourceResolver: getPlan,
		Multiplex:           client.ServiceAccountRegionMultiplexer("backup"),
		Transform:           transformers.TransformWithStruct(&backup.GetBackupPlanOutput{}),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
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
