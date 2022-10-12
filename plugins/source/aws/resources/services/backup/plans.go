// Code generated by codegen; DO NOT EDIT.

package backup

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Plans() *schema.Table {
	return &schema.Table{
		Name:                "aws_backup_plans",
		Description:         "https://docs.aws.amazon.com/aws-backup/latest/devguide/API_GetBackupPlan.html",
		Resolver:            fetchBackupPlans,
		PreResourceResolver: getPlan,
		Multiplex:           client.ServiceAccountRegionMultiplexer("backup"),
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
			{
				Name:     "advanced_backup_settings",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("AdvancedBackupSettings"),
			},
			{
				Name:     "backup_plan",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("BackupPlan"),
			},
			{
				Name:     "backup_plan_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("BackupPlanId"),
			},
			{
				Name:     "creation_date",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreationDate"),
			},
			{
				Name:     "creator_request_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CreatorRequestId"),
			},
			{
				Name:     "deletion_date",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("DeletionDate"),
			},
			{
				Name:     "last_execution_date",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("LastExecutionDate"),
			},
			{
				Name:     "version_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("VersionId"),
			},
			{
				Name:     "result_metadata",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ResultMetadata"),
			},
		},

		Relations: []*schema.Table{
			PlanSelections(),
		},
	}
}
